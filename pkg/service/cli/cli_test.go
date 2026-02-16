package cli

import (
	"bytes"
	"log/slog"
	"reflect"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestNewCLIService(t *testing.T) {
	type args struct {
		config cu.IM
		appLog *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want *CLIService
	}{
		{
			name: "success",
			args: args{
				config: cu.IM{},
				appLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			want: &CLIService{Config: cu.IM{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCLIService(tt.args.config, tt.args.appLog); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCLIService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ResetPassword(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &td.TestDriver{
							Config: cu.IM{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			cli.ResetPassword(tt.args.options, tt.args.requestData)
		})
	}
}

func TestCLIService_Database(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			cli.Database(tt.args.options, tt.args.requestData)
		})
	}
}

func TestCLIService_Function(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &td.TestDriver{
							Config: cu.IM{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			cli.Function(tt.args.options, tt.args.requestData)
		})
	}
}

func TestCLIService_View(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &td.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
					},
					"name": "VIEW_CUSTOMER_EVENTS",
					"filters": []any{
						map[string]any{"field": "like_subject", "value": "visit"},
						map[string]any{"field": "place", "value": "City1"},
					},
					"limit": 10,
				},
			},
			want: `{"code":200,"data":[{"id":1}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.View(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.View() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_Upgrade(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			cli.Upgrade(tt.args.options, tt.args.requestData)
		})
	}
}
