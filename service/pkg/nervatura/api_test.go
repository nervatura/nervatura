package nervatura

import (
	"errors"
	"testing"
	"time"

	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

func TestAPI_getHashvalue(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		refname string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "hash_value",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_HASHTABLE": testData.hashTable,
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{{"value": testData.validHash}}, nil
					},
				}, nil, nil),
			},
			args: args{
				refname: "testRefname",
			},
			wantErr: false,
		},
		{
			name: "hash_empty",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_HASHTABLE": testData.hashTable,
				}, nil, nil),
			},
			args: args{
				refname: "testRefname",
			},
			wantErr: false,
		},
		{
			name: "check_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"CheckHashtable": func() error {
						return errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				refname: "testRefname",
			},
			wantErr: true,
		},
		{
			name: "query_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_HASHTABLE": testData.hashTable,
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				refname: "testRefname",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.getHashvalue(tt.args.refname)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.getHashvalue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_authUser(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "auth_user",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: false,
		},
		{
			name: "auth_customer",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_ALIAS_TEST": "test",
					"Connection": func() struct {
						Alias     string
						Connected bool
						Engine    string
					} {
						return struct {
							Alias     string
							Connected bool
							Engine    string
						}{
							Alias:     "test",
							Connected: false,
							Engine:    "test",
						}
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custnumber": "DMCUST/00001"},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_opt_database",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Connection": func() struct {
						Alias     string
						Connected bool
						Engine    string
					} {
						return struct {
							Alias     string
							Connected bool
							Engine    string
						}{
							Alias:     "test",
							Connected: false,
							Engine:    "test",
						}
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
				},
			},
			wantErr: true,
		},
		{
			name: "missing_def_database",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Connection": func() struct {
						Alias     string
						Connected bool
						Engine    string
					} {
						return struct {
							Alias     string
							Connected bool
							Engine    string
						}{
							Alias:     "test",
							Connected: false,
							Engine:    "test",
						}
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "connection_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_ALIAS_TEST": "test",
					"Connection": func() struct {
						Alias     string
						Connected bool
						Engine    string
					} {
						return struct {
							Alias     string
							Connected bool
							Engine    string
						}{
							Alias:     "test",
							Connected: false,
							Engine:    "test",
						}
					},
					"CreateConnection": func() error {
						return errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_user",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "auth_user_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "auth_customer_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "unknown_user",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			if err := api.authUser(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("API.authUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_checkVersion(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "version_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			wantErr: false,
		},
		{
			name: "version_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			if err := api.checkVersion(); (err != nil) != tt.wantErr {
				t.Errorf("API.checkVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_TokenRefresh(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "refresh_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_HASHTABLE":         testData.hashTable,
					"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
				}, &User{Username: "admin"}, nil),
			},
			wantErr: false,
		},
		{
			name: "not_connect",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Connection": func() struct {
						Alias     string
						Connected bool
						Engine    string
					} {
						return struct {
							Alias     string
							Connected bool
							Engine    string
						}{
							Alias:     "test",
							Connected: false,
							Engine:    "test",
						}
					},
				}, nil, nil),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.TokenRefresh()
			if (err != nil) != tt.wantErr {
				t.Errorf("API.TokenRefresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_TokenLogin(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	adminToken, _ := ut.CreateToken("admin", "test", IM{"NT_TOKEN_PRIVATE_KEY": testData.tokenKey})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "login_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"token": adminToken,
					"keys":  map[string]map[string]string{},
				},
			},
			wantErr: false,
		},
		{
			name: "missing_token",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"keys": map[string]map[string]string{},
				},
			},
			wantErr: true,
		},
		{
			name: "token_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"token": adminToken,
					"keys":  map[string]map[string]string{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			if err := api.TokenLogin(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("API.TokenLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_UserPassword(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "username_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": 1},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "123",
					"confirm":  "123",
				},
			},
			wantErr: false,
		},
		{
			name: "username_ref_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, &User{Username: "admin", Scope: "admin"}, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "123",
					"confirm":  "123",
				},
			},
			wantErr: false,
		},
		{
			name: "custnumber_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": 1},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"custnumber": "DMCUST/00001",
					"password":   "123",
					"confirm":    "123",
				},
			},
			wantErr: false,
		},
		{
			name: "custnumber_ref_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, IM{"id": int64(1), "custnumber": "DMCUST/00001"}),
			},
			args: args{
				options: IM{
					"custnumber": "DMCUST/00001",
					"password":   "123",
					"confirm":    "123",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_username",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"password": "123",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_password",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_confirm",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "123",
				},
			},
			wantErr: true,
		},
		{
			name: "empty_password",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
		{
			name: "verify_password",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "321",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
		{
			name: "not_connect",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Connection": func() struct {
						Alias     string
						Connected bool
						Engine    string
					} {
						return struct {
							Alias     string
							Connected bool
							Engine    string
						}{
							Alias:     "test",
							Connected: false,
							Engine:    "test",
						}
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "123",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
		{
			name: "username_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "123",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
		{
			name: "unknown_user",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "123",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			if err := api.UserPassword(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("API.UserPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_UserLogin(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "login_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"value": testData.validHash},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: false,
		},
		{
			name: "default_database",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custnumber": "DMCUST/00001"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"value": testData.validHash},
						}, nil
					},
					"NT_ALIAS_DEFAULT": "test",
				}, nil, IM{"id": int64(1), "custnumber": "DMCUST/00001"}),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
				},
			},
			wantErr: false,
		},
		{
			name: "password_login_disabled",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custnumber": "DMCUST/00001"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"value": testData.validHash},
						}, nil
					},
					"NT_PASSWORD_LOGIN": false,
				}, nil, IM{"id": int64(1), "custnumber": "DMCUST/00001"}),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
				},
			},
			wantErr: true,
		},
		{
			name: "missing_database",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
				},
			},
			wantErr: true,
		},
		{
			name: "missing_user",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "version_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "hash_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_HASHTABLE": testData.hashTable,
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == testData.hashTable {
							return nil, errors.New("error")
						}
						return []IM{}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_password",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"value": testData.validHash},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "argon_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"value": "HS0123456789HS"},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": testData.validPass,
					"database": "test",
				},
			},
			wantErr: true,
		},
		{
			name: "wrong_password",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"username": "admin", "scope": "admin"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"value": testData.validHash},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"username": "admin",
					"password": "password",
					"database": "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, _, err := api.UserLogin(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_DatabaseCreate(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_ALIAS_TEST": "test",
				}, nil, nil),
			},
			args: args{
				options: IM{
					"database": "test",
					"demo":     false,
				},
			},
			wantErr: false,
		},
		{
			name: "create_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_ALIAS_TEST": "test",
					"CreateDatabase": func() ([]SM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"database": "test",
					"demo":     false,
				},
			},
			wantErr: true,
		},
		{
			name: "demo_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_ALIAS_TEST": "test",
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "groups" {
							return []IM{
								{"id": 1, "groupname": "nervatype", "groupvalue": "value"},
								{"id": 2, "groupname": "transtype", "groupvalue": "value"},
								{"id": 3, "groupname": "direction", "groupvalue": "value"},
								{"id": 4, "groupname": "filetype", "groupvalue": "value"},
								{"id": 5, "groupname": "fieldtype", "groupvalue": "value"},
								{"id": 6, "groupname": "wheretype", "groupvalue": "value"},
							}, nil
						}
						return []IM{}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"database": "test",
					"demo":     true,
				},
			},
			wantErr: true,
		},
		{
			name: "missing_database",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_ALIAS_TEST": "test",
				}, nil, nil),
			},
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
		{
			name: "conn_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_ALIAS_TEST": "test",
					"CreateConnection": func() error {
						return errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"database": "test",
					"demo":     false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.DatabaseCreate(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.DatabaseCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_demoDatabase(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*
			{
				name: "list_error",
				fields: fields{
					NStore: testData.getNervaStore(IM{
						"Query": func(queries []Query) ([]IM, error) {
							if queries[0].From == "ui_report" {
								return nil, errors.New("error")
							}
							return []IM{}, nil
						},
					}, nil, nil),
				},
				args: args{
					options: IM{
						"logData": []SM{},
					},
				},
				wantErr: true,
			},
			{
				name: "install_error",
				fields: fields{
					NStore: testData.getNervaStore(IM{
						"Query": func(queries []Query) ([]IM, error) {
							if queries[0].From == "groups" {
								return nil, errors.New("error")
							}
							return []IM{}, nil
						},
					}, nil, nil),
				},
				args: args{
					options: IM{
						"logData": []SM{},
					},
				},
				wantErr: true,
			},
		*/
		{
			name: "empty_reports",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "groups" {
							return nil, errors.New("error")
						}
						return []IM{}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"logData":    []SM{},
					"report_dir": "../../data/fonts",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.demoDatabase(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.demoDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_Delete(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "address",
					"id":        123,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			if err := api.Delete(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("API.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_Get(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "filters_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": 1},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": 1, "ref_id": 1, "fieldname": "fieldname",
								"fieldtype": "fieldtype", "value": "value", "notes": "notes"},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "custname;==;First Customer Co.|custnumber;in;DMCUST/00001,DMCUST/00002",
				},
			},
			wantErr: false,
		},
		{
			name: "ids_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
					"ids":       "2,4",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"metadata": true,
					"ids":      "2,4",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_nervatype",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "kalevala",
					"metadata":  true,
					"ids":       "2,4",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_filter_ids",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_filter",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "custname;==|custnumber;in;DMCUST/00001,DMCUST/00002",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_fieldname",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "kalevala;==;First Customer Co.|custnumber;in;DMCUST/00001,DMCUST/00002",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_comparison",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "custname;#=#;First Customer Co.|custnumber;in;DMCUST/00001,DMCUST/00002",
				},
			},
			wantErr: true,
		},
		{
			name: "Query_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "custname;==;First Customer Co.|custnumber;in;DMCUST/00001,DMCUST/00002",
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": 1},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "custname;==;First Customer Co.|custnumber;in;DMCUST/00001,DMCUST/00002",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.Get(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_View(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options []IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "view_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: []IM{
					{
						"key":    "customers",
						"text":   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
						"values": []interface{}{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "BeginTransaction_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"BeginTransaction": func() (interface{}, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: []IM{},
			},
			wantErr: true,
		},
		{
			name: "missing_key",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: []IM{
					{
						"text":   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
						"values": []interface{}{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing_values",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: []IM{
					{
						"key":  "customers",
						"text": "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing_text",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: []IM{
					{
						"key":    "customers",
						"values": []interface{}{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "QuerySQL_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: []IM{
					{
						"key":    "customers",
						"text":   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
						"values": []interface{}{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "RollbackTransaction_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"RollbackTransaction": func() error {
						return errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: []IM{
					{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.View(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.View() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_Function(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "function_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"key": "nextNumber",
					"values": IM{
						"numberkey": "custnumber",
						"step":      false,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing_key",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"values": IM{
						"numberkey": "custnumber",
						"step":      false,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing_values",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"key": "nextNumber",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.Function(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.Function() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_updateTransInfo(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		data []IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "info_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["trans_id"] == nil {
							return []IM{
								{"rtype": "trans", "transtype": 123, "custnumber": "DMCUST/00001"},
							}, nil
						}
						return []IM{
							{"rtype": "customer", "transtype": nil, "custnumber": "DMCUST/00001"},
							{"rtype": "groups", "transtype": "invoice", "custnumber": nil},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				data: []IM{
					{
						"id":        1,
						"transtype": 123,
						"keys": IM{
							"direction":   "out",
							"customer_id": "DMCUST/00001",
						},
					},
					{
						"customer_id": 234,
						"keys": IM{
							"transtype": "invoice",
							"direction": "out",
						},
					},
					{
						"customer_id": 234,
					},
					{
						"id":        1,
						"transtype": 123,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.updateTransInfo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.updateTransInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_updateSetKeys(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		nervatype string
		data      []IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "trans_info",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["extraInfo"] == true {
							return []IM{
								{"id": 1, "custtype": "own", "terms": 0, "custname": "custname",
									"taxnumber": "taxnumber", "zipcode": "zipcode", "city": "city", "street": "street"},
								{"id": 2, "custtype": "company", "terms": 0, "custname": "custname",
									"taxnumber": "taxnumber", "zipcode": "zipcode", "city": "city", "street": "street"},
							}, nil
						}
						return []IM{
							{"transtype": 1, "direction": 1, "digit": 2},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "numberkey": "numberkey", "prefix": "prefix", "curvalue": int64(1),
								"isyear": false, "sep": "/", "len": int64(5), "description": ""},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{
						"keys": IM{
							"id":          "DMINV/00003",
							"ref_id":      "customer/DMCUST/00001",
							"transtype":   "invoice",
							"customer_id": "DMCUST/00001",
							"transnumber": IL{"numberdef", "invoice_out"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "link_info",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": 1, "transtype": 1, "direction": 1, "digit": 0, "qty": 1,
								"discount": 0, "tax_id": 1, "rate": 1, "movetype": 1},
							{"id": 1, "transtype": 1, "direction": 1, "digit": 0, "qty": 1,
								"discount": 0, "tax_id": 1, "rate": 1, "movetype": 1},
							{"id": 1, "transtype": 1, "direction": 1, "digit": 0, "qty": 1,
								"discount": 0, "tax_id": 1, "rate": 1, "movetype": 1},
							{"id": 1, "transtype": 1, "direction": 1, "digit": 0, "qty": 1,
								"discount": 0, "tax_id": 1, "rate": 1, "movetype": 1},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "link",
				data: []IM{
					{
						"keys": IM{
							"nervatype_1": "movement",
							"ref_id_1":    "movement/DMDEL/00001~3",
							"nervatype_2": "item",
							"ref_id_2":    "item/DMORD/00001~2",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "event_numberdef",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "numberkey": "numberkey", "prefix": "prefix", "curvalue": int64(1),
								"isyear": false, "sep": "/", "len": int64(5), "description": ""},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "event",
				data: []IM{
					{
						"keys": IM{
							"calnumber": "numberdef",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "customer_refnumber",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": 1, "custtype": 1},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "customer",
				data: []IM{
					{
						"keys": IM{
							"custnumber": "DMCUST/00001",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "refnumber_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "event",
				data: []IM{
					{
						"keys": IM{
							"calnumber": "numberdef",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_nervatype",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				nervatype: "customer",
				data: []IM{
					{
						"keys": IM{
							"kalevala": "DMINV/00003",
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.updateSetKeys(tt.args.nervatype, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.updateSetKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_updateCheckInfo(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		nervatype string
		data      []IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "trans_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{"transdate": time.Now().Format("2006-01-02"), "transtype": 1, "direction": 2, "transtate": 3},
				},
			},
			wantErr: false,
		},
		{
			name: "trans_cruser_id",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, &User{Username: "admin"}, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{"transdate": time.Now().Format("2006-01-02"), "transtype": 1, "direction": 2, "transtate": 3},
				},
			},
			wantErr: false,
		},
		{
			name: "trans_crdate",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{"id": 1, "crdate": time.Now().Format("2006-01-02")},
				},
			},
			wantErr: false,
		},
		{
			name: "log_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				nervatype: "log",
				data: []IM{
					{"employee_id": 1, "logstate": 3},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.updateCheckInfo(tt.args.nervatype, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.updateCheckInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_Update(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		nervatype string
		data      []IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "trans" {
							return []IM{
								{"id": 1},
							}, nil
						}
						return []IM{}, nil
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{
						"id":    1,
						"notes": "notes",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid_nervatype",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				nervatype: "kalevala",
				data: []IM{
					{
						"id":    1,
						"notes": "notes",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "updateTransInfo_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{
						"id":    1,
						"notes": "notes",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "updateSetKeys_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{
						"notes": "notes",
						"keys": IM{
							"id": "DMINV/00003"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "updateCheckInfo_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				nervatype: "customer",
				data: []IM{
					{},
				},
			},
			wantErr: true,
		},
		{
			name: "BeginTransaction_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"BeginTransaction": func() (interface{}, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "customer",
				data: []IM{
					{
						"id": 1,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "update_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"RollbackTransaction": func() error {
						return errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				nervatype: "trans",
				data: []IM{
					{
						"id":    1,
						"notes": "notes",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.Update(tt.args.nervatype, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_Report(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "report_err",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.Report(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.Report() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_ReportList(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "static_list",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": 1, "reportkey": "ntr_invoice_en"},
						}, nil
					},
				}, nil, nil),
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "repdir_list",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"NT_REPORT_DIR": "../../data/test",
				}, nil, nil),
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.ReportList(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.ReportList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_ReportDelete(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete_ok",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": 1, "reportkey": "ntr_invoice_en"},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey": "ntr_invoice_en",
				},
			},
			wantErr: false,
		},
		{
			name: "delete_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey": "ntr_invoice_en",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_args",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
		{
			name: "missing_key",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey": "ntr_invoice_en",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			if err := api.ReportDelete(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("API.ReportDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_ReportInstall(t *testing.T) {
	type fields struct {
		NStore *NervaStore
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "static_install",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "groups" {
							return []IM{
								{"id": 1, "groupname": "nervatype", "groupvalue": "value"},
								{"id": 2, "groupname": "transtype", "groupvalue": "value"},
								{"id": 3, "groupname": "direction", "groupvalue": "value"},
								{"id": 4, "groupname": "filetype", "groupvalue": "value"},
								{"id": 5, "groupname": "fieldtype", "groupvalue": "value"},
								{"id": 6, "groupname": "wheretype", "groupvalue": "value"},
							}, nil
						}
						return []IM{}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey": "ntr_invoice_en",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_reportkey",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
		{
			name: "missing_file",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey":  "ntr_invoice_en",
					"report_dir": "../data",
				},
			},
			wantErr: true,
		},
		{
			name: "Query_error",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey": "ntr_invoice_en",
				},
			},
			wantErr: true,
		},
		{
			name: "exists_template",
			fields: fields{
				NStore: testData.getNervaStore(IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": 1},
						}, nil
					},
				}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey": "ntr_invoice_en",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_template",
			fields: fields{
				NStore: testData.getNervaStore(IM{}, nil, nil),
			},
			args: args{
				options: IM{
					"reportkey":  "test_client_config",
					"report_dir": "../../data",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				NStore: tt.fields.NStore,
			}
			_, err := api.ReportInstall(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.ReportInstall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
