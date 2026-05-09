package bible

// LoadTransliterateAndExport loads a testament, transliterates word by word an then exports the testament
func LoadTransliterateAndExport(loader TestamentLoader, transliterator Transliterator, exporter Exporter) error {
	testament, err := loader.Load()
	if err != nil {
		return err
	}
	transliteratedTestament := testament.Transliterate(transliterator)
	return exporter.Export(transliteratedTestament)
}
