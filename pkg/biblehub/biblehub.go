package biblehub

import (
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

func NewExtractor() _Extractor {
	return _Extractor{}
}

type Downloader interface {
	GetInterlinearChapter(book string, chapter int) (io.ReadCloser, error)
}

type _Extractor struct {
	Downloader Downloader
}

type _Downloader struct {
	BaseURL string
}

func (_Downloader) GetInterlinearChapter(book string, chapter int) (io.ReadCloser, error) {
	client := resty.New().SetBaseURL("https://biblehub.com")
	resp, err := client.R().
		SetDoNotParseResponse(true).
		Get(fmt.Sprintf("interlinear/%s/%d.htm", book, chapter))
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("invalid status %d", resp.StatusCode())
	}
	return resp.RawBody(), nil
}
