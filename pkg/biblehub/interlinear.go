package biblehub

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type InterlinearChapter struct {
	Title string
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

	topheading := doc.Find("#topheading")
	result.Title = strings.TrimSpace(topheading.Contents().Eq(1).Text())
	return result, nil
}
