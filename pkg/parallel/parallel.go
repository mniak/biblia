package parallel

import "sync"

type foridata[T any] struct {
	i int
	v T
}

func ForI[T any](initial, exclusiveMax int, fn func(int) (T, error)) ([]T, error) {
	cherr := make(chan error)
	defer close(cherr)

	chdata := make(chan foridata[T])
	defer close(chdata)

	chdone := make(chan bool)
	defer close(chdone)

	var wg sync.WaitGroup
	for i := initial; i < exclusiveMax; i++ {
		wg.Add(1)
		go func(i int) {
			obj, err := fn(i)
			if err != nil {
				cherr <- err
			}
			chdata <- foridata[T]{i: i - initial, v: obj}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		chdone <- true
	}()

	results := make([]T, exclusiveMax-initial)
	for {
		select {
		case data := <-chdata:
			results[data.i] = data.v

		case err := <-cherr:
			return nil, err

		case <-chdone:
			return results, nil
		}
	}
}
