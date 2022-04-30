package wlc

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/samber/lo"
)

func ParseAllBooks(directory string) (bible.Testament, error) {
	var test bible.Testament
	for _, bookname := range BookNames() {
		filename := filepath.Join(directory, bookname+".xml")
		filebytes, err := os.ReadFile(filename)
		if err != nil {
			return test, err
		}
		book, err := ParseBook(filebytes)
		if err != nil {
			return test, err
		}
		test.Books = append(test.Books, book)
	}
	return test, nil
}

func ParseBook(data []byte) (bible.Book, error) {
	var tanachFile TanachFile
	err := xml.Unmarshal(data, &tanachFile)
	if err != nil {
		return bible.Book{}, err
	}
	tanachBook := tanachFile.Tanach.Book
	book := bible.Book{
		Name:     tanachBook.Names.Name,
		Chapters: lo.Map(tanachBook.Chapters, convertChapter),
	}
	return book, nil
}

func convertChapter(ch TanachChapter, i int) bible.Chapter {
	return bible.Chapter{
		Verses: lo.Map[](),
	}
}
