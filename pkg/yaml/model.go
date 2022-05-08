package yaml

type verseModel struct {
	BookName string      `yaml:"book_name"`
	Chapter  int         `yaml:"chapter"`
	Verse    int         `yaml:"verse"`
	Words    []wordModel `yaml:"words"`
}

type wordModel struct {
	Text            string `yaml:"text"`
	Transliteration string `yaml:"tlit"`
}
