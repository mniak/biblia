package bible

type Bible struct {
	Books []Book
}

func (b Bible) Transliterate(t Transliterator) Bible {
	for idx, book := range b.Books {
		b.Books[idx] = book.Transliterate(t)
	}
	return b
}
