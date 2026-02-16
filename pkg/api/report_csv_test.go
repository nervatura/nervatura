package api

import (
	"testing"

	cu "github.com/nervatura/component/pkg/util"
)

func TestSetReportWhere(t *testing.T) {
	type args struct {
		reportTemplate cu.IM
		filters        cu.IM
		sources        []cu.SM
		params         map[string][]any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "fields",
			args: args{
				reportTemplate: cu.IM{
					"fields": cu.IM{
						"date1": cu.IM{
							"fieldtype": "date",
							"wheretype": "where",
						},
						"string1": cu.IM{
							"fieldtype": "string",
							"wheretype": "where",
							"sqlstr":    "select * from table where string1 = @string1",
						},
						"date2": cu.IM{
							"fieldtype": "date",
							"wheretype": "where",
							"dataset":   "dsdate",
						},
						"string2": cu.IM{
							"fieldtype": "string",
							"wheretype": "in",
						},
						"string3": cu.IM{
							"fieldtype": "string",
							"wheretype": "in",
							"sqlstr":    "select * from table where string1 = @string1",
						},
					},
				},
				filters: cu.IM{
					"date1":   "2021-12-24",
					"string1": "value",
					"date2":   "2021-12-24",
					"string2": "value",
					"string3": "value",
				},
				sources: []cu.SM{
					{
						"dataset": "dsdate",
						"sqlstr":  "select * from table",
					},
				},
				params: map[string][]any{
					"dsdate": make([]any, 0),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetReportWhere(tt.args.reportTemplate, tt.args.filters, tt.args.sources, tt.args.params)
		})
	}
}
