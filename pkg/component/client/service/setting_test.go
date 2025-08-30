package service

import (
	"errors"
	"log/slog"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	"golang.org/x/oauth2"
)

func TestClientService_settingData(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		ds   *api.DataStore
		user cu.IM
		in2  cu.IM
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
				user: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "config_data_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "config_data" {
								return []cu.IM{}, errors.New("error")
							}
							return []cu.IM{{"id": 1}}, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToType: func(data interface{}, result any) (err error) {
						return nil
					},
				},
				user: cu.IM{},
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
			_, err := cls.settingData(tt.args.ds, tt.args.user, tt.args.in2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.settingData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestClientService_settingResponse(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt ct.ResponseEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "form_ok",
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
									"setting": cu.IM{},
								},
							},
						},
					},
					Name: ct.FormEventOK,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "contacts", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{"name": "contact2"},
						"event": ct.FormEventOK},
				},
			},
			wantErr: false,
		},
		{
			name: "form_event",
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
									"setting": cu.IM{},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "contacts", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{"name": "contact2"},
						"event": ct.FormEventOK},
				},
			},
			wantErr: false,
		},
		{
			name: "side_editor_save_ok",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
								"editor": cu.IM{},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_save",
				},
			},
			wantErr: false,
		},
		{
			name: "theme",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "theme",
						"value": "dark",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "lang_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 0, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "lang",
						"value": "en",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "page_size",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "page_size",
						"value": 10,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "orientation",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "orientation",
						"value": "landscape",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "pagination",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "pagination",
						"value": 10,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "history",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "history",
						"value": 5,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "export_sep",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "export_sep",
						"value": ";",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "password",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "password",
						"value": "123456",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "confirm",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
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
									"setting": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "confirm",
						"value": "123456",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "change_password",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 12345, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
						CreatePasswordHash: func(password string) (hash string, err error) {
							return "123456", nil
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
									"setting": cu.IM{
										"password": "123456",
										"confirm":  "123456",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"code": "123456",
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name": "change_password",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "change_password_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 0, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return []byte{}, nil
						},
						CreatePasswordHash: func(password string) (hash string, err error) {
							return "123456", nil
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
									"setting": cu.IM{
										"password": "123456",
										"confirm":  "123456",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"code": "123456",
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name": "change_password",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing",
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
									"setting": cu.IM{
										"password": "123456",
										"confirm":  "123456",
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"code": "123456",
								"auth_map": cu.IM{
									"theme": "dark",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name": "missing",
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
			_, err := cls.settingResponse(tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.settingResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
