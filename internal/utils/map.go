package utils

func MapErr[A any, B any](list []A, mapfn func(A) (B, error)) ([]B, error) {
	result := make([]B, 0)
	for _, item := range list {
		book, err := mapfn(item)
		if err != nil {
			return nil, err
		}
		result = append(result, book)
	}
	return result, nil
}
