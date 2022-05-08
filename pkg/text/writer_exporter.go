package text

import (
	"github.com/mniak/biblia/pkg/bible"
)

type WriterExporter struct {
	Writer TextWriter
}

func StdoutExporter() WriterExporter {
	return WriterExporter{
		Writer: NewIndentedStdout(),
	}
}

func (e WriterExporter) Export(t bible.Testament) error {
	w := NewIndentedWriter(e.Writer)

	for _, book := range t.Books {
		_, err := w.Printlnf("Book of %s", book.Name)
		if err != nil {
			return err
		}
		w.Indent()
		for _, chapter := range book.Chapters {
			_, err = w.Printlnf("Chapter %d (%d vv.)",
				chapter.Number,
				len(chapter.Verses),
			)
			if err != nil {
				return err
			}

			w.Indent()
			for _, verse := range chapter.Verses {
				_, err = w.Printf("%d: ", verse.Number)
				for _, word := range verse.Words {
					w.Print(word)
					w.Print(" ")
				}
				if err != nil {
					return err
				}
				_, err = w.Println()
				if err != nil {
					return err
				}
			}
			_, err = w.Println()
			if err != nil {
				return err
			}
			w.Dedent()
		}
		w.Dedent()
	}
	return nil
}
