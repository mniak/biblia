package biblehub

import (
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

type Downloader interface {
	GetInterlinearChapter(chapter ChapterID) (io.ReadCloser, error)
}

type _Downloader struct {
	BaseURL string
}

func (_Downloader) GetInterlinearChapter(chapter ChapterID) (io.ReadCloser, error) {
	client := resty.New().SetBaseURL("https://biblehub.com")
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
