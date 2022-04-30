package wlc

import "encoding/xml"

type TanachChapter struct {
	Text   string `xml:",chardata"`
	Number int    `xml:"n,attr"`
	Verses []struct {
		Text   string `xml:",chardata"`
		Number int    `xml:"n,attr"`
		Words  []struct {
			Text string `xml:",chardata"`
			X    string `xml:"x"`
		} `xml:"w"`
		Pe     string   `xml:"pe"`
		Samekh string   `xml:"samekh"`
		K      []string `xml:"k"`
		Q      []string `xml:"q"`
	} `xml:"v"`
	VerseCount int `xml:"vs"`
}
type TanachBook struct {
	Text  string `xml:",chardata"`
	Names struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name"`
		Abbrev     string `xml:"abbrev"`
		Number     string `xml:"number"`
		Filename   string `xml:"filename"`
		Hebrewname string `xml:"hebrewname"`
	} `xml:"names"`
	Chapters []TanachChapter `xml:"c"`
	Vs       string          `xml:"vs"`
	Cs       string          `xml:"cs"`
}

type TanachRoot struct {
	Text string     `xml:",chardata"`
	Book TanachBook `xml:"book"`
}

type TanachFile struct {
	XMLName                   xml.Name `xml:"Tanach"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	TeiHeader                 struct {
		Text     string `xml:",chardata"`
		FileDesc struct {
			Text      string `xml:",chardata"`
			TitleStmt struct {
				Text  string `xml:",chardata"`
				Title []struct {
					Text  string `xml:",chardata"`
					Level string `xml:"level,attr"`
					Type  string `xml:"type,attr"`
				} `xml:"title"`
				Editor   []string `xml:"editor"`
				RespStmt struct {
					Text string `xml:",chardata"`
					Resp string `xml:"resp"`
					Name []struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"name"`
				} `xml:"respStmt"`
			} `xml:"titleStmt"`
			EditionStmt struct {
				Text    string `xml:",chardata"`
				Edition struct {
					Text          string `xml:",chardata"`
					Version       string `xml:"version"`
					Date          string `xml:"date"`
					Build         string `xml:"build"`
					BuildDateTime string `xml:"buildDateTime"`
				} `xml:"edition"`
				RespStmt struct {
					Text string   `xml:",chardata"`
					Resp []string `xml:"resp"`
				} `xml:"respStmt"`
			} `xml:"editionStmt"`
			Extent          string `xml:"extent"`
			PublicationStmt struct {
				Text        string `xml:",chardata"`
				Distributor struct {
					Text string `xml:",chardata"`
					Name []struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"name"`
				} `xml:"distributor"`
				Availability struct {
					Text   string `xml:",chardata"`
					Status string `xml:"status,attr"`
				} `xml:"availability"`
			} `xml:"publicationStmt"`
			NotesStmt struct {
				Text       string `xml:",chardata"`
				Note       string `xml:"note"`
				Correction []struct {
					Text        string `xml:",chardata"`
					Citation    string `xml:"citation"`
					Description string `xml:"description"`
					Author      string `xml:"author"`
					Filedate    string `xml:"filedate"`
					Date        string `xml:"date"`
					N           string `xml:"n"`
				} `xml:"correction"`
			} `xml:"notesStmt"`
			SourceDesc struct {
				Text     string `xml:",chardata"`
				BiblItem []struct {
					Text    string   `xml:",chardata"`
					Title   []string `xml:"title"`
					Editor  []string `xml:"editor"`
					Edition string   `xml:"edition"`
					Imprint struct {
						Text      string   `xml:",chardata"`
						Publisher string   `xml:"publisher"`
						PubPlace  []string `xml:"pubPlace"`
						Date      string   `xml:"date"`
					} `xml:"imprint"`
					Idno struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"idno"`
				} `xml:"biblItem"`
			} `xml:"sourceDesc"`
		} `xml:"fileDesc"`
		EncodingDesc string `xml:"encodingDesc"`
		ProfileDesc  struct {
			Text      string `xml:",chardata"`
			Creation  string `xml:"creation"`
			Date      string `xml:"date"`
			LangUsage struct {
				Text     string `xml:",chardata"`
				Language struct {
					Text  string `xml:",chardata"`
					Ident string `xml:"ident,attr"`
				} `xml:"language"`
			} `xml:"langUsage"`
		} `xml:"profileDesc"`
	} `xml:"teiHeader"`
	Tanach TanachRoot `xml:"tanach"`
	Notes  struct {
		Text string `xml:",chardata"`
		Note []struct {
			Text   string `xml:",chardata"`
			Code   string `xml:"code"`
			Gccode string `xml:"gccode"`
			Note   string `xml:"note"`
		} `xml:"note"`
	} `xml:"notes"`
}
