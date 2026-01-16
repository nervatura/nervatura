package service

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestShortcutService_Data(t *testing.T) {
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
			params:  cu.IM{},
			wantErr: false,
		},
		{
			name: "error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return nil, errors.New("error")
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
			wantErr: true,
		},
		{
			name: "report_error",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &td.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								if queries[0].From == "config_report" {
									return nil, errors.New("error")
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
			params:  cu.IM{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortcutService{cls: tt.cls}
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

func TestShortcutService_Response(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		evt     ct.ResponseEvent
		wantErr bool
	}{
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
						Data: cu.IM{
							"editor": cu.IM{
								"shortcut": cu.IM{"id": 12345},
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
			name: "side_shortcut_call_external",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &td.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{"result": "test"}`), nil
						},
						ConvertFromByte: func(data []byte, result interface{}) error {
							return cu.ConvertFromByte(data, result)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{})
						},
						NewRequest: func(method, url string, body io.Reader) (*http.Request, error) {
							return httptest.NewRequest("POST", "/", nil), nil
						},
						RequestDo: func(req *http.Request) (*http.Response, error) {
							json := `{"code": 0, "message": "success"}`
							recorder := httptest.NewRecorder()
							recorder.Header().Add("Content-Type", "application/json")
							recorder.WriteString(json)
							expectedResponse := recorder.Result()
							return expectedResponse, nil
						},
					}
				},
			},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"shortcut": cu.IM{"address": "https://www.google.com", "method": md.ShortcutMethodGET.String()},
								"params":   cu.IM{"q": "test"},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "shortcut_call",
			},
			wantErr: false,
		},
		{
			name: "side_shortcut_call_own",
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
								"shortcut": cu.IM{"func_name": "test"},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "shortcut_call",
			},
			wantErr: false,
		},
		{
			name: "side_shortcut_reset",
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
								"shortcut": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "shortcut_reset",
			},
			wantErr: false,
		},
		{
			name: "side_shortcut_list",
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
								"shortcut": cu.IM{"id": 12345},
							},
						},
					},
				},
				Name:  ct.ClientEventSideMenu,
				Value: "shortcut_list",
			},
			wantErr: false,
		},
		{
			name: "side_invalid_menu",
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
								"shortcut": cu.IM{"id": 12345},
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
			name: "list_event_call_shortcut",
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
								"shortcut": cu.IM{},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "shortcut",
					"value": cu.IM{
						"row": cu.IM{
							"func_name": "test",
							"lstype":    "shortcut",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "list_event_call_report",
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
								"shortcut": cu.IM{},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "shortcut",
					"value": cu.IM{
						"row": cu.IM{
							"report_key": "test",
							"lstype":     "report",
							"template":   "{\"fields\": {\"field_1\": {\"description\": \"field_1_description\", \"field_type\": \"field_type_1\"}}}",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "params",
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
								"shortcut": cu.IM{},
							},
						},
					},
				},
				Name: ct.EditorEventField,
				Value: cu.IM{"name": "params",
					"value": "value",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShortcutService(tt.cls)
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
