package component

import (
	"html/template"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
)

func TestQueueEditor_CustomTemplateCell(t *testing.T) {
	type args struct {
		sessionID string
		label     string
		inline    string
	}
	tests := []struct {
		name string
		s    *QueueEditor
		args args
		want func(row cu.IM, col ct.TableColumn, value any, rowIndex int64) template.HTML
	}{
		{
			name: "custom_template_cell",
			s:    &QueueEditor{},
			args: args{
				sessionID: "123",
				label:     "test",
				inline:    "true",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &QueueEditor{}
			got := s.CustomTemplateCell(tt.args.sessionID, tt.args.label, tt.args.inline)
			got(cu.IM{"code": "123"}, ct.TableColumn{}, "123", 0, nil)
		})
	}
}
