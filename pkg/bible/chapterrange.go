package bible

import (
	"fmt"
	"iter"
)

type ChapterRange struct {
	Book  BookCode
	Start int
	End   int
}

func NewChapterRange(book BookCode, chapter int) ChapterRange {
	return ChapterRange{Book: book, Start: chapter, End: chapter}
}

func NewChapterRangeFromTo(book BookCode, start, end int) ChapterRange {
	return ChapterRange{Book: book, Start: start, End: end}
}

func (r ChapterRange) Chapters() iter.Seq[ChapterRef] {
	return func(yield func(ChapterRef) bool) {
		for ch := r.Start; ch <= r.End; ch++ {
			if !yield(ChapterRef{Book: r.Book, Chapter: ch}) {
				return
			}
		}
	}
}

func (r ChapterRange) Contains(ref ChapterRef) bool {
	return ref.Book == r.Book && ref.Chapter >= r.Start && ref.Chapter <= r.End
}

func (r ChapterRange) ContainsChapter(chapter int) bool {
	return chapter >= r.Start && chapter <= r.End
}

func (r ChapterRange) Collect() []ChapterRef {
	var result []ChapterRef
	for ch := range r.Chapters() {
		result = append(result, ch)
	}
	return result
}

func (r ChapterRange) String() string {
	if r.Start == r.End {
		return fmt.Sprintf("%s%d", r.Book, r.Start)
	}
	return fmt.Sprintf("%s%d-%d", r.Book, r.Start, r.End)
}

func (r ChapterRange) Count() int {
	return r.End - r.Start + 1
}
