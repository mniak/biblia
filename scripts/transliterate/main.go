package main

import (
	"encoding/xml"
	"log"
	"os"
)

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	data, err := os.ReadFile("../../base/UnicodeXML_Westminster_Leningrad_Codex/Tanach/Genesis.xml")
	handle(err)

	w := NewIndentedStdout()
	transliterator := _Transliterator{}

	var tanach Tanach
	err = xml.Unmarshal(data, &tanach)
	handle(err)

	book := tanach.Tanach.Book
	w.Printlnf("Book of %s", book.Names.Name)
	w.Indent()
	for _, chapter := range book.Chapters[:1] {
		w.Printlnf("Chapter %d (%d vv.)",
			chapter.Number,
			chapter.VerseCount,
		)

		w.Indent()
		for _, verse := range chapter.Verses[:10] {
			w.Printf("%d: ", verse.Number)
			for _, word := range verse.Words {
				word.Text = transliterator.Transliterate(word.Text)
				w.Print(word.Text)
				w.Print(" ")
			}
			w.Println()
		}
		w.Dedent()
	}
	w.Dedent()
}
