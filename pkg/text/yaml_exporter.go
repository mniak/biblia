package text

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mniak/biblia/pkg/bible"
)

//go:embed yaml_exporter_verse.tmpl
var embedfs embed.FS

type YamlExporter struct {
	Directory string
}

func NewYamlExporter(directory string) YamlExporter {
	return YamlExporter{
		Directory: directory,
	}
}

func (e YamlExporter) Export(t bible.Testament) error {
	for _, book := range t.Books {
		normalizedBookName := strings.ReplaceAll(book.Name, " ", "_")
		bookdir := filepath.Join(e.Directory, normalizedBookName)
		for _, chapter := range book.Chapters {
			chapterdir := filepath.Join(bookdir, fmt.Sprintf("%s_%d", normalizedBookName, chapter.Number))
			err := os.MkdirAll(chapterdir, os.ModePerm)
			if err != nil {
				return err
			}

			for _, verse := range chapter.Verses {

				versepath := filepath.Join(chapterdir, fmt.Sprintf("%s_%d_%d.yaml", normalizedBookName, chapter.Number, verse.Number))
				w, err := os.Create(versepath)
				if err != nil {
					return err
				}
				defer w.Close()

				verseTemplate, err := template.ParseFS(embedfs, "yaml_exporter_verse.tmpl")
				if err != nil {
					return err
				}
				err = verseTemplate.Execute(w, map[string]any{
					"Book":    book,
					"Chapter": chapter,
					"Verse":   verse,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
