package biblehubtest

import (
	"embed"
	"fmt"
	"io"
)

//go:embed fake_downloader/*.htm
var fakeDownloaderFS embed.FS

const FakeDownloader fakeDownloader = false

type fakeDownloader bool

func (fakeDownloader) GetInterlinearChapter(book string, chapter int) (io.ReadCloser, error) {
	return fakeDownloaderFS.Open(fmt.Sprintf("fake_downloader/interlinear_%s_%d.htm", book, chapter))
}
