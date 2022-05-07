package exporters

import "github.com/mniak/biblia/pkg/bible"

type Exporter interface {
	Export(t bible.Testament) error
}
