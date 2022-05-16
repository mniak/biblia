package bible

type Verse struct {
	Number int
	Words  []Word
}

func (b Verse) Transliterate(t Transliterator) Verse {
	for idx, word := range b.Words {
		b.Words[idx].Romanized = t.TransliterateWord(word.Text)
	}
	return b
}
