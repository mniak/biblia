package flagutils

import (
	"fmt"
	"os"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/text"
	"github.com/mniak/biblia/pkg/yaml"
)

func Exporter(output, outputDir string) (bible.Exporter, error) {
	switch output {
	case "stdout":
		return text.NewExporter(os.Stdout), nil
	case "yaml":
		return yaml.NewFastExporter(outputDir), nil
	case "yaml-strict":
		return yaml.NewExporter(outputDir), nil
	default:
		return nil, fmt.Errorf("invalid format: %s", output)
	}
}
