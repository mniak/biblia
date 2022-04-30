package text

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func Sprintlnf(w io.StringWriter, template string, data ...interface{}) string {
	return fmt.Sprintln(
		fmt.Sprintf(template, data...),
	)
}

func Tprintf(tw TextWriter, template string, data ...interface{}) (int, error) {
	return tw.WriteString(
		fmt.Sprintf(template, data...),
	)
}

func Tprintln(tw TextWriter, text ...interface{}) (int, error) {
	return tw.WriteString(
		fmt.Sprintln(text...),
	)
}

func Tprintlnf(tw TextWriter, template string, data ...interface{}) (int, error) {
	return Tprintln(tw,
		fmt.Sprintf(template, data...),
	)
}

type TextWriter io.StringWriter

func NewIndentedWriter(textWriter TextWriter) *_IndentedTextWriter {
	return &_IndentedTextWriter{
		TextWriter:       textWriter,
		indentationSize:  2,
		shouldIndentNext: true,
	}
}

func NewIndentedStdout() *_IndentedTextWriter {
	return NewIndentedWriter(os.Stdout)
}

type _IndentedTextWriter struct {
	TextWriter
	indentationSize int

	indentationLevel int
	shouldIndentNext bool
}

func (tw *_IndentedTextWriter) Indent() {
	tw.indentationLevel++
}

func (tw *_IndentedTextWriter) Dedent() {
	if tw.indentationLevel > 0 {
		tw.indentationLevel--
	}
}

func (w *_IndentedTextWriter) WriteString(text string) (int, error) {
	var n1 int
	var err error

	if w.shouldIndentNext {
		n1, err = w.TextWriter.WriteString(strings.Repeat(" ", w.indentationLevel*w.indentationSize))
		w.shouldIndentNext = false
	}

	n2, err := w.TextWriter.WriteString(text)

	w.shouldIndentNext = shouldIndentAfter(text)
	return n1 + n2, err
}

func shouldIndentAfter(text string) bool {
	r, _ := utf8.DecodeLastRuneInString(text)
	// The Unicode standard defines a number of characters that
	// conforming applications should recognize as line terminators:
	//   LF:    Line Feed, U+000A
	//   VT:    Vertical Tab, U+000B
	//   FF:    Form Feed, U+000C
	//   CR:    Carriage Return, U+000D
	//   CR+LF: CR (U+000D) followed by LF (U+000A)
	//   NEL:   Next Line, U+0085
	//   LS:    Line Separator, U+2028
	//   PS:    Paragraph Separator, U+2029
	// Source: https://en.wikipedia.org/wiki/Newline#Unicode"
	switch r {
	case '\u000a':
		return true
	case '\u000b':
		return true
	case '\u000c':
		return true
	case '\u000d':
		return true
	case '\u0085':
		return true
	case '\u2028':
		return true
	case '\u2029':
		return true
	default:
		return false
	}
}

func (tw *_IndentedTextWriter) Print(text string) (int, error) {
	return tw.WriteString(text)
}

func (tw *_IndentedTextWriter) Printf(template string, data ...interface{}) (int, error) {
	return Tprintf(tw, template, data...)
}

func (tw *_IndentedTextWriter) Println(text ...interface{}) (int, error) {
	return Tprintln(tw, text...)
}

func (tw *_IndentedTextWriter) Printlnf(template string, data ...interface{}) (int, error) {
	return Tprintlnf(tw, template, data...)
}
