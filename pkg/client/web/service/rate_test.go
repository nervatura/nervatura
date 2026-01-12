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
)

func TestRateService_Data(t *testing.T) {
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
				"rate_id": 1,
			},
			wantErr: false,
		},
		{
			name: "rate_error",
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
				"rate_id": 1,
			},
			wantErr: true,
		},
		{
			name: "currency_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "currency" {
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
				"rate_id": 1,
			},
			wantErr: true,
		},
		{
			name: "place_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "place" {
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
				"rate_id": 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewRateService(tt.cls)
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

func TestRateService_Response(t *testing.T) {
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
								"rate": cu.IM{
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
								"rate": cu.IM{
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
								"rate": cu.IM{
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
								"rate": cu.IM{
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
								"rate": cu.IM{
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
								"rate": cu.IM{
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
						ConvertToType: ut.ConvertToType,
					}
				},
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"rate": cu.IM{"id": 12345},
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
						ConvertToType: ut.ConvertToType,
					}
				},
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"rate": cu.IM{"id": 12345, "code": "RAT1731101982N123", "place_code": "PLA1731101982N123", "currency_code": "EUR"},
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
			name: "side_editor_save_new",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 12345}}, nil
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
						ConvertToType: ut.ConvertToType,
					}
				},
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"rate": cu.IM{"place_code": "PLA1731101982N123", "currency_code": "EUR"},
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
								"rate": cu.IM{"id": 12345},
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
								"rate": cu.IM{"id": 12345},
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
								"rate":  cu.IM{"id": 12345},
								"dirty": true,
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
								"rate": cu.IM{"id": 12345},
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
								"rate": cu.IM{"id": 12345},
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
								"rate": cu.IM{"id": 12345},
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"rate_map": cu.IM{
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
					"value": cu.IM{"row": cu.IM{"id": 12345, "tag": "tag1"}, "index": 0, "view": "rate"},
					"event": ct.ListEventAddItem,
				},
			},
			wantErr: false,
		},
		{
			name: "rate_value",
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"view": "rate",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "rate_value",
					"value": 1234,
				},
			},
			wantErr: false,
		},
		{
			name: "rate_date",
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"view": "rate",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "rate_date",
					"value": "2021-01-01",
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"view": "rate",
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
			name: "rate_type",
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"view": "rate",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "rate_type",
					"value": "RATE_BUY",
				},
			},
			wantErr: false,
		},
		{
			name: "place_code",
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"view": "rate",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "place_code",
					"value": "PLA1731101982N123",
				},
			},
			wantErr: false,
		},
		{
			name: "currency_code",
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
								},
								"view": "rate",
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "currency_code",
					"value": "EUR",
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
								"rate": cu.IM{
									"id": 12345,
									"rate_meta": cu.IM{
										"tags": []string{"tag1", "tag2"},
									},
									"rate_map": cu.IM{
										"demo_string": "tag1",
									}},
								"view": "rate",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewRateService(tt.cls)
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
