package report

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/twooffive"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

const (
	_generator       = "gopdf"
	_title           = "Nervatura Report"
	_margin          = float64(36.85)
	_fontFamily      = "Cabin"
	_fontStyle       = ""
	_fontSize        = float64(9)
	_textColor       = uint8(0)
	_borderColor     = uint8(0)
	_backgroundColor = uint8(255)
	_padding         = float64(6.4)
	_format          = "a4"
	_orientation     = "p"
	_unit            = "pt"
	_mmPt            = 2.83465
	_align           = "L"
	_fontDir         = ""
	_regValue        = "={{(\\S*?)[^}}]*}}.*?|={{.*? /}}"
)

//IM is a map[string]interface{} type short alias
type IM = map[string]interface{}

//SM is a map[string]string type short alias
type SM = map[string]string

// Generator the PDF generator interface
type Generator interface {
	Init(rpt *Report)
	// GetPageSize returns the current page's width and height.
	GetPageSize() (width, height float64)
	// PageNo returns the current page number.
	PageNo() int
	// AddPage adds a new page to the document
	AddPage()
	// AddImage draws a image
	AddImage(image *Image, x, y float64, options IM)
	// AddFont imports a font and makes it available
	AddFont(familyStr, styleStr, fileStr string, rd io.Reader)
	// GetFontSize returns the size of the current font in points.
	GetFontSize() (ptSize float64)
	// SetFont sets the font used to print character strings
	SetFont(familyStr, styleStr string, size float64)
	// SetFontSize defines the size of the current font.
	SetFontSize(size float64)
	// GetTextWidth returns the length of a string in user units.
	GetTextWidth(s string) float64
	// SetDrawColor defines the color used for all drawing operations
	SetDrawColor(r, g, b int)
	// SetFillColor defines the color used for all filling operations
	SetFillColor(r, g, b int)
	// SetTextColor defines the color used for text.
	SetTextColor(r, g, b int)
	// SetProperties - general report props. (title, author etc.)
	SetProperties(rpt *Report)
	// Text - Write prints text from the current position.
	Text(txtStr string, pageBreak float64)
	// Rect outputs a rectangle of width w and height h with the upper left corner positioned at point (x, y)
	Rect(x, y, w, h float64, styleStr string)
	// Line draws a line between points (x1, y1) and (x2, y2) using the current draw color, line width and cap style.
	Line(x1, y1, x2, y2 float64)
	// GetX returns the abscissa of the current position.
	GetX() float64
	// GetY returns the ordinate of the current position.
	GetY() float64
	// SetX defines the abscissa of the current position.
	SetX(x float64)
	// SetY : set current position y
	SetY(y float64)
	// SetXY defines the abscissa and ordinate of the current position.
	SetXY(x, y float64)
	// Ln performs a line break.
	Ln(h float64)
	// Cell prints a rectangular cell with optional borders, background color and character string.
	Cell(options IM)
	// MultiCell supports printing text with line breaks.
	MultiCell(options IM)

	// Save2Pdf creates a PDF output.
	Save2Pdf() ([]byte, error)
	// Save2PdfFile writes the PDF document to file
	Save2PdfFile(filename string) error
}

var generators = make(map[string]Generator)

func registerGenerators(name string, gen Generator) {
	generators[name] = gen
}

var propMap SM = SM{
	"height": "Height", "hgap": "HGap", "visible": "Visible", "name": "Name", "value": "Value", "width": "Width",
	"border": "Border", "align": "Align", "multiline": "Multiline", "font-style": "FontStyle", "fontstyle": "FontStyle",
	"font-size": "FontSize", "fontsize": "FontSize", "color": "TextColor", "textcolor": "TextColor",
	"border-color": "BorderColor", "bordercolor": "BorderColor",
	"background-color": "BackgroundColor", "backgroundcolor": "BackgroundColor",
	"src": "Src", "code-type": "CodeType", "codetype": "CodeType",
	"visible-value": "VisibleValue", "visiblevalue": "VisibleValue", "wide": "Width", "narrow": "Height",
	"extend": "Extend", "gap": "Gap", "fieldname": "Fieldname", "html": "Value", "databind": "Databind",
	"merge": "Merge", "header-background": "HeaderBackground", "headerbackground": "HeaderBackground",
	"footer-background": "FooterBackground", "footerbackground": "FooterBackground", "label": "Label",
	"header-align": "HeaderAlign", "headeralign": "HeaderAlign",
	"footer-align": "FooterAlign", "footeralign": "FooterAlign", "footer": "Footer",
	"title": "Title", "author": "Author", "creator": "Creator",
	"subject": "Subject", "keywords": "Keywords", "leftmargin": "LeftMargin", "left-margin": "LeftMargin",
	"topmargin": "TopMargin", "top-margin": "TopMargin", "rightmargin": "RightMargin", "right-margin": "RightMargin",
	"bottommargin": "BottomMargin", "bottom-margin": "BottomMargin",
	"imagepath": "ImagePath", "image-path": "ImagePath",
	"fontfamily": "FontFamily", "font-family": "FontFamily",
}

func invalidErr(etype, evalue string) string {
	return fmt.Sprintf("invalid %s element: %s", etype, evalue)
}

// PageItem - element interface wrapper
type PageItem struct {
	ItemType string
	Item     interface{}
}

func (pi *PageItem) setPageItem(fieldname string, value interface{}) error {
	vmap := map[string]map[string]func(value interface{}){
		"row": {
			"Height": func(value interface{}) {
				pi.Item.(*Row).Height = ut.ToFloat(value, 0)
			},
			"HGap": func(value interface{}) {
				pi.Item.(*Row).HGap = ut.ToFloat(value, 0)
			},
			"Visible": func(value interface{}) {
				pi.Item.(*Row).Visible = ut.ToString(value, "")
			},
		},
		"cell": {
			"Name": func(value interface{}) {
				pi.Item.(*Cell).Name = ut.ToString(value, "")
			},
			"Value": func(value interface{}) {
				pi.Item.(*Cell).Value = ut.ToString(value, "")
			},
			"Width": func(value interface{}) {
				pi.Item.(*Cell).Width = ut.ToString(value, "")
			},
			"Border": func(value interface{}) {
				pi.Item.(*Cell).Border = ut.ToString(value, "")
			},
			"Align": func(value interface{}) {
				pi.Item.(*Cell).Align = ut.ToString(value, "L")
			},
			"Multiline": func(value interface{}) {
				pi.Item.(*Cell).Multiline = ut.ToBoolean(value, false)
			},
			"FontStyle": func(value interface{}) {
				pi.Item.(*Cell).FontStyle = ut.ToString(value, "")
			},
			"FontSize": func(value interface{}) {
				pi.Item.(*Cell).FontSize = ut.ToFloat(value, pi.Item.(*Cell).FontSize)
			},
			"TextColor": func(value interface{}) {
				pi.Item.(*Cell).TextColor = ut.ToRGBA(value, pi.Item.(*Cell).TextColor)
			},
			"BorderColor": func(value interface{}) {
				pi.Item.(*Cell).BorderColor = ut.ToRGBA(value, pi.Item.(*Cell).BorderColor)
			},
			"BackgroundColor": func(value interface{}) {
				pi.Item.(*Cell).BackgroundColor = ut.ToRGBA(value, pi.Item.(*Cell).BackgroundColor)
			},
		},
		"image": {
			"Src": func(value interface{}) {
				pi.Item.(*Image).Src = ut.ToString(value, "")
			},
			"Height": func(value interface{}) {
				pi.Item.(*Image).Height = ut.ToFloat(value, 0)
			},
		},
		"barcode": {
			"CodeType": func(value interface{}) {
				pi.Item.(*Barcode).CodeType = ut.ToString(value, "")
			},
			"Value": func(value interface{}) {
				pi.Item.(*Barcode).Value = ut.ToString(value, "")
			},
			"VisibleValue": func(value interface{}) {
				pi.Item.(*Barcode).VisibleValue = ut.ToBoolean(value, false)
			},
			"Width": func(value interface{}) {
				pi.Item.(*Barcode).Width = ut.ToFloat(value, 0)
			},
			"Height": func(value interface{}) {
				pi.Item.(*Barcode).Height = ut.ToFloat(value, 0)
			},
			"Extend": func(value interface{}) {
				pi.Item.(*Barcode).Extend = ut.ToBoolean(value, false)
			},
		},
		"separator": {
			"Gap": func(value interface{}) {
				pi.Item.(*Separator).Gap = ut.ToFloat(value, 0)
			},
		},
		"vgap": {
			"Height": func(value interface{}) {
				pi.Item.(*VGap).Height = ut.ToFloat(value, 0)
			},
			"Visible": func(value interface{}) {
				// Deprecated
			},
		},
		"hline": {
			"Width": func(value interface{}) {
				pi.Item.(*HLine).Width = ut.ToString(value, "")
			},
			"Gap": func(value interface{}) {
				pi.Item.(*HLine).Gap = ut.ToFloat(value, 0)
			},
			"BorderColor": func(value interface{}) {
				pi.Item.(*HLine).BorderColor = ut.ToRGBA(value, pi.Item.(*HLine).BorderColor)
			},
			"Visible": func(value interface{}) {
				// Deprecated
			},
		},
		"html": {
			"Fieldname": func(value interface{}) {
				pi.Item.(*HTML).Fieldname = ut.ToString(value, "")
			},
			"Value": func(value interface{}) {
				pi.Item.(*HTML).Value = ut.ToString(value, "")
			},
		},
		"datagrid": {
			"Name": func(value interface{}) {
				pi.Item.(*Datagrid).Name = ut.ToString(value, "")
			},
			"Databind": func(value interface{}) {
				pi.Item.(*Datagrid).Databind = ut.ToString(value, "")
			},
			"Width": func(value interface{}) {
				pi.Item.(*Datagrid).Width = ut.ToString(value, "")
			},
			"Merge": func(value interface{}) {
				pi.Item.(*Datagrid).Merge = ut.ToBoolean(value, false)
			},
			"Border": func(value interface{}) {
				pi.Item.(*Datagrid).Border = ut.ToString(value, "1")
			},
			"FontSize": func(value interface{}) {
				pi.Item.(*Datagrid).FontSize = ut.ToFloat(value, pi.Item.(*Datagrid).FontSize)
			},
			"TextColor": func(value interface{}) {
				pi.Item.(*Datagrid).TextColor = ut.ToRGBA(value, pi.Item.(*Datagrid).TextColor)
			},
			"BorderColor": func(value interface{}) {
				pi.Item.(*Datagrid).BorderColor = ut.ToRGBA(value, pi.Item.(*Datagrid).BorderColor)
			},
			"BackgroundColor": func(value interface{}) {
				pi.Item.(*Datagrid).BackgroundColor = ut.ToRGBA(value, pi.Item.(*Datagrid).BackgroundColor)
			},
			"HeaderBackground": func(value interface{}) {
				pi.Item.(*Datagrid).HeaderBackground = ut.ToRGBA(value, pi.Item.(*Datagrid).HeaderBackground)
			},
			"FooterBackground": func(value interface{}) {
				pi.Item.(*Datagrid).FooterBackground = ut.ToRGBA(value, pi.Item.(*Datagrid).FooterBackground)
			},
		},
		"column": {
			"Fieldname": func(value interface{}) {
				pi.Item.(*Column).Fieldname = ut.ToString(value, "")
			},
			"Label": func(value interface{}) {
				pi.Item.(*Column).Label = ut.ToString(value, "")
			},
			"Width": func(value interface{}) {
				pi.Item.(*Column).Width = ut.ToString(value, "")
			},
			"Align": func(value interface{}) {
				pi.Item.(*Column).Align = ut.ToString(value, "L")
			},
			"HeaderAlign": func(value interface{}) {
				pi.Item.(*Column).HeaderAlign = ut.ToString(value, "L")
			},
			"FooterAlign": func(value interface{}) {
				pi.Item.(*Column).FooterAlign = ut.ToString(value, "L")
			},
			"Footer": func(value interface{}) {
				pi.Item.(*Column).Footer = ut.ToString(value, "")
			},
		},
	}

	if _, found := vmap[pi.ItemType][propMap[strings.ToLower(fieldname)]]; found {
		vmap[pi.ItemType][propMap[strings.ToLower(fieldname)]](value)
		return nil
	}
	return errors.New(invalidErr(pi.ItemType, fieldname))
}

func (rpt *Report) getPageItem(etype string) (PageItem, error) {
	switch etype {
	case "barcode":
		return PageItem{
			ItemType: etype,
			Item: &Barcode{
				CodeType: "CODE_39"}}, nil
	case "cell":
		return PageItem{
			ItemType: etype,
			Item: &Cell{
				Border:          "",
				Align:           "L",
				FontStyle:       rpt.FontStyle,
				FontSize:        rpt.FontSize,
				TextColor:       rpt.TextColor,
				BorderColor:     rpt.BorderColor,
				BackgroundColor: rpt.BackgroundColor}}, nil
	case "column":
		return PageItem{
			ItemType: etype,
			Item: &Column{
				Align:       "L",
				HeaderAlign: "L",
				FooterAlign: "L"}}, nil
	case "datagrid":
		return PageItem{
			ItemType: etype,
			Item: &Datagrid{
				Border:           "",
				FontSize:         rpt.FontSize,
				TextColor:        rpt.TextColor,
				BorderColor:      rpt.BorderColor,
				BackgroundColor:  rpt.BackgroundColor,
				HeaderBackground: rpt.BackgroundColor,
				FooterBackground: rpt.BackgroundColor,
				Columns:          make([]PageItem, 0)}}, nil
	case "hline":
		return PageItem{
			ItemType: etype,
			Item: &HLine{
				BorderColor: rpt.BorderColor}}, nil
	case "html":
		return PageItem{
			ItemType: etype,
			Item:     &HTML{}}, nil
	case "image":
		return PageItem{
			ItemType: etype,
			Item:     &Image{}}, nil
	case "row":
		return PageItem{
			ItemType: etype,
			Item: &Row{
				Columns: make([]PageItem, 0)}}, nil
	case "separator":
		return PageItem{
			ItemType: etype,
			Item:     &Separator{}}, nil
	case "vgap":
		return PageItem{
			ItemType: etype,
			Item:     &VGap{}}, nil
	}
	return PageItem{}, errors.New(invalidErr("Element", etype))
}

// Row - Horizontal logical group. The last element width extends up to the right margin.
type Row struct {
	Height  float64    `xml:"height,attr" json:"height"`   //row height
	HGap    float64    `xml:"hgap,attr" json:"hgap"`       //default gap between these two elements
	Visible string     `xml:"visible,attr" json:"visible"` //table data source name
	Columns []PageItem `xml:"columns,attr" json:"columns"` //Cell, Image, Barcode, Separator
}

// Cell - Row unit
type Cell struct {
	Name            string     `xml:"name,attr" json:"name"`                         //XML output node name
	Value           string     `xml:"value,attr" json:"value"`                       //static text or databind value
	Width           string     `xml:"width,attr" json:"width"`                       //number or percent value (e.g. "10" or "10%")
	Border          string     `xml:"border,attr" json:"border"`                     //values: "0"(no border, default), "1"(all) or some or all of the following characters: "L"(left), "T"(top), "R"(right),"B"(bottom)
	Align           string     `xml:"align,attr" json:"align"`                       //values: "L" (default) or "left", "R" or "right", "C" or "center"
	Multiline       bool       `xml:"multiline,attr" json:"multiline"`               //if true, print text with line breaks (default false)
	FontStyle       string     `xml:"font-style,attr" json:"font-style"`             //values: "" (default), "bold", "italic", "bolditalic"
	FontSize        float64    `xml:"font-size,attr" json:"font-size"`               //Default value: Report.FontSize
	TextColor       color.RGBA `xml:"color,attr" json:"color"`                       //JSON or XML value: in hexadecimal (e.g. #A0522D) or in decimal (e.g 10506797), default "black"
	BorderColor     color.RGBA `xml:"border-color,attr" json:"border-color"`         //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
	BackgroundColor color.RGBA `xml:"background-color,attr" json:"background-color"` //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
}

// Image - Row unit
type Image struct {
	Src       string  `xml:"src,attr" json:"src"` //JPEG or PNG image file path and name (e.g. "test/logo.jpg") or image data
	Data      []byte  `xml:"data,attr" json:"data"`
	MaxWidth  float64 `xml:"max-width,attr" json:"max-width"`
	MaxHeight float64 `xml:"max-height,attr" json:"max-height"`
	Height    float64 `xml:"height,attr" json:"height"` //image height (default height of parent Row).
	Width     float64 `xml:"width,attr" json:"width"`   //image width will be calculated from the height dimension so that the aspect ratio is maintained.
}

// Barcode - Row unit
type Barcode struct {
	CodeType     string  `xml:"code-type,attr" json:"code-type"`         //Values: "CODE_39"/"code39", "ITF"/"i2of5", "CODE_128"/"code128", "EAN"/"ean", "QR"/"qr"
	Value        string  `xml:"value,attr" json:"value"`                 //barcode text value
	VisibleValue bool    `xml:"visible-value,attr" json:"visible-value"` //show or not the value of text
	Width        float64 `xml:"wide,attr" json:"wide"`                   //barcode width (default width of the value string + padding)
	Height       float64 `xml:"narrow,attr" json:"narrow"`               //barcode height (default 10).
	Extend       bool    `xml:"extend,attr" json:"extend"`               //barcode width extends up to the right margin (default false)
}

// Separator - Row unit, A horizontal separator line.
type Separator struct {
	Gap float64 `xml:"gap,attr" json:"gap"` //distance size
}

// VGap - a vertical gap.
type VGap struct {
	Height float64 `xml:"height,attr" json:"height"` //distance size
}

// HLine - a horizontal line.
type HLine struct {
	Width       string     `xml:"width,attr" json:"width"`               //number or percent value (e.g. "10" or "10%")
	Gap         float64    `xml:"gap,attr" json:"gap"`                   // greater than 0 then double line
	BorderColor color.RGBA `xml:"border-color,attr" json:"border-color"` //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
}

// HTML - a basic HTML elements rendering. It supports
// only hyperlinks and bold, italic and underscore attributes.
type HTML struct {
	Fieldname string `xml:"fieldname,attr" json:"fieldname"` //databind fieldname
	Value     string `xml:",cdata" json:"html"`              //html text
}

// Datagrid - Create a table from a data list.
type Datagrid struct {
	Name             string     `xml:"name,attr" json:"name"`                           //XML output node name
	Databind         string     `xml:"databind,attr" json:"databind"`                   //table data source name
	Width            string     `xml:"width,attr" json:"width"`                         //number or percent value (e.g. "10" or "10%")
	Merge            bool       `xml:"merge,attr" json:"merge"`                         //if true then all fields will be displayed in a single column (default false)
	Border           string     `xml:"border,attr" json:"border"`                       //values: "0"(no border, default), "1"(all) or some or all of the following characters: "L"(left), "T"(top), "R"(right),"B"(bottom)
	FontSize         float64    `xml:"font-size,attr" json:"font-size"`                 //Default value: Report.FontSize
	TextColor        color.RGBA `xml:"color,attr" json:"color"`                         //JSON or XML value: in hexadecimal (e.g. #A0522D) or in decimal (e.g 10506797), default "black"
	BorderColor      color.RGBA `xml:"border-color,attr" json:"border-color"`           //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
	BackgroundColor  color.RGBA `xml:"background-color,attr" json:"background-color"`   //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
	HeaderBackground color.RGBA `xml:"header-background,attr" json:"header-background"` //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
	FooterBackground color.RGBA `xml:"footer-background,attr" json:"footer-background"` //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
	Columns          []PageItem `xml:"columns" json:"columns"`                          //columns list of the datagrid
}

// Column - Datagrid unit
type Column struct {
	Fieldname   string `xml:"fieldname,attr" json:"fieldname"`       //datasource dictonary key (special value: "counter")
	Label       string `xml:"label,attr" json:"label"`               //Column caption
	Width       string `xml:"width,attr" json:"width"`               //number or percent value (e.g. "10" or "10%")
	Align       string `xml:"align,attr" json:"align"`               //values: "L" (default) or "left", "R" or "right", "C" or "center"
	HeaderAlign string `xml:"header-align,attr" json:"header-align"` //values: "L" (default) or "left", "R" or "right", "C" or "center"
	FooterAlign string `xml:"footer-align,attr" json:"footer-align"` //values: "L" (default) or "left", "R" or "right", "C" or "center"
	Footer      string `xml:"footer,attr" json:"footer"`             //static text or databind value
}

// Report is the principal structure for creating a single PDF document
type Report struct {
	pdf                                                 Generator
	orientation, format, fontDir, xmlHeader, xmlDetails string
	//header/footer elements: Row, VGap, HLine. Page elements: Row, VGap, HLine, HTML, Datagrid
	header, details, footer []PageItem
	//Valid datasource types: string or map[string]string (dictonary) or []map[string]string (record list)
	data                    IM
	footerHeight, pageBreak float64
	Title                   string     `xml:"title,attr" json:"title"`
	Author                  string     `xml:"author,attr" json:"author"`
	Creator                 string     `xml:"creator,attr" json:"creator"`
	Subject                 string     `xml:"subject,attr" json:"subject"`
	Keywords                string     `xml:"keywords,attr" json:"keywords"`
	LeftMargin              float64    `xml:"left-margin,attr" json:"left-margin"`
	RightMargin             float64    `xml:"right-margin,attr" json:"right-margin"`
	TopMargin               float64    `xml:"top-margin,attr" json:"top-margin"`
	BottomMargin            float64    `xml:"bottom-margin,attr" json:"bottom-margin"`
	FontFamily              string     `xml:"font-family,attr" json:"font-family"`           //values: "times"(default), "helvetica", "courier" or custom font
	FontStyle               string     `xml:"font-style,attr" json:"font-style"`             //values: "" (default), "bold", "italic", "bolditalic"
	FontSize                float64    `xml:"font-size,attr" json:"font-size"`               //Default value: 10
	TextColor               color.RGBA `xml:"color,attr" json:"color"`                       //JSON or XML value: in hexadecimal (e.g. #A0522D) or in decimal (e.g 10506797), default "black"
	BorderColor             color.RGBA `xml:"border-color,attr" json:"border-color"`         //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
	BackgroundColor         color.RGBA `xml:"background-color,attr" json:"background-color"` //JSON or XML value: integer gray color (in range from 0 "black" to 255 "white"), default "black"
	ImagePath               string     `xml:"image-path,attr" json:"image-path"`
}

// SetReportValue - You can set the Report properties safely and type independent.
func (rpt *Report) SetReportValue(fieldname string, value interface{}) error {

	vmap := map[string]func(value interface{}){
		"Title": func(value interface{}) {
			rpt.Title = ut.ToString(value, rpt.Title)
		},
		"Author": func(value interface{}) {
			rpt.Author = ut.ToString(value, rpt.Author)
		},
		"Creator": func(value interface{}) {
			rpt.Creator = ut.ToString(value, rpt.Creator)
		},
		"Subject": func(value interface{}) {
			rpt.Subject = ut.ToString(value, rpt.Subject)
		},
		"Keywords": func(value interface{}) {
			rpt.Keywords = ut.ToString(value, rpt.Keywords)
		},
		"LeftMargin": func(value interface{}) {
			rpt.LeftMargin = ut.ToFloat(value, rpt.LeftMargin) * _mmPt
		},
		"TopMargin": func(value interface{}) {
			rpt.TopMargin = ut.ToFloat(value, rpt.TopMargin) * _mmPt
		},
		"RightMargin": func(value interface{}) {
			rpt.RightMargin = ut.ToFloat(value, rpt.RightMargin) * _mmPt
		},
		"BottomMargin": func(value interface{}) {
			rpt.BottomMargin = ut.ToFloat(value, rpt.BottomMargin) * _mmPt
		},
		"FontStyle": func(value interface{}) {
			rpt.FontStyle = rpt.parseValue("FontStyle", value).(string)
		},
		"FontSize": func(value interface{}) {
			rpt.FontSize = ut.ToFloat(value, rpt.FontSize)
		},
		"TextColor": func(value interface{}) {
			rpt.TextColor = ut.ToRGBA(value, rpt.TextColor)
		},
		"BorderColor": func(value interface{}) {
			rpt.BorderColor = ut.ToRGBA(value, rpt.BorderColor)
		},
		"BackgroundColor": func(value interface{}) {
			rpt.BackgroundColor = ut.ToRGBA(value, rpt.BackgroundColor)
		},
		"ImagePath": func(value interface{}) {
			rpt.ImagePath = ut.ToString(value, rpt.ImagePath)
		},
	}

	if _, found := vmap[propMap[strings.ToLower(fieldname)]]; found {
		vmap[propMap[strings.ToLower(fieldname)]](value)
		return nil
	}
	err := fmt.Sprintf("missing report fieldname: %s", fieldname)
	return errors.New(err)
}

func (rpt *Report) createHeaderAndFooter() {
	createSection := func(section string, elements []PageItem) {
		for index := 0; index < len(elements); index++ {
			switch elements[index].Item.(type) {
			case *Row, *VGap, *HLine:
				rpt.createElement(section, elements[index].Item)
			}
		}
	}
	createSection("header", rpt.header)
	cx := rpt.pdf.GetX()
	cy := rpt.pdf.GetY()
	_, pageHeight := rpt.pdf.GetPageSize()
	rpt.pdf.SetY(pageHeight - rpt.BottomMargin - rpt.footerHeight)
	createSection("footer", rpt.footer)
	rpt.pdf.SetXY(cx, cy)
}

func (rpt *Report) checkPageBreak(nextHeight float64) bool {
	cy := rpt.pdf.GetY()
	dLine := rpt.pageBreak
	if cy < rpt.pageBreak-rpt.footerHeight {
		dLine -= rpt.footerHeight
	}
	if cy+nextHeight > dLine {
		return true
	}
	return false
}

func (rpt *Report) getFooterHeight() (fHeight float64) {
	for index := 0; index < len(rpt.footer); index++ {
		switch v := rpt.footer[index].Item.(type) {
		case *Row:
			fHeight += rpt.createRow("footer", v, true)
		case *VGap:
			fHeight += v.Height
		case *HLine:
			fHeight += (1 + v.Gap)
		}
	}
	_, pageHeight := rpt.pdf.GetPageSize()
	rpt.pageBreak = pageHeight - rpt.BottomMargin
	return fHeight
}

func (rpt *Report) setHTMLValue(value, fieldname string) string {
	r, _ := regexp.Compile(_regValue)
	if r.MatchString(value) {
		valueKey := r.FindString(value)
		valueF := rpt.setValue(valueKey)
		value = strings.Replace(value, valueKey, valueF, -1)
		rpt.addToXML("details", []string{fieldname, valueF, fieldname})
		if r.MatchString(value) {
			value = rpt.setHTMLValue(value, fieldname)
		}
		return value
	}
	return value
}

func (rpt *Report) setValue(value string) string {
	var getValue = func(valueGet string) string {
		if matched, _ := regexp.MatchString("{{page}}", valueGet); matched {
			valueGet = strings.ReplaceAll(valueGet, "{{page}}", strconv.Itoa(rpt.pdf.PageNo()))
		}
		dbv := strings.Split(valueGet, ".")
		storeData, isData := rpt.data[dbv[0]]
		if isData {
			if data, valid := storeData.([]SM); valid {
				if len(dbv) > 2 {
					row := ut.ToInteger(dbv[1], 0)
					if len(data) > int(row) {
						rowValue, isData := data[row][dbv[2]]
						if isData {
							return rowValue
						}
					}
					return ""
				}
				return valueGet
			}
			if data, valid := storeData.(SM); valid {
				if len(dbv) > 1 {
					dictValue, isData := data[dbv[1]]
					if isData {
						return dictValue
					}
					return ""
				}
				return valueGet
			}
			if data, valid := storeData.(string); valid {
				return data
			}
		}
		return valueGet
	}
	if matched, _ := regexp.MatchString(_regValue, value); matched {
		valueSet := value[strings.Index(value, "={{")+3 : strings.Index(value, "}}")]
		value = strings.Replace(value, "={{"+valueSet+"}}", getValue(valueSet), strings.Index(value, "}}")+2)
		if matched, _ := regexp.MatchString(_regValue, value); matched {
			return rpt.setValue(value)
		}
		return value
	}
	return getValue(value)
}

func (rpt *Report) getCellHeight(text string, width float64, options IM) float64 {
	if text == "" {
		text = "X"
	}
	rpt.pdf.SetFont(rpt.FontFamily, options["fontStyle"].(string), options["fontSize"].(float64))
	lines := rpt.wrapTextLines(text, width-_padding)
	lineHt := rpt.pdf.GetFontSize()
	return float64(len(lines)) * (lineHt + _padding)
}

func (rpt *Report) createGridHeader(headerOptions IM) {
	headerOptions["height"] = float64(0)
	for colIndex := 0; colIndex < len(headerOptions["columns"].([]IM)); colIndex++ {
		column := headerOptions["columns"].([]IM)[colIndex]
		if !headerOptions["merge"].(bool) {
			headerOptions["text"] = column["label"]
			headerOptions["ln"] = column["ln"]
			if column["columnWidth"] == 0 {
				column["columnWidth"] = (headerOptions["gridWidth"].(float64) - headerOptions["columnsWidth"].(float64)) / float64(len(headerOptions["columns"].([]IM)))
			}
			headerOptions["columnWidth"] = column["columnWidth"]
		} else {
			headerOptions["text"] = headerOptions["text"].(string) + " " + column["label"].(string)
		}
		headerOptions["align"] = column["headerAlign"]
		rpt.createCell(headerOptions)
	}
}

func (rpt *Report) createDatagrid(gridElement *Datagrid, virtual bool) bool {
	if len(gridElement.Columns) == 0 {
		return false
	}
	rows, valid := rpt.data[gridElement.Databind].([]SM)
	if !valid || len(rows) == 0 {
		return false
	}
	gridOptions := IM{
		"xname":           ut.ToString(gridElement.Name, "items"),
		"border":          ut.ToString(gridElement.Border, "1"),
		"fontFamily":      rpt.FontFamily,
		"fontStyle":       rpt.FontStyle,
		"fontSize":        gridElement.FontSize,
		"textColor":       gridElement.TextColor,
		"borderColor":     gridElement.BorderColor,
		"backgroundColor": gridElement.BackgroundColor,
		"virtual":         virtual,
	}
	headerOptions := IM{
		"fontSize": gridOptions["fontSize"], "textColor": gridOptions["textColor"],
		"borderColor": gridOptions["borderColor"], "border": gridOptions["border"],
		"backgroundColor": ut.ToRGBA(gridElement.HeaderBackground, gridElement.BackgroundColor),
		"merge":           ut.ToBoolean(gridElement.Merge, false),
		"fontFamily":      gridOptions["fontFamily"], "fontStyle": "B", "text": "", "height": float64(0),
		"columns":      make([]IM, 0),
		"columnsWidth": float64(0), "gridWidth": float64(0), "multiline": false, "virtual": virtual}

	pageWidth, _ := rpt.pdf.GetPageSize()
	nwidth := pageWidth - rpt.RightMargin - rpt.LeftMargin
	gridElement.Width = ut.ToString(gridElement.Width, "100%")
	if strings.HasSuffix(gridElement.Width, "%") {
		headerOptions["gridWidth"] = nwidth * ut.ToFloat(strings.Replace(gridElement.Width, "%", "", -1), 0) / 100
	} else {
		headerOptions["gridWidth"] = ut.ToFloat(gridElement.Width, 0)
		if headerOptions["gridWidth"].(float64) > nwidth {
			headerOptions["gridWidth"] = nwidth
		}
	}
	headerOptions["extend"] = (headerOptions["gridWidth"] == nwidth)
	gridOptions["extend"] = headerOptions["extend"]

	footerOptions := IM{
		"fontSize": gridOptions["fontSize"], "textColor": gridOptions["textColor"],
		"borderColor": gridOptions["borderColor"], "border": gridOptions["border"],
		"fontFamily": gridOptions["fontFamily"], "fontStyle": "B", "text": "", "height": float64(0),
		"backgroundColor": ut.ToRGBA(gridElement.FooterBackground, gridElement.BackgroundColor),
		"extend":          headerOptions["extend"], "multiline": false, "virtual": virtual}

	zcol, footers, footerWidth := 0, make([]IM, 0), float64(0)
	xCol := rpt.LeftMargin
	lnWidth := headerOptions["gridWidth"].(float64)
	for index := 0; index < len(gridElement.Columns); index++ {
		column := gridElement.Columns[index].Item.(*Column)
		if headerOptions["columnsWidth"].(float64) >= headerOptions["gridWidth"].(float64) {
			return false
		}
		columnOptions := IM{
			"fontFamily": gridOptions["fontFamily"], "fontStyle": gridOptions["fontStyle"],
			"fontSize": gridOptions["fontSize"], "textColor": gridOptions["textColor"],
			"borderColor": gridOptions["borderColor"], "border": gridOptions["border"],
			"fieldname": column.Fieldname, "multiline": true,
			"label":       rpt.setValue(column.Label),
			"columnWidth": float64(0)}
		if !headerOptions["merge"].(bool) {
			columnWidth := ut.ToString(column.Width, "")
			if columnWidth != "" {
				if strings.HasSuffix(columnWidth, "%") {
					columnOptions["columnWidth"] = headerOptions["gridWidth"].(float64) * ut.ToFloat(strings.Replace(columnWidth, "%", "", -1), 0) / 100
				} else {
					columnOptions["columnWidth"] = ut.ToFloat(columnWidth, 0)
				}
			} else {
				if len(gridElement.Columns)-1 == index {
					columnOptions["columnWidth"] = lnWidth
				} else {
					columnOptions["columnWidth"] = rpt.pdf.GetTextWidth(columnOptions["label"].(string)) + _padding
				}
				zcol++
			}
			if headerOptions["columnsWidth"].(float64)+columnOptions["columnWidth"].(float64) >= headerOptions["gridWidth"].(float64) {
				columnOptions["columnWidth"] = headerOptions["gridWidth"].(float64) - headerOptions["columnsWidth"].(float64)
			}
			headerOptions["columnsWidth"] = headerOptions["columnsWidth"].(float64) + columnOptions["columnWidth"].(float64)
			lnWidth = lnWidth - columnOptions["columnWidth"].(float64)
			columnOptions["xCol"] = xCol
			xCol += columnOptions["columnWidth"].(float64)
			cheight := rpt.getCellHeight(columnOptions["label"].(string), columnOptions["columnWidth"].(float64), columnOptions)
			if cheight > headerOptions["height"].(float64) {
				headerOptions["height"] = cheight
			}
		}
		columnOptions["headerAlign"] = ut.ToString(column.HeaderAlign, "L")
		columnOptions["align"] = ut.ToString(column.Align, "L")

		footerValue := rpt.setValue(ut.ToString(column.Footer, ""))
		footerAlign := ut.ToString(column.FooterAlign, "L")
		if footerValue != "" {
			if len(footers) == 0 {
				footers = append(footers, IM{
					"text": footerValue, "align": footerAlign,
					"columnWidth": footerWidth + columnOptions["columnWidth"].(float64)})
			} else {
				footers[len(footers)-1]["columnWidth"] = footers[len(footers)-1]["columnWidth"].(float64) + footerWidth
				footers = append(footers, IM{
					"text": footerValue, "align": footerAlign, "columnWidth": columnOptions["columnWidth"]})
			}
			footerWidth = 0
		} else {
			footerWidth = footerWidth + columnOptions["columnWidth"].(float64)
		}
		if len(gridElement.Columns)-1 == index {
			columnOptions["ln"] = 1
		} else {
			columnOptions["ln"] = 0
		}
		headerOptions["columns"] = append(headerOptions["columns"].([]IM), columnOptions)
	}
	if !headerOptions["merge"].(bool) {
		rpt.createGridHeader(headerOptions)
	}

	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]
		rpt.addToXML("details", []string{gridOptions["xname"].(string)})
		gridOptions["height"] = float64(0)
		gridOptions["text"] = ""
		for colIndex := 0; colIndex < len(headerOptions["columns"].([]IM)); colIndex++ {
			column := headerOptions["columns"].([]IM)[colIndex]
			if column["fieldname"] == "counter" {
				column["text"] = strconv.Itoa(rowIndex + 1)
			} else {
				if value, found := row[column["fieldname"].(string)]; found {
					column["text"] = value
				} else {
					column["text"] = ""
				}
			}
			if !headerOptions["merge"].(bool) {
				cheight := rpt.getCellHeight(column["text"].(string), column["columnWidth"].(float64), gridOptions)
				if cheight > gridOptions["height"].(float64) {
					gridOptions["height"] = cheight
				}
			} else {
				gridOptions["text"] = gridOptions["text"].(string) + " " + column["text"].(string)
				rpt.addToXML("details", []string{column["fieldname"].(string), column["text"].(string), column["fieldname"].(string)})
			}
		}
		if rpt.checkPageBreak(gridOptions["height"].(float64)) {
			rpt.addPage()
			if !headerOptions["merge"].(bool) {
				rpt.createGridHeader(headerOptions)
			}
		}
		if !headerOptions["merge"].(bool) {
			for colIndex := 0; colIndex < len(headerOptions["columns"].([]IM)); colIndex++ {
				column := headerOptions["columns"].([]IM)[colIndex]
				gridOptions["text"] = column["text"]
				gridOptions["columnWidth"] = column["columnWidth"]
				gridOptions["xCol"] = column["xCol"]
				gridOptions["align"] = column["align"]
				gridOptions["ln"] = column["ln"]
				gridOptions["multiline"] = column["multiline"]
				rpt.createCell(gridOptions)
				rpt.addToXML("details", []string{column["fieldname"].(string), column["text"].(string), column["fieldname"].(string)})
			}
		} else {
			gridOptions["text"] = strings.Trim(gridOptions["text"].(string), " ")
			gridOptions["columnWidth"] = float64(0)
			gridOptions["ln"] = 1
			gridOptions["multiline"] = true
			gridOptions["height"] = float64(0)
			rpt.createCell(gridOptions)
		}
		rpt.addToXML("details", []string{gridOptions["xname"].(string)})
	}
	if !headerOptions["merge"].(bool) {
		for colIndex := 0; colIndex < len(footers); colIndex++ {
			column := footers[colIndex]
			cheight := rpt.getCellHeight(column["text"].(string), column["columnWidth"].(float64), footerOptions)
			if cheight > footerOptions["height"].(float64) {
				footerOptions["height"] = cheight
			}
		}
		for colIndex := 0; colIndex < len(footers); colIndex++ {
			column := footers[colIndex]
			footerOptions["text"] = column["text"]
			footerOptions["columnWidth"] = column["columnWidth"]
			footerOptions["align"] = column["align"]
			if len(footers)-1 == colIndex {
				footerOptions["ln"] = 1
			} else {
				footerOptions["ln"] = 0
			}
			rpt.createCell(footerOptions)
			rpt.addToXML("footer", []string{gridOptions["xname"].(string), column["text"].(string), gridOptions["xname"].(string)})
		}
	}

	return true
}

func (rpt *Report) createCell(options IM) float64 {
	//x, y, width, height, text, border, ln, align, padding, multiline, extend
	rpt.setPageStyle(options)
	virtual := ut.ToBoolean(options["virtual"], false)
	text := ut.ToString(options["text"], "")
	multiline := ut.ToBoolean(options["multiline"], false)
	ln := ut.ToBoolean(options["ln"], false)
	padding := _padding
	border := ut.ToString(options["border"], "")
	align := ut.ToString(options["align"], "L")
	fill := false
	backgroundColor := ut.ToRGBA(options["backgroundColor"], rpt.BackgroundColor)
	if backgroundColor != rpt.BackgroundColor {
		fill = true
	}
	if !fill && border == "" && !multiline {
		padding = 0
	}

	pageWidth, _ := rpt.pdf.GetPageSize()
	lineHt := rpt.pdf.GetFontSize()
	startY := rpt.pdf.GetY()
	startX := rpt.pdf.GetX()

	xCol := ut.ToFloat(options["xCol"], startX)
	if xCol != startX {
		rpt.pdf.SetX(xCol)
	}

	width := ut.ToFloat(options["columnWidth"], 0)
	if ln && ut.ToBoolean(options["extend"], false) {
		width = pageWidth - rpt.RightMargin - xCol
	} else if width == 0 {
		widthStr := ut.ToString(options["width"], "")
		if strings.HasSuffix(widthStr, "%") {
			width = ut.ToFloat(strings.Replace(widthStr, "%", "", -1), width) / 100
			width = (pageWidth - rpt.LeftMargin - rpt.RightMargin) * width
		} else {
			width = ut.ToFloat(widthStr, width)
			if width > 0 {
				width += padding
			}
		}
		if width == 0 {
			width = rpt.pdf.GetTextWidth(text) + padding
		}
		if startX+padding+width > pageWidth-rpt.RightMargin {
			width = 0
		}
	}

	height := ut.ToFloat(options["height"], 0)

	if multiline {
		if height == 0 {
			height = rpt.getCellHeight(text, width, options)
		}
		if !virtual {
			//w, h, lineH, padding float64, txtStr, borderStr, alignStr string, fill bool
			rpt.pdf.MultiCell(IM{
				"w": width, "h": height, "lineH": lineHt + padding, "padding": padding,
				"txtStr": text, "borderStr": border, "alignStr": align, "fill": fill,
			})
		}
		if ln {
			rpt.pdf.SetY(startY + height)
		} else {
			rpt.pdf.SetY(startY)
		}
		return height
	}

	if lineHt+padding > height {
		height = lineHt + padding
	}
	if !virtual {
		rpt.pdf.Cell(IM{
			"w": width, "h": height, "padding": padding, "txtStr": text, "borderStr": border,
			"alignStr": align, "fill": fill, "ln": ln,
		})
	}
	if rpt.pdf.GetY()-startY > height {
		return rpt.pdf.GetY() - startY
	}
	return height
}

// encodeImage - base64 -> []byte
func (rpt *Report) encodeImage(data string, v *Image) {
	rawImage := string(data)[strings.Index(string(data), ",")+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(rawImage))
	m, iFormat, err := image.Decode(reader)
	if err == nil {
		v.MaxWidth = float64(m.Bounds().Max.X)
		v.MaxHeight = float64(m.Bounds().Max.Y)
		buf := new(bytes.Buffer)
		switch iFormat {
		case "jpeg":
			err = jpeg.Encode(buf, m, &jpeg.Options{Quality: 85})
		case "png":
			err = png.Encode(buf, m)
		}
		if err == nil && len(buf.Bytes()) > 0 {
			v.Data = buf.Bytes()
		}
	}
}

// setImageSize - Decode from reader to image format
func (rpt *Report) setImageSize(v *Image) {
	src := v.Src
	if rpt.ImagePath != "" {
		src = path.Join(rpt.ImagePath, v.Src)
	}
	reader, err := os.Open(filepath.Clean(src))
	if err == nil {
		m, _, err := image.Decode(reader)
		if err == nil {
			v.MaxWidth = float64(m.Bounds().Max.X)
			v.MaxHeight = float64(m.Bounds().Max.Y)
		}
	} else {
		v.Src = ""
	}
	defer func() {
		if rb_err := reader.Close(); rb_err != nil {
			return
		}
	}()
}

func (rpt *Report) drawImage(v *Image) {
	if v.Width == 0 {
		v.Width = v.Height * (v.MaxWidth / v.MaxHeight)
	}
	if rpt.checkPageBreak(v.Height) {
		rpt.addPage()
	}
	rpt.pdf.AddImage(v, rpt.pdf.GetX(), rpt.pdf.GetY(), IM{"ImagePath": rpt.ImagePath})
}

func (rpt *Report) createImage(v *Image, rowHeight float64, virtual bool) (float64, float64) {
	data := rpt.setValue(v.Src)
	if strings.HasPrefix(data, "data:image") && v.Data == nil {
		rpt.encodeImage(data, v)
	}
	if v.Data == nil && v.Width == 0 {
		rpt.setImageSize(v)
	}
	if v.Data != nil || v.Src != "" {
		height := float64(-1)
		if height <= 0 && v.Height > 0 {
			height = v.Height
			if height > rowHeight {
				rowHeight = height
			}
		}
		if height <= 0 && rowHeight > 0 {
			height = rowHeight - _padding/3
		}
		v.Height = height
		if !virtual {
			rpt.drawImage(v)
		}
		if height > rowHeight || rowHeight == 0 {
			rowHeight = height
		}
	}
	return rowHeight, v.Width
}

func (rpt *Report) createBarcode(v *Barcode, virtual, ln bool) (float64, float64) {
	pageWidth, _ := rpt.pdf.GetPageSize()
	rpt.pdf.SetTextColor(int(rpt.TextColor.R), int(rpt.TextColor.G), int(rpt.TextColor.B))
	width := v.Width
	strWidth := rpt.pdf.GetTextWidth(v.Value)
	if width == 0 {
		width = strWidth + 1.5*_padding
	}
	height := v.Height
	if height == 0 {
		height = 10 * _mmPt
	}
	startX := rpt.pdf.GetX()
	startY := rpt.pdf.GetY()
	lineHt := rpt.pdf.GetFontSize()
	if ln {
		if v.Extend {
			width = pageWidth - startX - rpt.RightMargin - _padding
		}
	}
	if rpt.checkPageBreak(height) && !virtual {
		rpt.addPage()
	}
	var bcode barcode.Barcode
	switch v.CodeType {
	case "CODE_39", "code39":
		bcode, _ = code39.Encode(v.Value, true, true)

	case "ITF", "i2of5":
		bcode, _ = twooffive.Encode(v.Value, true)

	case "CODE_128", "code128":
		bcode, _ = code128.Encode(v.Value)

	case "EAN", "ean":
		bcode, _ = ean.Encode(v.Value)

	case "QR", "qr":
		width = height
		bcode, _ = qr.Encode(v.Value, qr.H, qr.Unicode)

	}

	if bcode != nil {
		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, bcode, &jpeg.Options{Quality: 100})
		if err == nil {
			rpt.pdf.AddImage(&Image{Data: buf.Bytes(), Width: width, Height: height}, startX, startY, IM{})
		}
		if v.VisibleValue {
			rpt.pdf.SetXY(startX+(width-strWidth)/2, startY+height+1.5*_padding)
			rpt.pdf.Text(v.Value, rpt.pageBreak-rpt.footerHeight)
			height += lineHt + 1.5*_padding
		}
	}
	return height, width
}

func (rpt *Report) addToXML(section string, values []string) {
	if values[0] != "label" {
		switch section {
		case "header":
			rpt.xmlHeader += fmt.Sprintf("\n    <%s><![CDATA[%s]]></%s>", values[0], values[1], values[2])
		case "details":
			if len(values) == 1 {
				rpt.xmlDetails += fmt.Sprintf("\n    <%s>", values[0])
				return
			}
			rpt.xmlDetails += fmt.Sprintf("\n    <%s><![CDATA[%s]]></%s>", values[0], values[1], values[2])
		case "footer":
			rpt.xmlDetails += fmt.Sprintf("\n    <%s_footer"+"><![CDATA[%s]]></%s_footer"+">",
				values[0], values[1], values[2])
		}
	}
}

func (rpt *Report) createRow(section string, rowElement *Row, virtual bool) float64 {
	maxHeight := rowElement.Height
	for index := 0; index < len(rowElement.Columns); index++ {
		startY := rpt.pdf.GetY()
		if rpt.pdf.GetX() != rpt.LeftMargin {
			rpt.pdf.SetX(rpt.pdf.GetX() + rowElement.HGap)
		}
		startX := rpt.pdf.GetX()
		ln := len(rowElement.Columns)-1 == index
		element := rowElement.Columns[index].Item
		switch v := element.(type) {
		case *Cell:
			options := IM{
				"height":          maxHeight,
				"fontFamily":      rpt.FontFamily,
				"fontStyle":       v.FontStyle,
				"fontSize":        v.FontSize,
				"textColor":       v.TextColor,
				"borderColor":     v.BorderColor,
				"backgroundColor": v.BackgroundColor,
				"text":            rpt.setValue(v.Value),
				"width":           v.Width,
				"border":          v.Border,
				"align":           v.Align,
				"multiline":       false,
				"extend":          true,
				"virtual":         virtual,
			}
			options["ln"] = ln
			if section == "details" {
				options["multiline"] = v.Multiline
				if v.Multiline {
					options["height"] = 0
				}
			}
			cellHeight := rpt.createCell(options)
			if cellHeight > maxHeight || maxHeight == 0 {
				maxHeight = cellHeight
			}

			var xname = ut.ToString(v.Name, "head")
			rpt.addToXML(section, []string{xname, options["text"].(string), xname})

		case *Image:
			if v.Src != "" {
				height, width := rpt.createImage(v, maxHeight, virtual)
				if height > maxHeight {
					maxHeight = height
				}
				if len(rowElement.Columns)-1 == index {
					rpt.pdf.SetXY(rpt.LeftMargin, startY+maxHeight)
				} else {
					rpt.pdf.SetXY(startX+width, startY)
				}
			}
		case *Barcode:
			height, width := rpt.createBarcode(v, virtual, ln)
			if height > maxHeight || maxHeight == 0 {
				maxHeight = height
			}
			if len(rowElement.Columns)-1 == index {
				rpt.pdf.SetXY(rpt.LeftMargin, startY+maxHeight)
			} else {
				rpt.pdf.SetXY(startX+width+_padding, startY)
			}
		case *Separator:
			if !virtual {
				rpt.pdf.Line(rpt.pdf.GetX()+v.Gap, rpt.pdf.GetY(), rpt.pdf.GetX()+v.Gap, rpt.pdf.GetY()+maxHeight)
			}
			if len(rowElement.Columns)-1 == index {
				rpt.pdf.SetX(rpt.pdf.GetX() + v.Gap)
			}
			if v.Gap > maxHeight || maxHeight == 0 {
				maxHeight = v.Gap
			}
		}
	}
	return maxHeight
}

func (rpt *Report) createHTML(v *HTML) {
	lineHt := rpt.pdf.GetFontSize()
	htmlStr := v.Value
	fieldname := ut.ToString(v.Fieldname, "head")
	options := IM{
		"fontFamily":      rpt.FontFamily,
		"fontStyle":       rpt.FontStyle,
		"fontSize":        rpt.FontSize,
		"textColor":       rpt.TextColor,
		"borderColor":     rpt.BorderColor,
		"backgroundColor": rpt.BackgroundColor}
	htmlStr = rpt.setHTMLValue(htmlStr, fieldname)
	rpt.setPageStyle(options)
	rpt.writeHTML(lineHt, htmlStr)
	rpt.pdf.SetXY(rpt.LeftMargin, rpt.pdf.GetY()+lineHt+_padding)
}

func (rpt *Report) createLine(v *HLine, virtual bool) {
	width := float64(0)
	pageWidth, _ := rpt.pdf.GetPageSize()
	if strings.HasSuffix(v.Width, "%") {
		width = ut.ToFloat(strings.Replace(v.Width, "%", "", -1), width) / 100
		width = (pageWidth - rpt.LeftMargin - rpt.RightMargin) * width
	} else {
		width = ut.ToFloat(v.Width, width)
	}
	if width == 0 {
		width = pageWidth - rpt.LeftMargin - rpt.RightMargin
	}
	options := IM{"borderColor": v.BorderColor}
	rpt.setPageStyle(options)
	if !virtual {
		rpt.pdf.Line(rpt.pdf.GetX(), rpt.pdf.GetY(), rpt.pdf.GetX()+width, rpt.pdf.GetY())
		if v.Gap > 0 {
			rpt.pdf.Line(rpt.pdf.GetX(), rpt.pdf.GetY()+v.Gap, rpt.pdf.GetX()+width, rpt.pdf.GetY()+v.Gap)
		}
	}
}

func (rpt *Report) createElement(section string, element interface{}) {
	switch v := element.(type) {
	case *Row:
		if v.Visible != "" {
			if _, found := rpt.data[v.Visible]; found {
				if srows, valid := rpt.data[v.Visible].([]SM); !valid || len(srows) == 0 {
					return
				}
			}
		}
		rpt.createRow(section, v, false)
	case *VGap:
		if rpt.checkPageBreak(v.Height) {
			rpt.addPage()
		}
		rpt.pdf.Ln(v.Height)
	case *HLine:
		rpt.createLine(v, false)
	case *HTML:
		rpt.createHTML(v)
	case *Datagrid:
		rpt.createDatagrid(v, false)
	}
}

func (rpt *Report) setPageStyle(options IM) {
	//font-family, font-size, font-style, color,background-color,border-color
	fontStyle := ut.ToString(options["fontStyle"], "")
	fontSize := ut.ToFloat(options["fontSize"], rpt.FontSize)
	rpt.pdf.SetFont(rpt.FontFamily, fontStyle, fontSize)

	if textColor, textKey := options["textColor"]; textKey {
		rpt.pdf.SetTextColor(int(textColor.(color.RGBA).R), int(textColor.(color.RGBA).G), int(textColor.(color.RGBA).B))
	}
	if borderColor, borderKey := options["borderColor"]; borderKey {
		rpt.pdf.SetDrawColor(int(borderColor.(color.RGBA).R), int(borderColor.(color.RGBA).G), int(borderColor.(color.RGBA).B))
	}
	if backgroundColor, backgroundKey := options["backgroundColor"]; backgroundKey {
		rpt.pdf.SetFillColor(int(backgroundColor.(color.RGBA).R), int(backgroundColor.(color.RGBA).G), int(backgroundColor.(color.RGBA).B))
	}

}

func (rpt *Report) parseValue(vname string, value interface{}) interface{} {

	parseStringMap := func(value interface{}, defValue string) interface{} {
		svalue := ut.ToString(value, defValue)
		valid := SM{
			"R": "R", "C": "C", "L": "L", "J": "J", "left": "L", "center": "C", "right": "R", "justify": "L",
			"B": "B", "I": "I", "BI": "BI", "IB": "IB", "bold": "B", "italic": "I", "bolditalic": "BI", "normal": "",
			"p": "p", "l": "l", "portrait": "p", "landscape": "l",
			"a3": "a3", "a4": "a4", "a5": "a5", "letter": "letter", "legal": "legal",
		}
		if _, found := valid[svalue]; found {
			return valid[svalue]
		}
		return defValue
	}
	checkValue := map[string]func(value interface{}) interface{}{
		"Format": func(value interface{}) interface{} {
			return parseStringMap(value, _format)
		},
		"Orientation": func(value interface{}) interface{} {
			return parseStringMap(value, _orientation)
		},
		"FontSize": func(value interface{}) interface{} {
			return ut.ToFloat(value, rpt.FontSize)
		},
		"Height": func(value interface{}) interface{} {
			return ut.ToFloat(value, 0) * _mmPt
		},
		"Gap": func(value interface{}) interface{} {
			return ut.ToFloat(value, 0) * _mmPt
		},
		"HGap": func(value interface{}) interface{} {
			return ut.ToFloat(value, 0) * _mmPt
		},
		"Width": func(value interface{}) interface{} {
			switch v := value.(type) {
			case string:
				if strings.HasSuffix(v, "%") {
					svalue := strings.Replace(v, "%", "", -1)
					ivalue := ut.ToInteger(svalue, 0)
					if ivalue > 100 {
						ivalue = 100
					}
					return ut.ToString(ivalue, "0") + "%"
				}
				return ut.ToString(ut.ToFloat(value, 0)*_mmPt, "0")

			default:
				return ut.ToString(ut.ToFloat(value, 0)*_mmPt, "0")
			}
		},
		"Merge": func(value interface{}) interface{} {
			return ut.ToBoolean(value, false)
		},
		"VisibleValue": func(value interface{}) interface{} {
			return ut.ToBoolean(value, false)
		},
		"Multiline": func(value interface{}) interface{} {
			return ut.ToBoolean(value, false)
		},
		"TextColor": func(value interface{}) interface{} {
			return ut.ToRGBA(value, rpt.TextColor)
		},
		"Border": func(value interface{}) interface{} {
			borderStr := ut.ToString(value, "")
			if borderStr == "1" || strings.Contains(borderStr, "T") || strings.Contains(borderStr, "L") ||
				strings.Contains(borderStr, "R") || strings.Contains(borderStr, "B") {
				return value
			}
			return ""
		},
		"BorderColor": func(value interface{}) interface{} {
			return ut.ToRGBA(value, rpt.BorderColor)
		},
		"BackgroundColor": func(value interface{}) interface{} {
			return ut.ToRGBA(value, rpt.BackgroundColor)
		},
		"FooterBackground": func(value interface{}) interface{} {
			return ut.ToRGBA(value, rpt.BackgroundColor)
		},
		"HeaderBackground": func(value interface{}) interface{} {
			return ut.ToRGBA(value, rpt.BackgroundColor)
		},
		"FontStyle": func(value interface{}) interface{} {
			return parseStringMap(value, _fontStyle)
		},
		"Align": func(value interface{}) interface{} {
			return parseStringMap(value, _align)
		},
		"HeaderAlign": func(value interface{}) interface{} {
			return parseStringMap(value, _align)
		},
		"FooterAlign": func(value interface{}) interface{} {
			return parseStringMap(value, _align)
		},
	}

	if _, found := checkValue[vname]; found {
		return checkValue[vname](value)
	}
	return value
}

func (rpt *Report) setFont() bool {
	fontFile := map[string]func(family string) string{
		"REGULAR": func(family string) string {
			return family + "-Regular.ttf"
		},
		"BOLD": func(family string) string {
			return family + "-Bold.ttf"
		},
		"ITALIC": func(family string) string {
			return family + "-Italic.ttf"
		},
		"BOLDITALIC": func(family string) string {
			return family + "-BoldItalic.ttf"
		},
	}

	checkCustom := func() bool {
		if _, err := os.Stat(path.Join(rpt.fontDir, fontFile["REGULAR"](rpt.FontFamily))); err != nil {
			return false
		}
		if _, err := os.Stat(path.Join(rpt.fontDir, fontFile["BOLD"](rpt.FontFamily))); err != nil {
			return false
		}
		if _, err := os.Stat(path.Join(rpt.fontDir, fontFile["ITALIC"](rpt.FontFamily))); err != nil {
			return false
		}
		if _, err := os.Stat(path.Join(rpt.fontDir, fontFile["BOLDITALIC"](rpt.FontFamily))); err != nil {
			return false
		}
		return true
	}

	custom := false
	if rpt.FontFamily != _fontFamily && rpt.fontDir != "" {
		custom = checkCustom()
	}

	if custom {
		rpt.pdf.AddFont(rpt.FontFamily, "", path.Join(rpt.fontDir, fontFile["REGULAR"](rpt.FontFamily)), nil)
		rpt.pdf.AddFont(rpt.FontFamily, "B", path.Join(rpt.fontDir, fontFile["BOLD"](rpt.FontFamily)), nil)
		rpt.pdf.AddFont(rpt.FontFamily, "I", path.Join(rpt.fontDir, fontFile["ITALIC"](rpt.FontFamily)), nil)
		rpt.pdf.AddFont(rpt.FontFamily, "BI", path.Join(rpt.fontDir, fontFile["BOLDITALIC"](rpt.FontFamily)), nil)
	} else {
		rpt.FontFamily = _fontFamily
		font, _ := ut.Report.Open(path.Join("static", "fonts", fontFile["REGULAR"](_fontFamily)))
		rpt.pdf.AddFont(rpt.FontFamily, "", "", font)
		font, _ = ut.Report.Open(path.Join("static", "fonts", fontFile["BOLD"](_fontFamily)))
		rpt.pdf.AddFont(rpt.FontFamily, "B", "", font)
		font, _ = ut.Report.Open(path.Join("static", "fonts", fontFile["ITALIC"](_fontFamily)))
		rpt.pdf.AddFont(rpt.FontFamily, "I", "", font)
		font, _ = ut.Report.Open(path.Join("static", "fonts", fontFile["BOLDITALIC"](_fontFamily)))
		rpt.pdf.AddFont(rpt.FontFamily, "BI", "", font)
	}
	return true
}

func (rpt *Report) addPage() {
	rpt.pdf.AddPage()
	rpt.pdf.SetXY(rpt.LeftMargin, rpt.TopMargin)
	rpt.createHeaderAndFooter()
}

func (rpt *Report) onPage() {
	rpt.addPage()
}

/*
New returns a pointer to a new Report instance. Options:
 • orientation - Optional. Default value:"P" Values: "P","portrait","L","landscape".
 • format - Optional. Defaut value: "A4" Values: "A3","A4","A5","letter","legal".
 • fontFamily - Optional Default: Cabin
 • fontDir - Optional Default: ""

Example:
 rpt := report.New("P", "A4")
*/
func New(options ...string) (rpt *Report) {
	rpt = new(Report)
	if len(options) > 0 {
		rpt.orientation = rpt.parseValue("Orientation", options[0]).(string)
	} else {
		rpt.orientation = rpt.parseValue("Orientation", _orientation).(string)
	}
	if len(options) > 1 {
		rpt.format = rpt.parseValue("Format", options[1]).(string)
	} else {
		rpt.format = rpt.parseValue("Format", _format).(string)
	}
	rpt.FontFamily = _fontFamily
	if len(options) > 2 {
		rpt.FontFamily = options[2]
	}
	if len(options) > 3 {
		rpt.fontDir = options[3]
	}

	rpt.Title = _title
	rpt.LeftMargin = _margin
	rpt.TopMargin = _margin
	rpt.RightMargin = _margin
	rpt.BottomMargin = _margin
	rpt.FontStyle = _fontStyle
	rpt.FontSize = _fontSize
	rpt.TextColor = color.RGBA{_textColor, _textColor, _textColor, 0}
	rpt.BorderColor = color.RGBA{_borderColor, _borderColor, _borderColor, 0}
	rpt.BackgroundColor = color.RGBA{_backgroundColor, _backgroundColor, _backgroundColor, 0}
	rpt.header = make([]PageItem, 0)
	rpt.details = make([]PageItem, 0)
	rpt.footer = make([]PageItem, 0)
	rpt.data = make(IM)

	rpt.pdf = generators[_generator]
	rpt.pdf.Init(rpt)
	rpt.setFont()

	return
}

// CreateReport - the report template processing, databind replacement.
func (rpt *Report) CreateReport() bool {
	rpt.pdf.SetProperties(rpt)
	rpt.setPageStyle(make(IM))
	rpt.footerHeight = rpt.getFooterHeight()
	rpt.addPage()
	for index := 0; index < len(rpt.details); index++ {
		rpt.createElement("details", rpt.details[index].Item)
	}
	return true
}

func (rpt *Report) getJSONElements(edata interface{}) (el PageItem, err error) {
	for eName, eValue := range edata.(IM) {
		if el, err = rpt.getPageItem(eName); err != nil {
			return el, err
		}
		for ekey, eValueData := range eValue.(IM) {
			if ekey == "columns" {
				for colIndex := 0; colIndex < len(eValueData.([]interface{})); colIndex++ {
					coldata := eValueData.([]interface{})[colIndex]
					for cName, cValue := range coldata.(IM) {
						switch cName {
						case "cell", "image", "barcode", "separator", "column":
							el2, _ := rpt.getPageItem(cName)
							for ckey, cValueData := range cValue.(IM) {
								if err := el2.setPageItem(ckey, rpt.parseValue(propMap[strings.ToLower(ckey)], cValueData)); err != nil {
									return el, err
								}
							}
							if eName == "row" {
								el.Item.(*Row).Columns = append(el.Item.(*Row).Columns, el2)
							} else {
								el.Item.(*Datagrid).Columns = append(el.Item.(*Datagrid).Columns, el2)
							}
						default:
							return el, errors.New(invalidErr("Columns", cName))
						}
					}
				}
			} else {
				if err := el.setPageItem(ekey, rpt.parseValue(propMap[strings.ToLower(ekey)], eValueData)); err != nil {
					return el, err
				}
			}
		}
	}
	return el, nil
}

// LoadJSONDefinition load to the report an JSON definition.
func (rpt *Report) LoadJSONDefinition(jsonString string) error {

	if jsonString == "" {
		return errors.New("missing JSON")
	}
	var jsonData IM
	if err := ut.ConvertFromByte([]byte(jsonString), &jsonData); err != nil {
		return err
	}
	if report, found := jsonData["report"]; found {
		for valueKey, valueData := range report.(IM) {
			if err := rpt.SetReportValue(valueKey, rpt.parseValue(propMap[strings.ToLower(valueKey)], valueData)); err != nil {
				return err
			}
		}
	}
	if header, found := jsonData["header"]; found {
		for index := 0; index < len(header.([]interface{})); index++ {
			el, err := rpt.getJSONElements(header.([]interface{})[index])
			if err != nil {
				return err
			}
			rpt.header = append(rpt.header, el)
		}
	}
	if details, found := jsonData["details"]; found {
		for index := 0; index < len(details.([]interface{})); index++ {
			el, err := rpt.getJSONElements(details.([]interface{})[index])
			if err != nil {
				return err
			}
			rpt.details = append(rpt.details, el)
		}
	}
	if footer, found := jsonData["footer"]; found {
		for index := 0; index < len(footer.([]interface{})); index++ {
			el, err := rpt.getJSONElements(footer.([]interface{})[index])
			if err != nil {
				return err
			}
			rpt.footer = append(rpt.footer, el)
		}
	}
	if data, found := jsonData["data"]; found {
		for dKey, dValue := range data.(IM) {
			switch dValue.(type) {
			case []interface{}:
				rpt.data[dKey] = make([]SM, 0)
				for index := 0; index < len(dValue.([]interface{})); index++ {
					jRow := dValue.([]interface{})[index]
					dRow := SM{}
					for key, value := range jRow.(IM) {
						dRow[key] = ut.ToString(value, "")
					}
					rpt.data[dKey] = append(rpt.data[dKey].([]SM), dRow)
				}
			case IM:
				rpt.data[dKey] = SM{}
				for key, value := range dValue.(IM) {
					rpt.data[dKey].(SM)[key] = ut.ToString(value, "")
				}
			case string, SM, []SM:
				rpt.data[dKey] = dValue
			default:
				return errors.New("valid data types: string, map[string][string], []map[string][string] ")
			}
		}
	}
	return nil
}

/*
AppendElement - Append an element in the template.
 • parent - Optional. The parent elemnt. Values: "header","details","footer" or result value (row, datagrid) Default value: "details"
 • ename - Optional. An Element type: "row", "datagrid", "vgap", "hline", "html", "column", "cell", "image", "separator", "barcode". Default value: "row"
 • values - Optional. Element attributes
Example:
 row_data := rpt.AppendElement("header", "row", map[string]interface{}{"height": 10})
 rpt.AppendElement(row_data, "image", map[string]interface{}{"src": "test/logo.jpg"})

*/
func (rpt *Report) AppendElement(options ...interface{}) (*[]PageItem, error) {

	el, _ := rpt.getPageItem("row")
	parent := &rpt.details

	if len(options) > 0 {
		switch options[0].(type) {
		case string:
			switch options[0] {
			case "header":
				parent = &rpt.header
				if len(options) > 1 {
					ename := ut.ToString(options[1], "")
					if ut.Contains([]string{"row", "vgap", "hline"}, ename) {
						el, _ = rpt.getPageItem(ename)
					} else {
						return nil, errors.New(invalidErr("Header", ename))
					}
				}
			case "details":
				parent = &rpt.details
				if len(options) > 1 {
					ename := ut.ToString(options[1], "")
					if ut.Contains([]string{"row", "vgap", "hline", "html", "datagrid"}, ename) {
						el, _ = rpt.getPageItem(ename)
					} else {
						return nil, errors.New(invalidErr("Details", ename))
					}
				}
			case "footer":
				parent = &rpt.footer
				if len(options) > 1 {
					ename := ut.ToString(options[1], "")
					if ut.Contains([]string{"row", "vgap", "hline"}, ename) {
						el, _ = rpt.getPageItem(ename)
					} else {
						return nil, errors.New(invalidErr("Footer", ename))
					}
				}
			}
		case *[]PageItem:
			parent = options[0].(*[]PageItem)
			if len(options) > 1 {
				ename := ut.ToString(options[1], "")
				if ut.Contains([]string{"cell", "image", "barcode", "separator", "column"}, ename) {
					el, _ = rpt.getPageItem(ename)
				} else {
					return nil, errors.New(invalidErr("columns", ename))
				}
			}
		default:
			return nil, errors.New("valid parent values: 'header','details','footer' (string) or Columns of Row and Datagrid (*[]PageItem)")
		}
	}

	if len(options) > 2 {
		switch options[2].(type) {
		case IM:
			for key, value := range options[2].(IM) {
				if err := el.setPageItem(key, rpt.parseValue(propMap[strings.ToLower(key)], value)); err != nil {
					return nil, err
				}
			}
		default:
			return nil, errors.New("valid values type: map[string]interface{}")
		}

	}

	*parent = append(*parent, el)
	if el.ItemType == "row" {
		return &el.Item.(*Row).Columns, nil
	} else if el.ItemType == "datagrid" {
		return &el.Item.(*Datagrid).Columns, nil
	}
	return parent, nil
}

/*
SetData - Set the template data. Parameters:
 • key - string
 • value - interface{} Valid interface type: string or dictonary (map[string]string) or record list ([]map[string]string)
Example:
 rpt.SetData("items_footer", map[string]string{"items_total": "3 703 680"})
*/
func (rpt *Report) SetData(key string, value interface{}) (bool, error) {

	if _, found := rpt.data[key]; found {
		switch rpt.data[key].(type) {
		case SM:
			switch value.(type) {
			case SM:
				for valueKey, valueData := range value.(SM) {
					rpt.data[key].(SM)[valueKey] = valueData
				}
				return true, nil
			}
		}
	}
	switch value.(type) {
	case string, SM, []SM:
		rpt.data[key] = value
	default:
		return false, errors.New("valid value types: string, map[string][string], []map[string]string")
	}
	return true, nil
}

// Save2DataURLString creates a base64 data URI scheme.
func (rpt *Report) Save2DataURLString(filename string) (string, error) {
	pdf, err := rpt.Save2Pdf()
	if err != nil {
		return "", err
	}
	pdfStr := base64.URLEncoding.EncodeToString([]byte(pdf))
	if filename != "" {
		filename = "filename=" + filename + ";"
	}
	return "data:application/pdf;" + filename + "base64," + pdfStr, nil
}

// Save2Pdf creates a PDF output.
func (rpt *Report) Save2Pdf() ([]byte, error) {
	return rpt.pdf.Save2Pdf()
}

// Save2PdfFile creates or truncates the file specified by fileStr and
// writes the PDF document to it.
func (rpt *Report) Save2PdfFile(filename string) error {
	return rpt.pdf.Save2PdfFile(filename)
}

// Save2Xml creates an XML output. Only the values of cells and datagrid
// rows from header and details. The node name of the cell name (except when
// name="label"), or datagrid name/column fieldname.
func (rpt *Report) Save2Xml() string {
	return fmt.Sprintf("<data>%s%s\n</data>", rpt.xmlHeader, rpt.xmlDetails)
}
