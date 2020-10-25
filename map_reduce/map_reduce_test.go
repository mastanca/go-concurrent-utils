package map_reduce

import (
	"sort"
	"testing"

	go_concurrent_utils "github.com/mastanca/go-concurrent-utils"

	"github.com/stretchr/testify/assert"
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
		return a - b
	}
	ints := []int{1, 2, 3, 4}
	result := Reduce(ints, difference)
	assert.Equal(t, -10, result)
}

func TestMapReduce(t *testing.T) {
	strings := []string{"a", "ab", "abc", "abcd"}
	count := func(s string) int {
		return len(s)
	}
	sum := func(a, b int) int {
		return a + b
	}
	result := Reduce(Map(strings, count), sum)
	assert.Equal(t, 10, result)
}

func TestMapStream(t *testing.T) {
	strings := []string{"a", "ab", "abc", "abcd"}
	count := func(s string) int {
		return len(s)
	}

	out := MapStream(strings, count)
	var result []int
	for range strings {
		result = append(result, <-out)
	}
	sort.Ints(result)
	assert.Equal(t, []int{1, 2, 3, 4}, result)
}

func TestReduceStream(t *testing.T) {
	difference := func(a, b int) int {
		return a - b
	}
	ints := []int{1, 2, 3, 4}
	done := make(chan interface{})
	defer close(done)
	result := ReduceStream(go_concurrent_utils.IntGenerator(done, ints...), difference)
	assert.Equal(t, -10, result)
}

func TestMapReduceStream(t *testing.T) {
	strings := []string{"a", "ab", "abc", "abcd"}
	count := func(s string) int {
		return len(s)
	}
	sum := func(a, b int) int {
		return a + b
	}
	result := ReduceStream(MapStream(strings, count), sum)
	assert.Equal(t, 10, result)
}
