package app

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"log/slog"
	"os"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	ht "github.com/nervatura/nervatura/v6/pkg/host"
)

type testHost struct {
	raiseErr bool
	startErr bool
	stopErr  bool
	result   string
}

func (ts *testHost) StartServer(config cu.IM, appLogOut, httpLogOut io.Writer, interrupt chan os.Signal) error {
	if ts.startErr {
		return errors.New("error")
	}
	if ts.raiseErr {
		interrupt <- os.Interrupt
	}
	return nil
}

func (ts *testHost) Results() string {
	return ts.result
}

func (ts *testHost) StopServer(ctx context.Context) error {
	if ts.stopErr {
		return errors.New("error")
	}
	return nil
}

func generateRsaKeyPair() {
	//privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	privkey, _ := rsa.GenerateKey(rand.Reader, 2048)

	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	os.WriteFile("../../data/private.pem", privkey_pem, 0644)

	pubkey_bytes, _ := x509.MarshalPKIXPublicKey(&privkey.PublicKey)
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)
	os.WriteFile("../../data/public.pem", pubkey_pem, 0644)
}

func TestNew(t *testing.T) {
	type args struct {
		version string
		args    cu.SM
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "cli_app",
			args: args{
				version: "1.0.0",
				args: cu.SM{
					"cmd":              "help",
					"NT_APP_LOG_FILE":  "../../data/nervatura.log",
					"NT_HTTP_LOG_FILE": "../../data/http.log",
				},
			},
			wantErr: false,
		},
		{
			name: "cli_error",
			args: args{
				version: "dev",
				args: cu.SM{
					"NT_APP_LOG_FILE":  "data/nervatura.log",
					"NT_HTTP_LOG_FILE": "data/http.log",
				},
			},
			wantErr: true,
		},
		{
			name: "server_start",
			args: args{
				version: "test",
				args: cu.SM{
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
			_, err := New(tt.args.version, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestApp_setConfig(t *testing.T) {
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
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
				config: cu.IM{},
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
				config: cu.IM{},
				getEnv: func(key string) string {
					return ""
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			app.setConfig(tt.args.isSnap)
		})
	}
}

func Test_loadEnvFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				filename: "../../.env.example",
			},
			wantErr: false,
		},
		{
			name: "not_found",
			args: args{
				filename: "../../.env.missing",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := loadEnvFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadEnvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestApp_setEnv(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
	}
	type args struct {
		defaultEnvFile string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "env_file",
			fields: fields{
				config: cu.IM{},
				getEnv: func(key string) string {
					return ""
				},
			},
			args: args{
				defaultEnvFile: "../../.env.example",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			if tt.name == "env_file" {
				os.Args = append(os.Args, "-env", "../../.env.example", "-tray")
			}
			app.setEnv(tt.args.defaultEnvFile)
		})
	}
}

func TestApp_setTokenKeys(t *testing.T) {
	generateRsaKeyPair()
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
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
			name: "ok",
			fields: fields{
				config: cu.IM{
					"tokenKeys":            []cu.SM{},
					"NT_TOKEN_PRIVATE_KEY": "../../data/private.pem",
				},
				readFile: os.ReadFile,
			},
			args: args{
				keyType: "private",
			},
			wantErr: false,
		},
		{
			name: "file_error",
			fields: fields{
				config: cu.IM{
					"tokenKeys":            []cu.SM{},
					"NT_TOKEN_PRIVATE_KEY": "../../data/private.pem",
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				readFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
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
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			if err := app.setTokenKeys(tt.args.keyType); (err != nil) != tt.wantErr {
				t.Errorf("App.setTokenKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_setPublicTokenURLKeys(t *testing.T) {
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "url_error",
			fields: fields{
				config: cu.IM{
					"version":                 "test",
					"NT_TOKEN_PUBLIC_KEY_URL": "httpsx://www.googleapis.coma",
					"tokenKeys":               []cu.SM{},
				},
				appLog:  slog.New(slog.NewTextHandler(os.Stdout, nil)),
				readAll: io.ReadAll,
			},
		},
		{
			name: "body_error",
			fields: fields{
				config: cu.IM{
					"version":                 "test",
					"NT_TOKEN_PUBLIC_KEY_URL": "https://www.google.com",
					"tokenKeys":               []cu.SM{},
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				readAll: func(r io.Reader) ([]byte, error) {
					return nil, errors.New("body error")
				},
			},
		},
		{
			name: "ok",
			fields: fields{
				config: cu.IM{
					"version":                 "test",
					"NT_TOKEN_PUBLIC_KEY_URL": "https://www.google.com",
					"tokenKeys":               []cu.SM{},
				},
				appLog:  slog.New(slog.NewTextHandler(os.Stdout, nil)),
				readAll: io.ReadAll,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			app.setPublicTokenURLKeys()
		})
	}
}

func TestApp_GetResults(t *testing.T) {
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "result",
			fields: fields{
				hosts: map[string]ht.APIHost{
					"cli": &testHost{
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
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			if got := app.GetResults(); got != tt.want {
				t.Errorf("App.GetResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_openURL(t *testing.T) {
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
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
			fields: fields{
				hosts: map[string]ht.APIHost{
					"http": &testHost{
						startErr: true,
					},
					"grpc": &testHost{
						startErr: true,
					},
				},
			},
			args: args{
				goOS:   "linux",
				urlStr: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			if err := app.openURL(tt.args.goOS, tt.args.urlStr); (err != nil) != tt.wantErr {
				t.Errorf("App.openURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_onTrayMenu(t *testing.T) {
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
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
				config: cu.IM{
					"NT_TASK_SEC_KEY": "value",
					"NT_HTTP_PORT":    "8080",
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
			},
		},
		{
			name: "admin",
			args: args{
				mKey: "admin",
			},
			fields: fields{
				config: cu.IM{
					"NT_HTTP_PORT": "8080",
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
			},
		},
		{
			name: "default",
			args: args{
				mKey: "default",
			},
			fields: fields{
				config: cu.IM{
					"NT_HTTP_PORT": "8080",
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			app.onTrayMenu(tt.args.mKey)
		})
	}
}

func TestApp_startServer(t *testing.T) {
	type fields struct {
		config     cu.IM
		hosts      map[string]ht.APIHost
		showTray   bool
		taskSecKey string
		appLogOut  io.Writer
		httpLogOut io.Writer
		appLog     *slog.Logger
		getEnv     func(key string) string
		readFile   func(name string) ([]byte, error)
		readAll    func(r io.Reader) ([]byte, error)
	}
	type args struct {
		name      string
		interrupt chan os.Signal
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				name:      "test",
				interrupt: nil,
			},
			fields: fields{
				hosts: map[string]ht.APIHost{
					"test": &testHost{
						result: "value",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			if err := app.startServer(tt.args.name, tt.args.interrupt); (err != nil) != tt.wantErr {
				t.Errorf("App.startServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_backgroundServer(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "token_error",
			fields: fields{
				config: cu.IM{
					"tokenKeys":            []cu.SM{},
					"NT_TOKEN_PRIVATE_KEY": "../../data/private.pem",
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				readFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
			},
			wantErr: true,
		},
		{
			name: "interrupt",
			fields: fields{
				config: cu.IM{
					"tokenKeys":       []cu.SM{},
					"NT_HTTP_ENABLED": "true",
					"NT_GRPC_ENABLED": "true",
					"NT_HTTP_PORT":    5000,
					"NT_GRPC_PORT":    9200,
				},
				hosts: map[string]ht.APIHost{
					"http": &testHost{
						raiseErr: true,
					},
					"grpc": &testHost{
						startErr: false,
					},
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
			},
			wantErr: false,
		},
		{
			name: "cancel",
			fields: fields{
				config: cu.IM{
					"tokenKeys":       []cu.SM{},
					"NT_HTTP_ENABLED": "true",
					"NT_GRPC_ENABLED": "true",
					"NT_HTTP_PORT":    5000,
					"NT_GRPC_PORT":    9200,
				},
				hosts: map[string]ht.APIHost{
					"http": &testHost{
						startErr: true,
					},
					"grpc": &testHost{
						startErr: true,
					},
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
			},
			wantErr: false,
		},
		{
			name: "tray",
			fields: fields{
				config: cu.IM{
					"tokenKeys":       []cu.SM{},
					"NT_HTTP_ENABLED": "false",
					"NT_GRPC_ENABLED": "true",
					"NT_HTTP_PORT":    5000,
					"NT_GRPC_PORT":    9200,
				},
				hosts: map[string]ht.APIHost{
					"http": &testHost{
						startErr: true,
					},
					"grpc": &testHost{
						startErr: true,
					},
				},
				traySrv: &systemTray{
					app:          &App{},
					interrupt:    nil,
					ctx:          context.Background(),
					httpDisabled: true,
				},
				appLog:   slog.New(slog.NewTextHandler(os.Stdout, nil)),
				showTray: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				config:     tt.fields.config,
				hosts:      tt.fields.hosts,
				traySrv:    tt.fields.traySrv,
				showTray:   tt.fields.showTray,
				taskSecKey: tt.fields.taskSecKey,
				appLogOut:  tt.fields.appLogOut,
				httpLogOut: tt.fields.httpLogOut,
				appLog:     tt.fields.appLog,
				getEnv:     tt.fields.getEnv,
				readFile:   tt.fields.readFile,
				readAll:    tt.fields.readAll,
			}
			if err := app.backgroundServer(); (err != nil) != tt.wantErr {
				t.Errorf("App.backgroundServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
