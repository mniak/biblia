package bible

type Romanizer interface {
	RomanizeWord(word string) string
}
