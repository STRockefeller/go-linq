package linq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Int_Methods(t *testing.T) {
	assert := assert.New(t)
	si := NewLinq([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	{ // case ToSlice
		assert.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, si.ToSlice())
	}
	{ // case ToChannel
		ch := si.ToChannel()
		var elements []int
		for element := range ch {
			elements = append(elements, element)
		}
		assert.ElementsMatch(si, elements)
	}
	{ // case ToChannelWithBuffer
		ch := si.ToChannelWithBuffer(2)
		var elements []int
		for element := range ch {
			elements = append(elements, element)
		}
		assert.ElementsMatch(si, elements)
	}
	{ // case ToMapWithKey
		m := si.ToMapWithKey(func(i int) interface{} { return i * 10 })
		assert.Equal(5, m[50])
	}
	{ // case ConvertToMapWithKey
		m := ConvertToMapWithKey(si, func(i int) float32 { return float32(i) * 0.1 })
		assert.Equal(5, m[0.5])
	}
	{ // case ToMapWithKeyValue
		m := si.ToMapWithKeyValue(func(i int) interface{} { return i * 100 }, func(i int) interface{} { return i * 2 })
		assert.Equal(4, m[200])
	}
	{ // case ConvertToMapWithKeyValue
		m := ConvertToMapWithKeyValue(si, func(i int) int { return i * 100 }, func(i int) int { return i * 2 })
		assert.Equal(6, m[300])
	}
	{ // case Where
		actual := si.Where(func(i int) bool { return i%2 == 0 }).ToSlice()
		assert.Equal([]int{0, 2, 4, 6, 8}, actual)
	}
	{ // case Take
		actual := si.Take(3).ToSlice()
		assert.Equal([]int{0, 1, 2}, actual)
	}
	{ // case Skip
		actual := si.Skip(5).ToSlice()
		assert.Equal([]int{5, 6, 7, 8, 9}, actual)
	}
	{ // case TakeWhile
		actual := si.TakeWhile(func(i int) bool { return i < 5 }).ToSlice()
		assert.Equal([]int{0, 1, 2, 3, 4}, actual)
	}
	{ // case SkipWhile
		actual := si.SkipWhile(func(i int) bool { return i < 8 }).ToSlice()
		assert.Equal([]int{8, 9}, actual)
	}
	{ // case Contains
		actual := si.Contains(3)
		assert.Equal(true, actual)
	}
	{ // case Contains
		actual := si.Contains(10)
		assert.Equal(false, actual)
	}
	{ // case Any
		actual := si.Any(func(i int) bool { return i > 10 })
		assert.Equal(false, actual)
	}
	{ // case Any
		actual := si.Any(func(i int) bool { return i < 2 })
		assert.Equal(true, actual)
	}
	{ // case All
		actual := si.All(func(i int) bool { return i < 3 })
		assert.Equal(false, actual)
	}
	{ // case All
		actual := si.All(func(i int) bool { return i >= 0 })
		assert.Equal(true, actual)
	}
	{ // case TakeLast
		actual := si.TakeLast(3).ToSlice()
		assert.Equal([]int{7, 8, 9}, actual)
	}
	{ // case SkipLast
		actual := si.SkipLast(7).ToSlice()
		assert.Equal([]int{0, 1, 2}, actual)
	}
	{ // case Count
		actual := si.Count(func(i int) bool { return i%2 == 1 })
		assert.Equal(5, actual)
	}
	{ // case Append
		actual := si.Take(2).Append(3).ToSlice()
		assert.Equal([]int{0, 1, 3}, actual)
	}
	{ // case Append multiple value
		actual := si.Take(2).Append(3, 5, 7).ToSlice()
		assert.Equal([]int{0, 1, 3, 5, 7}, actual)
	}
	{ // case ElementAt
		actual := si.ElementAt(3)
		assert.Equal(3, actual)
	}
	{ // case ElementAtOrDefault common case
		actual := si.ElementAtOrDefault(3)
		assert.Equal(3, actual)
	}
	{ // case ElementAtOrDefault out of index
		actual := si.ElementAtOrDefault(300)
		assert.Equal(0, actual)
	}
	{ // case ElementAtOrDefault invalid index
		actual := si.ElementAtOrDefault(-3)
		assert.Equal(0, actual)
	}
	{ // case First
		actual := si.First(func(i int) bool { return i > 2 })
		assert.Equal(3, actual)
	}
	{ // case FirstOrDefault
		actual := si.FirstOrDefault(func(i int) bool { return i > 100 })
		assert.Equal(0, actual)
	}
	{ // case Last
		actual := si.Last(func(i int) bool { return i < 8 })
		assert.Equal(7, actual)
	}
	{ // case LastOrDefault
		actual := si.LastOrDefault(func(i int) bool { return i < 8 })
		assert.Equal(7, actual)
	}
	{ // case Prepend
		actual := si.Prepend(999).First(func(i int) bool { return true })
		assert.Equal(999, actual)
	}
	{ // case Prepend multiple value
		actual := si.Prepend(999, 888)[:2]
		assert.Equal(NewLinq([]int{999, 888}), actual)
	}
	{ // case Reverse
		actual := si.Reverse().ToSlice()
		assert.Equal([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, actual)
	}
	{ // case Single
		actual := si.Single(func(i int) bool { return i < 1 })
		assert.Equal(0, actual)
	}
	{ // case SingleOeDefault
		actual := si.SingleOrDefault(func(i int) bool { return i > 3 })
		assert.Equal(0, actual)
	}
	{ // case ForEach
		si.ForEach(func(i int) { fmt.Println("Foreach test ", i) })
	}
	{ // case Remove
		actual := NewLinq([]int{1, 2, 3, 4})
		actual2 := actual.Remove(3)
		assert.True(actual2)
		assert.Equal(NewLinq([]int{1, 2, 4}), actual)
	}
	{ // case RemoveAll
		actual := NewLinq([]int{1, 2, 3, 4, 5, 6, 7})
		actual2 := actual.RemoveAll(func(i int) bool { return i%2 == 1 })
		assert.Equal(4, actual2)
		assert.Equal(NewLinq([]int{2, 4, 6}), actual)
	}
	{ // case RemoveAt
		actual := NewLinq([]int{1, 2, 3, 4, 5})
		actual.RemoveAt(3)
		assert.Equal(NewLinq([]int{1, 2, 3, 5}), actual)
	}
	{ // case RemoveRange
		actual := NewLinq([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(-1, 3)
		assert.Equal(fmt.Errorf("argument out of range"), err)
	}
	{ // case RemoveRange
		actual := NewLinq([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(1, 50)
		assert.Equal(fmt.Errorf("argument out of range"), err)
	}
	{ // case RemoveRange
		actual := NewLinq([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(2, 2)
		assert.NoError(err)
		assert.Equal(NewLinq([]int{1, 2, 5, 6, 7, 8, 9}), actual)
	}
	{ // case Distinct
		actual := NewLinq([]int{1, 2, 3, 1, 5, 5, 2, 3, 8}).Distinct().ToSlice()
		assert.Equal([]int{1, 2, 3, 5, 8}, actual)
	}
	{ // OrderBy
		si := NewLinq([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := si.OrderBy(func(i int) int { return i })
		assert.Equal(NewLinq([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), orderedSi)
	}
	{ // OrderByDescending
		si := NewLinq([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := si.OrderByDescending(func(i int) int { return i })
		assert.Equal(NewLinq([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), orderedSi)
	}
	{ // another Order
		si := NewLinq([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := OrderBy(si, func(i int) int { return i })
		assert.Equal(NewLinq([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), orderedSi)
	}
	{ // another Order
		si := NewLinq([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := OrderBy(si, func(i int) int64 { return int64(i * 20) })
		assert.Equal(NewLinq([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), orderedSi)
	}
	{ // another OrderByDescending
		si := NewLinq([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := OrderByDescending(si, func(i int) int { return i })
		assert.Equal(NewLinq([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), orderedSi)
	}
	{ // Add
		si := NewLinq([]int{1, 2, 3})
		si.Add(4)
		assert.Equal(NewLinq([]int{1, 2, 3, 4}), si)
	}
	{ // Add Range
		si := NewLinq([]int{1, 2, 3})
		si.AddRange([]int{4, 5, 6})
		assert.Equal(si, NewLinq([]int{1, 2, 3, 4, 5, 6}))
	}
	{ // Clear
		si := NewLinq([]int{1, 2, 3})
		capacity := cap(si.ToSlice())
		si.Clear()
		assert.Equal(si, NewLinq(make([]int, capacity)))
	}
	{ // Exists
		si := NewLinq([]int{1, 2, 3})
		assert.True(si.Exists(func(i int) bool { return i == 2 }))
		assert.False(si.Exists(func(i int) bool { return i-10 > 0 }))
	}
	{ // Find
		si := NewLinq([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(6, si.Find(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindAll
		si := NewLinq([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(NewLinq([]int{6, 8}), si.FindAll(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindIndex
		si := NewLinq([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(3, si.FindIndex(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindLast
		si := NewLinq([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(8, si.FindLast(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindLastIndex
		si := NewLinq([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(5, si.FindLastIndex(func(i int) bool { return i%2 == 0 }))
	}
	{ // Sum
		nsi := NewNumberLinq[int, int](si)
		sum := nsi.Sum(func(i int) int { return i })
		assert.Equal(45, sum)
	}
	{ // Max
		nsi := NewNumberLinq[int, int](si)
		max := nsi.Max(func(i int) int { return i })
		assert.Equal(9, max)
	}
	{ // Min
		nsi := NewNumberLinq[int, int](si)
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
		expected := Linq[user]{
			user{
				name: "Rockefeller",
			},
			user{
				name: "Rockefeller",
			},
			user{
				name: "Rockefeller",
			},
		}
		assert.Equal(expected, actual)
	}
	{ // other
		users := Linq[user]{
			{name: "A1"},
			{name: "A2"},
			{name: "B1"},
			{name: "C1"},
			{name: "C2"},
			{name: "D1"},
			{name: "A3"},
			{name: "C3"},
			{name: "A99"},
		}

		actual := users.Skip(5).Where(func(u user) bool { return u.name[0] == 'A' }).Count(func(u user) bool { return true })
		assert.Equal(2, actual)
	}
	{ // Order by string
		ss := NewLinq([]user{{name: "abc"}, {name: "apple"}, {name: "a1234567"}, {name: "a"}})
		orderedSs := OrderBy(ss, func(u user) string { return u.name })
		assert.Equal(NewLinq([]user{{name: "a"}, {name: "a1234567"}, {name: "abc"}, {name: "apple"}}), orderedSs)
	}
	{ // Order by string length
		ss := NewLinq([]user{{name: "abc"}, {name: "apple"}, {name: "a1234567"}, {name: "a"}})
		orderedSs := OrderBy(ss, func(u user) int { return len(u.name) })
		assert.Equal(NewLinq([]user{{name: "a"}, {name: "abc"}, {name: "apple"}, {name: "a1234567"}}), orderedSs)
	}
}

func Test_Select(t *testing.T) {
	assert := assert.New(t)
	type user struct {
		name string
		age  int
	}

	users := Linq[user]{
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

	names := Select(users, func(u user) string { return u.name })
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

	act := GroupBy(NewLinq(testData),
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

	result := SelectMany([][]int{slice1, slice2, slice3}, func(x []int) []int {
		return SelectMany(x, selector)
	})

	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81}

	assert.Equal(NewLinq(expected), result)
}
