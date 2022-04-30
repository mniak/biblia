package bible

type Transliterator interface {
	TransliterateWord(word string) string
}
