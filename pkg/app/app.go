package app

import (
	"context"
	"errors"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"syscall"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	db "github.com/nervatura/nervatura/v6/pkg/driver"
	ht "github.com/nervatura/nervatura/v6/pkg/host"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
	"golang.org/x/sync/errgroup"
)

// App - Nervatura Application
type App struct {
	config     cu.IM
	hosts      map[string]ht.APIHost
	traySrv    trayService
	showTray   bool
	taskSecKey string
	appLogOut  io.Writer
	httpLogOut io.Writer
	appLog     *slog.Logger
	getEnv     func(key string) string
	readFile   func(name string) ([]byte, error)
	readAll    func(r io.Reader) ([]byte, error)
}

// trayService - interface for tray service
type trayService interface {
	Run(app *App, interrupt chan os.Signal, ctx context.Context, httpDisabled bool, onExit func())
}

// New - create new Nervatura application
func New(version string, args cu.SM) (app *App, err error) {
	app = &App{
		config: cu.IM{
			"version":   version,
			"args":      cu.ToSM(args, cu.SM{}),
			"tokenKeys": []cu.SM{},
		},
		hosts:      ht.Hosts,
		traySrv:    &systemTray{},
		showTray:   false,
		taskSecKey: cu.RandString(32),
		appLogOut:  os.Stdout,
		httpLogOut: os.Stdout,
		appLog:     slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		getEnv:     os.Getenv,
		readFile:   os.ReadFile,
		readAll:    io.ReadAll,
	}
	app.setEnv("./.env")

	app.config["NT_APP_LOG_FILE"] = cu.ToString(args["NT_APP_LOG_FILE"], app.getEnv("NT_APP_LOG_FILE"))
	if cu.ToString(app.config["NT_APP_LOG_FILE"], "") != "" {
		f, err := os.OpenFile(cu.ToString(app.config["NT_APP_LOG_FILE"], ""), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			app.setErrorLog("opening log file", err)
		} else {
			app.appLogOut = f
			app.appLog = slog.New(slog.NewJSONHandler(f, nil))
		}
		defer func() {
			if def_err := f.Close(); def_err != nil {
				return
			}
		}()
	}

	app.setConfig(app.isSnap())

	if cu.ToString(app.config["NT_HTTP_LOG_FILE"], "") != "" {
		f, err := os.OpenFile(cu.ToString(app.config["NT_HTTP_LOG_FILE"], ""), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			app.setErrorLog("opening log file", err)
		} else {
			app.httpLogOut = f
		}
		defer func() {
			if def_err := f.Close(); def_err != nil {
				return
			}
		}()
	}

	if err = app.startServer("cli", nil); err != nil {
		return nil, err
	}

	if app.hosts["cli"].Results() == "server" {
		return app, app.backgroundServer()
	}

	return app, nil
}

// setEnv - set environment variables
func (app *App) setEnv(defaultEnvFile string) {
	// Load env file if it exists
	if envMap, err := loadEnvFile(defaultEnvFile); err == nil {
		for key, value := range envMap {
			if _, exists := os.LookupEnv(key); !exists {
				os.Setenv(key, value)
			}
		}
	}
	// Load env file from args if it exists
	for index, arg := range os.Args[1:] {
		if arg == "-env" && len(os.Args[1:]) > index+1 {
			envFile := os.Args[1:][index+1]
			if envMap, err := loadEnvFile(envFile); err == nil {
				for key, value := range envMap {
					os.Setenv(key, value)
				}
			}
		}
		if arg == "-tray" {
			app.showTray = true
		}
	}
}

// loadEnvFile reads a .env file and returns a map of key/value pairs
func loadEnvFile(filename string) (cu.SM, error) {
	envMap := make(cu.SM)

	data, err := os.ReadFile(filename)
	if err != nil {
		return envMap, err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.SplitN(line, "=", 2)
		if line == "" || strings.HasPrefix(line, "#") || len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		value = strings.Trim(value, `"'`)

		envMap[key] = value
	}

	return envMap, nil
}

// setInfoLog - set info log
func (app *App) setInfoLog(message string, args ...any) {
	app.appLog.Info(message, args...)
}

// setErrorLog - set error log
func (app *App) setErrorLog(message string, args ...any) {
	app.appLog.Error(message, args...)
}

// setConfig - set configuration
func (app *App) setConfig(isSnap bool) {
	args := cu.ToSM(app.config["args"], cu.SM{})

	app.config["NT_GOOGLE_CLIENT_ID"] = cu.ToString(app.getEnv("NT_GOOGLE_CLIENT_ID"), "")
	app.config["NT_GOOGLE_CLIENT_SECRET"] = cu.ToString(app.getEnv("NT_GOOGLE_CLIENT_SECRET"), "")
	app.config["NT_FACEBOOK_CLIENT_ID"] = cu.ToString(app.getEnv("NT_FACEBOOK_CLIENT_ID"), "")
	app.config["NT_FACEBOOK_CLIENT_SECRET"] = cu.ToString(app.getEnv("NT_FACEBOOK_CLIENT_SECRET"), "")
	app.config["NT_GITHUB_CLIENT_ID"] = cu.ToString(app.getEnv("NT_GITHUB_CLIENT_ID"), "")
	app.config["NT_GITHUB_CLIENT_SECRET"] = cu.ToString(app.getEnv("NT_GITHUB_CLIENT_SECRET"), "")
	app.config["NT_MICROSOFT_CLIENT_ID"] = cu.ToString(app.getEnv("NT_MICROSOFT_CLIENT_ID"), "")
	app.config["NT_MICROSOFT_CLIENT_SECRET"] = cu.ToString(app.getEnv("NT_MICROSOFT_CLIENT_SECRET"), "")

	app.config["NT_API_KEY"] = cu.ToString(app.getEnv("NT_API_KEY"), cu.RandString(32))
	app.config["NT_AUTH_KEY"] = cu.ToString(app.getEnv("NT_AUTH_KEY"), cu.RandString(32))
	app.config["NT_RSA_PUBLIC"] = cu.ToString(app.getEnv("NT_RSA_PUBLIC"), "")

	app.config["NT_SESSION_EXP"] = cu.ToInteger(app.getEnv("NT_SESSION_EXP"), cu.ToInteger(st.DefaultConfig["session"]["exp"], 1))
	app.config["NT_SESSION_METHOD"] = cu.ToInteger(app.getEnv("NT_SESSION_METHOD"), cu.ToInteger(st.DefaultConfig["session"]["method"], 0))

	app.config["NT_TLS_CERT_FILE"] = cu.ToString(app.getEnv("NT_TLS_CERT_FILE"), cu.ToString(st.DefaultConfig["connection"]["tls_cert_file"], ""))
	app.config["NT_TLS_KEY_FILE"] = cu.ToString(app.getEnv("NT_TLS_KEY_FILE"), cu.ToString(st.DefaultConfig["connection"]["tls_key_file"], ""))
	app.config["NT_DEFAULT_ALIAS"] = cu.ToString(app.getEnv("NT_DEFAULT_ALIAS"), cu.ToString(st.DefaultConfig["connection"]["default_alias"], ""))
	app.config["NT_DEFAULT_ADMIN"] = cu.ToString(app.getEnv("NT_DEFAULT_ADMIN"), cu.ToString(st.DefaultConfig["connection"]["default_admin"], ""))

	app.config["NT_HTTP_ENABLED"] = cu.ToBoolean(args["NT_HTTP_ENABLED"], cu.ToBoolean(app.getEnv("NT_HTTP_ENABLED"), cu.ToBoolean(st.DefaultConfig["http"]["http_enabled"], true)))
	app.config["NT_HTTP_PORT"] = cu.ToInteger(app.getEnv("NT_HTTP_PORT"), cu.ToInteger(st.DefaultConfig["http"]["port"], 5000))
	app.config["NT_HTTP_TLS_ENABLED"] = cu.ToBoolean(app.getEnv("NT_HTTP_TLS_ENABLED"), cu.ToBoolean(st.DefaultConfig["http"]["tls_enabled"], false))
	app.config["NT_HTTP_READ_TIMEOUT"] = cu.ToFloat(app.getEnv("NT_HTTP_READ_TIMEOUT"), cu.ToFloat(st.DefaultConfig["http"]["read_timeout"], 30))
	app.config["NT_HTTP_WRITE_TIMEOUT"] = cu.ToFloat(app.getEnv("NT_HTTP_WRITE_TIMEOUT"), cu.ToFloat(st.DefaultConfig["http"]["write_timeout"], 30))
	app.config["NT_HTTP_HOME"] = cu.ToString(app.getEnv("NT_HTTP_HOME"), cu.ToString(st.DefaultConfig["http"]["home"], "/"))
	app.config["NT_HTTP_LOG_FILE"] = cu.ToString(args["NT_HTTP_LOG_FILE"], app.getEnv("NT_HTTP_LOG_FILE"))

	dataDir := "data"
	if isSnap {
		dataDir = "/var/snap/nervatura/common"
		if app.config["NT_HTTP_LOG_FILE"] == "" {
			app.config["NT_HTTP_LOG_FILE"] = dataDir + "/http.log"
		}
	}

	app.config["NT_GRPC_ENABLED"] = cu.ToBoolean(args["NT_GRPC_ENABLED"], cu.ToBoolean(app.getEnv("NT_GRPC_ENABLED"), cu.ToBoolean(st.DefaultConfig["grpc"]["enabled"], true)))
	app.config["NT_GRPC_PORT"] = cu.ToInteger(app.getEnv("NT_GRPC_PORT"), cu.ToInteger(st.DefaultConfig["grpc"]["port"], 9200))
	app.config["NT_GRPC_TLS_ENABLED"] = cu.ToBoolean(app.getEnv("NT_GRPC_TLS_ENABLED"), cu.ToBoolean(st.DefaultConfig["grpc"]["tls_enabled"], false))

	app.config["NT_REPORT_FONT_FAMILY"] = cu.ToString(app.getEnv("NT_REPORT_FONT_FAMILY"), cu.ToString(st.DefaultConfig["report"]["font_family"], ""))
	app.config["NT_REPORT_FONT_DIR"] = cu.ToString(app.getEnv("NT_REPORT_FONT_DIR"), cu.ToString(st.DefaultConfig["report"]["font_dir"], ""))
	app.config["NT_REPORT_DIR"] = cu.ToString(app.getEnv("NT_REPORT_DIR"), cu.ToString(st.DefaultConfig["report"]["dir"], ""))

	app.config["NT_SESSION_ALIAS"] = cu.ToString(app.getEnv("NT_SESSION_ALIAS"), cu.ToString(st.DefaultConfig["connection"]["default_alias"], ""))
	app.config["NT_SESSION_TABLE"] = cu.ToString(app.getEnv("NT_SESSION_TABLE"), cu.ToString(st.DefaultConfig["session"]["table"], ""))
	app.config["NT_SESSION_DIR"] = cu.ToString(app.getEnv("NT_SESSION_DIR"), cu.ToString(st.DefaultConfig["session"]["file_dir"], ""))

	app.config["SQL_MAX_OPEN_CONNS"] = cu.ToInteger(app.getEnv("SQL_MAX_OPEN_CONNS"), cu.ToInteger(st.DefaultConfig["sql"]["max_open_conns"], 10))
	app.config["SQL_MAX_IDLE_CONNS"] = cu.ToInteger(app.getEnv("SQL_MAX_IDLE_CONNS"), cu.ToInteger(st.DefaultConfig["sql"]["max_idle_conns"], 3))
	app.config["SQL_CONN_MAX_LIFETIME"] = cu.ToInteger(app.getEnv("SQL_CONN_MAX_LIFETIME"), cu.ToInteger(st.DefaultConfig["sql"]["conn_max_lifetime"], 15))

	app.config["NT_ALIAS_DEMO"] = cu.ToString(app.getEnv("NT_ALIAS_DEMO"), "")
	if app.config["NT_ALIAS_DEMO"] == "" && slices.Contains(db.Drivers, "sqlite") {
		if _, err := os.Stat("data"); err == nil || isSnap {
			app.config["NT_ALIAS_DEMO"] = "sqlite://file:" + dataDir + "/demo.db?cache=shared&mode=rwc"
		}
	}

	app.config["NT_TOKEN_ISS"] = cu.ToString(app.getEnv("NT_TOKEN_ISS"), cu.ToString(st.DefaultConfig["token"]["iss"], "nervatura"))
	app.config["NT_TOKEN_PRIVATE_KID"] = cu.ToString(app.getEnv("NT_TOKEN_PRIVATE_KID"), ut.GetHash("nervatura", "sha256"))
	isServer := func() bool {
		if _, found := args["cmd"]; found {
			if args["cmd"] != "server" {
				return true
			}
		}
		return (slices.Contains(os.Args, "-c") && !slices.Contains(os.Args, "server"))
	}
	if isServer() {
		app.config["NT_TOKEN_PRIVATE_KEY"] = cu.ToString(app.getEnv("NT_TOKEN_PRIVATE_KEY"), ut.GetHash(cu.RandString(16), "sha256"))
	} else {
		app.config["NT_TOKEN_PRIVATE_KEY"] = cu.ToString(app.getEnv("NT_TOKEN_PRIVATE_KEY"), cu.RandString(32))
	}
	app.config["NT_TOKEN_EXP"] = cu.ToFloat(app.getEnv("NT_TOKEN_EXP"), cu.ToFloat(st.DefaultConfig["token"]["exp"], 6))

	app.config["NT_TOKEN_PUBLIC_KEY"] = cu.ToString(app.getEnv("NT_TOKEN_PUBLIC_KEY"), "")
	app.config["NT_TOKEN_PUBLIC_KEY_URL"] = cu.ToString(app.getEnv("NT_TOKEN_PUBLIC_KEY_URL"), cu.ToString(st.DefaultConfig["token"]["public_key_url"], ""))

	app.config["NT_CORS_ENABLED"] = cu.ToBoolean(app.getEnv("NT_CORS_ENABLED"), true)
	app.config["NT_CORS_ALLOW_ORIGINS"] = strings.Split(cu.ToString(app.getEnv("NT_CORS_ALLOW_ORIGINS"), st.DefaultConfig["cors"]["allow_origins"]), ",")
	app.config["NT_CORS_ALLOW_METHODS"] = strings.Split(cu.ToString(app.getEnv("NT_CORS_ALLOW_METHODS"), st.DefaultConfig["cors"]["allow_methods"]), ",")
	app.config["NT_CORS_ALLOW_HEADERS"] = strings.Split(cu.ToString(app.getEnv("NT_CORS_ALLOW_HEADERS"), st.DefaultConfig["cors"]["allow_headers"]), ",")
	app.config["NT_CORS_EXPOSE_HEADERS"] = strings.Split(cu.ToString(app.getEnv("NT_CORS_EXPOSE_HEADERS"), st.DefaultConfig["cors"]["expose_headers"]), ",")
	app.config["NT_CORS_ALLOW_CREDENTIALS"] = cu.ToBoolean(app.getEnv("NT_CORS_ALLOW_CREDENTIALS"), false)
	app.config["NT_CORS_MAX_AGE"] = cu.ToInteger(app.getEnv("NT_CORS_MAX_AGE"), 0)
	app.config["NT_CSRF_TRUSTED_ORIGINS"] = strings.Split(cu.ToString(app.getEnv("NT_CSRF_TRUSTED_ORIGINS"), st.DefaultConfig["cors"]["trusted_origins"]), ",")

	app.config["NT_SMTP_HOST"] = cu.ToString(app.getEnv("NT_SMTP_HOST"), cu.ToString(st.DefaultConfig["smtp"]["host"], ""))
	app.config["NT_SMTP_PORT"] = cu.ToInteger(app.getEnv("NT_SMTP_PORT"), cu.ToInteger(st.DefaultConfig["smtp"]["port"], 465))
	app.config["NT_SMTP_TLS_MIN_VERSION"] = cu.ToInteger(app.getEnv("NT_SMTP_TLS_MIN_VERSION"), cu.ToInteger(st.DefaultConfig["smtp"]["tls_min_version"], 0))
	app.config["NT_SMTP_USER"] = cu.ToString(app.getEnv("NT_SMTP_USER"), cu.ToString(st.DefaultConfig["smtp"]["user"], ""))
	app.config["NT_SMTP_PASSWORD"] = cu.ToString(app.getEnv("NT_SMTP_PASSWORD"), "")

	info := []string{"NT_API_KEY", "NT_TOKEN_PRIVATE_KID", "NT_TOKEN_PRIVATE_KEY"}
	for i := 0; i < len(info); i++ {
		if app.getEnv(info[i]) == "" && len(args) == 0 {
			log.Println(info[i] + ": " + app.config[info[i]].(string))
		}
	}
}

// isDocker - check if running in docker
func (app *App) isDocker() bool {
	_, err := os.Stat("/.dockerenv")
	return (err == nil)
}

// isSnap - check if running in snap
func (app *App) isSnap() bool {
	current, _ := os.Executable()
	return strings.Contains(current, "snap/nervatura")
}

// setTokenKeys - set/load private and public token keys from file or environment variable
func (app *App) setTokenKeys(keyType string) error {
	pkey := cu.ToString(app.config["NT_TOKEN_"+strings.ToUpper(keyType)+"_KEY"], "")
	alg := cu.ToString(app.config["NT_TOKEN_ALG"], "HS256")
	algType := ut.TokenAlg[alg]
	if pkey != "" {
		//file or key?
		if _, err := os.Stat(pkey); err == nil {
			content, err := app.readFile(filepath.Clean(pkey))
			if err != nil {
				app.setErrorLog("reading token key file", err)
				return err
			}
			pkey = string(content)
		}
		if algType == "HMAC" || keyType == "public" {
			app.config["tokenKeys"] = append(app.config["tokenKeys"].([]cu.SM), cu.SM{
				"type":  keyType,
				"value": pkey,
			})
		}
		app.config["NT_TOKEN_"+strings.ToUpper(keyType)+"_KEY"] = pkey
	}
	return nil
}

// setPublicTokenURLKeys - set/load public token keys from URL
func (app *App) setPublicTokenURLKeys() {
	publicUrl := cu.ToString(app.config["NT_TOKEN_PUBLIC_KEY_URL"], "")
	if publicUrl != "" {
		res, err := http.Get(publicUrl)
		if err != nil {
			app.setErrorLog("reading external token", err)
			return
		}
		defer res.Body.Close()
		data, err := app.readAll(res.Body)
		if err != nil {
			app.setErrorLog("reading external token", err)
			return
		}
		app.config["tokenKeys"] = append(app.config["tokenKeys"].([]cu.SM), cu.SM{
			"type":  "public",
			"value": string(data),
		})
	}
}

// setTokenKeyRing - set token key ring
func (app *App) setTokenKeyRing() (err error) {
	for _, tv := range []string{"private", "public"} {
		if err = app.setTokenKeys(tv); err != nil {
			return err
		}
	}
	app.setPublicTokenURLKeys()
	return nil
}

// GetResults - get results
func (app *App) GetResults() string {
	return app.hosts["cli"].Results()
}

// openURL - open URL from tray menu
func (app *App) openURL(goOS, urlStr string) error {
	var cmd *exec.Cmd
	switch goOS {
	case "darwin":
		cmd = exec.Command("open", urlStr)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", urlStr)
	default:
		cmd = exec.Command("xdg-open", urlStr)
	}
	if app.hosts != nil {
		return cmd.Start()
	}
	return errors.New("internal error")
}

// onTrayMenu - handle tray menu click
func (app *App) onTrayMenu(mKey string) {
	var mURL string
	switch mKey {
	case "config":
		app.config["NT_TASK_SEC_KEY"] = cu.RandString(32)
		mURL = "http://localhost:" + cu.ToString(app.config["NT_HTTP_PORT"], "") + "/config/" + cu.ToString(app.config["NT_TASK_SEC_KEY"], "")
	case "admin":
		mURL = "http://localhost:" + cu.ToString(app.config["NT_HTTP_PORT"], "") + "/"
	}
	if err := app.openURL(runtime.GOOS, mURL); err != nil {
		app.setErrorLog("opening URL", err)
	}
}

// startServer - start server
func (app *App) startServer(name string, interrupt chan os.Signal) error {
	return app.hosts[name].StartServer(app.config, app.appLogOut, app.httpLogOut, interrupt)
}

// backgroundServer - start background http and/or grpc server
func (app *App) backgroundServer() error {
	app.setInfoLog("skipping cli")
	app.setInfoLog("enabled drivers", "drivers", strings.Join(db.Drivers, ","))

	if err := app.setTokenKeyRing(); err != nil {
		app.setErrorLog("setting token key ring", err)
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	// Setup HTTP server
	httpDisabled, configURL := app.setupHTTPServer(g, interrupt)

	// Setup gRPC server
	grpcDisabled := app.setupGRPCServer(g, interrupt)

	if httpDisabled && grpcDisabled {
		return nil
	}

	onExit := app.createExitHandler(cancel, g)

	return app.runServer(httpDisabled, configURL, interrupt, ctx, onExit)
}

func (app *App) setupHTTPServer(g *errgroup.Group, interrupt chan os.Signal) (bool, string) {
	httpDisabled := false
	configURL := "http://localhost:" + cu.ToString(app.config["NT_HTTP_PORT"], "") + "/admin/task/config/" + app.taskSecKey

	if _, found := app.hosts["http"]; found && cu.ToBoolean(app.config["NT_HTTP_ENABLED"], false) {
		g.Go(func() error {
			return app.startServer("http", interrupt)
		})
	} else {
		httpDisabled = true
		configURL = "http disabled"
		app.setInfoLog(configURL)
	}
	return httpDisabled, configURL
}

func (app *App) setupGRPCServer(g *errgroup.Group, interrupt chan os.Signal) bool {
	grpcDisabled := false
	if _, found := app.hosts["grpc"]; found && cu.ToBoolean(app.config["NT_GRPC_ENABLED"], false) {
		g.Go(func() error {
			return app.startServer("grpc", interrupt)
		})
	} else {
		grpcDisabled = true
		app.setInfoLog("grpc disabled")
	}
	return grpcDisabled
}

func (app *App) createExitHandler(cancel context.CancelFunc, g *errgroup.Group) func() {
	return func() {
		app.setInfoLog("shutdown signal")
		cancel()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		if _, found := app.hosts["http"]; found && cu.ToBoolean(app.config["NT_HTTP_ENABLED"], false) {
			_ = app.hosts["http"].StopServer(shutdownCtx)
		}
		if _, found := app.hosts["grpc"]; found && cu.ToBoolean(app.config["NT_GRPC_ENABLED"], false) {
			_ = app.hosts["grpc"].StopServer(shutdownCtx)
		}

		_ = g.Wait()
	}
}

func (app *App) runServer(httpDisabled bool, configURL string, interrupt chan os.Signal, ctx context.Context, onExit func()) error {
	trayIcon := app.showTray && !app.isDocker() && (app.traySrv != nil)
	if trayIcon {
		app.traySrv.Run(app, interrupt, ctx, httpDisabled, onExit)
	} else {
		log.Println("configuration values", "url", configURL)
		select {
		case <-interrupt:
		case <-ctx.Done():
		}
		onExit()
	}
	return nil
}
