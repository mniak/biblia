package bible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookRange_Books(t *testing.T) {
	r := BookRange{Genesis, Deuteronomy}

	var books []BookCode
	for book := range r.Books() {
		books = append(books, book)
	}

	assert.Equal(t, []BookCode{Genesis, Exodus, Leviticus, Numbers, Deuteronomy}, books)
}

func TestBookRange_Contains(t *testing.T) {
	r := Torah

	assert.True(t, r.Contains(Genesis))
	assert.True(t, r.Contains(Exodus))
	assert.True(t, r.Contains(Deuteronomy))
	assert.False(t, r.Contains(Joshua))
	assert.False(t, r.Contains(Matthew))
}

func TestBookRange_Collect(t *testing.T) {
	r := Gospels

	books := r.Collect()

	assert.Equal(t, []BookCode{Matthew, Mark, Luke, John}, books)
}

func TestPredefinedRanges(t *testing.T) {
	tests := []struct {
		name  string
		r     BookRange
		count int
		first BookCode
		last  BookCode
	}{
		{name: "WholeBible", r: WholeBible, count: 66, first: Genesis, last: Revelation},
		{name: "OldTestament", r: OldTestament, count: 39, first: Genesis, last: Malachi},
		{name: "NewTestament", r: NewTestament, count: 27, first: Matthew, last: Revelation},
		{name: "Torah", r: Torah, count: 5, first: Genesis, last: Deuteronomy},
		{name: "Pentateuch", r: Pentateuch, count: 5, first: Genesis, last: Deuteronomy},
		{name: "HistoricalBooksOT", r: HistoricalBooksOT, count: 12, first: Joshua, last: Esther},
		{name: "WisdomBooks", r: WisdomBooks, count: 5, first: Job, last: SongOfSongs},
		{name: "PoetryBooks", r: PoetryBooks, count: 5, first: Job, last: SongOfSongs},
		{name: "MajorProphets", r: MajorProphets, count: 5, first: Isaiah, last: Daniel},
		{name: "MinorProphets", r: MinorProphets, count: 12, first: Hosea, last: Malachi},
		{name: "Prophets", r: Prophets, count: 17, first: Isaiah, last: Malachi},
		{name: "Gospels", r: Gospels, count: 4, first: Matthew, last: John},
		{name: "HistoricalBooksNT", r: HistoricalBooksNT, count: 1, first: Acts, last: Acts},
		{name: "PaulineEpistles", r: PaulineEpistles, count: 0, first: Romans, last: Philemon},
		{name: "GeneralEpistles", r: GeneralEpistles, count: 0, first: Hebrews, last: Jude},
		{name: "Apocalyptic", r: Apocalyptic, count: 1, first: Revelation, last: Revelation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			books := tt.r.Collect()
			assert.Len(t, books, tt.count)
			assert.Equal(t, tt.first, books[0])
			assert.Equal(t, tt.last, books[len(books)-1])
		})
	}
}
