package yaml

import (
	"embed"
	"io"
	"text/template"

	"github.com/mniak/biblia/pkg/bible"
)

//go:embed verse.yaml.tmpl
var embedfs embed.FS

type fastYamlExporter struct {
	directory string
}

func NewFastExporter(directory string) fastYamlExporter {
	return fastYamlExporter{
		directory: directory,
	}
}

func (e fastYamlExporter) Export(t bible.Testament) error {
	verseTemplate, err := template.ParseFS(embedfs, "verse.yaml.tmpl")
	if err != nil {
		return err
	}

	return export(e.directory, t, func(w io.Writer, book bible.Book, chapter bible.Chapter, verse bible.Verse) error {
		return verseTemplate.Execute(w, map[string]any{
			"Book":    book,
			"Chapter": chapter,
			"Verse":   verse,
		})
	})
}
