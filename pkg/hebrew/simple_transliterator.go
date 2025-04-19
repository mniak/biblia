package hebrew

import (
	"strings"

	"github.com/mniak/biblia/pkg/runeutils"
)

type simpleTransliterator struct{}

func SimpleTransliterator() *simpleTransliterator {
	return &simpleTransliterator{}
}

func (t *simpleTransliterator) TransliterateWord(word string) string {
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

func (t *simpleTransliterator) getLastChar(walker runeutils.RuneWalker) string {
	basicTable := map[rune]string{
		'א': "",
		'ב': "b",
		'ג': "g",
		'ד': "d",
		'ה': "h",
		'ו': "w",
		'ז': "z",
		'ח': "ch",
		'ט': "t",
		'י': "y",
		'כ': "kh",
		'ך': "k",
		'ל': "l",
		'מ': "m",
		'ם': "m",
		'נ': "n",
		'ן': "n",
		'ס': "s",
		'ע': "",
		'פ': "p",
		'ף': "p",
		'צ': "ʦ",
		'ץ': "ts",
		'ק': "q",
		'ר': "r",
		'ש': "sh",
		'ת': "t",

		// Vowels
		QAMATS:   "a",
		PATAH:    "a",
		SEGOL:    "e",
		TSERE:    "e",
		HIRIK:    "i",
		HOLAM:    "o",
		QUBUTS:   "u",
		SHEVA:    "ᵉ",
		'\u05b2': "a", // HATAF_PATAH
		'\u05b3': "o", // HATAF QAMATS
		'\u05b1': "e", // HATAF_SEGOL

		// Punctuation
		// Hebrew 		|	Latin
		// -------------|-----------------
		// maqaf		|	hyphen
		// geresh		|	apostrophe
		// gershayim	|	quotation mark
		// meteg		|	comma
		// inverted nun	|	bracket
		// Source: https://en.wikipedia.org/wiki/Hebrew_punctuation

		// '\u05bd': ",", // Meteg
		// '\u05be': "-", // Maqaf
	}

	shinTable := map[rune]string{
		'\u05c2': "s",
		'\u05c1': "sh",
	}

	current := walker.Rune()

	// Mater Lectionis
	materLectionisTable := map[rune]map[rune]string{
		'ה': {
			QAMATS: "a",
		},
		'י': {
			TSERE: "e",
			SEGOL: "e",
			HIRIK: "i",
		},
		HOLAM: {
			'ו': "o",
		},
		DAGESH: {
			'ו': "u",
		},
	}
	if entry, ok := materLectionisTable[current]; ok {
		if walker.Walk() {
			if char, ok := entry[walker.Rune()]; ok {
				return char
			}

			walker.WalkBack()
		}
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
