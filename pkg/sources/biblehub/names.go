package biblehub

import "github.com/mniak/biblia/pkg/bible"

const UnknownBookName = "unknown"

func mapOTBookName(name bible.OldTestamentBook) string {
	switch name {
	case bible.Genesis:
		return "genesis"
	case bible.Exodus:
		return "exodus"
	case bible.Leviticus:
		return "leviticus"
	case bible.Numbers:
		return "numbers"
	case bible.Deuteronomy:
		return "deuteronomy"
	case bible.Joshua:
		return "joshua"
	case bible.Judges:
		return "judges"
	case bible.Ruth:
		return "ruth"
	case bible.Samuel_1:
		return "1_samuel"
	case bible.Samuel_2:
		return "2_samuel"
	case bible.Kings_1:
		return "1_kings"
	case bible.Kings_2:
		return "2_kings"
	case bible.Chronicles_1:
		return "1_chronicles"
	case bible.Chronicles_2:
		return "2_chronicles"
	case bible.Ezra:
		return "ezra"
	case bible.Nehemiah:
		return "nehemiah"
	case bible.Esther:
		return "esther"
	case bible.Job:
		return "job"
	case bible.Psalms:
		return "psalms"
	case bible.Proverbs:
		return "proverbs"
	case bible.Ecclesiastes:
		return "ecclesiastes"
	case bible.Song_of_Songs:
		return "songs"
	case bible.Isaiah:
		return "isaiah"
	case bible.Jeremiah:
		return "jeremiah"
	case bible.Lamentations:
		return "lamentations"
	case bible.Ezekiel:
		return "ezekiel"
	case bible.Daniel:
		return "daniel"
	case bible.Hosea:
		return "hosea"
	case bible.Joel:
		return "joel"
	case bible.Amos:
		return "amos"
	case bible.Obadiah:
		return "obadiah"
	case bible.Jonah:
		return "jonah"
	case bible.Micah:
		return "micah"
	case bible.Nahum:
		return "nahum"
	case bible.Habakkuk:
		return "habakkuk"
	case bible.Zephaniah:
		return "zephaniah"
	case bible.Haggai:
		return "haggai"
	case bible.Zechariah:
		return "zechariah"
	case bible.Malachi:
		return "malachi"
	}
	return UnknownBookName
}

// func mapNTBookName(name bible.NewTestamentBook) string {
// 	switch name {
// 	case bible.Matthew:
// 		return "matthew"
// 	case bible.Mark:
// 		return "mark"
// 	case bible.Luke:
// 		return "luke"
// 	case bible.John:
// 		return "john"
// 	case bible.Acts:
// 		return "acts"
// 	case bible.Romans:
// 		return "romans"
// 	case bible.Corinthians_1:
// 		return "1_corinthians"
// 	case bible.Corinthians_2:
// 		return "2_corinthians"
// 	case bible.Galatians:
// 		return "galatians"
// 	case bible.Ephesians:
// 		return "ephesians"
// 	case bible.Philippians:
// 		return "philippians"
// 	case bible.Colossians:
// 		return "colossians"
// 	case bible.Thessalonians_1:
// 		return "1_thessalonians"
// 	case bible.Thessalonians_2:
// 		return "2_thessalonians"
// 	case bible.Timothy_1:
// 		return "1_timothy"
// 	case bible.Timothy_2:
// 		return "2_timothy"
// 	case bible.Titus:
// 		return "titus"
// 	case bible.Philemon:
// 		return "philemon"
// 	case bible.Hebrews:
// 		return "hebrews"
// 	case bible.James:
// 		return "james"
// 	case bible.Peter_1:
// 		return "1_peter"
// 	case bible.Peter_2:
// 		return "2_peter"
// 	case bible.John_1:
// 		return "1_john"
// 	case bible.John_2:
// 		return "2_john"
// 	case bible.John_3:
// 		return "3_john"
// 	case bible.Jude:
// 		return "jude"
// 	case bible.Revelation:
// 		return "revelation"
// 	}
// 	return UnknownBookName
// }
