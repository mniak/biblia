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

//go:embed chapter_template.tmpl
var embedfs embed.FS

type txtExporter struct {
	directory string
}

func TxtExporter(directory string) txtExporter {
	return txtExporter{
		directory: directory,
	}
}

func (e txtExporter) Export(t bible.Testament) error {
	chapterTemplate, err := template.ParseFS(embedfs, "chapter_template.tmpl")
	if err != nil {
		return err
	}

	for _, book := range t.Books {
		normalizedBookName := strings.ReplaceAll(book.Name, " ", "_")
		dirname := filepath.Join(e.directory, normalizedBookName)

		for _, chapter := range book.Chapters {

			err := os.MkdirAll(dirname, os.ModePerm)
			if err != nil {
				return err
			}

			chapterWriter, err := os.Create(filepath.Join(dirname, fmt.Sprintf("%s_%d.txt", normalizedBookName, chapter.Number)))
			if err != nil {
				return err
			}

			err = chapterTemplate.Execute(chapterWriter, map[string]any{
				"Book":    book,
				"Chapter": chapter,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
