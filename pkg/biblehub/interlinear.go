package biblehub

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type InterlinearChapter struct {
	Title  string
	Verses []InterlinearVerse
}

type InterlinearVerse struct {
	Number int
	Words  []InterlinearWord
}
type InterlinearWord struct {
	StrongsNumber   string
	StrongsText     string
	Transliteration string
	English         string
	Hebrew          string
	Syntax          string
}

func GetInterlinearChapter(book string, chapter int) (InterlinearChapter, error) {
	var result InterlinearChapter
	client := resty.New().SetBaseURL("https://biblehub.com")
	resp, err := client.R().
		SetDoNotParseResponse(true).
		Get(fmt.Sprintf("interlinear/%s/%d.htm", book, chapter))
	if err != nil {
		return result, err
	}
	if !resp.IsSuccess() {
		return result, fmt.Errorf("invalid status %d", resp.StatusCode())
	}

	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return result, err
	}

	result.Title = strings.TrimSpace(doc.Find("#topheading").Children().Remove().End().Text())
	var currentVerse *InterlinearVerse
	var currentWordNumber int
	doc.Find(".tablefloatheb > tbody > tr > td").Each(func(i1 int, s1 *goquery.Selection) {
		spans := s1.Find("span")
		currentWordNumber++

		var word InterlinearWord

		strongs1 := spans.Filter(".strongs").Find("a").First()
		// strongs2 := strongs1.Next()
		word.StrongsNumber = strings.TrimSpace(strongs1.Text())
		word.StrongsText = sanitizeEnglish(strings.TrimSpace(strongs1.AttrOr("title", "")))

		word.Transliteration = strings.TrimSpace(spans.Filter(".translit").Text())
		word.English = sanitizeEnglish(strings.TrimSpace(spans.Filter(".eng").Text()))
		word.Hebrew = strings.TrimSpace(spans.Filter(".hebrew").Text())
		word.Syntax = strings.TrimSpace(spans.Filter(".strongsnt").Text())
		spans.Filter(".refheb").First().Each(func(i int, refheb *goquery.Selection) {
			if currentVerse != nil {
				result.Verses = append(result.Verses, *currentVerse)
			}
			currentVerse = new(InterlinearVerse)
			hebrewNumber, _ := strconv.Atoi(strings.TrimSpace(refheb.Text()))
			currentVerse.Number = hebrewNumber
			currentWordNumber = 0
		})
		if currentVerse.Number == 49 {
			fmt.Printf("[%d = %s] ", currentWordNumber, word.English)
		}
		currentVerse.Words = append(currentVerse.Words, word)
	})

	if currentVerse != nil {
		result.Verses = append(result.Verses, *currentVerse)
	}
	return result, nil
}

func sanitizeEnglish(text string) string {
	text = strings.ReplaceAll(text, "\u00a0", " ")
	text = strings.ReplaceAll(text, "<BR>", "\n")
	return text
}
