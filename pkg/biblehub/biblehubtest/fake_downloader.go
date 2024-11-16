package biblehubtest

import (
	"embed"
	"fmt"
	"io"

	"github.com/mniak/biblia/pkg/biblehub"
)

//go:embed fake_downloader/*.htm
var fakeDownloaderFS embed.FS

const FakeDownloader fakeDownloader = false

type fakeDownloader bool

func (fakeDownloader) GetInterlinearChapter(chapter biblehub.ChapterID) (io.ReadCloser, error) {
	return fakeDownloaderFS.Open(fmt.Sprintf("fake_downloader/interlinear_%s_%d.htm", chapter.Book, chapter.Chapter))
}
