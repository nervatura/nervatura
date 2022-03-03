package report

import (
	"image/color"
	"io/ioutil"
	"path"
	"testing"

	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

func createGoReport(t *testing.T) (rpt *Report) {
	var appendElement = func(parent interface{}, ename string, values IM) *[]PageItem {
		eParent, err := rpt.AppendElement(parent, ename, values)
		if err != nil {
			t.Fatal(err)
		}
		return eParent
	}
	var setData = func(key string, value interface{}) {
		if _, err := rpt.SetData(key, value); err != nil {
			t.Fatal(err)
		}
	}

	//rpt = New("p", "A4", "NotoSans", "../../data/fonts")
	rpt = New("p", "A4")

	//default values
	rpt.SetReportValue("title", "Go Report")
	rpt.SetReportValue("fontSize", float64(9))

	//header
	rowData := appendElement("header", "row", IM{"height": 10})
	appendElement(rowData, "image", IM{"src": "logo"})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "labels.title",
		"font-style": "bolditalic", "font-size": 26, "color": "#D8DBDA"})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "Go Sample", "font-style": "bold", "align": "right"})
	appendElement("header", "vgap", IM{"height": 2})
	appendElement("header", "hline", IM{"border-color": 218})
	appendElement("header", "vgap", IM{"height": 2})

	//details
	appendElement("details", "vgap", IM{"height": 2})
	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "label", "width": "50%", "font-style": "bold", "value": "labels.left_text", "border": "LT",
		"border-color": 218, "background-color": 245})
	appendElement(rowData, "cell", IM{
		"name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LTR",
		"border-color": 218, "background-color": 245})

	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "width": "50%", "value": "head.short_text", "border": "L", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "value": "head.short_text", "border": "LR", "border-color": 218})
	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "width": "50%", "value": "head.short_text", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "value": "head.short_text", "border": "LBR", "border-color": 218})

	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "label", "width": "40", "font-style": "bold", "value": "labels.left_text", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "label", "align": "center", "width": "30", "font-style": "bold", "value": "labels.center_text",
		"border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "label", "align": "right", "width": "40", "font-style": "bold", "value": "labels.right_text",
		"border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LBR", "border-color": 218})

	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "width": "40", "value": "head.short_text", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "date", "align": "center", "width": "30", "value": "head.date", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "amount", "align": "right", "width": "40", "value": "head.number", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "value": "head.short_text", "border": "LBR", "border-color": 218})

	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "width": "50", "value": "head.short_text", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LB", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "value": "head.short_text", "border": "LBR", "border-color": 218})

	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "long_text", "multiline": true, "value": "head.long_text",
		"border": "LBR", "border-color": 218})

	appendElement("details", "vgap", IM{"height": 2})
	rowData = appendElement("details", "row", IM{"hgap": 2})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "labels.left_text", "font-style": "bold", "border": "1", "border-color": 218,
		"background-color": 245})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "value": "head.short_text", "border": "1", "border-color": 218})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "labels.left_text", "font-style": "bold", "border": "1", "border-color": 218, "background-color": 245})
	appendElement(rowData, "cell", IM{
		"name": "short_text", "value": "head.short_text", "border": "1", "border-color": 218})

	appendElement("details", "vgap", IM{"height": 2})
	rowData = appendElement("details", "row", IM{"hgap": 2})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "labels.long_text", "font-style": "bold", "border": "1", "border-color": 218, "background-color": 245})
	appendElement(rowData, "cell", IM{
		"name": "long_text", "multiline": true, "value": "head.long_text", "border": "1", "border-color": 218})

	appendElement("details", "vgap", IM{"height": 2})
	appendElement("details", "hline", IM{"border-color": 218})
	appendElement("details", "vgap", IM{"height": 2})

	rowData = appendElement("details", "row", IM{"hgap": 5})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "Barcode (code 128)", "font-style": "bold", "font-size": 9, "width": "40",
		"border": "1", "border-color": 245, "background-color": 245})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "Barcode (Interleaved 2of5)", "font-style": "bold", "font-size": 9, "width": "40",
		"border": "1", "border-color": 245, "background-color": 245})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "Barcode (EAN)", "font-style": "bold", "font-size": 9, "width": "40",
		"border": "1", "border-color": 245, "background-color": 245})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "Barcode (Code 39)", "font-style": "bold", "font-size": 9, "width": "40",
		"border": "1", "border-color": 245, "background-color": 245})

	rowData = appendElement("details", "row", IM{"hgap": 5})
	appendElement(rowData, "barcode",
		IM{"code-type": "CODE_128", "value": "1234567890ABCDEF", "width": 40, "visible-value": 1})
	appendElement(rowData, "barcode",
		IM{"code-type": "ITF", "value": "1234567890", "width": 40, "visible-value": 1})
	appendElement(rowData, "barcode",
		IM{"code-type": "EAN", "value": "96385074", "width": 40, "visible-value": 1})
	appendElement(rowData, "barcode",
		IM{"code-type": "CODE_39", "value": "1234567890ABCDEF", "width": 40, "extend": true, "visible-value": 1})

	appendElement("details", "vgap", IM{"height": 3})

	rowData = appendElement("details", "row", IM{"hgap": 5})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "QR code: Hello Go Report!", "font-size": 9,
		"border": "1", "border-color": 245, "background-color": 245})
	appendElement(rowData, "barcode",
		IM{"code-type": "QR", "value": "Hello Go Report!", "height": 10})
	appendElement(rowData, "cell", IM{})
	appendElement("details", "vgap", IM{"height": 3})

	rowData = appendElement("details", "row", IM{})
	appendElement(rowData, "cell", IM{
		"name": "label", "value": "Datagrid Sample", "align": "center", "font-style": "bold",
		"border": "1", "border-color": 245, "background-color": 245})
	appendElement("details", "vgap", IM{"height": 2})

	var gridData = appendElement("details", "datagrid", IM{
		"name": "items", "databind": "items", "border": "1", "border-color": 218, "header-background": 245, "footer-background": 245})
	appendElement(gridData, "column", IM{
		"width": "8%", "fieldname": "counter", "align": "right", "label": "labels.counter", "footer": "labels.total"})
	appendElement(gridData, "column", IM{
		"width": "20%", "fieldname": "date", "align": "center", "label": "labels.center_text"})
	appendElement(gridData, "column", IM{
		"width": "15%", "fieldname": "number", "align": "right", "label": "labels.right_text",
		"footer": "items_footer.items_total", "footer-align": "right"})
	appendElement(gridData, "column", IM{
		"fieldname": "text", "label": "labels.left_text"})

	appendElement("details", "vgap", IM{"height": 5})
	appendElement("details", "html", IM{"fieldname": "html_text",
		"html": "<i>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</i> ={{html_text}} <p>Nulla a <b><i>pretium</i></b> nunc, in <u>cursus</u> quam.</p>"})

	//footer
	appendElement("footer", "vgap", IM{"height": 2})
	appendElement("footer", "hline", IM{"border-color": 218})
	rowData = appendElement("footer", "row", IM{"height": 10})
	appendElement(rowData, "cell", IM{"value": "Nervatura Report Template", "font-style": "bolditalic"})
	appendElement(rowData, "cell", IM{"value": "{{page}}", "align": "right", "font-style": "bold"})

	//data
	setData("labels", SM{"title": "REPORT TEMPLATE", "left_text": "Short text", "center_text": "Centered text",
		"right_text": "Right text", "long_text": "Long text", "counter": "No.", "total": "Total"})
	setData("head", SM{"short_text": "Lorem éáőűúóüö dolor", "number": "123 456", "date": "2015.01.01",
		"long_text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim. Nulla a pretium nunc, in cursus quam."})
	setData("html_text", "<p><b>Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim.</b></p>")
	setData("items_footer", SM{"items_total": "3 703 680"})
	setData("logo", "data:image/jpg;base64,/9j/4AAQSkZJRgABAQIA7ADsAAD/2wBDAAoHBwgHBgoICAgLCgoLDhgQDg0NDh0VFhEYIx8lJCIfIiEmKzcvJik0KSEiMEExNDk7Pj4+JS5ESUM8SDc9Pjv/2wBDAQoLCw4NDhwQEBw7KCIoOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozv/wgARCABAAEADAREAAhEBAxEB/8QAGwABAAIDAQEAAAAAAAAAAAAAAAQGAQMFAgf/xAAYAQEBAQEBAAAAAAAAAAAAAAAAAQIDBP/aAAwDAQACEAMQAAABuYAAAAAAPCQZno3eCBMc+YjyRJnmzndenq6F3VMefg55Abl2W37p68Hznl4sHolXcWYs2+9h12wUPn5MGCLM9C7u/T1ejmzEaZr2eOiTr66WfXfbaAAAAP/EACMQAAICAgEDBQEAAAAAAAAAAAIDAQQABRATITAREhQVIjH/2gAIAQEAAQUC8xEIRXd8hvDLtdWFtl5O2LC2VksMzZNJfSq5sLkkfMh7QrK61jJ7D/eIiSn8oyZkp1aPQeG1yW/prHCd2ypVKyyIgYy1Fks+tCVMpvVMKZOV9YZYtYqDxf/EACERAAEDAwUBAQAAAAAAAAAAAAEAAhEQE0EDEiAwUSEy/9oACAEDAQE/Ae8Gal4CuhXVcciZTBApqPwOEQmiTy/NNMZqW/YUDK3eUY3dV27CtiEWEKCm6fqAjr//xAAbEQACAgMBAAAAAAAAAAAAAAABEQAQEiAwQP/aAAgBAgEBPwHxOZTKOhROo4C1oBZiipdP/8QAKRAAAQMCBAUEAwAAAAAAAAAAAQACEQMhEDFBYRIiMHGRICNRoTJCYv/aAAgBAQAGPwLrS4gDdPe38G2G+MF8n4C5abj3VqQ8qxDewUvcXd0wam5wNFhhoz39EuzOQTGb3wJxgCStHVPpqkmSjWOthi6l4k6L3Ks7MuuGm3gb9nD+BmUALAYNFB0A5ohziah/cqDTJ3F1am7wprco+NVwMEAdP//EACUQAQACAAUEAQUAAAAAAAAAAAEAERAhMUFRMGFxkdEgobHw8f/aAAgBAQABPyHrOTrdVBvNQ/Ld/GLlPyMF90BNn/N4bX69vLQTlXL0KD3YJyybN30VyV5muWcEtvDfB9gIqlW1wMuTQJWt2evyMcuTNWMdz9XDWU0Ms0APK2bd7O3vSUMzZNN+RwoGZ/Chl0KA2wrmyj3JrMv3JVUcNItSXgcXG7fV8QIHbHT/AP/aAAwDAQACAAMAAAAQAAAAAAANEty9DknkCjioltiArqvAAAAA/8QAIBEBAAICAQQDAAAAAAAAAAAAAQARECExMEFRYSCRsf/aAAgBAwEBPxDrKG2X1OM8kx7BL+ItE5MoTCL8Bqt5laZW94BWia9n8iq2yst3y4xK+X1HShRhV6gVowFRNiF35iXE9cV3ACjp/wD/xAAcEQACAgIDAAAAAAAAAAAAAAABEQAQIEEwMVH/2gAIAQIBAT8Q5wXZARKPCSYCFahgoDOXVDuyhUQHcflMttRVCQiMHuALj//EACUQAQABAwMDBQEBAAAAAAAAAAERACExQWGBEHGhMFGRscEg0f/aAAgBAQABPxD1svViA+auiElpLweHV2B5GZspY5aAZ7T8WaeMjpN9AqIC2f8AVbp1cO04rVUIIZuv2IOOi6wWIdY7GN/4ljjsNngDoZc4iXhJN2N/AjmgAgsU4GUT4pyyFVyvR9TwOVe1QZRsCHb6bWDekkDISrUKcbppbvLBx0QCJI2aOgLdIIkBj9ofKRqE5wOFqEJMI/nTYg26G0Kjt+zd8Zo+oglgYOjrVq4PZZdM4vTGblwoNInHvN6ZaDOXkxzFaDikH6rUecg/h5O1BLC32Lq7+n//2Q==")
	items := make([]SM, 0)
	items = append(items, SM{"text": "Lorem ipsum dolor", "number": "123 456", "date": "Lorem ipsum dolorjkjkjl jhkjh"})
	for index := 0; index < 20; index++ {
		items = append(items, SM{"text": "Lorem ipsum dolor", "number": "123 456", "date": "2015.01.01"})
	}
	items = append(items, SM{"text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim. Nulla a pretium nunc, in cursus quam.", "number": "123 456", "date": "2015.01.01"})
	for index := 0; index < 20; index++ {
		items = append(items, SM{"text": "Lorem ipsum dolor", "number": "123 456", "date": "2015.01.01"})
	}
	setData("items", items)

	rpt.CreateReport()
	return rpt
}

func TestCreateGoReport(t *testing.T) {
	rpt := createGoReport(t)
	if err := rpt.Save2PdfFile("../../data/test/go.pdf"); err != nil {
		t.Fatal(err)
	}
}

func TestXMLData(t *testing.T) {
	rpt := createGoReport(t)
	dataXml := rpt.Save2Xml()
	if err := ioutil.WriteFile("../../data/test/data.xml", []byte(dataXml), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestJSONReport(t *testing.T) {
	json, _ := ut.Report.ReadFile(path.Join("static", "templates", "sample.json"))
	rpt := New("L")
	if err := rpt.LoadJSONDefinition(string(json)); err != nil {
		t.Fatal(err)
	}
	rpt.CreateReport()
	pdf, err := rpt.Save2Pdf()
	if err != nil {
		t.Fatal(err)
	}
	if err = ioutil.WriteFile("../../data/test/json.pdf", pdf, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestBase64Report(t *testing.T) {
	json, _ := ut.Report.ReadFile(path.Join("static", "templates", "sample.json"))
	rpt := New()
	if err := rpt.LoadJSONDefinition(string(json)); err != nil {
		t.Fatal(err)
	}
	rpt.CreateReport()
	base64Str, err := rpt.Save2DataURLString("Report.pdf")
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile("../../data/test/base64.txt", []byte(base64Str), 0644); err != nil {
		t.Fatal(err)
	}

}

func TestPageItem_setPageItem(t *testing.T) {
	type fields struct {
		ItemType string
		Item     interface{}
	}
	type args struct {
		fieldname string
		value     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "row_Visible",
			fields: fields{
				ItemType: "row",
				Item: &Row{
					Columns: make([]PageItem, 0),
				},
			},
			args: args{
				fieldname: "Visible",
				value:     "true",
			},
			wantErr: false,
		},
		{
			name: "image_Height",
			fields: fields{
				ItemType: "image",
				Item:     &Image{},
			},
			args: args{
				fieldname: "Height",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "separator_Gap",
			fields: fields{
				ItemType: "separator",
				Item:     &Separator{},
			},
			args: args{
				fieldname: "Gap",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "vgap_Visible",
			fields: fields{
				ItemType: "vgap",
				Item:     &VGap{},
			},
			args: args{
				fieldname: "Visible",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "hline_Visible",
			fields: fields{
				ItemType: "hline",
				Item:     &HLine{},
			},
			args: args{
				fieldname: "Width",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "hline_Gap",
			fields: fields{
				ItemType: "hline",
				Item:     &HLine{},
			},
			args: args{
				fieldname: "Gap",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "hline_BorderColor",
			fields: fields{
				ItemType: "hline",
				Item:     &HLine{},
			},
			args: args{
				fieldname: "BorderColor",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "hline_Visible",
			fields: fields{
				ItemType: "hline",
				Item:     &HLine{},
			},
			args: args{
				fieldname: "Visible",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "datagrid_Width",
			fields: fields{
				ItemType: "datagrid",
				Item: &Datagrid{
					Columns: make([]PageItem, 0)},
			},
			args: args{
				fieldname: "Width",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "datagrid_Merge",
			fields: fields{
				ItemType: "datagrid",
				Item: &Datagrid{
					Columns: make([]PageItem, 0)},
			},
			args: args{
				fieldname: "Merge",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "datagrid_FontSize",
			fields: fields{
				ItemType: "datagrid",
				Item: &Datagrid{
					Columns: make([]PageItem, 0)},
			},
			args: args{
				fieldname: "FontSize",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "datagrid_TextColor",
			fields: fields{
				ItemType: "datagrid",
				Item: &Datagrid{
					Columns: make([]PageItem, 0)},
			},
			args: args{
				fieldname: "TextColor",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "datagrid_BackgroundColor",
			fields: fields{
				ItemType: "datagrid",
				Item: &Datagrid{
					Columns: make([]PageItem, 0)},
			},
			args: args{
				fieldname: "BackgroundColor",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "column_HeaderAlign",
			fields: fields{
				ItemType: "column",
				Item: &Column{
					Align:       "L",
					HeaderAlign: "L",
					FooterAlign: "L"},
			},
			args: args{
				fieldname: "HeaderAlign",
				value:     0,
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				ItemType: "table",
				Item:     nil,
			},
			args: args{
				fieldname: "Header",
				value:     0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := &PageItem{
				ItemType: tt.fields.ItemType,
				Item:     tt.fields.Item,
			}
			if err := pi.setPageItem(tt.args.fieldname, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("PageItem.setPageItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReport_getPageItem(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		etype string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "separator",
			fields: fields{},
			args: args{
				etype: "separator",
			},
			wantErr: false,
		},
		{
			name:   "error",
			fields: fields{},
			args: args{
				etype: "error",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			_, err := rpt.getPageItem(tt.args.etype)
			if (err != nil) != tt.wantErr {
				t.Errorf("Report.getPageItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestReport_SetReportValue(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		fieldname string
		value     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Author",
			fields: fields{},
			args: args{
				fieldname: "Author",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "Creator",
			fields: fields{},
			args: args{
				fieldname: "Creator",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "Subject",
			fields: fields{},
			args: args{
				fieldname: "Subject",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "Keywords",
			fields: fields{},
			args: args{
				fieldname: "Keywords",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "BottomMargin",
			fields: fields{},
			args: args{
				fieldname: "BottomMargin",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "FontStyle",
			fields: fields{},
			args: args{
				fieldname: "FontStyle",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "TextColor",
			fields: fields{},
			args: args{
				fieldname: "TextColor",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "BorderColor",
			fields: fields{},
			args: args{
				fieldname: "BorderColor",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "BackgroundColor",
			fields: fields{},
			args: args{
				fieldname: "BackgroundColor",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "ImagePath",
			fields: fields{},
			args: args{
				fieldname: "ImagePath",
				value:     "",
			},
			wantErr: false,
		},
		{
			name:   "error",
			fields: fields{},
			args: args{
				fieldname: "Error",
				value:     "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if err := rpt.SetReportValue(tt.args.fieldname, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Report.SetReportValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReport_setHTMLValue(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		value     string
		fieldname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "rek_value",
			fields: fields{},
			args: args{
				fieldname: "html_text1",
				value:     "<i>Lorem ipsum ... elit.</i> ={{html_text1}} <p>Nulla a <b><i>pretium</i></b> nunc, ={{html_text2}} in <u>cursus</u> quam.</p>",
			},
			want: "<i>Lorem ipsum ... elit.</i> html_text1 <p>Nulla a <b><i>pretium</i></b> nunc, html_text2 in <u>cursus</u> quam.</p>",
		},
		{
			name:   "not_found",
			fields: fields{},
			args: args{
				fieldname: "html_text",
				value:     "<i>Lorem ipsum ... elit.</i>",
			},
			want: "<i>Lorem ipsum ... elit.</i>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if got := rpt.setHTMLValue(tt.args.value, tt.args.fieldname); got != tt.want {
				t.Errorf("Report.setHTMLValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_setValue(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "len_3",
			fields: fields{
				data: map[string]interface{}{
					"item": []map[string]string{
						{
							"value": "value",
						},
					},
				},
			},
			args: args{
				value: "item.0.value",
			},
			want: "value",
		},
		{
			name: "len_2",
			fields: fields{
				data: map[string]interface{}{
					"item": []map[string]string{
						{
							"value": "value",
						},
					},
				},
			},
			args: args{
				value: "item.0",
			},
			want: "item.0",
		},
		{
			name: "dict_0",
			fields: fields{
				data: map[string]interface{}{
					"item": map[string]string{},
				},
			},
			args: args{
				value: "item.value",
			},
			want: "",
		},
		{
			name: "dict_err",
			fields: fields{
				data: map[string]interface{}{
					"item": map[string]string{},
				},
			},
			args: args{
				value: "item",
			},
			want: "item",
		},
		{
			name: "dict",
			fields: fields{
				data: map[string]interface{}{
					"item": map[string]string{
						"value1": "value",
						"value2": "value",
					},
				},
			},
			args: args{
				value: "value ={{value1}} value ={{value2}}",
			},
			want: "value value1 value value2",
		},
		{
			name: "invalid_index",
			fields: fields{
				data: map[string]interface{}{
					"item": []map[string]string{
						{
							"value": "value",
						},
					},
				},
			},
			args: args{
				value: "item.2.value",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if got := rpt.setValue(tt.args.value); got != tt.want {
				t.Errorf("Report.setValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_getCellHeight(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		text    string
		width   float64
		options IM
	}
	rpt := New("p", "A4")
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "empty",
			fields: fields{
				pdf:        rpt.pdf,
				FontFamily: rpt.FontFamily,
			},
			args: args{
				text: "",
				options: IM{
					"fontStyle": "",
					"fontSize":  float64(10),
				},
			},
			want: float64(16.4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if got := rpt.getCellHeight(tt.args.text, tt.args.width, tt.args.options); got != tt.want {
				t.Errorf("Report.getCellHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_createGridHeader(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		headerOptions IM
	}
	rpt := New("p", "A4")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "columnWidth_0",
			fields: fields{
				pdf: rpt.pdf,
			},
			args: args{
				headerOptions: IM{
					"merge":        false,
					"gridWidth":    float64(0),
					"columnsWidth": float64(0),
					"virtual":      true,
					"columns": []IM{
						{
							"label":       "label",
							"ln":          0,
							"columnWidth": 0,
							"headerAlign": "L",
						},
					},
				},
			},
		},
		{
			name: "merge",
			fields: fields{
				pdf: rpt.pdf,
			},
			args: args{
				headerOptions: IM{
					"merge":        true,
					"gridWidth":    float64(0),
					"columnsWidth": float64(0),
					"virtual":      true,
					"text":         "text",
					"columns": []IM{
						{
							"label":       "label",
							"ln":          0,
							"columnWidth": 0,
							"headerAlign": "L",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			rpt.createGridHeader(tt.args.headerOptions)
		})
	}
}

func TestReport_createDatagrid(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		gridElement *Datagrid
	}
	rpt := New("p", "A4")
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Columns_0",
			fields: fields{
				pdf: rpt.pdf,
			},
			args: args{
				gridElement: &Datagrid{
					Columns: make([]PageItem, 0),
				},
			},
			want: false,
		},
		{
			name: "Columns_0",
			fields: fields{
				pdf:  rpt.pdf,
				data: make(map[string]interface{}),
			},
			args: args{
				gridElement: &Datagrid{
					Columns: []PageItem{
						{},
						{},
					},
					Databind: "db",
				},
			},
			want: false,
		},
		{
			name: "Columns_gridWidth",
			fields: fields{
				pdf: rpt.pdf,
				data: IM{
					"db": []SM{
						{"id": ""},
					},
				},
			},
			args: args{
				gridElement: &Datagrid{
					Width: "600",
					Columns: []PageItem{
						{
							ItemType: "column",
							Item: &Column{
								Fieldname:   "field",
								Label:       "label",
								Width:       "",
								Footer:      "",
								Align:       "L",
								HeaderAlign: "L",
								FooterAlign: "L"},
						},
					},
					Databind: "db",
				},
			},
			want: true,
		},
		{
			name: "merge",
			fields: fields{
				pdf: rpt.pdf,
				data: IM{
					"db": []SM{
						{"id": ""},
					},
				},
			},
			args: args{
				gridElement: &Datagrid{
					Columns: []PageItem{
						{
							ItemType: "column",
							Item: &Column{
								Fieldname:   "field",
								Label:       "label",
								Width:       "",
								Footer:      "",
								Align:       "L",
								HeaderAlign: "L",
								FooterAlign: "L"},
						},
					},
					Databind: "db",
					Merge:    true,
				},
			},
			want: true,
		},
		{
			name: "columnWidth",
			fields: fields{
				pdf: rpt.pdf,
				data: IM{
					"db": []SM{
						{"id": ""},
					},
				},
			},
			args: args{
				gridElement: &Datagrid{
					Columns: []PageItem{
						{
							ItemType: "column",
							Item: &Column{
								Fieldname:   "field",
								Label:       "label",
								Width:       "20",
								Footer:      "",
								Align:       "L",
								HeaderAlign: "L",
								FooterAlign: "L"},
						},
					},
					Databind: "db",
				},
			},
			want: true,
		},
		{
			name: "cols",
			fields: fields{
				pdf: rpt.pdf,
				data: IM{
					"db": []SM{
						{"id": ""},
					},
				},
			},
			args: args{
				gridElement: &Datagrid{
					Columns: []PageItem{
						{
							ItemType: "column",
							Item: &Column{
								Fieldname:   "field",
								Label:       "label",
								Width:       "",
								Footer:      "",
								Align:       "L",
								HeaderAlign: "L",
								FooterAlign: "L"},
						},
						{
							ItemType: "column",
							Item: &Column{
								Fieldname:   "field",
								Label:       "label",
								Width:       "",
								Footer:      "",
								Align:       "L",
								HeaderAlign: "L",
								FooterAlign: "L"},
						},
					},
					Databind: "db",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if got := rpt.createDatagrid(tt.args.gridElement, true); got != tt.want {
				t.Errorf("Report.createDatagrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_setImageSize(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		v *Image
	}
	rpt := New("p", "A4")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "decode_ok",
			fields: fields{
				pdf:       rpt.pdf,
				ImagePath: "../utils/static/client",
			},
			args: args{
				v: &Image{
					Src: "icon-192.png",
				},
			},
		},
		{
			name: "decode_error",
			fields: fields{
				pdf:       rpt.pdf,
				ImagePath: "",
			},
			args: args{
				v: &Image{
					Src: "icon-192.png",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			rpt.setImageSize(tt.args.v)
		})
	}
}

func TestReport_addToXML(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		section string
		values  []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "header",
			fields: fields{
				xmlHeader: "",
			},
			args: args{
				section: "header",
				values:  []string{"text", "value", "text"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			rpt.addToXML(tt.args.section, tt.args.values)
		})
	}
}

func TestReport_createRow(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		section    string
		rowElement *Row
		virtual    bool
	}
	rpt := New("p", "A4")
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "Separator",
			fields: fields{
				pdf: rpt.pdf,
			},
			args: args{
				rowElement: &Row{
					Columns: []PageItem{
						{
							ItemType: "separator",
							Item: &Separator{
								Gap: float64(5),
							},
						},
					},
				},
				virtual: true,
			},
			want: float64(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if got := rpt.createRow(tt.args.section, tt.args.rowElement, tt.args.virtual); got != tt.want {
				t.Errorf("Report.createRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_createLine(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		v       *HLine
		virtual bool
	}
	rpt := New("p", "A4")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "pc_width",
			fields: fields{
				pdf: rpt.pdf,
			},
			args: args{
				v: &HLine{
					Width: "30%",
				},
				virtual: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			rpt.createLine(tt.args.v, true)
		})
	}
}

func TestReport_createElement(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		section string
		element interface{}
	}
	rpt := New("p", "A4")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Row_Visible",
			fields: fields{
				pdf: rpt.pdf,
				data: IM{
					"ds": []SM{},
				},
			},
			args: args{
				section: "",
				element: &Row{
					Visible: "ds",
					Columns: make([]PageItem, 0)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			rpt.createElement(tt.args.section, tt.args.element)
		})
	}
}

func TestReport_setFont(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	rpt := New("p", "A4", "Roboto", "../../data/fonts")
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "path_font",
			fields: fields{
				pdf: rpt.pdf,
			},
			want: true,
		},
		{
			name: "path_font_error",
			fields: fields{
				pdf:     rpt.pdf,
				fontDir: "../../data",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if got := rpt.setFont(); got != tt.want {
				t.Errorf("Report.setFont() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_SetData(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "SM",
			fields: fields{
				data: IM{
					"ds": SM{},
				},
			},
			args: args{
				key: "ds",
				value: SM{
					"key": "value",
				},
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "error",
			fields: fields{
				data: IM{
					"ds": SM{},
				},
			},
			args: args{
				key:   "ds",
				value: int64(0),
			},
			wantErr: true,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			got, err := rpt.SetData(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Report.SetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Report.SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_AppendElement(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		options []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "header_error",
			fields: fields{},
			args: args{
				options: []interface{}{
					"header", "ename",
				},
			},
			wantErr: true,
		},
		{
			name:   "details_error",
			fields: fields{},
			args: args{
				options: []interface{}{
					"details", "ename",
				},
			},
			wantErr: true,
		},
		{
			name:   "footer_error",
			fields: fields{},
			args: args{
				options: []interface{}{
					"footer", "ename",
				},
			},
			wantErr: true,
		},
		{
			name:   "pitem_error",
			fields: fields{},
			args: args{
				options: []interface{}{
					&[]PageItem{}, "ename",
				},
			},
			wantErr: true,
		},
		{
			name:   "type_error",
			fields: fields{},
			args: args{
				options: []interface{}{
					int64(0), "ename",
				},
			},
			wantErr: true,
		},
		{
			name:   "params_error",
			fields: fields{},
			args: args{
				options: []interface{}{
					"footer", "row", int64(0),
				},
			},
			wantErr: true,
		},
		{
			name:   "setPageItem_error",
			fields: fields{},
			args: args{
				options: []interface{}{
					"footer", "row", IM{"field": "value"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			_, err := rpt.AppendElement(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Report.AppendElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestReport_getJSONElements(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		edata interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "invalid_element",
			fields: fields{},
			args: args{
				edata: IM{
					"invalid": IM{},
				},
			},
			wantErr: true,
		},
		{
			name:   "columns_setPageItem_error",
			fields: fields{},
			args: args{
				edata: IM{
					"column": IM{
						"columns": []interface{}{
							IM{
								"image": IM{
									"png": "value",
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "columns_invalid_error",
			fields: fields{},
			args: args{
				edata: IM{
					"column": IM{
						"columns": []interface{}{
							IM{
								"datagrid": IM{},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "setPageItem_error",
			fields: fields{},
			args: args{
				edata: IM{
					"column": IM{
						"column": []interface{}{
							IM{
								"datagrid": IM{},
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			_, err := rpt.getJSONElements(tt.args.edata)
			if (err != nil) != tt.wantErr {
				t.Errorf("Report.getJSONElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestReport_LoadJSONDefinition(t *testing.T) {
	type fields struct {
		pdf             Generator
		orientation     string
		format          string
		fontDir         string
		xmlHeader       string
		xmlDetails      string
		header          []PageItem
		details         []PageItem
		footer          []PageItem
		data            IM
		footerHeight    float64
		pageBreak       float64
		Title           string
		Author          string
		Creator         string
		Subject         string
		Keywords        string
		LeftMargin      float64
		RightMargin     float64
		TopMargin       float64
		BottomMargin    float64
		FontFamily      string
		FontStyle       string
		FontSize        float64
		TextColor       color.RGBA
		BorderColor     color.RGBA
		BackgroundColor color.RGBA
		ImagePath       string
	}
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "missing_JSON",
			fields: fields{},
			args: args{
				jsonString: "",
			},
			wantErr: true,
		},
		{
			name:   "convert_JSON_error",
			fields: fields{},
			args: args{
				jsonString: "[{''",
			},
			wantErr: true,
		},
		{
			name:   "report_SetReportValue_error",
			fields: fields{},
			args: args{
				jsonString: `{"report":{"field":"value"}}`,
			},
			wantErr: true,
		},
		{
			name:   "header_SetReportValue_error",
			fields: fields{},
			args: args{
				jsonString: `{"header":[{"field":"value"}]}`,
			},
			wantErr: true,
		},
		{
			name:   "details_SetReportValue_error",
			fields: fields{},
			args: args{
				jsonString: `{"details":[{"field":"value"}]}`,
			},
			wantErr: true,
		},
		{
			name:   "footer_SetReportValue_error",
			fields: fields{},
			args: args{
				jsonString: `{"footer":[{"field":"value"}]}`,
			},
			wantErr: true,
		},
		{
			name:   "data_type_error",
			fields: fields{},
			args: args{
				jsonString: `{"data":{"field":true}}`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpt := &Report{
				pdf:             tt.fields.pdf,
				orientation:     tt.fields.orientation,
				format:          tt.fields.format,
				fontDir:         tt.fields.fontDir,
				xmlHeader:       tt.fields.xmlHeader,
				xmlDetails:      tt.fields.xmlDetails,
				header:          tt.fields.header,
				details:         tt.fields.details,
				footer:          tt.fields.footer,
				data:            tt.fields.data,
				footerHeight:    tt.fields.footerHeight,
				pageBreak:       tt.fields.pageBreak,
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Creator:         tt.fields.Creator,
				Subject:         tt.fields.Subject,
				Keywords:        tt.fields.Keywords,
				LeftMargin:      tt.fields.LeftMargin,
				RightMargin:     tt.fields.RightMargin,
				TopMargin:       tt.fields.TopMargin,
				BottomMargin:    tt.fields.BottomMargin,
				FontFamily:      tt.fields.FontFamily,
				FontStyle:       tt.fields.FontStyle,
				FontSize:        tt.fields.FontSize,
				TextColor:       tt.fields.TextColor,
				BorderColor:     tt.fields.BorderColor,
				BackgroundColor: tt.fields.BackgroundColor,
				ImagePath:       tt.fields.ImagePath,
			}
			if err := rpt.LoadJSONDefinition(tt.args.jsonString); (err != nil) != tt.wantErr {
				t.Errorf("Report.LoadJSONDefinition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
