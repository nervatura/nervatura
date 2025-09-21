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
		{
			name: "config_map_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "config_map" {
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
		{
			name: "config_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "config" {
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
		{
			name: "currency_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "currency" {
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
		{
			name: "tax_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "tax" {
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
		{
			name: "auth_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "auth" {
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
			name: "config_map_edit",
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
						"name":  "config_map",
						"event": ct.ListEventEditItem,
						"value": cu.IM{
							"row": cu.IM{
								"id": 1,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "config_map_delete",
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
						"name":  "config_map",
						"event": ct.ListEventDelete,
						"value": cu.IM{
							"row": cu.IM{
								"id":   1,
								"code": "123456",
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "config_map_add",
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
						"name":  "config_map",
						"event": ct.ListEventAddItem,
						"value": cu.IM{
							"row": cu.IM{
								"id":   1,
								"code": "123456",
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "auth_edit",
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
						"name":  "auth",
						"event": ct.ListEventEditItem,
						"value": cu.IM{
							"row": cu.IM{
								"id": 1,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "auth_add",
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
						"name":  "auth",
						"event": ct.ListEventAddItem,
						"value": cu.IM{
							"row": cu.IM{
								"id":   1,
								"code": "123456",
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_row_selected",
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
						"name": ct.TableEventRowSelected,
						"value": cu.IM{
							"row": cu.IM{
								"id":   1,
								"code": "123456",
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_config_data_edit",
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
									"view": "config_data",
									"config_values": []cu.IM{{
										"id":   1,
										"code": "123456",
										"data": cu.IM{},
									}},
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
						"name": ct.TableEventFormUpdate,
						"value": cu.IM{
							"row": cu.IM{
								"id":          1,
								"config_code": "123456",
								"config_key":  "config_key",
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_config_data_new",
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
									"view": "config_data",
									"config_values": []cu.IM{{
										"id":   0,
										"code": "123456",
										"data": cu.IM{},
									}},
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
						"name": ct.TableEventFormUpdate,
						"value": cu.IM{
							"row": cu.IM{
								"id":          0,
								"config_code": "123456",
								"config_key":  "config_key",
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_currency_edit",
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
									"view": "currency",
									"currency": []cu.IM{{
										"id":            1,
										"code":          "USD",
										"currency_meta": cu.IM{},
									}},
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
						"name": ct.TableEventFormUpdate,
						"value": cu.IM{
							"row": cu.IM{
								"id":          1,
								"code":        "USD",
								"description": "United States Dollar",
								"digit":       2,
								"cash_round":  0,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_currency_new",
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
									"view": "currency",
									"currency": []cu.IM{{
										"id":            0,
										"code":          "USD",
										"currency_meta": cu.IM{},
									}},
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
						"name": ct.TableEventFormUpdate,
						"value": cu.IM{
							"row": cu.IM{
								"id":          0,
								"code":        "USD",
								"description": "United States Dollar",
								"digit":       2,
								"cash_round":  0,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_tax_edit",
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
									"view": "tax",
									"tax": []cu.IM{{
										"id":       1,
										"code":     "VAT",
										"tax_meta": cu.IM{},
									}},
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
						"name": ct.TableEventFormUpdate,
						"value": cu.IM{
							"row": cu.IM{
								"id":          1,
								"code":        "VAT",
								"description": "Value Added Tax",
								"rate_value":  0.2,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_tax_new",
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
									"view": "tax",
									"tax": []cu.IM{{
										"id":       0,
										"code":     "VAT",
										"tax_meta": cu.IM{},
									}},
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
						"name": ct.TableEventFormUpdate,
						"value": cu.IM{
							"row": cu.IM{
								"id":          0,
								"code":        "VAT",
								"description": "Value Added Tax",
								"rate_value":  0.2,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_delete_tax",
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
									"view": "tax",
									"tax": []cu.IM{{
										"id":       1,
										"code":     "VAT",
										"tax_meta": cu.IM{},
									}},
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
						"name": ct.TableEventFormDelete,
						"value": cu.IM{
							"row": cu.IM{
								"id":          1,
								"code":        "VAT",
								"description": "Value Added Tax",
								"rate_value":  0.2,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_delete_currency",
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
									"view": "currency",
									"currency": []cu.IM{{
										"id":            1,
										"code":          "USD",
										"currency_meta": cu.IM{},
									}},
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
						"name": ct.TableEventFormDelete,
						"value": cu.IM{
							"row": cu.IM{
								"id":          1,
								"code":        "USD",
								"description": "United States Dollar",
								"digit":       2,
								"cash_round":  0,
							},
							"index": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_currency",
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
									"view": "currency",
									"currency": []cu.IM{{
										"id":            1,
										"code":          "USD",
										"currency_meta": cu.IM{},
									}},
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
						"name":  ct.TableEventAddItem,
						"value": cu.IM{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_tax",
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
									"view": "tax",
									"tax": []cu.IM{{
										"id":       1,
										"code":     "VAT",
										"tax_meta": cu.IM{},
									}},
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
						"name":  ct.TableEventAddItem,
						"value": cu.IM{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_ok_config_map",
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
									"config_values": []cu.IM{{
										"id":   1,
										"code": "123456",
										"data": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{
								"index": 0, "key": "config_map",
								"data": cu.IM{
									"id":          1,
									"code":        "123456",
									"field_name":  "field_name",
									"field_type":  "field_type",
									"description": "description",
								},
							},
							"data": cu.IM{"code": "123456"}},
						"value": cu.IM{"code": "123456"},
						"event": ct.FormEventOK,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_ok_auth",
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
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_map":  cu.IM{},
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{
								"index": 0, "key": "auth",
								"data": cu.IM{
									"id":          1,
									"code":        "123456",
									"user_name":   "user_name",
									"user_group":  "user_group",
									"disabled":    false,
									"auth_meta":   cu.IM{},
									"auth_map":    cu.IM{},
									"description": "description",
								},
							},
							"data": cu.IM{"code": "123456"}},
						"value": cu.IM{"code": "123456"},
						"event": ct.FormEventOK,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_ok_auth_new",
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
									"auth": []cu.IM{{
										"id":        0,
										"code":      "123456",
										"auth_map":  cu.IM{},
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{
								"index": 0, "key": "auth",
								"data": cu.IM{
									"id":          0,
									"code":        "123456",
									"user_name":   "user_name",
									"user_group":  "user_group",
									"disabled":    false,
									"auth_meta":   cu.IM{},
									"auth_map":    cu.IM{},
									"description": "description",
								},
							},
							"data": cu.IM{"code": "123456"}},
						"value": cu.IM{"code": "123456"},
						"event": ct.FormEventOK,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_cancel_config_map_delete",
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
									"config_values": []cu.IM{{
										"id":   1,
										"code": "",
										"data": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{
								"index": 0, "key": "config_map",
								"data": cu.IM{
									"id":          1,
									"code":        "",
									"field_name":  "field_name",
									"field_type":  "field_type",
									"description": "description",
								},
							},
							"data": cu.IM{"code": ""}},
						"value": cu.IM{"code": "", "form_delete": true},
						"event": ct.FormEventCancel,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_cancel_auth_delete",
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
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_map":  cu.IM{},
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{
								"index": 0, "key": "auth",
								"data": cu.IM{
									"id":          1,
									"code":        "123456",
									"user_name":   "user_name",
									"user_group":  "user_group",
									"disabled":    false,
									"auth_meta":   cu.IM{},
									"auth_map":    cu.IM{},
									"description": "description",
								},
							},
							"data": cu.IM{"code": "123456"}},
						"value": cu.IM{"code": "123456", "form_delete": true},
						"event": ct.FormEventCancel,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_invalid",
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
									"config_values": []cu.IM{{
										"id":   1,
										"code": "",
										"data": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{
								"index": 0, "key": "config_map",
								"data": cu.IM{
									"id":          1,
									"code":        "",
									"field_name":  "field_name",
									"field_type":  "field_type",
									"description": "description",
								},
							},
							"data": cu.IM{"code": ""}},
						"value": cu.IM{"code": "", "form_delete": true},
						"event": "invalid",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_config_map_tags",
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
									"config_values": []cu.IM{{
										"id":   1,
										"code": "123456",
										"data": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "config_map",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "value",
						"event": ct.FormEventChange, "name": "tags",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_config_map_filter",
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
									"view": "config_map",
									"config_values": []cu.IM{{
										"id":   1,
										"code": "123456",
										"data": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "config_map",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "value",
						"event": ct.FormEventChange, "name": "filter",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_config_map_field_name",
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
									"view": "config_map",
									"config_values": []cu.IM{{
										"id":   1,
										"code": "123456",
										"data": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "config_map",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "value",
						"event": ct.FormEventChange, "name": "field_name",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_auth_tags",
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
									"view": "auth",
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "auth",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "value",
						"event": ct.FormEventChange, "name": "tags",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_auth_filter",
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
									"view": "auth",
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "auth",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "value",
						"event": ct.FormEventChange, "name": "filter",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_auth_user_name",
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
									"view": "auth",
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "auth",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "value",
						"event": ct.FormEventChange, "name": "user_name",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_auth_user_group",
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
									"view": "auth",
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "auth",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "GROUP_ADMIN",
						"event": ct.FormEventChange, "name": "user_group",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_auth_description",
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
									"view": "auth",
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "auth",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "description",
						"event": ct.FormEventChange, "name": "description",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_event_change_auth_password",
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
									"view": "auth",
									"auth": []cu.IM{{
										"id":        1,
										"code":      "123456",
										"auth_meta": cu.IM{},
									}},
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
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data": cu.IM{
							"form": cu.IM{"index": 0, "key": "auth",
								"data": cu.IM{"data": cu.IM{"tags": []string{"tag1", "tag2"}}}}, "data": cu.IM{"code": "123456"}},
						"value": "password",
						"event": ct.FormEventChange, "name": "password",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "form_add_tag_meta_ok",
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
									"customer": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
								},
							},
						},
					},
					Name: ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "form_add_tag", "frm_key": "view", "frm_index": 0, "name": "tags",
						"row": cu.IM{"tags": []string{"tag1", "tag2"},
							"view_meta": cu.IM{"tags": []string{"tag1", "tag2"}}}},
						"value": cu.IM{
							"value": "tag3",
						}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_add_tag_ok",
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
									"customer": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
								},
							},
						},
					},
					Name: ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "form_add_tag", "frm_key": "view", "frm_index": 0, "name": "tags",
						"row": cu.IM{"tags": []string{"tag1", "tag2"}}},
						"value": cu.IM{
							"value": "tag3",
						}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_add_tag_cancel",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "form_add_tag", "name": "tags"}, "value": cu.IM{"value": ""}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_config_delete",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"config_values": []cu.IM{{
										"id":   1,
										"code": "123456",
										"data": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "config_delete", "code": "123456"}, "value": cu.IM{"value": "123456"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_auth_delete",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"auth": []cu.IM{{
										"id":   1,
										"code": "123456",
										"auth_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "auth_delete", "code": "123456"}, "value": cu.IM{"value": "123456"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_currency_delete",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"currency": []cu.IM{{
										"id":   1,
										"code": "USD",
										"currency_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "currency_delete", "code": "USD"}, "value": cu.IM{"value": "USD"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_currency_delete_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, errors.New("error")
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"currency": []cu.IM{{
										"id":   1,
										"code": "USD",
										"currency_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "currency_delete", "code": "USD"}, "value": cu.IM{"value": "USD"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_tax_delete",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tax": []cu.IM{{
										"id":   1,
										"code": "VAT",
										"tax_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "tax_delete", "code": "VAT"}, "value": cu.IM{"value": "VAT"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_tax_delete_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, errors.New("error")
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tax": []cu.IM{{
										"id":   1,
										"code": "VAT",
										"tax_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "tax_delete", "code": "VAT"}, "value": cu.IM{"value": "VAT"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_password_reset",
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
									"auth": []cu.IM{{
										"id":   1,
										"code": "123456",
										"auth_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "password_reset", "code": "123456"}, "value": cu.IM{"value": "123456"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_tax_add_exists",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tax": []cu.IM{{
										"id":   1,
										"code": "VAT",
										"tax_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "tax_add", "code": "VAT"}, "value": cu.IM{"value": "VAT"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_tax_add",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].Filters[1].Field == "id" {
									return []cu.IM{{"id": 1}}, nil
								}
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tax": []cu.IM{{
										"id":   1,
										"code": "VAT",
										"tax_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "tax_add", "code": "VAT"}, "value": cu.IM{"value": "VAT"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_tax_add_update_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, errors.New("error")
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].Filters[1].Field == "id" {
									return []cu.IM{{"id": 1}}, nil
								}
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tax": []cu.IM{{
										"id":   1,
										"code": "VAT",
										"tax_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "tax_add", "code": "VAT"}, "value": cu.IM{"value": "VAT"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_auth_add_exists",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"auth": []cu.IM{{
										"id":   1,
										"code": "123456",
										"auth_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "auth_add", "code": "123456"}, "value": cu.IM{"value": "123456"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_auth_add",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].Filters[1].Field == "id" {
									return []cu.IM{{"id": 1}}, nil
								}
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"auth": []cu.IM{{
										"id":   1,
										"code": "123456",
										"auth_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "auth_add", "code": "123456"}, "value": cu.IM{"value": "123456"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_auth_add_update_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, errors.New("error")
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].Filters[1].Field == "id" {
									return []cu.IM{{"id": 1}}, nil
								}
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"auth": []cu.IM{{
										"id":   1,
										"code": "123456",
										"auth_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "auth_add", "code": "123456"}, "value": cu.IM{"value": "123456"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_currency_add_exists",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"currency": []cu.IM{{
										"id":   1,
										"code": "USD",
										"currency_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "currency_add", "code": "USD"}, "value": cu.IM{"value": "USD"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_currency_add_invalid",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"currency": []cu.IM{{
										"id":   1,
										"code": "USD",
										"currency_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "currency_add", "code": "USD"}, "value": cu.IM{"value": "12345"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_currency_add",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].Filters[1].Field == "id" {
									return []cu.IM{{"id": 1}}, nil
								}
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"currency": []cu.IM{{
										"id":   1,
										"code": "USD",
										"currency_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "currency_add", "code": "USD"}, "value": cu.IM{"value": "USD"}},
				},
			},
			wantErr: false,
		},
		{
			name: "form_currency_add_update_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, errors.New("error")
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].Filters[1].Field == "id" {
									return []cu.IM{{"id": 1}}, nil
								}
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"currency": []cu.IM{{
										"id":   1,
										"code": "USD",
										"currency_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									}},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "currency_add", "code": "USD"}, "value": cu.IM{"value": "USD"}},
				},
			},
			wantErr: false,
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
