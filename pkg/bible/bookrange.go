package bible

import "iter"

type BookRange struct {
	Start BookCode
	End   BookCode
}

func (r BookRange) Books() iter.Seq[BookCode] {
	return func(yield func(BookCode) bool) {
		startIdx := r.Start.Index()
		endIdx := r.End.Index()
		if startIdx == 0 || endIdx == 0 {
			return
		}
		for i := startIdx; i <= endIdx; i++ {
			if !yield(i.BookCode()) {
				return
			}
		}
	}
}

func (r BookRange) Contains(code BookCode) bool {
	idx := code.Index()
	return idx >= r.Start.Index() && idx <= r.End.Index()
}

func (r BookRange) Collect() []BookCode {
	var result []BookCode
	for book := range r.Books() {
		result = append(result, book)
	}
	return result
}

var (
	WholeBible = BookRange{Genesis, Revelation}

	OldTestament      = BookRange{Genesis, Malachi}
	Torah             = BookRange{Genesis, Deuteronomy}
	Pentateuch        = Torah
	HistoricalBooksOT = BookRange{Joshua, Esther}
	WisdomBooks       = BookRange{Job, SongOfSongs}
	PoetryBooks       = WisdomBooks
	MajorProphets     = BookRange{Isaiah, Daniel}
	MinorProphets     = BookRange{Hosea, Malachi}
	Prophets          = BookRange{Isaiah, Malachi}

	NewTestament      = BookRange{Matthew, Revelation}
	Gospels           = BookRange{Matthew, John}
	HistoricalBooksNT = BookRange{Acts, Acts}
	PaulineEpistles   = BookRange{Romans, Philemon}
	GeneralEpistles   = BookRange{Hebrews, Jude}
	Apocalyptic       = BookRange{Revelation, Revelation}
)
