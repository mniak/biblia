package bible

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// ChapterRef represents a reference to a specific chapter in a book
type ChapterRef struct {
	Book    BookCode
	Chapter int
}

func (r ChapterRef) String() string {
	return fmt.Sprintf("%s%d", r.Book, r.Chapter)
}

// VerseRef represents a reference to a specific verse in a chapter
type VerseRef struct {
	Book    BookCode
	Chapter int
	Verse   int
}

func (r VerseRef) String() string {
	return fmt.Sprintf("%s%d:%d", r.Book, r.Chapter, r.Verse)
}

// ChapterRefOf returns a ChapterRef from the VerseRef
func (r VerseRef) ChapterRef() ChapterRef {
	return ChapterRef{
		Book:    r.Book,
		Chapter: r.Chapter,
	}
}

// ParseChapterRef parses a chapter reference like "Genesis1", "MAT5", "IJoao3"
// The aliases map provides additional book name mappings
func ParseChapterRef(s string, aliases map[string]BookCode) (ChapterRef, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return ChapterRef{}, fmt.Errorf("empty chapter reference")
	}

	book, chapter, err := splitBookAndNumber(s, aliases)
	if err != nil {
		return ChapterRef{}, fmt.Errorf("invalid chapter reference %q: %w", s, err)
	}

	return ChapterRef{Book: book, Chapter: chapter}, nil
}

// ParseVerseRef parses a verse reference like "Genesis1:1", "MAT5:3", "iJoao4:8"
// The aliases map provides additional book name mappings
func ParseVerseRef(s string, aliases map[string]BookCode) (VerseRef, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return VerseRef{}, fmt.Errorf("empty verse reference")
	}

	// Find the colon separator for verse
	colonIdx := strings.LastIndex(s, ":")
	if colonIdx == -1 {
		return VerseRef{}, fmt.Errorf("invalid verse reference %q: missing colon", s)
	}

	chapterPart := s[:colonIdx]
	versePart := s[colonIdx+1:]

	verse, err := strconv.Atoi(versePart)
	if err != nil {
		return VerseRef{}, fmt.Errorf("invalid verse number %q: %w", versePart, err)
	}

	if verse < 1 {
		return VerseRef{}, fmt.Errorf("verse must be positive: %d", verse)
	}

	book, chapter, err := splitBookAndNumber(chapterPart, aliases)
	if err != nil {
		return VerseRef{}, fmt.Errorf("invalid verse reference %q: %w", s, err)
	}

	return VerseRef{Book: book, Chapter: chapter, Verse: verse}, nil
}

// splitBookAndNumber splits "Genesis1" into (Genesis, 1)
// Handles tricky cases like "1Samuel1" -> (1SA, 1), "IJoao3" -> (1JN, 3)
func splitBookAndNumber(s string, aliases map[string]BookCode) (BookCode, int, error) {
	// Find where the trailing number starts
	numStart := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			numStart = i
		} else {
			break
		}
	}

	if numStart == len(s) {
		return "", 0, fmt.Errorf("missing chapter number")
	}

	if numStart == 0 {
		return "", 0, fmt.Errorf("missing book name")
	}

	bookPart := s[:numStart]
	numPart := s[numStart:]

	num, err := strconv.Atoi(numPart)
	if err != nil {
		return "", 0, fmt.Errorf("invalid number %q: %w", numPart, err)
	}

	if num < 1 {
		return "", 0, fmt.Errorf("chapter must be positive: %d", num)
	}

	// Try to resolve the book
	book, err := ParseBookCode(bookPart, aliases)
	if err != nil {
		return "", 0, err
	}

	return book, num, nil
}

// ParseBookCode parses a string into a BookCode using:
// 1. Direct UBS code match (case-insensitive)
// 2. Aliases map (case-insensitive)
func ParseBookCode(s string, aliases map[string]BookCode) (BookCode, error) {
	s = strings.TrimSpace(s)

	// Try as UBS code
	code := BookCode(strings.ToUpper(s))
	if code.Index() != 0 {
		return code, nil
	}

	// Try aliases (case-insensitive)
	if aliases != nil {
		lower := strings.ToLower(s)
		for alias, bookCode := range aliases {
			if strings.ToLower(alias) == lower {
				return bookCode, nil
			}
		}
	}

	return "", fmt.Errorf("unknown book: %s", s)
}

// ParseBookIndex parses a string into a BookIndex using:
// 1. Direct UBS code match (case-insensitive)
// 2. Aliases map (case-insensitive)
func ParseBookIndex(s string, aliases map[string]BookCode) (BookIndex, error) {
	code, err := ParseBookCode(s, aliases)
	if err != nil {
		return 0, err
	}
	return code.Index(), nil
}
