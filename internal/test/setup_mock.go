package test

type _mock[E any] interface {
	EXPECT() *E
}

func Setup[T _mock[E], E any](mock T, setup func(e *E)) T {
	expect := mock.EXPECT()
	setup(expect)
	return mock
}
