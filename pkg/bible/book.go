package bible

type Book struct {
	Name     string
	Chapters []Chapter
}

func (b Book) Transliterate(t Transliterator) Book {
	for idx, chapter := range b.Chapters {
		b.Chapters[idx] = chapter.Transliterate(t)
	}
	return b
}
