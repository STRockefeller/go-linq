package linq

func NoPredict[T any]() func(T) bool {
	return func(T) bool {
		return true
	}
}
