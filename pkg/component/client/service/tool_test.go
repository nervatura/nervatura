package service

import (
	"errors"
	"log/slog"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestToolService_Data(t *testing.T) {
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
				"tool_id": 1,
			},
			wantErr: false,
		},
		{
			name: "tool_error",
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
				"tool_id": 1,
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
						Db: &md.TestDriver{Config: cu.IM{
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
				"tool_id": 1,
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
						Db: &md.TestDriver{Config: cu.IM{
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
				"tool_id": 1,
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
						Db: &md.TestDriver{Config: cu.IM{
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
				"tool_id": 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewToolService(tt.cls)
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

func TestToolService_Response(t *testing.T) {
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "editor_delete"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "next_product",
			cls: &ClientService{
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "product"}, "value": cu.IM{}},
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
						Database: "test",
						User:     cu.IM{},
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
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
						Database: "test",
						User:     cu.IM{},
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
						Database: "test",
						User:     cu.IM{},
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
						Database: "test",
						User:     cu.IM{"id": 1},
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
						Database: "test",
						User:     cu.IM{"id": 1},
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
						Database: "test",
						User:     cu.IM{},
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
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "editor_map_value"}, "value": cu.IM{"value": "code", "model": "tool", "map_field": "tags"}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "tags", "form_event": ct.ListEventAddItem},
			},
			wantErr: false,
		},
		{
			name: "client_form_change_delete_meta",
			cls: &ClientService{
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data":  cu.IM{"form": cu.IM{"index": 0, "key": "events", "data": cu.IM{"tags": []string{"tag1", "tag2"}}}, "data": cu.IM{"name": "contact2"}},
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "editor_save",
			},
			wantErr: false,
		},
		{
			name: "editor_delete",
			cls: &ClientService{
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
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool":  cu.IM{"id": 12345},
								"dirty": true,
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database:  "test",
						User:      cu.IM{},
						SessionID: "sess_123",
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool": cu.IM{"id": 12345},
								"config_report": []cu.IM{
									{"report_key": "report_key"},
								},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool": cu.IM{"id": 12345},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventRowSelected, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool": cu.IM{"id": 12345, "events": []cu.IM{
									{"id": 12345},
								}},
								"view": "events",
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_events_base",
			cls: &ClientService{
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_map_field_tool",
			cls: &ClientService{
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
								"tool":      cu.IM{"id": 12345},
								"view":      "maps",
								"map_field": "demo_string",
								"config_map": []cu.IM{
									{"field_name": "demo_string", "field_type": "FIELD_ENUM", "tags": []string{"tag1", "tag2"}},
								},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormDelete,
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string"}, "index": 0, "view": "events"}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormUpdate,
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "customer_ref", "field_type": "FIELD_CUSTOMER", "value": "CUS12345"}, "index": 0, "view": "events"}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormUpdate,
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormChange,
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormCancel,
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "map_field",
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
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
							"modal": cu.IM{
								"data": cu.IM{},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "queue",
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
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
							"modal": cu.IM{
								"data": cu.IM{},
							},
						},
					},
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "queue",
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "tags",
					"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "events"},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "tags",
					"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "events"},
					"event": ct.ListEventAddItem,
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
			name: "serial_number",
			cls: &ClientService{
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "serial_number",
					"value": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "description",
			cls: &ClientService{
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "description",
					"value": "value",
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
			name: "product_code_selected",
			cls: &ClientService{
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "product_code",
					"event": ct.SelectorEventSelected,
					"value": cu.IM{"row": cu.IM{"id": 12345, "meta": cu.IM{}}},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
						Db:     &md.TestDriver{Config: cu.IM{}},
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
					Ticket: ct.Ticket{
						Database: "test",
						User:     cu.IM{},
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
			s := NewToolService(tt.cls)
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

func TestToolService_update(t *testing.T) {
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
			name: "new_tool",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{}},
					}
				},
			},
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
			wantErr: false,
		},
		{
			name: "update_tool",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{}},
					}
				},
			},
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewToolService(tt.cls)
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

func TestToolService_ResponseNew(t *testing.T) {
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
							Database: "test",
							User:     cu.IM{},
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
				cls: &ClientService{
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
			},
			args: args{
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
				cls: &ClientService{
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
				cls: &ClientService{
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
							Database: "test",
							User:     cu.IM{},
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
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
							Database: "test",
							User:     cu.IM{},
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
							Database: "test",
							User:     cu.IM{},
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
							Database: "test",
							User:     cu.IM{"id": 1},
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
							Database: "test",
							User:     cu.IM{"id": 1},
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
							Database: "test",
							User:     cu.IM{},
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
							Database: "test",
							User:     cu.IM{},
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
				cls: &ClientService{
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
			},
			args: args{
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
			name: "side_editor_save_err",
			fields: fields{
				cls: &ClientService{
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "editor_save",
				},
			},
			wantErr: false,
		},
		{
			name: "editor_delete_side",
			fields: fields{
				cls: &ClientService{
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"tool": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"tool":  cu.IM{"id": 12345},
									"dirty": true,
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"tool": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							Database:  "test",
							User:      cu.IM{},
							SessionID: "sess_123",
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"tool": cu.IM{"id": 12345},
									"config_report": []cu.IM{
										{"report_key": "report_key"},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"tool": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"tool": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"tool": cu.IM{"id": 12345},
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
				cls: &ClientService{
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
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
						},
					},
					Name:  ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "events"}},
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
								"modal": cu.IM{
									"data": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
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
				cls: &ClientService{
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
								"modal": cu.IM{
									"data": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							Database: "test",
							User:     cu.IM{},
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "queue",
						"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "demo_string", "field_type": "FIELD_STRING", "value": "value"}, "index": 0, "view": "events"}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToolService{
				cls: tt.fields.cls,
			}
			_, err := s.Response(tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToolService.Response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
