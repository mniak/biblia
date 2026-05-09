package bible

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseChapterRef(t *testing.T) {
	aliases := map[string]BookCode{
		"Genesis": Genesis,
		"Gn":      Genesis,
		"IJoao":   IJohn,
		"ijoao":   IJohn,
		"1Joao":   IJohn,
	}

	tests := []struct {
		name    string
		input   string
		want    ChapterRef
		wantErr bool
	}{
		{
			name:  "by UBS code",
			input: "GEN1",
			want:  ChapterRef{Book: Genesis, Chapter: 1},
		},
		{
			name:  "by UBS code lowercase",
			input: "gen1",
			want:  ChapterRef{Book: Genesis, Chapter: 1},
		},
		{
			name:  "by full name",
			input: "Genesis1",
			want:  ChapterRef{Book: Genesis, Chapter: 1},
		},
		{
			name:  "by alias",
			input: "Gn1",
			want:  ChapterRef{Book: Genesis, Chapter: 1},
		},
		{
			name:  "multi-digit chapter",
			input: "GEN50",
			want:  ChapterRef{Book: Genesis, Chapter: 50},
		},
		{
			name:  "book with roman numeral prefix",
			input: "IJoao3",
			want:  ChapterRef{Book: IJohn, Chapter: 3},
		},
		{
			name:  "book with roman numeral lowercase",
			input: "ijoao4",
			want:  ChapterRef{Book: IJohn, Chapter: 4},
		},
		{
			name:  "new testament book",
			input: "MAT5",
			want:  ChapterRef{Book: Matthew, Chapter: 5},
		},
		{
			name:    "missing chapter number",
			input:   "Genesis",
			wantErr: true,
		},
		{
			name:    "missing book name",
			input:   "123",
			wantErr: true,
		},
		{
			name:    "unknown book",
			input:   "Unknown1",
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseChapterRef(tt.input, aliases)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseVerseRef(t *testing.T) {
	aliases := map[string]BookCode{
		"Genesis": Genesis,
		"Gn":      Genesis,
		"IJoao":   IJohn,
		"ijoao":   IJohn,
	}

	tests := []struct {
		name    string
		input   string
		want    VerseRef
		wantErr bool
	}{
		{
			name:  "by UBS code",
			input: "GEN1:1",
			want:  VerseRef{Book: Genesis, Chapter: 1, Verse: 1},
		},
		{
			name:  "by UBS code lowercase",
			input: "gen1:1",
			want:  VerseRef{Book: Genesis, Chapter: 1, Verse: 1},
		},
		{
			name:  "by full name",
			input: "Genesis1:1",
			want:  VerseRef{Book: Genesis, Chapter: 1, Verse: 1},
		},
		{
			name:  "by alias",
			input: "Gn1:1",
			want:  VerseRef{Book: Genesis, Chapter: 1, Verse: 1},
		},
		{
			name:  "multi-digit chapter and verse",
			input: "GEN50:26",
			want:  VerseRef{Book: Genesis, Chapter: 50, Verse: 26},
		},
		{
			name:  "book with roman numeral prefix",
			input: "IJoao4:8",
			want:  VerseRef{Book: IJohn, Chapter: 4, Verse: 8},
		},
		{
			name:  "book with roman numeral lowercase",
			input: "ijoao4:8",
			want:  VerseRef{Book: IJohn, Chapter: 4, Verse: 8},
		},
		{
			name:  "new testament book",
			input: "MAT5:3",
			want:  VerseRef{Book: Matthew, Chapter: 5, Verse: 3},
		},
		{
			name:    "missing colon",
			input:   "Genesis1",
			wantErr: true,
		},
		{
			name:    "missing verse number",
			input:   "Genesis1:",
			wantErr: true,
		},
		{
			name:    "missing chapter number",
			input:   "Genesis:1",
			wantErr: true,
		},
		{
			name:    "unknown book",
			input:   "Unknown1:1",
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseVerseRef(tt.input, aliases)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestChapterRef_String(t *testing.T) {
	ref := ChapterRef{Book: Genesis, Chapter: 1}
	assert.Equal(t, "GEN1", ref.String())
}

func TestVerseRef_String(t *testing.T) {
	ref := VerseRef{Book: Genesis, Chapter: 1, Verse: 1}
	assert.Equal(t, "GEN1:1", ref.String())
}

func TestVerseRef_ChapterRef(t *testing.T) {
	verse := VerseRef{Book: Genesis, Chapter: 1, Verse: 5}
	chapter := verse.ChapterRef()
	assert.Equal(t, ChapterRef{Book: Genesis, Chapter: 1}, chapter)
}
