package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RunInAsync(t *testing.T) {
	assert := assert.New(t)
	nums := Linq[int]{1, 2, 3, 4, 5}
	double := RunInAsync(nums, func(i int) int { return 2 * i })
	assert.Equal([]int{2, 4, 6, 8, 10}, double)
}

func Test_RunInAsyncWithRoutineLimit(t *testing.T) {
	assert := assert.New(t)
	nums := Linq[int]{1, 2, 3, 4, 5}
	outputs := make(chan int, 5)
	nums.RunInAsyncWithRoutineLimit(func(i int) { outputs <- i * 2 }, 3)
	close(outputs)
	actual := []int{}
	for out := range outputs {
		actual = append(actual, out)
	}
	assert.ElementsMatch([]int{2, 4, 6, 8, 10}, actual)
}
