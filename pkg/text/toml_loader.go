package text

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/mniak/biblia/pkg/bible"
)

func NewTomlLoader(dir string) TomlLoader {
	return TomlLoader{
		Directory: dir,
	}
}

type TomlLoader struct {
	Directory string
}

func (l TomlLoader) Load() (bible.Testament, error) {
	var result bible.Testament
	for _, book := range bible.AllOldTestamentBooks() {
		normalizedBookName := book.NormalizedFileName()
		bookdir := filepath.Join(l.Directory, normalizedBookName)
		for chapter := 1; chapter <= book.ChapterCount(); chapter++ {

			chapterdir := filepath.Join(bookdir, fmt.Sprintf("%s_%d", normalizedBookName, chapter))
			err := os.MkdirAll(chapterdir, os.ModePerm)
			if err != nil {
				return result, err
			}

			for verse := 1; verse <= book.VerseCount(chapter); verse++ {
				versepath := filepath.Join(chapterdir, fmt.Sprintf("%s_%d_%d.toml", normalizedBookName, chapter, verse))
				w, err := os.Create(versepath)
				if err != nil {
					return result, err
				}
				defer w.Close()

				enc := toml.NewEncoder(w)
				err = enc.Encode(verse)
				if err != nil {
					return result, err
				}
			}
		}
	}
	return result, nil

	// defer r.Closer()

	// dec := toml.NewDecoder(r)

	// var result bible.Testament
	// dec.Decode(&result)
}
