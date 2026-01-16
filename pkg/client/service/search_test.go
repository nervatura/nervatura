package service

import (
	"log/slog"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	cp "github.com/nervatura/nervatura/v6/pkg/client/component"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestSearchService_Data(t *testing.T) {
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
			name: "customer",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "name": "test"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: cp.NewClientComponent(map[string]cu.SM{}),
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
				"view": "customer",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
					Filter: "field like 'abc'",
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSearchService(tt.cls)
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

func TestSearchService_Response(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		evt     ct.ResponseEvent
		wantErr bool
	}{
		{
			name: "side_menu_group_open",
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
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			},
			evt: ct.ResponseEvent{
				Name: ct.ClientEventSideMenu,
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"search": cu.IM{
								"view":       "customer",
								"side_group": "group_customer",
							},
						},
					},
					Ticket: ct.Ticket{
						User: cu.IM{},
					},
				},
				Value: "group_product",
			},
			wantErr: false,
		},
		{
			name: "side_menu_group_close",
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
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			},
			evt: ct.ResponseEvent{
				Name: ct.ClientEventSideMenu,
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"search": cu.IM{
								"view":       "customer",
								"side_group": "group_value_1",
							},
						},
					},
					Ticket: ct.Ticket{
						User: cu.IM{},
					},
				},
				Value: "group_value_1",
			},
			wantErr: false,
		},
		{
			name: "side_menu_default",
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
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			},
			evt: ct.ResponseEvent{
				Name: ct.ClientEventSideMenu,
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
				Value: "value_2",
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
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"search": cu.IM{
								"view":     "customer",
								"customer": cu.IM{},
							},
						},
					},
					Ticket: ct.Ticket{
						User: cu.IM{
							"id": 1,
						},
					},
					Lang:            "en",
					CustomFunctions: &cp.ClientComponent{},
				},
				Name:  ct.FormEventOK,
				Value: cu.IM{"data": cu.IM{"next": "bookmark_add"}, "value": cu.IM{"value": "label"}},
			},
			wantErr: false,
		},
		{
			name: "bookmark_new",
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
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"search": cu.IM{
								"view":     "customer",
								"customer": cu.IM{},
							},
						},
					},
					Ticket: ct.Ticket{
						User: cu.IM{
							"id": 1,
						},
					},
					Lang:            "en",
					CustomFunctions: &cp.ClientComponent{},
				},
				Name:  ct.BrowserEventBookmark,
				Value: cu.IM{"data": cu.IM{"next": "bookmark_add"}, "value": cu.IM{"value": "label"}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSearchService(tt.cls)
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
