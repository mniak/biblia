package loaders

import (
	"encoding/xml"
	"os"

	"github.com/mniak/biblia/pkg/bible"
)

func LoadWestminsterLeningradCodex() (bible.Testament, error) {
	bookBytes, err := os.ReadFile("sources/UnicodeXML_Westminster_Leningrad_Codex/Tanach/Genesis.xml")
	if err != nil {
		return bible.Testament{}, err
	}

	err = xml.Unmarshal(data, &tanach)

	return bible.Testament{
		// Books: ,
	}, nil
}
