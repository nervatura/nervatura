package api

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func TestDataStore_TokenRefresh(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		user md.Auth
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestTokenRefresh",
			fields: fields{
				Db:     &md.TestDriver{},
				Alias:  "test",
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
			args: args{
				user: md.Auth{
					Code:     "test",
					UserName: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "TestTokenRefreshError",
			fields: fields{
				Db:     &md.TestDriver{},
				Alias:  "test",
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "", errors.New("test")
				},
			},
			args: args{
				user: md.Auth{
					Code:     "test",
					UserName: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.TokenRefresh(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.TokenRefresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_AuthUser(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		uid      string
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "name": "test"}}, nil
						},
					},
				},
				Alias:  "test",
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
			},
			args: args{
				uid:      "test",
				username: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.AuthUser(tt.args.uid, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.AuthUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_UserLogin(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		username    string
		password    string
		createToken bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantToken string
		wantErr   bool
	}{
		{
			name: "success",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
						},
					},
				},
				Config: cu.IM{
					"NT_API_KEY": "test",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
			args: args{
				username:    "test",
				password:    "test",
				createToken: true,
			},
			wantToken: "test",
			wantErr:   false,
		},
		{
			name: "default admin password",
			fields: fields{
				Config: cu.IM{
					"NT_API_KEY": "test",
				},
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "usref" {
								return []cu.IM{}, nil
							}
							return []cu.IM{{"id": 1, "code": st.DefaultConfig["connection"]["default_admin"], "user_group": md.UserGroupAdmin}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToType: func(data interface{}, result any) (err error) {
					return ut.ConvertToType(data, result)
				},
			},
			args: args{
				username:    "admin",
				password:    "test",
				createToken: false,
			},
			wantToken: "",
			wantErr:   false,
		},
		{
			name: "wrong default admin password",
			fields: fields{
				Config: cu.IM{
					"NT_API_KEY": "test",
				},
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "usref" {
								return []cu.IM{}, nil
							}
							return []cu.IM{{"id": 1, "code": st.DefaultConfig["connection"]["default_admin"], "user_group": md.UserGroupAdmin}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToType: func(data interface{}, result any) (err error) {
					return ut.ConvertToType(data, result)
				},
			},
			args: args{
				username:    "admin",
				password:    "test1",
				createToken: false,
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "wrong password",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return errors.New("error")
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
			args: args{
				username:    "test",
				password:    "test1",
				createToken: true,
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "missing user",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreatePasswordHash: func(password string) (hash string, err error) {
					return "test", nil
				},
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
			args: args{
				username:    "test",
				password:    "",
				createToken: true,
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "missing password",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "usref" {
								return []cu.IM{}, nil
							}
							return []cu.IM{{"id": 1, "code": st.DefaultConfig["connection"]["default_admin"], "user_group": md.UserGroupAdmin}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreatePasswordHash: func(password string) (hash string, err error) {
					return "test", nil
				},
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
			args: args{
				username:    "test",
				password:    "",
				createToken: true,
			},
			wantToken: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			gotToken, err := ds.UserLogin(tt.args.username, tt.args.password, tt.args.createToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("DataStore.UserLogin() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestDataStore_TokenLogin(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
						},
					},
				},
				Config: cu.IM{
					"tokenKeys": []cu.SM{{"user_code": "user_code"}, {"user_name": "user_name"}},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ParseToken: func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
					return cu.IM{"user_code": "test", "user_name": "test"}, nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
			},
			args: args{
				token: "test",
			},
			wantErr: false,
		},
		{
			name: "missing token",
			fields: fields{
				Db: &md.TestDriver{},
			},
			args: args{
				token: "",
			},
			wantErr: true,
		},
		{
			name: "invalid token",
			fields: fields{
				Db: &md.TestDriver{},
				Config: cu.IM{
					"tokenKeys": []cu.SM{{"user_code": "user_code"}, {"user_name": "user_name"}},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ParseToken: func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
					return cu.IM{}, errors.New("error")
				},
			},
			args: args{
				token: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.TokenLogin(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.TokenLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_UserPassword(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		userCode string
		password string
		confirm  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success update",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"value": "test"}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 100, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreatePasswordHash: func(password string) (hash string, err error) {
					return "test", nil
				},
			},
			args: args{
				userCode: "test",
				password: "test",
				confirm:  "test",
			},
			wantErr: false,
		},
		{
			name: "success insert",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 100, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreatePasswordHash: func(password string) (hash string, err error) {
					return "test", nil
				},
			},
			args: args{
				userCode: "test",
				password: "test",
				confirm:  "test",
			},
			wantErr: false,
		},
		{
			name: "query error",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, errors.New("error")
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreatePasswordHash: func(password string) (hash string, err error) {
					return "test", nil
				},
			},
			args: args{
				userCode: "test",
				password: "test",
				confirm:  "test",
			},
			wantErr: true,
		},
		{
			name: "create password hash error",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreatePasswordHash: func(password string) (hash string, err error) {
					return "", errors.New("error")
				},
			},
			args: args{
				userCode: "test",
				password: "test",
				confirm:  "test",
			},
			wantErr: true,
		},
		{
			name: "verify password error",
			fields: fields{
				Db: &md.TestDriver{},
			},
			args: args{
				userCode: "test",
				password: "test",
				confirm:  "test2",
			},
			wantErr: true,
		},
		{
			name: "empty password error",
			fields: fields{
				Db: &md.TestDriver{},
			},
			args: args{
				userCode: "test",
				password: "",
				confirm:  "test2",
			},
			wantErr: true,
		},
		{
			name: "missing user code error",
			fields: fields{
				Db: &md.TestDriver{},
			},
			args: args{
				userCode: "",
				password: "test",
				confirm:  "test2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			if err := ds.UserPassword(tt.args.userCode, tt.args.password, tt.args.confirm); (err != nil) != tt.wantErr {
				t.Errorf("DataStore.UserPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
