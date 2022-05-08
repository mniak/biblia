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

//go:embed verse_template.tmpl
var embedfs embed.FS

type yamlExporter struct {
	directory string
}

func YamlExporter(directory string) yamlExporter {
	return yamlExporter{
		directory: directory,
	}
}

func (e yamlExporter) Export(t bible.Testament) error {
	chapterTemplate, err := template.ParseFS(embedfs, "verse_template.tmpl")
	if err != nil {
		return err
	}

	for _, book := range t.Books {
		normalizedBookName := strings.ReplaceAll(book.Name, " ", "_")
		bookdir := filepath.Join(e.directory, normalizedBookName)
		for _, chapter := range book.Chapters {
			chapterdir := filepath.Join(bookdir, fmt.Sprintf("%s_%d", normalizedBookName, chapter.Number))
			err := os.MkdirAll(chapterdir, os.ModePerm)
			if err != nil {
				return err
			}

			for _, verse := range chapter.Verses {

				versepath := filepath.Join(chapterdir, fmt.Sprintf("%s_%d_%d.yaml", normalizedBookName, chapter.Number, verse.Number))
				chapterWriter, err := os.Create(versepath)
				if err != nil {
					return err
				}

				err = chapterTemplate.Execute(chapterWriter, map[string]any{
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
