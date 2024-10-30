package biblehub

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type InterlinearChapter struct {
	Title  string
	Verses []InterlinearVerse
}

type InterlinearVerse struct {
	Transliteration string
	Words           []InterlinearWord
}
type InterlinearWord struct {
	Strongs         string
	Transliteration string
	English         string
	Hebrew          string
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
	doc.Find(".tablefloatheb > tbody > tr > td").Each(func(i int, s1 *goquery.Selection) {
		s1c := s1.Children()

		s1c.Find(".refheb").Each(func(i int, s2 *goquery.Selection) {
			if currentVerse != nil {
				result.Verses = append(result.Verses, *currentVerse)
			}
			currentVerse = new(InterlinearVerse)
			fmt.Printf("\n\n%s: ", s2.Text())
		})

		word := InterlinearWord{
			Strongs:         strings.TrimSpace(s1c.Filter(".strongs").First().Text()),
			Transliteration: strings.TrimSpace(s1c.Filter(".translit").Text()),
			English:         strings.TrimSpace(s1c.Filter(".eng").Text()),
			Hebrew:          strings.TrimSpace(s1c.Filter(".hebrew").Text()),
		}
		if word.English != "" {
			fmt.Print(word.English + " ")
		}
		currentVerse.Words = append(currentVerse.Words, word)
	})

	if currentVerse != nil {
		result.Verses = append(result.Verses, *currentVerse)
	}
	return result, nil
}
