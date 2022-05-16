package biblehub

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/mniak/biblia/pkg/bible"
)

type interlinearOTLoader struct{}

func NewInterlinearOldTestamentLoader() bible.TestamentLoader {
	return interlinearOTLoader{}
}

var otBookNames = map[string]int{
	"genesis": 50,
}

func loadOTBook(bookname string, chapterCount int) (bible.Book, error) {
	result := bible.Book{}
	res, err := http.Get(fmt.Sprintf("https://biblehub.com/interlinear/%s/1.htm", bookname))
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
	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
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
}
