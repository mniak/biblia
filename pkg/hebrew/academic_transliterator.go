package hebrew

import (
	"strings"

	"github.com/mniak/biblia/pkg/runeutils"
)

type academicTransliterator struct{}

func AcademicTransliterator() *academicTransliterator {
	return &academicTransliterator{}
}

func (t *academicTransliterator) TransliterateWord(word string) string {
	walker := runeutils.NewReverseRuneWalker(word)
	walker.Filter(func(r rune) bool {
		_, ignored := ignoredSet[r]
		return !ignored
	})

	resultChars := make([]string, 0)
	for walker.Walk() {
		resultChars = append(resultChars, t.getLastChar(walker))
	}

	var sb strings.Builder
	for i := len(resultChars) - 1; i >= 0; i-- {
		sb.WriteString(resultChars[i])
	}
	return sb.String()
}

func (t *academicTransliterator) getLastChar(walker runeutils.RuneWalker) string {
	current := walker.Rune()

	// Maitres lectiones
	if entry, ok := maitresLectionesTable[current]; ok {
		if walker.Walk() {
			if char, ok := entry[walker.Rune()]; ok {
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

		if char, ok := dageshTable[walker.Rune()]; ok {
			return char
		}
		char := t.getLastChar(walker)
		return char + char
	}

	// Shin
	if current == '\u05c2' || current == '\u05c1' {
		if !walker.Walk() {
			return INVALID
		}

		if walker.Rune() == 'ש' {
			return shinTable[current]
		}

		return t.getLastChar(walker) + INVALID
	}

	if char, ok := basicTable[current]; ok {
		return char
	}

	return ""
}
