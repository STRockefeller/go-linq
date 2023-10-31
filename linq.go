package linq

import (
	"fmt"
	"reflect"
	"sort"

	"golang.org/x/exp/constraints"
)

func equal[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

// linq simulates C# System.Linq Enumerable methods and System.Collections.Generic List methods.
// Methods of linq will panic when something goes wrong.
type linq[T any] []T

// linq constructor
func New[T any](slice []T) linq[T] {
	return linq[T](slice)
}

// linq constructor
func NewFromMap[K comparable, V any, T any](m map[K]V, delegate func(K, V) T) linq[T] {
	res := make([]T, 0, len(m))
	for k, v := range m {
		res = append(res, delegate(k, v))
	}
	return res
}

// linq constructor
// ! Make sure to close the channel when you are done sending elements to it.
func NewFromChannel[T any](c <-chan T) linq[T] {
	res := make([]T, 0)
	for v := range c {
		res = append(res, v)
	}
	return res
}

// Contains determines whether a sequence contains a specified element.
func (l linq[T]) Contains(target T) bool {
	for _, elem := range l {
		if equal(elem, target) {
			return true
		}
	}
	return false
}

// Count returns a number that represents how many elements in the specified sequence satisfy a condition.
func (l linq[T]) Count(predicate func(T) bool) int {
	var count int
	for _, elem := range l {
		if predicate(elem) {
			count++
		}
	}
	return count
}

// Distinct returns distinct elements from a sequence by using the default equality comparer to compare values.
func (l linq[T]) Distinct() linq[T] {
	res := l.Empty()
	for _, elem := range l {
		if !res.Contains(elem) {
			res = res.Append(elem)
		}
	}
	return res
}

// Any determines whether any element of a sequence satisfies a condition.
func (l linq[T]) Any(predicate func(T) bool) bool {
	for _, elem := range l {
		if predicate(elem) {
			return true
		}
	}
	return false
}

// All determines whether all elements of a sequence satisfy a condition.
func (l linq[T]) All(predicate func(T) bool) bool {
	for _, elem := range l {
		if predicate(elem) {
			continue
		} else {
			return false
		}
	}
	return true
}

// Append appends a value to the end of the sequence.
func (l linq[T]) Append(t ...T) linq[T] {
	return append(l, t...)
}

// Prepend adds a value to the beginning of the sequence.
func (l linq[T]) Prepend(t ...T) linq[T] {
	return append(t, l.ToSlice()...)
}

// ElementAt returns the element at a specified index in a sequence.
// ! this method panics when index is out of range.
func (l linq[T]) ElementAt(index int) T {
	if index >= len(l) {
		panic("linq: ElementAt() out of index")
	}
	return l[index]
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
func (l linq[T]) ElementAtOrDefault(index int) T {
	var defaultValue T
	if index >= len(l) || index < 0 {
		return defaultValue
	}
	return l[index]
}

// Empty returns an empty linq[T] that has the specified type argument.
func (l linq[T]) Empty() linq[T] {
	return linq[T]{}
}

// First returns the first element in a sequence that satisfies a specified condition.
// ! this method panics when no element is found.
func (l linq[T]) First(predicate func(T) bool) T {
	if len(l) == 0 {
		panic("linq: First() empty set")
	}
	for _, elem := range l {
		if predicate(elem) {
			return elem
		}
	}
	panic("linq: First() no match element in the slice")
}

// FirstOrDefault returns the first element of a sequence, or a default value if the sequence contains no elements.
func (l linq[T]) FirstOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if len(l) == 0 {
		return defaultValue
	}
	for _, elem := range l {
		if predicate(elem) {
			return elem
		}
	}
	return defaultValue
}

// Last returns the last element of a sequence.
// ! this method panics when no element is found.
func (l linq[T]) Last(predicate func(T) bool) T {
	if len(l) == 0 {
		panic("linq: Last() empty set")
	}
	for i := len(l) - 1; i >= 0; i-- {
		if predicate(l[i]) {
			return l[i]
		}
	}
	panic("linq: Last() no match element in the slice")
}

// LastOrDefault returns the last element of a sequence, or a specified default value if the sequence contains no elements.
func (l linq[T]) LastOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if len(l) == 0 {
		return defaultValue
	}
	for i := len(l) - 1; i >= 0; i-- {
		if predicate(l[i]) {
			return l[i]
		}
	}
	return defaultValue
}

// Single returns the only element of a sequence that satisfies a specified condition, and panics if more than one such element exists.
func (l linq[T]) Single(predicate func(T) bool) T {
	if l.Count(predicate) == 1 {
		return l.First(predicate)
	}
	panic("linq: Single() eligible data count is not unique")
}

// SingleOrDefault returns the only element of a sequence, or a default value of T if the sequence is empty.
func (l linq[T]) SingleOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if l.Count(predicate) == 1 {
		return l.First(predicate)
	}
	return defaultValue
}

// Where filters a sequence of values based on a predicate.
func (l linq[T]) Where(predicate func(T) bool) linq[T] {
	res := []T{}
	for _, elem := range l {
		if predicate(elem) {
			res = append(res, elem)
		}
	}
	return res
}

// Reverse inverts the order of the elements in a sequence.
func (l linq[T]) Reverse() linq[T] {
	res := linq[T](make([]T, len(l)))
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = l[j], l[i]
	}
	return res
}

// Take returns a specified number of contiguous elements from the start of a sequence.
// ! this method panics when count is out of range.
func (l linq[T]) Take(count int) linq[T] {
	if count < 0 || count >= len(l) {
		panic("Linq: Take() out of index")
	}
	res := []T{}
	for i := 0; i < count; i++ {
		res = append(res, l[i])
	}
	return res
}

// TakeWhile returns elements from a sequence as long as a specified condition is true. The element's index is used in the logic of the predicate function.
func (l linq[T]) TakeWhile(predicate func(T) bool) linq[T] {
	res := []T{}
	for i := 0; i < len(l); i++ {
		if predicate(l[i]) {
			res = append(res, l[i])
		} else {
			return res
		}
	}
	return res
}

// TakeLast returns a new enumerable collection that contains the last count elements from source.
// ! this method panics when count is out of range.
func (l linq[T]) TakeLast(count int) linq[T] {
	if count < 0 || count >= len(l) {
		panic("Linq: TakeLast() out of index")
	}
	return l.Skip(len(l) - count)
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
// ! this method panics when count is out of range.
func (l linq[T]) Skip(count int) linq[T] {
	if count < 0 || count >= len(l) {
		panic("Linq: Skip() out of index")
	}
	return l[count:]
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements. The element's index is used in the logic of the predicate function.
func (l linq[T]) SkipWhile(predicate func(T) bool) linq[T] {
	for i := 0; i < len(l); i++ {
		if predicate(l[i]) {
			continue
		} else {
			return l[i:]
		}
	}
	return linq[T]{}
}

// SkipLast returns a new enumerable collection that contains the elements from source with the last count elements of the source collection omitted.
// ! this method panics when count is out of range.
func (l linq[T]) SkipLast(count int) linq[T] {
	if count < 0 || count > len(l) {
		panic("Linq: SkipLast() out of index")
	}
	return l.Take(len(l) - count)
}

// Select projects each element of linq into a new form by incorporating the element's index.
func Select[T, S any](l linq[T], delegate func(T) S) linq[S] {
	res := make([]S, len(l))
	for i, elem := range l {
		res[i] = delegate(elem)
	}
	return res
}

// SelectMany takes a slice of slices and a selector function,
// and returns a flattened slice of elements selected by the selector function.
func SelectMany[T any, U any](l linq[T], selector func(T) []U) linq[U] {
	var res []U

	l.ForEach(func(t T) {
		res = append(res, selector(t)...)
	})

	return res
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func OrderBy[L any, O constraints.Ordered](l linq[L], comparer func(L) O) linq[L] {
	sort.SliceStable(l, func(i, j int) bool {
		return comparer(l[i]) < comparer(l[j])
	})
	return l
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func OrderByDescending[L any, O constraints.Ordered](l linq[L], comparer func(L) O) linq[L] {
	sort.SliceStable(l, func(i, j int) bool {
		return comparer(l[i]) > comparer(l[j])
	})
	return l
}

func GroupBy[L any, K comparable, E any](l linq[L], key func(L) K, element func(L) E) map[K][]E {
	res := make(map[K][]E)
	l.ForEach(func(l L) {
		elem := element(l)
		if _, ok := res[key(l)]; ok {
			res[key(l)] = append(res[key(l)], elem)
		} else {
			res[key(l)] = []E{elem}
		}
	})
	return res
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func (l linq[T]) OrderBy(comparer func(T) int) linq[T] {
	sort.SliceStable(l, func(i, j int) bool {
		return comparer(l[i]) < comparer(l[j])
	})
	return l
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func (l linq[T]) OrderByDescending(comparer func(T) int) linq[T] {
	sort.SliceStable(l, func(i, j int) bool {
		return comparer(l[i]) > comparer(l[j])
	})
	return l
}

// Repeat generates a sequence that contains one repeated value.
func Repeat[T any](element T, count int) linq[T] {
	if count <= 0 {
		return []T{}
	}
	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = element
	}
	return res
}

// ToSlice creates a slice from a linq[T].
func (l linq[T]) ToSlice() []T {
	res := make([]T, len(l))
	copy(res, l)
	return res
}

// ToChannel creates a channel with values in linq[T]
func (l linq[T]) ToChannel() <-chan T {
	res := make(chan T, len(l))
	l.ForEach(func(t T) {
		res <- t
	})
	close(res)
	return res
}

// ToChannelWithBuffer creates a channel with values in linq[T] with specified buffer. (async method)
func (l linq[T]) ToChannelWithBuffer(buffer int) <-chan T {
	res := make(chan T, buffer)
	go func() {
		l.ForEach(func(t T) {
			res <- t
		})
		close(res)
	}()
	return res
}

// Creates a map[interface{}]T from an linq[T] according to a specified key selector function.
func (l linq[T]) ToMapWithKey(keySelector func(T) interface{}) map[interface{}]T {
	res := make(map[interface{}]T)
	l.ForEach(func(t T) {
		res[keySelector(t)] = t
	})
	return res
}

// Creates a map[TKey]TSource from an linq[TSource] according to a specified key selector function.
func ConvertToMapWithKey[TSource any, TKey comparable](l linq[TSource], keySelector func(TSource) TKey) map[TKey]TSource {
	res := make(map[TKey]TSource)
	l.ForEach(func(t TSource) {
		res[keySelector(t)] = t
	})
	return res
}

// Creates a map[interface{}]interface from an linq[T] according to a specified key selector function.
func (l linq[T]) ToMapWithKeyValue(keySelector func(T) interface{}, valueSelector func(T) interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	l.ForEach(func(t T) {
		res[keySelector(t)] = valueSelector(t)
	})
	return res
}

// Creates a map[TKey,TValue] from an linq[TSource] according to specified key selector and element selector functions.
func ConvertToMapWithKeyValue[TSource any, TKey comparable, TValue any](l linq[TSource], keySelector func(TSource) TKey, valueSelector func(TSource) TValue) map[TKey]TValue {
	res := make(map[TKey]TValue)
	l.ForEach(func(t TSource) {
		res[keySelector(t)] = valueSelector(t)
	})
	return res
}

// #region not linq

// Add adds an object to the end of the linq[T].
func (l *linq[T]) Add(element T) {
	*l = append(*l, element)
}

// AddRange adds the elements of the specified collection to the end of the linq[T].
func (l *linq[T]) AddRange(collection linq[T]) {
	*l = append(*l, collection...)
}

// Clear removes all elements from the linq[T].
func (l *linq[T]) Clear() {
	*l = linq[T](make([]T, cap(l.ToSlice())))
}

// Clone returns a copy of linq[T]
func (l linq[T]) Clone() linq[T] {
	return l.ToSlice()
}

// Exists determines whether the linq[T] contains elements that match the conditions defined by the specified predicate.
func (l linq[T]) Exists(predicate func(T) bool) bool {
	return l.Any(predicate)
}

// Find Searches for an element that matches the conditions defined by the specified predicate, and returns the first occurrence within the entire linq[T].
func (l linq[T]) Find(predicate func(T) bool) T {
	return l.FirstOrDefault(predicate)
}

// FindAll retrieves all the elements that match the conditions defined by the specified predicate.
func (l linq[T]) FindAll(predicate func(T) bool) linq[T] {
	return l.Where(predicate)
}

// FindIndex searches for an element that matches the conditions defined by the specified predicate, and returns the zero-based index of the first occurrence within the entire linq[T].
func (l linq[T]) FindIndex(predicate func(T) bool) int {
	for i, elem := range l {
		if predicate(elem) {
			return i
		}
	}
	return -1
}

// FindLast searches for an element that matches the conditions defined by the specified predicate, and returns the last occurrence within the entire linq[T].
func (l linq[T]) FindLast(predicate func(T) bool) T {
	return l.LastOrDefault(predicate)
}

// FindLastIndex searches for an element that matches the conditions defined by a specified predicate, and returns the zero-based index of the last occurrence within the linq[T] or a portion of it.
func (l linq[T]) FindLastIndex(predicate func(T) bool) int {
	res := -1
	for i, elem := range l {
		if predicate(elem) {
			res = i
		}
	}
	return res
}

// ForEach performs the specified action on each element of the linq[T].
func (l linq[T]) ForEach(callBack func(T)) {
	for _, elem := range l {
		callBack(elem)
	}
}

// ReplaceAll replaces old values by new values
func (l linq[T]) ReplaceAll(oldValue, newValue T) linq[T] {
	res := linq[T](make([]T, 0, len(l)))
	for _, elem := range l {
		if equal(elem, oldValue) {
			res = res.Append(newValue)
		} else {
			res = res.Append(elem)
		}
	}
	return res
}

// Remove removes the first occurrence of a specific object from the linq[T].
func (l *linq[T]) Remove(item T) bool {
	res := linq[T]([]T{})
	var isRemoved bool
	for _, elem := range *l {
		if equal(elem, item) && !isRemoved {
			isRemoved = true
			continue
		}
		res = res.Append(elem)
	}
	*l = res
	return isRemoved
}

// RemoveAll removes all the elements that match the conditions defined by the specified predicate.
func (l *linq[T]) RemoveAll(predicate func(T) bool) int {
	var count int
	res := linq[T]([]T{})
	for _, elem := range *l {
		if predicate(elem) {
			count++
			continue
		}
		res = res.Append(elem)
	}
	*l = res
	return count
}

// RemoveAt removes the element at the specified index of the linq[T].
func (l *linq[T]) RemoveAt(index int) {
	res := linq[T]([]T{})
	for i := 0; i < len(*l); i++ {
		if i == index {
			continue
		}
		res = res.Append((*l)[i])
	}
	*l = res
}

// RemoveRange removes a range of elements from the linq[T].
func (l *linq[T]) RemoveRange(index, count int) error {
	if index < 0 || count < 0 || index+count > len(*l) {
		return fmt.Errorf("argument out of range")
	}
	res := linq[T]([]T{})
	for i := 0; i < len(*l); i++ {
		if i >= index && count != 0 {
			count--
			continue
		}
		res = res.Append((*l)[i])
	}
	*l = res
	return nil
}

// #endregion not linq
