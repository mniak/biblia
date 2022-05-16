package text

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mniak/biblia/pkg/bible"
)

type TomlExporter struct {
	directory string
}

func NewTomlExporter(directory string) TomlExporter {
	return TomlExporter{
		directory: directory,
	}
}

func (e TomlExporter) Export(t bible.Testament) error {
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

				versepath := filepath.Join(chapterdir, fmt.Sprintf("%s_%d_%d.toml", normalizedBookName, chapter.Number, verse.Number))
				w, err := os.Create(versepath)
				if err != nil {
					return err
				}
				defer w.Close()

				enc := toml.NewEncoder(w)
				err = enc.Encode(verse)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
