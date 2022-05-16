package biblehub

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mniak/biblia/pkg/bible"
	"github.com/pkg/errors"
)

type interlinearOTLoader struct{}

func NewInterlinearOldTestamentLoader() bible.TestamentLoader {
	return interlinearOTLoader{}
}

var otBookNames = map[string]int{
	"genesis": 1,
}

func loadOTBook(bookname string, chapterCount int) (bible.Book, error) {
	book := bible.Book{
		Name:     bookname,
		Chapters: make([]bible.Chapter, chapterCount),
	}
	for chapterNumber := 1; chapterNumber <= chapterCount; chapterNumber++ {
		chapter, err := loadOTChapter(bookname, chapterNumber)
		if err != nil {
			return book, errors.Wrapf(err, "failed to load chapter %d", chapterNumber)
		}
		book.Chapters[chapterNumber-1] = chapter
	}
	return book, nil
}

type wordEntry struct {
	verse          int
	strongs        int
	hebrew         string
	english        string
	transliterated string
}

type compositeWordEntry struct {
	verse          []int
	strongs        []int
	hebrew         []string
	english        []string
	transliterated []string
}

func composeWordEntries(words ...wordEntry) compositeWordEntry {
	var result compositeWordEntry
	for idx := len(words) - 1; idx >= 0; idx-- {
		word := words[idx]

		if word.verse != 0 {
			result.verse = append(result.verse, word.verse)
		}
		if word.strongs != 0 {
			result.strongs = append(result.strongs, word.strongs)
		}
		if word.hebrew != "" {
			result.hebrew = append(result.hebrew, word.hebrew)
		}
		if word.english != "" {
			result.english = append(result.english, word.english)
		}
		if word.transliterated != "" {
			result.transliterated = append(result.transliterated, word.transliterated)
		}
	}
	return result
}

func loadOTChapter(bookname string, chapter int) (bible.Chapter, error) {
	result := bible.Chapter{
		Number: chapter,
		Verses: make([]bible.Verse, 0),
	}
	res, err := http.Get(fmt.Sprintf("https://biblehub.com/interlinear/%s/%d.htm", bookname, chapter))
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return result, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return result, err
	}
	var chapterEntries []compositeWordEntry
	doc.Find("table.tablefloatheb").Each(func(i int, s *goquery.Selection) {
		var entries []wordEntry
		s.
			ChildrenFiltered("tbody").
			ChildrenFiltered("tr").
			ChildrenFiltered("td").
			Each(func(i int, s *goquery.Selection) {
				var entry wordEntry
				entry.strongs, _ = strconv.Atoi(normalize(s.Find("span.strongs").First().Text()))
				verseStr := normalize(s.Find("span.refheb").First().Text())
				entry.verse, _ = strconv.Atoi(verseStr)
				entry.hebrew = normalize(s.Find("span.hebrew").Text())
				entry.transliterated = normalize(s.Find("span.translit").Text())
				entry.english = normalize(s.Find("span.eng").Text())

				entries = append(entries, entry)
			})

		chapterEntries = append(chapterEntries, composeWordEntries(entries...))
	})

	var currentVerse *bible.Verse
	for _, entry := range chapterEntries {
		if len(entry.verse) > 0 {
			if currentVerse != nil {
				result.Verses = append(result.Verses, *currentVerse)
			}
			currentVerse = &bible.Verse{
				Words: make([]bible.Word, 0),
			}
			currentVerse.Number = entry.verse[0]
		}

		var word bible.Word
		word.Text = strings.Join(entry.hebrew, " ")
		word.Romanized = strings.Join(entry.transliterated, " ")
		word.English = strings.Join(entry.english, " ")
		currentVerse.Words = append(currentVerse.Words, word)
	}
	if currentVerse != nil {
		result.Verses = append(result.Verses, *currentVerse)
	}

	// for _, verse := range result.Verses {
	// 	fmt.Println(verse.Number, verse.Words)
	// }
	return result, nil
}

func (l interlinearOTLoader) Load() (bible.Testament, error) {
	result := bible.Testament{}
	for bookName, chapterCount := range otBookNames {
		book, err := loadOTBook(bookName, chapterCount)
		if err != nil {
			return result, err
		}
		result.Books = append(result.Books, book)
	}
	return result, nil
}
