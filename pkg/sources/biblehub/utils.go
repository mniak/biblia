package biblehub

import "strings"

func normalize(text string) string {
	text = strings.ReplaceAll(text, "\u00a0", " ")
	text = strings.Trim(text, " ")
	return text
}
