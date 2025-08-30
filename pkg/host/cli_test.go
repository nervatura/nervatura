package server

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	cli "github.com/nervatura/nervatura/v6/pkg/service/cli"
)

func Test_cliHost_StartServer(t *testing.T) {
	type fields struct {
		config cu.IM
		appLog *slog.Logger
		srv    *cli.CLIService
		args   cu.SM
		result string
	}
	type args struct {
		config    cu.IM
		interrupt chan os.Signal
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "parse_flags",
			fields: fields{
				config: cu.IM{"args": cu.SM{}},
			},
			args: args{
				config:    cu.IM{"args": cu.SM{}},
				interrupt: make(chan os.Signal),
			},
			wantErr: false,
		},
		{
			name: "server",
			fields: fields{
				config: cu.IM{"args": cu.SM{"cmd": "server"}},
			},
			args: args{
				config:    cu.IM{"args": cu.SM{"cmd": "server"}},
				interrupt: make(chan os.Signal),
			},
			wantErr: false,
		},
		{
			name: "help",
			fields: fields{
				config: cu.IM{"args": cu.SM{"cmd": "help"}},
			},
			args: args{
				config:    cu.IM{"args": cu.SM{"cmd": "help"}},
				interrupt: make(chan os.Signal),
			},
			wantErr: false,
		},
		{
			name: "required_error",
			fields: fields{
				config: cu.IM{"args": cu.SM{"cmd": "update"}},
			},
			args: args{
				config:    cu.IM{"args": cu.SM{"cmd": "update"}},
				interrupt: make(chan os.Signal),
			},
			wantErr: true,
		},
		{
			name: "parse_error",
			fields: fields{
				config: cu.IM{"args": cu.SM{"cmd": "query", "model": "customer", "options": "{]"}},
			},
			args: args{
				config:    cu.IM{"args": cu.SM{"cmd": "query", "model": "customer", "options": "{]"}},
				interrupt: make(chan os.Signal),
			},
			wantErr: true,
		},
	}
	osArgs := os.Args
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cliHost{
				config: tt.fields.config,
				appLog: tt.fields.appLog,
				srv:    tt.fields.srv,
				args:   tt.fields.args,
				result: tt.fields.result,
			}
			appLogOut := &bytes.Buffer{}
			httpLogOut := &bytes.Buffer{}
			os.Args = osArgs
			if tt.name == "parse_flags" {
				os.Args = append(os.Args, "-help")
				os.Args = append(os.Args, "-c", "update")
				os.Args = append(os.Args, "-o", "{}")
				os.Args = append(os.Args, "-d", "{}")
				os.Args = append(os.Args, "-m", "customer")
			}
			if err := h.StartServer(tt.args.config, appLogOut, httpLogOut, tt.args.interrupt); (err != nil) != tt.wantErr {
				t.Errorf("cliHost.StartServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_cliHost_StopServer(t *testing.T) {
	type fields struct {
		config cu.IM
		appLog *slog.Logger
		srv    *cli.CLIService
		args   cu.SM
		result string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "stop",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cliHost{
				config: tt.fields.config,
				appLog: tt.fields.appLog,
				srv:    tt.fields.srv,
				args:   tt.fields.args,
				result: tt.fields.result,
			}
			if err := h.StopServer(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("cliHost.StopServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cliHost_Results(t *testing.T) {
	type fields struct {
		config cu.IM
		appLog *slog.Logger
		srv    *cli.CLIService
		args   cu.SM
		result string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "results",
			fields: fields{
				result: "results",
			},
			want: "results",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cliHost{
				config: tt.fields.config,
				appLog: tt.fields.appLog,
				srv:    tt.fields.srv,
				args:   tt.fields.args,
				result: tt.fields.result,
			}
			if got := h.Results(); got != tt.want {
				t.Errorf("cliHost.Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cliHost_checkRequired(t *testing.T) {
	type fields struct {
		config cu.IM
		appLog *slog.Logger
		srv    *cli.CLIService
		args   cu.SM
		result string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "create_err",
			fields: fields{
				args: cu.SM{"cmd": "create", "model": "customer", "options": "{}"},
			},
			wantErr: true,
		},
		{
			name: "query_err",
			fields: fields{
				args: cu.SM{"cmd": "query", "model": "customer"},
			},
			wantErr: true,
		},
		{
			name: "database_err",
			fields: fields{
				args: cu.SM{"cmd": "database"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cliHost{
				config: tt.fields.config,
				appLog: tt.fields.appLog,
				srv:    tt.fields.srv,
				args:   tt.fields.args,
				result: tt.fields.result,
			}
			if err := h.checkRequired(); (err != nil) != tt.wantErr {
				t.Errorf("cliHost.checkRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cliHost_parseCommand(t *testing.T) {
	type fields struct {
		config cu.IM
		appLog *slog.Logger
		srv    *cli.CLIService
		args   cu.SM
		result string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "found",
			fields: fields{
				config: cu.IM{},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				srv: &cli.CLIService{
					Config: cu.IM{},
				},
				args: cu.SM{"cmd": "view", "options": "{}"},
			},
			wantErr: false,
		},
		{
			name: "not_found",
			fields: fields{
				config: cu.IM{},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				srv: &cli.CLIService{
					Config: cu.IM{},
				},
				args: cu.SM{"cmd": "get", "model": "baba", "options": "{}"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cliHost{
				config: tt.fields.config,
				appLog: tt.fields.appLog,
				srv:    tt.fields.srv,
				args:   tt.fields.args,
				result: tt.fields.result,
			}
			_, err := h.parseCommand()
			if (err != nil) != tt.wantErr {
				t.Errorf("cliHost.parseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
