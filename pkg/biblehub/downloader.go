package biblehub

import (
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

type Downloader interface {
	GetInterlinearChapter(chapter ChapterID) (io.ReadCloser, error)
}

type WebDownloader struct {
	BaseURL string
}

var DefaultDownloader = WebDownloader{
	BaseURL: "https://biblehub.com",
}

func (d WebDownloader) GetInterlinearChapter(chapter ChapterID) (io.ReadCloser, error) {
	client := resty.New().SetBaseURL(d.BaseURL)
	resp, err := client.R().
		SetDoNotParseResponse(true).
		Get(fmt.Sprintf("interlinear/%s/%d.htm", chapter.Book, chapter.Chapter))
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("invalid status %d", resp.StatusCode())
	}
	return resp.RawBody(), nil
}
