package api

import (
	"errors"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestCreateDatabase(t *testing.T) {
	type args struct {
		options cu.IM
		config  cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": testConn,
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
						},
					},
				},
			},
		},
		{
			name: "failed to check connection",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config:  cu.IM{},
			},
		},
		{
			name: "failed to begin transaction",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": testConn,
							"BeginTransaction": func() (interface{}, error) {
								return nil, errors.New("failed to begin transaction")
							},
						},
					},
				},
			},
		},
		{
			name: "UpdateSQL failed",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": testConn,
							"UpdateSQL": func(sqlString string, transaction interface{}) error {
								return errors.New("failed to update")
							},
						},
					},
				},
			},
		},
		{
			name: "ReadSQL failed",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{},
					},
				},
			},
		},
		{
			name: "Report install failed",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": testConn,
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, errors.New("failed to query")
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateDatabase(tt.args.options, tt.args.config)
		})
	}
}

func TestUpgradeDatabase(t *testing.T) {
	type args struct {
		options cu.IM
		config  cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": testConn,
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
						},
					},
				},
			},
		},
		{
			name: "failed to check connection",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config:  cu.IM{},
			},
		},
		{
			name: "failed to begin transaction",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": testConn,
							"BeginTransaction": func() (interface{}, error) {
								return nil, errors.New("failed to begin transaction")
							},
						},
					},
				},
			},
		},
		{
			name: "UpdateSQL failed",
			args: args{
				options: cu.IM{"alias": "test", "demo": true},
				config: cu.IM{
					"NT_ALIAS_TEST": "test",
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": testConn,
							"UpdateSQL": func(sqlString string, transaction interface{}) error {
								return errors.New("failed to update")
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpgradeDatabase(tt.args.options, tt.args.config)
		})
	}
}
