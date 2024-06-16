package transliterators

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
