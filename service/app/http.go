//go:build http || all
// +build http all

package app

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	ct "github.com/nervatura/component/pkg/static"
	db "github.com/nervatura/nervatura/service/pkg/database"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	srv "github.com/nervatura/nervatura/service/pkg/service"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
	"github.com/unrolled/secure"
)

type httpServer struct {
	app        *App
	mux        *chi.Mux
	service    srv.HTTPService
	admin      srv.AdminService
	result     string
	server     *http.Server
	tlsEnabled bool
	readAll    func(r io.Reader) ([]byte, error)
	tokenLogin func(r *http.Request) (ctx context.Context, err error)
}

func init() {
	registerService("http", &httpServer{
		readAll: io.ReadAll,
	})
}

func (s *httpServer) StartService() error {
	s.mux = chi.NewRouter()
	s.service = srv.HTTPService{
		Config:        s.app.config,
		GetNervaStore: s.app.GetNervaStore,
		GetParam:      chi.URLParam,
		GetTokenKeys:  s.app.GetTokenKeys,
	}
	s.tokenLogin = s.service.TokenLogin

	s.admin = srv.AdminService{
		Config:          s.app.config,
		GetNervaStore:   s.app.GetNervaStore,
		GetParam:        chi.URLParam,
		GetTokenKeys:    s.app.GetTokenKeys,
		GetTaskSecKey:   s.app.GetTaskSecKey,
		ReadFile:        os.ReadFile,
		ConvertFromByte: ut.ConvertFromByte,
		CreateFile:      os.Create,
		ConvertToWriter: ut.ConvertToWriter,
		Session: nt.SessionService{
			Config:          s.app.config,
			Conn:            &db.SQLDriver{Config: s.app.config},
			CreateDir:       os.Mkdir,
			CreateFile:      os.Create,
			ReadFile:        os.ReadFile,
			FileStat:        os.Stat,
			ConvertToByte:   ut.ConvertToByte,
			ConvertFromByte: ut.ConvertFromByte,
			RemoveFile:      os.Remove,
		},
	}

	s.setPublicKeys()
	s.setMiddleware()
	s.setRoutes()

	// Start API server.
	return s.startServer()
}

func (s *httpServer) setPublicKeys() {
	publicUrl := s.app.config["NT_TOKEN_PUBLIC_KEY_URL"].(string)
	if publicUrl != "" {
		res, err := http.Get(publicUrl)
		if err != nil {
			s.app.errorLog.Printf(ut.GetMessage("error_external_token"), err)
			return
		}
		defer res.Body.Close()
		data, err := s.readAll(res.Body)
		if err != nil {
			s.app.errorLog.Printf(ut.GetMessage("error_external_token"), err)
			return
		}
		var tokenKeys map[string]string
		err = ut.ConvertFromByte(data, &tokenKeys)
		if err != nil {
			s.app.errorLog.Printf(ut.GetMessage("error_external_token"), err)
		}
		for key, value := range tokenKeys {
			s.app.tokenKeys[key] = map[string]string{
				"type":  "public",
				"value": value,
			}
		}
	}
}

func (s *httpServer) startServer() error {
	s.server = &http.Server{
		Handler:      s.mux,
		Addr:         fmt.Sprintf(":%d", s.app.config["NT_HTTP_PORT"].(int64)),
		ReadTimeout:  time.Duration(s.app.config["NT_HTTP_READ_TIMEOUT"].(float64)) * time.Second,
		WriteTimeout: time.Duration(s.app.config["NT_HTTP_WRITE_TIMEOUT"].(float64)) * time.Second,
	}
	s.tlsEnabled = s.app.config["NT_HTTP_TLS_ENABLED"].(bool) &&
		s.app.config["NT_TLS_CERT_FILE"] != "" && s.app.config["NT_TLS_KEY_FILE"] != ""

	s.app.infoLog.Printf(ut.GetMessage("http_serving"), s.app.config["NT_HTTP_PORT"].(int64), s.tlsEnabled)
	if s.tlsEnabled {
		return s.server.ListenAndServeTLS(s.app.config["NT_TLS_CERT_FILE"].(string), s.app.config["NT_TLS_KEY_FILE"].(string))
	}
	return s.server.ListenAndServe()
}

func (s *httpServer) StopService(ctx interface{}) error {
	if s.server != nil {
		s.app.infoLog.Println(ut.GetMessage("http_stopping"))
		return s.server.Shutdown(ctx.(context.Context))
	}
	return nil
}

func (s *httpServer) Results() string {
	return s.result
}

func (s *httpServer) ConnectApp(app interface{}) {
	s.app = app.(*App)
}

func (s *httpServer) Logger(next http.Handler) http.Handler {
	color := !(runtime.GOOS == "windows")
	DefaultLogger := middleware.RequestLogger(
		&middleware.DefaultLogFormatter{Logger: s.app.httpLog, NoColor: !color})
	return DefaultLogger(next)
}

// Register middleware.
func (s *httpServer) setMiddleware() {

	s.mux.Use(s.Logger)
	s.mux.Use(middleware.RequestID)
	s.mux.Use(middleware.Recoverer)

	s.mux.Use(middleware.CleanPath)
	s.mux.Use(middleware.StripSlashes)
	s.mux.Use(middleware.NoCache)

	if s.app.config["NT_CORS_ENABLED"].(bool) {
		s.mux.Use(cors.Handler(cors.Options{
			AllowedOrigins:   s.app.config["NT_CORS_ALLOW_ORIGINS"].([]string),
			AllowedMethods:   s.app.config["NT_CORS_ALLOW_METHODS"].([]string),
			AllowedHeaders:   s.app.config["NT_CORS_ALLOW_HEADERS"].([]string),
			ExposedHeaders:   s.app.config["NT_CORS_EXPOSE_HEADERS"].([]string),
			AllowCredentials: s.app.config["NT_CORS_ALLOW_CREDENTIALS"].(bool),
			MaxAge:           int(s.app.config["NT_CORS_MAX_AGE"].(int64)),
		}))
	}

	if s.app.config["NT_SECURITY_ENABLED"].(bool) {
		s.mux.Use(secure.New(secure.Options{
			AllowedHosts:            s.app.config["NT_SECURITY_ALLOWED_HOSTS"].([]string),
			AllowedHostsAreRegex:    s.app.config["NT_SECURITY_ALLOWED_HOSTS_ARE_REGEX"].(bool),
			HostsProxyHeaders:       s.app.config["NT_SECURITY_HOSTS_PROXY_HEADERS"].([]string),
			SSLRedirect:             s.app.config["NT_SECURITY_SSL_REDIRECT"].(bool),
			SSLTemporaryRedirect:    s.app.config["NT_SECURITY_SSL_TEMPORARY_REDIRECT"].(bool),
			SSLHost:                 s.app.config["NT_SECURITY_SSL_HOST"].(string),
			STSSeconds:              s.app.config["NT_SECURITY_STS_SECONDS"].(int64),
			STSIncludeSubdomains:    s.app.config["NT_SECURITY_STS_INCLUDE_SUBDOMAINS"].(bool),
			STSPreload:              s.app.config["NT_SECURITY_STS_PRELOAD"].(bool),
			ForceSTSHeader:          s.app.config["NT_SECURITY_FORCE_STS_HEADER"].(bool),
			FrameDeny:               s.app.config["NT_SECURITY_FRAME_DENY"].(bool),
			CustomFrameOptionsValue: s.app.config["NT_SECURITY_CUSTOM_FRAME_OPTIONS_VALUE"].(string),
			ContentTypeNosniff:      s.app.config["NT_SECURITY_CONTENT_TYPE_NOSNIFF"].(bool),
			BrowserXssFilter:        s.app.config["NT_SECURITY_BROWSER_XSS_FILTER"].(bool),
			ContentSecurityPolicy:   s.app.config["NT_SECURITY_CONTENT_SECURITY_POLICY"].(string),
			//PublicKey:               s.app.config["NT_SECURITY_PUBLIC_KEY"].(string),
			ReferrerPolicy: s.app.config["NT_SECURITY_REFERRER_POLICY"].(string),
			FeaturePolicy:  s.app.config["NT_SECURITY_FEATURE_POLICY"].(string),
			//ExpectCTHeader:          s.app.config["NT_SECURITY_EXPECT_CT_HEADER"].(string),
			IsDevelopment: s.app.config["NT_SECURITY_DEVELOPMENT"].(bool),
		}).Handler)
	}

}

func (s *httpServer) tokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := s.tokenLogin(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *httpServer) homeRoute(w http.ResponseWriter, r *http.Request) {
	home := s.app.config["NT_HTTP_HOME"].(string)
	if home != "/" {
		http.Redirect(w, r, home, http.StatusSeeOther)
	}
}

// Register API routes.
func (s *httpServer) setRoutes() {
	// Register static dirs.
	var publicFS, _ = fs.Sub(ut.Public, "static")
	var adminFS, _ = fs.Sub(ct.Static, ".")
	s.fileServer("/", http.FS(publicFS))
	s.fileServer("/static", http.FS(adminFS))

	s.mux.Get("/", s.homeRoute)

	s.mux.Route("/admin", func(r chi.Router) {
		r.Get("/", s.admin.Home)
		r.Post("/event", s.admin.AppEvent)
		r.Get("/task/{taskName}/{secKey}", s.admin.Task)
	})
	s.mux.Route("/api", func(r chi.Router) {
		r.Post("/database", s.service.DatabaseCreate)
		r.Get("/config", s.service.ClientConfig)
		r.Group(func(r chi.Router) {
			r.Use(s.tokenAuth)
			r.Post("/function", s.service.Function)
			r.Post("/view", s.service.View)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", s.service.UserLogin)
			r.Group(func(r chi.Router) {
				r.Use(s.tokenAuth)
				r.Post("/password", s.service.UserPassword)
				r.Get("/refresh", s.service.TokenRefresh)
				r.Get("/validate", s.service.TokenValidate)
			})
		})

		r.Route("/{nervatype}", func(r chi.Router) {
			r.Use(s.tokenAuth)
			r.Get("/", s.service.GetFilter)
			r.Get("/{IDs}", s.service.GetIds)
			r.Post("/", s.service.Update)
			r.Delete("/", s.service.Delete)
		})

		r.Route("/report", func(r chi.Router) {
			r.Use(s.tokenAuth)
			r.Get("/", s.service.Report)
			r.Post("/", s.service.Report)
			r.Get("/list", s.service.ReportList)
			r.Post("/install", s.service.ReportInstall)
			r.Delete("/delete", s.service.ReportDelete)
		})

	})

}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func (s *httpServer) fileServer(path string, root http.FileSystem) {

	if strings.ContainsAny(path, "{}*") {
		s.app.errorLog.Println(ut.GetMessage("error_fileserver"))
		return
	}

	if path != "/" && path[len(path)-1] != '/' {
		s.mux.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	s.mux.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
