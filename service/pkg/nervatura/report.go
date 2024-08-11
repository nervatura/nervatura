package nervatura

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"strings"
	"time"

	ut "github.com/nervatura/nervatura/service/pkg/utils"
	"github.com/nervatura/report"
)

func (nstore *NervaStore) getReportRefValues(options IM) (nervatype string, refValues IM, err error) {
	nervatype = ut.ToString(options["nervatype"], "")
	refnumber := ut.ToString(options["refnumber"], "")
	refValues = IM{}
	if nervatype != "" && refnumber != "" {
		refValues, err = nstore.GetInfofromRefnumber(IM{"nervatype": nervatype, "refnumber": refnumber})
	}
	return nervatype, refValues, err
}

func (nstore *NervaStore) getReportDefault(options IM) (report IM, err error) {
	nervatype, refValues, err := nstore.getReportRefValues(options)
	if err == nil && nervatype != "" {
		params := IM{
			"qkey":      "default_report",
			"nervatype": nervatype}
		if _, found := refValues["transtype"]; found {
			params["transtype"] = ut.ToString(refValues["transtype"], "")
		}
		if _, found := refValues["direction"]; found {
			params["direction"] = ut.ToString(refValues["direction"], "")
		}
		rdata, err := nstore.ds.QueryKey(params, nil)
		if err != nil {
			return nil, err
		}
		if len(rdata) == 0 {
			return nil, errors.New(ut.GetMessage("not_exist"))
		}
		data := rdata[0]
		data["ref_id"] = refValues["id"]
		return data, nil
	}
	return IM{}, nil
}

func (nstore *NervaStore) getReportHead(options IM) (report IM, err error) {
	if report, valid := options["report"].(IM); valid {
		return report, nil
	}
	reportKey := ut.ToString(options["reportkey"], "")
	reportId := ut.ToFloat(options["report_id"], 0)
	params := IM{"qkey": "default_report"}
	if reportId > 0 {
		params["report_id"] = reportId
	} else if reportKey != "" {
		params["reportkey"] = reportKey
	} else {
		return nstore.getReportDefault(options)
	}
	rdata, err := nstore.ds.QueryKey(params, nil)
	if err != nil {
		return nil, err
	}
	if len(rdata) == 0 {
		return nil, errors.New(ut.GetMessage("not_exist"))
	}
	data := rdata[0]
	nervatype, refValues, err := nstore.getReportRefValues(options)
	if err == nil && nervatype != "" {
		data["ref_id"] = refValues["id"]
	}
	return data, err
}

func (nstore *NervaStore) getReportDataWhere(reportTemplate, filters IM, sources []SM) (SM, error) {
	whereStr := SM{}
	fields := IM{}
	if tFields, found := reportTemplate["fields"].(IM); found {
		fields = tFields
	}

	setWhere := func(wkey, fieldname, rel string) {
		fstr := ""
		if fields[fieldname].(IM)["sql"] == nil || fields[fieldname].(IM)["sql"] == "" {
			fstr = fieldname + rel + ut.ToString(filters[fieldname], "")
		} else {
			fstr = strings.ReplaceAll(fstr, "@"+fieldname, ut.ToString(filters[fieldname], ""))
		}
		if _, found := whereStr[wkey]; !found {
			whereStr[wkey] = " and " + fstr
		} else {
			whereStr[wkey] = whereStr[wkey] + " and " + fstr
		}
	}

	for fieldname, value := range filters {
		if _, found := fields[fieldname]; !found {
			if fieldname == "@id" {
				for index := 0; index < len(sources); index++ {
					ds := sources[index]
					ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], "@id", ut.ToString(value, ""))
				}
			} else {
				return whereStr, errors.New(ut.GetMessage("invalid_fieldname") + ": " + fieldname)
			}
		} else {
			rel := " = "
			if fields[fieldname].(IM)["fieldtype"] == "date" {
				filters[fieldname] = "'" + ut.ToString(filters[fieldname], "") + "'"
			}
			if fields[fieldname].(IM)["fieldtype"] == "string" {
				fieldtype := ut.ToString(filters[fieldname], "")
				if !strings.HasPrefix(fieldtype, "'") {
					filters[fieldname] = "'" + fieldtype + "'"
				}
				rel = " like "
			}
			if fields[fieldname].(IM)["wheretype"] == "where" && fields[fieldname].(IM)["dataset"] == nil {
				setWhere("nods", fieldname, rel)
			}
			for index := 0; index < len(sources); index++ {
				ds := sources[index]
				if fields[fieldname].(IM)["wheretype"] == "where" {
					if fields[fieldname].(IM)["dataset"] == ds["dataset"] {
						setWhere(ut.ToString(fields[fieldname].(IM)["dataset"], ""), fieldname, rel)
					}
				} else {
					if fields[fieldname].(IM)["sql"] == nil || fields[fieldname].(IM)["sql"] == "" {
						ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], "@"+fieldname, ut.ToString(filters[fieldname], ""))
					} else {
						fstr := strings.ReplaceAll(ut.ToString(fields[fieldname].(IM)["sql"], ""), "@"+fieldname, ut.ToString(filters[fieldname], ""))
						ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], "@"+fieldname, fstr)
					}
				}
			}
		}
	}
	return whereStr, nil
}

func (nstore *NervaStore) getReportData(reportTemplate, filters IM, sources []SM) (datarows IM, err error) {
	datarows = IM{}

	if labels, found := reportTemplate["data"].(IM)["labels"].(IM); found {
		for key, label := range labels {
			for si := 0; si < len(sources); si++ {
				sources[si]["sqlstr"] = strings.ReplaceAll(
					sources[si]["sqlstr"], "={{"+key+"}}", ut.ToString(label, ""))
			}
		}
	}

	whereStr, err := nstore.getReportDataWhere(reportTemplate, filters, sources)
	if err != nil {
		return datarows, err
	}

	trows := 0
	const whereKey = "@where_str"
	for index := 0; index < len(sources); index++ {
		ds := sources[index]
		if _, found := whereStr[ds["dataset"]]; found {
			ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], whereKey, whereStr[ds["dataset"]])
		}
		if _, found := whereStr["nods"]; found {
			ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], whereKey, whereStr["nods"])
		}
		ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], whereKey, "")
		params := make([]interface{}, 0)
		datarows[ds["dataset"]], err = nstore.ds.QuerySQL(ds["sqlstr"], params, nil)
		if err != nil {
			return datarows, err
		}
		trows += len(datarows[ds["dataset"]].([]IM))
	}
	datarows["title"] = reportTemplate["meta"].(IM)["repname"]
	datarows["crtime"] = time.Now().Format(TimeLayout)
	if trows == 0 {
		return datarows, errors.New(ut.GetMessage("nodata"))
	}
	if _, found := datarows["ds"]; found {
		if len(datarows["ds"].([]IM)) == 0 {
			return datarows, errors.New(ut.GetMessage("nodata"))
		}
	}
	return datarows, nil
}

func (nstore *NervaStore) getReportCSV(reportTemplate, datarows IM, base64Encoding bool) (result IM, err error) {
	rows := make([][]string, 0)
	labels, lbFound := reportTemplate["data"].(IM)["labels"].(IM)
	if details, valid := reportTemplate["details"].([]interface{}); valid {
		for di := 0; di < len(details); di++ {
			databind := ut.ToString(details[di].(IM)["databind"], "")
			columns, cols := details[di].(IM)["columns"].([]interface{})
			if data, found := datarows[databind].([]IM); found && cols {
				row := make([]string, 0)
				for ci := 0; ci < len(columns); ci++ {
					fieldname := ut.ToString(columns[ci], "")
					if lbFound {
						if label, lbValue := labels[fieldname]; lbValue {
							fieldname = ut.ToString(label, "")
						}
					}
					row = append(row, fieldname)
				}
				rows = append(rows, row)

				for i := 0; i < len(data); i++ {
					row := make([]string, 0)
					for ci := 0; ci < len(columns); ci++ {
						if value, vcol := data[i][ut.ToString(columns[ci], "")]; vcol {
							row = append(row, ut.ToString(value, ""))
						}
					}
					rows = append(rows, row)
				}
			}
		}
	}
	var b bytes.Buffer
	writr := csv.NewWriter(&b)
	if err = writr.WriteAll(rows); err != nil {
		return result, err
	}
	if base64Encoding {
		return IM{"filetype": "csv",
			"template": base64.URLEncoding.EncodeToString(b.Bytes()), "data": nil}, nil
	}
	return IM{"filetype": "csv", "template": b.String(), "data": nil}, nil
}

func (nstore *NervaStore) getReportPDF(options, datarows IM, jsonTemplate string) (result IM, err error) {
	orientation := ut.ToString(options["orientation"], "p")
	size := ut.ToString(options["size"], "a4")
	rpt := report.New(orientation, size,
		ut.ToString(nstore.config["NT_FONT_FAMILY"], ""), ut.ToString(nstore.config["NT_FONT_DIR"], ""))
	rpt.ImagePath = ut.ToString(nstore.config["NT_REPORT_DIR"], "")
	if err = rpt.LoadJSONDefinition(jsonTemplate); err != nil {
		return result, err
	}
	for key, value := range datarows {
		switch v := value.(type) {
		case string, map[string]string, []map[string]string:
			_, err = rpt.SetData(key, v)
		case map[string]interface{}:
			values := SM{}
			for skey, ivalue := range v {
				values[skey] = ut.ToString(ivalue, "")
			}
			_, err = rpt.SetData(key, values)
		case []map[string]interface{}:
			ivalues := []SM{}
			for index := 0; index < len(v); index++ {
				values := SM{}
				for skey, ivalue := range v[index] {
					values[skey] = ut.ToString(ivalue, "")
				}
				ivalues = append(ivalues, values)
			}
			_, err = rpt.SetData(key, ivalues)
		}
		if err != nil {
			return result, err
		}
	}
	rpt.CreateReport()

	switch options["output"] {
	case "xml":
		xml := rpt.Save2Xml()
		return IM{"filetype": "xml", "template": xml, "data": nil}, nil

	case "base64":
		pdf, err := rpt.Save2DataURLString("Report.pdf")
		if err != nil {
			return result, err
		}
		return IM{"filetype": "base64", "template": pdf, "data": nil}, nil

	default:
		pdf, err := rpt.Save2Pdf()
		if err != nil {
			return result, err
		}
		return IM{"filetype": "pdf", "template": pdf, "data": nil}, nil
	}
}

// getReport - server-side PDF and CSV report generation
func (nstore *NervaStore) getReport(options IM) (results IM, err error) {

	results = IM{
		"report":   IM{},
		"datarows": IM{},
	}
	filters := IM{}
	if ofilters, valid := options["filters"].(IM); valid {
		filters = ofilters
	}
	results["report"], err = nstore.getReportHead(options)
	if err != nil {
		return results, err
	}
	reportkey := ut.ToString(results["report"].(IM)["reportkey"], "")
	jsonTemplate := ut.ToString(options["template"], ut.ToString(results["report"].(IM)["report"], ""))
	reportTemplate := IM{}
	err = ut.ConvertFromByte([]byte(jsonTemplate), &reportTemplate)
	if err != nil {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": report_id or reportkey")
	}

	if reportkey != "" {
		if ref_id, found := results["report"].(IM)["ref_id"]; found && ut.ToString(filters["@id"], "") == "" {
			filters["@id"] = ut.ToString(ref_id, "")
		}

		sources := make([]SM, 0)
		cengine := strings.ReplaceAll(nstore.ds.Connection().Engine, "3", "")
		if tSources, found := reportTemplate["sources"].(IM); found {
			for dkey, ds := range tSources {
				dsValues := SM{"dataset": dkey, "sqlstr": ""}
				for engine, sql := range ds.(IM) {
					if (engine == "default" && dsValues["sqlstr"] == "") || engine == cengine {
						dsValues["sqlstr"] = ut.ToString(sql, "")
					}
				}
				sources = append(sources, dsValues)
			}
		}

		results["datarows"], err = nstore.getReportData(reportTemplate, filters, sources)
		if err != nil {
			return results, err
		}
	}

	if options["output"] == "data" {
		return IM{
			"filetype": ut.ToString(results["report"].(IM)["reptype"], ""),
			"template": ut.ToString(results["report"].(IM)["report"], ""),
			"data":     results["datarows"]}, nil
	}
	reptype := ut.ToString(results["report"].(IM)["reptype"], "")
	if reptype == "csv" {
		base64Encoding := (options["output"] == "base64")
		return nstore.getReportCSV(reportTemplate, results["datarows"].(IM), base64Encoding)
	}
	return nstore.getReportPDF(
		options, results["datarows"].(IM), jsonTemplate)
}
