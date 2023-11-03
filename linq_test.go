package linq

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Int_Methods(t *testing.T) {
	assert := assert.New(t)
	si := New([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	{ // ToSlice
		assert.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, si.ToSlice())
	}
	{ // ToChannel
		ch := si.ToChannel()
		var elements []int
		for element := range ch {
			elements = append(elements, element)
		}
		assert.ElementsMatch(si.ToSlice(), elements)
	}
	{ // ToChannelWithBuffer
		ch := si.ToChannelWithBuffer(2)
		var elements []int
		for element := range ch {
			elements = append(elements, element)
		}
		assert.ElementsMatch(si.ToSlice(), elements)
	}
	{ // ToMapWithKey
		m := si.ToMapWithKey(func(i int) interface{} { return i * 10 })
		assert.Equal(5, m[50])
	}
	{ // ConvertToMapWithKey
		m := ConvertToMapWithKey(si, func(i int) float32 { return float32(i) * 0.1 })
		assert.Equal(5, m[0.5])
	}
	{ // ToMapWithKeyValue
		m := si.ToMapWithKeyValue(func(i int) interface{} { return i * 100 }, func(i int) interface{} { return i * 2 })
		assert.Equal(4, m[200])
	}
	{ // ConvertToMapWithKeyValue
		m := ConvertToMapWithKeyValue(si, func(i int) int { return i * 100 }, func(i int) int { return i * 2 })
		assert.Equal(6, m[300])
	}
	{ // Where
		actual := si.Where(func(i int) bool { return i%2 == 0 }).ToSlice()
		assert.Equal([]int{0, 2, 4, 6, 8}, actual)
	}
	{ // Take
		actual := si.Take(3).ToSlice()
		assert.Equal([]int{0, 1, 2}, actual)
	}
	{ // Skip
		actual := si.Skip(5).ToSlice()
		assert.Equal([]int{5, 6, 7, 8, 9}, actual)
	}
	{ // TakeWhile
		actual := si.TakeWhile(func(i int) bool { return i < 5 }).ToSlice()
		assert.Equal([]int{0, 1, 2, 3, 4}, actual)
	}
	{ // SkipWhile
		actual := si.SkipWhile(func(i int) bool { return i < 8 }).ToSlice()
		assert.Equal([]int{8, 9}, actual)
	}
	{ // Contains
		actual := si.Contains(3)
		assert.Equal(true, actual)
	}
	{ // Contains
		actual := si.Contains(10)
		assert.Equal(false, actual)
	}
	{ // Any
		actual := si.Any(func(i int) bool { return i > 10 })
		assert.Equal(false, actual)
	}
	{ // Any
		actual := si.Any(func(i int) bool { return i < 2 })
		assert.Equal(true, actual)
	}
	{ // All
		actual := si.All(func(i int) bool { return i < 3 })
		assert.Equal(false, actual)
	}
	{ // All
		actual := si.All(func(i int) bool { return i >= 0 })
		assert.Equal(true, actual)
	}
	{ // TakeLast
		actual := si.TakeLast(3).ToSlice()
		assert.Equal([]int{7, 8, 9}, actual)
	}
	{ // SkipLast
		actual := si.SkipLast(7).ToSlice()
		assert.Equal([]int{0, 1, 2}, actual)
	}
	{ // Count
		actual := si.Count(func(i int) bool { return i%2 == 1 })
		assert.Equal(5, actual)
	}
	{ // Append
		actual := si.Take(2).Append(3).ToSlice()
		assert.Equal([]int{0, 1, 3}, actual)
	}
	{ // Append multiple value
		actual := si.Take(2).Append(3, 5, 7).ToSlice()
		assert.Equal([]int{0, 1, 3, 5, 7}, actual)
	}
	{ // ElementAt
		actual := si.ElementAt(3)
		assert.Equal(3, actual)
	}
	{ // ElementAtOrDefault common case
		actual := si.ElementAtOrDefault(3)
		assert.Equal(3, actual)
	}
	{ // ElementAtOrDefault out of index
		actual := si.ElementAtOrDefault(300)
		assert.Equal(0, actual)
	}
	{ // ElementAtOrDefault invalid index
		actual := si.ElementAtOrDefault(-3)
		assert.Equal(0, actual)
	}
	{ // First
		actual := si.First(func(i int) bool { return i > 2 })
		assert.Equal(3, actual)
	}
	{ // FirstOrDefault
		actual := si.FirstOrDefault(func(i int) bool { return i > 100 })
		assert.Equal(0, actual)
	}
	{ // Last
		actual := si.Last(func(i int) bool { return i < 8 })
		assert.Equal(7, actual)
	}
	{ // LastOrDefault
		actual := si.LastOrDefault(func(i int) bool { return i < 8 })
		assert.Equal(7, actual)
	}
	{ // Prepend
		actual := si.Prepend(999).First(func(i int) bool { return true })
		assert.Equal(999, actual)
	}
	{ // Prepend multiple value
		actual := si.Prepend(999, 888).ToSlice()[:2]
		assert.Equal([]int{999, 888}, actual)
	}
	{ // Reverse
		actual := si.Reverse().ToSlice()
		assert.Equal([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, actual)
	}
	{ // Single
		actual := si.Single(func(i int) bool { return i < 1 })
		assert.Equal(0, actual)
	}
	{ // SingleOeDefault
		actual := si.SingleOrDefault(func(i int) bool { return i > 3 })
		assert.Equal(0, actual)
	}
	{ // ForEach
		si.ForEach(func(i int) { fmt.Println("Foreach test ", i) })
	}
	{ // Remove
		actual := New([]int{1, 2, 3, 4})
		actual2 := actual.Remove(3)
		assert.True(actual2)
		assert.Equal(New([]int{1, 2, 4}), actual)
	}
	{ // RemoveAll
		actual := New([]int{1, 2, 3, 4, 5, 6, 7})
		actual2 := actual.RemoveAll(func(i int) bool { return i%2 == 1 })
		assert.Equal(4, actual2)
		assert.Equal(New([]int{2, 4, 6}), actual)
	}
	{ // RemoveAt
		actual := New([]int{1, 2, 3, 4, 5})
		actual.RemoveAt(3)
		assert.Equal(New([]int{1, 2, 3, 5}), actual)
	}
	{ // RemoveRange
		actual := New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(-1, 3)
		assert.Equal(fmt.Errorf("argument out of range"), err)
	}
	{ // RemoveRange
		actual := New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(1, 50)
		assert.Equal(fmt.Errorf("argument out of range"), err)
	}
	{ // RemoveRange
		actual := New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(2, 2)
		assert.NoError(err)
		assert.Equal(New([]int{1, 2, 5, 6, 7, 8, 9}), actual)
	}
	{ // Distinct
		actual := New([]int{1, 2, 3, 1, 5, 5, 2, 3, 8}).Distinct().ToSlice()
		assert.Equal([]int{1, 2, 3, 5, 8}, actual)
	}
	{ // OrderBy
		si := New([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := si.OrderBy(func(i int) int { return i })
		assert.Equal(New([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), orderedSi)
	}
	{ // OrderByDescending
		si := New([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := si.OrderByDescending(func(i int) int { return i })
		assert.Equal(New([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), orderedSi)
	}
	{ // another Order
		si := New([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := OrderBy(si, func(i int) int { return i })
		assert.Equal(New([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), orderedSi)
	}
	{ // another Order
		si := New([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := OrderBy(si, func(i int) int64 { return int64(i * 20) })
		assert.Equal(New([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), orderedSi)
	}
	{ // another OrderByDescending
		si := New([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := OrderByDescending(si, func(i int) int { return i })
		assert.Equal(New([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), orderedSi)
	}
	{ // Add
		si := New([]int{1, 2, 3})
		si.Add(4)
		assert.Equal(New([]int{1, 2, 3, 4}), si)
	}
	{ // Add Range
		si := New([]int{1, 2, 3})
		si.AddRange([]int{4, 5, 6})
		assert.Equal(si, New([]int{1, 2, 3, 4, 5, 6}))
	}
	{ // Clear
		si := New([]int{1, 2, 3})
		capacity := cap(si.ToSlice())
		si.Clear()
		assert.Equal(si, New(make([]int, capacity)))
	}
	{ // Exists
		si := New([]int{1, 2, 3})
		assert.True(si.Exists(func(i int) bool { return i == 2 }))
		assert.False(si.Exists(func(i int) bool { return i-10 > 0 }))
	}
	{ // Find
		si := New([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(6, si.Find(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindAll
		si := New([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(New([]int{6, 8}), si.FindAll(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindIndex
		si := New([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(3, si.FindIndex(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindLast
		si := New([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(8, si.FindLast(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindLastIndex
		si := New([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(5, si.FindLastIndex(func(i int) bool { return i%2 == 0 }))
	}
	{ // GroupBy common case
		si := New([]int{1, 2, 3, 4, 5})
		keyFunc := func(i int) int { return i % 2 }
		elementFunc := func(i int) string { return strconv.Itoa(i) }
		expected := map[int][]string{
			0: {"2", "4"},
			1: {"1", "3", "5"},
		}

		actual := GroupBy(si, keyFunc, elementFunc)

		assert.Equal(expected, actual)
	}
	{ // GroupBy returns an empty map when the input Linq is empty
		si := New([]int{})
		keyFunc := func(i int) int { return i % 2 }
		elementFunc := func(i int) string { return strconv.Itoa(i) }
		expected := map[int][]string{}

		actual := GroupBy(si, keyFunc, elementFunc)
		assert.Equal(expected, actual)
	}
	{ // Repeat numbers
		element := 5
		count := 3
		actual := Repeat(element, count)
		assert.Equal(New([]int{5, 5, 5}), actual)
	}
	{ // Repeat with negative count
		element := 5
		count := -3
		actual := Repeat(element, count)
		assert.Equal(New([]int{}), actual)
	}
	{ // ReplaceAll
		si := New([]int{1, 2, 3, 4, 5})
		oldValue := 2
		newValue := 10

		actual := si.ReplaceAll(oldValue, newValue)

		expected := New([]int{1, 10, 3, 4, 5})
		assert.Equal(expected, actual)
	}
	{ // Sum
		nsi := NewNumberLinq[int, int](si.ToSlice())
		sum := nsi.Sum(func(i int) int { return i })
		assert.Equal(45, sum)
	}
	{ // Max
		nsi := NewNumberLinq[int, int](si.ToSlice())
		max := nsi.Max(func(i int) int { return i })
		assert.Equal(9, max)
	}
	{ // Min
		nsi := NewNumberLinq[int, int](si.ToSlice())
		min := nsi.Min(func(i int) int { return i })
		assert.Equal(0, min)
	}
}

func Test_Struct_Methods(t *testing.T) {
	assert := assert.New(t)

	type user struct {
		name string
	}

	{ // repeat
		actual := Repeat(user{
			name: "Rockefeller",
		}, 3)
		expected := New([]user{
			{
				name: "Rockefeller",
			},
			{
				name: "Rockefeller",
			},
			{
				name: "Rockefeller",
			},
		})
		assert.Equal(expected, actual)
	}
	{ // other
		users := New([]user{
			{name: "A1"},
			{name: "A2"},
			{name: "B1"},
			{name: "C1"},
			{name: "C2"},
			{name: "D1"},
			{name: "A3"},
			{name: "C3"},
			{name: "A99"},
		})

		actual := users.Skip(5).Where(func(u user) bool { return u.name[0] == 'A' }).Count(func(u user) bool { return true })
		assert.Equal(2, actual)
	}
	{ // Order by string
		ss := New([]user{{name: "abc"}, {name: "apple"}, {name: "a1234567"}, {name: "a"}})
		orderedSs := OrderBy(ss, func(u user) string { return u.name })
		assert.Equal(New([]user{{name: "a"}, {name: "a1234567"}, {name: "abc"}, {name: "apple"}}), orderedSs)
	}
	{ // Order by string length
		ss := New([]user{{name: "abc"}, {name: "apple"}, {name: "a1234567"}, {name: "a"}})
		orderedSs := OrderBy(ss, func(u user) int { return len(u.name) })
		assert.Equal(New([]user{{name: "a"}, {name: "abc"}, {name: "apple"}, {name: "a1234567"}}), orderedSs)
	}
}

func Test_Select(t *testing.T) {
	assert := assert.New(t)
	type user struct {
		name string
		age  int
	}

	users := []user{
		{
			name: "Ann",
			age:  12,
		},
		{
			name: "Jack",
			age:  11,
		},
		{
			name: "Ian",
			age:  15,
		},
	}

	names := Select(New(users), func(u user) string { return u.name })
	assert.Equal([]string{"Ann", "Jack", "Ian"}, names.ToSlice())
}

func Test_GroupBy(t *testing.T) {
	assert := assert.New(t)
	type str struct {
		key   string
		value int
	}
	testData := []str{
		{
			key:   "hello",
			value: 0,
		},
		{
			key:   "hello",
			value: 2,
		},
		{
			key:   "world",
			value: 9,
		},
		{
			key:   "world",
			value: 7,
		},
		{
			key:   "hello",
			value: 4,
		},
		{
			key:   "world",
			value: 5,
		},
	}

	act := GroupBy(New(testData),
		func(s str) string { return s.key },
		func(s str) int { return s.value })

	assert.Equal(map[string][]int{
		"hello": {0, 2, 4},
		"world": {9, 7, 5},
	}, act)
}

func TestSelectMany(t *testing.T) {
	assert := assert.New(t)
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{7, 8, 9}

	selector := func(x int) []int {
		return []int{x * x}
	}

	result := SelectMany(New([][]int{slice1, slice2, slice3}), func(x []int) []int {
		return SelectMany(New(x), selector).ToSlice()
	})

	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81}

	assert.Equal(New(expected), result)
}

// Call the 'NewFromMap' function with an empty map[K]V and assert that the returned linq[T] is empty.
func TestNewFromMap_EmptyMap_ReturnsEmptyLinq(t *testing.T) {
	assert := assert.New(t)
	m := make(map[int]string)
	delegate := func(k int, v string) string {
		return v + strconv.Itoa(k)
	}

	result := NewFromMap(m, delegate)

	assert.Empty(result.ToSlice(), "Expected empty linq")
}

// Call the 'NewFromMap' function with a nil map[K]V and assert that the returned linq[T] is empty.
func Test_new_from_map_with_nil_map(t *testing.T) {
	assert := assert.New(t)

	var m map[int]string
	delegate := func(k int, v string) string {
		return v + strconv.Itoa(k)
	}

	assert.Empty(NewFromMap(m, delegate).ToSlice())
}

// Call the 'NewFromMap' function with a map[K]V containing one key-value pair and assert that the length of the returned linq[T] is 1.
func TestNewFromMap_WithOneKeyValuePair_ReturnsLinqWithLengthOne(t *testing.T) {
	assert := assert.New(t)

	m := map[int]string{1: "value"}
	delegate := func(k int, v string) string {
		return v + strconv.Itoa(k)
	}

	result := NewFromMap(m, delegate)

	assert.Equal(1, result.Length())
}

// Should panic when given a nil delegate function
func TestNewFromMap_NilDelegate_Panics(t *testing.T) {
	assert := assert.New(t)
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	var delegate func(int, string) string

	assert.Panics(func() {
		NewFromMap(m, delegate)
	})
}

// Returns an empty linq[T] when the channel is closed before any value is sent
func Test_empty_linq_when_channel_closed_before_any_value_sent(t *testing.T) {
	c := make(chan int)
	close(c)

	result := NewFromChannel(c)

	assert.Empty(t, result.ToSlice())
}

// Returns a linq[T] with all values sent through the channel
func Test_linq_with_all_values_sent_through_channel(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	result := NewFromChannel(c)

	assert.ElementsMatch(t, []int{1, 2, 3}, result.ToSlice())
}
