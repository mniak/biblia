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
	basicTable := map[rune]string{
		'א': "ʾ",
		'ב': "ḇ",
		'ג': "ḡ",
		'ד': "ḏ",
		'ה': "h",
		'ו': "w",
		'ז': "z",
		'ח': "ḥ",
		'ט': "ṭ",
		'י': "y",
		'כ': "ḵ",
		'ך': "ḵ",
		'ל': "l",
		'מ': "m",
		'ם': "m",
		'נ': "n",
		'ן': "n",
		'ס': "s",
		'ע': "ʿ",
		'פ': "p̄",
		'ף': "p̄",
		'צ': "ṣ",
		'ץ': "ṣ",
		'ק': "q",
		'ר': "r",
		'ש': "š",
		'ת': "ṯ",

		// Vowels
		QAMATS:   "ā",
		PATAH:    "a",
		SEGOL:    "e",
		TSERE:    "ē",
		HIRIK:    "i",
		HOLAM:    "o",
		QUBUTS:   "u",
		SHEVA:    "'",
		'\u05b2': "ă", // HATAF_PATAH
		'\u05b3': "ŏ", // HATAF QAMATS
		'\u05b1': "ĕ", // HATAF_SEGOL

		// Punctuation
		// Hebrew 		|	Latin
		// -------------|-----------------
		// maqaf		|	hyphen
		// geresh		|	apostrophe
		// gershayim	|	quotation mark
		// meteg		|	comma
		// inverted nun	|	bracket
		// Source: https://en.wikipedia.org/wiki/Hebrew_punctuation

		'\u05bd': ",", // Meteg
		'\u05be': "-", // Maqaf
	}

	dageshTable := map[rune]string{
		'ב': "b",
		'ג': "g",
		'ד': "d",
		'כ': "k",
		'פ': "p",
		'ת': "t",
	}

	shinTable := map[rune]string{
		'\u05c2': "ś",
		'\u05c1': "š",
	}

	materLectionisTable := map[rune]map[rune]string{
		'ה': {
			QAMATS: "â",
		},
		'י': {
			TSERE: "ê",
			SEGOL: "ê",
			HIRIK: "î",
		},
		HOLAM: {
			'ו': "ô",
		},
		DAGESH: {
			'ו': "û",
		},
	}

	current := walker.Rune()

	// Mater Lectionis
	if entry, ok := materLectionisTable[current]; ok {
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
