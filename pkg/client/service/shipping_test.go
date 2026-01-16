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
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestShippingService_Data(t *testing.T) {
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
								return []cu.IM{{"id": 1, "direction": md.DirectionOut.String()}}, nil
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
				"trans_id": 1,
			},
			wantErr: false,
		},
		{
			name: "trans_error",
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
				"trans_id": 1,
			},
			wantErr: true,
		},
		{
			name: "items_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "item_shipping" {
									return []cu.IM{}, errors.New("error")
								}
								return []cu.IM{{"id": 1, "direction": md.DirectionOut.String()}}, nil
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
				"trans_id": 1,
			},
			wantErr: true,
		},
		{
			name: "movements_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if strings.Contains(queries[0].From, "movement_view") {
									return []cu.IM{}, errors.New("error")
								}
								return []cu.IM{{"id": 1, "direction": md.DirectionOut.String()}}, nil
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
				"trans_id": 1,
			},
			wantErr: true,
		},
		{
			name: "places_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "place" {
									return []cu.IM{}, errors.New("error")
								}
								return []cu.IM{{"id": 1, "direction": md.DirectionOut.String()}}, nil
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
				"trans_id": 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShippingService(tt.cls)
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

func TestShippingService_Response(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		evt     ct.ResponseEvent
		wantErr bool
	}{
		{
			name: "next_editor_cancel",
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
			name: "next_create_all",
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
								return []cu.IM{{"id": 1, "code": "1234567890"}}, nil
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
								"items": []cu.IM{
									{
										"code":       "1234567890",
										"difference": 10,
									},
								},
								"shipping": cu.IM{
									"code":          "1234567890",
									"direction":     md.DirectionOut.String(),
									"place_code":    "1234567890",
									"shipping_time": "2021-01-01",
								},
							},
						},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "create_all"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "next_create_item",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, errors.New("error")
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234567890"}}, nil
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
								"items": []cu.IM{
									{
										"code": "1234567890",
										"qty":  10,
									},
								},
								"shipping": cu.IM{
									"code":          "1234567890",
									"direction":     md.DirectionOut.String(),
									"place_code":    "1234567890",
									"shipping_time": "2021-01-01",
								},
							},
						},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "create_item"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "next_create_item_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								if data.Model == "movement" {
									return 0, errors.New("error")
								}
								return 1, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234567890"}}, nil
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
								"items": []cu.IM{
									{
										"code": "1234567890",
										"qty":  10,
									},
								},
								"shipping": cu.IM{
									"code":          "1234567890",
									"direction":     md.DirectionOut.String(),
									"place_code":    "1234567890",
									"shipping_time": "2021-01-01",
								},
							},
						},
					},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "create_item"}, "value": cu.IM{}},
			},
			wantErr: false,
		},
		{
			name: "next_invalid",
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
				Value: cu.IM{"data": cu.IM{"next": "invalid_next"}, "value": cu.IM{}},
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
					}
				},
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{},
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
					}
				},
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
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
			name: "side_create_all",
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
				Name:  ct.ClientEventSideMenu,
				Value: "create_all",
			},
			wantErr: false,
		},
		{
			name: "side_create_item",
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
				Name:  ct.ClientEventSideMenu,
				Value: "create_item",
			},
			wantErr: false,
		},
		{
			name: "side_invalid",
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
				Name:  ct.ClientEventSideMenu,
				Value: "invalid_side",
			},
			wantErr: false,
		},
		{
			name: "field_place_code",
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
								"shipping": cu.IM{
									"place_code": "1234567890",
								},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "place_code",
					"value": "1234567890",
				},
			},
			wantErr: false,
		},
		{
			name: "field_shipping_time",
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
								"shipping": cu.IM{
									"place_code": "1234567890",
								},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "shipping_time",
					"value": "2021-01-01",
				},
			},
			wantErr: false,
		},
		{
			name: "field_qty_trigger",
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
								"shipping": cu.IM{
									"place_code": "1234567890",
								},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormChange,
					"value": cu.IM{"row": cu.IM{"qty": 10}, "index": 0, "value": 10},
					"trigger": &ct.Table{
						Rows: []cu.IM{{"qty": 0, "difference": 5}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "field_qty_trigger2",
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
								"shipping": cu.IM{
									"place_code": "1234567890",
								},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormChange,
					"value": cu.IM{"row": cu.IM{"qty": 10}, "index": 0, "value": -10},
					"trigger": &ct.Table{
						Rows: []cu.IM{{"qty": 0, "difference": 5}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "field_update_row",
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
								"shipping": cu.IM{
									"place_code": "1234567890",
								},
								"items": []cu.IM{
									{
										"code": "1234567890",
										"qty":  10,
									},
								},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": ct.TableEventFormUpdate,
					"value": cu.IM{"row": cu.IM{"qty": 10, "code": "1234567890"}, "index": 0},
				},
			},
			wantErr: false,
		},
		{
			name: "field_invalid",
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
								"shipping": cu.IM{
									"place_code": "1234567890",
								},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "invalid_field",
					"value": "1234567890",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShippingService(tt.cls)
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
