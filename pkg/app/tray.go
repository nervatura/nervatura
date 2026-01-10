package app

import (
	"context"
	"os"

	"fyne.io/systray"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/static/icon"
)

// systemTray implements the Nervatura system tray icon and menu
type systemTray struct {
	app          *App
	interrupt    chan os.Signal
	ctx          context.Context
	httpDisabled bool
}

// onReady - initialize the system tray icon and menu
func (st *systemTray) onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Nervatura")
	systray.SetTooltip("Nervatura " + cu.ToString(st.app.config["version"], ""))
	mnuConfig := systray.AddMenuItem("Configuration values", "Configuration values")
	mnuAdmin := systray.AddMenuItem("Open Admin GUI", "Open Admin GUI")
	if st.httpDisabled {
		mnuConfig.Disable()
		mnuAdmin.Disable()
	}
	systray.AddSeparator()
	mnuExit := systray.AddMenuItem("Shut down and quit", "Shut down and quit")
	/*
		go func() {
			for range mnuExit.ClickedCh {
				systray.Quit()
			}
		}()
	*/
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
