package main

import (
	"fmt"
	"strings"
	"unicode"
)

func normalizeIdentifier(text string) string {
	if text == "" {
		return text
	}
	text = strings.Trim(text, " ")
	splitted := strings.SplitN(text, " ", 2)
	if len(splitted) == 2 {
		firstRune := rune(splitted[0][0])
		if unicode.IsDigit(firstRune) {
			return normalizeIdentifier(fmt.Sprintf("%s_%s", splitted[1], splitted[0]))
		}
	}
	var countDigits int
	for _, ch := range text {
		if unicode.IsDigit(ch) || ch == ' ' {
			countDigits++
		} else {
			break
		}
	}
	text = text[countDigits:] + text[:countDigits]
	text = strings.ReplaceAll(text, " ", "_")
	return text
}
