package app

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	srv "github.com/nervatura/nervatura/service/pkg/service"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

type cliServer struct {
	app     *App
	service srv.CLIService
	args    nt.SM
	result  string
}

func init() {
	registerService("cli", &cliServer{})
}

// StartService - Start Nervatura CLI server
func (s *cliServer) StartService() error {
	s.service = srv.CLIService{
		Config:        s.app.config,
		GetNervaStore: s.app.GetNervaStore,
	}
	s.args = s.app.args
	if s.app.args == nil {
		s.parseFlags()
	}

	if s.args["cmd"] == "server" {
		s.result = "server"
		return nil
	}

	err := s.checkRequired()
	if err != nil {
		if s.app.args == nil {
			flag.Usage()
		}
		return err
	}

	s.result, err = s.parseCommand()
	if err != nil {
		return err
	}
	if s.app.args == nil {
		fmt.Println(s.result)
	}
	return nil
}

func (s *cliServer) Results() string {
	return s.result
}

func (s *cliServer) ConnectApp(app interface{}) {
	s.app = app.(*App)
}

func (s *cliServer) StopService(interface{}) error {
	return nil
}

func (s *cliServer) parseFlags() {
	s.args = make(nt.SM)
	var cmds = []string{"server", "Delete", "Function", "Get", "Update", "View",
		"UserPassword", "TokenLogin", "TokenRefresh", "UserLogin", "DatabaseCreate",
		"Report", "ReportDelete", "ReportInstall", "ReportList", "TokenDecode"}
	var cmdsO = []string{"Delete", "Function", "Get", "Update", "UserPassword",
		"UserLogin", "DatabaseCreate", "Report", "ReportDelete", "ReportInstall", "ReportList"}
	var cmdsNT = []string{"Update"}
	var cmdsD = []string{"Update", "View"}
	var cmdsK = []string{"DatabaseCreate"}

	var help bool
	flag.BoolVar(&help, "help", false, ut.GetMessage("cli_usage"))
	var env string
	flag.StringVar(&env, "env", "", ut.GetMessage("cli_flag_env"))
	var cmd string
	flag.StringVar(&cmd, "c", "server", ut.GetMessage("cli_flag_c")+strings.Join(cmds[:10], ", ")+",\n"+strings.Join(cmds[10:], ", "))
	var token string
	flag.StringVar(&token, "t", "", ut.GetMessage("cli_flag_t"))
	var options string
	flag.StringVar(&options, "o", "", ut.GetMessage("cli_flag_o")+strings.Join(cmdsO[:8], ", ")+",\n"+strings.Join(cmdsO[8:], ", "))
	var ntype string
	flag.StringVar(&ntype, "nt", "", ut.GetMessage("cli_flag_nt")+strings.Join(cmdsNT, ", "))
	var data string
	flag.StringVar(&data, "d", "", ut.GetMessage("cli_flag_d")+strings.Join(cmdsD, ", "))
	var key string
	flag.StringVar(&key, "k", "", ut.GetMessage("cli_flag_k")+strings.Join(cmdsK, ", "))

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	s.args["cmd"] = cmd
	if help {
		s.args["cmd"] = "help"
	}
	if token != "" {
		s.args["token"] = token
	}
	if options != "" {
		s.args["options"] = options
	}
	if data != "" {
		s.args["data"] = data
	}
	if ntype != "" {
		s.args["nervatype"] = ntype
	}
	if key != "" {
		s.args["key"] = key
	}
}

func (s *cliServer) checkRequired() (err error) {
	if value, found := s.args["token"]; (!found || value == "") &&
		s.args["cmd"] != "UserLogin" && s.args["cmd"] != "DatabaseCreate" && s.args["cmd"] != "help" {
		return errors.New(ut.GetMessage("missing_parameter") + ": token(-t)")
	}
	switch s.args["cmd"] {
	case "Delete", "Function", "Get", "UserPassword",
		"UserLogin", "Report", "ReportDelete", "ReportInstall", "ReportList":
		if value, found := s.args["options"]; !found || value == "" {
			return errors.New(ut.GetMessage("missing_parameter") + ": options(-o)")
		}

	case "DatabaseCreate":
		if value, found := s.args["options"]; !found || value == "" {
			return errors.New(ut.GetMessage("missing_parameter") + ": options(-o)")
		}
		if value, found := s.args["key"]; !found || value == "" {
			return errors.New(ut.GetMessage("missing_parameter") + ": API key(-k)")
		}

	case "View":
		if value, found := s.args["data"]; !found || value == "" {
			return errors.New(ut.GetMessage("missing_parameter") + ": data(-d)")
		}

	case "Update":
		if value, found := s.args["nervatype"]; !found || value == "" {
			return errors.New(ut.GetMessage("missing_parameter") + ": nervatype(-nt)")
		}
		if value, found := s.args["data"]; !found || value == "" {
			return errors.New(ut.GetMessage("missing_parameter") + ": data(-d)")
		}

	}
	return nil
}

// parseCommand - Parse s.args from command line parameters
func (s *cliServer) parseCommand() (result string, err error) {
	if s.args["cmd"] == "help" {
		flag.Usage()
		return "", nil
	}

	var api *nt.API
	if _, found := s.args["token"]; found {
		if s.args["cmd"] == "TokenDecode" {
			return s.service.TokenDecode(s.args["token"]), nil
		}
		api, result = s.service.TokenLogin(s.args["token"], s.app.tokenKeys)
		if api == nil || s.args["cmd"] == "TokenLogin" {
			return result, nil
		}
	}

	var options nt.IM
	if _, found := s.args["options"]; found {
		if err = ut.ConvertFromByte([]byte(s.args["options"]), &options); err != nil {
			return "", errors.New(ut.GetMessage("invalid_json"))
		}
	}

	var data []nt.IM
	if _, found := s.args["data"]; found {
		if err = ut.ConvertFromByte([]byte(s.args["data"]), &data); err != nil {
			return "", errors.New(ut.GetMessage("invalid_json"))
		}
	}

	apiMap := map[string]func(api *nt.API) string{
		"UserLogin": func(api *nt.API) string {
			return s.service.UserLogin(options)
		},
		"UserPassword": func(api *nt.API) string {
			return s.service.UserPassword(api, options)
		},
		"TokenRefresh": func(api *nt.API) string {
			return s.service.TokenRefresh(api)
		},
		"Get": func(api *nt.API) string {
			return s.service.Get(api, options)
		},
		"View": func(api *nt.API) string {
			return s.service.View(api, data)
		},
		"Function": func(api *nt.API) string {
			return s.service.Function(api, options)
		},
		"Update": func(api *nt.API) string {
			return s.service.Update(api, s.args["nervatype"], data)
		},
		"Delete": func(api *nt.API) string {
			return s.service.Delete(api, options)
		},
		"DatabaseCreate": func(api *nt.API) string {
			return s.service.DatabaseCreate(s.args["key"], options)
		},
		"Report": func(api *nt.API) string {
			return s.service.Report(api, options)
		},
		"ReportList": func(api *nt.API) string {
			return s.service.ReportList(api, options)
		},
		"ReportInstall": func(api *nt.API) string {
			return s.service.ReportInstall(api, options)
		},
		"ReportDelete": func(api *nt.API) string {
			return s.service.ReportDelete(api, options)
		},
	}
	if _, found := apiMap[s.args["cmd"]]; !found {
		return "", errors.New(ut.GetMessage("invalid_command") + ": " + s.args["cmd"] + " (-c)")
	}
	return apiMap[s.args["cmd"]](api), nil
}
