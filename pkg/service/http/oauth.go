package http

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

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

func loginTemplate(w http.ResponseWriter, params cu.SM) {
	data := cu.IM{
		"title":          "OAuth Authorization",
		"subtitle":       cu.ToString(params["subtitle"], "Sign in to your account"),
		"username_label": "Username or email",
		"password_label": "Password",
		"login_button":   "Sign in",
		"username":       cu.ToString(params["username"], ""),
		"session_id":     cu.ToString(params["session_id"], ""),
		"error_msg":      cu.ToString(params["error_msg"], ""),
		"code_label":     "Code",
		"state_label":    "State",
		"code":           cu.ToString(params["code"], ""),
		"state":          cu.ToString(params["state"], ""),
	}
	tmp, _ := template.New("auth").Parse(st.AuthPage)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tmp.ExecuteTemplate(w, "auth", data)
}

// OAuth authorization (login page)
func OAuthAuthorization(w http.ResponseWriter, r *http.Request) {
	config := r.Context().Value(md.ConfigCtxKey).(cu.IM)
	sessionService := r.Context().Value(md.SessionServiceCtxKey).(*api.SessionService)
	sessionID := cu.GetComponentID()
	sessionData := cu.SM{
		"response_type":         cu.ToString(r.URL.Query().Get("response_type"), "code"),
		"client_id":             cu.ToString(r.URL.Query().Get("client_id"), cu.ToString(config["NT_AUTH_CLIENT_ID"], "")),
		"redirect_uri":          r.URL.Query().Get("redirect_uri"),
		"state":                 r.URL.Query().Get("state"),
		"code_challenge":        r.URL.Query().Get("code_challenge"),
		"code_challenge_method": r.URL.Query().Get("code_challenge_method"),
		"scope":                 r.URL.Query().Get("scope"),
	}
	if sessionData["response_type"] != "code" {
		http.Error(w, "unsupported_response_type", http.StatusBadRequest)
		return
	}
	if _, err := url.ParseRequestURI(sessionData["redirect_uri"]); err != nil && sessionData["redirect_uri"] != "" {
		http.Error(w, "invalid_redirect_uri", http.StatusBadRequest)
		return
	}
	if sessionData["client_id"] != "" && sessionData["client_id"] != cu.ToString(config["NT_AUTH_CLIENT_ID"], "") {
		http.Error(w, "invalid_client_id", http.StatusBadRequest)
		return
	}
	sessionService.SaveSession(sessionID, sessionData)
	loginTemplate(w, cu.SM{"session_id": sessionID})
}

// OAuth authorization (redirects to authorization server)
func OAuthLogin(w http.ResponseWriter, r *http.Request) {
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

func OAuthValidate(w http.ResponseWriter, r *http.Request) {
	sessionService := r.Context().Value(md.SessionServiceCtxKey).(*api.SessionService)
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	//errMsg := ut.GetMessage("en", "failed_authentication")

	var err error
	if err = r.ParseForm(); err != nil {
		loginTemplate(w, cu.SM{"error_msg": "Invalid parameters"})
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	//database := cu.ToString(r.Form.Get("database"), cu.ToString(config["NT_DEFAULT_ALIAS"], ""))
	sessionID := r.Form.Get("session_id")

	var sessionData any
	if sessionData, err = sessionService.LoadSession(sessionID, &cu.IM{}); err != nil {
		loginTemplate(w, cu.SM{"error_msg": "Invalid session ID", "username": username})
		return
	}
	session := cu.ToSM(sessionData, cu.SM{})

	var user md.Auth
	var token string
	if user, err = ds.AuthUser("", username); err != nil {
		loginTemplate(w, cu.SM{"error_msg": "Invalid email or username", "username": username, "session_id": sessionID})
		return
	}
	if token, err = ds.UserLogin(username, password, true); err != nil {
		loginTemplate(w, cu.SM{"error_msg": err.Error(), "username": username, "session_id": sessionID})
		return
	}
	session["token"] = token
	session["scope"] = user.UserGroup.String()
	sessionService.SaveSession(sessionID, session)

	callbackURL := cu.ToString(session["redirect_uri"], "")
	state := cu.ToString(session["state"], "")
	if callbackURL != "" {
		callbackURL = fmt.Sprintf("%s?code=%s", callbackURL, sessionID)
		if state != "" {
			callbackURL = fmt.Sprintf("%s&state=%s", callbackURL, state)
		}
		http.Redirect(w, r, callbackURL, http.StatusMovedPermanently)
		return
	}

	loginTemplate(w, cu.SM{"subtitle": "Successfully authenticated", "code": sessionID, "state": state})
}

// Token exchange with OAuth provider
func OAuthToken(w http.ResponseWriter, r *http.Request) {
	config := r.Context().Value(md.ConfigCtxKey).(cu.IM)
	sessionService := r.Context().Value(md.SessionServiceCtxKey).(*api.SessionService)

	var err error
	if err = r.ParseForm(); err != nil {
		RespondMessage(w, http.StatusBadRequest,
			cu.IM{"error": "invalid_request", "error_description": "Invalid request parameters"}, http.StatusBadRequest, nil)
		return
	}

	if r.Form.Get("grant_type") != "authorization_code" {
		RespondMessage(w, http.StatusBadRequest,
			cu.IM{"error": "unsupported_grant_type", "error_description": "Unsupported grant type"}, http.StatusBadRequest, nil)
		return
	}
	if r.Form.Get("code") == "" {
		RespondMessage(w, http.StatusBadRequest,
			cu.IM{"error": "invalid_request", "error_description": "Invalid or missing code"}, http.StatusBadRequest, nil)
		return
	}
	if r.Form.Get("client_id") != cu.ToString(config["NT_AUTH_CLIENT_ID"], "") {
		RespondMessage(w, http.StatusBadRequest,
			cu.IM{"error": "invalid_request", "error_description": "Invalid or missing client_id"}, http.StatusBadRequest, nil)
		return
	}
	var sessionData any
	if sessionData, err = sessionService.LoadSession(r.Form.Get("code"), &cu.IM{}); err != nil {
		RespondMessage(w, http.StatusBadRequest,
			cu.IM{"error": "invalid_request", "error_description": "Invalid or missing code"}, http.StatusBadRequest, nil)
		return
	}
	session := cu.ToSM(sessionData, cu.SM{})
	response := cu.IM{
		"access_token": session["token"],
		"token_type":   "bearer",
		"expires_in":   time.Now().Add(time.Duration(cu.ToFloat(config["NT_SESSION_EXP"], 1)) * time.Hour).Unix(),
		"scope":        session["scope"],
	}
	RespondMessage(w, http.StatusOK, response, http.StatusInternalServerError, nil)
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
