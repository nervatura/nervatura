package report

import (
	"bytes"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/signintech/gopdf"
)

type genGoPDF struct {
	pdf          gopdf.GoPdf
	format       string
	orientation  string
	fontFamily   string
	fontStyle    string
	fontSize     float64
	leftMargin   float64
	topMargin    float64
	rightMargin  float64
	bottomMargin float64
	splitText    func(txt string, w float64) []string
	onPage       func()
	pageSize     gopdf.Rect
}

var sizeMap = map[string]gopdf.Rect{
	"a3":     *gopdf.PageSizeA3,
	"a4":     *gopdf.PageSizeA4,
	"a5":     *gopdf.PageSizeA5,
	"letter": *gopdf.PageSizeLetter,
	"legal":  *gopdf.PageSizeLegal,
}

func init() {
	registerGenerators("gopdf", &genGoPDF{})
}

func (gen *genGoPDF) Init(rpt *Report) {
	gen.fontFamily = rpt.FontFamily
	gen.fontSize = rpt.FontSize
	gen.pdf = gopdf.GoPdf{}
	gen.format = strings.ToLower(rpt.format)
	gen.orientation = strings.ToLower(rpt.orientation)
	gen.leftMargin = rpt.LeftMargin
	gen.topMargin = rpt.TopMargin
	gen.rightMargin = rpt.RightMargin
	gen.bottomMargin = rpt.BottomMargin
	gen.splitText = func(txt string, w float64) []string {
		return rpt.wrapTextLines(txt, w)
	}
	gen.onPage = func() {
		rpt.onPage()
	}
	size := sizeMap[strings.ToLower(gen.format)]
	if gen.orientation == "p" {
		gen.pageSize = size
	} else {
		gen.pageSize = gopdf.Rect{W: size.H, H: size.W}
	}
	gen.pdf.Start(gopdf.Config{
		Unit:     1,
		PageSize: gen.pageSize,
	})
}

// GetPageSize returns the current page's width and height.
func (gen *genGoPDF) GetPageSize() (width, height float64) {
	return gen.pageSize.W, gen.pageSize.H
}

// PageNo returns the current page number.
func (gen *genGoPDF) PageNo() int {
	return gen.pdf.GetNumberOfPages()
}

// AddPage adds a new page to the document
func (gen *genGoPDF) AddPage() {
	gen.pdf.AddPage()
}

// AddImage draws an image
func (gen *genGoPDF) AddImage(image *Image, x, y float64, options IM) {
	if image.Data != nil {
		img, err := gopdf.ImageHolderByBytes(image.Data)
		if err != nil {
			return
		}
		_ = gen.pdf.ImageByHolder(img, x, y, &gopdf.Rect{H: image.Height, W: image.Width})
		return
	}
	if image.Src != "" {
		src := image.Src
		if options["ImagePath"].(string) != "" {
			src = path.Join(options["ImagePath"].(string), image.Src)
		}
		_ = gen.pdf.Image(src, x, y, &gopdf.Rect{H: image.Height, W: image.Width})
	}

}

// AddFont imports a font and makes it available
func (gen *genGoPDF) AddFont(familyStr, styleStr, fileStr string, rd io.Reader) {
	style := func(check string, value int) int {
		if strings.Contains(styleStr, check) {
			return value
		}
		return gopdf.Regular
	}
	if rd != nil {
		_ = gen.pdf.AddTTFFontByReaderWithOption(familyStr, rd,
			gopdf.TtfOption{Style: style("B", gopdf.Bold) | style("I", gopdf.Italic)})
		return
	}
	_ = gen.pdf.AddTTFFontWithOption(familyStr, fileStr,
		gopdf.TtfOption{Style: style("B", gopdf.Bold) | style("I", gopdf.Italic)})
}

// GetFontSize returns the size of the current font in points.
func (gen *genGoPDF) GetFontSize() (ptSize float64) {
	return gen.fontSize
}

// SetFont sets the font used to print character strings
func (gen *genGoPDF) SetFont(familyStr, styleStr string, size float64) {
	if familyStr == "" {
		familyStr = gen.fontFamily
	}
	if size == 0 {
		size = gen.fontSize
	}
	err := gen.pdf.SetFont(familyStr, styleStr, int(size))
	if err == nil {
		gen.fontFamily = familyStr
		gen.fontStyle = styleStr
		gen.fontSize = size
	}
}

// SetFontSize defines the size of the current font.
func (gen *genGoPDF) SetFontSize(size float64) {
	if err := gen.pdf.SetFont(gen.fontFamily, gen.fontStyle, int(size)); err != nil {
		return
	}
}

// GetTextWidth returns the length of a string in user units.
func (gen *genGoPDF) GetTextWidth(s string) float64 {
	w, err := gen.pdf.MeasureTextWidth(s)
	if err != nil {
		return 0
	}
	return w
}

// SetDrawColor defines the color used for all drawing operations
func (gen *genGoPDF) SetDrawColor(r, g, b int) {
	gen.pdf.SetStrokeColor(uint8(r), uint8(g), uint8(b))
}

// SetFillColor defines the color used for all filling operations
func (gen *genGoPDF) SetFillColor(r, g, b int) {
	gen.pdf.SetFillColor(uint8(r), uint8(g), uint8(b))
}

// SetTextColor defines the color used for text.
func (gen *genGoPDF) SetTextColor(r, g, b int) {
	gen.pdf.SetTextColor(uint8(r), uint8(g), uint8(b))
}

// SetProperties - general report props. (title, author etc.)
func (gen *genGoPDF) SetProperties(rpt *Report) {
	gen.pdf.SetInfo(gopdf.PdfInfo{
		Title: rpt.Title, Author: rpt.Author, Subject: rpt.Subject, Creator: rpt.Creator,
	})
	gen.pdf.SetMargins(rpt.LeftMargin, rpt.TopMargin, rpt.RightMargin, rpt.BottomMargin)
}

// Text - Write prints text from the current position.
func (gen *genGoPDF) Text(txtStr string, pageBreak float64) {
	checkBreak := func(height float64) {
		cy := gen.pdf.GetY()
		if cy+height > pageBreak {
			gen.onPage()
			gen.Ln(height)
		}
	}
	tw := gen.GetTextWidth(txtStr)
	cw := gen.currentWidth()
	lineHt := gen.GetFontSize()
	if tw > cw {
		txt := gen.splitText(txtStr, cw)[0]
		checkBreak(lineHt)
		if err := gen.pdf.Text(txt); err != nil {
			return
		}
		txtStr = strings.TrimLeft(txtStr[len(txt)-1:], " ")
		pw, _ := gen.GetPageSize()
		lines := gen.splitText(txtStr, pw-gen.leftMargin-gen.rightMargin)
		for i := 0; i < len(lines); i++ {
			gen.Ln(lineHt)
			checkBreak(lineHt)
			if err := gen.pdf.Text(lines[i]); err != nil {
				return
			}
		}
	} else {
		checkBreak(lineHt)
		if err := gen.pdf.Text(txtStr); err != nil {
			return
		}
	}
}

// Rect outputs a rectangle of width w and height h with the upper left corner positioned at point (x, y)
func (gen *genGoPDF) Rect(x, y, w, h float64, styleStr string) {
	gen.pdf.RectFromUpperLeftWithStyle(x, y, w, h, styleStr)
}

// Line draws a line between points (x1, y1) and (x2, y2) using the current draw color, line width and cap style.
func (gen *genGoPDF) Line(x1, y1, x2, y2 float64) {
	gen.pdf.Line(x1, y1, x2, y2)
}

// GetX returns the abscissa of the current position.
func (gen *genGoPDF) GetX() float64 {
	return gen.pdf.GetX()
}

// GetY returns the ordinate of the current position.
func (gen *genGoPDF) GetY() float64 {
	return gen.pdf.GetY()
}

// SetX defines the abscissa of the current position.
func (gen *genGoPDF) SetX(x float64) {
	gen.pdf.SetX(x)
}

// SetY : set current position y
func (gen *genGoPDF) SetY(y float64) {
	gen.pdf.SetY(y)
}

// SetXY defines the abscissa and ordinate of the current position.
func (gen *genGoPDF) SetXY(x, y float64) {
	gen.pdf.SetX(x)
	gen.pdf.SetY(y)
}

// Ln performs a line break.
func (gen *genGoPDF) Ln(h float64) {
	gen.pdf.Br(h)
}

func (gen *genGoPDF) setBorder(borderStr string, w, h, x, y float64) {
	if borderStr == "1" || strings.Contains(borderStr, "T") {
		gen.pdf.Line(x, y, x+w, y)
	}
	if borderStr == "1" || strings.Contains(borderStr, "B") {
		gen.pdf.Line(x, y+h, x+w, y+h)
	}
	if borderStr == "1" || strings.Contains(borderStr, "L") {
		gen.pdf.Line(x, y, x, y+h)
	}
	if borderStr == "1" || strings.Contains(borderStr, "R") {
		gen.pdf.Line(x+w, y, x+w, y+h)
	}
}

func (gen *genGoPDF) currentWidth() float64 {
	pw, _ := gen.GetPageSize()
	return pw - gen.rightMargin - gen.pdf.GetX()
}

// Cell prints a rectangular cell with optional borders, background color and character string.
func (gen *genGoPDF) Cell(options IM) {
	if options["w"].(float64) == 0 {
		options["w"] = gen.currentWidth()
	}
	cx := gen.pdf.GetX()
	cy := gen.pdf.GetY()
	if options["fill"].(bool) {
		gen.pdf.RectFromUpperLeftWithStyle(cx, cy, options["w"].(float64), options["h"].(float64), "F")
	}
	if options["borderStr"] != "" {
		gen.setBorder(options["borderStr"].(string), options["w"].(float64), options["h"].(float64), cx, cy)
	}
	if options["alignStr"] == "L" {
		gen.pdf.SetX(cx + options["padding"].(float64)/2)
		if err := gen.pdf.CellWithOption(&gopdf.Rect{W: options["w"].(float64), H: options["h"].(float64)}, options["txtStr"].(string),
			gopdf.CellOption{Align: gopdf.Left | gopdf.Middle, Float: gopdf.Right}); err != nil {
			return
		}
		gen.pdf.SetX(cx + options["w"].(float64))
	}
	if options["alignStr"] == "C" {
		if err := gen.pdf.CellWithOption(&gopdf.Rect{W: options["w"].(float64), H: options["h"].(float64)}, options["txtStr"].(string),
			gopdf.CellOption{Align: gopdf.Center | gopdf.Middle, Float: gopdf.Right}); err != nil {
			return
		}
	}
	if options["alignStr"] == "R" {
		tw := gen.GetTextWidth(options["txtStr"].(string))
		gen.pdf.SetX(cx + (options["w"].(float64) - tw - options["padding"].(float64)/2))
		if err := gen.pdf.CellWithOption(&gopdf.Rect{W: options["w"].(float64), H: options["h"].(float64)}, options["txtStr"].(string),
			gopdf.CellOption{Align: gopdf.Left | gopdf.Middle, Float: gopdf.Right}); err != nil {
			return
		}
		gen.pdf.SetX(cx + options["w"].(float64))
	}
	if options["ln"].(bool) {
		gen.pdf.Br(options["h"].(float64))
	}
}

// MultiCell supports printing text with line breaks.
func (gen *genGoPDF) MultiCell(options IM) {
	cx := gen.pdf.GetX()
	cy := gen.pdf.GetY()
	if options["fill"].(bool) {
		gen.pdf.RectFromUpperLeftWithStyle(cx, gen.pdf.GetY(), options["w"].(float64), options["h"].(float64), "F")
	}
	if options["borderStr"] != "" {
		gen.setBorder(options["borderStr"].(string), options["w"].(float64), options["h"].(float64), cx, cy)
	}
	lines := gen.splitText(options["txtStr"].(string), options["w"].(float64)-options["padding"].(float64))
	for i := 0; i < len(lines); i++ {
		txt := lines[i]
		gen.pdf.SetX(cx)
		gen.Cell(IM{
			"w": options["w"].(float64), "h": options["lineH"].(float64), "padding": options["padding"].(float64),
			"txtStr": txt, "borderStr": "", "alignStr": options["alignStr"].(string), "fill": false, "ln": true,
		})
	}
}

// Save2Pdf creates a PDF output.
func (gen *genGoPDF) Save2Pdf() ([]byte, error) {
	ba := bytes.NewBuffer([]byte{})
	if err := gen.pdf.Write(ba); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(ba)
}

// Save2PdfFile writes the PDF document to file
func (gen *genGoPDF) Save2PdfFile(filename string) error {
	return gen.pdf.WritePdf(filename)
}
