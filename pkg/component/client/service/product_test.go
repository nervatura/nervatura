package service

import (
	"errors"
	"log/slog"
	"strings"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestProductService_Data(t *testing.T) {
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
				"product_id": 1,
			},
			wantErr: false,
		},
		{
			name: "new",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{
									{"id": 1, "config_key": "default_taxcode"},
									{"id": 2, "config_key": "default_unit"},
								}, nil
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
			params:  cu.IM{},
			wantErr: false,
		},
		{
			name: "product_error",
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
				"product_id": 1,
			},
			wantErr: true,
		},
		{
			name: "price_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "price" {
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
				"product_id": 1,
			},
			wantErr: true,
		},
		{
			name: "components_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if strings.Contains(queries[0].From, "link") {
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
				"product_id": 1,
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
				"product_id": 1,
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
				"product_id": 1,
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
				"product_id": 1,
			},
			wantErr: true,
		},
		{
			name: "tax_view_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "tax_view" {
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
				"product_id": 1,
			},
			wantErr: true,
		},
		{
			name: "currency_view_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "currency_view" {
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
				"product_id": 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProductService(tt.cls)
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

func TestProductService_Response(t *testing.T) {
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
								"product": cu.IM{
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
			name: "customer",
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
								"product": cu.IM{
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
			name: "product_next",
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
								"product": cu.IM{
									"id": 1,
								},
							},
						},
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
				Value: cu.IM{"data": cu.IM{"next": "editor_map_value"}, "value": cu.IM{"value": "code", "model": "product", "map_field": "tags"}},
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
								"product": cu.IM{
									"id": 1,
									"product_meta": cu.IM{
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
								"product": cu.IM{
									"id": 1,
									"product_meta": cu.IM{
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
								"product": cu.IM{
									"id": 1,
									"product_meta": cu.IM{
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
							"link_code_2":   "P0000002",
							"product_name":  "Product 2",
							"unit":          "PC",
						}},
					},
					"value": cu.IM{"price_value": 1000, "qty": 100, "name": "event1",
						"component_qty": 100, "component_notes": "notes", "link_code_2": "P0000002", "product_name": "Product 2", "unit": "PC"},
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
								"product": cu.IM{
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
								"product": cu.IM{
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
								"product": cu.IM{
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
			name: "client_form_price_value",
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
								"product": cu.IM{
									"id": 1,
									"meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"prices": []cu.IM{
										{"id": 1, "price_value": 100},
									},
								},
							},
						},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data":  cu.IM{"form": cu.IM{"index": 0, "key": "prices", "data": cu.IM{"price_value": 100}}, "data": cu.IM{"name": "price2"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "price_value", "form_event": ct.ListEventEditItem},
			},
			wantErr: false,
		},
		{
			name: "client_form_component_qty",
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
								"product": cu.IM{
									"id": 1,
									"meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"components": []cu.IM{
										{"id": 1, "qty": 100},
									},
								},
							},
						},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data":  cu.IM{"form": cu.IM{"index": 0, "key": "components", "data": cu.IM{"qty": 100}}, "data": cu.IM{"name": "component2"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "component_qty", "form_event": ct.ListEventEditItem},
			},
			wantErr: false,
		},
		{
			name: "client_form_customer_code_selected",
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
								"product": cu.IM{
									"id": 1,
									"meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"prices": []cu.IM{
									{"id": 1, "price_value": 100, "customer_code": "12345"},
								},
								"form": cu.IM{
									"index": 0,
									"key":   "prices",
									"data": cu.IM{
										"id": 1, "price_value": 100, "customer_code": "12345",
									},
								},
							},
						},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "prices",
						"data": cu.IM{"customer_code": "12345"}}, "data": cu.IM{"name": "price2"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "customer_code", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "client_form_product_code_selected",
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
								"product": cu.IM{
									"id": 1,
									"meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"prices": []cu.IM{
									{"id": 1, "price_value": 100, "customer_code": "12345"},
								},
								"components": []cu.IM{
									{"id": 1, "link_code_2": "P0000002", "product_name": "Product 2", "unit": "PC"},
								},
								"form": cu.IM{
									"index": 0,
									"key":   "components",
									"data": cu.IM{
										"id": 1, "product_code": "P0000002", "product_name": "Product 2", "unit": "PC",
									},
								},
							},
						},
					},
				},
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "prices",
						"data": cu.IM{"product_code": "P0000002", "product_name": "Product 2", "unit": "PC"}}, "data": cu.IM{"name": "price2"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "product_code", "form_event": ct.SelectorEventSelected},
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
								"product": cu.IM{
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345},
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
								"product": cu.IM{"id": 12345,
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
			name: "table_add_item_prices",
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
								"product": cu.IM{"id": 12345},
								"view":    "prices",
								"prices": []cu.IM{
									{"id": 12345},
								},
								"currencies": []cu.IM{
									{"code": "USD"},
								},
							},
						},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "prices"}},
			},
			wantErr: false,
		},
		{
			name: "table_add_item_components",
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
								"product": cu.IM{"id": 12345},
								"view":    "components",
								"components": []cu.IM{
									{"id": 12345},
								},
								"currencies": []cu.IM{
									{"code": "USD"},
								},
							},
						},
					},
				},
				Name:  ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventAddItem, "value": cu.IM{"row": cu.IM{"id": 12345}, "index": 0, "view": "components"}},
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
								"product": cu.IM{"id": 12345, "events": []cu.IM{
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
			name: "table_add_item_map_field_customer",
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
								"product":   cu.IM{"id": 12345},
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
								"product":   cu.IM{"id": 12345},
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
			name: "table_edit_cell",
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "components",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventEditCell,
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "product_code", "value": "value"}, "index": 0, "view": "components"}},
			},
			wantErr: false,
		},
		{
			name: "table_edit_cell_dirty",
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view":  "components",
								"dirty": true,
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventEditCell,
					"value": cu.IM{"row": cu.IM{"id": 12345, "field_name": "product_code", "value": "value"}, "index": 0, "view": "components"}},
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
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
			name: "barcode_qty",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "barcode_qty",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "inactive",
					"value": "true",
				},
			},
			wantErr: false,
		},
		{
			name: "code",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "code",
					"value": "code",
				},
			},
			wantErr: false,
		},
		{
			name: "product_name",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "product_name",
					"value": "product_name",
				},
			},
			wantErr: false,
		},
		{
			name: "product_type",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "product_type",
					"value": "product_type",
				},
			},
			wantErr: false,
		},
		{
			name: "tax_code",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "tax_code",
					"value": "tax_code",
				},
			},
			wantErr: false,
		},
		{
			name: "unit",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "unit",
					"value": "unit",
				},
			},
			wantErr: false,
		},
		{
			name: "barcode_type",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "barcode_type",
					"value": "barcode_type",
				},
			},
			wantErr: false,
		},
		{
			name: "barcode",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "barcode",
					"value": "barcode",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "notes",
					"value": "notes",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
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
								"product": cu.IM{
									"id": 12345,
									"product_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"product_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "product",
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
			s := NewProductService(tt.cls)
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

func TestProductService_update(t *testing.T) {
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
			name: "new_product",
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
				"product": cu.IM{
					"id": 0, "code": "test",
				},
				"prices": []cu.IM{
					{
						"id": 0, "price_type": "PRICE_CUSTOMER",
						"valid_from": "2025-01-01", "valid_to": "2025-01-01",
						"customer_code": "CUS0000001",
						"price_meta":    cu.IM{"tags": []string{"tag1", "tag2"}},
					},
					{
						"id": 1, "price_type": "PRICE_CUSTOMER",
						"valid_from": "2025-01-01", "valid_to": "2025-01-01",
						"customer_code": "CUS0000001",
						"price_meta":    cu.IM{"tags": []string{"tag1", "tag2"}},
					},
				},
				"prices_delete": []cu.IM{
					{
						"id": 1,
					},
				},
				"components": []cu.IM{
					{
						"id": 12345, "link_type_1": "LINK_PRODUCT",
						"link_code_1": "P0000001",
						"link_type_2": "LINK_PRODUCT",
						"link_code_2": "P0000002",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "update_product",
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
				"product": cu.IM{
					"id": 1, "code": "test",
				},
			},
			wantErr: false,
		},
		{
			name: "price_error",
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
						if data.Model == "price" {
							return 0, errors.New("error")
						}
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
				"product": cu.IM{
					"id": 0, "code": "test",
				},
				"prices": []cu.IM{
					{
						"id": 0, "price_type": "PRICE_CUSTOMER",
						"valid_from": "2025-01-01", "valid_to": "2025-01-01",
						"customer_code": "CUS0000001",
						"price_meta":    cu.IM{"tags": []string{"tag1", "tag2"}},
					},
					{
						"id": 1, "price_type": "PRICE_CUSTOMER",
						"valid_from": "2025-01-01", "valid_to": "2025-01-01",
						"customer_code": "CUS0000001",
						"price_meta":    cu.IM{"tags": []string{"tag1", "tag2"}},
					},
				},
				"prices_delete": []cu.IM{
					{
						"id": 1,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "components_error",
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
						if data.Model == "link" {
							return 0, errors.New("error")
						}
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
				"product": cu.IM{
					"id": 0, "code": "test",
				},
				"components": []cu.IM{
					{
						"id": 12345, "link_type_1": "LINK_PRODUCT",
						"link_code_1": "P0000001",
						"link_type_2": "LINK_PRODUCT",
						"link_code_2": "P0000002",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "price_delete_error",
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
						if queries[0].From == "price" {
							return []cu.IM{}, errors.New("error")
						}
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
				"product": cu.IM{
					"id": 0, "code": "test",
				},
				"prices": []cu.IM{
					{
						"id": 0, "price_type": "PRICE_CUSTOMER",
						"valid_from": "2025-01-01", "valid_to": "2025-01-01",
						"customer_code": "CUS0000001",
						"price_meta":    cu.IM{"tags": []string{"tag1", "tag2"}},
					},
					{
						"id": 1, "price_type": "PRICE_CUSTOMER",
						"valid_from": "2025-01-01", "valid_to": "2025-01-01",
						"customer_code": "CUS0000001",
						"price_meta":    cu.IM{"tags": []string{"tag1", "tag2"}},
					},
				},
				"prices_delete": []cu.IM{
					{
						"id": 1,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProductService(tt.cls)
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
