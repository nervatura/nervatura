package mcp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	"golang.org/x/time/rate"
)

func NewMCPServer(config cu.IM) (server *mcp.Server) {
	opts := &mcp.ServerOptions{
		Instructions: "Nervatura MCP Server",
		GetSessionID: cu.GetComponentID,
	}
	server = mcp.NewServer(&mcp.Implementation{
		Name: "nervatura", Title: "Nervatura MCP Server",
		Version: cu.ToString(config["version"], "0.0.0"),
	}, opts)

	// Add tools that exercise different features of the protocol.
	mcp.AddTool(server, &customerUpdateTool, modelUpdate)
	mcp.AddTool(server, &customerQueryTool, modelQuery)
	mcp.AddTool(server, &productUpdateTool, modelUpdate)
	mcp.AddTool(server, &productQueryTool, modelQuery)

	mcp.AddTool(server, &reportDataCodeTool, reportDataCode)

	mcp.AddTool(server, &queryCodeTool, queryCode)
	mcp.AddTool(server, &queryElicitingTool, queryEliciting)
	mcp.AddTool(server, &queryParametersTool, queryParameters)
	mcp.AddTool(server, &deleteCodeTool, deleteCode)

	server.AddResource(&ntrCustomerEnResource, templateResource)

	if prompts, ok := config["prompts"].(map[string]PromptData); ok {
		for _, prompt := range prompts {
			server.AddPrompt(&mcp.Prompt{
				Name:        prompt.Name,
				Title:       prompt.Title,
				Description: prompt.Description,
				Arguments:   prompt.Arguments,
			}, promptHandler)
		}
	}

	server.AddReceivingMiddleware(receivingHandler)
	server.AddSendingMiddleware(sendingHandler)
	server.AddReceivingMiddleware(globalRateLimiterMiddleware(rate.NewLimiter(rate.Every(time.Second/5), 10)))
	server.AddReceivingMiddleware(perSessionRateLimiterMiddleware(rate.Every(time.Second/5), 10))
	server.AddReceivingMiddleware(perMethodRateLimiterMiddleware(map[string]*rate.Limiter{
		//"tools/call":  rate.NewLimiter(rate.Every(time.Second), 5),  // once a second with a burst up to 5
		//"listTools": rate.NewLimiter(rate.Every(time.Minute), 10), // once a minute with a burst up to 20
	}))
	return server
}

func receivingHandler(next mcp.MethodHandler) mcp.MethodHandler {
	return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
		extra := req.GetExtra()
		result, err := next(ctx, method, req)
		if method == "tools/call" &&
			slices.Contains(extra.TokenInfo.Scopes, md.UserGroupGuest.String()) {
			return nil, errors.New(http.StatusText(http.StatusUnauthorized))
		}
		return result, err
	}
}

func sendingHandler(next mcp.MethodHandler) mcp.MethodHandler {
	return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
		return next(ctx, method, req)
	}
}

// GlobalRateLimiterMiddleware creates a middleware that applies a global rate limit.
// Every request attempting to pass through will try to acquire a token.
// If a token cannot be acquired immediately, the request will be rejected.
func globalRateLimiterMiddleware(limiter *rate.Limiter) mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			if !limiter.Allow() {
				return nil, errors.New("JSON RPC overloaded")
			}
			return next(ctx, method, req)
		}
	}
}

// PerMethodRateLimiterMiddleware creates a middleware that applies rate limiting
// on a per-method basis.
// Methods not specified in limiters will not be rate limited by this middleware.
func perMethodRateLimiterMiddleware(limiters map[string]*rate.Limiter) mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			if limiter, ok := limiters[method]; ok {
				if !limiter.Allow() {
					return nil, errors.New("JSON RPC overloaded")
				}
			}
			return next(ctx, method, req)
		}
	}
}

// PerSessionRateLimiterMiddleware creates a middleware that applies rate limiting
// on a per-session basis for receiving requests.
func perSessionRateLimiterMiddleware(limit rate.Limit, burst int) mcp.Middleware {
	// A map to store limiters, keyed by the session ID.
	var (
		sessionLimiters = make(map[string]*rate.Limiter)
		mu              sync.Mutex
	)

	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			// It's possible that session.ID() may be empty at this point in time
			// for some transports (e.g., stdio) or until the MCP initialize handshake
			// has completed.
			sessionID := req.GetSession().ID()
			if sessionID == "" {
				// In this situation, you could apply a single global identifier
				// if session ID is empty or bypass the rate limiter.
				// In this example, we bypass the rate limiter.
				log.Printf("Warning: Session ID is empty for method %q. Skipping per-session rate limiting.", method)
				return next(ctx, method, req) // Skip limiting if ID is unavailable
			}
			mu.Lock()
			limiter, ok := sessionLimiters[sessionID]
			if !ok {
				limiter = rate.NewLimiter(limit, burst)
				sessionLimiters[sessionID] = limiter
			}
			mu.Unlock()
			if !limiter.Allow() {
				return nil, errors.New("JSON RPC overloaded")
			}
			return next(ctx, method, req)
		}
	}
}

func GetServer(config cu.IM) func(*http.Request) *mcp.Server {
	return func(req *http.Request) *mcp.Server {
		return NewMCPServer(config)
	}
}

func TokenAuth(opt md.AuthOptions) (*auth.TokenInfo, error) {
	alias := cu.ToString(opt.Config["NT_DEFAULT_ALIAS"], "")
	userCode := cu.ToString(opt.Config["NT_DEFAULT_ADMIN"], "")
	var user md.Auth = md.Auth{
		UserGroup: md.UserGroupAdmin,
		Code:      userCode,
		UserName:  "admin",
	}
	if slices.Contains(strings.Split(cu.ToString(opt.Config["NT_API_KEY"], ""), ","), opt.TokenString) {
		ds := api.NewDataStore(opt.Config, alias, opt.AppLog)
		return &auth.TokenInfo{
			Scopes:     []string{md.UserGroupAdmin.String()}, // User permissions
			Expiration: time.Now().Add(24 * time.Hour),       // 24 hour expiration
			Extra: cu.IM{
				"user":   user,
				"config": opt.Config,
				"ds":     ds,
			},
		}, nil
	}

	var tokenData cu.IM
	var err error
	if tokenData, err = opt.ParseToken(opt.TokenString, opt.Config["tokenKeys"].([]cu.SM), opt.Config); err != nil {
		return nil, fmt.Errorf("%w: %v", auth.ErrInvalidToken, err)
	}
	alias = cu.ToString(tokenData["alias"], alias)
	userCode = cu.ToString(tokenData["user_code"], "")
	userName := cu.ToString(tokenData["user_name"], cu.ToString(tokenData["email"], ""))
	ds := api.NewDataStore(opt.Config, alias, opt.AppLog)
	if user, err = ds.AuthUser(userCode, userName); err != nil {
		return nil, fmt.Errorf("%w: %v", errors.New("error authenticating user"), err)
	}
	return &auth.TokenInfo{
		Scopes:     []string{user.UserGroup.String()},               // User permissions
		Expiration: time.Unix(cu.ToInteger(tokenData["exp"], 0), 0), // Token expiration time
		Extra: cu.IM{
			"user":   user,
			"token":  tokenData,
			"config": opt.Config,
			"ds":     ds,
		},
	}, nil
}

/*
func greetingPrompt(ctx context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {


	return &mcp.GetPromptResult{
		Description: "Hi prompt",
		Messages: []*mcp.PromptMessage{
			{
				Role:    "user",
				Content: &mcp.TextContent{Text: "Say hi to " + req.Params.Arguments["name"]},
			},
			{
				Role:    "user",
				Content: &mcp.ResourceLink{},
			},
		},
	}, nil
}
*/
