package go_concurrent_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntGenerator(t *testing.T) {
	t.Run("1 int", func(t *testing.T) {
		ints := []int{1}
		intsChan := IntGenerator(nil, ints...)
		expectedResult := make(chan int)
		go func() {
			defer close(expectedResult)
			expectedResult <- 1
		}()
		assert.Equal(t, <-expectedResult, <-intsChan)
	})
	t.Run("many ints", func(t *testing.T) {
		ints := []int{1, 2, 3, 4, 5}
		intsChan := IntGenerator(nil, ints...)
		expectedResult := make(chan int)
		go func() {
			defer close(expectedResult)
			for i := 1; i < 6; i++ {
				expectedResult <- i
			}
		}()
		assert.Equal(t, <-expectedResult, <-intsChan)
	})
}
