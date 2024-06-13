package transliterators

const (
	INVALID = "�"
)

type Vowel rune

const (
	SHEVA     Vowel = '\u05b0'
	HIRIK     Vowel = '\u05b4'
	TSERE     Vowel = '\u05b5'
	SEGOL     Vowel = '\u05b6'
	PATAH     Vowel = '\u05b7'
	QAMATS    Vowel = '\u05b8'
	HOLAM     Vowel = '\u05b9'
	HOLAM_VAV Vowel = '\u05ba'
	QUBUTS    Vowel = '\u05bb'
	DAGESH    Vowel = '\u05bc'

	HATAF_SEGOL  Vowel = '\u05b1'
	HATAF_PATAH  Vowel = '\u05b2'
	HATAF_QAMATS Vowel = '\u05b3'
)

var vowels = []Vowel{
	SHEVA,
	DAGESH,
	QAMATS,
	PATAH,
	SEGOL,
	TSERE,
	HIRIK,
	HOLAM,
	HOLAM_VAV,
	QUBUTS,
	HATAF_SEGOL,
	HATAF_PATAH,
	HATAF_QAMATS,
}

func basicConvert(r rune) string {
	switch r {
	case 'א':
		return "ʾ"
	case 'ב':
		return "ḇ"
	case 'ג':
		return "ḡ"
	case 'ד':
		return "ḏ"
	case 'ה':
		return "h"
	case 'ו':
		return "w"
	case 'ז':
		return "z"
	case 'ח':
		return "ḥ"
	case 'ט':
		return "ṭ"
	case 'י':
		return "y"
	case 'כ':
		return "ḵ"
	case 'ך':
		return "ḵ"
	case 'ל':
		return "l"
	case 'מ':
		return "m"
	case 'ם':
		return "m"
	case 'נ':
		return "n"
	case 'ן':
		return "n"
	case 'ס':
		return "s"
	case 'ע':
		return "ʿ"
	case 'פ':
		return "p̄"
	case 'ף':
		return "p̄"
	case 'צ':
		return "ṣ"
	case 'ץ':
		return "ṣ"
	case 'ק':
		return "q"
	case 'ר':
		return "r"
	case 'ש':
		return "š"
	case 'ת':
		return "ṯ"

	// Vowels
	case rune(QAMATS):
		return "ā"
	case rune(PATAH):
		return "a"
	case rune(SEGOL):
		return "e"
	case rune(TSERE):
		return "ē"
	case rune(HIRIK):
		return "i"
	case rune(HOLAM):
		return "o"
	case rune(QUBUTS):
		return "u"
	case rune(SHEVA):
		return "'"
	case rune(HATAF_SEGOL):
		return "ĕ"
	case rune(HATAF_PATAH):
		return "ă"
	case rune(HATAF_QAMATS):
		return "ŏ"

	// Punctuation
	// Hebrew 		|	Latin
	// -------------|-----------------
	// maqaf		|	hyphen
	// geresh		|	apostrophe
	// gershayim	|	quotation mark
	// meteg		|	comma
	// inverted nun	|	bracket
	// Source: https://en.wikipedia.org/wiki/Hebrew_punctuation

	case '\u05bd':
		return "," // Meteg
	case '\u05be':
		return "-" // Maqaf
	}
	return INVALID
}

var dageshTable = map[rune]string{
	'ב': "b",
	'ג': "g",
	'ד': "d",
	'כ': "k",
	'פ': "p",
	'ת': "t",
}

var shinTable = map[rune]string{
	'\u05c2': "ś",
	'\u05c1': "š",
}

var maitresLectionesTable = map[rune]map[rune]string{
	'ה': {
		rune(QAMATS): "â",
	},
	'י': {
		rune(TSERE): "ê",
		rune(SEGOL): "ê",
		rune(HIRIK): "î",
	},
	rune(HOLAM): {
		'ו': "ô",
	},
	rune(DAGESH): {
		'ו': "û",
	},
}

var ignoredSet = map[rune]interface{}{
	// Accents
	'\u0591': nil,
	'\u0592': nil,
	'\u0593': nil,
	'\u0594': nil,
	'\u0595': nil,
	'\u0596': nil,
	'\u0597': nil,
	'\u0598': nil,
	'\u0599': nil,
	'\u059a': nil,
	'\u059b': nil,
	'\u059c': nil,
	'\u059d': nil,
	'\u059e': nil,
	'\u059f': nil,
	'\u05a1': nil,
	'\u05a2': nil,
	'\u05a3': nil,
	'\u05a4': nil,
	'\u05a5': nil,
	'\u05a6': nil,
	'\u05a7': nil,
	'\u05a8': nil,
	'\u05a9': nil,
	'\u05aa': nil,
	'\u05ab': nil,
	'\u05ac': nil,
	'\u05ad': nil,
	'\u05ae': nil,

	// Punctuation
	'\u05bd': nil,
	'\u05bf': nil,
	'\u05c0': nil,

	'\u05c3': nil,
	'\u05c4': nil,
	'\u05c5': nil,
	'\u05c6': nil,

	// Symbols
	'\ufffd': nil, // REPLACEMENT CHARACTER
}
