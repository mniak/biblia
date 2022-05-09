package bible

type Testament struct {
	Books []Book
}

func (t Testament) Romanize(tra Romanizer) Testament {
	for idx, book := range t.Books {
		t.Books[idx] = book.Romanize(tra)
	}
	return t
}
