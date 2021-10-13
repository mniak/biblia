package main

import (
	"strings"
)

type Transliterator interface {
	Transliterate(word string) string
}

type _Transliterator struct{}

func ignore(r rune) bool {
	if r >= '\u0591' && r <= '\u05ae' { // Accents
		return true
	}
	if r >= '\u05bd' && r <= '\u05c0' { // Some points
		return true
	}
	if r >= '\u05c3' && r <= '\u05c6' { // More points
		return true
	}

	return false
}

func walk(walker *_ReverseRuneWalker) bool {
	if !walker.Walk() {
		return false
	}
	r := walker.Rune
	if ignore(r) {
		return walk(walker)
	}
	return true
}

func (t *_Transliterator) Transliterate(word string) string {
	walker := ReverseRuneWalker(word)
	resultWords := make([]string, 0)
	for walk(&walker) {
		resultWords = append(resultWords, getLastChar(&walker))
	}

	var sb strings.Builder
	for i := len(resultWords) - 1; i >= 0; i-- {
		sb.WriteString(resultWords[i])
	}
	return sb.String()
}

func getLastChar(walker *_ReverseRuneWalker) string {
	current := walker.Rune
	if current == DAGESH {
		if !walk(walker) {
			return INVALID
		}

		if char, ok := dageshTable[walker.Rune]; ok {
			return char
		}
		char := getLastChar(walker)
		return char + char
	}

	if current == '\u05c2' || current == '\u05c1' {
		if !walk(walker) {
			return INVALID
		}

		if walker.Rune == 'ש' {
			return shinTable[current]
		}

		return getLastChar(walker) + INVALID
	}

	if entry, ok := maitresLectionesTable[current]; ok {
		if !walk(walker) {
			return getBasic(current)
		}

		if char, ok := entry[walker.Rune]; ok {
			return char
		}

		return getLastChar(walker) + getBasic(current)
	}

	if char, ok := basicTable[current]; ok {
		return char
	}

	// return string(walker.Rune)
	// return fmt.Sprintf("<\\u%04x>", walker.Rune)
	return ""
}

func getBasic(r rune) string {
	if char, ok := basicTable[r]; ok {
		return char
	}

	// return string(walker.Rune)
	// return fmt.Sprintf("<\\u%04x>", walker.Rune)
	return ""
}
