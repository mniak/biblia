package bible

type Chapter struct {
	Number int
	Verses []Verse
}

func (c Chapter) Romanize(t Romanizer) Chapter {
	for idx, verse := range c.Verses {
		c.Verses[idx] = verse.Romanize(t)
	}
	return c
}
