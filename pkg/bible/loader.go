package bible

type TestamentLoader interface {
	Load() (Testament, error)
}
