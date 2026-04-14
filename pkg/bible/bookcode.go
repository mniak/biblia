package bible

// BookCode represents a standardized book code following UBS (United Bible Societies) conventions
type BookCode string

// Old Testament book codes (UBS standard)
const (
	// Torah / Pentateuch
	GEN BookCode = "GEN" // Genesis
	EXO BookCode = "EXO" // Exodus
	LEV BookCode = "LEV" // Leviticus
	NUM BookCode = "NUM" // Numbers
	DEU BookCode = "DEU" // Deuteronomy

	// Historical Books
	JOS BookCode = "JOS" // Joshua
	JDG BookCode = "JDG" // Judges
	RUT BookCode = "RUT" // Ruth
	SA1 BookCode = "1SA" // 1 Samuel
	SA2 BookCode = "2SA" // 2 Samuel
	KI1 BookCode = "1KI" // 1 Kings
	KI2 BookCode = "2KI" // 2 Kings
	CH1 BookCode = "1CH" // 1 Chronicles
	CH2 BookCode = "2CH" // 2 Chronicles
	EZR BookCode = "EZR" // Ezra
	NEH BookCode = "NEH" // Nehemiah
	EST BookCode = "EST" // Esther

	// Wisdom / Poetry Books
	JOB BookCode = "JOB" // Job
	PSA BookCode = "PSA" // Psalms
	PRO BookCode = "PRO" // Proverbs
	ECC BookCode = "ECC" // Ecclesiastes
	SNG BookCode = "SNG" // Song of Songs

	// Major Prophets
	ISA BookCode = "ISA" // Isaiah
	JER BookCode = "JER" // Jeremiah
	LAM BookCode = "LAM" // Lamentations
	EZK BookCode = "EZK" // Ezekiel
	DAN BookCode = "DAN" // Daniel

	// Minor Prophets
	HOS BookCode = "HOS" // Hosea
	JOL BookCode = "JOL" // Joel
	AMO BookCode = "AMO" // Amos
	OBA BookCode = "OBA" // Obadiah
	JON BookCode = "JON" // Jonah
	MIC BookCode = "MIC" // Micah
	NAM BookCode = "NAM" // Nahum
	HAB BookCode = "HAB" // Habakkuk
	ZEP BookCode = "ZEP" // Zephaniah
	HAG BookCode = "HAG" // Haggai
	ZEC BookCode = "ZEC" // Zechariah
	MAL BookCode = "MAL" // Malachi
)

// New Testament book codes (UBS standard)
const (
	// Gospels
	MAT BookCode = "MAT" // Matthew
	MRK BookCode = "MRK" // Mark
	LUK BookCode = "LUK" // Luke
	JHN BookCode = "JHN" // John

	// History
	ACT BookCode = "ACT" // Acts

	// Pauline Epistles
	ROM BookCode = "ROM" // Romans
	CO1 BookCode = "1CO" // 1 Corinthians
	CO2 BookCode = "2CO" // 2 Corinthians
	GAL BookCode = "GAL" // Galatians
	EPH BookCode = "EPH" // Ephesians
	PHP BookCode = "PHP" // Philippians
	COL BookCode = "COL" // Colossians
	TH1 BookCode = "1TH" // 1 Thessalonians
	TH2 BookCode = "2TH" // 2 Thessalonians
	TI1 BookCode = "1TI" // 1 Timothy
	TI2 BookCode = "2TI" // 2 Timothy
	TIT BookCode = "TIT" // Titus
	PHM BookCode = "PHM" // Philemon

	// General Epistles
	HEB BookCode = "HEB" // Hebrews
	JAS BookCode = "JAS" // James
	PE1 BookCode = "1PE" // 1 Peter
	PE2 BookCode = "2PE" // 2 Peter
	JN1 BookCode = "1JN" // 1 John
	JN2 BookCode = "2JN" // 2 John
	JN3 BookCode = "3JN" // 3 John
	JUD BookCode = "JUD" // Jude

	// Apocalyptic
	REV BookCode = "REV" // Revelation
)

// OldTestamentBooks returns all Old Testament book codes in canonical order
func OldTestamentBooks() []BookCode {
	return []BookCode{
		GEN, EXO, LEV, NUM, DEU,
		JOS, JDG, RUT, SA1, SA2, KI1, KI2, CH1, CH2, EZR, NEH, EST,
		JOB, PSA, PRO, ECC, SNG,
		ISA, JER, LAM, EZK, DAN,
		HOS, JOL, AMO, OBA, JON, MIC, NAM, HAB, ZEP, HAG, ZEC, MAL,
	}
}

// NewTestamentBooks returns all New Testament book codes in canonical order
func NewTestamentBooks() []BookCode {
	return []BookCode{
		MAT, MRK, LUK, JHN,
		ACT,
		ROM, CO1, CO2, GAL, EPH, PHP, COL, TH1, TH2, TI1, TI2, TIT, PHM,
		HEB, JAS, PE1, PE2, JN1, JN2, JN3, JUD,
		REV,
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
		GEN: "Genesis",
		EXO: "Exodus",
		LEV: "Leviticus",
		NUM: "Numbers",
		DEU: "Deuteronomy",
		JOS: "Joshua",
		JDG: "Judges",
		RUT: "Ruth",
		SA1: "1 Samuel",
		SA2: "2 Samuel",
		KI1: "1 Kings",
		KI2: "2 Kings",
		CH1: "1 Chronicles",
		CH2: "2 Chronicles",
		EZR: "Ezra",
		NEH: "Nehemiah",
		EST: "Esther",
		JOB: "Job",
		PSA: "Psalms",
		PRO: "Proverbs",
		ECC: "Ecclesiastes",
		SNG: "Song of Songs",
		ISA: "Isaiah",
		JER: "Jeremiah",
		LAM: "Lamentations",
		EZK: "Ezekiel",
		DAN: "Daniel",
		HOS: "Hosea",
		JOL: "Joel",
		AMO: "Amos",
		OBA: "Obadiah",
		JON: "Jonah",
		MIC: "Micah",
		NAM: "Nahum",
		HAB: "Habakkuk",
		ZEP: "Zephaniah",
		HAG: "Haggai",
		ZEC: "Zechariah",
		MAL: "Malachi",

		// New Testament
		MAT: "Matthew",
		MRK: "Mark",
		LUK: "Luke",
		JHN: "John",
		ACT: "Acts",
		ROM: "Romans",
		CO1: "1 Corinthians",
		CO2: "2 Corinthians",
		GAL: "Galatians",
		EPH: "Ephesians",
		PHP: "Philippians",
		COL: "Colossians",
		TH1: "1 Thessalonians",
		TH2: "2 Thessalonians",
		TI1: "1 Timothy",
		TI2: "2 Timothy",
		TIT: "Titus",
		PHM: "Philemon",
		HEB: "Hebrews",
		JAS: "James",
		PE1: "1 Peter",
		PE2: "2 Peter",
		JN1: "1 John",
		JN2: "2 John",
		JN3: "3 John",
		JUD: "Jude",
		REV: "Revelation",
	}

	if name, ok := names[c]; ok {
		return name
	}
	return string(c)
}
