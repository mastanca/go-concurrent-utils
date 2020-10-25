package go_concurrent_utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrDone(t *testing.T) {
	t.Run("finish not by done", func(t *testing.T) {
		done := make(chan interface{})
		ints := IntGenerator(done, []int{1, 1, 1, 1}...)
		iface := make(chan interface{})
		go func() {
			defer close(iface)
			for i := range ints {
				iface <- i
			}
		}()
		for val := range OrDone(done, iface) {
			assert.Equal(t, 1, val.(int))
		}
	})

	t.Run("finish by done", func(t *testing.T) {
		done := make(chan interface{})
		ints := IntGenerator(done, []int{1, 1, 1, 1}...)
		iface := make(chan interface{})
		go func() {
			defer close(iface)
			for i := range ints {
				iface <- i
			}
		}()
		close(done)
		finishedByDone := false
	loop:
		for {
			select {
			case <-OrDone(done, iface):
				finishedByDone = true
				break loop
			case <-time.After(5 * time.Second):
				assert.Fail(t, "not finished by done")
				break loop
			}
		}
		assert.True(t, finishedByDone)
	})
}
