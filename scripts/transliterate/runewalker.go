package main

import "unicode/utf8"

type _RuneWalker struct {
	Rune rune

	runes  []rune
	cursor int
}

func RuneWalker(text string) _RuneWalker {
	runes := getRunes([]byte(text))
	return RuneWalkerFromRunes(runes)
}

func ReverseRuneWalker(text string) _RuneWalker {
	runes := getReverseRunes([]byte(text))
	return RuneWalkerFromRunes(runes)
}

func RuneWalkerFromRunes(runes []rune) _RuneWalker {
	return _RuneWalker{
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

func (rw *_RuneWalker) Walk() bool {
	if rw.cursor+1 == len(rw.runes) {
		return false
	}
	rw.cursor++
	rw.Rune = rw.runes[rw.cursor]
	return true
}

func (rw *_RuneWalker) WalkBack() bool {
	if rw.cursor == 0 {
		return false
	}
	rw.cursor--
	rw.Rune = rw.runes[rw.cursor]
	return true
}
