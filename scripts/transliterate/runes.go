package main

import (
	"unicode/utf8"
)

type _ReverseRuneWalker struct {
	Bytes     []byte
	Rune      rune
	RuneWidth int
}

func ReverseRuneWalker(text string) _ReverseRuneWalker {
	return _ReverseRuneWalker{
		Bytes: []byte(text),
	}
}

func (rw *_ReverseRuneWalker) Walk() bool {
	rw.Rune, rw.RuneWidth = utf8.DecodeLastRune(rw.Bytes)
	rw.Bytes = rw.Bytes[:len(rw.Bytes)-rw.RuneWidth]
	return rw.RuneWidth > 0
}

// func DecodeRunes(text string) []rune {
// 	runes := make([]rune, 0, utf8.RuneCountInString(text))
// 	rw := ReverseRuneWalker(text)
// 	for rw.Walk() {
// 		runes = append(runes, rw.Rune)
// 	}
// 	return runes
// }

// func EncodeRunes(runes []rune) string {
// 	var sb strings.Builder
// 	for _, r := range runes {
// 		sb.WriteRune(r)
// 	}
// 	return sb.String()
// }
