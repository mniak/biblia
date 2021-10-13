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
		_, ignored := ignoredSet[r]
		return !ignored
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

	// return fmt.Sprintf("<\\u%04x>", walker.Rune)
	return ""
}
