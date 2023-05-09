package service

import (
	"errors"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

// CLIService implements the Nervatura API service
type CLIService struct {
	Config        map[string]interface{}
	GetNervaStore func(database string) *nt.NervaStore
}

func respondData(code int, data interface{}, errCode int, err error) string {
	if err != nil {
		data = nt.IM{"code": errCode, "message": err.Error()}
	}
	if data == nil {
		data = nt.IM{"code": code, "message": "OK"}
	}
	jdata, _ := ut.ConvertToByte(data)
	return string(jdata)
}

func (srv *CLIService) TokenDecode(token string) string {
	claim, err := ut.TokenDecode(token)
	return respondData(200, claim, 400, err)
}

func (srv *CLIService) TokenLogin(token string, tokenKeys map[string]map[string]string) (*nt.API, string) {
	claim, err := ut.TokenDecode(token)
	if err != nil {
		return nil, respondData(0, nil, 401, errors.New(ut.GetMessage("error_unauthorized")))
	}
	database := ut.ToString(claim["database"], "")
	nstore := srv.GetNervaStore(database)
	api := (&nt.API{NStore: nstore})
	err = api.TokenLogin(nt.IM{"token": token, "keys": tokenKeys})
	return api, respondData(200, api.NStore.User, 401, err)
}

func (srv *CLIService) UserLogin(options nt.IM) string {
	if _, found := options["database"]; !found {
		return respondData(0, nil, 400, errors.New(ut.GetMessage("missing_database")))
	}
	nstore := srv.GetNervaStore(options["database"].(string))
	token, engine, err := (&nt.API{NStore: nstore}).UserLogin(options)
	return respondData(200, nt.SM{"token": token, "engine": engine, "version": srv.Config["version"].(string)}, 400, err)
}

func (srv *CLIService) checkUser(api *nt.API, admin bool) error {
	if api == nil {
		return errors.New(ut.GetMessage("error_unauthorized"))
	}
	if api.NStore.User == nil {
		return errors.New(ut.GetMessage("error_unauthorized"))
	}
	if admin && api.NStore.User.Scope != "admin" {
		return errors.New(ut.GetMessage("error_unauthorized"))
	}
	return nil
}

func (srv *CLIService) UserPassword(api *nt.API, options nt.IM) string {
	username := ut.ToString(options["username"], "")
	custnumber := ut.ToString(options["custnumber"], "")

	if err := srv.checkUser(api, ((username != "" && api.NStore.User != nil && username != api.NStore.User.Username) || custnumber != "")); err != nil {
		return respondData(0, nil, 401, err)
	}
	if username == "" && custnumber == "" {
		options["username"] = api.NStore.User.Username
	}
	err := api.UserPassword(options)
	return respondData(204, nil, 400, err)
}

func (srv *CLIService) TokenRefresh(api *nt.API) string {
	if api == nil {
		return respondData(0, nil, 401, errors.New(ut.GetMessage("error_unauthorized")))
	}
	token, err := api.TokenRefresh()
	return respondData(200, nt.SM{"token": token}, 400, err)
}

func (srv *CLIService) Get(api *nt.API, options nt.IM) string {
	if err := srv.checkUser(api, false); err != nil {
		return respondData(0, nil, 401, err)
	}
	results, err := api.Get(options)
	return respondData(200, results, 400, err)
}

func (srv *CLIService) View(api *nt.API, data []nt.IM) string {
	if err := srv.checkUser(api, false); err != nil {
		return respondData(0, nil, 401, err)
	}
	results, err := api.View(data)
	return respondData(200, results, 400, err)
}

func (srv *CLIService) Function(api *nt.API, options nt.IM) string {
	if err := srv.checkUser(api, false); err != nil {
		return respondData(0, nil, 401, err)
	}
	results, err := api.Function(options)
	return respondData(200, results, 400, err)
}

func (srv *CLIService) Update(api *nt.API, nervatype string, data []nt.IM) string {
	if err := srv.checkUser(api, false); err != nil {
		return respondData(0, nil, 401, err)
	}
	results, err := api.Update(nervatype, data)
	return respondData(200, results, 400, err)
}

func (srv *CLIService) Delete(api *nt.API, options nt.IM) string {
	if err := srv.checkUser(api, false); err != nil {
		return respondData(0, nil, 401, err)
	}
	err := api.Delete(options)
	return respondData(204, nil, 400, err)
}

func (srv *CLIService) DatabaseCreate(apiKey string, options nt.IM) string {
	if srv.Config["NT_API_KEY"] != apiKey {
		return respondData(0, nil, 401, errors.New(ut.GetMessage("error_unauthorized")))
	}
	nstore := srv.GetNervaStore(options["database"].(string))
	log, err := (&nt.API{NStore: nstore}).DatabaseCreate(options)
	return respondData(200, log, 400, err)
}

func (srv *CLIService) ReportList(api *nt.API, options nt.IM) string {
	if err := srv.checkUser(api, true); err != nil {
		return respondData(0, nil, 401, err)
	}
	results, err := api.ReportList(options)
	return respondData(200, results, 400, err)
}

func (srv *CLIService) ReportInstall(api *nt.API, options nt.IM) string {
	if err := srv.checkUser(api, true); err != nil {
		return respondData(0, nil, 401, err)
	}
	results, err := api.ReportInstall(options)
	return respondData(200, results, 400, err)
}

func (srv *CLIService) ReportDelete(api *nt.API, options nt.IM) string {
	if err := srv.checkUser(api, true); err != nil {
		return respondData(0, nil, 401, err)
	}
	err := api.ReportDelete(options)
	return respondData(204, nil, 400, err)
}

func (srv *CLIService) Report(api *nt.API, options nt.IM) string {
	if err := srv.checkUser(api, false); err != nil {
		return respondData(0, nil, 401, err)
	}
	if _, found := options["output"]; !found || (options["output"] != "xml") {
		options["output"] = "base64"
	}
	results, err := api.Report(options)
	return respondData(200, results, 400, err)
}
