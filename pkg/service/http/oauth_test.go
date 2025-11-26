package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestProtectedResource(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/oauth-protected-resource", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			ProtectedResource(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthorizationServer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-authorization-server", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			AuthorizationServer(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOpenIDConfiguration(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/openid-configuration", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			OpenIDConfiguration(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestJwks(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/jwks.json", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			Jwks(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOAuthAuthorization(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-authorization", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			ctx = context.WithValue(ctx, md.SessionServiceCtxKey, &api.SessionService{})
			OAuthAuthorization(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOAuthToken(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-token", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OAuthToken(tt.w, tt.r)
		})
	}
}

func TestOAuthRegistration(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-registration", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OAuthRegistration(tt.w, tt.r)
		})
	}
}

func TestOAuthCallback(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-callback?code=test&state=SES012345", nil),
		},
		{
			name: "error",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-callback?code=test&state=SES012346", nil),
		},
	}
	ses := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &md.TestDriver{Config: cu.IM{}},
	}
	ses.SaveSession("SES012345", cu.IM{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			ctx = context.WithValue(ctx, md.SessionServiceCtxKey, ses)
			OAuthCallback(tt.w, tt.r.WithContext(ctx))
		})
	}
}
