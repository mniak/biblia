package greek

type noopTransliterator struct{}

func NoopTransliterator() *noopTransliterator {
	return &noopTransliterator{}
}

func (noopTransliterator) TransliterateWord(word string) string {
	return word
}
