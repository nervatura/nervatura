//go:build linux || windows
// +build linux windows

package app

import (
	"context"
	"os"

	"fyne.io/systray"
	"github.com/nervatura/nervatura/service/pkg/icon"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

// systemTray implements the Nervatura system tray icon and menu
type systemTray struct {
	app          *App
	interrupt    chan os.Signal
	ctx          context.Context
	httpDisabled bool
}

func init() {
	traySrv = &systemTray{}
}

func (st *systemTray) onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Nervatura")
	systray.SetTooltip("Nervatura " + ut.ToString(st.app.config["version"], ""))
	mnuConfig := systray.AddMenuItem(ut.GetMessage("view_configuration"), ut.GetMessage("view_configuration"))
	mnuAdmin := systray.AddMenuItem(ut.GetMessage("task_admin"), ut.GetMessage("task_admin"))
	if st.httpDisabled {
		mnuConfig.Disable()
		mnuAdmin.Disable()
	}
	systray.AddSeparator()
	mnuExit := systray.AddMenuItem(ut.GetMessage("task_exit"), ut.GetMessage("task_exit"))
	go func() {
		for {
			select {
			case <-mnuConfig.ClickedCh:
				st.app.onTrayMenu("config")

			case <-mnuAdmin.ClickedCh:
				st.app.onTrayMenu("admin")

			case <-mnuExit.ClickedCh:
				systray.Quit()

			case <-st.interrupt:
				systray.Quit()

			case <-st.ctx.Done():
				systray.Quit()
			}
		}
	}()
}

// Run initializes GUI and starts the event loop, then invokes the onReady callback.
func (st *systemTray) Run(app *App, interrupt chan os.Signal, ctx context.Context, httpDisabled bool, onExit func()) {
	st.app = app
	st.interrupt = interrupt
	st.ctx = ctx
	st.httpDisabled = httpDisabled
	systray.Run(st.onReady, onExit)
}
