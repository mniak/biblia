package greek

type simpleTransliterator struct{}

func SimpleTransliterator() *simpleTransliterator {
	return nil
}

func (simpleTransliterator) TransliterateWord(word string) string {
	return word
}
