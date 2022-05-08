package bible

type Chapter struct {
	Number int
	Verses []Verse
}

func (c Chapter) Transliterate(t Transliterator) Chapter {
	for idx, verse := range c.Verses {
		c.Verses[idx] = verse.Transliterate(t)
	}
	return c
}
