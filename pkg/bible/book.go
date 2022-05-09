package bible

type Book struct {
	Name     string
	Chapters []Chapter
}

func (b Book) Romanize(t Romanizer) Book {
	for idx, chapter := range b.Chapters {
		b.Chapters[idx] = chapter.Romanize(t)
	}
	return b
}
