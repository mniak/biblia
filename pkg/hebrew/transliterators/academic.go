package transliterators

import (
	"strings"

	"github.com/mniak/biblia/pkg/runeutils"
)

type academic struct{}

func Academic() *academic {
	return new(academic)
}

func (academic) TransliterateWord(word string) string {
	walker := runeutils.NewReverseRuneWalker(word)
	walker.Filter(func(r rune) bool {
		_, ignored := ignoredSet[r]
		return !ignored
	})

	resultChars := make([]string, 0)
	for walker.Walk() {
		resultChars = append(resultChars, getLastChar(walker))
	}

	var sb strings.Builder
	for i := len(resultChars) - 1; i >= 0; i-- {
		sb.WriteString(resultChars[i])
	}
	return sb.String()
}
