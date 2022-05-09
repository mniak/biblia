package yaml

import (
	"io"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

type yamlExporter struct {
	directory string
}

func NewExporter(directory string) yamlExporter {
	return yamlExporter{
		directory: directory,
	}
}

func (e yamlExporter) Export(t bible.Testament) error {
	return export(e.directory, t, func(w io.Writer, book bible.Book, chapter bible.Chapter, verse bible.Verse) error {
		enc := yaml.NewEncoder(w)
		defer enc.Close()

		return enc.Encode(verseModel{
			BookName: book.Name,
			Chapter:  chapter.Number,
			Verse:    verse.Number,
			Words: lo.Map(verse.Words, func(w bible.Word, i int) wordModel {
				return wordModel{
					Text:      w.Text,
					Romanized: w.Romanized,
				}
			}),
		})
	})
}
