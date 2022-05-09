package text

import (
	"io"

	"github.com/mniak/biblia/pkg/bible"
)

type textExporter struct {
	writer *_IndentedTextWriter
}

func NewExporter(writer io.StringWriter) textExporter {
	return textExporter{
		writer: newAutoIndentWriter(writer),
	}
}

func (e textExporter) Export(t bible.Testament) error {
	for _, book := range t.Books {
		_, err := e.writer.Printf("Book of %s\n", book.Name)
		if err != nil {
			return err
		}
		e.writer.Indent()
		for _, chapter := range book.Chapters {
			_, err = e.writer.Printf("Chapter %d (%d vv.)\n", chapter.Number, len(chapter.Verses))
			if err != nil {
				return err
			}

			e.writer.Indent()
			for _, verse := range chapter.Verses {
				_, err = e.writer.Printf("%d: ", verse.Number)
				for _, word := range verse.Words {
					e.writer.Print(word.Romanized)
					e.writer.Print(" ")
				}
				if err != nil {
					return err
				}
				_, err = e.writer.Println()
				if err != nil {
					return err
				}
			}
			_, err = e.writer.Println()
			if err != nil {
				return err
			}
			e.writer.Dedent()
		}
		e.writer.Dedent()
	}
	return nil
}
