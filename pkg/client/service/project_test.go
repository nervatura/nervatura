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

func TestProjectService_Data(t *testing.T) {
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
						Database: "test",
						User:     cu.IM{},
					},
				},
			},
			params: cu.IM{
				"project_id": 1,
			},
			wantErr: false,
		},
		{
			name: "project_error",
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
						Database: "test",
						User:     cu.IM{},
					},
				},
			},
			params: cu.IM{
				"project_id": 1,
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
						Database: "test",
						User:     cu.IM{},
					},
				},
			},
			params: cu.IM{
				"project_id": 1,
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
						Database: "test",
						User:     cu.IM{},
					},
				},
			},
			params: cu.IM{
				"project_id": 1,
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
						Database: "test",
						User:     cu.IM{},
					},
				},
			},
			params: cu.IM{
				"project_id": 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProjectService(tt.cls)
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

func TestProjectService_Response(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		evt     ct.ResponseEvent
		wantErr bool
	}{
		{
			name: "editor_cancel",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "editor_cancel"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "editor_delete_error",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
			},
			wantErr: true,
		},
		{
			name: "editor_delete",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 1,
								},
							},
						},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "next_customer",
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
					}
				},
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 1,
								},
							},
						},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "customer"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "editor_add_tag_ok",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "editor_add_tag_cancel",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "form_add_tag_meta_ok",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "form_add_tag_ok",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "form_add_tag_cancel",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "bookmark_add",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "bookmark_add_error",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "editor_map_value_invalid",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "editor_map_value_update",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
				Value: cu.IM{"data": cu.IM{"next": "editor_map_value"}, "value": cu.IM{"value": "code", "model": "project", "map_field": "tags"}},
			},
			wantErr: false,
		},
		{
			name: "invalid",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "invalid"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "client_form_delete_out",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 1,
									"project_meta": cu.IM{
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
			wantErr: false,
		},
		{
			name: "client_form_delete_in",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 1,
									"project_meta": cu.IM{
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
			wantErr: false,
		},
		{
			name: "client_form_cancel",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 1,
									"project_meta": cu.IM{
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
			wantErr: false,
		},
		{
			name: "client_form_ok",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
					"value": cu.IM{"name": "contact2"},
					"event": ct.FormEventOK},
			},
			wantErr: false,
		},
		{
			name: "client_form_change_add",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "client_form_change_default",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "client_form_change_skip",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
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
			wantErr: false,
		},
		{
			name: "client_form_invalid",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 1,
									"meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"contacts": []cu.IM{},
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
			wantErr: false,
		},
		{
			name: "side_editor_save_err",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_save",
			},
			wantErr: true,
		},
		{
			name: "side_editor_save_ok",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345, "code": "123456", "customer_code": "CUS1731101982N123"},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_save",
			},
			wantErr: false,
		},
		{
			name: "editor_delete_side_menu",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_delete",
			},
			wantErr: false,
		},
		{
			name: "side_editor_cancel",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_cancel",
			},
			wantErr: false,
		},
		{
			name: "side_editor_cancel_dirty",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
								"dirty":   true,
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_cancel",
			},
			wantErr: false,
		},
		{
			name: "editor_new",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_new",
			},
			wantErr: false,
		},
		{
			name: "editor_report",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
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
			wantErr: false,
		},
		{
			name: "editor_bookmark",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_bookmark",
			},
			wantErr: false,
		},
		{
			name: "invalid_menu",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "invalid_menu",
			},
			wantErr: false,
		},
		{
			name: "table_row_selected",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventRowSelected, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_addresses",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345,
									"addresses": []cu.IM{
										{"id": 12345},
									}},
								"view": "addresses",
							},
						},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_contacts",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345},
								"view":    "contacts",
								"contacts": []cu.IM{
									{"id": 12345},
								},
							},
						},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "contacts"}},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_events",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{"id": 12345, "events": []cu.IM{
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
			wantErr: false,
		},
		{
			name: "table_add_item_map_field_project",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project":   cu.IM{"id": 12345},
								"view":      "maps",
								"map_field": "ref_customer",
								"config_map": []cu.IM{
									{"field_name": "ref_customer", "field_type": "FIELD_CUSTOMER"},
								},
							},
						},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_map_field_enum",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project":   cu.IM{"id": 12345},
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
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "table_form_delete",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string"}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "table_form_update_validate_err",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "customer_ref", "field_type": "FIELD_CUSTOMER", "value": "CUS12345"}, "index": 0, "view": "addresses"}},
			},
			wantErr: true,
		},
		{
			name: "table_form_update_ok",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "table_form_change",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "table_form_cancel",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "map_field",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "queue_err",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
			},
			wantErr: true,
		},
		{
			name: "queue_ok",
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"project": cu.IM{
									"id": 12345,
									"project_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "addresses"}},
			},
			wantErr: false,
		},
		{
			name: "tag_delete",
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
			wantErr: false,
		},
		{
			name: "tag_add",
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
			wantErr: false,
		},
		{
			name: "project_name",
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
				Value: cu.IM{"name": "project_name",
					"value": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "notes",
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
				Value: cu.IM{"name": "notes",
					"value": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "start_date",
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
				Value: cu.IM{"name": "start_date",
					"value": "2025-01-01",
				},
			},
			wantErr: false,
		},
		{
			name: "end_date",
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
				Value: cu.IM{"name": "end_date",
					"value": "2025-01-01",
				},
			},
			wantErr: false,
		},
		{
			name: "customer_code_delete",
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
			wantErr: false,
		},
		{
			name: "inactive",
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
				Value: cu.IM{"name": "inactive",
					"value": true,
				},
			},
			wantErr: false,
		},
		{
			name: "report_orientation",
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
				Value: cu.IM{"name": "orientation",
					"value": "portrait",
				},
			},
			wantErr: false,
		},
		{
			name: "skip",
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
				Value: cu.IM{"name": "invalid",
					"value": "value",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProjectService(tt.cls)
			_, gotErr := s.Response(tt.evt)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Response() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Response() succeeded unexpectedly")
			}
		})
	}
}

func TestProjectService_update(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		ds      *api.DataStore
		data    cu.IM
		wantErr bool
	}{
		{
			name: "new_project",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{}},
					}
				},
			},
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
				"project": cu.IM{
					"id": 0, "code": "test", "customer_code": "test",
				},
			},
			wantErr: false,
		},
		{
			name: "update_project",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{}},
					}
				},
			},
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
				"project": cu.IM{
					"id": 1, "code": "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProjectService(tt.cls)
			_, gotErr := s.update(tt.ds, tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("update() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("update() succeeded unexpectedly")
			}
		})
	}
}
