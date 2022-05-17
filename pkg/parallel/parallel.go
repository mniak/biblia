package parallel

import (
	"context"
	"sync"
)

type foridata[T any] struct {
	i int
	v T
}

func ForI[T any](initial, exclusiveMax int, fn func(int) (T, error)) ([]T, error) {
	return ForIContext(context.Background(), initial, exclusiveMax, func(ctx context.Context, i int) (T, error) {
		return fn(i)
	})
}

func ForIContext[T any](ctx context.Context, initial, exclusiveMax int, fn func(context.Context, int) (T, error)) ([]T, error) {
	cherr := make(chan error)
	chdata := make(chan foridata[T])
	chdone := make(chan bool)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	for i := initial; i < exclusiveMax; i++ {
		wg.Add(1)
		go func(i int) {
			obj, err := fn(ctx, i)
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
