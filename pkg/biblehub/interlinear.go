package biblehub

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"github.com/pkg/errors"
)

type Language string

const (
	Hebrew  Language = "Hebrew"
	Greek   Language = "Greek"
	Unknown Language = "Unknown"
)

type ChapterID struct {
	Book    string
	Chapter int
}

type InterlinearChapter struct {
	ChapterID
	Title    string
	Verses   []InterlinearVerse
	Language Language
	Next     *ChapterID
}

type InterlinearVerse struct {
	Number int
	Words  []InterlinearWord
}
type InterlinearWord struct {
	Original          string
	Transliteration   string
	StrongsNumber     string
	StrongsText       string
	SyntaxCode        string
	SyntaxDescription string
	English           string
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

func (ex *Scraper) GetInterlinearChapter(chapter ChapterID) (InterlinearChapter, error) {
	var ch InterlinearChapter
	ch.ChapterID = chapter
	body, err := ex.Downloader.GetInterlinearChapter(chapter)
	if err != nil {
		return ch, err
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return ch, err
	}

	ch.Title = strings.TrimSpace(doc.Find("#topheading").Children().Remove().End().Text())
	rightHref := doc.Find("div#right > a").AttrOr("href", "")
	if rightHref != "" {
		rightHref, _ = strings.CutPrefix(rightHref, "../")
		rightHref, _ = strings.CutSuffix(rightHref, ".htm")
		segments := strings.Split(rightHref, "/")
		if len(segments) != 2 {
			return ch, fmt.Errorf("could not detect next chapter: wrong number of sements in the url: %d", len(segments))
		}

		if segments[0] == "" {
			return ch, errors.New("could not detect next chapter: book name is empty")

		}
		num, err := strconv.Atoi(segments[1])
		if err != nil {
			return ch, errors.WithMessage(err, "could not detect next chapter: invalid chapter number")
		}
		ch.Next = &ChapterID{
			Book:    segments[0],
			Chapter: num,
		}
	}

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
			word.StrongsNumber = text(strongs1)
			word.StrongsText = sanitizeEnglish(strings.TrimSpace(strongs1.AttrOr("title", "")))

			word.Transliteration = text(spans, ".translit")
			word.English = sanitizeEnglish(text(spans, ".eng"))
			word.Original = text(spans, originalSelector)
			strongsnt := spans.Filter(lang.Select(".strongsnt", ".strongsnt2")).Last()
			word.SyntaxCode = strongsnt.Text()
			word.SyntaxDescription = strongsnt.Find("a").AttrOr("title", "")
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
