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
							"sqlstr":    "string1",
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
							"sqlstr":    "string1",
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetReportWhere(tt.args.reportTemplate, tt.args.filters, tt.args.sources)
		})
	}
}
