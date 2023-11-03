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
type linq[T any] struct {
	items []T
}

// linq constructor
func New[T any](slice []T) Linq[T] {
	return &linq[T]{
		items: slice,
	}
}

// linq constructor
func NewFromMap[K comparable, V any, T any](m map[K]V, delegate func(K, V) T) Linq[T] {
	res := make([]T, 0, len(m))
	for k, v := range m {
		res = append(res, delegate(k, v))
	}
	return New(res)
}

// linq constructor
// ! Make sure to close the channel when you are done sending elements to it.
func NewFromChannel[T any](c <-chan T) Linq[T] {
	res := make([]T, 0)
	for v := range c {
		res = append(res, v)
	}
	return New(res)
}

// Contains determines whether a sequence contains a specified element.
func (l linq[T]) Contains(target T) bool {
	for _, elem := range l.items {
		if equal(elem, target) {
			return true
		}
	}
	return false
}

// Count returns a number that represents how many elements in the specified sequence satisfy a condition.
func (l linq[T]) Count(predicate func(T) bool) int {
	var count int
	for _, elem := range l.items {
		if predicate(elem) {
			count++
		}
	}
	return count
}

// Distinct returns distinct elements from a sequence by using the default equality comparer to compare values.
func (l linq[T]) Distinct() Linq[T] {
	res := l.Empty()
	for _, elem := range l.items {
		if !res.Contains(elem) {
			res = res.Append(elem)
		}
	}
	return res
}

// Any determines whether any element of a sequence satisfies a condition.
func (l linq[T]) Any(predicate func(T) bool) bool {
	for _, elem := range l.items {
		if predicate(elem) {
			return true
		}
	}
	return false
}

// All determines whether all elements of a sequence satisfy a condition.
func (l linq[T]) All(predicate func(T) bool) bool {
	for _, elem := range l.items {
		if predicate(elem) {
			continue
		} else {
			return false
		}
	}
	return true
}

// Append appends a value to the end of the sequence.
func (l linq[T]) Append(t ...T) Linq[T] {
	return New(append(l.items, t...))
}

// Prepend adds a value to the beginning of the sequence.
func (l linq[T]) Prepend(t ...T) Linq[T] {
	return New(append(t, l.ToSlice()...))
}

// ElementAt returns the element at a specified index in a sequence.
// ! this method panics when index is out of range.
func (l linq[T]) ElementAt(index int) T {
	if index >= len(l.items) {
		panic("linq: ElementAt() out of index")
	}
	return l.items[index]
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
func (l linq[T]) ElementAtOrDefault(index int) T {
	var defaultValue T
	if index >= len(l.items) || index < 0 {
		return defaultValue
	}
	return l.items[index]
}

// Empty returns an empty linq[T] that has the specified type argument.
func (l linq[T]) Empty() Linq[T] {
	return New([]T{})
}

// First returns the first element in a sequence that satisfies a specified condition.
// ! this method panics when no element is found.
func (l linq[T]) First(predicate func(T) bool) T {
	if len(l.items) == 0 {
		panic("linq: First() empty set")
	}
	for _, elem := range l.items {
		if predicate(elem) {
			return elem
		}
	}
	panic("linq: First() no match element in the slice")
}

// FirstOrDefault returns the first element of a sequence, or a default value if the sequence contains no elements.
func (l linq[T]) FirstOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if len(l.items) == 0 {
		return defaultValue
	}
	for _, elem := range l.items {
		if predicate(elem) {
			return elem
		}
	}
	return defaultValue
}

// Last returns the last element of a sequence.
// ! this method panics when no element is found.
func (l linq[T]) Last(predicate func(T) bool) T {
	if len(l.items) == 0 {
		panic("linq: Last() empty set")
	}
	for i := len(l.items) - 1; i >= 0; i-- {
		if predicate(l.items[i]) {
			return l.items[i]
		}
	}
	panic("linq: Last() no match element in the slice")
}

// LastOrDefault returns the last element of a sequence, or a specified default value if the sequence contains no elements.
func (l linq[T]) LastOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if len(l.items) == 0 {
		return defaultValue
	}
	for i := len(l.items) - 1; i >= 0; i-- {
		if predicate(l.items[i]) {
			return l.items[i]
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
func (l linq[T]) Where(predicate func(T) bool) Linq[T] {
	res := []T{}
	for _, elem := range l.items {
		if predicate(elem) {
			res = append(res, elem)
		}
	}
	return New(res)
}

// Reverse inverts the order of the elements in a sequence.
func (l linq[T]) Reverse() Linq[T] {
	res := make([]T, len(l.items))
	for i, j := 0, len(l.items)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = l.items[j], l.items[i]
	}
	return New(res)
}

// Take returns a specified number of contiguous elements from the start of a sequence.
// ! this method panics when count is out of range.
func (l linq[T]) Take(count int) Linq[T] {
	if count < 0 || count >= len(l.items) {
		panic("Linq: Take() out of index")
	}
	res := []T{}
	for i := 0; i < count; i++ {
		res = append(res, l.items[i])
	}
	return New(res)
}

// TakeWhile returns elements from a sequence as long as a specified condition is true. The element's index is used in the logic of the predicate function.
func (l linq[T]) TakeWhile(predicate func(T) bool) Linq[T] {
	res := []T{}
	for i := 0; i < len(l.items); i++ {
		if predicate(l.items[i]) {
			res = append(res, l.items[i])
		} else {
			return New(res)
		}
	}
	return New(res)
}

// TakeLast returns a new enumerable collection that contains the last count elements from source.
// ! this method panics when count is out of range.
func (l linq[T]) TakeLast(count int) Linq[T] {
	if count < 0 || count >= len(l.items) {
		panic("Linq: TakeLast() out of index")
	}
	return l.Skip(len(l.items) - count)
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
// ! this method panics when count is out of range.
func (l linq[T]) Skip(count int) Linq[T] {
	if count < 0 || count >= len(l.items) {
		panic("Linq: Skip() out of index")
	}
	return New(l.items[count:])
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements. The element's index is used in the logic of the predicate function.
func (l linq[T]) SkipWhile(predicate func(T) bool) Linq[T] {
	for i := 0; i < len(l.items); i++ {
		if predicate(l.items[i]) {
			continue
		} else {
			return New(l.items[i:])
		}
	}
	return l.Empty()
}

// SkipLast returns a new enumerable collection that contains the elements from source with the last count elements of the source collection omitted.
// ! this method panics when count is out of range.
func (l linq[T]) SkipLast(count int) Linq[T] {
	if count < 0 || count > len(l.items) {
		panic("Linq: SkipLast() out of index")
	}
	return l.Take(len(l.items) - count)
}

// Select projects each element of linq into a new form by incorporating the element's index.
func Select[T, S any](items []T, delegate func(T) S) Linq[S] {
	res := make([]S, len(items))
	for i, elem := range items {
		res[i] = delegate(elem)
	}
	return New(res)
}

// SelectMany takes a slice of slices and a selector function,
// and returns a flattened slice of elements selected by the selector function.
func SelectMany[T any, U any](items []T, selector func(T) []U) Linq[U] {
	var res []U

	for _, t := range items {
		res = append(res, selector(t)...)
	}

	return New(res)
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func OrderBy[L any, O constraints.Ordered](items []L, comparer func(L) O) Linq[L] {
	sort.SliceStable(items, func(i, j int) bool {
		return comparer(items[i]) < comparer(items[j])
	})
	return New(items)
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func OrderByDescending[L any, O constraints.Ordered](items []L, comparer func(L) O) Linq[L] {
	sort.SliceStable(items, func(i, j int) bool {
		return comparer(items[i]) > comparer(items[j])
	})
	return New(items)
}

func GroupBy[L any, K comparable, E any](items []L, key func(L) K, element func(L) E) map[K][]E {
	res := make(map[K][]E)

	for _, item := range items {
		elem := element(item)
		if _, ok := res[key(item)]; ok {
			res[key(item)] = append(res[key(item)], elem)
		} else {
			res[key(item)] = []E{elem}
		}
	}
	return res
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func (l linq[T]) OrderBy(comparer func(T) int) Linq[T] {
	sort.SliceStable(l.items, func(i, j int) bool {
		return comparer(l.items[i]) < comparer(l.items[j])
	})
	return &l
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func (l linq[T]) OrderByDescending(comparer func(T) int) Linq[T] {
	sort.SliceStable(l.items, func(i, j int) bool {
		return comparer(l.items[i]) > comparer(l.items[j])
	})
	return &l
}

// Repeat generates a sequence that contains one repeated value.
func Repeat[T any](element T, count int) Linq[T] {
	if count <= 0 {
		return New([]T{})
	}
	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = element
	}
	return New(res)
}

// ToSlice creates a slice from a linq[T].
func (l linq[T]) ToSlice() []T {
	res := make([]T, len(l.items))
	copy(res, l.items)
	return res
}

// ToChannel creates a channel with values in linq[T]
func (l linq[T]) ToChannel() <-chan T {
	res := make(chan T, len(l.items))
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
func ConvertToMapWithKey[TSource any, TKey comparable](items []TSource, keySelector func(TSource) TKey) map[TKey]TSource {
	res := make(map[TKey]TSource)
	for _, item := range items {
		res[keySelector(item)] = item
	}
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
func ConvertToMapWithKeyValue[TSource any, TKey comparable, TValue any](items []TSource, keySelector func(TSource) TKey, valueSelector func(TSource) TValue) map[TKey]TValue {
	res := make(map[TKey]TValue)

	for _, item := range items {
		res[keySelector(item)] = valueSelector(item)
	}
	return res
}

// #region not linq

// Add adds an object to the end of the linq[T].
func (l *linq[T]) Add(element T) {
	l.items = append(l.items, element)
}

// AddRange adds the elements of the specified collection to the end of the linq[T].
func (l *linq[T]) AddRange(collection []T) {
	l.items = append(l.items, collection...)
}

// Clear removes all elements from the linq[T].
func (l *linq[T]) Clear() {
	l.items = make([]T, cap(l.ToSlice()))
}

// Clone returns a copy of linq[T]
func (l linq[T]) Clone() Linq[T] {
	return New(l.ToSlice())
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
func (l linq[T]) FindAll(predicate func(T) bool) Linq[T] {
	return l.Where(predicate)
}

// FindIndex searches for an element that matches the conditions defined by the specified predicate, and returns the zero-based index of the first occurrence within the entire linq[T].
func (l linq[T]) FindIndex(predicate func(T) bool) int {
	for i, elem := range l.items {
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
	for i, elem := range l.items {
		if predicate(elem) {
			res = i
		}
	}
	return res
}

// ForEach performs the specified action on each element of the linq[T].
func (l linq[T]) ForEach(callBack func(T)) {
	for _, elem := range l.items {
		callBack(elem)
	}
}

// ReplaceAll replaces old values by new values
func (l linq[T]) ReplaceAll(oldValue, newValue T) Linq[T] {
	res := New(make([]T, 0, len(l.items)))
	for _, elem := range l.items {
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
	res := l.Empty()
	var isRemoved bool
	for _, elem := range l.items {
		if equal(elem, item) && !isRemoved {
			isRemoved = true
			continue
		}
		res = res.Append(elem)
	}
	*l = *res.(*linq[T])
	return isRemoved
}

// RemoveAll removes all the elements that match the conditions defined by the specified predicate.
func (l *linq[T]) RemoveAll(predicate func(T) bool) int {
	var count int
	res := l.Empty()
	for _, elem := range l.items {
		if predicate(elem) {
			count++
			continue
		}
		res = res.Append(elem)
	}
	*l = *res.(*linq[T])
	return count
}

// RemoveAt removes the element at the specified index of the linq[T].
func (l *linq[T]) RemoveAt(index int) {
	res := l.Empty()
	for i := 0; i < len(l.items); i++ {
		if i == index {
			continue
		}
		res = res.Append(l.items[i])
	}
	*l = *res.(*linq[T])
}

// RemoveRange removes a range of elements from the linq[T].
func (l *linq[T]) RemoveRange(index, count int) error {
	if index < 0 || count < 0 || index+count > len(l.items) {
		return fmt.Errorf("argument out of range")
	}
	res := l.Empty()
	for i := 0; i < len(l.items); i++ {
		if i >= index && count != 0 {
			count--
			continue
		}
		res = res.Append(l.items[i])
	}
	*l = *res.(*linq[T])
	return nil
}

// Length returns the number of items in the linq[T] collection.
func (l linq[T]) Length() int {
	return len(l.items)
}

// #endregion not linq
