package service

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	cp "github.com/nervatura/nervatura/v6/pkg/component/client/component"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	"golang.org/x/oauth2"
)

func TestNewClientService(t *testing.T) {
	type args struct {
		config     cu.IM
		appLog     *slog.Logger
		memSession map[string]md.MemoryStore
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				config: cu.IM{
					"NT_GOOGLE_CLIENT_ID":     "1234567890",
					"NT_GOOGLE_CLIENT_SECRET": "1234567890",
				},
				appLog:     slog.Default(),
				memSession: map[string]md.MemoryStore{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewClientService(tt.args.config, tt.args.appLog, tt.args.memSession)
		})
	}
}

func TestClientService_GetClient(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
		UI           *cp.ClientComponent
	}
	type args struct {
		host      string
		sessionID string
		eventURL  string
		lang      string
		theme     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				Config: cu.IM{},
				AuthConfigs: map[string]*oauth2.Config{
					"NT_GOOGLE_CLIENT": {
						ClientID:     "1234567890",
						ClientSecret: "1234567890",
					},
				},
				AppLog: slog.Default(),
				UI:     cp.NewClientComponent(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
				UI:           tt.fields.UI,
			}
			cls.GetClient(tt.args.host, tt.args.sessionID, tt.args.eventURL, tt.args.lang, tt.args.theme)
		})
	}
}

func TestClientService_LoadSession(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
		UI           *cp.ClientComponent
	}
	type args struct {
		sessionID string
	}
	client := &ct.Client{
		BaseComponent: ct.BaseComponent{
			Data: cu.IM{},
		},
		Ticket: ct.Ticket{
			User: cu.IM{"uid": "UID012345"},
		},
	}
	ses := api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &md.TestDriver{Config: cu.IM{}},
	}
	ses.SaveSession("SES012345", client)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "mem",
			fields: fields{
				Session: &ses,
				UI:      cp.NewClientComponent(),
			},
			args: args{
				sessionID: "SES012345",
			},
			wantErr: false,
		},
		{
			name: "file",
			fields: fields{
				Session: &api.SessionService{
					Config: api.SessionConfig{
						Method: md.SessionMethodFile,
					},
					ReadFile: func(name string) ([]byte, error) {
						app, _ := json.Marshal(client)
						return []byte(app), nil
					},
					CreateFile: func(name string) (*os.File, error) {
						return os.NewFile(0, name), nil
					},
					CreateDir: func(name string, perm fs.FileMode) error {
						return nil
					},
					FileStat: func(name string) (fs.FileInfo, error) {
						return nil, nil
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						bt, _ := cu.ConvertToByte(client)
						cu.ConvertFromByte(bt, result)
						return nil
					},
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				sessionID: "SES012345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
				UI:           tt.fields.UI,
			}
			_, err := cls.LoadSession(tt.args.sessionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.LoadSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_AuthUser(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		database string
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "demo",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "name": "test"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
					}
				},
			},
			args: args{
				database: "demo",
				username: "admin",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.AuthUser(tt.args.database, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.AuthUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_userLogin(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		database string
		username string
		password string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser cu.IM
		wantErr  bool
	}{
		{
			name: "demo",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "name": "test"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
					}
				},
			},
			args: args{
				database: "demo",
				username: "admin",
				password: "****",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.userLogin(tt.args.database, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.userLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

/*
func TestClientService_moduleData(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt    ct.ResponseEvent
		ds     *api.DataStore
		mKey   string
		user   cu.IM
		params cu.IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "customer",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
				},
				mKey: "customer",
				user: cu.IM{},
				params: cu.IM{
					"customer_id": 1,
				},
			},
			wantErr: false,
		},
		{
			name: "product",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToType: func(data interface{}, result any) (err error) {
						return nil
					},
				},
				mKey: "product",
				user: cu.IM{},
				params: cu.IM{
					"product_id": 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.moduleData(tt.args.evt, tt.args.ds, tt.args.mKey, tt.args.user, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.moduleData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
*/

/*
func TestClientService_moduleResponse(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		mKey string
		evt  ct.ResponseEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "search",
			args: args{
				mKey: "search",
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "customer",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{}},
					}
				},
			},
			args: args{
				mKey: "customer",
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "demo",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "product",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{}},
					}
				},
			},
			args: args{
				mKey: "product",
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"product": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "demo",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "tool",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{}},
					}
				},
			},
			args: args{
				mKey: "tool",
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "demo",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "project",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{}},
					}
				},
			},
			args: args{
				mKey: "project",
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "demo",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "employee",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{}},
					}
				},
			},
			args: args{
				mKey: "employee",
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"employee": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "demo",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.moduleResponse(tt.args.mKey, tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.moduleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
*/

func TestClientService_evtMsg(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		name        string
		triggerName string
		value       string
		toastType   string
		timeout     int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "search",
			args: args{
				name:        "search",
				triggerName: "search",
				value:       "search",
				toastType:   "success",
				timeout:     1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			cls.evtMsg(tt.args.name, tt.args.triggerName, tt.args.value, tt.args.toastType, tt.args.timeout)
		})
	}
}

func TestClientService_SetEditor(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt    ct.ResponseEvent
		module string
		params cu.IM
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "module_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, errors.New("error")
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer", "row": cu.IM{"id": 1234}},
				},
				module: "customer",
				params: cu.IM{},
			},
		},
		{
			name: "module_ok",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer", "row": cu.IM{"id": 1234}},
				},
				module: "customer",
				params: cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			cls.setEditor(tt.args.evt, tt.args.module, tt.args.params)
		})
	}
}

func TestClientService_searchEvent(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
		Modules      map[string]ServiceModule
		UI           *cp.ClientComponent
	}
	type args struct {
		evt ct.ResponseEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "module_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"search": NewSearchService(&ClientService{
						Config: cu.IM{},
						AppLog: slog.Default(),
						NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
							return &api.DataStore{
								Db: &md.TestDriver{Config: cu.IM{
									"Query": func(queries []md.Query) ([]cu.IM, error) {
										return []cu.IM{}, errors.New("error")
									},
								}},
								Config: config,
								AppLog: appLog,
							}
						},
						UI: cp.NewClientComponent(),
					}),
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer", "row": cu.IM{"id": 1234}},
				},
			},
		},
		{
			name: "module_ok",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"search": NewSearchService(&ClientService{
						Config: cu.IM{},
						AppLog: slog.Default(),
						NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
							return &api.DataStore{
								Db: &md.TestDriver{Config: cu.IM{
									"Query": func(queries []md.Query) ([]cu.IM, error) {
										return []cu.IM{{"id": 1}}, nil
									},
								}},
								Config: config,
								AppLog: appLog,
							}
						},
						UI: cp.NewClientComponent(),
					}),
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer", "row": cu.IM{"id": 1234}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
				Modules:      tt.fields.Modules,
				UI:           tt.fields.UI,
			}
			cls.searchEvent(tt.args.evt)
		})
	}
}

func TestClientService_insertPrintQueue(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		ds   *api.DataStore
		data cu.IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "insert_print_queue",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			if err := cls.insertPrintQueue(tt.args.ds, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ClientService.insertPrintQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientService_MainResponse(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
		Modules      map[string]ServiceModule
		UI           *cp.ClientComponent
	}
	type args struct {
		evt ct.ResponseEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "search_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"search": NewSearchService(&ClientService{
						Config: cu.IM{},
						AppLog: slog.Default(),
						NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
							return &api.DataStore{
								Db: &md.TestDriver{Config: cu.IM{
									"Query": func(queries []md.Query) ([]cu.IM, error) {
										return []cu.IM{}, errors.New("error")
									},
								}},
								Config: config,
								AppLog: appLog,
							}
						},
						UI: cp.NewClientComponent(),
					}),
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.BrowserEventSearch,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
				},
			},
		},
		{
			name: "browser",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"search": NewSearchService(&ClientService{
						Config: cu.IM{},
						AppLog: slog.Default(),
						NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
							return &api.DataStore{
								Db: &md.TestDriver{Config: cu.IM{
									"Query": func(queries []md.Query) ([]cu.IM, error) {
										return []cu.IM{}, nil
									},
								}},
								Config: config,
								AppLog: appLog,
							}
						},
						UI: cp.NewClientComponent(),
					}),
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.BrowserEventSearch,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
				},
			},
		},
		{
			name: "search",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"search": NewSearchService(&ClientService{
						Config: cu.IM{},
						AppLog: slog.Default(),
						NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
							return &api.DataStore{
								Db: &md.TestDriver{Config: cu.IM{
									"Query": func(queries []md.Query) ([]cu.IM, error) {
										return []cu.IM{}, nil
									},
								}},
								Config: config,
								AppLog: appLog,
							}
						},
						UI: cp.NewClientComponent(),
					}),
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.SearchEventSearch,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
				},
			},
		},
		{
			name: "table_edit_cell_code",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "code", "row": cu.IM{"id": 1234, "editor": "customer"}},
				},
			},
		},
		{
			name: "table_edit_cell_ref_code",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"customer": &CustomerService{
						cls: &ClientService{
							Config: cu.IM{},
							AppLog: slog.Default(),
							NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
								return &api.DataStore{
									Db: &md.TestDriver{Config: cu.IM{
										"Query": func(queries []md.Query) ([]cu.IM, error) {
											return []cu.IM{}, errors.New("error")
										},
									}},
									Config: config,
									AppLog: appLog,
								}
							},
						},
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer_code", "row": cu.IM{"id": 1234, "editor": "customer"}},
				},
			},
		},
		{
			name: "table_edit_cell_url",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "value", "row": cu.IM{"id": 1234, "field_type": "FIELD_URL"}},
				},
			},
		},
		{
			name: "table_edit_cell_model",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventEditCell,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "value", "row": cu.IM{"id": 1234, "field_type": "FIELD_CUSTOMER"}},
				},
			},
		},
		{
			name: "table_add_item",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventAddItem,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "transmovement_formula",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer", "row": cu.IM{"id": 1234}},
				},
			},
		},
		{
			name: "table_add_transitem",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.TableEventAddItem,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "transitem",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer", "row": cu.IM{"id": 1234}},
				},
			},
		},
		{
			name: "browser_edit_row",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.BrowserEventEditRow,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer", "editor_id": 1234, "editor": "customer", "editor_view": "customer"},
				},
			},
		},
		{
			name: "search_selected",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, errors.New("error")
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.SearchEventSelected,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer_simple",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"fieldname": "customer",
						"row": cu.IM{"editor_id": 1234, "editor": "customer", "editor_view": "customer"}},
				},
			},
		},
		{
			name: "client_module_search",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventModule,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "search",
				},
			},
		},
		{
			name: "client_module_editor",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventModule,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "customer",
				},
			},
		},
		{
			name: "client_module_bookmark",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventModule,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"bookmarks": []cu.IM{
									{
										"label":         "test",
										"code":          "test",
										"bookmark_type": "browser",
									},
								},
							},
						},
					},
					Value: "bookmark",
				},
			},
		},
		{
			name: "client_module_none",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventModule,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "search",
				},
			},
		},
		{
			name: "client_module_dirty",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventModule,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"dirty": true,
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "search",
				},
			},
		},
		{
			name: "form_ok_set_search",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventOK,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"data": cu.IM{"next": "set_search"}, "value": cu.IM{}},
				},
			},
		},
		{
			name: "form_ok_set_setting",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ReadFile: func(name string) ([]byte, error) {
							return []byte(`{"meta": {"report_key": "ntr_customer_en", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}`), nil
						},
						ConvertFromByte: func(data []byte, result interface{}) error {
							return cu.ConvertFromByte([]byte(`{"meta": {"report_key": "ntr_customer_en", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}`), result)
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventOK,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"data": cu.IM{"next": "set_setting"}, "value": cu.IM{}},
				},
			},
		},
		{
			name: "form_ok_set_bookmark",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventOK,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"bookmarks": []cu.IM{
									{
										"label":         "test",
										"code":          "test",
										"bookmark_type": "browser",
									},
								},
							},
						},
					},
					Value: cu.IM{"data": cu.IM{"next": "set_bookmark"}, "value": cu.IM{}},
				},
			},
		},
		{
			name: "form_ok_trans_new",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"trans": &TransService{
						cls: &ClientService{
							Config: cu.IM{},
							AppLog: slog.Default(),
							NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
								return &api.DataStore{
									Db: &md.TestDriver{Config: cu.IM{}},
								}
							},
						},
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventOK,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"data": cu.IM{"next": "transitem_new"}, "value": cu.IM{}},
				},
			},
		},
		{
			name: "form_ok_default",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventOK,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "customer",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
				},
			},
		},
		{
			name: "form_change",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventChange,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "missing",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
				},
			},
		},
		{
			name: "form_change_bookmark_filter",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "user_name": "test", "auth_meta": cu.IM{
									"tags": []string{"test"},
									"bookmarks": []cu.IM{
										{
											"label": "test",
										},
									},
								}}}, nil
							},
							"Update": func(update md.Update) (int64, error) {
								return 1, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return ut.ConvertToType(data, result)
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventChange,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "missing",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"bookmarks": []cu.IM{
									{
										"label":         "test",
										"code":          "test",
										"bookmark_type": "browser",
									},
								},
							},
						},
					},
					Value: cu.IM{"data": cu.IM{},
						"name":  "bookmark",
						"event": "list_filter_change",
						"value": cu.IM{
							"value": cu.IM{
								"value": cu.IM{
									"index": 0,
								},
							},
						}},
				},
			},
		},
		{
			name: "form_change_delete_bookmark",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "user_name": "test", "auth_meta": cu.IM{
									"tags": []string{"test"},
									"bookmarks": []cu.IM{
										{
											"label": "test",
										},
									},
								}}}, nil
							},
							"Update": func(update md.Update) (int64, error) {
								return 1, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return ut.ConvertToType(data, result)
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventChange,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "missing",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"bookmarks": []cu.IM{
									{
										"label":         "test",
										"code":          "test",
										"bookmark_type": "browser",
									},
								},
							},
						},
					},
					Value: cu.IM{"data": cu.IM{},
						"name":  "bookmark",
						"event": "list_delete",
						"value": cu.IM{
							"value": cu.IM{
								"value": cu.IM{
									"index": 0,
								},
							},
						}},
				},
			},
		},
		{
			name: "form_change_delete_bookmark_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "user_name": "test", "auth_meta": cu.IM{
									"tags": []string{"test"},
									"bookmarks": []cu.IM{
										{
											"label": "test",
										},
									},
								}}}, nil
							},
							"Update": func(update md.Update) (int64, error) {
								return 0, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return ut.ConvertToType(data, result)
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventChange,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "missing",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"bookmarks": []cu.IM{
									{
										"label":         "test",
										"code":          "test",
										"bookmark_type": "browser",
									},
								},
							},
						},
					},
					Value: cu.IM{"data": cu.IM{},
						"name":  "bookmark",
						"event": "list_delete",
						"value": cu.IM{
							"value": cu.IM{
								"value": cu.IM{
									"index": 0,
								},
							},
						}},
				},
			},
		},
		{
			name: "form_change_set_bookmark",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "user_name": "test", "auth_meta": cu.IM{
									"tags": []string{"test"},
									"bookmarks": []cu.IM{
										{
											"label": "test",
										},
									},
								}}}, nil
							},
							"Update": func(update md.Update) (int64, error) {
								return 1, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return ut.ConvertToType(data, result)
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventChange,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "missing",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"bookmarks": []cu.IM{
									{
										"label":         "test",
										"code":          "test",
										"bookmark_type": "browser",
									},
								},
							},
						},
					},
					Value: cu.IM{"data": cu.IM{},
						"name":  "bookmark",
						"event": ct.ListEventEditItem,
						"value": cu.IM{
							"value": cu.IM{
								"value": cu.IM{
									"index": 0,
								},
							},
						}},
				},
			},
		},
		{
			name: "form_change_set_bookmark_editor",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "user_name": "test", "auth_meta": cu.IM{
									"tags": []string{"test"},
									"bookmarks": []cu.IM{
										{
											"label": "test",
										},
									},
								}}}, nil
							},
							"Update": func(update md.Update) (int64, error) {
								return 1, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return ut.ConvertToType(data, result)
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.FormEventChange,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "missing",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"bookmarks": []cu.IM{
									{
										"label":         "test",
										"code":          "test",
										"bookmark_type": md.BookmarkTypeEditor.String(),
										"key":           "customer",
									},
								},
							},
						},
					},
					Value: cu.IM{"data": cu.IM{},
						"name":  "bookmark",
						"event": ct.ListEventEditItem,
						"value": cu.IM{
							"value": cu.IM{
								"value": cu.IM{
									"index": 0,
								},
							},
						}},
				},
			},
		},
		{
			name: "client_form",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventForm,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "customer",
									"form": cu.IM{
										"key": "contacts",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
				},
			},
		},
		{
			name: "client_side_menu",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
				Modules: map[string]ServiceModule{
					"customer": &CustomerService{
						cls: &ClientService{
							Config: cu.IM{},
							AppLog: slog.Default(),
							NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
								return &api.DataStore{
									Db: &md.TestDriver{Config: cu.IM{}},
								}
							},
						},
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventSideMenu,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key": "customer",
									"customer": cu.IM{
										"id": 12345,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "editor_save",
				},
			},
		},
		{
			name: "browser_bookmark",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.BrowserEventBookmark,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "",
				},
			},
		},
		{
			name: "editor_field",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key":      "customer",
									"customer": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"name": "notes", "value": "..."},
				},
			},
		},
		{
			name: "editor_field_setting",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key":     "setting",
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"name": "lang", "value": "en"},
				},
			},
		},
		{
			name: "editor_field_cell",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.EditorEventField,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"key":      "customer",
									"customer": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"name": "table_edit_cell", "value": cu.IM{"id": 1234}},
				},
			},
		},
		{
			name: "login_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, errors.New("error")
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.LoginEventLogin,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"username": "admin", "password": "admin"},
				},
			},
		},
		{
			name: "login_ok",
			fields: fields{
				Config: cu.IM{"NT_API_KEY": "test"},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ComparePasswordAndHash: func(password string, hash string) (err error) {
							return nil
						},
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
						CreateLoginToken: func(code, userName, database string, config cu.IM) (result string, err error) {
							return "test", nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.LoginEventLogin,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: cu.IM{"username": "admin", "password": "test"},
				},
			},
		},
		{
			name: "login_auth_err",
			fields: fields{
				Config: cu.IM{"NT_API_KEY": "test"},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.LoginEventAuth,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "google",
				},
			},
		},
		{
			name: "login_auth_ok",
			fields: fields{
				Config: cu.IM{"NT_API_KEY": "test"},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				AuthConfigs: map[string]*oauth2.Config{
					"testaut": {
						ClientID:     "TEST_CLIENT_ID",
						ClientSecret: "TEST_CLIENT_SECRET",
						RedirectURL:  "REDIRECT_URL",
						Scopes:       []string{"email"},
						Endpoint: oauth2.Endpoint{
							AuthURL: "testAuthCodeURL",
						},
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.LoginEventAuth,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "testaut",
				},
			},
		},
		{
			name: "client_logout",
			fields: fields{
				Config: cu.IM{"NT_API_KEY": "test"},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventLogOut,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "",
				},
			},
		},
		{
			name: "other_event",
			fields: fields{
				Config: cu.IM{"NT_API_KEY": "test"},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventTheme,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
				Modules:      tt.fields.Modules,
				UI:           tt.fields.UI,
			}
			cls.MainResponse(tt.args.evt)
		})
	}
}

func TestClientService_editorFormTags(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		params cu.IM
		evt    ct.ResponseEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "client_form_change_add",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				params: cu.IM{"row_field": "tags"},
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"contacts": []cu.IM{
											{
												"id":   1,
												"name": "contact1",
												"tags": []string{"tag1", "tag2"},
											},
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "contacts", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{},
						"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventAddItem},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_change_delete",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				params: cu.IM{"row_field": "tags"},
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"contacts": []cu.IM{
											{
												"id":   1,
												"name": "contact1",
												"tags": []string{"tag1", "tag2"},
											},
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{"form": cu.IM{"index": 0, "key": "contacts", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{
							"row":   cu.IM{"tag": "tag1"},
							"index": 0,
						},
						"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventDelete},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_change_delete_meta",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				params: cu.IM{"row_field": "tags"},
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"view": []cu.IM{
											{
												"id":   1,
												"name": "contact1",
												"view_meta": cu.IM{
													"tags": []string{"tag1", "tag2"},
												},
											},
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "view",
								"data": cu.IM{"view_meta": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{
							"row":   cu.IM{"tag": "tag1"},
							"index": 0,
						},
						"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventDelete},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_change_skip",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				params: cu.IM{"row_field": "tags"},
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"contacts": []cu.IM{
											{
												"id":   1,
												"name": "contact1",
												"tags": []string{"tag1", "tag2"},
											},
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "contacts", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{},
						"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventEditItem},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.editorFormTags(tt.args.params, tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.editorFormTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_editorFormOK(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt          ct.ResponseEvent
		rows         []cu.IM
		customValues map[string]func(value any)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "client_form_ok",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"product": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"events": []cu.IM{
											{
												"id":   1,
												"name": "contact1",
												"tags": []string{"tag1", "tag2"},
											},
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{"form": cu.IM{"index": 0, "key": "events",
							"data": cu.IM{
								"tags": []string{"tag1", "tag2"},
								"price_meta": cu.IM{
									"tags": []string{"tag1", "tag2"},
								},
								"customer_code": "CUS0000001",
								"price_value":   1000,
							}},
						},
						"value": cu.IM{"price_value": 1000, "qty": 100, "name": "event1"},
						"event": ct.FormEventOK},
				},
				rows: []cu.IM{
					{
						"id":            1,
						"name":          "event1",
						"tags":          []string{"tag1", "tag2"},
						"price_value":   1000,
						"qty":           100,
						"customer_code": "CUS0000001",
					},
				},
				customValues: map[string]func(value any){
					"frm_price_value": func(value any) {
					},
					"base_price_value": func(value any) {
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.editorFormOK(tt.args.evt, tt.args.rows, tt.args.customValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.editorFormOK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_addMapField(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt          ct.ResponseEvent
		editorMap    cu.IM
		resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "table_add_item_map_field_tool",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool":      cu.IM{"id": 12345},
									"view":      "maps",
									"map_field": "ref_tool",
									"config_map": []cu.IM{
										{"field_name": "ref_tool", "field_type": "FIELD_CUSTOMER"},
									},
								},
							},
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
				},
				editorMap: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_map_field_enum",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool":      cu.IM{"id": 12345},
									"view":      "maps",
									"map_field": "demo_string",
									"config_map": []cu.IM{
										{"field_name": "demo_string", "field_type": "FIELD_ENUM", "tags": []string{"tag1", "tag2"}},
									},
								},
							},
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
				},
				editorMap: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.addMapField(tt.args.evt, tt.args.editorMap, tt.args.resultUpdate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.addMapField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_updateMapField(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt          ct.ResponseEvent
		editorMap    cu.IM
		resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "table_form_update_validate_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, errors.New("error")
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
										"id": 12345,
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventFormUpdate,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "customer_ref", "field_type": "FIELD_CUSTOMER", "value": "CUS12345"}, "index": 0, "view": "events"}},
				},
				editorMap: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "table_form_update_ok",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, errors.New("error")
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
										"id": 12345,
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventFormUpdate,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
				},
				editorMap: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.updateMapField(tt.args.evt, tt.args.editorMap, tt.args.resultUpdate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.updateMapField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_showReportSelector(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt     ct.ResponseEvent
		refType string
		code    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "editor_report",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{"id": 12345},
									"config_report": []cu.IM{
										{"report_key": "report_key"},
									},
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_report",
				},
				refType: "TOOL",
				code:    "TOOL12345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.showReportSelector(tt.args.evt, tt.args.refType, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.showReportSelector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_editorTags(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt          ct.ResponseEvent
		editorMeta   cu.IM
		resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "tag_delete",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "tags",
						"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "addresses"},
						"event": ct.ListEventDelete,
					},
				},
				editorMeta: cu.IM{
					"tags": []string{"tag1", "tag2"},
				},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "tag_add",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "tags",
						"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "addresses"},
						"event": ct.ListEventAddItem,
					},
				},
				editorMeta: cu.IM{
					"tags": []string{"tag1", "tag2"},
				},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.editorTags(tt.args.evt, tt.args.editorMeta, tt.args.resultUpdate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.editorTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_editorCodeSelector(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
		Modules      map[string]ServiceModule
		UI           *cp.ClientComponent
	}
	type args struct {
		evt          ct.ResponseEvent
		editor       string
		codeType     string
		editorData   cu.IM
		resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "customer_code_search_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
				Modules: map[string]ServiceModule{
					"search": NewSearchService(&ClientService{
						Config: cu.IM{},
						AppLog: slog.Default(),
						NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
							return &api.DataStore{
								Db: &md.TestDriver{Config: cu.IM{
									"Query": func(queries []md.Query) ([]cu.IM, error) {
										return []cu.IM{}, errors.New("error")
									},
								}},
								Config: config,
								AppLog: appLog,
							}
						},
						UI: cp.NewClientComponent(),
					}),
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{
										"id": 12345,
										"project_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"project_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "project",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_code",
						"value": "value",
						"event": ct.SelectorEventSearch,
					},
				},
				codeType:   "customer",
				editorData: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "customer_code_search_ok",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
				Modules: map[string]ServiceModule{
					"search": NewSearchService(&ClientService{
						Config: cu.IM{},
						AppLog: slog.Default(),
						NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
							return &api.DataStore{
								Db: &md.TestDriver{Config: cu.IM{
									"Query": func(queries []md.Query) ([]cu.IM, error) {
										return []cu.IM{{"code": "value"}}, nil
									},
								}},
								Config: config,
								AppLog: appLog,
							}
						},
						UI: cp.NewClientComponent(),
					}),
				},
				UI: cp.NewClientComponent(),
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{
										"id": 12345,
										"project_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"project_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "project",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_code",
						"value": "value",
						"event": ct.SelectorEventSearch,
					},
				},
				codeType:   "customer",
				editorData: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "customer_code_selected",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{
										"id": 12345,
										"project_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"project_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "project",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "meta": cu.IM{}}},
					},
				},
				codeType:   "customer",
				editorData: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "customer_code_delete",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{
										"id": 12345,
										"project_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"project_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "project",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_code",
						"event": ct.SelectorEventDelete,
						"value": cu.IM{"row": cu.IM{"id": 12345, "meta": cu.IM{}}},
					},
				},
				codeType:   "customer",
				editorData: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "customer_code_link_dirty",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{
										"id": 12345,
										"project_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"project_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":  "project",
									"dirty": true,
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_code",
						"event": ct.SelectorEventLink, "value": ct.SelectOption{Value: "code", Text: "code"},
					},
				},
				codeType:   "customer",
				editorData: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "customer_code_link",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"code": "value"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{
										"id": 12345,
										"project_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"project_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "project",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_code",
						"event": ct.SelectorEventLink, "value": ct.SelectOption{Value: "code", Text: "code"},
					},
				},
				codeType:   "customer",
				editorData: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "customer_code_show_modal",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"code": "value"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"project": cu.IM{
										"id": 12345,
										"project_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"project_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "project",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_code",
						"event": ct.SelectorEventShowModal, "value": "",
					},
				},
				codeType:   "customer",
				editorData: cu.IM{},
				resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
					return ct.ResponseEvent{}, nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
				Modules:      tt.fields.Modules,
				UI:           tt.fields.UI,
			}
			_, err := cls.editorCodeSelector(tt.args.evt, tt.args.editor, tt.args.codeType, tt.args.editorData, tt.args.resultUpdate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.editorCodeSelector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_codeName(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		config cu.IM
		appLog *slog.Logger
		// Named input parameters for target function.
		ds         *api.DataStore
		code       string
		model      string
		want       string
		memSession map[string]md.MemoryStore
	}{
		{
			name: "success",
			ds: &api.DataStore{
				Db: &md.TestDriver{Config: cu.IM{
					"Query": func(queries []md.Query) ([]cu.IM, error) {
						return []cu.IM{{"customer_name": "name"}}, nil
					},
				}},
			},
			code:       "value",
			model:      "customer",
			want:       "name",
			memSession: map[string]md.MemoryStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := NewClientService(tt.config, tt.appLog, tt.memSession)
			got := cls.codeName(tt.ds, tt.code, tt.model)
			if got != tt.want {
				t.Errorf("codeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
