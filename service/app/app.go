package app

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	db "github.com/nervatura/nervatura-service/pkg/database"
	nt "github.com/nervatura/nervatura-service/pkg/nervatura"
	srv "github.com/nervatura/nervatura-service/pkg/service"
	ut "github.com/nervatura/nervatura-service/pkg/utils"
	"golang.org/x/sync/errgroup"
)

// App - Nervatura Application
type App struct {
	services  map[string]srv.APIService
	defConn   nt.DataDriver
	infoLog   *log.Logger
	errorLog  *log.Logger
	httpLog   *log.Logger
	args      map[string]string
	tokenKeys map[string]map[string]string
	config    map[string]interface{}
}

var services = make(map[string]srv.APIService)

func registerService(name string, server srv.APIService) {
	services[name] = server
}

func New(version string, args nt.SM) (app *App, err error) {
	app = &App{
		config: nt.IM{
			"version": version,
		},
		args:      args,
		services:  services,
		tokenKeys: make(map[string]map[string]string),
	}

	app.infoLog = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	app.errorLog = log.New(os.Stdout, "ERROR: ", log.LstdFlags)
	app.httpLog = log.New(os.Stdout, "", log.LstdFlags)
	app.config["NT_APP_LOG_FILE"] = ut.ToString(args["NT_APP_LOG_FILE"], os.Getenv("NT_APP_LOG_FILE"))
	if app.config["NT_APP_LOG_FILE"] != "" {
		f, err := os.OpenFile(app.config["NT_APP_LOG_FILE"].(string), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			app.errorLog.Printf(ut.GetMessage("error_opening_log"), err)
		} else {
			app.infoLog = log.New(f, "INFO: ", log.LstdFlags)
			app.errorLog = log.New(f, "ERROR: ", log.LstdFlags)
		}
		defer func() {
			if def_err := f.Close(); def_err != nil {
				return
			}
		}()
	}
	app.setConfig()
	if app.config["NT_HTTP_LOG_FILE"] != "" {
		f, err := os.OpenFile(app.config["NT_HTTP_LOG_FILE"].(string), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			app.errorLog.Printf(ut.GetMessage("error_opening_log"), err)
		} else {
			app.httpLog = log.New(f, "", log.LstdFlags)
		}
		defer func() {
			if def_err := f.Close(); def_err != nil {
				return
			}
		}()
	}

	err = app.checkDefaultConn()
	if err != nil {
		app.errorLog.Printf(ut.GetMessage("error_checking_def_db"), err)
		return nil, err
	}

	err = app.setPrivateKey()
	if err != nil {
		app.errorLog.Printf(ut.GetMessage("error_private_key"), err)
		return nil, err
	}

	err = app.startService("cli")
	if err != nil {
		app.errorLog.Printf(ut.GetMessage("error_starting_cli"), err)
		return nil, err
	}

	if services["cli"].Results() == "server" {
		return app, app.startServer()
	}

	return app, err
}

func (app *App) setConfig() {
	app.config["NT_TLS_CERT_FILE"] = ut.ToString(os.Getenv("NT_TLS_CERT_FILE"), "")
	app.config["NT_TLS_KEY_FILE"] = ut.ToString(os.Getenv("NT_TLS_KEY_FILE"), "")

	app.config["NT_HTTP_ENABLED"] = ut.ToBoolean(app.args["NT_HTTP_ENABLED"], ut.ToBoolean(os.Getenv("NT_HTTP_ENABLED"), true))
	app.config["NT_HTTP_PORT"] = ut.ToInteger(os.Getenv("NT_HTTP_PORT"), 5000)
	app.config["NT_HTTP_TLS_ENABLED"] = ut.ToBoolean(os.Getenv("NT_HTTP_TLS_ENABLED"), false)
	app.config["NT_HTTP_READ_TIMEOUT"] = ut.ToFloat(os.Getenv("NT_HTTP_READ_TIMEOUT"), 30)
	app.config["NT_HTTP_WRITE_TIMEOUT"] = ut.ToFloat(os.Getenv("NT_HTTP_WRITE_TIMEOUT"), 30)
	app.config["NT_HTTP_HOME"] = ut.ToString(os.Getenv("NT_HTTP_HOME"), "/admin")
	app.config["NT_HTTP_LOG_FILE"] = ut.ToString(app.args["NT_APP_LOG_FILE"], os.Getenv("NT_HTTP_LOG_FILE"))
	if app.config["NT_HTTP_LOG_FILE"] == "" && os.Getenv("SNAP_COMMON") != "" && strings.Contains(os.Getenv("SNAP_COMMON"), "nervatura") {
		app.config["NT_HTTP_LOG_FILE"] = os.Getenv("SNAP_COMMON") + "/http.log"
	}

	app.config["NT_GRPC_ENABLED"] = ut.ToBoolean(app.args["NT_GRPC_ENABLED"], ut.ToBoolean(os.Getenv("NT_GRPC_ENABLED"), true))
	app.config["NT_GRPC_PORT"] = ut.ToInteger(os.Getenv("NT_GRPC_PORT"), 9200)
	app.config["NT_GRPC_TLS_ENABLED"] = ut.ToBoolean(os.Getenv("NT_GRPC_TLS_ENABLED"), false)

	app.config["NT_CLIENT_CONFIG"] = ut.ToString(os.Getenv("NT_CLIENT_CONFIG"), "")

	app.config["NT_FONT_FAMILY"] = ut.ToString(os.Getenv("NT_FONT_FAMILY"), "")
	app.config["NT_FONT_DIR"] = ut.ToString(os.Getenv("NT_FONT_DIR"), "")
	app.config["NT_REPORT_DIR"] = ut.ToString(os.Getenv("NT_REPORT_DIR"), "")

	if os.Getenv("NT_API_KEY") == "" && (app.config["version"] == "test" || app.config["version"] == "debug") {
		app.config["NT_API_KEY"] = "TEST_API_KEY"
	} else {
		app.config["NT_API_KEY"] = ut.ToString(os.Getenv("NT_API_KEY"), ut.RandString(32))
	}

	app.config["NT_TOKEN_ISS"] = ut.ToString(os.Getenv("NT_TOKEN_ISS"), "nervatura")
	app.config["NT_TOKEN_KID"] = ut.ToString(os.Getenv("NT_TOKEN_KID"), ut.GetHash("nervatura"))
	app.config["NT_TOKEN_PRIVATE_KEY_TYPE"] = ut.ToString(os.Getenv("NT_TOKEN_PRIVATE_KEY_TYPE"), "KEY")
	isServer := func() bool {
		if ut.Contains(os.Args, "-c") && !ut.Contains(os.Args, "server") {
			return true
		}
		if _, found := app.args["cmd"]; found {
			if app.args["cmd"] != "server" {
				return true
			}
		}
		return false
	}
	if isServer() {
		app.config["NT_TOKEN_PRIVATE_KEY"] = ut.ToString(os.Getenv("NT_TOKEN_PRIVATE_KEY"), ut.GetHash(time.Now().Format("20060102")))
	} else {
		app.config["NT_TOKEN_PRIVATE_KEY"] = ut.ToString(os.Getenv("NT_TOKEN_PRIVATE_KEY"), ut.RandString(32))
	}
	app.config["NT_TOKEN_EXP"] = ut.ToFloat(os.Getenv("NT_TOKEN_EXP"), 6)
	app.config["NT_TOKEN_PUBLIC_KEY_TYPE"] = ut.ToString(os.Getenv("NT_TOKEN_PUBLIC_KEY_TYPE"), "RSA")
	app.config["NT_TOKEN_PUBLIC_KEY_URL"] = ut.ToString(os.Getenv("NT_TOKEN_PUBLIC_KEY_URL"), "")

	app.config["NT_HASHTABLE"] = ut.ToString(os.Getenv("NT_HASHTABLE"), "ref17890714")

	app.config["NT_SMTP_HOST"] = ut.ToString(os.Getenv("NT_SMTP_HOST"), "")
	app.config["NT_SMTP_PORT"] = ut.ToInteger(os.Getenv("NT_SMTP_PORT"), 465)
	app.config["NT_SMTP_TLS_MIN_VERSION"] = ut.ToInteger(os.Getenv("NT_SMTP_TLS_MIN_VERSION"), 0)
	app.config["NT_SMTP_USER"] = ut.ToString(os.Getenv("NT_SMTP_USER"), "")
	app.config["NT_SMTP_PASSWORD"] = ut.ToString(os.Getenv("NT_SMTP_PASSWORD"), "")

	app.config["SQL_MAX_OPEN_CONNS"] = ut.ToInteger(os.Getenv("SQL_MAX_OPEN_CONNS"), 10)
	app.config["SQL_MAX_IDLE_CONNS"] = ut.ToInteger(os.Getenv("SQL_MAX_IDLE_CONNS"), 3)
	app.config["SQL_CONN_MAX_LIFETIME"] = ut.ToInteger(os.Getenv("SQL_CONN_MAX_LIFETIME"), 15)

	app.config["NT_ALIAS_DEFAULT"] = ut.ToString(app.args["NT_ALIAS_DEFAULT"], os.Getenv("NT_ALIAS_DEFAULT"))

	app.config["NT_CORS_ENABLED"] = ut.ToBoolean(os.Getenv("NT_CORS_ENABLED"), true)
	app.config["NT_CORS_ALLOW_ORIGINS"] = strings.Split(ut.ToString(os.Getenv("NT_CORS_ALLOW_ORIGINS"), "*"), ",")
	app.config["NT_CORS_ALLOW_METHODS"] = strings.Split(ut.ToString(os.Getenv("NT_CORS_ALLOW_METHODS"), "GET,POST,DELETE,OPTIONS"), ",")
	app.config["NT_CORS_ALLOW_HEADERS"] = strings.Split(ut.ToString(os.Getenv("NT_CORS_ALLOW_HEADERS"), "Accept,Authorization,Content-Type,X-CSRF-Token,X-Api-Key"), ",")
	app.config["NT_CORS_EXPOSE_HEADERS"] = strings.Split(ut.ToString(os.Getenv("NT_CORS_EXPOSE_HEADERS"), ""), ",")
	app.config["NT_CORS_ALLOW_CREDENTIALS"] = ut.ToBoolean(os.Getenv("NT_CORS_ALLOW_CREDENTIALS"), false)
	app.config["NT_CORS_MAX_AGE"] = ut.ToInteger(os.Getenv("NT_CORS_MAX_AGE"), 0)

	app.config["NT_SECURITY_ENABLED"] = ut.ToBoolean(os.Getenv("NT_SECURITY_ENABLED"), false)
	app.config["NT_SECURITY_ALLOWED_HOSTS"] = strings.Split(ut.ToString(os.Getenv("NT_SECURITY_ALLOWED_HOSTS"), ""), ",")
	app.config["NT_SECURITY_HOSTS_PROXY_HEADERS"] = strings.Split(ut.ToString(os.Getenv("NT_SECURITY_HOSTS_PROXY_HEADERS"), ""), ",")
	app.config["NT_SECURITY_ALLOWED_HOSTS_ARE_REGEX"] = ut.ToBoolean(os.Getenv("NT_SECURITY_ALLOWED_HOSTS_ARE_REGEX"), false)
	app.config["NT_SECURITY_SSL_REDIRECT"] = ut.ToBoolean(os.Getenv("NT_SECURITY_SSL_REDIRECT"), false)
	app.config["NT_SECURITY_SSL_TEMPORARY_REDIRECT"] = ut.ToBoolean(os.Getenv("NT_SECURITY_SSL_TEMPORARY_REDIRECT"), false)
	app.config["NT_SECURITY_SSL_HOST"] = ut.ToString(os.Getenv("NT_SECURITY_SSL_HOST"), "")
	app.config["NT_SECURITY_PROXY_HEADERS"] = strings.Split(ut.ToString(os.Getenv("NT_SECURITY_PROXY_HEADERS"), ""), ",")
	app.config["NT_SECURITY_STS_SECONDS"] = ut.ToInteger(os.Getenv("NT_SECURITY_STS_SECONDS"), 0)
	app.config["NT_SECURITY_STS_INCLUDE_SUBDOMAINS"] = ut.ToBoolean(os.Getenv("NT_SECURITY_STS_INCLUDE_SUBDOMAINS"), false)
	app.config["NT_SECURITY_STS_PRELOAD"] = ut.ToBoolean(os.Getenv("NT_SECURITY_STS_PRELOAD"), false)
	app.config["NT_SECURITY_FORCE_STS_HEADER"] = ut.ToBoolean(os.Getenv("NT_SECURITY_FORCE_STS_HEADER"), false)
	app.config["NT_SECURITY_FRAME_DENY"] = ut.ToBoolean(os.Getenv("NT_SECURITY_FRAME_DENY"), false)
	app.config["NT_SECURITY_CUSTOM_FRAME_OPTIONS_VALUE"] = ut.ToString(os.Getenv("NT_SECURITY_CUSTOM_FRAME_OPTIONS_VALUE"), "")
	app.config["NT_SECURITY_CONTENT_TYPE_NOSNIFF"] = ut.ToBoolean(os.Getenv("NT_SECURITY_CONTENT_TYPE_NOSNIFF"), false)
	app.config["NT_SECURITY_BROWSER_XSS_FILTER"] = ut.ToBoolean(os.Getenv("NT_SECURITY_BROWSER_XSS_FILTER"), false)
	app.config["NT_SECURITY_CONTENT_SECURITY_POLICY"] = ut.ToString(os.Getenv("NT_SECURITY_CONTENT_SECURITY_POLICY"), "")
	app.config["NT_SECURITY_PUBLIC_KEY"] = ut.ToString(os.Getenv("NT_SECURITY_PUBLIC_KEY"), "")
	app.config["NT_SECURITY_REFERRER_POLICY"] = ut.ToString(os.Getenv("NT_SECURITY_REFERRER_POLICY"), "")
	app.config["NT_SECURITY_FEATURE_POLICY"] = ut.ToString(os.Getenv("NT_SECURITY_FEATURE_POLICY"), "")
	app.config["NT_SECURITY_EXPECT_CT_HEADER"] = ut.ToString(os.Getenv("NT_SECURITY_EXPECT_CT_HEADER"), "")
	app.config["NT_SECURITY_DEVELOPMENT"] = ut.ToBoolean(os.Getenv("NT_SECURITY_DEVELOPMENT"), false)

	app.config["NT_ALIAS_DEMO"] = ut.ToString(os.Getenv("NT_ALIAS_DEMO"), "")
	if app.config["NT_ALIAS_DEMO"] == "" && ut.Contains(db.Drivers, "sqlite") {
		if _, err := os.Stat("data"); err == nil {
			app.config["NT_ALIAS_DEMO"] = "sqlite://file:data/demo.db?cache=shared&mode=rwc"
		}
		if os.Getenv("SNAP_COMMON") != "" && strings.Contains(os.Getenv("SNAP_COMMON"), "nervatura") {
			app.config["NT_ALIAS_DEMO"] = "sqlite://file:" + os.Getenv("SNAP_COMMON") + "/demo.db?cache=shared&mode=rwc"
		}
	}

	info := []string{"NT_API_KEY", "NT_TOKEN_KID", "NT_TOKEN_PRIVATE_KEY"}
	for i := 0; i < len(info); i++ {
		if os.Getenv(info[i]) == "" && app.args == nil {
			app.infoLog.Println(info[i] + ": " + app.config[info[i]].(string))
		}
	}
}

func (app *App) setPrivateKey() error {
	pkey := app.config["NT_TOKEN_PRIVATE_KEY"].(string)
	kid := app.config["NT_TOKEN_KID"].(string)
	ktype := app.config["NT_TOKEN_PRIVATE_KEY_TYPE"].(string)
	if pkey != "" {
		//file or key?
		if _, err := os.Stat(pkey); err == nil {
			content, err := ioutil.ReadFile(filepath.Clean(pkey))
			if err != nil {
				return err
			}
			pkey = string(content)
		}
		app.tokenKeys[kid] = nt.SM{
			"type":  "private",
			"ktype": ktype,
			"value": pkey,
		}
	}
	return nil
}

func (app *App) startServer() error {
	app.infoLog.Println(ut.GetMessage("skipping_cli"))
	app.infoLog.Printf(ut.GetMessage("enabled_drivers"), strings.Join(db.Drivers, ","))

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	httpDisabled := false
	if _, found := services["http"]; found && app.config["NT_HTTP_ENABLED"].(bool) {
		g.Go(func() error {
			return app.startService("http")
		})
	} else {
		httpDisabled = true
		app.infoLog.Println(ut.GetMessage("http_disabled"))
	}

	grpcDisabled := false
	if _, found := services["grpc"]; found && app.config["NT_GRPC_ENABLED"].(bool) {
		g.Go(func() error {
			return app.startService("grpc")
		})
	} else {
		grpcDisabled = true
		app.infoLog.Println(ut.GetMessage("grpc_disabled"))
	}

	if httpDisabled && grpcDisabled {
		return nil
	}

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	app.infoLog.Println(ut.GetMessage("shutdown_signal"))

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if _, found := services["http"]; found && app.config["NT_HTTP_ENABLED"].(bool) {
		_ = services["http"].StopService(shutdownCtx)
	}
	if _, found := services["grpc"]; found && app.config["NT_GRPC_ENABLED"].(bool) {
		_ = services["grpc"].StopService(nil)
	}

	err := g.Wait()
	if err != nil {
		app.errorLog.Printf(ut.GetMessage("application_error"), err)
		return err
	}
	return nil
}

func (app *App) startService(name string) error {
	services[name].ConnectApp(app)
	return services[name].StartService()
}

func (app *App) checkDefaultConn() (err error) {
	connStr := ""
	alias := ""
	if app.config["NT_ALIAS_DEFAULT"] != "" {
		connStr = os.Getenv("NT_ALIAS_" + strings.ToUpper(app.config["NT_ALIAS_DEFAULT"].(string)))
		alias = strings.ToLower(app.config["NT_ALIAS_DEFAULT"].(string))
	}
	if connStr != "" {
		app.defConn = &db.SQLDriver{Config: app.config}
		return app.defConn.CreateConnection(alias, connStr)
	}
	return nil
}

func (app *App) GetNervaStore(database string) (nstore *nt.NervaStore) {
	if app.defConn != nil {
		if app.defConn.Connection().Alias == database {
			return nt.New(app.defConn, app.config)
		}
	}
	return nt.New(&db.SQLDriver{Config: app.config}, app.config)
}

func (app *App) GetResults() string {
	return app.services["cli"].Results()
}
