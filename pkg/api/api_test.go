package api

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestNewDataStore(t *testing.T) {
	type args struct {
		config cu.IM
		alias  string
		appLog *slog.Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_new_datastore",
			args: args{
				config: cu.IM{
					"db": &md.TestDriver{},
				},
				alias:  "test",
				appLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewDataStore(tt.args.config, tt.args.alias, tt.args.appLog)
		})
	}
}

func TestDataStore_SetError(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		sysErr    error
		publicErr error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "hide_false",
			fields: fields{
				Config: cu.IM{
					"NT_DEV_HIDE_ERROR": false,
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				sysErr:    errors.New("error"),
				publicErr: errors.New("error"),
			},
			wantErr: true,
		},
		{
			name: "hide_true",
			fields: fields{
				Config: cu.IM{
					"NT_DEV_HIDE_ERROR": true,
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				sysErr:    errors.New("error"),
				publicErr: errors.New("error"),
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			if err := ds.SetError(tt.args.sysErr, tt.args.publicErr); (err != nil) != tt.wantErr {
				t.Errorf("DataStore.SetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var testConn func() struct {
	Alias     string
	Connected bool
	Engine    string
} = func() struct {
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
		Engine:    "sqlite",
	}
}

func TestDataStore_checkConnection(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "connected",
			fields: fields{
				Alias:  "test",
				Db:     &md.TestDriver{Config: cu.IM{}},
				Config: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "create_connection",
			fields: fields{
				Alias: "test",
				Db: &md.TestDriver{Config: cu.IM{
					"Connection": testConn,
				}},
				Config: cu.IM{
					"NT_ALIAS_TEST": "sqlite3://file::memory:?cache=shared",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_alias",
			fields: fields{
				Alias: "test",
				Db: &md.TestDriver{Config: cu.IM{
					"Connection": testConn,
				}},
				Config: cu.IM{},
			},
			wantErr: true,
		},
		{
			name: "missing_conn_str",
			fields: fields{
				Alias: "test",
				Db: &md.TestDriver{Config: cu.IM{
					"Connection": testConn,
				}},
				Config: cu.IM{},
			},
			wantErr: true,
		},
		{
			name: "create_connection_error",
			fields: fields{
				Alias: "test",
				Db: &md.TestDriver{Config: cu.IM{
					"Connection": testConn,
					"CreateConnection": func() error {
						return errors.New("error")
					},
				}},
				Config: cu.IM{
					"NT_ALIAS_TEST": "sqlite3://file::memory:?cache=shared",
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			if err := ds.checkConnection(); (err != nil) != tt.wantErr {
				t.Errorf("DataStore.checkConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataStore_StoreDataUpdate(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		update md.Update
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create_connection_error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Connection": testConn,
					"CreateConnection": func() error {
						return errors.New("error")
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				update: md.Update{},
			},
			wantErr: true,
		},
		{
			name: "update_error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Update": func(data md.Update) (int64, error) {
						return 0, errors.New("error")
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				update: md.Update{},
			},
			wantErr: true,
		},
		{
			name: "not_found",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Update": func(data md.Update) (int64, error) {
						return 0, nil
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				update: md.Update{},
			},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Update": func(data md.Update) (int64, error) {
						return 100, nil
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				update: md.Update{},
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.StoreDataUpdate(tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.StoreDataUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_StoreDataQuery(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		query    md.Query
		foundErr bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create_connection_error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Connection": testConn,
					"CreateConnection": func() error {
						return errors.New("error")
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				query: md.Query{},
			},
			wantErr: true,
		},
		{
			name: "query_error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{}, errors.New("error")
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				query: md.Query{},
			},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{}, nil
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				query:    md.Query{},
				foundErr: true,
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.StoreDataQuery(tt.args.query, tt.args.foundErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.StoreDataQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_StoreDataQueries(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		queries []md.Query
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create_connection_error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Connection": testConn,
					"CreateConnection": func() error {
						return errors.New("error")
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				queries: []md.Query{},
			},
			wantErr: true,
		},
		{
			name: "query_error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{}, errors.New("error")
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				queries: []md.Query{},
			},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{}, nil
					},
				}},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				queries: []md.Query{},
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.StoreDataQueries(tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.StoreDataQueries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_StoreDataGet(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		params   cu.IM
		foundErr bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sqlite",
			fields: fields{
				Alias: "test",
				Db: &md.TestDriver{Config: cu.IM{
					"alias":  "test",
					"engine": "sqlite",
					"Get": func(params cu.IM) ([]cu.IM, error) {
						return []cu.IM{}, nil
					},
				}},
				Config: cu.IM{
					"NT_ALIAS_TEST": "test://",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				params: cu.IM{
					"fields":     []string{"id", "name"},
					"model":      "test",
					"limit":      10,
					"offset":     10,
					"tag":        "test",
					"fieldNames": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "postgres",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"alias":  "test",
					"engine": "postgres",
				}},
			},
			args: args{
				params: cu.IM{
					"fields":     []string{"id", "name"},
					"model":      "test",
					"limit":      10,
					"offset":     10,
					"tag":        "test",
					"fieldNames": "value",
				},
			},
		},
		{
			name: "mysql",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{
					"alias":  "test",
					"engine": "mysql",
				}},
			},
			args: args{
				params: cu.IM{
					"fields":     []string{"id", "name"},
					"model":      "test",
					"limit":      10,
					"offset":     10,
					"tag":        "test",
					"fieldNames": "value",
				},
			},
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.StoreDataGet(tt.args.params, tt.args.foundErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.StoreDataGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_GetBodyData(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		modelName string
		body      io.ReadCloser
		modelData any
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
				Db: &md.TestDriver{Config: cu.IM{}},
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte([]byte(`{"id": 1, "name": "test", "test_meta": {"id": 1, "name": "test"}}`), result)
				},
			},
			args: args{
				modelName: "test",
				body:      io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "test", "test_meta": {"id": 1, "name": "test"}}`)),
				modelData: &cu.IM{},
			},
		},
		{
			name: "model convert error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{}},
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					if result == nil {
						return errors.New("error")
					}
					return cu.ConvertFromByte([]byte(`{"id": 1, "name": "test", "test_meta": {"id": 1, "name": "test"}}`), result)
				},
			},
			args: args{
				modelName: "test",
				body:      io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "test"}`)),
				modelData: nil,
			},
			wantErr: true,
		},
		{
			name: "body convert error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{}},
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					if result != nil {
						return errors.New("error")
					}
					return cu.ConvertFromByte([]byte(`{"id": 1, "name": "test", "test_meta": {"id": 1, "name": "test"}}`), result)
				},
			},
			args: args{
				modelName: "test",
				body:      io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "test"}`)),
				modelData: nil,
			},
			wantErr: true,
		},
		{
			name: "read all error",
			fields: fields{
				Db: &md.TestDriver{Config: cu.IM{}},
				ReadAll: func(r io.Reader) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
			args: args{
				modelName: "test",
				body:      io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "test"}`)),
				modelData: nil,
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, _, _, err := ds.GetBodyData(tt.args.modelName, tt.args.body, tt.args.modelData)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.GetBodyData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_SetUpdateValue(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		modelName   string
		item        cu.IM
		inputData   any
		inputFields []string
		setValue    func(modelName string, itemRow cu.IM, values cu.IM, inputName string, fieldValue interface{}) cu.IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return "", nil
				},
			},
			args: args{
				inputData:   cu.IM{"code": "123456", "addresses": []cu.IM{{"address": "123456"}}, "customer_type": "1"},
				inputFields: []string{"code", "missing"},
				setValue:    nil,
			},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test", "customer_map": {"value2": "test"}}`), nil
				},
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return JSONName, data.(cu.IM)[JSONName]
				},
			},
			args: args{
				modelName: "customer",
				inputData: cu.IM{
					"code":          "123456",
					"addresses":     []cu.IM{{"address": "123456"}},
					"customer_type": "1",
					"barcode_type":  1,
					"link_type_1":   1,
					"link_type_2":   "1",
					"shipping_time": "2024-01-01T00:00:00Z",
					"paid_date":     "2024-01-01",
					"valid_from":    md.TimeDate{},
					"default":       "value",
					"customer_map":  cu.IM{"value1": "test"},
				},
				inputFields: []string{"code", "addresses", "customer_type", "barcode_type", "link_type_1",
					"link_type_2", "shipping_time", "paid_date", "valid_from", "default", "customer_map"},
				setValue: nil,
			},
			wantErr: false,
		},
		{
			name: "shipping_time",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return JSONName, data.(cu.IM)[JSONName]
				},
			},
			args: args{
				inputData:   cu.IM{"shipping_time": md.TimeDateTime{}},
				inputFields: []string{"shipping_time"},
				setValue:    nil,
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.SetUpdateValue(tt.args.modelName, tt.args.item, tt.args.inputData, tt.args.inputFields, tt.args.setValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.SetUpdateValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_MergeMetaData(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		modelName  string
		item       cu.IM
		inputMeta  any
		metaFields []string
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
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return JSONName, data.(cu.IM)[JSONName]
				},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
			},
			args: args{
				modelName:  "customer",
				item:       cu.IM{"customer_meta": cu.IM{"id": 1, "name": "test"}},
				inputMeta:  cu.IM{"id": 1, "name": "test"},
				metaFields: []string{"id", "name"},
			},
		},
		{
			name: "error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return "", nil
				},
			},
			args: args{
				modelName:  "customer",
				item:       cu.IM{"customer_meta": cu.IM{"id": 1, "name": "test"}},
				inputMeta:  cu.IM{"id": 1, "name": "test"},
				metaFields: []string{"missing"},
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.MergeMetaData(tt.args.modelName, tt.args.item, tt.args.inputMeta, tt.args.metaFields)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.MergeMetaData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_GetDataByID(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		model    string
		id       int64
		code     string
		foundErr bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok_id",
			fields: fields{
				Config: cu.IM{},
				Db:     &md.TestDriver{Config: cu.IM{}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				model:    "customer",
				id:       1,
				code:     "",
				foundErr: false,
			},
			wantErr: false,
		},
		{
			name: "ok_code",
			fields: fields{
				Config: cu.IM{},
				Db:     &md.TestDriver{Config: cu.IM{}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				model:    "customer",
				id:       0,
				code:     "123456",
				foundErr: false,
			},
			wantErr: false,
		},
		{
			name:    "missing id or code",
			args:    args{},
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			_, err := ds.GetDataByID(tt.args.model, tt.args.id, tt.args.code, tt.args.foundErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.GetDataByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_UpdateData(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		options md.UpdateDataOptions
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
				Config: cu.IM{},
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{{"id": 1, "name": "test"}}, nil
					},
					"Update": func(data md.Update) (int64, error) {
						return 100, nil
					},
				}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return JSONName, "test"
				},
			},
			args: args{
				options: md.UpdateDataOptions{
					Model:      "customer",
					IDKey:      1,
					Code:       "123456",
					Data:       cu.IM{"name": "test", "customer_meta": cu.IM{"id": 1, "name": "test"}},
					Fields:     []string{"name"},
					MetaFields: []string{"id", "name"},
					SetValue:   nil,
				},
			},
			wantErr: false,
		},
		{
			name: "merge meta data error",
			fields: fields{
				Config: cu.IM{},
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{{"id": 1, "name": "test"}}, nil
					},
					"Update": func(data md.Update) (int64, error) {
						return 100, nil
					},
				}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), errors.New("error")
				},
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return JSONName, "test"
				},
			},
			args: args{
				options: md.UpdateDataOptions{
					Model:      "customer",
					IDKey:      1,
					Code:       "123456",
					Data:       cu.IM{"name": "test", "customer_meta": cu.IM{"id": 1, "name": "test"}},
					Fields:     []string{"name"},
					MetaFields: []string{"id", "name"},
					SetValue:   nil,
				},
			},
			wantErr: true,
		},
		{
			name: "get data by id error",
			fields: fields{
				Config: cu.IM{},
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{{"id": 1, "name": "test"}}, errors.New("error")
					},
					"Update": func(data md.Update) (int64, error) {
						return 100, nil
					},
				}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return JSONName, "test"
				},
			},
			args: args{
				options: md.UpdateDataOptions{
					Model:      "customer",
					IDKey:      1,
					Code:       "123456",
					Data:       cu.IM{"name": "test", "customer_meta": cu.IM{"id": 1, "name": "test"}},
					Fields:     []string{"name"},
					MetaFields: []string{"id", "name"},
					SetValue:   nil,
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			if _, err := ds.UpdateData(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("DataStore.UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataStore_DataDelete(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		model string
		id    int64
		code  string
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
				Config: cu.IM{},
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{{"id": 1, "name": "test"}}, nil
					},
					"Update": func(data md.Update) (int64, error) {
						return 100, nil
					},
				}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				model: "customer",
				id:    1,
				code:  "123456",
			},
			wantErr: false,
		},
		{
			name: "get data by id error",
			fields: fields{
				Config: cu.IM{},
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{{"id": 1, "name": "test"}}, errors.New("error")
					},
				}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				model: "customer",
				id:    1,
				code:  "123456",
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			if err := ds.DataDelete(tt.args.model, tt.args.id, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("DataStore.DataDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataStore_GetData(t *testing.T) {
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
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
	}
	type args struct {
		query  md.Query
		result any
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
				Config: cu.IM{},
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{{"id": 1, "name": "test"}}, nil
					},
				}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
			},
			args: args{
				query:  md.Query{},
				result: nil,
			},
			wantErr: false,
		},
		{
			name: "no data",
			fields: fields{
				Config: cu.IM{},
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{}, nil
					},
				}},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				query:  md.Query{},
				result: nil,
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
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
			}
			if err := ds.GetData(tt.args.query, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("DataStore.GetData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
