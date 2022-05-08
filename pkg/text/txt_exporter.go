package text

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mniak/biblia/pkg/bible"
)

type txtExporter struct {
	directory string
}

func TxtExporter(directory string) txtExporter {
	return txtExporter{
		directory: directory,
	}
}

func (e txtExporter) Export(t bible.Testament) error {
	chapterTemplate, err := template.New("chapter").Parse(``)
	if err != nil {
		return err
	}

	for _, book := range t.Books {
		for _, chapter := range book.Chapters {

			normalizedBookName := strings.ReplaceAll(book.Name, " ", "_")
			dirname := filepath.Join(e.directory, normalizedBookName)

			err := os.MkdirAll(dirname, os.ModePerm)
			if err != nil {
				return err
			}

			chapterWriter, err := os.Create(filepath.Join(dirname, fmt.Sprintf("%s_%d.txt", normalizedBookName, chapter.Number)))
			if err != nil {
				return err
			}

			err = chapterTemplate.Execute(chapterWriter, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
