package biblehub

import (
	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/text"
)

const localInterlinearOTDirectory = ""

func NewLocalInterlinearOTLoader() bible.TestamentLoader {
	return text.NewTomlLoader(localInterlinearOTDirectory)
}

func NewLocalInterlinearOTExporter() bible.Exporter {
	return text.NewTomlExporter(localInterlinearOTDirectory)
}
