package yaml

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mniak/biblia/pkg/bible"
)

type exportFn func(w io.Writer, book bible.Book, chapter bible.Chapter, verse bible.Verse) error

func export(directory string, t bible.Testament, savefn exportFn) error {
	for _, book := range t.Books {
		normalizedBookName := strings.ReplaceAll(book.Name, " ", "_")
		bookdir := filepath.Join(directory, normalizedBookName)
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

				err = savefn(w, book, chapter, verse)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
