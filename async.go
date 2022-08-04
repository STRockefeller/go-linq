package linq

import "sync"

func RunInAsync[I comparable, O any](inputs Linq[I], delegate func(I) O) []O {
	res := make([]O, len(inputs))
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(inputs))
	for i, input := range inputs {
		i := i
		input := input
		go func() {
			defer wg.Done()
			result := delegate(input)
			m.Lock()
			res[i] = result
			m.Unlock()
		}()
	}
	wg.Wait()
	return res
}

func (linq Linq[T]) RunInAsyncWithRoutineLimit(delegate func(T), limit int) {
	inputs := make(chan T, 5)
	go func() {
		defer close(inputs)
		linq.ForEach(func(t T) {
			inputs <- t
		})
	}()

	consumer := func() {
		for inp := range inputs {
			delegate(inp)
		}
	}

	for i := 0; i < limit; i++ {
		consumer()
	}
}
