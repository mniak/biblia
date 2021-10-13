package main

import (
	"strings"
)

type Transliterator interface {
	Transliterate(word string) string
}

type _Transliterator struct{}

func (t *_Transliterator) Transliterate(word string) string {
	walker := ReverseRuneWalker(word)
	walker.Filter(func(r rune) bool {
		if r >= '\u0591' && r <= '\u05ae' { // Accents
			return false
		}
		if r >= '\u05bd' && r <= '\u05c0' { // Some points
			return false
		}
		if r >= '\u05c3' && r <= '\u05c6' { // More points
			return false
		}
		if r == '\ufffd' {
			return false
		}
		return true
	})

	resultChars := make([]string, 0)
	for walker.Walk() {
		resultChars = append(resultChars, getLastChar(&walker))
	}

	var sb strings.Builder
	for i := len(resultChars) - 1; i >= 0; i-- {
		sb.WriteString(resultChars[i])
	}
	return sb.String()
}

func getLastChar(walker *_RuneWalker) string {
	current := walker.Rune

	// Maitres lectiones
	if entry, ok := maitresLectionesTable[current]; ok {
		if walker.Walk() {
			if char, ok := entry[walker.Rune]; ok {
				return char
			}

			walker.WalkBack()
		}
	}

	// Dagesh
	if current == DAGESH {
		if !walker.Walk() {
			return INVALID
		}

		if char, ok := dageshTable[walker.Rune]; ok {
			return char
		}
		char := getLastChar(walker)
		return char + char
	}

	// Shin
	if current == '\u05c2' || current == '\u05c1' {
		if !walker.Walk() {
			return INVALID
		}

		if walker.Rune == 'ש' {
			return shinTable[current]
		}

		return getLastChar(walker) + INVALID
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
