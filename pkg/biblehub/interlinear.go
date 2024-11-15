package biblehub

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
)

type Language string

const (
	Hebrew  Language = "Hebrew"
	Greek   Language = "Greek"
	Unknown Language = "Unknown"
)

type InterlinearChapter struct {
	Title    string
	Verses   []InterlinearVerse
	Language Language
}

type InterlinearVerse struct {
	Number int
	Words  []InterlinearWord
}
type InterlinearWord struct {
	StrongsNumber   string
	StrongsText     string
	Transliteration string
	English         string
	Original        string
	Syntax          string
}

func (lang Language) Select(ifHebrew, ifGreek string) string {
	switch lang {
	case Hebrew:
		return ifHebrew
	case Greek:
		return ifGreek
	default:
		return ""
	}
}

func (lang Language) Select2(ifHebrew, ifGreek goquery.Matcher) goquery.Matcher {
	switch lang {
	case Hebrew:
		return ifHebrew
	case Greek:
		return ifGreek
	default:
		return goquery.Single("~")
	}
}

func detectLanguage(doc *goquery.Document) Language {
	if doc.FindMatcher(goquery.Single(".tablefloatheb")).Length() > 0 {
		return Hebrew
	}
	if doc.FindMatcher(goquery.Single(".tablefloat")).Length() > 0 {
		return Greek
	}
	return Unknown
}

func text(sel *goquery.Selection, selectors ...string) string {
	for _, s := range selectors {
		sel = sel.Filter(s)
	}
	return strings.TrimSpace(sel.Text())
}

func (ex *_Extractor) GetInterlinearChapter(book string, chapter int) (InterlinearChapter, error) {
	var ch InterlinearChapter
	body, err := ex.Downloader.GetInterlinearChapter(book, chapter)
	if err != nil {
		return ch, err
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return ch, err
	}

	ch.Title = strings.TrimSpace(doc.Find("#topheading").Children().Remove().End().Text())
	var currentVerse *InterlinearVerse
	var currentWordNumber int

	lang := detectLanguage(doc)
	ch.Language = lang

	strongsSelector := lang.Select2(
		cascadia.MustCompile(".strongs"),
		cascadia.MustCompile(".pos"),
	)
	originalSelector := lang.Select(".hebrew", ".greek")
	refSelector := lang.Select(".refheb", ".refmain")

	doc.Find(lang.Select(".tablefloatheb", ".tablefloat")).
		ChildrenFiltered("tbody").
		ChildrenFiltered("tr").
		ChildrenFiltered("td").
		Each(func(i1 int, s1 *goquery.Selection) {
			spans := s1.Find("span")
			currentWordNumber++

			var word InterlinearWord

			strongs1 := spans.FilterMatcher(strongsSelector).Find("a").First()
			// strongs2 := strongs1.Next()
			word.StrongsNumber = text(strongs1)
			word.StrongsText = sanitizeEnglish(strings.TrimSpace(strongs1.AttrOr("title", "")))

			word.Transliteration = text(spans, ".translit")
			word.English = sanitizeEnglish(text(spans, ".eng"))
			word.Original = text(spans, originalSelector)
			word.Syntax = text(spans, ".strongsnt")
			spans.Filter(refSelector).
				First().
				Each(func(i int, ref *goquery.Selection) {
					if currentVerse != nil {
						ch.Verses = append(ch.Verses, *currentVerse)
					}
					currentVerse = new(InterlinearVerse)
					hebrewNumber, _ := strconv.Atoi(text(ref))
					currentVerse.Number = hebrewNumber
					currentWordNumber = 0
				})
			if currentVerse != nil {
				currentVerse.Words = append(currentVerse.Words, word)
			}
		})

	if currentVerse != nil {
		ch.Verses = append(ch.Verses, *currentVerse)
	}
	return ch, nil
}

func sanitizeEnglish(text string) string {
	text = strings.ReplaceAll(text, "\u00a0", " ")
	text = strings.ReplaceAll(text, "<BR>", "\n")
	return text
}
