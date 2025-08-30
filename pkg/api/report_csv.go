package api

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"strings"

	cu "github.com/nervatura/component/pkg/util"
)

func SetReportWhere(reportTemplate, filters cu.IM, sources []cu.SM) cu.SM {
	// Pre-allocate maps with initial capacity
	whereStr := make(cu.SM, len(filters))
	fields := cu.ToIM(reportTemplate["fields"], cu.IM{})

	setWhere := func(wkey, fieldname, rel string, filterValue interface{}) {
		field := cu.ToIM(fields[fieldname], cu.IM{})
		sqlStr := cu.ToString(field["sqlstr"], "")
		filterStr := cu.ToString(filterValue, "")

		var fstr string
		if sqlStr == "" {
			fstr = fieldname + rel + filterStr
		} else {
			fstr = strings.ReplaceAll(sqlStr, "@"+fieldname, filterStr)
		}

		if existing, found := whereStr[wkey]; found {
			whereStr[wkey] = existing + " and " + fstr
		} else {
			whereStr[wkey] = " and " + fstr
		}
	}

	for fieldname, filterValue := range filters {
		if field, found := fields[fieldname]; found {
			fieldMap := cu.ToIM(field, cu.IM{})
			fieldtype := cu.ToString(fieldMap["fieldtype"], "")

			// Process filter value based on field type
			switch fieldtype {
			case "date", "string":
				strValue := cu.ToString(filterValue, "")
				if !strings.HasPrefix(strValue, "'") {
					filters[fieldname] = "'" + strValue + "'"
				}
			}

			rel := " = "
			if fieldtype == "string" {
				rel = " like "
			}

			// Handle where conditions
			if cu.ToString(fieldMap["wheretype"], "") == "where" {
				dataset := cu.ToString(fieldMap["dataset"], "")
				if dataset == "" {
					setWhere("nods", fieldname, rel, filters[fieldname])
				} else {
					// Process sources
					for _, ds := range sources {
						if dataset == ds["dataset"] {
							setWhere(dataset, fieldname, rel, filters[fieldname])
						}
					}
				}
				continue
			}

			// Handle non-where conditions
			for _, ds := range sources {
				sqlStr := cu.ToString(fieldMap["sqlstr"], "")
				filterStr := cu.ToString(filters[fieldname], "")

				if sqlStr == "" {
					ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], "@"+fieldname, filterStr)
				} else {
					fstr := strings.ReplaceAll(sqlStr, "@"+fieldname, filterStr)
					ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], "@"+fieldname, fstr)
				}
			}
		}
	}
	return whereStr
}

func createHeaderRow(columns []interface{}, labels cu.IM, hasLabels bool) []string {
	headerRow := make([]string, 0, len(columns))

	for _, col := range columns {
		fieldname := cu.ToString(col, "")
		if hasLabels {
			if label, exists := labels[fieldname]; exists {
				fieldname = cu.ToString(label, "")
			}
		}
		headerRow = append(headerRow, fieldname)
	}

	return headerRow
}

func createDataRows(data []cu.IM, columns []interface{}) [][]string {
	rows := make([][]string, 0, len(data))

	for _, item := range data {
		row := make([]string, 0, len(columns))
		for _, col := range columns {
			colName := cu.ToString(col, "")
			if value, exists := item[colName]; exists {
				row = append(row, cu.ToString(value, ""))
			}
		}
		rows = append(rows, row)
	}

	return rows
}

func processReportDetails(details []interface{}, datarows cu.IM, labels cu.IM, hasLabels bool) [][]string {
	var rows [][]string

	for _, detail := range details {
		detailMap := detail.(cu.IM)
		databind := cu.ToString(detailMap["databind"], "")
		if columns, ok := detailMap["columns"].([]interface{}); ok {
			if data, found := datarows[databind].([]cu.IM); found {
				// Add header row
				headerRow := createHeaderRow(columns, labels, hasLabels)
				rows = append(rows, headerRow)

				// Add data rows
				dataRows := createDataRows(data, columns)
				rows = append(rows, dataRows...)
			}
		}
	}

	return rows
}

func generateCSVContent(rows [][]string) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	err := writer.WriteAll(rows)
	return buf.Bytes(), err
}

func CreateReportCSV(reportTemplate, datarows cu.IM, base64Encoding bool) (result cu.IM, err error) {
	result = cu.IM{
		"content_type": "text/csv",
		"data":         nil,
	}
	rows := make([][]string, 0)
	labels, lbFound := reportTemplate["data"].(cu.IM)["labels"].(cu.IM)

	if details, valid := reportTemplate["details"].([]interface{}); valid {
		rows = processReportDetails(details, datarows, labels, lbFound)
	}

	var csvContent []byte
	if csvContent, err = generateCSVContent(rows); err == nil {
		result["template"] = string(csvContent)
		if base64Encoding {
			result["template"] = base64.URLEncoding.EncodeToString(csvContent)
		}
	}

	return result, err
}
