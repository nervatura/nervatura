package app

import (
	"errors"
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	db "github.com/nervatura/nervatura/service/pkg/database"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	srv "github.com/nervatura/nervatura/service/pkg/service"
)

const testDatabase = "sqlite://file::memory:?cache=shared"

// const testDatabase = "postgres://postgres:admin@172.25.0.1:5432/nervatura?sslmode=disable"

type testServer struct {
	app    *App
	result string
}

func (ts *testServer) StartService() error {
	return nil
}
func (ts *testServer) Results() string {
	return ts.result
}
func (ts *testServer) ConnectApp(app interface{}) {
	ts.app = app.(*App)
}
func (ts *testServer) StopService(interface{}) error {
	return nil
}

func TestNew(t *testing.T) {
	type args struct {
		version string
		args    nt.SM
	}
	os.Setenv("NT_ALIAS_TEST", "sqlite5://file:../data/empty.db")
	os.Setenv("SNAP_COMMON", "/nervatura")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "cli_app",
			args: args{
				version: "test",
				args: nt.SM{
					"cmd":              "help",
					"NT_APP_LOG_FILE":  "../data/nervatura.log",
					"NT_HTTP_LOG_FILE": "../data/http.log",
				},
			},
			wantErr: false,
		},
		{
			name: "cli_error",
			args: args{
				version: "dev",
				args: nt.SM{
					"NT_APP_LOG_FILE":  "data/nervatura.log",
					"NT_HTTP_LOG_FILE": "data/http.log",
				},
			},
			wantErr: true,
		},
		{
			name: "def_conn_error",
			args: args{
				version: "test",
				args: nt.SM{
					"NT_ALIAS_DEFAULT": "TEST",
				},
			},
			wantErr: true,
		},
		{
			name: "server_start",
			args: args{
				version: "test",
				args: nt.SM{
					"cmd":             "server",
					"NT_HTTP_ENABLED": "false",
					"NT_GRPC_ENABLED": "false",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.args.version, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestApp_setConfig(t *testing.T) {
	type fields struct {
		services   map[string]srv.APIService
		defConn    nt.DataDriver
		infoLog    *log.Logger
		errorLog   *log.Logger
		httpLog    *log.Logger
		args       nt.SM
		tokenKeys  map[string]nt.SM
		config     map[string]interface{}
		readFile   func(name string) ([]byte, error)
		getEnv     func(key string) string
		tray       bool
		taskSecKey string
	}
	type args struct {
		isSnap bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			args: args{
				isSnap: false,
			},
			fields: fields{
				config:  nt.IM{},
				infoLog: log.New(os.Stdout, "INFO: ", log.LstdFlags),
				getEnv: func(key string) string {
					return ""
				},
			},
		},
		{
			name: "snap",
			args: args{
				isSnap: true,
			},
			fields: fields{
				config:  nt.IM{},
				infoLog: log.New(os.Stdout, "INFO: ", log.LstdFlags),
				getEnv: func(key string) string {
					return ""
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:   tt.fields.services,
				defConn:    tt.fields.defConn,
				infoLog:    tt.fields.infoLog,
				errorLog:   tt.fields.errorLog,
				httpLog:    tt.fields.httpLog,
				args:       tt.fields.args,
				tokenKeys:  tt.fields.tokenKeys,
				config:     tt.fields.config,
				readFile:   tt.fields.readFile,
				getEnv:     tt.fields.getEnv,
				tray:       tt.fields.tray,
				taskSecKey: tt.fields.taskSecKey,
			}
			app.setConfig(tt.args.isSnap)
		})
	}
}

func TestApp_setTokenKey(t *testing.T) {
	type fields struct {
		services  map[string]srv.APIService
		defConn   nt.DataDriver
		infoLog   *log.Logger
		errorLog  *log.Logger
		httpLog   *log.Logger
		args      map[string]string
		tokenKeys map[string]map[string]string
		config    map[string]interface{}
		readFile  func(name string) ([]byte, error)
	}
	type args struct {
		keyType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "private_key_file",
			fields: fields{
				tokenKeys: map[string]map[string]string{},
				config: nt.IM{
					"NT_TOKEN_PRIVATE_KEY": "../data/x509/server_key.pem",
					"NT_TOKEN_PRIVATE_KID": "",
				},
				readFile: func(name string) ([]byte, error) {
					return []byte("ok"), nil
				},
			},
			args: args{
				keyType: "private",
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				tokenKeys: map[string]map[string]string{},
				config: nt.IM{
					"NT_TOKEN_PRIVATE_KEY": "../data/x509/server_key.pem",
					"NT_TOKEN_PRIVATE_KID": "",
				},
				readFile: func(name string) ([]byte, error) {
					return nil, errors.New("error")
				},
				errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
			},
			args: args{
				keyType: "private",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:  tt.fields.services,
				defConn:   tt.fields.defConn,
				infoLog:   tt.fields.infoLog,
				errorLog:  tt.fields.errorLog,
				httpLog:   tt.fields.httpLog,
				args:      tt.fields.args,
				tokenKeys: tt.fields.tokenKeys,
				config:    tt.fields.config,
				readFile:  tt.fields.readFile,
			}
			if err := app.setTokenKey(tt.args.keyType); (err != nil) != tt.wantErr {
				t.Errorf("App.setTokenKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_startServer(t *testing.T) {
	type fields struct {
		services   map[string]srv.APIService
		defConn    nt.DataDriver
		infoLog    *log.Logger
		errorLog   *log.Logger
		httpLog    *log.Logger
		args       map[string]string
		tokenKeys  map[string]map[string]string
		config     map[string]interface{}
		tray       bool
		taskSecKey string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "disabled",
			fields: fields{
				config: nt.IM{
					"NT_HTTP_ENABLED": false,
					"NT_GRPC_ENABLED": false,
				},
				services: make(map[string]srv.APIService),
				infoLog:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
				errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
			},
			wantErr: false,
		},
		{
			name: "http_server",
			fields: fields{
				config: nt.IM{
					"version":                 "test",
					"NT_GRPC_ENABLED":         true,
					"NT_GRPC_TLS_ENABLED":     false,
					"NT_GRPC_PORT":            int64(-1),
					"NT_HTTP_ENABLED":         true,
					"NT_HTTP_TLS_ENABLED":     false,
					"NT_HTTP_PORT":            int64(-1),
					"NT_HTTP_READ_TIMEOUT":    float64(30),
					"NT_HTTP_WRITE_TIMEOUT":   float64(30),
					"NT_TOKEN_PUBLIC_KEY_URL": "https://www.googleapis.com/oauth2/v1/certs",
					"NT_CORS_ENABLED":         false,
					"NT_SECURITY_ENABLED":     false,
					"NT_HTTP_HOME":            "/admin",
					"NT_TLS_CERT_FILE":        "",
					"NT_TLS_KEY_FILE":         "",
					"NT_CLIENT_CONFIG":        "../data/test_client_config.json",
				},
				services: map[string]srv.APIService{
					"http": &testServer{},
					"grpc": &testServer{},
				},
				infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
				errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
				tokenKeys: make(map[string]map[string]string),
				tray:      false,
			},
			wantErr: false,
		},
		{
			name: "tray icon",
			fields: fields{
				config: nt.IM{
					"version":                 "test",
					"NT_GRPC_ENABLED":         true,
					"NT_GRPC_TLS_ENABLED":     false,
					"NT_GRPC_PORT":            int64(-1),
					"NT_HTTP_ENABLED":         true,
					"NT_HTTP_TLS_ENABLED":     false,
					"NT_HTTP_PORT":            int64(-1),
					"NT_HTTP_READ_TIMEOUT":    float64(30),
					"NT_HTTP_WRITE_TIMEOUT":   float64(30),
					"NT_TOKEN_PUBLIC_KEY_URL": "https://www.googleapis.com/oauth2/v1/certs",
					"NT_CORS_ENABLED":         false,
					"NT_SECURITY_ENABLED":     false,
					"NT_HTTP_HOME":            "/admin",
					"NT_TLS_CERT_FILE":        "",
					"NT_TLS_KEY_FILE":         "",
					"NT_CLIENT_CONFIG":        "../data/test_client_config.json",
				},
				services: map[string]srv.APIService{
					"http": &testServer{},
					"grpc": &testServer{},
				},
				infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
				errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
				tokenKeys: make(map[string]map[string]string),
				tray:      true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:   tt.fields.services,
				defConn:    tt.fields.defConn,
				infoLog:    tt.fields.infoLog,
				errorLog:   tt.fields.errorLog,
				httpLog:    tt.fields.httpLog,
				args:       tt.fields.args,
				tokenKeys:  tt.fields.tokenKeys,
				config:     tt.fields.config,
				tray:       tt.fields.tray,
				taskSecKey: tt.fields.taskSecKey,
			}
			if err := app.startServer(); (err != nil) != tt.wantErr {
				t.Errorf("App.startServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_checkDefaultConn(t *testing.T) {
	type fields struct {
		services  map[string]srv.APIService
		defConn   nt.DataDriver
		infoLog   *log.Logger
		errorLog  *log.Logger
		httpLog   *log.Logger
		args      map[string]string
		tokenKeys map[string]map[string]string
		config    map[string]interface{}
		readFile  func(name string) ([]byte, error)
		getEnv    func(key string) string
	}
	os.Setenv("NT_ALIAS_TEST", testDatabase)
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test_db",
			fields: fields{
				config: nt.IM{
					"NT_ALIAS_DEFAULT":     "TEST",
					"NT_TOKEN_PRIVATE_KEY": "../data/x509/server_key.pem",
					"NT_TOKEN_PRIVATE_KID": "",
					"NT_TOKEN_PUBLIC_KID":  "public",
					"NT_TOKEN_PUBLIC_KEY":  "",
				},
				tokenKeys: map[string]map[string]string{},
				readFile: func(name string) ([]byte, error) {
					return []byte("ok"), nil
				},
				getEnv: func(key string) string {
					return ""
				},
			},
			wantErr: false,
		},
		{
			name: "token_error",
			fields: fields{
				config: nt.IM{
					"NT_ALIAS_DEFAULT":     "TEST",
					"NT_TOKEN_PRIVATE_KEY": "../data/x509/server_key.pem",
					"NT_TOKEN_PRIVATE_KID": "",
					"NT_TOKEN_PUBLIC_KID":  "public",
					"NT_TOKEN_PUBLIC_KEY":  "",
				},
				tokenKeys: map[string]map[string]string{},
				readFile: func(name string) ([]byte, error) {
					return nil, errors.New("error")
				},
				errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:  tt.fields.services,
				defConn:   tt.fields.defConn,
				infoLog:   tt.fields.infoLog,
				errorLog:  tt.fields.errorLog,
				httpLog:   tt.fields.httpLog,
				args:      tt.fields.args,
				tokenKeys: tt.fields.tokenKeys,
				config:    tt.fields.config,
				readFile:  tt.fields.readFile,
				getEnv:    tt.fields.getEnv,
			}
			if err := app.checkDefaultConn(); (err != nil) != tt.wantErr {
				t.Errorf("App.checkDefaultConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_GetNervaStore(t *testing.T) {
	type fields struct {
		services  map[string]srv.APIService
		defConn   nt.DataDriver
		infoLog   *log.Logger
		errorLog  *log.Logger
		httpLog   *log.Logger
		args      map[string]string
		tokenKeys map[string]map[string]string
		config    map[string]interface{}
	}
	type args struct {
		database string
	}
	defConn := &db.SQLDriver{Config: nt.IM{}}
	_ = defConn.CreateConnection("test", testDatabase)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "simple",
		},
		{
			name: "defconn",
			args: args{
				database: "test",
			},
			fields: fields{
				defConn: defConn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:  tt.fields.services,
				defConn:   tt.fields.defConn,
				infoLog:   tt.fields.infoLog,
				errorLog:  tt.fields.errorLog,
				httpLog:   tt.fields.httpLog,
				args:      tt.fields.args,
				tokenKeys: tt.fields.tokenKeys,
				config:    tt.fields.config,
			}
			app.GetNervaStore(tt.args.database)
		})
	}
}

func TestApp_GetResults(t *testing.T) {
	type fields struct {
		services  map[string]srv.APIService
		defConn   nt.DataDriver
		infoLog   *log.Logger
		errorLog  *log.Logger
		httpLog   *log.Logger
		args      map[string]string
		tokenKeys map[string]map[string]string
		config    map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "result",
			fields: fields{
				services: map[string]srv.APIService{
					"cli": &testServer{
						result: "value",
					},
				},
			},
			want: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:  tt.fields.services,
				defConn:   tt.fields.defConn,
				infoLog:   tt.fields.infoLog,
				errorLog:  tt.fields.errorLog,
				httpLog:   tt.fields.httpLog,
				args:      tt.fields.args,
				tokenKeys: tt.fields.tokenKeys,
				config:    tt.fields.config,
			}
			if got := app.GetResults(); got != tt.want {
				t.Errorf("App.GetResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_setEnv(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	type fields struct {
		services  map[string]srv.APIService
		defConn   nt.DataDriver
		infoLog   *log.Logger
		errorLog  *log.Logger
		httpLog   *log.Logger
		args      map[string]string
		tokenKeys map[string]map[string]string
		config    map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "env_file",
			fields: fields{
				infoLog:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
				errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
			},
		},
	}
	//osArgs := os.Args
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:  tt.fields.services,
				defConn:   tt.fields.defConn,
				infoLog:   tt.fields.infoLog,
				errorLog:  tt.fields.errorLog,
				httpLog:   tt.fields.httpLog,
				args:      tt.fields.args,
				tokenKeys: tt.fields.tokenKeys,
				config:    tt.fields.config,
			}
			if tt.name == "env_file" {
				os.Args = append(os.Args, "-env", ".env.example")
			}
			app.setEnv()
		})
	}
}

func TestApp_GetTokenKeys(t *testing.T) {
	type fields struct {
		services  map[string]srv.APIService
		defConn   nt.DataDriver
		infoLog   *log.Logger
		errorLog  *log.Logger
		httpLog   *log.Logger
		args      map[string]string
		tokenKeys map[string]map[string]string
		config    map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]nt.SM
	}{
		{
			name: "ok",
			fields: fields{
				tokenKeys: map[string]map[string]string{},
			},
			want: map[string]map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:  tt.fields.services,
				defConn:   tt.fields.defConn,
				infoLog:   tt.fields.infoLog,
				errorLog:  tt.fields.errorLog,
				httpLog:   tt.fields.httpLog,
				args:      tt.fields.args,
				tokenKeys: tt.fields.tokenKeys,
				config:    tt.fields.config,
			}
			if got := app.GetTokenKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.GetTokenKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestApp_trayIcon(t *testing.T) {
	type fields struct {
		services   map[string]srv.APIService
		defConn    nt.DataDriver
		infoLog    *log.Logger
		errorLog   *log.Logger
		httpLog    *log.Logger
		args       nt.SM
		tokenKeys  map[string]nt.SM
		config     map[string]interface{}
		readFile   func(name string) ([]byte, error)
		getEnv     func(key string) string
		tray       bool
		taskSecKey string
	}
	type args struct {
		httpDisabled bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "disabled_menu",
			args: args{
				httpDisabled: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:   tt.fields.services,
				defConn:    tt.fields.defConn,
				infoLog:    tt.fields.infoLog,
				errorLog:   tt.fields.errorLog,
				httpLog:    tt.fields.httpLog,
				args:       tt.fields.args,
				tokenKeys:  tt.fields.tokenKeys,
				config:     tt.fields.config,
				readFile:   tt.fields.readFile,
				getEnv:     tt.fields.getEnv,
				tray:       tt.fields.tray,
				taskSecKey: tt.fields.taskSecKey,
			}
			app.trayIcon(tt.args.httpDisabled)
		})
	}
}
*/

func TestApp_onTrayMenu(t *testing.T) {
	type fields struct {
		services   map[string]srv.APIService
		defConn    nt.DataDriver
		infoLog    *log.Logger
		errorLog   *log.Logger
		httpLog    *log.Logger
		args       nt.SM
		tokenKeys  map[string]nt.SM
		config     map[string]interface{}
		readFile   func(name string) ([]byte, error)
		getEnv     func(key string) string
		tray       bool
		taskSecKey string
	}
	type args struct {
		mKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "config",
			args: args{
				mKey: "config",
			},
			fields: fields{
				errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
			},
		},
		{
			name: "admin",
			args: args{
				mKey: "admin",
			},
			fields: fields{
				errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
			},
		},
		{
			name: "default",
			args: args{
				mKey: "default",
			},
			fields: fields{
				errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:   tt.fields.services,
				defConn:    tt.fields.defConn,
				infoLog:    tt.fields.infoLog,
				errorLog:   tt.fields.errorLog,
				httpLog:    tt.fields.httpLog,
				args:       tt.fields.args,
				tokenKeys:  tt.fields.tokenKeys,
				config:     tt.fields.config,
				readFile:   tt.fields.readFile,
				getEnv:     tt.fields.getEnv,
				tray:       tt.fields.tray,
				taskSecKey: tt.fields.taskSecKey,
			}
			app.onTrayMenu(tt.args.mKey)
		})
	}
}

func TestApp_openURL(t *testing.T) {
	type fields struct {
		services   map[string]srv.APIService
		defConn    nt.DataDriver
		infoLog    *log.Logger
		errorLog   *log.Logger
		httpLog    *log.Logger
		args       nt.SM
		tokenKeys  map[string]nt.SM
		config     map[string]interface{}
		readFile   func(name string) ([]byte, error)
		getEnv     func(key string) string
		tray       bool
		taskSecKey string
	}
	type args struct {
		goOS   string
		urlStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "darwin",
			args: args{
				goOS:   "darwin",
				urlStr: "",
			},
			wantErr: true,
		},
		{
			name: "windows",
			args: args{
				goOS:   "windows",
				urlStr: "",
			},
			wantErr: true,
		},
		{
			name: "linux",
			args: args{
				goOS:   "linux",
				urlStr: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:   tt.fields.services,
				defConn:    tt.fields.defConn,
				infoLog:    tt.fields.infoLog,
				errorLog:   tt.fields.errorLog,
				httpLog:    tt.fields.httpLog,
				args:       tt.fields.args,
				tokenKeys:  tt.fields.tokenKeys,
				config:     tt.fields.config,
				readFile:   tt.fields.readFile,
				getEnv:     tt.fields.getEnv,
				tray:       tt.fields.tray,
				taskSecKey: tt.fields.taskSecKey,
			}
			if err := app.openURL(tt.args.goOS, tt.args.urlStr); (err != nil) != tt.wantErr {
				t.Errorf("App.openURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_GetTaskSecKey(t *testing.T) {
	type fields struct {
		services   map[string]srv.APIService
		defConn    nt.DataDriver
		infoLog    *log.Logger
		errorLog   *log.Logger
		httpLog    *log.Logger
		args       nt.SM
		tokenKeys  map[string]nt.SM
		config     map[string]interface{}
		readFile   func(name string) ([]byte, error)
		getEnv     func(key string) string
		tray       bool
		taskSecKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				taskSecKey: "SEC01234",
			},
			want: "SEC01234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				services:   tt.fields.services,
				defConn:    tt.fields.defConn,
				infoLog:    tt.fields.infoLog,
				errorLog:   tt.fields.errorLog,
				httpLog:    tt.fields.httpLog,
				args:       tt.fields.args,
				tokenKeys:  tt.fields.tokenKeys,
				config:     tt.fields.config,
				readFile:   tt.fields.readFile,
				getEnv:     tt.fields.getEnv,
				tray:       tt.fields.tray,
				taskSecKey: tt.fields.taskSecKey,
			}
			if got := app.GetTaskSecKey(); got != tt.want {
				t.Errorf("App.GetTaskSecKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
