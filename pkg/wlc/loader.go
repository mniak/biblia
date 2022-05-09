package wlc

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mniak/biblia/internal/utils"
	"github.com/mniak/biblia/pkg/bible"
)

type wlcLoader struct {
	Directory string
}

func NewLoader(dir string) wlcLoader {
	return wlcLoader{
		Directory: dir,
	}
}

func (l wlcLoader) Load() (bible.Testament, error) {
	books, err := utils.MapErr(BookNames(), func(bookname string) (bible.Book, error) {
		bookBytes, err := os.ReadFile(filepath.Join(l.Directory, fmt.Sprintf("%s.xml", bookname)))
		if err != nil {
			return bible.Book{}, err
		}
		return ParseBook(bookBytes)
	})

	return bible.Testament{
		Books: books,
	}, err
}
