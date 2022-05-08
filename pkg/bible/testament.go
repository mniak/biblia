package bible

type Testament struct {
	Books []Book
}

func (t Testament) Transliterate(tra Transliterator) Testament {
	for idx, book := range t.Books {
		t.Books[idx] = book.Transliterate(tra)
	}
	return t
}
