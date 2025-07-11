// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package templates

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/russross/blackfriday"
)

const linebreak = "\n"

// ASCIIRenderer implements blackfriday.Renderer.
var _ blackfriday.Renderer = &ASCIIRenderer{}

// ASCIIRenderer is a blackfriday.Renderer intended for rendering markdown
// documents as plain text, well suited for human reading on terminals.
type ASCIIRenderer struct {
	Indentation string

	listItemCount uint
	listLevel     uint
}

// NormalText gets a text chunk *after* the markdown syntax was already
// processed and does a final cleanup on things we don't expect here, like
// removing linebreaks on things that are not a paragraph break (auto unwrap).
func (r *ASCIIRenderer) NormalText(out *bytes.Buffer, text []byte) {
	raw := string(text)
	lines := strings.Split(raw, linebreak)
	for _, line := range lines {
		trimmed := strings.Trim(line, " \n\t")
		if len(trimmed) > 0 && trimmed[0] != '_' {
			out.WriteString(" ")
		}
		out.WriteString(trimmed)
	}
}

// List renders the start and end of a list.
func (r *ASCIIRenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	r.listLevel++
	out.WriteString(linebreak)
	text()
	r.listLevel--
}

// ListItem renders list items and supports both ordered and unordered lists.
func (r *ASCIIRenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	if flags&blackfriday.LIST_ITEM_BEGINNING_OF_LIST != 0 {
		r.listItemCount = 1
	} else {
		r.listItemCount++
	}
	indent := strings.Repeat(r.Indentation, int(r.listLevel))
	var bullet string
	if flags&blackfriday.LIST_TYPE_ORDERED != 0 {
		bullet += fmt.Sprintf("%d.", r.listItemCount)
	} else {
		bullet += "*"
	}
	out.WriteString(indent + bullet + " ")
	r.fw(out, text)
	out.WriteString(linebreak)
}

// Paragraph renders the start and end of a paragraph.
func (r *ASCIIRenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	out.WriteString(linebreak)
	text()
	out.WriteString(linebreak)
}

// BlockCode renders a chunk of text that represents source code.
func (r *ASCIIRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	out.WriteString(linebreak)
	lines := []string{}
	for _, line := range strings.Split(string(text), linebreak) {
		indented := r.Indentation + line
		lines = append(lines, indented)
	}
	out.WriteString(strings.Join(lines, linebreak))
}

func (r *ASCIIRenderer) GetFlags() int { return 0 }
func (r *ASCIIRenderer) HRule(out *bytes.Buffer) {
	out.WriteString(linebreak + "----------" + linebreak)
}
func (r *ASCIIRenderer) LineBreak(out *bytes.Buffer)                                      { out.WriteString(linebreak) }
func (r *ASCIIRenderer) TitleBlock(out *bytes.Buffer, text []byte)                        { r.fw(out, text) }
func (r *ASCIIRenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) { text() }
func (r *ASCIIRenderer) BlockHtml(out *bytes.Buffer, text []byte)                         { r.fw(out, text) }
func (r *ASCIIRenderer) BlockQuote(out *bytes.Buffer, text []byte)                        { r.fw(out, text) }
func (r *ASCIIRenderer) TableRow(out *bytes.Buffer, text []byte)                          { r.fw(out, text) }
func (r *ASCIIRenderer) TableHeaderCell(out *bytes.Buffer, text []byte, align int)        { r.fw(out, text) }
func (r *ASCIIRenderer) TableCell(out *bytes.Buffer, text []byte, align int)              { r.fw(out, text) }
func (r *ASCIIRenderer) Footnotes(out *bytes.Buffer, text func() bool)                    { text() }
func (r *ASCIIRenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	r.fw(out, text)
}
func (r *ASCIIRenderer) AutoLink(out *bytes.Buffer, link []byte, kind int)         { r.fw(out, link) }
func (r *ASCIIRenderer) CodeSpan(out *bytes.Buffer, text []byte)                   { r.fw(out, text) }
func (r *ASCIIRenderer) DoubleEmphasis(out *bytes.Buffer, text []byte)             { r.fw(out, text) }
func (r *ASCIIRenderer) Emphasis(out *bytes.Buffer, text []byte)                   { r.fw(out, text) }
func (r *ASCIIRenderer) RawHtmlTag(out *bytes.Buffer, text []byte)                 { r.fw(out, text) }
func (r *ASCIIRenderer) TripleEmphasis(out *bytes.Buffer, text []byte)             { r.fw(out, text) }
func (r *ASCIIRenderer) StrikeThrough(out *bytes.Buffer, text []byte)              { r.fw(out, text) }
func (r *ASCIIRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int)         { r.fw(out, ref) }
func (r *ASCIIRenderer) Entity(out *bytes.Buffer, entity []byte)                   { r.fw(out, entity) }
func (r *ASCIIRenderer) Smartypants(out *bytes.Buffer, text []byte)                { r.fw(out, text) }
func (r *ASCIIRenderer) DocumentHeader(out *bytes.Buffer)                          {}
func (r *ASCIIRenderer) DocumentFooter(out *bytes.Buffer)                          {}
func (r *ASCIIRenderer) TocHeaderWithAnchor(text []byte, level int, anchor string) {}
func (r *ASCIIRenderer) TocHeader(text []byte, level int)                          {}
func (r *ASCIIRenderer) TocFinalize()                                              {}

func (r *ASCIIRenderer) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	r.fw(out, header, body)
}

func (r *ASCIIRenderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.WriteString(" ")
	r.fw(out, link)
}

func (r *ASCIIRenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	r.fw(out, link)
}

func (r *ASCIIRenderer) fw(out io.Writer, text ...[]byte) {
	for _, t := range text {
		out.Write(t)
	}
}
