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

func TestTransService_Data(t *testing.T) {
	type fields struct {
		cls *ClientService
	}
	type args struct {
		evt    ct.ResponseEvent
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
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1, "trans_type": md.TransTypeCash.String()}}, nil
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
							Data: cu.IM{},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
				},
				params: cu.IM{
					"trans_id":   1,
					"trans_type": md.TransTypeCash.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "trans_invoice_new",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{
										{"id": 1, "config_key": "default_taxcode", "config_value": "VAT00"},
										{"id": 2, "config_key": "default_currency", "config_value": "USD"},
										{"id": 3, "config_key": "default_deadline", "config_value": "30"},
										{"id": 4, "config_key": "default_paidtype", "config_value": "ONLINE"},
										{"id": 5, "config_key": "default_bank", "config_value": "1234567890"},
										{"id": 6, "config_key": "default_chest", "config_value": "1234567890"},
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
				},
				params: cu.IM{
					"trans_type": md.TransTypeInvoice.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "trans_bank_new",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{
										{"id": 1, "config_key": "default_taxcode", "config_value": "VAT00"},
										{"id": 2, "config_key": "default_currency", "config_value": "USD"},
										{"id": 3, "config_key": "default_deadline", "config_value": "30"},
										{"id": 4, "config_key": "default_paidtype", "config_value": "ONLINE"},
										{"id": 5, "config_key": "default_bank", "config_value": "1234567890"},
										{"id": 6, "config_key": "default_chest", "config_value": "1234567890"},
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
				},
				params: cu.IM{
					"trans_type": md.TransTypeBank.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "trans_cash_new",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{
										{"id": 1, "config_key": "default_taxcode", "config_value": "VAT00"},
										{"id": 2, "config_key": "default_currency", "config_value": "USD"},
										{"id": 3, "config_key": "default_deadline", "config_value": "30"},
										{"id": 4, "config_key": "default_paidtype", "config_value": "ONLINE"},
										{"id": 5, "config_key": "default_bank", "config_value": "1234567890"},
										{"id": 6, "config_key": "default_chest", "config_value": "1234567890"},
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
				},
				params: cu.IM{
					"trans_type": md.TransTypeCash.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "trans_error",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									if queries[0].From == "trans" {
										return []cu.IM{}, errors.New("error")
									}
									return []cu.IM{{"id": 1, "trans_type": md.TransTypeCash.String()}}, nil
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
			},
			wantErr: true,
		},
		{
			name: "query_base_error",
			fields: fields{
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
									return []cu.IM{{"id": 1, "trans_type": md.TransTypeInvoice.String()}}, nil
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
			},
			wantErr: true,
		},
		{
			name: "query_ext_error",
			fields: fields{
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
									return []cu.IM{{"id": 1, "trans_type": md.TransTypeInvoice.String()}}, nil
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
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{
				cls: tt.fields.cls,
			}
			_, err := s.Data(tt.args.evt, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransService.Data() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransService_Response(t *testing.T) {
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
							Db:     &md.TestDriver{Config: cu.IM{}},
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
							Db: &md.TestDriver{Config: cu.IM{}},
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "editor_map_value"}, "value": cu.IM{"value": "code", "model": "trans", "map_field": "tags"}},
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
									"trans": cu.IM{
										"id": 1,
										"trans_meta": cu.IM{
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
									"trans": cu.IM{
										"id": 1,
										"trans_meta": cu.IM{
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
									"trans": cu.IM{
										"id": 1,
										"trans_meta": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{
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
									"trans": cu.IM{"id": 12345},
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
		/*
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
										"trans": cu.IM{"id": 12345},
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
		*/
		{
			name: "editor_delete",
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
									"trans": cu.IM{"id": 12345},
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "product"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_employee",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "employee"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_project",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "project"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_shipping",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "shipping"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_trans",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "trans"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_customer",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "customer"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_transmovement_new",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "transmovement_new"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_transpayment_new",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "transpayment_new"}, "value": cu.IM{"create_trans_type": md.TransTypeBank.String()}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_transitem_new",
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
									"trans": cu.IM{
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
					Value: cu.IM{"data": cu.IM{"next": "transitem_new"}, "value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_trans_copy",
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
								"Update": func(data md.Update) (int64, error) {
									return 1, nil
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
									"trans": cu.IM{
										"id":            1,
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"trans_date":    "2025-01-01",
										"currency_code": "USD",
										"trans_meta": cu.IM{
											"due_time": "2025-01-01",
										},
									},
									"config_data": []cu.IM{
										{"config_key": "default_deadline", "config_value": 8},
									},
									"items": []cu.IM{
										{"id": 1, "item_meta": cu.IM{
											"qty":        1,
											"net_amount": 100,
											"vat_amount": 10,
											"amount":     110,
										}},
									},
									"movements": []cu.IM{
										{"id": 1,
											"product_code": "P1",
											"place_code":   "L1",
											"movement_meta": cu.IM{
												"qty":        1,
												"net_amount": 100,
												"vat_amount": 10,
												"amount":     110,
											}},
									},
									"payments": []cu.IM{
										{"id": 1, "payment_meta": cu.IM{
											"amount": 100,
										}},
									},
									"payment_link": []cu.IM{
										{"id": 1, "link_meta": cu.IM{
											"amount": 100,
										}},
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
					Value: cu.IM{"data": cu.IM{"next": "trans_copy"},
						"value": cu.IM{}},
				},
			},
			wantErr: false,
		},
		{
			name: "next_trans_create",
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
								"Update": func(data md.Update) (int64, error) {
									return 1, nil
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
									"trans": cu.IM{
										"id":            1,
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"trans_date":    "2025-01-01",
										"currency_code": "USD",
										"trans_meta": cu.IM{
											"due_time": "2025-01-01",
										},
									},
									"config_data": []cu.IM{
										{"config_key": "default_deadline", "config_value": 8},
									},
									"items": []cu.IM{
										{"id": 1,
											"product_code": "P1",
											"item_meta": cu.IM{
												"qty":        1,
												"net_amount": 100,
												"vat_amount": 10,
												"amount":     110,
											}},
									},
									"movements": []cu.IM{
										{"id": 1,
											"product_code": "P1",
											"place_code":   "L1",
											"movement_meta": cu.IM{
												"qty":        1,
												"net_amount": 100,
												"vat_amount": 10,
												"amount":     110,
											}},
									},
									"payments": []cu.IM{
										{"id": 1, "payment_meta": cu.IM{
											"amount": 100,
										}},
									},
									"payment_link": []cu.IM{
										{"id": 1, "link_meta": cu.IM{
											"amount": 100,
										}},
									},
									"products": []cu.IM{
										{"id": 1, "code": "P1", "product_meta": cu.IM{
											"unit":         "PC",
											"barcode_type": "CODE_128",
											"barcode_data": "1234567890",
											"barcode_qty":  1,
											"notes":        "Test product",
											"inactive":     false,
											"tags":         []string{"test"},
										}},
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
					Value: cu.IM{"data": cu.IM{"next": "trans_create"},
						"value": cu.IM{
							"create_trans_type": md.TransTypeWorksheet.String(), "create_direction": md.DirectionOut.String(),
							"status": md.TransStatusCancellation.String(), "create_netto": true,
						}},
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
									"trans": cu.IM{"id": 12345},
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
									"trans": cu.IM{"id": 12345},
									"dirty": true,
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
									"trans": cu.IM{"id": 12345},
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
									"trans": cu.IM{"id": 12345},
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
									"trans": cu.IM{"id": 12345},
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
									"trans": cu.IM{"id": 12345},
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
									"trans": cu.IM{"id": 12345},
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
									"trans": cu.IM{"id": 12345,
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
									"trans": cu.IM{"id": 12345},
									"view":  "contacts",
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
									"trans": cu.IM{"id": 12345, "events": []cu.IM{
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
									"trans":     cu.IM{"id": 12345},
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
									"trans":     cu.IM{"id": 12345},
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
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
									"trans": cu.IM{
										"id": 12345,
										"trans_map": cu.IM{
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
									"trans": cu.IM{
										"id": 12345,
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
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
									"trans": cu.IM{
										"id": 12345,
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
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
			name: "trans_type",
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
									"trans": cu.IM{
										"id": 12345,
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "trans",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "trans_type",
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
									"trans": cu.IM{
										"id": 12345,
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "trans",
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
			name: "closed",
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
									"trans": cu.IM{
										"id": 12345,
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "trans",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "closed",
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
									"trans": cu.IM{
										"id": 12345,
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "trans",
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
									"trans": cu.IM{
										"id": 12345,
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
											"demo_string": "tag1",
										}},
									"view": "trans",
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
			s := &TransService{
				cls: tt.fields.cls,
			}
			_, err := s.Response(tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransService.Response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
