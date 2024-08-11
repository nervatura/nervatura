//go:build linux || windows
// +build linux windows

package app

/*
func Test_systemTray_onReady(t *testing.T) {
	type fields struct {
		app          *App
		interrupt    chan os.Signal
		ctx          context.Context
		httpDisabled bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "disabled",
			fields: fields{
				app:          &App{},
				ctx:          context.Background(),
				httpDisabled: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &systemTray{
				app:          tt.fields.app,
				interrupt:    tt.fields.interrupt,
				ctx:          tt.fields.ctx,
				httpDisabled: tt.fields.httpDisabled,
			}
			st.onReady()
		})
	}
}
*/
