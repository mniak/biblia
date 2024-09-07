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
	if current == rune(DAGESH) {
		if !walker.Walk() {
			return string(InvalidChar)
		}

		// Dagesh Forte - BeGaD KePhaT (phonetic phenomenon)
		if char, isBegadKephat := dageshTable[Letter(walker.Rune())]; isBegadKephat {
			return char
		}

		// Dagesh Lene - Double
		char := getLastChar(walker)
		return char + char
	}

	// Shin
	const SinDot = 'a'
	const ShinDot = 'b'
	if current == '\u05c2' || current == '\u05c1' {
		if !walker.Walk() {
			return string(InvalidChar)
		}

		if walker.Rune() == 'ש' {
			return shinTable[current]
		}

		return getLastChar(walker) + string(InvalidChar)
	}

	if char := basicConvert(current); char != string(InvalidChar) {
		return char
	}

	return ""
}
