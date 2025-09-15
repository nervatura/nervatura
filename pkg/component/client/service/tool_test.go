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

func TestClientService_toolData(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		ds     *api.DataStore
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
				params: cu.IM{
					"tool_id": 1,
				},
			},
			wantErr: false,
		},
		{
			name: "tool_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, errors.New("error")
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToType: func(data interface{}, result any) (err error) {
						return nil
					},
				},
				user: cu.IM{},
				params: cu.IM{
					"tool_id": 1,
				},
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
				params: cu.IM{
					"tool_id": 1,
				},
			},
			wantErr: true,
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
				params: cu.IM{
					"tool_id": 1,
				},
			},
			wantErr: true,
		},
		{
			name: "config_report_error",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "config_report" {
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
				params: cu.IM{
					"tool_id": 1,
				},
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
			_, err := cls.toolData(tt.args.ds, tt.args.user, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.toolData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_toolUpdate(t *testing.T) {
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
			name: "new_tool",
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
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToType: func(data interface{}, result any) (err error) {
						return nil
					},
				},
				data: cu.IM{
					"tool": cu.IM{
						"id": 0, "code": "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "update_tool",
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
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToType: func(data interface{}, result any) (err error) {
						return nil
					},
				},
				data: cu.IM{
					"tool": cu.IM{
						"id": 1, "code": "test",
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
			_, err := cls.toolUpdate(tt.args.ds, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.toolUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_toolResponse(t *testing.T) {
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
			name: "editor_cancel",
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
							Data: cu.IM{},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "editor_cancel"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_delete_error",
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
							Data: cu.IM{},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
				},
			},
			wantErr: true,
		},
		{
			name: "editor_delete",
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
									"tool": cu.IM{
										"id": 1,
									},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_product",
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
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
										"id": 1,
									},
								},
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "product"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_add_tag_ok",
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
									"tool": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "editor_add_tag"}, "value": cu.IM{"value": "value"}},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_add_tag_cancel",
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
									"tool": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "editor_add_tag"}, "value": cu.IM{"value": ""}},
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
									"tool": cu.IM{
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
									"tool": cu.IM{
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
									"tool": cu.IM{
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
			name: "bookmark_add",
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
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"id": 1,
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "bookmark_add"}, "value": cu.IM{"value": "label"}},
				},
			},
			wantErr: false,
		},
		{
			name: "bookmark_add_error",
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
								return []cu.IM{{"id": 1}}, errors.New("error")
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
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"id": 1,
							},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "bookmark_add"}, "value": cu.IM{"value": "label"}},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_map_value_invalid",
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
									"tool": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "editor_map_value"}, "value": cu.IM{"value": ""}},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_map_value_update",
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
									"tool": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "editor_map_value"}, "value": cu.IM{"value": "code", "model": "tool", "map_field": "tags"}},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid",
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
							Data: cu.IM{},
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "invalid"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_delete_out",
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
									"tool": cu.IM{
										"id": 1,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
									"view": []cu.IM{
										{
											"id":   1,
											"name": "contact1",
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "view"}, "data": cu.IM{}},
						"value": cu.IM{"form_delete": "form_delete"},
						"event": ct.FormEventCancel},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_delete_in",
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
									"tool": cu.IM{
										"id": 1,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"view": []cu.IM{
											{
												"id":   1,
												"name": "contact1",
											},
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "view"}, "data": cu.IM{}},
						"value": cu.IM{"form_delete": "form_delete"},
						"event": ct.FormEventCancel},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_cancel",
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
									"tool": cu.IM{
										"id": 1,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
									},
									"view": []cu.IM{
										{
											"id":   1,
											"name": "contact1",
										},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "view"}, "data": cu.IM{}},
						"value": cu.IM{"name": ""},
						"event": ct.FormEventCancel},
				},
			},
			wantErr: false,
		},
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
									"tool": cu.IM{
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
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{"name": "contact2"},
						"event": ct.FormEventOK},
				},
			},
			wantErr: false,
		},
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
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
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
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{},
						"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventAddItem},
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
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
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
						"value": "value",
						"event": ct.FormEventChange, "name": "default"},
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
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"tool": cu.IM{
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
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{},
						"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventEditItem},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_invalid",
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
									"tool": cu.IM{
										"id": 1,
										"meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"events": []cu.IM{},
									},
								},
							},
						},
					},
					Name: ct.ClientEventForm,
					Value: cu.IM{
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{},
						"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventEditItem},
				},
			},
			wantErr: false,
		},
		{
			name: "side_editor_save_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 0, errors.New("error")
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
									"tool": cu.IM{"id": 12345},
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_save",
				},
			},
			wantErr: true,
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
								"editor": cu.IM{
									"tool": cu.IM{"id": 12345},
								},
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
			name: "editor_delete",
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
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_delete",
				},
			},
			wantErr: false,
		},
		{
			name: "side_editor_cancel",
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
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_cancel",
				},
			},
			wantErr: false,
		},
		{
			name: "side_editor_cancel_dirty",
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
									"tool":  cu.IM{"id": 12345},
									"dirty": true,
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_cancel",
				},
			},
			wantErr: false,
		},
		{
			name: "editor_new",
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
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_new",
				},
			},
			wantErr: false,
		},
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
			},
			wantErr: false,
		},
		{
			name: "editor_bookmark",
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
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_bookmark",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid_menu",
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
								},
							},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "invalid_menu",
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
								},
							},
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventRowSelected, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_events",
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
									"tool": cu.IM{"id": 12345, "events": []cu.IM{
										{"id": 12345},
									}},
									"view": "events",
								},
							},
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_events_base",
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
									"events": []cu.IM{
										{"id": 12345},
									},
									"view": "events",
								},
							},
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
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
			},
			wantErr: false,
		},
		{
			name: "table_form_delete",
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
					Value: cu.IM{"name": ct.TableEventFormDelete,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string"}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
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
			},
			wantErr: false,
		},
		{
			name: "table_form_change",
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
					Value: cu.IM{"name": ct.TableEventFormChange,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_cancel",
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
					Value: cu.IM{"name": ct.TableEventFormCancel,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
		{
			name: "map_field",
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
					Value: cu.IM{"name": "map_field",
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
		{
			name: "queue_err",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 0, errors.New("error")
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
					Value: cu.IM{"name": "queue",
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
				},
			},
			wantErr: true,
		},
		{
			name: "queue_ok",
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
					Value: cu.IM{"name": "queue",
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
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
					Value: cu.IM{"name": "tags",
						"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "events"},
						"event": ct.ListEventDelete,
					},
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
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
					Value: cu.IM{"name": "tags",
						"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "events"},
						"event": ct.ListEventAddItem,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "notes",
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "tool",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "notes",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "serial_number",
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "tool",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "serial_number",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "description",
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "tool",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "description",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "inactive",
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "tool",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "inactive",
						"value": true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "product_code_selected",
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "tool",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "product_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "report_orientation",
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "tool",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "orientation",
						"value": "portrait",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skip",
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
									"tool": cu.IM{
										"id": 12345,
										"tool_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"tool_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "tool",
								},
							},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "invalid",
						"value": "value",
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
			_, err := cls.toolResponse(tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.toolResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
