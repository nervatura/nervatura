package mcp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	"golang.org/x/time/rate"
)

type McpServer struct {
	scope  string
	config cu.IM
	appLog *slog.Logger
	scopes []string
}

var toolDataMap map[string]ToolData = map[string]ToolData{}

func (ms *McpServer) NewMCPServer(scope string) (server *mcp.Server) {
	opts := &mcp.ServerOptions{
		Instructions: cu.ToString(ScopeInstruction[scope], ScopeInstruction["root"]),
		GetSessionID: cu.GetComponentID,
	}
	server = mcp.NewServer(&mcp.Implementation{
		Name: "nervatura", Title: "Nervatura MCP Server",
		Version: cu.ToString(ms.config["version"], "0.0.0"),
	}, opts)

	for key, td := range toolDataMap {
		if slices.Contains(td.Scopes, scope) || len(td.Scopes) == 0 || scope == "all" {
			addTool(key, server, scope)
		}
	}

	if resources, ok := ms.config["resources"].(map[string]ResourceData); ok {
		for _, resource := range resources {
			if slices.Contains(resource.Scopes, scope) || len(resource.Scopes) == 0 || scope == "all" {
				server.AddResource(&mcp.Resource{
					Name:        resource.Name,
					Title:       resource.Title,
					Description: resource.Description,
					MIMEType:    resource.MIMEType,
					URI:         resource.URI,
				}, resourceHandler)
			}
		}
	}

	if prompts, ok := ms.config["prompts"].(map[string]PromptData); ok {
		for _, prompt := range prompts {
			if slices.Contains(prompt.Scopes, scope) || len(prompt.Scopes) == 0 || scope == "all" {
				server.AddPrompt(&mcp.Prompt{
					Name:        prompt.Name,
					Title:       prompt.Title,
					Description: prompt.Description,
					Arguments:   prompt.Arguments,
				}, promptHandler)
			}
		}
	}

	server.AddReceivingMiddleware(ms.receivingHandler)
	server.AddSendingMiddleware(ms.sendingHandler)
	server.AddReceivingMiddleware(ms.globalRateLimiterMiddleware(rate.NewLimiter(rate.Every(time.Second/5), 10)))
	//server.AddReceivingMiddleware(ms.perSessionRateLimiterMiddleware(rate.Every(time.Second/5), 10))
	server.AddReceivingMiddleware(ms.perMethodRateLimiterMiddleware(map[string]*rate.Limiter{
		//"tools/call":  rate.NewLimiter(rate.Every(time.Second), 5),  // once a second with a burst up to 5
		//"listTools": rate.NewLimiter(rate.Every(time.Minute), 10), // once a minute with a burst up to 20
	}))
	return server
}

func (ms *McpServer) receivingHandler(next mcp.MethodHandler) mcp.MethodHandler {
	return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
		extra := req.GetExtra()
		if method == "tools/call" && extra.TokenInfo != nil &&
			slices.Contains(extra.TokenInfo.Scopes, md.UserGroupGuest.String()) {
			return nil, errors.New(http.StatusText(http.StatusUnauthorized))
		}

		if extra.TokenInfo != nil {
			for _, scope := range extra.TokenInfo.Scopes {
				if slices.Contains(ms.scopes, scope) {
					if !slices.Contains(extra.TokenInfo.Scopes, ms.scope) {
						return nil, errors.New(http.StatusText(http.StatusUnauthorized))
					}
				}
			}
		}

		ms.config["MCP_SCOPE"] = ms.scope
		alias := cu.ToString(ms.config["NT_DEFAULT_ALIAS"], "")
		if extra.TokenInfo != nil {
			alias = cu.ToString(extra.TokenInfo.Extra["alias"], alias)
		}
		ds := api.NewDataStore(ms.config, alias, ms.appLog)
		ctx = context.WithValue(ctx, md.DataStoreCtxKey, ds)
		ctx = context.WithValue(ctx, md.ConfigCtxKey, ms.config)
		result, err := next(ctx, method, req)
		return result, err
	}
}

func (ms *McpServer) sendingHandler(next mcp.MethodHandler) mcp.MethodHandler {
	return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
		return next(ctx, method, req)
	}
}

// GlobalRateLimiterMiddleware creates a middleware that applies a global rate limit.
// Every request attempting to pass through will try to acquire a token.
// If a token cannot be acquired immediately, the request will be rejected.
func (ms *McpServer) globalRateLimiterMiddleware(limiter *rate.Limiter) mcp.Middleware {
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
func (ms *McpServer) perMethodRateLimiterMiddleware(limiters map[string]*rate.Limiter) mcp.Middleware {
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

/*
// PerSessionRateLimiterMiddleware creates a middleware that applies rate limiting
// on a per-session basis for receiving requests.
func (ms *McpServer) perSessionRateLimiterMiddleware(limit rate.Limit, burst int) mcp.Middleware {
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
*/

func GetServer(scope string, config cu.IM, appLog *slog.Logger) func(*http.Request) *mcp.Server {
	return func(req *http.Request) *mcp.Server {
		ms := &McpServer{
			scope:  scope,
			config: config,
			appLog: appLog,
			scopes: []string{},
		}
		for key := range ScopeInstruction {
			ms.scopes = append(ms.scopes, key)
		}
		return ms.NewMCPServer(scope)
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
		return &auth.TokenInfo{
			Scopes:     []string{md.UserGroupAdmin.String()}, // User permissions
			Expiration: time.Now().Add(24 * time.Hour),       // 24 hour expiration
			Extra: cu.IM{
				"alias": alias,
				"user":  user,
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
			"alias": alias,
			"user":  user,
			"token": tokenData,
		},
	}, nil
}
