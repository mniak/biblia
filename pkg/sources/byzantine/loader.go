package byzantine

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/pkg/errors"
)

type byzantineLoader struct {
	Directory string
}

func Loader(dir string) byzantineLoader {
	return byzantineLoader{
		Directory: dir,
	}
}

func (l byzantineLoader) Load() (bible.Testament, error) {
	var testament bible.Testament
	for _, bookCode := range bookCodes {
		book, err := l.loadBook(bookCode)
		if err != nil {
			return testament, errors.WithMessagef(err, "failed to load book %s", bookCode)
		}
		testament.Books = append(testament.Books, book)
	}
	return testament, nil
}

func (l byzantineLoader) loadBook(code string) (bible.Book, error) {
	var book bible.Book
	file, err := os.Open(filepath.Join(l.Directory, code+".csv"))
	if err != nil {
		return book, err
	}
	r := csv.NewReader(file)
	var lineCounter int
	var chapter *bible.Chapter
	for {
		lineCounter++
		line, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return book, errors.WithMessage(err, "failed to parse csv")
		}

		if lineCounter == 1 {
			continue // skip header
		}

		if len(line) < 3 {
			return book, fmt.Errorf("too few fields: line=%d, fields=%d", lineCounter, len(line))
		}

		chapterNumber, err := strconv.Atoi(line[0])
		if err != nil {
			return book, errors.WithMessage(err, "chapter number is not a number")
		}
		verseNumber, err := strconv.Atoi(line[1])
		if err != nil {
			return book, errors.WithMessage(err, "verse number is not a number")
		}
		text := line[2]

		if chapter == nil {
			chapter = new(bible.Chapter)
		} else if chapterNumber > chapter.Number {
			book.Chapters = append(book.Chapters, *chapter)
			chapter = new(bible.Chapter)
		}
		chapter.Number = chapterNumber
		chapter.Verses = append(chapter.Verses, bible.Verse{
			Number: verseNumber,
			Words:  strings.Split(text, " "),
		})
	}
	if chapter != nil {
		book.Chapters = append(book.Chapters, *chapter)
	}
	return book, nil
}
