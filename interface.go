package linq

type Linq[T any] interface {
	// All determines whether all elements of a sequence satisfy a condition.
	All(predicate func(T) bool) bool
	// Any determines whether any element of a sequence satisfies a condition.
	Any(predicate func(T) bool) bool
	// Append appends a value to the end of the sequence.
	Append(t ...T) Linq[T]
	// Clone returns a copy of linq[T]
	Clone() Linq[T]
	// Contains determines whether a sequence contains a specified element.
	Contains(target T) bool
	// Count returns a number that represents how many elements in the specified sequence satisfy a condition.
	Count(predicate func(T) bool) int
	// Distinct returns distinct elements from a sequence by using the default equality comparer to compare values.
	Distinct() Linq[T]
	// ElementAt returns the element at a specified index in a sequence.
	// ! this method panics when index is out of range.
	ElementAt(index int) T
	// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
	ElementAtOrDefault(index int) T
	// Empty returns an empty linq[T] that has the specified type argument.
	Empty() Linq[T]
	// Exists determines whether the linq[T] contains elements that match the conditions defined by the specified predicate.
	Exists(predicate func(T) bool) bool
	// Find Searches for an element that matches the conditions defined by the specified predicate, and returns the first occurrence within the entire linq[T].
	Find(predicate func(T) bool) T
	// FindAll retrieves all the elements that match the conditions defined by the specified predicate.
	FindAll(predicate func(T) bool) Linq[T]
	// FindIndex searches for an element that matches the conditions defined by the specified predicate, and returns the zero-based index of the first occurrence within the entire linq[T].
	FindIndex(predicate func(T) bool) int
	// FindLast searches for an element that matches the conditions defined by the specified predicate, and returns the last occurrence within the entire linq[T].
	FindLast(predicate func(T) bool) T
	// FindLastIndex searches for an element that matches the conditions defined by a specified predicate, and returns the zero-based index of the last occurrence within the linq[T] or a portion of it.
	FindLastIndex(predicate func(T) bool) int
	// First returns the first element in a sequence that satisfies a specified condition.
	// ! this method panics when no element is found.
	First(predicate func(T) bool) T
	// FirstOrDefault returns the first element of a sequence, or a default value if the sequence contains no elements.
	FirstOrDefault(predicate func(T) bool) T
	// ForEach performs the specified action on each element of the linq[T].
	ForEach(callBack func(T))
	// Last returns the last element of a sequence.
	// ! this method panics when no element is found.
	Last(predicate func(T) bool) T
	// LastOrDefault returns the last element of a sequence, or a specified default value if the sequence contains no elements.
	LastOrDefault(predicate func(T) bool) T
	// OrderBy sorts the elements of a sequence in ascending order according to a key.
	OrderBy(comparer func(T) int) Linq[T]
	// OrderByDescending sorts the elements of a sequence in descending order according to a key.
	OrderByDescending(comparer func(T) int) Linq[T]
	// Prepend adds a value to the beginning of the sequence.
	Prepend(t ...T) Linq[T]
	// ReplaceAll replaces old values by new values
	ReplaceAll(oldValue T, newValue T) Linq[T]
	// Reverse inverts the order of the elements in a sequence.
	Reverse() Linq[T]
	RunInAsyncWithRoutineLimit(delegate func(T), limit int)
	// Single returns the only element of a sequence that satisfies a specified condition, and panics if more than one such element exists.
	Single(predicate func(T) bool) T
	// SingleOrDefault returns the only element of a sequence, or a default value of T if the sequence is empty.
	SingleOrDefault(predicate func(T) bool) T
	// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
	// ! this method panics when count is out of range.
	Skip(count int) Linq[T]
	// SkipLast returns a new enumerable collection that contains the elements from source with the last count elements of the source collection omitted.
	// ! this method panics when count is out of range.
	SkipLast(count int) Linq[T]
	// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements. The element's index is used in the logic of the predicate function.
	SkipWhile(predicate func(T) bool) Linq[T]
	// Take returns a specified number of contiguous elements from the start of a sequence.
	// ! this method panics when count is out of range.
	Take(count int) Linq[T]
	// TakeLast returns a new enumerable collection that contains the last count elements from source.
	// ! this method panics when count is out of range.
	TakeLast(count int) Linq[T]
	// TakeWhile returns elements from a sequence as long as a specified condition is true. The element's index is used in the logic of the predicate function.
	TakeWhile(predicate func(T) bool) Linq[T]
	// ToChannel creates a channel with values in linq[T]
	ToChannel() <-chan T
	// ToChannelWithBuffer creates a channel with values in linq[T] with specified buffer. (async method)
	ToChannelWithBuffer(buffer int) <-chan T
	// Creates a map[interface{}]T from an linq[T] according to a specified key selector function.
	ToMapWithKey(keySelector func(T) interface{}) map[interface{}]T
	// Creates a map[interface{}]interface from an linq[T] according to a specified key selector function.
	ToMapWithKeyValue(keySelector func(T) interface{}, valueSelector func(T) interface{}) map[interface{}]interface{}
	// ToSlice creates a slice from a linq[T].
	ToSlice() []T
	// Where filters a sequence of values based on a predicate.
	Where(predicate func(T) bool) Linq[T]
	// Length returns the number of items in the linq[T] collection.
	Length() int

	/* ------------------------ pointer receiver methods ------------------------ */

	// Add adds an object to the end of the linq[T].
	Add(element T)
	// AddRange adds the elements of the specified collection to the end of the linq[T].
	AddRange(collection []T)
	// Remove removes the first occurrence of a specific object from the linq[T].
	Remove(item T) bool
	// RemoveAll removes all the elements that match the conditions defined by the specified predicate.
	RemoveAll(predicate func(T) bool) int
	// RemoveAt removes the element at the specified index of the linq[T].
	RemoveAt(index int)
	// RemoveRange removes a range of elements from the linq[T].
	RemoveRange(index int, count int) error
	// Clear removes all elements from the linq[T].
	Clear()
}
