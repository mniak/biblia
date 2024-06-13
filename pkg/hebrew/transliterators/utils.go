package transliterators

import "github.com/mniak/biblia/pkg/runeutils"

func getLastChar(walker runeutils.RuneWalker) string {
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
		char := getLastChar(walker)
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

		return getLastChar(walker) + INVALID
	}

	if char, ok := basicTable[current]; ok {
		return char
	}

	// return string(walker.Rune)
	// return fmt.Sprintf("<\\u%04x>", walker.Rune)
	return ""
}
