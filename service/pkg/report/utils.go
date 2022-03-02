package report

import (
	"regexp"
	"strings"
	"unicode"
)

// segmentType defines a segment of literal text in which the current
// attributes do not vary, or an open tag or a close tag.
type segmentType struct {
	Cat  byte              // 'O' open tag, 'C' close tag, 'T' text
	Str  string            // Literal text unchanged, tags are lower case
	Attr map[string]string // Attribute keys are lower case
}

// basicTokenize returns a list of HTML tags and literal elements. This is
// done with regular expressions, so the result is only marginally better than
// useless.
func basicTokenize(htmlStr string) (list []segmentType) {
	// This routine is adapted from http://www.fpdf.org/
	list = make([]segmentType, 0, 16)
	htmlStr = strings.Replace(htmlStr, "\n", " ", -1)
	htmlStr = strings.Replace(htmlStr, "\r", "", -1)
	tagRe, _ := regexp.Compile(`(?U)<.*>`)
	attrRe, _ := regexp.Compile(`([^=]+)=["']?([^"']+)`)
	capList := tagRe.FindAllStringIndex(htmlStr, -1)
	if capList != nil {
		var seg segmentType
		var parts []string
		pos := 0
		for _, cap := range capList {
			if pos < cap[0] {
				seg.Cat = 'T'
				seg.Str = htmlStr[pos:cap[0]]
				seg.Attr = nil
				list = append(list, seg)
			}
			if htmlStr[cap[0]+1] == '/' {
				seg.Cat = 'C'
				seg.Str = strings.ToLower(htmlStr[cap[0]+2 : cap[1]-1])
				seg.Attr = nil
				list = append(list, seg)
			} else {
				// Extract attributes
				parts = strings.Split(htmlStr[cap[0]+1:cap[1]-1], " ")
				if len(parts) > 0 {
					for j, part := range parts {
						if j == 0 {
							seg.Cat = 'O'
							seg.Str = strings.ToLower(parts[0])
							seg.Attr = make(map[string]string)
						} else {
							attrList := attrRe.FindAllStringSubmatch(part, -1)
							for _, attr := range attrList {
								seg.Attr[strings.ToLower(attr[1])] = attr[2]
							}
						}
					}
					list = append(list, seg)
				}
			}
			pos = cap[1]
		}
		if len(htmlStr) > pos {
			seg.Cat = 'T'
			seg.Str = htmlStr[pos:]
			seg.Attr = nil
			list = append(list, seg)
		}
	} else {
		list = append(list, segmentType{Cat: 'T', Str: htmlStr, Attr: nil})
	}
	return
}

// writeHTML prints text from the current position using the currently selected
// font. The text can be encoded with a basic subset of HTML
// that includes tags for italic (I), bold (B), underscore
// (U) attributes. When the right margin is reached a line
// break occurs and text continues from the left margin. Upon method exit, the
// current position is left at the end of the text.
func (rpt *Report) writeHTML(lineHt float64, htmlStr string) {
	var boldLvl, italicLvl, underscoreLvl int
	setStyle := func(boldAdj, italicAdj, underscoreAdj int) {
		styleStr := ""
		boldLvl += boldAdj
		if boldLvl > 0 {
			styleStr += "B"
		}
		italicLvl += italicAdj
		if italicLvl > 0 {
			styleStr += "I"
		}
		underscoreLvl += underscoreAdj
		if underscoreLvl > 0 {
			styleStr += "U"
		}
		rpt.pdf.SetFont("", styleStr, 0)
	}
	list := basicTokenize(htmlStr)
	for _, el := range list {
		switch el.Cat {
		case 'T':
			rpt.pdf.Text(el.Str, rpt.pageBreak-rpt.footerHeight)
		case 'O':
			switch el.Str {
			case "b", "strong":
				setStyle(1, 0, 0)
			case "i", "em":
				setStyle(0, 1, 0)
			//case "u":
			//	setStyle(0, 0, 1)
			case "br", "p", "div":
				rpt.pdf.Ln(lineHt)
			}
		case 'C':
			switch el.Str {
			case "b", "strong":
				setStyle(-1, 0, 0)
			case "i", "em":
				setStyle(0, -1, 0)
				//case "u":
				//setStyle(0, 0, -1)
			}
		}
	}
}

// wrapTextLines splits a string into multiple lines so that the text
// fits in the specified width. The text is wrapped on word boundaries.
// Newline characters ("\r" and "\n") also cause text to be split.
// You can find out the number of lines needed to wrap some
// text by checking the length of the returned array.
func (rpt *Report) wrapTextLines(text string, width float64) (ret []string) {
	// isWhiteSpace returns true if all the chars. in 's' are white-spaces
	isWhiteSpace := func(s string) bool {
		for _, r := range s {
			if !unicode.IsSpace(r) {
				return false
			}
		}
		return len(s) > 0
	}

	// splitLines splits 's' into several lines using line breaks in 's'
	splitLines := func(s string) []string {
		split := func(lines []string, sep string) (ret []string) {
			for _, line := range lines {
				if strings.Contains(line, sep) {
					ret = append(ret, strings.Split(line, sep)...)
					continue
				}
				ret = append(ret, line)
			}
			return ret
		}
		return split(split(split([]string{s}, "\r\n"), "\r"), "\n")
	} //

	fit := func(s string, step, n int, width float64) int {
		for max := len(s); n > 0 && n <= max; {
			w := rpt.pdf.GetTextWidth(s[:n])
			switch step {
			case 1, 3: //       keep halving (or - 1) until n chars fit in width
				if w <= width {
					return n
				}
				n--
				if step == 1 {
					n /= 2
				}
			case 2: //               increase n until n chars won't fit in width
				if w > width {
					return n
				}
				n = 1 + int((float64(n) * 1.1)) //    increase n by 1 + 20% of n
			}
		}
		return 0
	}
	// split text into lines. then break lines based on text width
	for _, line := range splitLines(text) {
		for rpt.pdf.GetTextWidth(line) > width {
			n := len(line) //    reduce, increase, then reduce n to get best fit
			for i := 1; i <= 3; i++ {
				n = fit(line, i, n, width)
			}
			// move to the last word (if white-space is found)
			found, max := false, n
			for n > 0 {
				if isWhiteSpace(line[n-1 : n]) {
					found = true
					break
				}
				n--
			}
			if !found {
				n = max
			}
			if n <= 0 {
				break
			}
			ret = append(ret, line[:n])
			line = line[n:]
		}
		ret = append(ret, line)
	}
	return ret
}
