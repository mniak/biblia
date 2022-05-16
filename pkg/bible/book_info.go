package bible

import "strings"

type OldTestamentBook string

var (
	allOtBooks  []OldTestamentBook
	otBooksData map[OldTestamentBook][]int
)

func AllOldTestamentBooks() []OldTestamentBook {
	return allOtBooks
}

func (otb OldTestamentBook) NormalizedFileName() string {
	value := string(otb)
	value = strings.ReplaceAll(value, " ", "_")
	value = strings.ToLower(value)
	return value
}

func (otb OldTestamentBook) ChapterCount() int {
	return len(otBooksData[otb])
}

func (otb OldTestamentBook) VerseCount(chapter int) int {
	return otBooksData[otb][chapter-1]
}
