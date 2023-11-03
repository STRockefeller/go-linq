package linq

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

// NumberLinq provides following new methods:
//  - Sum
//  - Max
//  - Min
type NumberLinq[T any, N number] struct {
	linq[T]
}

func NewNumberLinq[T any, N number](items []T) NumberLinq[T, N] {
	return NumberLinq[T, N]{linq[T]{items: items}}
}

func (nl NumberLinq[T, N]) Sum(selector func(T) N) N {
	var sum N
	for _, elem := range nl.items {
		sum += selector(elem)
	}
	return sum
}

func (nl NumberLinq[T, N]) Max(selector func(T) N) N {
	var max N
	for i, elem := range nl.items {
		num := selector(elem)
		if i == 0 || num > max {
			max = num
		}
	}
	return max
}

func (nl NumberLinq[T, N]) Min(selector func(T) N) N {
	var min N
	for i, elem := range nl.items {
		num := selector(elem)
		if i == 0 || num < min {
			min = num
		}
	}
	return min
}
