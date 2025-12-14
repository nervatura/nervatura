package mcp

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	"golang.org/x/time/rate"
)

func TestTokenAuth(t *testing.T) {
	type args struct {
		opt md.AuthOptions
	}
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Authorization", "Bearer test")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid token",
			args: args{
				opt: md.AuthOptions{
					TokenString: "test",
					Request:     req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
						"db": &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
					},
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{"alias": "test", "user_code": "123456", "user_name": "test"}, nil
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				},
			},
		},
		{
			name: "api_key token",
			args: args{
				opt: md.AuthOptions{
					TokenString: "test",
					Request:     req,
					Config: cu.IM{
						"NT_API_KEY": "test",
						"tokenKeys":  []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
					},
				},
			},
		},
		{
			name: "invalid token",
			args: args{
				opt: md.AuthOptions{
					TokenString: "test",
					Request:     req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
						"db": &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
					},
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{}, errors.New("invalid token")
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				},
			},
			wantErr: true,
		},
		{
			name: "auth error",
			args: args{
				opt: md.AuthOptions{
					TokenString: "test",
					Request:     req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
						"db": &md.TestDriver{
							Config: cu.IM{},
						},
					},
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{"alias": "test", "user_code": "123456", "user_name": "test"}, nil
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TokenAuth(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetServer(t *testing.T) {
	type args struct {
		scope   string
		config  cu.IM
		appLog  *slog.Logger
		session *api.SessionService
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "all scope",
			args: args{
				scope: "all",
				config: cu.IM{
					"resources": map[string]ResourceData{
						"test": {
							Resource: mcp.Resource{
								Name:        "test",
								Title:       "test",
								Description: "test",
							},
							Scopes: []string{"customer"},
						},
					},
					"prompts": map[string]PromptData{
						"test": {
							Name:        "test",
							Title:       "test",
							Description: "test",
							Scopes:      []string{"customer"},
						},
					},
				},
				appLog:  slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				session: &api.SessionService{},
			},
		},
		{
			name: "customer scope",
			args: args{
				scope: "customer",
				config: cu.IM{
					"resources": map[string]ResourceData{
						"test": {
							Resource: mcp.Resource{
								Name:        "test",
								Title:       "test",
								Description: "test",
							},
							Scopes: []string{"customer"},
						},
					},
					"prompts": map[string]PromptData{
						"test": {
							Name:        "test",
							Title:       "test",
							Description: "test",
							Scopes:      []string{"customer"},
						},
					},
				},
				appLog:  slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				session: &api.SessionService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetServer(tt.args.scope, tt.args.config, tt.args.appLog, tt.args.session)(httptest.NewRequest("POST", "/", nil))
		})
	}
}

func TestMcpServer_sendingHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		next mcp.MethodHandler
	}{
		{
			name: "valid request",
			next: func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
				return nil, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ms McpServer
			ms.sendingHandler(tt.next)(context.Background(), "test", &mcp.ServerRequest[*mcp.CallToolParams]{
				Session: &mcp.ServerSession{},
				Params:  &mcp.CallToolParams{},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{},
					Header:    http.Header{},
				},
			})
			// TODO: update the condition below to compare got with tt.want.
		})
	}
}

func TestMcpServer_receivingHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		next      mcp.MethodHandler
		tokenInfo *auth.TokenInfo
	}{
		{
			name: "guest user",
			next: func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
				return nil, nil
			},
			tokenInfo: &auth.TokenInfo{
				Scopes: []string{md.UserGroupGuest.String()},
			},
		},
		{
			name: "restricted scope",
			next: func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
				return nil, nil
			},
			tokenInfo: &auth.TokenInfo{
				Scopes: []string{md.UserGroupUser.String(), "customer"},
			},
		},
		{
			name: "valid",
			next: func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
				return nil, nil
			},
			tokenInfo: &auth.TokenInfo{
				Scopes: []string{md.UserGroupAdmin.String()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var ms McpServer = McpServer{
				config: cu.IM{
					"NT_DEFAULT_ALIAS": "test",
				},
				appLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				scope:  "product",
				scopes: []string{"product", "customer"},
			}
			ms.receivingHandler(tt.next)(context.Background(), "tools/call", &mcp.ServerRequest[*mcp.CallToolParams]{
				Session: &mcp.ServerSession{},
				Params:  &mcp.CallToolParams{},
				Extra: &mcp.RequestExtra{
					TokenInfo: tt.tokenInfo,
					Header:    http.Header{},
				},
			})
		})
	}
}

func TestMcpServer_globalRateLimiterMiddleware(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		limiter *rate.Limiter
	}{
		{
			name:    "allow",
			limiter: rate.NewLimiter(rate.Every(time.Second), 1),
		},
		{
			name:    "deny",
			limiter: rate.NewLimiter(rate.Every(time.Second), 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var ms McpServer
			got := ms.globalRateLimiterMiddleware(tt.limiter)
			got2 := got(func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
				return nil, nil
			})
			got2(context.Background(), "tools/call", &mcp.ServerRequest[*mcp.CallToolParams]{
				Session: &mcp.ServerSession{},
				Params:  &mcp.CallToolParams{},
				Extra:   &mcp.RequestExtra{},
			})

		})
	}
}

func TestMcpServer_perMethodRateLimiterMiddleware(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		limiters map[string]*rate.Limiter
	}{
		{
			name: "success",
			limiters: map[string]*rate.Limiter{
				"tools/call": rate.NewLimiter(rate.Every(time.Second), 1),
			},
		},
		{
			name: "deny",
			limiters: map[string]*rate.Limiter{
				"tools/call": rate.NewLimiter(rate.Every(time.Second), 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var ms McpServer
			got := ms.perMethodRateLimiterMiddleware(tt.limiters)
			got2 := got(func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
				return nil, nil
			})
			got2(context.Background(), "tools/call", &mcp.ServerRequest[*mcp.CallToolParams]{
				Session: &mcp.ServerSession{},
				Params:  &mcp.CallToolParams{},
				Extra:   &mcp.RequestExtra{},
			})
		})
	}
}

func TestCatalog(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Catalog(tt.w, tt.r)
		})
	}
}
