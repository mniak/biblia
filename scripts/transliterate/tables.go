package main

const (
	INVALID   = "�"
	DAGESH    = '\u05bc'
	QAMATS    = '\u05b8'
	PATAH     = '\u05b7'
	SEGOL     = '\u05b6'
	TSERE     = '\u05b5'
	HIRIK     = '\u05b4'
	HOLAM     = '\u05b9'
	HOLAM_VAV = '\u05ba'
	QUBUTS    = '\u05bb'
	SHEVA     = '\u05b0'
)

var basicTable = map[rune]string{
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
		QAMATS: "â",
	},
	'י': {
		TSERE: "ê",
		SEGOL: "ê",
		HIRIK: "î",
	},
	'ו': {
		// HOLAM:     "ô",
		HOLAM_VAV: "ô",
	},
}
