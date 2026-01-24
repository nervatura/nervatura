package app

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"fyne.io/systray"
	cu "github.com/nervatura/component/pkg/util"
)

func Test_systemTray_onMenuAdmin(t *testing.T) {
	type fields struct {
		app          *App
		interrupt    chan os.Signal
		ctx          context.Context
		httpDisabled bool
		mnuConfig    *systray.MenuItem
		mnuAdmin     *systray.MenuItem
		mnuExit      *systray.MenuItem
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				app: &App{
					config: cu.IM{},
					appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				},
				mnuAdmin:  systray.AddMenuItem("Open Admin GUI", "Open Admin GUI"),
				mnuExit:   systray.AddMenuItem("Shut down and quit", "Shut down and quit"),
				mnuConfig: systray.AddMenuItem("Configuration values", "Configuration values"),
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
				mnuConfig:    tt.fields.mnuConfig,
				mnuAdmin:     tt.fields.mnuAdmin,
				mnuExit:      tt.fields.mnuExit,
			}
			go st.onMenuAdmin()
			go st.onMenuExit()
			go st.onMenuConfig()
			time.Sleep(1 * time.Second)
			tt.fields.mnuAdmin.ClickedCh <- struct{}{}
			tt.fields.mnuExit.ClickedCh <- struct{}{}
			tt.fields.mnuConfig.ClickedCh <- struct{}{}
		})
	}
}

func Test_systemTray_onInterruptOrDone(t *testing.T) {
	type fields struct {
		app          *App
		interrupt    chan os.Signal
		ctx          context.Context
		httpDisabled bool
		mnuConfig    *systray.MenuItem
		mnuAdmin     *systray.MenuItem
		mnuExit      *systray.MenuItem
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				app: &App{
					config: cu.IM{},
					appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				},
				interrupt: make(chan os.Signal),
				ctx:       context.Background(),
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
				mnuConfig:    tt.fields.mnuConfig,
				mnuAdmin:     tt.fields.mnuAdmin,
				mnuExit:      tt.fields.mnuExit,
			}
			go st.onInterruptOrDone()
			time.Sleep(1 * time.Second)
			tt.fields.interrupt <- os.Interrupt
		})
	}
}
