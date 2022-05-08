package text

import (
	"fmt"
	"io"
)

type basicWriter struct {
	io.StringWriter
}

func (w basicWriter) Printf(template string, data ...interface{}) (int, error) {
	return w.WriteString(
		fmt.Sprintf(template, data...),
	)
}

func (w basicWriter) Println(text ...interface{}) (int, error) {
	return w.WriteString(
		fmt.Sprintln(text...),
	)
}
