//go:build http || all
// +build http all

package server

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	cst "github.com/nervatura/component/pkg/static"
	cu "github.com/nervatura/component/pkg/util"
	docs "github.com/nervatura/nervatura/v6/docs6"
	cl "github.com/nervatura/nervatura/v6/pkg/component/client/service"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	src "github.com/nervatura/nervatura/v6/pkg/service/gui"
	srv "github.com/nervatura/nervatura/v6/pkg/service/http"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type httpServer struct {
	config     cu.IM
	appLog     *slog.Logger
	mux        *http.ServeMux
	server     *http.Server
	tlsEnabled bool
	result     string
}

func init() {
	registerHost("http", &httpServer{})
}

func (s *httpServer) StartServer(config cu.IM, appLogOut, httpLogOut io.Writer, interrupt chan os.Signal) error {
	s.config = config
	s.appLog = slog.New(slog.NewJSONHandler(appLogOut, nil))
	s.mux = http.NewServeMux()

	s.setRoutes()

	rootHandler := handlers.CompressHandler(handlers.RecoveryHandler()(handlers.CombinedLoggingHandler(httpLogOut, s.mux)))
	s.server = &http.Server{
		Handler:      rootHandler,
		Addr:         fmt.Sprintf(":%d", cu.ToInteger(s.config["NT_HTTP_PORT"], 0)),
		ReadTimeout:  time.Duration(cu.ToFloat(s.config["NT_HTTP_READ_TIMEOUT"], 0)) * time.Second,
		WriteTimeout: time.Duration(cu.ToFloat(s.config["NT_HTTP_WRITE_TIMEOUT"], 0)) * time.Second,
	}
	s.tlsEnabled = cu.ToBoolean(s.config["NT_HTTP_TLS_ENABLED"], false) &&
		cu.ToString(s.config["NT_TLS_CERT_FILE"], "") != "" && cu.ToString(s.config["NT_TLS_KEY_FILE"], "") != ""

	s.appLog.Info(fmt.Sprintf("HTTP server serving at: %d. SSL/TLS authentication: %v.",
		cu.ToInteger(s.config["NT_HTTP_PORT"], 0), s.tlsEnabled))
	if s.tlsEnabled {
		return s.server.ListenAndServeTLS(
			cu.ToString(s.config["NT_TLS_CERT_FILE"], ""),
			cu.ToString(s.config["NT_TLS_KEY_FILE"], ""))
	}
	return s.server.ListenAndServe()
}

func (s *httpServer) StopServer(ctx context.Context) error {
	if s.server != nil {
		s.appLog.Info("stopping HTTP server")
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *httpServer) Results() string {
	return s.result
}

// Register API routes.
func (s *httpServer) setRoutes() {
	clientCSRF := src.ClientCsrfProtect(s.config)
	CORS := handlers.CORS(
		handlers.AllowedOrigins(s.config["NT_CORS_ALLOW_ORIGINS"].([]string)),
		handlers.AllowedMethods(s.config["NT_CORS_ALLOW_METHODS"].([]string)),
		handlers.AllowedHeaders(s.config["NT_CORS_ALLOW_HEADERS"].([]string)),
		handlers.ExposedHeaders(s.config["NT_CORS_EXPOSE_HEADERS"].([]string)),
		handlers.AllowCredentials(),
		handlers.MaxAge(int(cu.ToInteger(s.config["NT_CORS_MAX_AGE"], 0))),
	)

	s.mux.HandleFunc("/", s.homeRoute)
	s.mux.HandleFunc("/config/{secKey}", s.configRoute)

	s.mux.Handle(st.ClientPath+"/", clientCSRF(s.headerClient(s.clientUIRoutes())))
	s.mux.Handle(st.ClientPath+"/api/", CORS(s.headerClient(s.clientAPIRoutes())))

	s.mux.Handle(st.ApiPath+"/", CORS(s.headerAPI(s.apiRoutes())))

	// Register static dirs.
	// app css files
	var publicFS, _ = fs.Sub(st.Public, "public")
	// components css files
	var staticFS, _ = fs.Sub(cst.Static, ".")
	s.mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.FS(publicFS))))
	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
	s.mux.Handle("/docs6/", http.StripPrefix("/docs6/", http.FileServer(http.FS(docs.Docs))))
}

func (s *httpServer) homeRoute(w http.ResponseWriter, r *http.Request) {
	home := cu.ToString(s.config["NT_HTTP_HOME"], "/")
	if home != "/" {
		http.Redirect(w, r, home, http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, st.ClientPath+"/auth/", http.StatusSeeOther)
}

func (s *httpServer) envList() []cu.IM {
	envResult := make([]cu.IM, 0)
	keys := make([]string, 0)
	configs := cu.IM{}
	for key, value := range s.config {
		keys = append(keys, key)
		configs[key] = value
	}
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "NT_ALIAS_") {
			keys = append(keys, strings.Split(env, "=")[0])
			configs[strings.Split(env, "=")[0]] = strings.Split(env, "=")[1]
		}
	}

	sort.Strings(keys)
	for _, key := range keys {
		envResult = append(envResult, cu.IM{"envkey": strings.ToUpper(key), "envvalue": cu.ToString(configs[key], "")})
	}
	return envResult
}

func (s *httpServer) configRoute(w http.ResponseWriter, r *http.Request) {
	secKey := cu.ToString(r.PathValue("secKey"), "")
	if secKey != cu.ToString(s.config["NT_TASK_SEC_KEY"], "") {
		http.Error(w, "Missing or invalid authentication key", http.StatusUnauthorized)
		return
	}
	data := cu.IM{
		"title":      "Configuration values",
		"env_result": s.envList(),
	}
	tmp, _ := template.New("task").Parse(st.TaskPage)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tmp.ExecuteTemplate(w, "task", data)
}

func (s *httpServer) headerClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := cl.NewClientService(s.config, s.appLog)
		ctx := context.WithValue(r.Context(), md.ClientServiceCtxKey, client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *httpServer) headerAPI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		state := func() string {
			if strings.Contains(r.URL.Path, "/auth/login") {
				return "LOGIN"
			}
			if r.Header.Get("X-API-KEY") != "" {
				return "API_KEY"
			}
			return "TOKEN"
		}
		authOptions := md.AuthOptions{
			Request: r, Config: s.config, AppLog: s.appLog, ParseToken: ut.ParseToken, ConvertFromReader: cu.ConvertFromReader,
		}

		var ctx context.Context = r.Context()
		var errCode int
		switch state() {
		case "LOGIN":
			ctx = context.WithValue(r.Context(), md.AuthOptionsCtxKey, authOptions)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		case "API_KEY":
			ctx, errCode = srv.ApiKeyAuth(authOptions)
		default:
			ctx, errCode = srv.TokenAuth(authOptions)
		}

		if errCode > 0 {
			http.Error(w, http.StatusText(errCode), errCode)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *httpServer) apiRoutes() http.Handler {
	apiMux := http.NewServeMux()

	apiMux.HandleFunc("POST /auth/login", srv.AuthLogin)
	apiMux.HandleFunc("POST /auth", srv.AuthPost)
	apiMux.HandleFunc("GET /auth/{id_code}", srv.AuthGet)
	apiMux.HandleFunc("PUT /auth/{id_code}", srv.AuthPut)
	apiMux.HandleFunc("GET /auth/me", srv.AuthQuery)
	apiMux.HandleFunc("PUT /auth/me", srv.AuthPut)
	apiMux.HandleFunc("POST /auth/me/password", srv.AuthPassword)
	apiMux.HandleFunc("POST /auth/me/reset", srv.AuthReset)
	apiMux.HandleFunc("GET /auth/me/token", srv.AuthToken)
	apiMux.HandleFunc("PUT /auth/reset/{id_code}", srv.AuthReset)

	apiMux.HandleFunc("GET /config", srv.ConfigQuery)
	apiMux.HandleFunc("POST /config", srv.ConfigPost)
	apiMux.HandleFunc("GET /config/{id_code}", srv.ConfigGet)
	apiMux.HandleFunc("PUT /config/{id_code}", srv.ConfigPut)
	apiMux.HandleFunc("DELETE /config/{id_code}", srv.ConfigDelete)

	apiMux.HandleFunc("GET /currency", srv.CurrencyQuery)
	apiMux.HandleFunc("POST /currency", srv.CurrencyPost)
	apiMux.HandleFunc("GET /currency/{id_code}", srv.CurrencyGet)
	apiMux.HandleFunc("PUT /currency/{id_code}", srv.CurrencyPut)
	apiMux.HandleFunc("DELETE /currency/{id_code}", srv.CurrencyDelete)

	apiMux.HandleFunc("GET /customer", srv.CustomerQuery)
	apiMux.HandleFunc("POST /customer", srv.CustomerPost)
	apiMux.HandleFunc("GET /customer/{id_code}", srv.CustomerGet)
	apiMux.HandleFunc("PUT /customer/{id_code}", srv.CustomerPut)
	apiMux.HandleFunc("DELETE /customer/{id_code}", srv.CustomerDelete)

	apiMux.HandleFunc("GET /employee", srv.EmployeeQuery)
	apiMux.HandleFunc("POST /employee", srv.EmployeePost)
	apiMux.HandleFunc("GET /employee/{id_code}", srv.EmployeeGet)
	apiMux.HandleFunc("PUT /employee/{id_code}", srv.EmployeePut)
	apiMux.HandleFunc("DELETE /employee/{id_code}", srv.EmployeeDelete)

	apiMux.HandleFunc("GET /item", srv.ItemQuery)
	apiMux.HandleFunc("POST /item", srv.ItemPost)
	apiMux.HandleFunc("GET /item/{id_code}", srv.ItemGet)
	apiMux.HandleFunc("PUT /item/{id_code}", srv.ItemPut)
	apiMux.HandleFunc("DELETE /item/{id_code}", srv.ItemDelete)

	apiMux.HandleFunc("GET /link", srv.LinkQuery)
	apiMux.HandleFunc("POST /link", srv.LinkPost)
	apiMux.HandleFunc("GET /link/{id_code}", srv.LinkGet)
	apiMux.HandleFunc("PUT /link/{id_code}", srv.LinkPut)
	apiMux.HandleFunc("DELETE /link/{id_code}", srv.LinkDelete)

	apiMux.HandleFunc("GET /log", srv.LogQuery)
	apiMux.HandleFunc("GET /log/{id_code}", srv.LogGet)

	apiMux.HandleFunc("GET /movement", srv.MovementQuery)
	apiMux.HandleFunc("POST /movement", srv.MovementPost)
	apiMux.HandleFunc("GET /movement/{id_code}", srv.MovementGet)
	apiMux.HandleFunc("PUT /movement/{id_code}", srv.MovementPut)
	apiMux.HandleFunc("DELETE /movement/{id_code}", srv.MovementDelete)

	apiMux.HandleFunc("GET /payment", srv.PaymentQuery)
	apiMux.HandleFunc("POST /payment", srv.PaymentPost)
	apiMux.HandleFunc("GET /payment/{id_code}", srv.PaymentGet)
	apiMux.HandleFunc("PUT /payment/{id_code}", srv.PaymentPut)
	apiMux.HandleFunc("DELETE /payment/{id_code}", srv.PaymentDelete)

	apiMux.HandleFunc("GET /place", srv.PlaceQuery)
	apiMux.HandleFunc("POST /place", srv.PlacePost)
	apiMux.HandleFunc("GET /place/{id_code}", srv.PlaceGet)
	apiMux.HandleFunc("PUT /place/{id_code}", srv.PlacePut)
	apiMux.HandleFunc("DELETE /place/{id_code}", srv.PlaceDelete)

	apiMux.HandleFunc("GET /price", srv.PriceQuery)
	apiMux.HandleFunc("POST /price", srv.PricePost)
	apiMux.HandleFunc("GET /price/{id_code}", srv.PriceGet)
	apiMux.HandleFunc("PUT /price/{id_code}", srv.PricePut)
	apiMux.HandleFunc("DELETE /price/{id_code}", srv.PriceDelete)

	apiMux.HandleFunc("GET /product", srv.ProductQuery)
	apiMux.HandleFunc("POST /product", srv.ProductPost)
	apiMux.HandleFunc("GET /product/{id_code}", srv.ProductGet)
	apiMux.HandleFunc("PUT /product/{id_code}", srv.ProductPut)
	apiMux.HandleFunc("DELETE /product/{id_code}", srv.ProductDelete)

	apiMux.HandleFunc("GET /project", srv.ProjectQuery)
	apiMux.HandleFunc("POST /project", srv.ProjectPost)
	apiMux.HandleFunc("GET /project/{id_code}", srv.ProjectGet)
	apiMux.HandleFunc("PUT /project/{id_code}", srv.ProjectPut)
	apiMux.HandleFunc("DELETE /project/{id_code}", srv.ProjectDelete)

	apiMux.HandleFunc("GET /rate", srv.RateQuery)
	apiMux.HandleFunc("POST /rate", srv.RatePost)
	apiMux.HandleFunc("GET /rate/{id_code}", srv.RateGet)
	apiMux.HandleFunc("PUT /rate/{id_code}", srv.RatePut)
	apiMux.HandleFunc("DELETE /rate/{id_code}", srv.RateDelete)

	apiMux.HandleFunc("GET /tax", srv.TaxQuery)
	apiMux.HandleFunc("POST /tax", srv.TaxPost)
	apiMux.HandleFunc("GET /tax/{id_code}", srv.TaxGet)
	apiMux.HandleFunc("PUT /tax/{id_code}", srv.TaxPut)
	apiMux.HandleFunc("DELETE /tax/{id_code}", srv.TaxDelete)

	apiMux.HandleFunc("GET /tool", srv.ToolQuery)
	apiMux.HandleFunc("POST /tool", srv.ToolPost)
	apiMux.HandleFunc("GET /tool/{id_code}", srv.ToolGet)
	apiMux.HandleFunc("PUT /tool/{id_code}", srv.ToolPut)
	apiMux.HandleFunc("DELETE /tool/{id_code}", srv.ToolDelete)

	apiMux.HandleFunc("GET /trans", srv.TransQuery)
	apiMux.HandleFunc("POST /trans", srv.TransPost)
	apiMux.HandleFunc("GET /trans/{id_code}", srv.TransGet)
	apiMux.HandleFunc("PUT /trans/{id_code}", srv.TransPut)
	apiMux.HandleFunc("DELETE /trans/{id_code}", srv.TransDelete)

	apiMux.HandleFunc("POST /service/database", srv.Database)
	apiMux.HandleFunc("POST /service/function", srv.Function)
	apiMux.HandleFunc("POST /service/view", srv.View)

	return http.StripPrefix(st.ApiPath, apiMux)
}

func (s *httpServer) clientUIRoutes() http.Handler {
	clientMux := http.NewServeMux()
	clientMux.HandleFunc("GET /", src.ClientAuth)
	clientMux.HandleFunc("GET /auth/", src.ClientAuth)
	clientMux.HandleFunc("POST /auth/event", src.ClientAuthEvent)

	clientMux.HandleFunc("GET /session/{session_id}", src.ClientSession)
	clientMux.HandleFunc("GET /session/export/browser/{session_id}", src.ClientExportBrowser)
	clientMux.HandleFunc("GET /session/export/report/modal/{session_id}", src.ClientExportModalReport)
	clientMux.HandleFunc("POST /session/event", src.ClientSessionEvent)

	return http.StripPrefix(st.ClientPath, clientMux)
}

func (s *httpServer) clientAPIRoutes() http.Handler {
	clientMux := http.NewServeMux()
	clientMux.HandleFunc("GET /auth/callback", src.ClientAuthCallback)
	clientMux.HandleFunc("POST /session", src.ClientSessionCreate)

	return http.StripPrefix(st.ClientPath+"/api", clientMux)
}
