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
	return convertBook(tanachFile.Tanach.Book), nil
}

func convertBook(book TanachBook) bible.Book {
	return bible.Book{
		Name:     book.Names.Name,
		Chapters: lo.Map(book.Chapters, convertChapter),
	}
}

func convertChapter(chapter TanachChapter, i int) bible.Chapter {
	return bible.Chapter{
		Number: chapter.Number,
		Verses: lo.Map(chapter.Verses, convertVerse),
	}
}

func convertVerse(verse TanachVerse, i int) bible.Verse {
	return bible.Verse{
		Number: verse.Number,
		Words:  lo.Map(verse.Words, convertWord),
	}
}

func convertWord(word TanachWord, i int) string {
	return word.Text
}
