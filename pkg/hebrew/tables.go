package hebrew

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
