package bible

type Verse struct {
	Number int
	Words  []string
}

func (b Verse) Transliterate(t Transliterator) Verse {
	for idx, word := range b.Words {
		b.Words[idx] = t.TransliterateWord(word)
	}
	return b
}
