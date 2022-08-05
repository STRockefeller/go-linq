package linq

import (
	"fmt"
	"sort"
)

type LinqableType interface {
	comparable
}

type Orderable interface {
	number | string
}

// Linq simulates C# System.Linq Enumerable methods and System.Collections.Generic List methods.
// Methods of Linq will panic when something goes wrong.
type Linq[T LinqableType] []T

// Contains determines whether a sequence contains a specified element.
func (linq Linq[T]) Contains(target T) bool {
	for _, elem := range linq {
		if elem == target {
			return true
		}
	}
	return false
}

// Count returns a number that represents how many elements in the specified sequence satisfy a condition.
func (linq Linq[T]) Count(predicate func(T) bool) int {
	var count int
	for _, elem := range linq {
		if predicate(elem) {
			count++
		}
	}
	return count
}

// Distinct returns distinct elements from a sequence by using the default equality comparer to compare values.
func (linq Linq[T]) Distinct() Linq[T] {
	res := linq.Empty()
	for _, elem := range linq {
		if !res.Contains(elem) {
			res = res.Append(elem)
		}
	}
	return res
}

// Any determines whether any element of a sequence satisfies a condition.
func (linq Linq[T]) Any(predicate func(T) bool) bool {
	for _, elem := range linq {
		if predicate(elem) {
			return true
		}
	}
	return false
}

// All determines whether all elements of a sequence satisfy a condition.
func (linq Linq[T]) All(predicate func(T) bool) bool {
	for _, elem := range linq {
		if predicate(elem) {
			continue
		} else {
			return false
		}
	}
	return true
}

// Append appends a value to the end of the sequence.
func (linq Linq[T]) Append(t ...T) Linq[T] {
	return append(linq, t...)
}

// Prepend adds a value to the beginning of the sequence.
func (linq Linq[T]) Prepend(t ...T) Linq[T] {
	return append(t, linq.ToSlice()...)
}

// ElementAt returns the element at a specified index in a sequence.
func (linq Linq[T]) ElementAt(index int) T {
	if index >= len(linq) {
		panic("linq: ElementAt() out of index")
	}
	return linq[index]
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
func (linq Linq[T]) ElementAtOrDefault(index int) T {
	var defaultValue T
	if index >= len(linq) {
		return defaultValue
	}
	return linq[index]
}

// Empty returns an empty Linq[T] that has the specified type argument.
func (linq Linq[T]) Empty() Linq[T] {
	return Linq[T]{}
}

// First returns the first element in a sequence that satisfies a specified condition.
func (linq Linq[T]) First(predicate func(T) bool) T {
	if len(linq) <= 0 {
		panic("linq: First() empty set")
	}
	for _, elem := range linq {
		if predicate(elem) {
			return elem
		}
	}
	panic("linq: First() no match element in the slice")
}

// FirstOrDefault returns the first element of a sequence, or a default value if the sequence contains no elements.
func (linq Linq[T]) FirstOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if len(linq) <= 0 {
		return defaultValue
	}
	for _, elem := range linq {
		if predicate(elem) {
			return elem
		}
	}
	return defaultValue
}

// Last returns the last element of a sequence.
func (linq Linq[T]) Last(predicate func(T) bool) T {
	if len(linq) <= 0 {
		panic("linq: Last() empty set")
	}
	for i := len(linq) - 1; i >= 0; i-- {
		if predicate(linq[i]) {
			return linq[i]
		}
	}
	panic("linq: Last() no match element in the slice")
}

// LastOrDefault returns the last element of a sequence, or a specified default value if the sequence contains no elements.
func (linq Linq[T]) LastOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if len(linq) <= 0 {
		return defaultValue
	}
	for i := len(linq) - 1; i >= 0; i-- {
		if predicate(linq[i]) {
			return linq[i]
		}
	}
	return defaultValue
}

// Single returns the only element of a sequence that satisfies a specified condition, and panics if more than one such element exists.
func (linq Linq[T]) Single(predicate func(T) bool) T {
	if linq.Count(predicate) == 1 {
		return linq.First(predicate)
	}
	panic("linq: Single() eligible data count is not unique")
}

// SingleOrDefault returns the only element of a sequence, or a default value of T if the sequence is empty.
func (linq Linq[T]) SingleOrDefault(predicate func(T) bool) T {
	var defaultValue T
	if linq.Count(predicate) == 1 {
		return linq.First(predicate)
	}
	return defaultValue
}

// Where filters a sequence of values based on a predicate.
func (linq Linq[T]) Where(predicate func(T) bool) Linq[T] {
	res := []T{}
	for _, elem := range linq {
		if predicate(elem) {
			res = append(res, elem)
		}
	}
	return res
}

// Reverse inverts the order of the elements in a sequence.
func (linq Linq[T]) Reverse() Linq[T] {
	res := Linq[T](make([]T, len(linq)))
	for i, j := 0, len(linq)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = linq[j], linq[i]
	}
	return res
}

// Take returns a specified number of contiguous elements from the start of a sequence.
func (linq Linq[T]) Take(n int) Linq[T] {
	if n < 0 || n >= len(linq) {
		panic("Linq: Take() out of index")
	}
	res := []T{}
	for i := 0; i < n; i++ {
		res = append(res, linq[i])
	}
	return res
}

// TakeWhile returns elements from a sequence as long as a specified condition is true. The element's index is used in the logic of the predicate function.
func (linq Linq[T]) TakeWhile(predicate func(T) bool) Linq[T] {
	res := []T{}
	for i := 0; i < len(linq); i++ {
		if predicate(linq[i]) {
			res = append(res, linq[i])
		} else {
			return res
		}
	}
	return res
}

// TakeLast returns a new enumerable collection that contains the last count elements from source.
func (linq Linq[T]) TakeLast(n int) Linq[T] {
	if n < 0 || n >= len(linq) {
		panic("Linq: TakeLast() out of index")
	}
	return linq.Skip(len(linq) - n)
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
func (linq Linq[T]) Skip(n int) Linq[T] {
	if n < 0 || n >= len(linq) {
		panic("Linq: Skip() out of index")
	}
	return linq[n:]
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements. The element's index is used in the logic of the predicate function.
func (linq Linq[T]) SkipWhile(predicate func(T) bool) Linq[T] {
	for i := 0; i < len(linq); i++ {
		if predicate(linq[i]) {
			continue
		} else {
			return linq[i:]
		}
	}
	return Linq[T]{}
}

// SkipLast returns a new enumerable collection that contains the elements from source with the last count elements of the source collection omitted.
func (linq Linq[T]) SkipLast(n int) Linq[T] {
	if n < 0 || n > len(linq) {
		panic("Linq: SkipLast() out of index")
	}
	return linq.Take(len(linq) - n)
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func OrderBy[L LinqableType, O Orderable](linq Linq[L], comparer func(L) O) Linq[L] {
	sort.SliceStable(linq, func(i, j int) bool {
		return comparer(linq[i]) < comparer(linq[j])
	})
	return linq
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func OrderByDescending[L LinqableType, O Orderable](linq Linq[L], comparer func(L) O) Linq[L] {
	sort.SliceStable(linq, func(i, j int) bool {
		return comparer(linq[i]) > comparer(linq[j])
	})
	return linq
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func (linq Linq[T]) OrderBy(comparer func(T) int) Linq[T] {
	sort.SliceStable(linq, func(i, j int) bool {
		return comparer(linq[i]) < comparer(linq[j])
	})
	return linq
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func (linq Linq[T]) OrderByDescending(comparer func(T) int) Linq[T] {
	sort.SliceStable(linq, func(i, j int) bool {
		return comparer(linq[i]) > comparer(linq[j])
	})
	return linq
}

// Repeat generates a sequence that contains one repeated value.
func Repeat[T LinqableType](element T, count int) Linq[T] {
	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = element
	}
	return res
}

// ToSlice creates a slice from a Linq[T].
func (linq Linq[T]) ToSlice() []T {
	res := make([]T, len(linq))
	copy(res, linq)
	return res
}

// #region not linq

// Add adds an object to the end of the Linq[T].
func (linq *Linq[T]) Add(element T) {
	*linq = append(*linq, element)
}

// AddRange adds the elements of the specified collection to the end of the Linq[T].
func (linq *Linq[T]) AddRange(collection Linq[T]) {
	*linq = append(*linq, collection...)
}

// Clear removes all elements from the Linq[T].
func (linq *Linq[T]) Clear() {
	*linq = Linq[T](make([]T, cap(linq.ToSlice())))
}

// Clone returns a copy of Linq[T]
func (linq Linq[T]) Clone() Linq[T] {
	return linq.ToSlice()
}

// Exists determines whether the Linq[T] contains elements that match the conditions defined by the specified predicate.
func (linq Linq[T]) Exists(predicate func(T) bool) bool {
	return linq.Any(predicate)
}

// Find Searches for an element that matches the conditions defined by the specified predicate, and returns the first occurrence within the entire Linq[T].
func (linq Linq[T]) Find(predicate func(T) bool) T {
	return linq.FirstOrDefault(predicate)
}

// FindAll retrieves all the elements that match the conditions defined by the specified predicate.
func (linq Linq[T]) FindAll(predicate func(T) bool) Linq[T] {
	return linq.Where(predicate)
}

// FindIndex searches for an element that matches the conditions defined by the specified predicate, and returns the zero-based index of the first occurrence within the entire Linq[T].
func (linq Linq[T]) FindIndex(predicate func(T) bool) int {
	for i, elem := range linq {
		if predicate(elem) {
			return i
		}
	}
	return -1
}

// FindLast searches for an element that matches the conditions defined by the specified predicate, and returns the last occurrence within the entire Linq[T].
func (linq Linq[T]) FindLast(predicate func(T) bool) T {
	return linq.LastOrDefault(predicate)
}

// FindLastIndex searches for an element that matches the conditions defined by a specified predicate, and returns the zero-based index of the last occurrence within the Linq[T] or a portion of it.
func (linq Linq[T]) FindLastIndex(predicate func(T) bool) int {
	res := -1
	for i, elem := range linq {
		if predicate(elem) {
			res = i
		}
	}
	return res
}

// ForEach performs the specified action on each element of the Linq[T].
func (linq Linq[T]) ForEach(callBack func(T)) {
	for _, elem := range linq {
		callBack(elem)
	}
}

// ReplaceAll replaces old values by new values
func (linq Linq[T]) ReplaceAll(oldValue, newValue T) Linq[T] {
	res := Linq[T]([]T{})
	for _, elem := range linq {
		if elem == oldValue {
			res = res.Append(newValue)
		} else {
			res = res.Append(elem)
		}
	}
	return res
}

// Remove removes the first occurrence of a specific object from the Linq[T].
func (linq *Linq[T]) Remove(item T) bool {
	res := Linq[T]([]T{})
	var isRemoved bool
	for _, elem := range *linq {
		if elem == item && !isRemoved {
			isRemoved = true
			continue
		}
		res = res.Append(elem)
	}
	*linq = res
	return isRemoved
}

// RemoveAll removes all the elements that match the conditions defined by the specified predicate.
func (linq *Linq[T]) RemoveAll(predicate func(T) bool) int {
	var count int
	res := Linq[T]([]T{})
	for _, elem := range *linq {
		if predicate(elem) {
			count++
			continue
		}
		res = res.Append(elem)
	}
	*linq = res
	return count
}

// RemoveAt removes the element at the specified index of the Linq[T].
func (linq *Linq[T]) RemoveAt(index int) {
	res := Linq[T]([]T{})
	for i := 0; i < len(*linq); i++ {
		if i == index {
			continue
		}
		res = res.Append((*linq)[i])
	}
	*linq = res
}

// RemoveRange removes a range of elements from the Linq[T].
func (linq *Linq[T]) RemoveRange(index, count int) error {
	if index < 0 || count < 0 || index+count > len(*linq) {
		return fmt.Errorf("argument out of range")
	}
	res := Linq[T]([]T{})
	for i := 0; i < len(*linq); i++ {
		if i >= index && count != 0 {
			count--
			continue
		}
		res = res.Append((*linq)[i])
	}
	*linq = res
	return nil
}

// #endregion not linq
