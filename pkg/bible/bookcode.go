package bible

// BookCode represents a standardized book code following UBS (United Bible Societies) conventions
type BookCode string

// Old Testament book codes (UBS standard)
const (
	// Torah / Pentateuch
	Genesis     BookCode = "GEN"
	Exodus      BookCode = "EXO"
	Leviticus   BookCode = "LEV"
	Numbers     BookCode = "NUM"
	Deuteronomy BookCode = "DEU"

	// Historical Books
	Joshua       BookCode = "JOS"
	Judges       BookCode = "JDG"
	Ruth         BookCode = "RUT"
	ISamuel      BookCode = "1SA"
	IISamuel     BookCode = "2SA"
	IKings       BookCode = "1KI"
	IIKings      BookCode = "2KI"
	IChronicles  BookCode = "1CH"
	IIChronicles BookCode = "2CH"
	Ezra         BookCode = "EZR"
	Nehemiah     BookCode = "NEH"
	Esther       BookCode = "EST"

	// Wisdom / Poetry Books
	Job          BookCode = "JOB"
	Psalms       BookCode = "PSA"
	Proverbs     BookCode = "PRO"
	Ecclesiastes BookCode = "ECC"
	SongOfSongs  BookCode = "SNG"

	// Major Prophets
	Isaiah       BookCode = "ISA"
	Jeremiah     BookCode = "JER"
	Lamentations BookCode = "LAM"
	Ezekiel      BookCode = "EZK"
	Daniel       BookCode = "DAN"

	// Minor Prophets
	Hosea     BookCode = "HOS"
	Joel      BookCode = "JOL"
	Amos      BookCode = "AMO"
	Obadiah   BookCode = "OBA"
	Jonah     BookCode = "JON"
	Micah     BookCode = "MIC"
	Nahum     BookCode = "NAM"
	Habakkuk  BookCode = "HAB"
	Zephaniah BookCode = "ZEP"
	Haggai    BookCode = "HAG"
	Zechariah BookCode = "ZEC"
	Malachi   BookCode = "MAL"
)

// New Testament book codes (UBS standard)
const (
	// Gospels
	Matthew BookCode = "MAT"
	Mark    BookCode = "MRK"
	Luke    BookCode = "LUK"
	John    BookCode = "JHN"

	// History
	Acts BookCode = "ACT"

	// Pauline Epistles
	Romans          BookCode = "ROM"
	ICorinthians    BookCode = "1CO"
	IICorinthians   BookCode = "2CO"
	Galatians       BookCode = "GAL"
	Ephesians       BookCode = "EPH"
	Philippians     BookCode = "PHP"
	Colossians      BookCode = "COL"
	IThessalonians  BookCode = "1TH"
	IIThessalonians BookCode = "2TH"
	ITimothy        BookCode = "1TI"
	IITimothy       BookCode = "2TI"
	Titus           BookCode = "TIT"
	Philemon        BookCode = "PHM"

	// General Epistles
	Hebrews BookCode = "HEB"
	James   BookCode = "JAS"
	IPeter  BookCode = "1PE"
	IIPeter BookCode = "2PE"
	IJohn   BookCode = "1JN"
	IIJohn  BookCode = "2JN"
	IIIJohn BookCode = "3JN"
	Jude    BookCode = "JUD"

	// Apocalyptic
	Revelation BookCode = "REV"
)

// OldTestamentBooks returns all Old Testament book codes in canonical order
func OldTestamentBooks() []BookCode {
	return []BookCode{
		Genesis, Exodus, Leviticus, Numbers, Deuteronomy,
		Joshua, Judges, Ruth, ISamuel, IISamuel, IKings, IIKings, IChronicles, IIChronicles, Ezra, Nehemiah, Esther,
		Job, Psalms, Proverbs, Ecclesiastes, SongOfSongs,
		Isaiah, Jeremiah, Lamentations, Ezekiel, Daniel,
		Hosea, Joel, Amos, Obadiah, Jonah, Micah, Nahum, Habakkuk, Zephaniah, Haggai, Zechariah, Malachi,
	}
}

// NewTestamentBooks returns all New Testament book codes in canonical order
func NewTestamentBooks() []BookCode {
	return []BookCode{
		Matthew, Mark, Luke, John,
		Acts,
		Romans, ICorinthians, IICorinthians, Galatians, Ephesians, Philippians, Colossians, IThessalonians, IIThessalonians, ITimothy, IITimothy, Titus, Philemon,
		Hebrews, James, IPeter, IIPeter, IJohn, IIJohn, IIIJohn, Jude,
		Revelation,
	}
}

// AllBooks returns all book codes in canonical order
func AllBooks() []BookCode {
	return append(OldTestamentBooks(), NewTestamentBooks()...)
}

// String returns the string representation of the book code
func (c BookCode) String() string {
	return string(c)
}

// Name returns the full English name of the book
func (c BookCode) Name() string {
	names := map[BookCode]string{
		// Old Testament
		Genesis:      "Genesis",
		Exodus:       "Exodus",
		Leviticus:    "Leviticus",
		Numbers:      "Numbers",
		Deuteronomy:  "Deuteronomy",
		Joshua:       "Joshua",
		Judges:       "Judges",
		Ruth:         "Ruth",
		ISamuel:      "1 Samuel",
		IISamuel:     "2 Samuel",
		IKings:       "1 Kings",
		IIKings:      "2 Kings",
		IChronicles:  "1 Chronicles",
		IIChronicles: "2 Chronicles",
		Ezra:         "Ezra",
		Nehemiah:     "Nehemiah",
		Esther:       "Esther",
		Job:          "Job",
		Psalms:       "Psalms",
		Proverbs:     "Proverbs",
		Ecclesiastes: "Ecclesiastes",
		SongOfSongs:  "Song of Songs",
		Isaiah:       "Isaiah",
		Jeremiah:     "Jeremiah",
		Lamentations: "Lamentations",
		Ezekiel:      "Ezekiel",
		Daniel:       "Daniel",
		Hosea:        "Hosea",
		Joel:         "Joel",
		Amos:         "Amos",
		Obadiah:      "Obadiah",
		Jonah:        "Jonah",
		Micah:        "Micah",
		Nahum:        "Nahum",
		Habakkuk:     "Habakkuk",
		Zephaniah:    "Zephaniah",
		Haggai:       "Haggai",
		Zechariah:    "Zechariah",
		Malachi:      "Malachi",

		// New Testament
		Matthew:         "Matthew",
		Mark:            "Mark",
		Luke:            "Luke",
		John:            "John",
		Acts:            "Acts",
		Romans:          "Romans",
		ICorinthians:    "1 Corinthians",
		IICorinthians:   "2 Corinthians",
		Galatians:       "Galatians",
		Ephesians:       "Ephesians",
		Philippians:     "Philippians",
		Colossians:      "Colossians",
		IThessalonians:  "1 Thessalonians",
		IIThessalonians: "2 Thessalonians",
		ITimothy:        "1 Timothy",
		IITimothy:       "2 Timothy",
		Titus:           "Titus",
		Philemon:        "Philemon",
		Hebrews:         "Hebrews",
		James:           "James",
		IPeter:          "1 Peter",
		IIPeter:         "2 Peter",
		IJohn:           "1 John",
		IIJohn:          "2 John",
		IIIJohn:         "3 John",
		Jude:            "Jude",
		Revelation:      "Revelation",
	}

	if name, ok := names[c]; ok {
		return name
	}
	return string(c)
}
