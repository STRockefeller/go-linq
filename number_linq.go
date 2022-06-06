package linq

type number interface {
	int | int16 | int32 | int64 | float32 | float64
}

type NumberLinq[T LinqableType, N number] struct {
	Linq[T]
}

func (linq NumberLinq[T, N]) Sum(selector func(T) N) N {
	var sum N
	for _, elem := range linq.Linq {
		sum += selector(elem)
	}
	return sum
}

func (linq NumberLinq[T, N]) Max(selector func(T) N) N {
	var max N
	for i, elem := range linq.Linq {
		num := selector(elem)
		if i == 0 || num > max {
			max = num
		}
	}
	return max
}

func (linq NumberLinq[T, N]) Min(selector func(T) N) N {
	var min N
	for i, elem := range linq.Linq {
		num := selector(elem)
		if i == 0 || num < min {
			min = num
		}
	}
	return min
}
