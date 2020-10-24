package map_reduce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	strings := []string{"a", "ab", "abc", "abcd"}
	count := func(s string) int {
		return len(s)
	}

	result := Map(strings, count)
	assert.EqualValues(t, []int{1, 2, 3, 4}, result)
}

func TestReduce(t *testing.T) {
	difference := func(a, b int) int {
		return a-b
	}
	ints := []int{1, 2, 3, 4}
	result := Reduce(ints, difference)
	assert.Equal(t, -10, result)
}

func TestMapReduce(t *testing.T)  {
	strings := []string{"a", "ab", "abc", "abcd"}
	count := func(s string) int {
		return len(s)
	}
	sum := func(a, b int) int {
		return a+b
	}
	result := Reduce(Map(strings, count), sum)
	assert.Equal(t, 10, result)
}
