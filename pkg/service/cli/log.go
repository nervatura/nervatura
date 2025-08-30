package cli

import (
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
)

// LogQuery - get logs
func (cli *CLIService) LogQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "log"}
	for _, v := range []string{"log_type", "ref_type", "ref_code", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}

	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// LogGet - get log by id or code
func (cli *CLIService) LogGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	logID := cu.ToInteger(options["id"], 0)
	logCode := cu.ToString(options["code"], "")
	var err error
	var logs []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if logs, err = ds.GetDataByID("log", logID, logCode, true); err == nil {
		response = logs[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
