package greek

import (
	"bytes"
)

type simpleTransliterator struct{}

func SimpleTransliterator() *simpleTransliterator {
	return &simpleTransliterator{}
}

func (t simpleTransliterator) TransliterateWord(word string) string {
	var b bytes.Buffer
	for _, r := range word {
		if str, ok := simpleTable[r]; ok {
			b.WriteString(str)
			b.WriteRune(r) // If character is not in map, just keep it
		}
	}
	return b.String()
}

var simpleTable = map[rune]string{
	'Α': "A", 'α': "a",
	'Β': "B", 'β': "b",
	'Γ': "G", 'γ': "g",
	'Δ': "D", 'δ': "d",
	'Ε': "E", 'ε': "e",
	'Ζ': "Z", 'ζ': "z",
	'Η': "H", 'η': "h",
	'Θ': "Th", 'θ': "th",
	'Ι': "I", 'ι': "i",
	'Κ': "K", 'κ': "k",
	'Λ': "L", 'λ': "l",
	'Μ': "M", 'μ': "m",
	'Ν': "N", 'ν': "n",
	'Ξ': "X", 'ξ': "x",
	'Ο': "O", 'ο': "o",
	'Π': "P", 'π': "p",
	'Ρ': "R", 'ρ': "r",
	'Σ': "S", 'σ': "s", 'ς': "s",
	'Τ': "T", 'τ': "t",
	'Υ': "Y", 'υ': "y",
	'Φ': "F", 'φ': "f",
	'Χ': "Ch", 'χ': "ch",
	'Ψ': "Ps", 'ψ': "ps",
	'Ω': "O", 'ω': "o",
}
