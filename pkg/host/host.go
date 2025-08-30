package server

import (
	"context"
	"io"
	"os"

	cu "github.com/nervatura/component/pkg/util"
)

// APIHost Nervatura API interface
type APIHost interface {
	StartServer(config cu.IM, appLogOut, httpLogOut io.Writer, interrupt chan os.Signal) error
	Results() string
	StopServer(ctx context.Context) error
}

var Hosts = make(map[string]APIHost)

func registerHost(name string, server APIHost) {
	Hosts[name] = server
}
