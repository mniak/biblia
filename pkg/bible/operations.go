package bible

// LoadRomanizeAndExport loads a testament, romanizes word by word an then exports the testament
func LoadRomanizeAndExport(loader TestamentLoader, romanizer Romanizer, exporter Exporter) error {
	testament, err := loader.Load()
	if err != nil {
		return err
	}
	romanizedTestament := testament.Romanize(romanizer)
	if err != nil {
		return err
	}
	return exporter.Export(romanizedTestament)
}
