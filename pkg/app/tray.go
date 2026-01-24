package app

import (
	"context"
	"os"

	"fyne.io/systray"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/static/icon"
)

// systemTray implements the Nervatura system tray icon and menu
type systemTray struct {
	app          *App
	interrupt    chan os.Signal
	ctx          context.Context
	httpDisabled bool
	mnuConfig    *systray.MenuItem
	mnuAdmin     *systray.MenuItem
	mnuExit      *systray.MenuItem
}

func (st *systemTray) onMenuExit() {
	for range st.mnuExit.ClickedCh {
		systray.Quit()
	}
}

func (st *systemTray) onMenuConfig() {
	for range st.mnuConfig.ClickedCh {
		st.app.onTrayMenu("config")
	}
}

func (st *systemTray) onMenuAdmin() {
	for range st.mnuAdmin.ClickedCh {
		st.app.onTrayMenu("admin")
	}
}

func (st *systemTray) onInterruptOrDone() {
	for {
		select {
		case <-st.interrupt:
			systray.Quit()

		case <-st.ctx.Done():
			systray.Quit()
		}
	}
}

// onReady - initialize the system tray icon and menu
func (st *systemTray) onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Nervatura")
	systray.SetTooltip("Nervatura " + cu.ToString(st.app.config["version"], ""))
	if !st.httpDisabled {
		st.mnuConfig = systray.AddMenuItem("Configuration values", "Configuration values")
		go st.onMenuConfig()
		st.mnuAdmin = systray.AddMenuItem("Open Admin GUI", "Open Admin GUI")
		go st.onMenuAdmin()
	}
	systray.AddSeparator()
	st.mnuExit = systray.AddMenuItem("Shut down and quit", "Shut down and quit")
	go st.onMenuExit()

	go st.onInterruptOrDone()
}

// Run initializes GUI and starts the event loop, then invokes the onReady callback.
func (st *systemTray) Run(app *App, interrupt chan os.Signal, ctx context.Context, httpDisabled bool, onExit func()) {
	st.app = app
	st.interrupt = interrupt
	st.ctx = ctx
	st.httpDisabled = httpDisabled
	systray.Run(st.onReady, onExit)
}
