package wlc

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mniak/biblia/internal/utils"
	"github.com/mniak/biblia/pkg/bible"
)

type Loader struct {
	Directory string
}

func NewLoader(dir string) Loader {
	return Loader{
		Directory: dir,
	}
}

func (l Loader) Load() (bible.Testament, error) {
	books, err := utils.MapErr(BookNames(), func(bookname string) (bible.Book, error) {
		bookBytes, err := os.ReadFile(filepath.Join(l.Directory, "Tanach", fmt.Sprintf("%s.xml", bookname)))
		if err != nil {
			return bible.Book{}, err
		}
		return ParseBook(bookBytes)
	})

	return bible.Testament{
		Books: books,
	}, err
}
