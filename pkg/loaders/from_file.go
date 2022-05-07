package loaders

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mniak/biblia/internal/utils"
	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/sources/wlc"
)

type WLCLoader struct {
	Directory string
}

func (l WLCLoader) Load() (bible.Testament, error) {
	books, err := utils.MapErr(wlc.BookNames(), func(bookname string) (bible.Book, error) {
		bookBytes, err := os.ReadFile(filepath.Join(l.Directory, fmt.Sprintf("%s.xml", bookname)))
		if err != nil {
			return bible.Book{}, err
		}
		return wlc.ParseBook(bookBytes)
	})

	return bible.Testament{
		Books: books,
	}, err
}
