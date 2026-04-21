package bible

import (
	"fmt"
	"iter"
)

// VerseRange represents a range of verses within a single chapter
type VerseRange struct {
	Book    BookCode
	Chapter int
	Start   int
	End     int
}

// NewVerseRange creates a VerseRange for a single verse
func NewVerseRange(book BookCode, chapter, verse int) VerseRange {
	return VerseRange{Book: book, Chapter: chapter, Start: verse, End: verse}
}

// NewVerseRangeFromTo creates a VerseRange from start to end
func NewVerseRangeFromTo(book BookCode, chapter, start, end int) VerseRange {
	return VerseRange{Book: book, Chapter: chapter, Start: start, End: end}
}

// Verses returns an iterator over all verse references in the range
func (r VerseRange) Verses() iter.Seq[VerseRef] {
	return func(yield func(VerseRef) bool) {
		for v := r.Start; v <= r.End; v++ {
			if !yield(VerseRef{Book: r.Book, Chapter: r.Chapter, Verse: v}) {
				return
			}
		}
	}
}

// Contains checks if a verse reference is within the range
func (r VerseRange) Contains(ref VerseRef) bool {
	return ref.Book == r.Book &&
		ref.Chapter == r.Chapter &&
		ref.Verse >= r.Start &&
		ref.Verse <= r.End
}

// ContainsVerse checks if a verse number is within the range (same book/chapter assumed)
func (r VerseRange) ContainsVerse(verse int) bool {
	return verse >= r.Start && verse <= r.End
}

// Collect returns all verse references as a slice
func (r VerseRange) Collect() []VerseRef {
	var result []VerseRef
	for v := range r.Verses() {
		result = append(result, v)
	}
	return result
}

// String returns a string representation like "GEN1:1-5" or "GEN1:1" for single verse
func (r VerseRange) String() string {
	if r.Start == r.End {
		return fmt.Sprintf("%s%d:%d", r.Book, r.Chapter, r.Start)
	}
	return fmt.Sprintf("%s%d:%d-%d", r.Book, r.Chapter, r.Start, r.End)
}

// Count returns the number of verses in the range
func (r VerseRange) Count() int {
	return r.End - r.Start + 1
}

// ChapterRef returns the chapter reference for this verse range
func (r VerseRange) ChapterRef() ChapterRef {
	return ChapterRef{Book: r.Book, Chapter: r.Chapter}
}

// Well-known verse ranges
var (
	// Ten Commandments
	TenCommandments = VerseRange{Exodus, 20, 1, 17}

	// Lord's Prayer
	LordsPrayer = VerseRange{Matthew, 6, 9, 13}

	// Beatitudes
	Beatitudes = VerseRange{Matthew, 5, 3, 12}

	// Love Chapter
	LoveChapter = VerseRange{ICorinthians, 13, 1, 13}

	// Armor of God
	ArmorOfGod = VerseRange{Ephesians, 6, 10, 18}

	// Fruit of the Spirit
	FruitOfTheSpirit = VerseRange{Galatians, 5, 22, 23}

	// Great Commission
	GreatCommission = VerseRange{Matthew, 28, 18, 20}

	// John 3:16
	John316 = VerseRange{John, 3, 16, 16}

	// Psalm 23
	Psalm23 = VerseRange{Psalms, 23, 1, 6}

	// Romans Road
	RomansRoad323 = VerseRange{Romans, 3, 23, 23}
	RomansRoad623 = VerseRange{Romans, 6, 23, 23}
	RomansRoad58  = VerseRange{Romans, 5, 8, 8}
	RomansRoad109 = VerseRange{Romans, 10, 9, 10}
)
