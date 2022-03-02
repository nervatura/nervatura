package app

import (
	"os"
	"testing"

	nt "github.com/nervatura/nervatura-service/pkg/nervatura"
	srv "github.com/nervatura/nervatura-service/pkg/service"
)

func Test_cliServer_StartService(t *testing.T) {
	type fields struct {
		app     *App
		service srv.CLIService
		args    nt.SM
		result  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "start_server",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
					args: nt.SM{
						"cmd": "server",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "start_ok",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
					args: nt.SM{
						"cmd": "help",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "parse_flags",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
			},
			wantErr: false,
		},
		{
			name: "required_error",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
					args: nt.SM{
						"cmd": "UserLogin",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "parse_error",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
					args: nt.SM{
						"cmd":     "UserLogin",
						"options": "options",
					},
				},
			},
			wantErr: true,
		},
	}
	osArgs := os.Args
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cliServer{
				app:     tt.fields.app,
				service: tt.fields.service,
				args:    tt.fields.args,
				result:  tt.fields.result,
			}
			os.Args = osArgs
			if tt.name == "parse_flags" {
				os.Args = append(os.Args, "-help")
				os.Args = append(os.Args, "-t", "token")
				os.Args = append(os.Args, "-o", "options")
				os.Args = append(os.Args, "-d", "data")
				os.Args = append(os.Args, "-nt", "nervatype")
				os.Args = append(os.Args, "-k", "key")
			}
			if err := s.StartService(); (err != nil) != tt.wantErr {
				t.Errorf("cliServer.StartService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cliServer_Results(t *testing.T) {
	type fields struct {
		app     *App
		service srv.CLIService
		args    nt.SM
		result  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "result",
			fields: fields{
				result: "value",
			},
			want: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cliServer{
				app:     tt.fields.app,
				service: tt.fields.service,
				args:    tt.fields.args,
				result:  tt.fields.result,
			}
			if got := s.Results(); got != tt.want {
				t.Errorf("cliServer.Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cliServer_ConnectApp(t *testing.T) {
	type fields struct {
		app     *App
		service srv.CLIService
		args    nt.SM
		result  string
	}
	type args struct {
		app interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "app",
			args: args{
				app: &App{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cliServer{
				app:     tt.fields.app,
				service: tt.fields.service,
				args:    tt.fields.args,
				result:  tt.fields.result,
			}
			s.ConnectApp(tt.args.app)
		})
	}
}

func Test_cliServer_StopService(t *testing.T) {
	type fields struct {
		app     *App
		service srv.CLIService
		args    nt.SM
		result  string
	}
	type args struct {
		in0 interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "stop_server",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cliServer{
				app:     tt.fields.app,
				service: tt.fields.service,
				args:    tt.fields.args,
				result:  tt.fields.result,
			}
			if err := s.StopService(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("cliServer.StopService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cliServer_checkRequired(t *testing.T) {
	type fields struct {
		app     *App
		service srv.CLIService
		args    nt.SM
		result  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "missing_token",
			fields: fields{
				args: nt.SM{
					"cmd": "TokenLogin",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_options",
			fields: fields{
				args: nt.SM{
					"cmd": "UserLogin",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_dbs_options",
			fields: fields{
				args: nt.SM{
					"cmd": "DatabaseCreate",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_dbs_api_key",
			fields: fields{
				args: nt.SM{
					"cmd":     "DatabaseCreate",
					"options": "options",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_view_data",
			fields: fields{
				args: nt.SM{
					"cmd":   "View",
					"token": "token",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				args: nt.SM{
					"cmd":   "Update",
					"token": "token",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_update_data",
			fields: fields{
				args: nt.SM{
					"cmd":       "Update",
					"token":     "token",
					"nervatype": "nervatype",
				},
			},
			wantErr: true,
		},
		{
			name: "check_ok",
			fields: fields{
				args: nt.SM{
					"cmd":   "View",
					"token": "token",
					"data":  "data",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cliServer{
				app:     tt.fields.app,
				service: tt.fields.service,
				args:    tt.fields.args,
				result:  tt.fields.result,
			}
			if err := s.checkRequired(); (err != nil) != tt.wantErr {
				t.Errorf("cliServer.checkRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cliServer_parseCommand(t *testing.T) {
	type fields struct {
		app     *App
		service srv.CLIService
		args    nt.SM
		result  string
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult string
		wantErr    bool
	}{
		{
			name: "TokenDecode",
			fields: fields{
				args: nt.SM{
					"cmd":   "TokenDecode",
					"token": "token",
				},
			},
			wantErr: false,
		},
		{
			name: "TokenLogin",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd":   "TokenLogin",
					"token": "token",
				},
			},
			wantErr: false,
		},
		{
			name: "options_error",
			fields: fields{
				args: nt.SM{
					"options": "options",
				},
			},
			wantErr: true,
		},
		{
			name: "data_error",
			fields: fields{
				args: nt.SM{
					"data": "data",
				},
			},
			wantErr: true,
		},
		{
			name: "UserLogin",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "UserLogin",
				},
			},
			wantErr: false,
		},
		{
			name: "UserPassword",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "UserPassword",
				},
			},
			wantErr: false,
		},
		{
			name: "TokenRefresh",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "TokenRefresh",
				},
			},
			wantErr: false,
		},
		{
			name: "Get",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "Get",
				},
			},
			wantErr: false,
		},
		{
			name: "View",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "View",
				},
			},
			wantErr: false,
		},
		{
			name: "Function",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "Function",
				},
			},
			wantErr: false,
		},
		{
			name: "Update",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "Update",
				},
			},
			wantErr: false,
		},
		{
			name: "Delete",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "Delete",
				},
			},
			wantErr: false,
		},
		{
			name: "DatabaseCreate",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "DatabaseCreate",
				},
			},
			wantErr: false,
		},
		{
			name: "Report",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "Report",
				},
			},
			wantErr: false,
		},
		{
			name: "ReportList",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "ReportList",
				},
			},
			wantErr: false,
		},
		{
			name: "ReportInstall",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "ReportInstall",
				},
			},
			wantErr: false,
		},
		{
			name: "ReportDelete",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "ReportDelete",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid_command",
			fields: fields{
				app: &App{
					config: make(map[string]interface{}),
				},
				args: nt.SM{
					"cmd": "Kalevala",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cliServer{
				app:     tt.fields.app,
				service: tt.fields.service,
				args:    tt.fields.args,
				result:  tt.fields.result,
			}
			_, err := s.parseCommand()
			if (err != nil) != tt.wantErr {
				t.Errorf("cliServer.parseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
