package http

import (
	"context"
	"fmt"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
	"golang.org/x/oauth2"
)

func getServers(r *http.Request) (publicHost, authServer string) {
	config := r.Context().Value(md.ConfigCtxKey).(cu.IM)
	localServer := fmt.Sprintf("http://%s", r.Host)
	publicHost = cu.ToString(config["NT_PUBLIC_HOST"], localServer)
	authServer = cu.ToString(config["NT_AUTH_SERVER"], localServer)
	return publicHost, authServer
}

func ProtectedResource(w http.ResponseWriter, r *http.Request) {
	publicHost, authServer := getServers(r)
	resourceMetadata := cu.IM{
		"resource":              publicHost,
		"authorization_servers": []string{authServer},
		"scopes_supported":      []string{"openid", "email"},
	}
	RespondMessage(w, http.StatusOK, resourceMetadata, http.StatusInternalServerError, nil)
}

// OAuth 2.1 Authorization Server Metadata endpoint
func AuthorizationServer(w http.ResponseWriter, r *http.Request) {
	_, authServer := getServers(r)
	serverMetadata := cu.IM{
		"issuer":                                authServer,
		"authorization_endpoint":                authServer + "/oauth/authorization",
		"token_endpoint":                        authServer + "/oauth/token",
		"registration_endpoint":                 authServer + "/oauth/registration",
		"jwks_uri":                              authServer + "/.well-known/jwks.json",
		"response_types_supported":              []string{"code"},
		"code_challenge_methods_supported":      []string{"S256"},
		"token_endpoint_auth_methods_supported": []string{"client_secret_post"},
		"grant_types_supported":                 []string{"authorization_code"},
	}
	RespondMessage(w, http.StatusOK, serverMetadata, http.StatusInternalServerError, nil)
}

// OpenID Connect Discovery endpoint
func OpenIDConfiguration(w http.ResponseWriter, r *http.Request) {
	_, authServer := getServers(r)
	serverMetadata := cu.IM{
		"issuer":                                authServer,
		"authorization_endpoint":                authServer + "/oauth/authorization",
		"token_endpoint":                        authServer + "/oauth/token",
		"registration_endpoint":                 authServer + "/oauth/registration",
		"jwks_uri":                              authServer + "/.well-known/jwks.json",
		"response_types_supported":              []string{"code"},
		"subject_types_supported":               []string{"public"},
		"id_token_signing_alg_values_supported": []string{"RS256"},
	}
	RespondMessage(w, http.StatusOK, serverMetadata, http.StatusInternalServerError, nil)
}

// JSON Web Key Set (JWKS) public keys endpoint
func Jwks(w http.ResponseWriter, r *http.Request) {
	config := r.Context().Value(md.ConfigCtxKey).(cu.IM)
	serverMetadata := cu.IM{
		"keys": []cu.IM{
			{
				"kid": cu.ToString(config["NT_TOKEN_PUBLIC_KID"], "ntt-public-key"),
				"kty": "RSA",
				"use": "sig",
				"alg": "RS256",
				"e":   "AQAB",
				"n":   cu.ToString(config["NT_TOKEN_PUBLIC_KEY"], ""),
			},
		},
	}
	RespondMessage(w, http.StatusOK, serverMetadata, http.StatusInternalServerError, nil)
}

func getAuthConfig(r *http.Request) *oauth2.Config {
	config := r.Context().Value(md.ConfigCtxKey).(cu.IM)
	localServer := fmt.Sprintf("http://%s", r.Host)
	publicHost := cu.ToString(config["NT_PUBLIC_HOST"], localServer)
	redirectURL := fmt.Sprintf("%s%s", publicHost, st.OAuthAuthRedirectURL)
	return &oauth2.Config{
		ClientID:     cu.ToString(config["NT_AUTH_CLIENT_ID"], ""),
		ClientSecret: cu.ToString(config["NT_AUTH_CLIENT_SECRET"], ""),
		RedirectURL:  redirectURL,
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:       cu.ToString(config["NT_AUTH_AUTHORIZATION_ENDPOINT"], ""),
			TokenURL:      cu.ToString(config["NT_AUTH_TOKEN_ENDPOINT"], ""),
			DeviceAuthURL: cu.ToString(config["NT_AUTH_DEVICE_ENDPOINT"], ""),
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

// OAuth authorization (redirects to authorization server)
func OAuthAuthorization(w http.ResponseWriter, r *http.Request) {
	sessionService := r.Context().Value(md.SessionServiceCtxKey).(*api.SessionService)
	sessionID := cu.GetComponentID()
	verifier := oauth2.GenerateVerifier()
	sessionData := cu.IM{
		"session_id": sessionID,
		"verifier":   verifier,
	}
	sessionService.SaveSession(sessionID, sessionData)
	authConfig := getAuthConfig(r)
	url := authConfig.AuthCodeURL(sessionID, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Token exchange with OAuth provider
func OAuthToken(w http.ResponseWriter, r *http.Request) {
	RespondMessage(w, http.StatusOK, cu.IM{}, http.StatusInternalServerError, nil)
}

// Dynamic client registration
func OAuthRegistration(w http.ResponseWriter, r *http.Request) {
	RespondMessage(w, http.StatusOK, cu.IM{}, http.StatusInternalServerError, nil)
}

// Authorized redirect URIs for the OAuth 2.0 client where the API server redirects the user after the user completes the authorization flow.
func OAuthCallback(w http.ResponseWriter, r *http.Request) {
	sessionService := r.Context().Value(md.SessionServiceCtxKey).(*api.SessionService)

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	var sessionData any
	var err error
	if sessionData, err = sessionService.LoadSession(state, &cu.IM{}); err != nil {
		RespondMessage(w, http.StatusInternalServerError, cu.IM{
			"error": err.Error(),
		}, http.StatusInternalServerError, nil)
		return
	}
	verifier := cu.ToString(cu.ToIM(sessionData, cu.IM{})["verifier"], "")
	authConfig := getAuthConfig(r)
	var token *oauth2.Token
	token, err = authConfig.Exchange(context.Background(), code, oauth2.VerifierOption(verifier))
	RespondMessage(w, http.StatusOK, token, http.StatusInternalServerError, err)
}
