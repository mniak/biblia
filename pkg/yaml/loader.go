package yaml

type yamlTestamentLoader struct {
	directory string
}

func NewLoader(directory string) yamlTestamentLoader {
	return yamlTestamentLoader{
		directory: directory,
	}
}
