package api

import (
	"path"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func TestCreateReportPDF(t *testing.T) {
	type args struct {
		options      cu.IM
		datarows     cu.IM
		config       cu.IM
		jsonTemplate string
	}
	sample_json, _ := st.Report.ReadFile(path.Join("template", "sample.json"))
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base64",
			args: args{
				options: cu.IM{
					"output": "base64",
				},
				datarows: cu.IM{
					"labels": cu.IM{
						"label": "value",
					},
				},
				jsonTemplate: string(sample_json),
			},
			wantErr: false,
		},
		{
			name: "xml",
			args: args{
				options: cu.IM{
					"output": "xml",
				},
				datarows: cu.IM{
					"labels": []cu.IM{
						{"label": "value"},
					},
				},
				jsonTemplate: string(sample_json),
			},
			wantErr: false,
		},
		{
			name: "set_data_error",
			args: args{
				options: cu.IM{
					"output": "pdf",
				},
				datarows: cu.IM{
					"error": []interface{}{
						[]byte("value"),
					},
				},
				jsonTemplate: string(sample_json),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateReportPDF(tt.args.options, tt.args.datarows, tt.args.config, tt.args.jsonTemplate)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateReportPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
