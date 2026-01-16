package service

import (
	"errors"
	"log/slog"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestCustomerService_Data(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		evt     ct.ResponseEvent
		params  cu.IM
		wantErr bool
	}{
		{
			name: "success",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
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
			params: cu.IM{
				"customer_id": 1,
			},
			wantErr: false,
		},
		{
			name: "customer_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, errors.New("error")
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
			params: cu.IM{
				"customer_id": 1,
			},
			wantErr: true,
		},
		{
			name: "config_map_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "config_map" {
									return []cu.IM{}, errors.New("error")
								}
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
			params: cu.IM{
				"customer_id": 1,
			},
			wantErr: true,
		},
		{
			name: "config_data_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "config_data" {
									return []cu.IM{}, errors.New("error")
								}
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
			params: cu.IM{
				"customer_id": 1,
			},
			wantErr: true,
		},
		{
			name: "config_report_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "config_report" {
									return []cu.IM{}, errors.New("error")
								}
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
			params: cu.IM{
				"customer_id": 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomerService(tt.cls)
			_, gotErr := s.Data(tt.evt, tt.params)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Data() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Data() succeeded unexpectedly")
			}
		})
	}
}

func TestCustomerService_Response(t *testing.T) {
	type fields struct {
		cls *ClientService
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToType: func(data interface{}, result any) (err error) {
								return nil
							},
						}
					},
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
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "editor_cancel"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_delete_error",
			fields: fields{
				cls: &ClientService{Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToType: func(data interface{}, result any) (err error) {
								return nil
							},
						}
					},
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
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
				},
			},
			wantErr: true,
		},
		{
			name: "editor_delete",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 1,
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_add_tag_ok",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{}},
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
										Connected: true,
										Engine:    "sqlite",
									}
								},
								"QuerySQL": func(sqlString string) ([]cu.IM, error) {
									return []cu.IM{{"id": 1, "name": "test"}}, nil
								},
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
										Connected: true,
										Engine:    "sqlite",
									}
								},
								"QuerySQL": func(sqlString string) ([]cu.IM, error) {
									return []cu.IM{{"id": 1, "name": "test"}}, nil
								},
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "editor_map_value"}, "value": cu.IM{"value": "code", "model": "customer", "map_field": "tags"}},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "invalid"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "client_form_delete_out",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
										"customer_meta": cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
										"customer_meta": cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
										"customer_meta": cu.IM{
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
			name: "client_form_change_add",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
			name: "client_form_change_default",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
		{
			name: "client_form_invalid",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
						}
					},
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
										"contacts": []cu.IM{},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
		{
			name: "side_editor_save_err",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									return 0, errors.New("error")
								},
							}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
							ConvertToType: ut.ConvertToType,
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									return 12345, nil
								},
							}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
							ConvertToType: ut.ConvertToType,
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345, "code": "123456"},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
									"dirty":    true,
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
									"config_report": []cu.IM{
										{"report_key": "report_key"},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventRowSelected, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_addresses",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
							ConvertToType: ut.ConvertToType,
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345,
										"addresses": []cu.IM{
											{"id": 12345},
										}},
									"view": "addresses",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_contacts",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
							ConvertToType: ut.ConvertToType,
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345},
									"view":     "contacts",
									"contacts": []cu.IM{
										{"id": 12345},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "contacts"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_events",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
							ConvertToType: ut.ConvertToType,
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{"id": 12345, "events": []cu.IM{
										{"id": 12345},
									}},
									"view": "events",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_map_field_customer",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer":  cu.IM{"id": 12345},
									"view":      "maps",
									"map_field": "ref_customer",
									"config_map": []cu.IM{
										{"field_name": "ref_customer", "field_type": "FIELD_CUSTOMER"},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_map_field_enum",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer":  cu.IM{"id": 12345},
									"view":      "maps",
									"map_field": "demo_string",
									"config_map": []cu.IM{
										{"field_name": "demo_string", "field_type": "FIELD_ENUM", "tags": []string{"tag1", "tag2"}},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_delete",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventFormDelete,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_validate_err",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventFormUpdate,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "customer_ref", "field_type": "FIELD_CUSTOMER", "value": "CUS12345"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: true,
		},
		{
			name: "table_form_update_ok",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventFormUpdate,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_change",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventFormChange,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_form_cancel",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventFormCancel,
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "map_field",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "map_field",
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "queue_err",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "queue",
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: true,
		},
		{
			name: "queue_ok",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &td.TestDriver{Config: cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"customer": cu.IM{
										"id": 12345,
										"customer_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view":      "maps",
									"map_field": "demo_string",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "queue",
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
				},
			},
			wantErr: false,
		},
		{
			name: "tag_delete",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "tags",
						"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "addresses"},
						"event": ct.ListEventDelete,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "tag_add",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "tags",
						"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "addresses"},
						"event": ct.ListEventAddItem,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "customer_name",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_name",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "customer_type",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "customer_type",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "notes",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
			name: "terms",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "terms",
						"value": 1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "credit_limit",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "credit_limit",
						"value": 1000,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "discount",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "discount",
						"value": 10,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "inactive",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
			name: "tax_free",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "tax_free",
						"value": true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "report_orientation",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db:     &td.TestDriver{Config: cu.IM{}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
						}
					},
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
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
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
			s := &CustomerService{
				cls: tt.fields.cls,
			}
			_, err := s.Response(tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerService.Response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCustomerService_update(t *testing.T) {
	type fields struct {
		cls *ClientService
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
			name: "new_customer",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
				},
			},
			args: args{
				ds: &api.DataStore{
					Db: &td.TestDriver{Config: cu.IM{
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
					"customer": cu.IM{
						"id": 0, "code": "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "update_customer",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
				},
			},
			args: args{
				ds: &api.DataStore{
					Db: &td.TestDriver{Config: cu.IM{
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
					"customer": cu.IM{
						"id": 1, "code": "test",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CustomerService{
				cls: tt.fields.cls,
			}
			_, err := s.update(tt.args.ds, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerService.customerUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
