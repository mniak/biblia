package yaml

import "github.com/mniak/biblia/pkg/bible"

var _ bible.Exporter = fastYamlExporter{}
