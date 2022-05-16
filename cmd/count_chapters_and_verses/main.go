package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/mniak/biblia/pkg/sources/wlc"
)

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//go:embed template.go.tmpl
var templateString string

func main() {
	loader := wlc.NewLoader(wlc.DefaultDirectory)
	oldTestament, err := loader.Load()
	handle(err)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"NormalizeIdentifier": normalizeIdentifier,
		}).
		Parse(templateString)
	handle(err)

	var buffer strings.Builder
	err = tmpl.Execute(&buffer, map[string]any{
		"OldTestament": oldTestament,
	})
	handle(err)

	fmt.Println(buffer.String())
}
