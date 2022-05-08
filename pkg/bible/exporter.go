package bible

type Exporter interface {
	Export(t Testament) error
}
