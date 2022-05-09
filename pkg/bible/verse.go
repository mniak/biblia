package bible

type Verse struct {
	Number int
	Words  []Word
}

func (b Verse) Romanize(t Romanizer) Verse {
	for idx, word := range b.Words {
		b.Words[idx].Romanized = t.RomanizeWord(word.Text)
	}
	return b
}
