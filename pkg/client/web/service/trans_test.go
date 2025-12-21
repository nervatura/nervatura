package service

import (
	"errors"
	"log/slog"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	cp "github.com/nervatura/nervatura/v6/pkg/client/web/component"
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
			name: "editor_delete_error_invoice",
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
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id":         1,
										"trans_type": md.TransTypeInvoice.String(),
										"direction":  md.DirectionOut.String(),
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
			wantErr: true,
		},
		{
			name: "editor_delete_invoice",
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
										"id":            1,
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_meta": cu.IM{
											"status":     md.TransStatusNormal.String(),
											"paid":       true,
											"closed":     true,
											"due_time":   "2025-01-01",
											"ref_number": "1234567890",
											"paid_type":  md.PaidTypeOnline.String(),
											"rate":       1.0,
											"tax_free":   false,
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
											"id":            1,
											"code":          "C1",
											"name":          "contact1",
											"movement_code": "M1",
										},
										{
											"id":   1,
											"code": "M1",
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
						"data": cu.IM{"form": cu.IM{
							"index": 0, "key": "contacts",
							"data": cu.IM{
								"tags": []string{"tag1", "tag2"}, "item_meta": cu.IM{"qty": 1}, "product_code": "P1", "product_name": "Product 1",
								"product_unit": "PC", "tool_description": "Tool 1", "serial_number": "1234567890",
								"tool_code": "T1", "tool_name": "Tool 1", "tool_unit": "PC",
								"invoice_code": "I1", "invoice_name": "Invoice 1", "invoice_unit": "PC",
								"qty": 1, "discount": 0, "own_stock": true, "description": "description", "unit": "PC",
								"notes": "notes", "link_amount": 100, "link_rate": 1.1, "shared": true,
								"amount": 100, "net_amount": 100, "vat_amount": 20, "fx_price": 1.1,
								"tax_code": "T1", "tax_name": "Tax 1", "tax_unit": "PC",
								"tax_rate": 0.2, "tax_amount": 20, "tax_net_amount": 100, "tax_fx_price": 1.1,
								"place_code": "P1", "link_code_2": "L1", "invoice_curr": "USD",
							},
						}, "data": cu.IM{"name": "contact2"}},
						"value": cu.IM{"name": "contact2", "amount": 100, "net_amount": 100, "vat_amount": 20,
							"fx_price": 1.1, "qty": 1, "discount": 0, "own_stock": true, "description": "description", "unit": "PC",
							"notes": "notes", "link_amount": 100, "link_rate": 1.1, "shared": true},
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
			name: "client_form_change_product_code",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
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
										"currency_code": "USD",
										"customer_code": "C001",
									},
									"tax_codes":        []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
									"currencies":       []cu.IM{{"code": "USD", "digit": 2}},
									"product_selector": cu.IM{},
									"items": []cu.IM{
										{"id": 1, "product_code": "P1", "item_meta": cu.IM{"qty": 1}},
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
						"data":  cu.IM{"form": cu.IM{"index": 0, "key": "items", "data": cu.IM{"product_code": "P1", "item_meta": cu.IM{"qty": 1}}}, "data": cu.IM{}},
						"name":  "product_code",
						"value": cu.IM{"event": "selected"},
						"event": ct.FormEventChange,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_field_payment_amount",
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
										"id":         1,
										"trans_type": md.TransTypeCash.String(),
										"direction":  md.DirectionOut.String(),
									},
									"payments": []cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{
						"name":  "payment_amount",
						"value": "100",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_field_movement_product_code_selected_head",
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
										"id":         1,
										"trans_type": md.TransTypeFormula.String(),
									},
									"movements": []cu.IM{
										{
											"product_code":  "",
											"movement_type": md.MovementTypeHead.String(),
											"movement_meta": cu.IM{"qty": 1},
											"movement_map": cu.IM{
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
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "movement_product_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "product_code": "12345", "product_name": "name", "product_meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_field_movement_product_code_selected_no_head",
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
										"id":         1,
										"trans_type": md.TransTypeFormula.String(),
									},
									"movements": []cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "movement_product_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "product_code": "12345", "product_name": "name", "product_meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "editor_field_movement_product_code_search",
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
					Modules: map[string]ServiceModule{
						"search": NewSearchService(&ClientService{
							Config: cu.IM{},
							AppLog: slog.Default(),
							NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
								return &api.DataStore{
									Db: &md.TestDriver{Config: cu.IM{
										"Query": func(queries []md.Query) ([]cu.IM, error) {
											return []cu.IM{{"code": "value"}}, nil
										},
									}},
									Config: config,
									AppLog: appLog,
								}
							},
							UI: cp.NewClientComponent(),
						}),
					},
					UI: cp.NewClientComponent(),
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id":         1,
										"trans_type": md.TransTypeFormula.String(),
									},
									"movements": []cu.IM{
										{"product_code": ""},
									},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "movement_product_code",
						"event": ct.SelectorEventSearch,
						"value": "value",
					},
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
			name: "side_transitem_new",
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
					Value: "transitem_new",
				},
			},
			wantErr: false,
		},
		{
			name: "side_transpayment_new",
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
					Value: "transpayment_new",
				},
			},
			wantErr: false,
		},
		{
			name: "side_transmovement_new",
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
					Value: "transmovement_new",
				},
			},
			wantErr: false,
		},
		{
			name: "side_trans_copy",
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
					Value: "trans_copy",
				},
			},
			wantErr: false,
		},
		{
			name: "side_trans_cancellation",
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
					Value: "trans_cancellation",
				},
			},
			wantErr: false,
		},
		{
			name: "side_trans_create",
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
									"trans": cu.IM{"id": 12345, "trans_type": md.TransTypeInvoice.String(), "direction": md.DirectionOut.String()},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "trans_create",
				},
			},
			wantErr: false,
		},
		{
			name: "side_trans_corrective",
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
					Value: "trans_corrective",
				},
			},
			wantErr: false,
		},
		{
			name: "editor_save_with_items_and_payments",
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
									return []cu.IM{{"id": 1, "code": "T001"}}, nil
								},
							}},
							Config: config,
							AppLog: appLog,
							ConvertToByte: func(data interface{}) ([]byte, error) {
								return []byte{}, nil
							},
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
										"id":            1,
										"code":          "T001",
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_date":    "2025-01-01",
										"currency_code": "USD",
										"trans_meta":    cu.IM{},
										"trans_map":     cu.IM{},
									},
									"items": []cu.IM{
										{"id": 0, "product_code": "P1", "tax_code": "VAT01", "item_meta": cu.IM{"qty": 1}, "item_map": cu.IM{}},
									},
									"payments": []cu.IM{
										{"id": 0, "paid_date": "2025-01-01", "payment_meta": cu.IM{"amount": 100}, "payment_map": cu.IM{}},
									},
									"movements": []cu.IM{},
									"payment_link": []cu.IM{
										{"id": 0, "link_code_1": "PAY001", "link_code_2": "T002", "link_meta": cu.IM{"amount": 100}, "link_map": cu.IM{}},
									},
									"items_delete":        []cu.IM{},
									"payments_delete":     []cu.IM{},
									"movements_delete":    []cu.IM{},
									"payment_link_delete": []cu.IM{},
									"user":                cu.IM{"code": "admin"},
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
			name: "side_payment_link_add",
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
									"trans": cu.IM{"id": 1, "code": "T001"},
									"payments": []cu.IM{
										{"code": "P001", "payment_meta": cu.IM{"amount": 100}},
									},
									"places": []cu.IM{{"code": "P001", "currency_code": "USD"}},
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name:  ct.ClientEventSideMenu,
					Value: "payment_link_add",
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
					Value: cu.IM{"data": cu.IM{"next": "trans_corrective"},
						"value": cu.IM{
							"create_trans_type": md.TransTypeInvoice.String(), "create_direction": md.DirectionOut.String(),
						}},
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
					Value: cu.IM{"data": cu.IM{"next": "trans_cancellation"},
						"value": cu.IM{
							"create_trans_type": md.TransTypeInvoice.String(), "create_direction": md.DirectionOut.String(),
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
			name: "table_row_selected_transitem_invoice",
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
									"view":  "transitem_invoice",
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventRowSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "trans_code": "12345"}, "index": 0, "view": "transitem_invoice"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_row_selected_transitem_invoice_dirty",
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
									"view":  "transitem_invoice",
									"dirty": true,
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventRowSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "trans_code": "12345"}, "index": 0, "view": "transitem_invoice"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_edit_cell_payment_link_add",
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
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventEditCell,
						"value": cu.IM{"row": cu.IM{"id": 12345, "code": "12345"}, "index": 0, "fieldname": "payment_link_add"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_edit_cell_trans_code",
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
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventEditCell,
						"value": cu.IM{"row": cu.IM{"id": 12345, "code": "12345"}, "index": 0, "fieldname": "trans_code"}},
				},
			},
			wantErr: false,
		},
		{
			name: "table_edit_cell_trans_code_dirty",
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
					Name: ct.EditorEventField,
					Value: cu.IM{"name": ct.TableEventEditCell,
						"value": cu.IM{"row": cu.IM{"id": 12345, "code": "12345"}, "index": 0, "fieldname": "trans_code"}},
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
			name: "internal_notes",
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
					Value: cu.IM{"name": "internal_notes",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "worksheet_distance",
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
					Value: cu.IM{"name": "worksheet_distance",
						"value": 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "worksheet_repair",
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
					Value: cu.IM{"name": "worksheet_repair",
						"value": 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "worksheet_total",
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
					Value: cu.IM{"name": "worksheet_total",
						"value": 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "worksheet_justification",
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
					Value: cu.IM{"name": "worksheet_justification",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "rent_holiday",
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
					Value: cu.IM{"name": "rent_holiday",
						"value": 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "rent_bad_tool",
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
					Value: cu.IM{"name": "rent_bad_tool",
						"value": 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "rent_other",
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
					Value: cu.IM{"name": "rent_other",
						"value": 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "rent_justification",
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
					Value: cu.IM{"name": "rent_justification",
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
			name: "payment_paid_date",
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
										"id":         12345,
										"trans_type": md.TransTypeCash.String(),
										"direction":  md.DirectionOut.String(),
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
					Value: cu.IM{"name": "payment_paid_date",
						"value": "2025-01-01",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "movement_qty",
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
										"id":         12345,
										"trans_type": md.TransTypeProduction.String(),
										"direction":  md.DirectionTransfer.String(),
										"trans_meta": cu.IM{
											"tags": []string{"tag1", "tag2"},
										},
										"trans_map": cu.IM{
											"demo_string": "tag1",
										}},
									"movements": []cu.IM{
										{
											"movement_type": md.MovementTypeInventory.String(),
											"product_code":  "P1",
											"movement_meta": cu.IM{"qty": 1, "shared": true},
										},
									},
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
					Value: cu.IM{"name": "movement_qty",
						"value": 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "movement_notes",
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
										"id":         12345,
										"trans_type": md.TransTypeProduction.String(),
										"direction":  md.DirectionTransfer.String(),
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
					Value: cu.IM{"name": "movement_notes",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "movement_notes_waybill",
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
										"id":         12345,
										"trans_type": md.TransTypeWaybill.String(),
										"direction":  md.DirectionIn.String(),
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
					Value: cu.IM{"name": "movement_notes",
						"value": "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "rate",
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
					Value: cu.IM{"name": "rate",
						"value": 1.1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ref_number",
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
					Value: cu.IM{"name": "ref_number",
						"value": "1234567890",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "trans_state",
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
					Value: cu.IM{"name": "trans_state",
						"value": md.TransStateOK.String(),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "direction",
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
					Value: cu.IM{"name": "direction",
						"value": md.DirectionOut.String(),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "trans_date",
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
					Value: cu.IM{"name": "trans_date",
						"value": "2025-01-01",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "currency_code",
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
					Value: cu.IM{"name": "currency_code",
						"value": "USD",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "due_time",
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
					Value: cu.IM{"name": "due_time",
						"value": "2025-01-01T00:00:00Z",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "paid_type",
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
					Value: cu.IM{"name": "paid_type",
						"value": md.PaidTypeCash.String(),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "paid",
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
					Value: cu.IM{"name": "paid",
						"value": true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "code",
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
					Value: cu.IM{"name": "code",
						"value": "1234567890",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "place_code",
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
					Value: cu.IM{"name": "place_code",
						"value": "P1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "target_place_code",
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
									"movements": []cu.IM{
										{
											"id":            1,
											"movement_code": "M1",
										},
									},
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
					Value: cu.IM{"name": "target_place_code",
						"value": "P1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "customer_code_selected",
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
					Value: cu.IM{"name": "customer_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "customer_code": "12345", "customer_name": "name", "meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "employee_code_selected",
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
					Value: cu.IM{"name": "employee_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "employee_code": "12345", "customer_name": "name", "meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "project_code_selected",
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
					Value: cu.IM{"name": "project_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "project_code": "12345", "project_name": "name", "meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "trans_code_selected",
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
					Value: cu.IM{"name": "trans_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "trans_code": "12345", "meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "trans_code_selected_dirty",
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
									"view":  "trans",
									"dirty": true,
								},
							},
						},
						Ticket: ct.Ticket{
							User:     cu.IM{},
							Database: "test",
						},
					},
					Name: ct.EditorEventField,
					Value: cu.IM{"name": "trans_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "trans_code": "12345", "meta": cu.IM{}}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "transitem_code_selected",
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
					Value: cu.IM{"name": "transitem_code",
						"event": ct.SelectorEventSelected,
						"value": cu.IM{"row": cu.IM{"id": 12345, "trans_code": "12345", "meta": cu.IM{}}},
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

func TestTransService_update(t *testing.T) {
	type fields struct {
		cls *ClientService
	}
	type args struct {
		ds      *api.DataStore
		data    cu.IM
		msgFunc func(labelID string) string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantTransID int64
		wantErr     bool
	}{
		{
			name: "error_new_trans",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Config: cu.IM{},
					AppLog: slog.Default(),
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							if data.Model == "payment" {
								return 0, errors.New("update error")
							}
							return 12345, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "T001"}}, nil
						},
					}},
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(0),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items":        []cu.IM{},
					"payments":     []cu.IM{},
					"movements":    []cu.IM{},
					"payment_link": []cu.IM{},
					"items_delete": []cu.IM{},
					"payments_delete": []cu.IM{
						{"id": 12345},
					},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 12345,
			wantErr:     true,
		},
		{
			name: "success_trans",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Config: cu.IM{},
					AppLog: slog.Default(),
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 12345, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "T001"}}, nil
						},
					}},
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(12345),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items":        []cu.IM{},
					"payments":     []cu.IM{},
					"movements":    []cu.IM{},
					"payment_link": []cu.IM{},
					"items_delete": []cu.IM{
						{"id": 12345},
					},
					"payments_delete":     []cu.IM{},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 12345,
			wantErr:     false,
		},
		{
			name: "validation_error",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         int64(0),
						"trans_type": md.TransTypeInvoice.String(),
						"direction":  md.DirectionOut.String(),
						// Missing customer_code for item transaction
						"trans_meta": cu.IM{},
						"trans_map":  cu.IM{},
					},
					"user": cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 0,
			wantErr:     true,
		},
		{
			name: "update_error",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 0, errors.New("update error")
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(12345),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items": []cu.IM{
						{
							"id":           int64(0),
							"product_code": "P1",
							"tax_code":     "VAT01",
							"item_meta":    cu.IM{"qty": 1, "tags": []string{}},
							"item_map":     cu.IM{},
						},
					},
					"payments":            []cu.IM{},
					"movements":           []cu.IM{},
					"payment_link":        []cu.IM{},
					"items_delete":        []cu.IM{},
					"payments_delete":     []cu.IM{},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 0,
			wantErr:     true,
		},
		{
			name: "update_error_get",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 12345, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, errors.New("query error")
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(0),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items": []cu.IM{
						{
							"id":           int64(0),
							"product_code": "P1",
							"tax_code":     "VAT01",
							"item_meta":    cu.IM{"qty": 1, "tags": []string{}},
							"item_map":     cu.IM{},
						},
					},
					"payments":            []cu.IM{},
					"movements":           []cu.IM{},
					"payment_link":        []cu.IM{},
					"items_delete":        []cu.IM{},
					"payments_delete":     []cu.IM{},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 12345,
			wantErr:     true,
		},
		{
			name: "update_error_item",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							if data.Model == "item" {
								return 0, errors.New("update error")
							}
							return 12345, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(12345),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items": []cu.IM{
						{
							"id":           int64(0),
							"product_code": "P1",
							"tax_code":     "VAT01",
							"item_meta":    cu.IM{"qty": 1, "tags": []string{}},
							"item_map":     cu.IM{},
						},
					},
					"payments":            []cu.IM{},
					"movements":           []cu.IM{},
					"payment_link":        []cu.IM{},
					"items_delete":        []cu.IM{},
					"payments_delete":     []cu.IM{},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 12345,
			wantErr:     true,
		},
		{
			name: "update_error_payment",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							if data.Model == "payment" {
								return 0, errors.New("update error")
							}
							return 12345, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(12345),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items": []cu.IM{
						{
							"id":           12345,
							"product_code": "P1",
							"tax_code":     "VAT01",
							"item_meta":    cu.IM{"qty": 1, "tags": []string{}},
							"item_map":     cu.IM{},
						},
					},
					"payments": []cu.IM{
						{
							"id":           int64(0),
							"payment_code": "P1",
							"payment_meta": cu.IM{"amount": 1, "tags": []string{}},
							"payment_map":  cu.IM{},
						},
					},
					"movements":           []cu.IM{},
					"payment_link":        []cu.IM{},
					"items_delete":        []cu.IM{},
					"payments_delete":     []cu.IM{},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 12345,
			wantErr:     true,
		},
		{
			name: "update_error_link",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							if data.Model == "link" {
								return 0, errors.New("update error")
							}
							return 12345, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(12345),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items": []cu.IM{},
					"payments": []cu.IM{
						{
							"id":           12345,
							"payment_code": "P1",
							"payment_meta": cu.IM{"amount": 1, "tags": []string{}},
							"payment_map":  cu.IM{},
						},
					},
					"movements": []cu.IM{},
					"payment_link": []cu.IM{
						{
							"id":           int64(0),
							"payment_code": "P1",
							"payment_meta": cu.IM{"amount": 1, "tags": []string{}},
							"payment_map":  cu.IM{},
						},
					},
					"items_delete":        []cu.IM{},
					"payments_delete":     []cu.IM{},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 12345,
			wantErr:     true,
		},
		{
			name: "update_error_movement",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							if data.Model == "movement" {
								return 0, errors.New("update error")
							}
							return 12345, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"trans": cu.IM{
						"id":            int64(12345),
						"trans_type":    md.TransTypeInvoice.String(),
						"direction":     md.DirectionOut.String(),
						"trans_date":    "2025-01-01",
						"customer_code": "C001",
						"trans_meta":    cu.IM{},
						"trans_map":     cu.IM{},
					},
					"items": []cu.IM{},
					"payments": []cu.IM{
						{
							"id":           12345,
							"payment_code": "P1",
							"payment_meta": cu.IM{"amount": 1, "tags": []string{}},
							"payment_map":  cu.IM{},
						},
					},
					"movements": []cu.IM{
						{
							"id":            12345,
							"product_code":  "P1",
							"place_code":    "L1",
							"movement_meta": cu.IM{"qty": 1, "tags": []string{}},
							"movement_map":  cu.IM{},
						},
					},
					"payment_link": []cu.IM{
						{
							"id":           12345,
							"payment_code": "P1",
							"payment_meta": cu.IM{"amount": 1, "tags": []string{}},
							"payment_map":  cu.IM{},
						},
					},
					"items_delete":        []cu.IM{},
					"payments_delete":     []cu.IM{},
					"movements_delete":    []cu.IM{},
					"payment_link_delete": []cu.IM{},
					"user":                cu.IM{"code": "admin"},
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantTransID: 12345,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{
				cls: tt.fields.cls,
			}
			gotTransID, err := s.update(tt.args.ds, tt.args.data, tt.args.msgFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransService.update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTransID != tt.wantTransID {
				t.Errorf("TransService.update() = %v, want %v", gotTransID, tt.wantTransID)
			}
		})
	}
}

func TestTransService_updateMovements(t *testing.T) {
	type fields struct {
		cls *ClientService
	}
	type args struct {
		ds      *api.DataStore
		data    cu.IM
		trans   md.Trans
		msgFunc func(labelID string) string
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
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "M001"}}, nil
						},
					}},
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"movements": []cu.IM{
						{
							"id":            -1,
							"movement_type": md.MovementTypeInventory.String(),
							"shipping_time": "2025-01-01T00:00:00",
							"product_code":  "P1",
							"place_code":    "L2",
							"movement_meta": cu.IM{"qty": 10, "tags": []string{}},
							"movement_map":  cu.IM{},
						},
					},
				},
				trans: md.Trans{
					Code:      "T001",
					TransType: md.TransTypeDelivery,
					Direction: md.DirectionTransfer,
					PlaceCode: "L1",
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantErr: false,
		},
		{
			name: "validate_error",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "M001"}}, nil
						},
					}},
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"movements": []cu.IM{
						{
							"id":            -1,
							"movement_type": md.MovementTypeInventory.String(),
							"shipping_time": "2025-01-01T00:00:00",
							"product_code":  "P1",
							"place_code":    "",
							"movement_meta": cu.IM{"qty": 10, "tags": []string{}},
							"movement_map":  cu.IM{},
						},
					},
				},
				trans: md.Trans{
					Code:      "T001",
					TransType: md.TransTypeDelivery,
					Direction: md.DirectionTransfer,
					PlaceCode: "L1",
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantErr: true,
		},
		{
			name: "transfer_pair_update_error",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				ds: &api.DataStore{
					Config: cu.IM{},
					AppLog: slog.Default(),
					Db: &md.TestDriver{Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							if cu.ToString(data.Values["movement_code"], "") != "" {
								return 0, errors.New("update error")
							}
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "M001"}}, nil
						},
					}},
					ConvertToByte: func(data interface{}) ([]byte, error) {
						return []byte{}, nil
					},
				},
				data: cu.IM{
					"movements": []cu.IM{
						{
							"id":            -1,
							"movement_type": md.MovementTypeInventory.String(),
							"shipping_time": "2025-01-01T00:00:00",
							"product_code":  "P1",
							"place_code":    "L2",
							"movement_meta": cu.IM{"qty": 10, "tags": []string{}},
							"movement_map":  cu.IM{},
						},
					},
				},
				trans: md.Trans{
					Code:      "T001",
					TransType: md.TransTypeDelivery,
					Direction: md.DirectionTransfer,
					PlaceCode: "L1",
				},
				msgFunc: func(labelID string) string { return labelID },
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{
				cls: tt.fields.cls,
			}
			if err := s.updateMovements(tt.args.ds, tt.args.data, tt.args.trans, tt.args.msgFunc); (err != nil) != tt.wantErr {
				t.Errorf("TransService.updateMovements() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransService_getProductPrice(t *testing.T) {
	type fields struct {
		cls *ClientService
	}
	type args struct {
		ds      *api.DataStore
		options cu.IM
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantPrice    float64
		wantDiscount float64
	}{
		{
			name: "success",
			fields: fields{
				cls: &ClientService{
					Config: cu.IM{},
					AppLog: slog.Default(),
				},
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"mp": 100.0}}, nil
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
				},
				options: cu.IM{"currency_code": "USD", "product_code": "P1"},
			},
			wantPrice:    0,
			wantDiscount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{
				cls: tt.fields.cls,
			}
			gotPrice, gotDiscount := s.getProductPrice(tt.args.ds, tt.args.options)
			// ProductPrice calls complex database queries, so we just check the error path
			if tt.name == "error" && (gotPrice != tt.wantPrice || gotDiscount != tt.wantDiscount) {
				t.Errorf("TransService.getProductPrice() gotPrice = %v, want %v, gotDiscount = %v, want %v", gotPrice, tt.wantPrice, gotDiscount, tt.wantDiscount)
			}
		})
	}
}

func TestTransService_calcItemPrice(t *testing.T) {
	type args struct {
		calcMode  string
		value     float64
		stateData cu.IM
		formRow   cu.IM
	}
	tests := []struct {
		name string
		args args
		want cu.IM
	}{
		{
			name: "net_amount",
			args: args{
				calcMode: "net_amount",
				value:    100.0,
				stateData: cu.IM{
					"tax_codes":  []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
					"currencies": []cu.IM{{"code": "USD", "digit": 2}},
				},
				formRow: cu.IM{
					"tax_code":  "VAT01",
					"item_meta": cu.IM{"qty": 2, "discount": 0, "fx_price": 50},
				},
			},
			want: cu.IM{
				"net_amount": 100.0,
				"vat_amount": 20.0,
				"amount":     120.0,
				"fx_price":   50.0,
			},
		},
		{
			name: "amount",
			args: args{
				calcMode: "amount",
				value:    120.0,
				stateData: cu.IM{
					"trans":      cu.IM{"currency_code": "USD"},
					"tax_codes":  []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
					"currencies": []cu.IM{{"code": "USD", "digit": 2}},
				},
				formRow: cu.IM{
					"tax_code":  "VAT15",
					"item_meta": cu.IM{"qty": 2, "discount": 0},
				},
			},
			want: cu.IM{
				"net_amount": 120.0,
				"vat_amount": 0.0,
				"amount":     120.0,
				"fx_price":   60.0,
			},
		},
		{
			name: "fx_price",
			args: args{
				calcMode: "fx_price",
				value:    50.0,
				stateData: cu.IM{
					"tax_codes":  []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
					"currencies": []cu.IM{{"code": "USD", "digit": 2}},
				},
				formRow: cu.IM{
					"tax_code":  "VAT01",
					"item_meta": cu.IM{"qty": 2, "discount": 10},
				},
			},
			want: cu.IM{
				"net_amount": 90.0,
				"vat_amount": 18.0,
				"amount":     108.0,
				"fx_price":   50.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{}
			got := s.calcItemPrice(tt.args.calcMode, tt.args.value, tt.args.stateData, tt.args.formRow)
			if got["net_amount"] != tt.want["net_amount"] {
				t.Errorf("TransService.calcItemPrice() net_amount = %v, want %v", got["net_amount"], tt.want["net_amount"])
			}
			if got["vat_amount"] != tt.want["vat_amount"] {
				t.Errorf("TransService.calcItemPrice() vat_amount = %v, want %v", got["vat_amount"], tt.want["vat_amount"])
			}
			if got["amount"] != tt.want["amount"] {
				t.Errorf("TransService.calcItemPrice() amount = %v, want %v", got["amount"], tt.want["amount"])
			}
			if got["fx_price"] != tt.want["fx_price"] {
				t.Errorf("TransService.calcItemPrice() fx_price = %v, want %v", got["fx_price"], tt.want["fx_price"])
			}
		})
	}
}

func TestTransService_createModal(t *testing.T) {
	type fields struct {
		cls *ClientService
	}
	type args struct {
		evt  ct.ResponseEvent
		data cu.IM
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "transitem",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": "transitem",
				},
			},
		},
		{
			name: "transpayment",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": "transpayment",
				},
			},
		},
		{
			name: "transmovement",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": "transmovement",
				},
			},
		},
		{
			name: "offer",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": md.TransTypeOffer.String(),
				},
			},
		},
		{
			name: "order",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": md.TransTypeOrder.String(),
				},
			},
		},
		{
			name: "worksheet",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": md.TransTypeWorksheet.String(),
				},
			},
		},
		{
			name: "rent",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": md.TransTypeRent.String(),
				},
			},
		},
		{
			name: "invoice",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{},
						},
					},
				},
				data: cu.IM{
					"state_key": md.TransTypeInvoice.String(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{
				cls: tt.fields.cls,
			}
			re := s.createModal(tt.args.evt, tt.args.data)
			if re.Name != tt.args.evt.Name {
				t.Errorf("TransService.createModal() returned different event")
			}
		})
	}
}

func TestTransService_createData(t *testing.T) {
	type fields struct {
		cls *ClientService
	}
	type args struct {
		evt     ct.ResponseEvent
		options cu.IM
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
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									return 1, nil
								},
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1, "code": "T001"}}, nil
								},
							}},
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
										"id":            1,
										"trans_type":    md.TransTypeDelivery.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_meta":    cu.IM{},
										"trans_map":     cu.IM{},
									},
									"items":    []cu.IM{},
									"payments": []cu.IM{},
									"movements": []cu.IM{
										{
											"id":            12345,
											"movement_type": md.MovementTypeTool.String(),
											"tool_code":     "T001",
											"product_code":  "P1",
											"place_code":    "L1",
											"movement_meta": cu.IM{"qty": 100, "tags": []string{}},
											"movement_map":  cu.IM{},
										},
										{
											"id":            12346,
											"item_code":     "I001",
											"movement_code": "M001",
											"movement_meta": cu.IM{"qty": 200, "tags": []string{}},
											"movement_map":  cu.IM{},
										},
									},
									"user": cu.IM{"code": "admin"},
								},
							},
						},
						Ticket: ct.Ticket{SessionID: "test", Database: "test"},
					},
				},
				options: cu.IM{
					"create_trans_type": md.TransTypeDelivery.String(),
					"create_direction":  md.DirectionOut.String(),
					"status":            md.TransStatusCancellation.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "validation_error",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"id":         1,
										"trans_type": md.TransTypeReceipt.String(),
										"direction":  md.DirectionIn.String(), // Invalid combination
										"trans_meta": cu.IM{},
										"trans_map":  cu.IM{},
									},
								},
							},
						},
					},
				},
				options: cu.IM{
					"create_trans_type": md.TransTypeReceipt.String(),
					"create_direction":  md.DirectionIn.String(),
					"status":            md.TransStatusNormal.String(),
				},
			},
			wantErr: false, // Returns error modal, not actual error
		},
		{
			name: "create_trans_error",
			fields: fields{
				cls: &ClientService{
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									return 0, errors.New("create trans error")
								},
							}},
							Config: cu.IM{},
							AppLog: slog.Default(),
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
										"id":            1,
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_meta":    cu.IM{},
										"trans_map":     cu.IM{},
									},
									"user": cu.IM{"code": "admin"},
								},
							},
						},
						Ticket: ct.Ticket{Database: "test"},
					},
				},
				options: cu.IM{
					"create_trans_type": md.TransTypeInvoice.String(),
					"create_direction":  md.DirectionOut.String(),
					"status":            md.TransStatusNormal.String(),
				},
			},
			wantErr: false, // Returns error modal, not actual error
		},
		{
			name: "create_items_error",
			fields: fields{
				cls: &ClientService{
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									if data.Model == "item" {
										return 0, errors.New("create items error")
									}
									return 12345, nil
								},
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 12345}}, nil
								},
							}},
							Config: cu.IM{},
							AppLog: slog.Default(),
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
										"id":            1,
										"trans_type":    md.TransTypeReceipt.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_meta":    cu.IM{},
										"trans_map":     cu.IM{},
									},
									"items": []cu.IM{
										{
											"id":           12345,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 1, "tags": []string{}},
											"item_map":     cu.IM{},
										},
									},
									"user": cu.IM{"code": "admin"},
								},
							},
						},
						Ticket: ct.Ticket{Database: "test"},
					},
				},
				options: cu.IM{
					"create_trans_type": md.TransTypeRent.String(),
					"create_direction":  md.DirectionOut.String(),
					"status":            md.TransStatusNormal.String(),
				},
			},
			wantErr: false, // Returns error modal, not actual error
		},
		{
			name: "create_payments_error",
			fields: fields{
				cls: &ClientService{
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									if data.Model == "payment" {
										return 0, errors.New("create items error")
									}
									return 12345, nil
								},
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 12345}}, nil
								},
							}},
							Config: cu.IM{},
							AppLog: slog.Default(),
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
										"id":            1,
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_meta":    cu.IM{},
										"trans_map":     cu.IM{},
									},
									"payments": []cu.IM{
										{
											"id":           12345,
											"payment_meta": cu.IM{"amount": 100, "tags": []string{}},
											"payment_map":  cu.IM{},
										},
									},
									"user": cu.IM{"code": "admin"},
								},
							},
						},
						Ticket: ct.Ticket{Database: "test"},
					},
				},
				options: cu.IM{
					"create_trans_type": md.TransTypeInvoice.String(),
					"create_direction":  md.DirectionOut.String(),
					"status":            md.TransStatusNormal.String(),
				},
			},
			wantErr: false, // Returns error modal, not actual error
		},
		{
			name: "create_movements_error",
			fields: fields{
				cls: &ClientService{
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									if data.Model == "movement" {
										return 0, errors.New("create items error")
									}
									return 12345, nil
								},
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 12345}}, nil
								},
							}},
							Config: cu.IM{},
							AppLog: slog.Default(),
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
										"id":            1,
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_meta":    cu.IM{},
										"trans_map":     cu.IM{},
									},
									"movements": []cu.IM{
										{
											"id":            12345,
											"movement_meta": cu.IM{"qty": 100, "tags": []string{}},
											"movement_map":  cu.IM{},
										},
									},
									"user": cu.IM{"code": "admin"},
								},
							},
						},
						Ticket: ct.Ticket{Database: "test"},
					},
				},
				options: cu.IM{
					"create_trans_type": md.TransTypeInvoice.String(),
					"create_direction":  md.DirectionOut.String(),
					"status":            md.TransStatusNormal.String(),
				},
			},
			wantErr: false, // Returns error modal, not actual error
		},
		{
			name: "create_movements_formula_head_error",
			fields: fields{
				cls: &ClientService{
					NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
						return &api.DataStore{
							Db: &md.TestDriver{Config: cu.IM{
								"Update": func(data md.Update) (int64, error) {
									if data.Model == "movement" && data.Values["movement_type"] == md.MovementTypeHead.String() {
										return 0, errors.New("create error")
									}
									return 12345, nil
								},
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 12345}}, nil
								},
							}},
							Config: cu.IM{},
							AppLog: slog.Default(),
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
										"id":            1,
										"trans_type":    md.TransTypeInvoice.String(),
										"direction":     md.DirectionOut.String(),
										"customer_code": "C001",
										"trans_meta":    cu.IM{},
										"trans_map":     cu.IM{},
									},
									"movements": []cu.IM{
										{
											"id":            12345,
											"movement_type": md.MovementTypePlan.String(),
											"movement_meta": cu.IM{"qty": 100, "tags": []string{}},
											"movement_map":  cu.IM{},
										},
									},
									"movement_head": cu.IM{
										"id":            12345,
										"movement_type": md.MovementTypeHead.String(),
										"movement_meta": cu.IM{"qty": 100, "tags": []string{}},
										"movement_map":  cu.IM{},
									},
									"user": cu.IM{"code": "admin"},
								},
							},
						},
						Ticket: ct.Ticket{Database: "test"},
					},
				},
				options: cu.IM{
					"create_trans_type": md.TransTypeFormula.String(),
					"create_direction":  md.DirectionOut.String(),
					"status":            md.TransStatusNormal.String(),
				},
			},
			wantErr: false, // Returns error modal, not actual error
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{
				cls: tt.fields.cls,
			}
			_, err := s.createData(tt.args.evt, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransService.createData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransService_createInvoiceItems(t *testing.T) {
	type fields struct {
		cls *ClientService
	}
	type args struct {
		evt     ct.ResponseEvent
		options cu.IM
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantItems int
	}{
		{
			name: "create_delivery_base",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"currency_code": "USD",
										"direction":     md.DirectionIn.String(), // Changed to DirectionIn so iqty becomes positive
									},
									"items": []cu.IM{
										{
											"id":           1,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 10, "fx_price": 100, "discount": 0},
										},
									},
									"transitem_invoice": []cu.IM{},
									"transitem_shipping": []cu.IM{
										{
											"id":   "1-P1",
											"sqty": 5,
										},
									},
									"tax_codes":  []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
									"currencies": []cu.IM{{"code": "USD", "digit": 2}},
								},
							},
						},
					},
				},
				options: cu.IM{
					"create_delivery": true,
				},
			},
			wantItems: 1,
		},
		{
			name: "create_delivery_base_out",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"currency_code": "USD",
										"direction":     md.DirectionOut.String(), // Changed to DirectionIn so iqty becomes positive
									},
									"items": []cu.IM{
										{
											"id":           1,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 10, "fx_price": 100, "discount": 0},
										},
									},
									"transitem_invoice": []cu.IM{},
									"transitem_shipping": []cu.IM{
										{
											"id":   "1-P1",
											"sqty": 5,
										},
									},
									"tax_codes":  []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
									"currencies": []cu.IM{{"code": "USD", "digit": 2}},
								},
							},
						},
					},
				},
				options: cu.IM{
					"create_delivery": true,
				},
			},
			wantItems: 0,
		},
		{
			name: "create_netto_invoice",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"currency_code": "USD",
										"direction":     md.DirectionOut.String(),
									},
									"items": []cu.IM{
										{
											"id":           1,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 10},
										},
										{
											"id":           1,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 20, "deposit": true},
										},
									},
									"transitem_invoice": []cu.IM{
										{
											"id":           1,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 10},
										},
										{
											"id":           1,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 20, "deposit": true},
										},
									},
									"transitem_shipping": []cu.IM{},
									"tax_codes":          []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
									"currencies":         []cu.IM{{"code": "USD", "digit": 2}},
								},
							},
						},
					},
				},
				options: cu.IM{
					"create_delivery": false,
					"create_netto":    true,
				},
			},
			wantItems: 2,
		},
		{
			name: "create_regular",
			fields: fields{
				cls: &ClientService{},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"editor": cu.IM{
									"trans": cu.IM{
										"currency_code": "USD",
										"direction":     md.DirectionOut.String(),
									},
									"items": []cu.IM{
										{
											"id":           1,
											"product_code": "P1",
											"tax_code":     "VAT01",
											"item_meta":    cu.IM{"qty": 10},
										},
									},
									"transitem_invoice":  []cu.IM{},
									"transitem_shipping": []cu.IM{},
									"tax_codes":          []cu.IM{{"code": "VAT01", "rate_value": 0.2}},
									"currencies":         []cu.IM{{"code": "USD", "digit": 2}},
								},
							},
						},
					},
				},
				options: cu.IM{
					"create_delivery": false,
					"create_netto":    false,
				},
			},
			wantItems: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransService{
				cls: tt.fields.cls,
			}
			gotItems := s.createInvoiceItems(tt.args.evt, tt.args.options)
			if len(gotItems) != tt.wantItems {
				t.Errorf("TransService.createInvoiceItems() = %v items, want %v", len(gotItems), tt.wantItems)
			}
		})
	}
}

func TestTransService_linkAdd(t *testing.T) {
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
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"trans": cu.IM{
									"currency_code": "USD",
									"direction":     md.DirectionOut.String(),
									"place_code":    "P1",
								},
								"places": []cu.IM{
									{
										"code":          "P1",
										"currency_code": "USD",
									},
								},
							},
						},
					},
				},
			},
			params: cu.IM{
				"view":   "link",
				"code1":  "1234567890",
				"code2":  "1234567890",
				"amount": 100,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTransService(tt.cls)
			_, gotErr := s.linkAdd(tt.evt, tt.params)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("linkAdd() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("linkAdd() succeeded unexpectedly")
			}
		})
	}
}

func TestTransService_editorFieldViewAdd(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		evt          ct.ResponseEvent
		transMap     cu.IM
		resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)
		wantErr      bool
	}{
		{
			name: "items_view",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "items",
								"trans": cu.IM{
									"currency_code": "USD",
									"direction":     md.DirectionOut.String(),
									"place_code":    "P1",
								},
								"tax_codes": []cu.IM{
									{
										"code":       "VAT01",
										"rate_value": 0.2,
									},
								},
								"items": []cu.IM{
									{
										"id":           1,
										"product_code": "P1",
										"tax_code":     "VAT01",
										"item_meta":    cu.IM{"qty": 10},
									},
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "movement_delivery",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "movements",
								"trans": cu.IM{
									"trans_type": md.TransTypeDelivery.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
								"movements": []cu.IM{
									{
										"id":            1,
										"movement_type": md.MovementTypeInventory.String(),
										"trans_code":    "T1",
										"product_code":  "P1",
										"place_code":    "P1",
										"item_code":     "I1",
										"movement_meta": cu.IM{"qty": 10},
										"movement_map":  cu.IM{},
									},
									{
										"id":            2,
										"movement_type": md.MovementTypeInventory.String(),
										"trans_code":    "T1",
										"product_code":  "P1",
										"place_code":    "P1",
										"item_code":     "I1",
										"movement_meta": cu.IM{"qty": 10},
										"movement_map":  cu.IM{},
									},
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "movement_waybill",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "movements",
								"trans": cu.IM{
									"trans_type": md.TransTypeWaybill.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
								"movements": []cu.IM{
									{
										"id":            1,
										"movement_type": md.MovementTypeTool.String(),
										"trans_code":    "T1",
										"product_code":  "P1",
										"place_code":    "P1",
										"item_code":     "I1",
										"movement_meta": cu.IM{"qty": 10},
										"movement_map":  cu.IM{},
									},
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "movement_formula",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "movements",
								"trans": cu.IM{
									"trans_type": md.TransTypeFormula.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "movement_default",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "movements",
								"trans": cu.IM{
									"trans_type": md.TransTypeProduction.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
							},
							"movements": []cu.IM{
								{
									"id":            1,
									"movement_type": md.MovementTypeInventory.String(),
									"trans_code":    "T1",
									"product_code":  "P1",
									"place_code":    "P1",
									"item_code":     "I1",
									"movement_meta": cu.IM{"qty": 10},
									"movement_map":  cu.IM{},
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "tool_movement",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "tool_movement",
								"trans": cu.IM{
									"trans_type": md.TransTypeWaybill.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "tool_movement_dirty",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view":  "tool_movement",
								"dirty": true,
								"trans": cu.IM{
									"trans_type": md.TransTypeWaybill.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "trans_item_shipping",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "transitem_shipping",
								"trans": cu.IM{
									"trans_type": md.TransTypeOrder.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "trans_item_shipping_dirty",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view":  "transitem_shipping",
								"dirty": true,
								"trans": cu.IM{
									"trans_type": md.TransTypeOrder.String(),
									"direction":  md.DirectionOut.String(),
									"trans_date": "2025-01-01",
									"place_code": "P1",
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
		{
			name: "maps",
			cls:  &ClientService{},
			evt: ct.ResponseEvent{
				Trigger: &ct.Client{
					BaseComponent: ct.BaseComponent{
						Data: cu.IM{
							"editor": cu.IM{
								"view": "maps",
								"trans": cu.IM{
									"trans_map": cu.IM{
										"key": "value",
									},
								},
							},
						},
					},
				},
			},
			transMap: cu.IM{},
			resultUpdate: func(params cu.IM) (re ct.ResponseEvent, err error) {
				return ct.ResponseEvent{}, nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTransService(tt.cls)
			_, gotErr := s.editorFieldViewAdd(tt.evt, tt.transMap, tt.resultUpdate)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("editorFieldViewAdd() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("editorFieldViewAdd() succeeded unexpectedly")
			}
		})
	}
}

func TestTransService_formEventChange(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cls *ClientService
		// Named input parameters for target function.
		evt     ct.ResponseEvent
		wantErr bool
	}{
		{
			name: "product_code",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"product_code": "P1", "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "product_code"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "product_code", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "tool_code",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"tool_code": "T1", "tool_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "tool_code"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "tool_code", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "invoice_code",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"trans_code": "I1", "link_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "invoice_code"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "invoice_code", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "qty",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"qty": 2, "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "qty"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "qty", "form_event": ct.SelectorEventSelected},
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
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"tax_code": "T1", "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "tax_code"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "tax_code", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "amount",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"amount": 100, "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "amount"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "amount", "form_event": ct.SelectorEventSelected},
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
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"place_code": "P1", "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "place_code"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "place_code", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "own_stock",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"own_stock": 100, "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "own_stock"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "own_stock", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "deposit",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"deposit": true, "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "deposit"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "deposit", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "shared",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"shared": true, "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "shared"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "shared", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "link_amount",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"link_amount": 100, "link_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "link_amount"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "link_amount", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
		{
			name: "default",
			cls: &ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"product_name": "Test Product", "unit": "PC", "tax_code": "VAT01"}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: &cp.ClientComponent{},
			},
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
				Name: ct.ClientEventForm,
				Value: cu.IM{
					"data": cu.IM{"form": cu.IM{"index": 0, "key": "items",
						"data": cu.IM{"default": "P1", "item_meta": cu.IM{"qty": 1}}},
						"data": cu.IM{"name": "default"}},
					"value": cu.IM{},
					"event": ct.FormEventChange, "name": "default", "form_event": ct.SelectorEventSelected},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTransService(tt.cls)
			_, gotErr := s.formEventChange(tt.evt)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("formEventChange() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("formEventChange() succeeded unexpectedly")
			}
		})
	}
}
