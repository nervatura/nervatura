package api

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func (ds *DataStore) ReportInstall(reportKey, reportDir string) (configID int64, err error) {
	reportDir = cu.ToString(reportDir, cu.ToString(ds.Config["NT_REPORT_DIR"], ""))
	var file []byte
	if reportDir == "" {
		file, err = st.Report.ReadFile(path.Join("template", reportKey+".json"))
	} else {
		file, err = ds.ReadFile(filepath.Clean(filepath.Join(reportDir, reportKey+".json")))
	}
	if err != nil {
		return configID, err
	}

	var temp cu.IM
	if err = ds.ConvertFromByte(file, &temp); err != nil {
		return configID, err
	}

	metaData := cu.ToIM(temp["meta"], cu.IM{})
	for _, v := range []string{"report_key", "report_name", "report_type", "file_type"} {
		if _, ok := metaData[v]; !ok {
			return configID, errors.New("invalid template")
		}
	}
	var configs []cu.IM
	if configs, err = ds.GetDataByID("config", 0, cu.ToString(metaData["report_key"], ""), false); err != nil || len(configs) > 0 {
		if len(configs) > 0 {
			err = errors.New("the template already exists")
		}
		return configID, err
	}

	var report md.ConfigReport = md.ConfigReport{
		ReportKey:   cu.ToString(metaData["report_key"], ""),
		ReportType:  cu.ToString(metaData["report_type"], ""),
		TransType:   cu.ToString(metaData["trans_type"], ""),
		Direction:   cu.ToString(metaData["direction"], ""),
		ReportName:  cu.ToString(metaData["report_name"], ""),
		Description: cu.ToString(metaData["description"], ""),
		Label:       cu.ToString(metaData["label"], ""),
		FileType:    md.FileType(0).Value(cu.ToString(metaData["file_type"], "")),
		Template:    string(file[:]),
	}
	var values cu.IM = cu.IM{
		"code":        report.ReportKey,
		"config_type": md.ConfigTypeReport.String(),
	}
	if configDataByte, err := ds.ConvertToByte(report); err == nil {
		values["data"] = string(configDataByte[:])
	}

	return ds.StoreDataUpdate(md.Update{Values: values, Model: "config"})
}

func (ds *DataStore) ReportList(reportDir, filter string) (results []cu.IM, err error) {

	var reports []cu.IM
	var installed cu.IM = cu.IM{}
	query := md.Query{Fields: []string{"*"}, From: "config_report"}
	reports, err = ds.StoreDataQuery(query, false)
	for _, report := range reports {
		installed[cu.ToString(report["report_key"], "")] = cu.ToInteger(report["id"], 0)
	}

	fileInfo := func(file []byte, fileName string) {
		temp := cu.IM{}
		if err = ds.ConvertFromByte(file, &temp); err == nil {
			if meta, found := temp["meta"].(cu.IM); found {
				report := cu.IM{"installed": false, "label": ""}
				report["report_key"] = meta["report_key"]
				report["code"] = cu.ToString(report["report_key"], "")
				if id, found := installed[cu.ToString(meta["report_key"], "")]; found {
					report["installed"] = true
					report["id"] = cu.ToInteger(id, 0)
				}
				report["report_name"] = meta["report_name"]
				report["description"] = meta["description"]
				report["file_type"] = meta["file_type"]
				report["label"] = cu.ToString(meta["report_type"], "")
				report["label"] = cu.ToString(meta["trans_type"], cu.ToString(report["label"], ""))
				report["file_name"] = fileName
				if (filter == "") || (filter == cu.ToString(report["label"], "")) {
					results = append(results, report)
				}
			}
		}
	}

	var file []byte
	if reportDir == "" {
		err = fs.WalkDir(st.Report, path.Join("template"), func(path string, d fs.DirEntry, err error) error {
			if filepath.Ext(path) == ".json" && err == nil {
				if file, err = st.Report.ReadFile(path); err == nil {
					fileInfo(file, d.Name())
				}
			}
			return err
		})
	} else {
		err = filepath.Walk(reportDir, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".json" && err == nil {
				if file, err = ds.ReadFile(filepath.Clean(path)); err == nil {
					fileInfo(file, info.Name())
				}
			}
			return err
		})
	}

	return results, err
}

func setSourceLabel(sources []cu.SM, labels cu.IM) []cu.SM {
	for key, label := range labels {
		for si := 0; si < len(sources); si++ {
			sources[si]["sqlstr"] = strings.ReplaceAll(
				sources[si]["sqlstr"], "={{"+key+"}}", cu.ToString(label, ""))
		}
	}
	return sources
}

func (ds *DataStore) getReportData(reportTemplate, filters cu.IM, sources []cu.SM) (datarows cu.IM, err error) {
	datarows = make(cu.IM, len(sources)+2)

	reportData := cu.ToIM(reportTemplate["data"], cu.IM{})
	reportMeta := cu.ToIM(reportTemplate["meta"], cu.IM{})
	sources = setSourceLabel(sources, cu.ToIM(reportData["labels"], cu.IM{}))

	// Determine where clause once
	isReport := cu.ToString(reportMeta["report_type"], "") == "REPORT"
	whereStr := cu.SM{}
	SetPDFWhere(filters, sources)
	if isReport {
		whereStr = SetReportWhere(reportTemplate, filters, sources)
	}

	trows := 0
	const whereKey = "@where_str"
	for index := 0; index < len(sources); index++ {
		dsc := sources[index]
		if _, found := whereStr[dsc["dataset"]]; found {
			dsc["sqlstr"] = strings.ReplaceAll(dsc["sqlstr"], whereKey, whereStr[dsc["dataset"]])
		}
		if _, found := whereStr["nods"]; found {
			dsc["sqlstr"] = strings.ReplaceAll(dsc["sqlstr"], whereKey, whereStr["nods"])
		}
		dsc["sqlstr"] = strings.ReplaceAll(dsc["sqlstr"], whereKey, "")
		params := make([]interface{}, 0)
		datarows[dsc["dataset"]], err = ds.Db.QuerySQL(dsc["sqlstr"], params, nil)
		if err != nil {
			return datarows, err
		}
		trows += len(datarows[dsc["dataset"]].([]cu.IM))
	}

	datarows["title"] = cu.ToString(reportMeta["report_name"], "")
	datarows["crtime"] = time.Now().Format("2006-01-02 15:04")
	if trows == 0 {
		return datarows, errors.New(ut.GetMessage("en", "nodata"))
	}
	if _, found := datarows["ds"]; found {
		if len(datarows["ds"].([]cu.IM)) == 0 {
			return datarows, errors.New(ut.GetMessage("en", "nodata"))
		}
	}
	return datarows, nil
}

func getDataSources(reportSource cu.IM, conEngine string) (sources []cu.SM) {
	sources = make([]cu.SM, 0)
	for dkey, ds := range reportSource {
		dsValues := cu.SM{"dataset": dkey, "sqlstr": ""}
		for engine, sql := range ds.(cu.IM) {
			if (engine == "default" && dsValues["sqlstr"] == "") || engine == conEngine {
				dsValues["sqlstr"] = cu.ToString(sql, "")
			}
		}
		sources = append(sources, dsValues)
	}
	return sources
}

// getReport - server-side PDF and CSV report generation
func (ds *DataStore) GetReport(options cu.IM) (result cu.IM, err error) {

	filters := cu.ToIM(options["filters"], cu.IM{})
	filters["@code"] = cu.ToString(options["code"], "")
	var configs []cu.IM
	var report, datarows, reportTemplate cu.IM
	var jsonTemplate string = cu.ToString(options["template"], "")

	if jsonTemplate == "" {
		if configs, err = ds.GetDataByID(
			"config", cu.ToInteger(options["report_id"], 0), cu.ToString(options["report_key"], ""), true); err == nil {
			report = cu.ToIM(configs[0]["data"], cu.IM{})
			jsonTemplate = cu.ToString(report["template"], "")
		}
	}
	if err == nil {
		if err = ds.ConvertFromByte([]byte(jsonTemplate), &reportTemplate); err == nil {
			sources := getDataSources(cu.ToIM(reportTemplate["sources"], cu.IM{}), strings.ReplaceAll(ds.Db.Connection().Engine, "3", ""))
			if len(sources) > 0 {
				datarows, err = ds.getReportData(reportTemplate, filters, sources)
			}
		}
	}

	if err != nil {
		return result, err
	}

	if options["output"] == "data" {
		return cu.IM{
			"content_type": "application/json",
			"template":     jsonTemplate,
			"data":         datarows}, nil
	}

	reportMeta := cu.ToIM(reportTemplate["meta"], cu.IM{})
	fileType := cu.ToString(reportMeta["file_type"], "")
	if fileType == md.FileTypeCSV.String() {
		base64Encoding := (options["output"] == "base64")
		return CreateReportCSV(reportTemplate, datarows, base64Encoding)
	}

	return CreateReportPDF(
		options, datarows, ds.Config, jsonTemplate)
}
