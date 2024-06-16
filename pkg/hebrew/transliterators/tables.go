package transliterators

const (
	InvalidChar = '�'
)

func basicConvert(r rune) string {
	switch r {
	case rune(Alef):
		return "ʾ"
	case rune(Bet):
		return "ḇ"
	case rune(Gimel):
		return "ḡ"
	case rune(Dalet):
		return "ḏ"
	case rune(He):
		return "h"
	case rune(Vav):
		return "w"
	case rune(Zayin):
		return "z"
	case rune(Het):
		return "ḥ"
	case rune(Tet):
		return "ṭ"
	case rune(Yod):
		return "y"
	case rune(Kaf), rune(FinalKaf):
		return "ḵ"
	case rune(Lamed):
		return "l"
	case rune(Mem), rune(FinalMem):
		return "m"
	case rune(Nun), rune(FinalNun):
		return "n"
	case rune(Samekh):
		return "s"
	case rune(Ayin):
		return "ʿ"
	case rune(Pe), rune(FinalPe):
		return "p̄"
	case rune(Tsadi), rune(FinalTsadi):
		return "ṣ"
	case rune(Qof):
		return "q"
	case rune(Resh):
		return "r"
	case rune(Shin):
		return "š"
	case rune(Tav):
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
	return string(InvalidChar)
}

var dageshTable = map[Letter]string{
	Bet:   "b",
	Gimel: "g",
	Dalet: "d",
	Kaf:   "k",
	Pe:    "p",
	Tav:   "t",
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
