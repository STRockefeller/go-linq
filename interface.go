package linq

type Linq[T any] interface {
	All(predicate func(T) bool) bool
	Any(predicate func(T) bool) bool
	Append(t ...T) Linq[T]
	Clone() Linq[T]
	Contains(target T) bool
	Count(predicate func(T) bool) int
	Distinct() Linq[T]
	ElementAt(index int) T
	ElementAtOrDefault(index int) T
	Empty() Linq[T]
	Exists(predicate func(T) bool) bool
	Find(predicate func(T) bool) T
	FindAll(predicate func(T) bool) Linq[T]
	FindIndex(predicate func(T) bool) int
	FindLast(predicate func(T) bool) T
	FindLastIndex(predicate func(T) bool) int
	First(predicate func(T) bool) T
	FirstOrDefault(predicate func(T) bool) T
	ForEach(callBack func(T))
	Last(predicate func(T) bool) T
	LastOrDefault(predicate func(T) bool) T
	OrderBy(comparer func(T) int) Linq[T]
	OrderByDescending(comparer func(T) int) Linq[T]
	Prepend(t ...T) Linq[T]
	ReplaceAll(oldValue T, newValue T) Linq[T]
	Reverse() Linq[T]
	RunInAsyncWithRoutineLimit(delegate func(T), limit int)
	Single(predicate func(T) bool) T
	SingleOrDefault(predicate func(T) bool) T
	Skip(count int) Linq[T]
	SkipLast(count int) Linq[T]
	SkipWhile(predicate func(T) bool) Linq[T]
	Take(count int) Linq[T]
	TakeLast(count int) Linq[T]
	TakeWhile(predicate func(T) bool) Linq[T]
	ToChannel() <-chan T
	ToChannelWithBuffer(buffer int) <-chan T
	ToMapWithKey(keySelector func(T) interface{}) map[interface{}]T
	ToMapWithKeyValue(keySelector func(T) interface{}, valueSelector func(T) interface{}) map[interface{}]interface{}
	ToSlice() []T
	Where(predicate func(T) bool) Linq[T]
	Length() int

	// pointer receiver methods
	Add(element T)
	AddRange(collection []T)
	Remove(item T) bool
	RemoveAll(predicate func(T) bool) int
	RemoveAt(index int)
	RemoveRange(index int, count int) error
	Clear()
}
