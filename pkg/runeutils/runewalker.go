package runeutils

import "unicode/utf8"

type RuneWalker interface {
	Rune() rune
	Walk() bool
	WalkBack() bool
}

type _RuneWalker struct {
	current rune
	runes   []rune
	cursor  int
}

func NewRuneWalker(text string) *_RuneWalker {
	runes := getRunes([]byte(text))
	return NewRuneWalkerFromRunes(runes)
}

func NewReverseRuneWalker(text string) *_RuneWalker {
	runes := getReverseRunes([]byte(text))
	return NewRuneWalkerFromRunes(runes)
}

func NewRuneWalkerFromRunes(runes []rune) *_RuneWalker {
	return &_RuneWalker{
		runes:  runes,
		cursor: -1,
	}
}

func getRunes(bytes []byte) []rune {
	var runes []rune
	r, w := utf8.DecodeRune(bytes)
	for w > 0 {
		runes = append(runes, r)
		bytes = bytes[1:]
		r, w = utf8.DecodeRune(bytes)
	}
	return runes
}

func getReverseRunes(bytes []byte) []rune {
	var runes []rune
	r, w := utf8.DecodeLastRune(bytes)
	for w > 0 {
		runes = append(runes, r)
		bytes = bytes[:len(bytes)-1]
		r, w = utf8.DecodeLastRune(bytes)
	}
	return runes
}

func (rw *_RuneWalker) Rune() rune {
	return rw.current
}

func (rw *_RuneWalker) Filter(validFunc func(rune) bool) {
	filtered := make([]rune, 0, len(rw.runes))
	for _, r := range rw.runes {
		if validFunc(r) {
			filtered = append(filtered, r)
		}
	}
	rw.runes = filtered
}

func (rw *_RuneWalker) Walk() bool {
	if rw.cursor+1 == len(rw.runes) {
		return false
	}
	rw.cursor++
	rw.current = rw.runes[rw.cursor]
	return true
}

func (rw *_RuneWalker) WalkBack() bool {
	if rw.cursor == 0 {
		return false
	}
	rw.cursor--
	rw.current = rw.runes[rw.cursor]
	return true
}
