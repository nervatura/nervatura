package api

import (
	"strings"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/report"
)

func SetPDFWhere(filters cu.IM, sources []cu.SM, params map[string][]any) {
	codeValue := cu.ToString(filters["@code"], "")
	for _, ds := range sources {
		params[ds["dataset"]] = make([]any, 0)
		count := strings.Count(ds["sqlstr"], "@code")
		for i := 0; i < count; i++ {
			params[ds["dataset"]] = append(params[ds["dataset"]], codeValue)
		}
		ds["sqlstr"] = strings.ReplaceAll(ds["sqlstr"], "@code", "?")
	}
}

// Helper function to handle data setting
func setReportData(rpt *report.Report, key string, value interface{}) error {
	if iValue, found := value.(cu.IM); found {
		value = cu.IMToSM(iValue)
	}
	if iValues, found := value.([]cu.IM); found {
		values := []cu.SM{}
		for _, iValue := range iValues {
			values = append(values, cu.IMToSM(iValue))
		}
		value = values
	}
	_, err := rpt.SetData(key, value)
	return err
}

func CreateReportPDF(options, datarows, config cu.IM, jsonTemplate string) (result cu.IM, err error) {
	// Initialize result with default values
	result = cu.IM{
		"content_type": "application/pdf",
		"data":         nil,
	}

	// Create report instance
	rpt := report.New(
		cu.ToString(options["orientation"], "p"),
		cu.ToString(options["size"], "a4"),
		cu.ToString(config["NT_REPORT_FONT_FAMILY"], ""),
		cu.ToString(config["NT_REPORT_FONT_DIR"], ""),
	)
	rpt.ImagePath = cu.ToString(config["NT_REPORT_DIR"], "")

	// Load template
	if err = rpt.LoadJSONDefinition(jsonTemplate); err == nil {
		// Process data rows
		for key, value := range datarows {
			if err = setReportData(rpt, key, value); err != nil {
				return result, err
			}
		}
		rpt.CreateReport()
		// Handle output format
		switch options["output"] {
		case "xml":
			result["template"] = rpt.Save2Xml()
			result["content_type"] = "application/xml"
		case "base64":
			result["template"], err = rpt.Save2DataURLString("Report.pdf")
		default:
			result["template"], err = rpt.Save2Pdf()
		}
	}

	return result, err
}
