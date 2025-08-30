package cli

import (
	"log/slog"
	"net/http"
	"strings"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

// CLIService implements the Nervatura API service
type CLIService struct {
	Config map[string]interface{}
}

func NewCLIService(config cu.IM, appLog *slog.Logger) *CLIService {
	return &CLIService{Config: config}
}

func (cli *CLIService) respondString(code int, data interface{}, errCode int, err error) string {
	result := cu.IM{"code": code, "data": http.StatusText(code)}
	if err != nil {
		result["code"] = errCode
		result["data"] = err.Error()
	}
	if err == nil && data != nil {
		result["data"] = data
	}

	jdata, _ := cu.ConvertToByte(result)
	return string(jdata)
}

func (cli *CLIService) ResetPassword(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	authCode := cu.ToString(options["code"], "")

	password := cu.RandString(16)
	err := ds.UserPassword(authCode, password, password)
	return cli.respondString(http.StatusOK, cu.IM{"password": password}, http.StatusUnprocessableEntity, err)
}

// Database - create a new database schema
func (cli *CLIService) Database(options cu.IM, requestData string) string {
	response := api.CreateDatabase(options, cli.Config)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, nil)
}

// Upgrade - upgrade the database schema
func (cli *CLIService) Upgrade(options cu.IM, requestData string) string {
	response := api.UpgradeDatabase(options, cli.Config)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, nil)
}

// Call a server side function and get the result
// Example: create new PDF invoice, send email or get a product price
func (cli *CLIService) Function(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	response, err := ds.Function(cu.ToString(options["name"], ""), options)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// View - get a predefined view
func (cli *CLIService) View(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var response []cu.IM
	var err error
	query := md.Query{
		Fields:  []string{"*"},
		From:    strings.ToLower(strings.TrimPrefix(cu.ToString(options["name"], ""), "VIEW_")),
		Filter:  cu.ToString(options["filter"], ""),
		OrderBy: strings.Split(cu.ToString(options["order_by"], ""), ","),
		Limit:   cu.ToInteger(options["limit"], 0),
		Offset:  cu.ToInteger(options["offset"], 0),
	}
	response, err = ds.StoreDataQuery(query, false)

	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}
