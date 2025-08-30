package server

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"
	"strings"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	cli "github.com/nervatura/nervatura/v6/pkg/service/cli"
)

type cliHost struct {
	config cu.IM
	appLog *slog.Logger
	srv    *cli.CLIService
	args   cu.SM
	result string
}

func init() {
	registerHost("cli", &cliHost{})
}

func (h *cliHost) StartServer(config cu.IM, appLogOut, httpLogOut io.Writer, interrupt chan os.Signal) error {
	h.config = config
	h.appLog = slog.New(slog.NewJSONHandler(appLogOut, nil))
	appArgs := cu.ToSM(config["args"], cu.SM{})
	h.args = cu.ToSM(config["args"], cu.SM{})

	h.srv = cli.NewCLIService(h.config, h.appLog)

	if len(appArgs) == 0 {
		h.parseFlags()
	}

	if h.args["cmd"] == "server" {
		h.result = "server"
		return nil
	}

	var err error
	if err = h.checkRequired(); err != nil {
		return err
	}

	if h.result, err = h.parseCommand(); err != nil {
		return err
	}
	if len(appArgs) == 0 {
		fmt.Println(h.result)
	}

	return nil
}

func (h *cliHost) StopServer(ctx context.Context) error {
	return nil
}

func (h *cliHost) Results() string {
	return h.result
}

func (h *cliHost) parseFlags() {
	cmds := []string{"server", "create", "update", "delete", "get", "query", "database", "function", "reset", "view", "upgrade"}
	models := []string{"auth", "config", "currency", "customer", "employee", "item", "link", "log",
		"movement", "payment", "place", "price", "product", "project", "rate", "tax", "tool", "trans"}

	var help bool
	flag.BoolVar(&help, "help", false, "Program usage")
	flag.BoolVar(&help, "h", false, "Program usage")
	var tray bool
	flag.BoolVar(&tray, "tray", false, "Show the system tray icon and menu. Disabled in the Docker container!")
	var env string
	flag.StringVar(&env, "env", "", "Application configuration .env file. Examples:\n-env sample.env \n-env /home/user/conf/.env")
	var cmd string
	flag.StringVar(&cmd, "c", "server", "Available commands: "+strings.Join(cmds, ", ")+"\n")
	var model string
	flag.StringVar(&model, "m", "", "Available models:\n"+strings.Join(models, ", ")+"\n")
	var options string
	flag.StringVar(&options, "o", "", "Options JSON Object string. Required for the following commands:\n"+strings.Join(cmds[1:], ", ")+",\n")
	var data string
	flag.StringVar(&data, "d", "", "Data JSON Object string. Required for the following commands: create, update\n")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	h.args["cmd"] = cmd
	if help {
		h.args["cmd"] = "help"
	}
	if model != "" {
		h.args["model"] = model
	}
	if options != "" {
		h.args["options"] = options
	}
	if data != "" {
		h.args["data"] = data
	}
}

func (h *cliHost) checkRequired() (err error) {
	switch h.args["cmd"] {
	case "create", "update":
		if h.args["model"] == "" || h.args["options"] == "" || h.args["data"] == "" {
			return errors.New("required parameters are missing: -m, -o, -d")
		}
	case "get", "query":
		if h.args["model"] == "" || h.args["options"] == "" {
			return errors.New("required parameters are missing: -m, -o")
		}
	case "database", "function", "reset", "view", "upgrade":
		if h.args["options"] == "" {
			return errors.New("required parameters are missing: -o")
		}
	}
	return nil
}

func (h *cliHost) parseCommand() (result string, err error) {
	if h.args["cmd"] == "help" {
		flag.Usage()
		return "", nil
	}

	var options cu.IM
	if cu.ConvertFromByte([]byte(h.args["options"]), &options) != nil {
		return "", errors.New("failed to convert options")
	}
	h.srv.Config = cu.MergeIM(h.srv.Config, options)

	if cmd, found := h.cliMap(h.args["cmd"], h.args["model"]); found {
		options["ds"] = api.NewDataStore(h.srv.Config, cu.ToString(options["alias"], ""), h.appLog)
		return cmd(options, cu.ToString(h.args["data"], "")), nil
	}
	return "", errors.New("command or model not found: " + h.args["cmd"] + " (-c) " + h.args["model"] + " (-m)")
}

func (h *cliHost) cliMap(cmd, model string) (cmdFunc func(options cu.IM, requestData string) string, found bool) {
	cmdMap := map[string]map[string]func(options cu.IM, requestData string) string{
		"auth": {
			"create": h.srv.AuthInsert,
			"update": h.srv.AuthUpdate,
			"get":    h.srv.AuthGet,
			"query":  h.srv.AuthQuery,
		},
		"config": {
			"create": h.srv.ConfigInsert,
			"update": h.srv.ConfigUpdate,
			"delete": h.srv.ConfigDelete,
			"get":    h.srv.ConfigGet,
			"query":  h.srv.ConfigQuery,
		},
		"currency": {
			"create": h.srv.CurrencyInsert,
			"update": h.srv.CurrencyUpdate,
			"delete": h.srv.CurrencyDelete,
			"get":    h.srv.CurrencyGet,
			"query":  h.srv.CurrencyQuery,
		},
		"customer": {
			"create": h.srv.CustomerInsert,
			"update": h.srv.CustomerUpdate,
			"delete": h.srv.CustomerDelete,
			"get":    h.srv.CustomerGet,
			"query":  h.srv.CustomerQuery,
		},
		"employee": {
			"create": h.srv.EmployeeInsert,
			"update": h.srv.EmployeeUpdate,
			"delete": h.srv.EmployeeDelete,
			"get":    h.srv.EmployeeGet,
			"query":  h.srv.EmployeeQuery,
		},
		"item": {
			"create": h.srv.ItemInsert,
			"update": h.srv.ItemUpdate,
			"delete": h.srv.ItemDelete,
			"get":    h.srv.ItemGet,
			"query":  h.srv.ItemQuery,
		},
		"link": {
			"create": h.srv.LinkInsert,
			"update": h.srv.LinkUpdate,
			"delete": h.srv.LinkDelete,
			"get":    h.srv.LinkGet,
			"query":  h.srv.LinkQuery,
		},
		"log": {
			"query": h.srv.LogQuery,
			"get":   h.srv.LogGet,
		},
		"movement": {
			"create": h.srv.MovementInsert,
			"update": h.srv.MovementUpdate,
			"delete": h.srv.MovementDelete,
			"get":    h.srv.MovementGet,
			"query":  h.srv.MovementQuery,
		},
		"payment": {
			"create": h.srv.PaymentInsert,
			"update": h.srv.PaymentUpdate,
			"delete": h.srv.PaymentDelete,
			"get":    h.srv.PaymentGet,
			"query":  h.srv.PaymentQuery,
		},
		"place": {
			"create": h.srv.PlaceInsert,
			"update": h.srv.PlaceUpdate,
			"delete": h.srv.PlaceDelete,
			"get":    h.srv.PlaceGet,
			"query":  h.srv.PlaceQuery,
		},
		"price": {
			"create": h.srv.PriceInsert,
			"update": h.srv.PriceUpdate,
			"delete": h.srv.PriceDelete,
			"get":    h.srv.PriceGet,
			"query":  h.srv.PriceQuery,
		},
		"product": {
			"create": h.srv.ProductInsert,
			"update": h.srv.ProductUpdate,
			"delete": h.srv.ProductDelete,
			"get":    h.srv.ProductGet,
			"query":  h.srv.ProductQuery,
		},
		"project": {
			"create": h.srv.ProjectInsert,
			"update": h.srv.ProjectUpdate,
			"delete": h.srv.ProjectDelete,
			"get":    h.srv.ProjectGet,
			"query":  h.srv.ProjectQuery,
		},
		"rate": {
			"create": h.srv.RateInsert,
			"update": h.srv.RateUpdate,
			"delete": h.srv.RateDelete,
			"get":    h.srv.RateGet,
			"query":  h.srv.RateQuery,
		},
		"tax": {
			"create": h.srv.TaxInsert,
			"update": h.srv.TaxUpdate,
			"delete": h.srv.TaxDelete,
			"get":    h.srv.TaxGet,
			"query":  h.srv.TaxQuery,
		},
		"tool": {
			"create": h.srv.ToolInsert,
			"update": h.srv.ToolUpdate,
			"delete": h.srv.ToolDelete,
			"get":    h.srv.ToolGet,
			"query":  h.srv.ToolQuery,
		},
		"trans": {
			"create": h.srv.TransInsert,
			"update": h.srv.TransUpdate,
			"delete": h.srv.TransDelete,
			"get":    h.srv.TransGet,
			"query":  h.srv.TransQuery,
		},
		"reset": {
			"reset": h.srv.ResetPassword,
		},
		"database": {
			"database": h.srv.Database,
		},
		"function": {
			"function": h.srv.Function,
		},
		"view": {
			"view": h.srv.View,
		},
		"upgrade": {
			"upgrade": h.srv.Upgrade,
		},
	}
	if slices.Contains([]string{"database", "function", "reset", "view", "upgrade"}, cmd) {
		model = cmd
	}
	cmdFunc, found = cmdMap[model][cmd]
	return
}
