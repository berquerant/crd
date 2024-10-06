package util_test

import (
	"testing"

	"github.com/berquerant/crd/util"
	"github.com/stretchr/testify/assert"
)

func TestRing(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		_, err := util.NewRing[int]()
		assert.NotNil(t, err)
	})

	t.Run("one element", func(t *testing.T) {
		v, err := util.NewRing(10)
		assert.Nil(t, err)
		for i := range 3 {
			assert.Equal(t, 10, v.At(-i), "at %d", -i)
			assert.Equal(t, 10, v.At(i), "at %d", i)
		}
	})

	t.Run("two elements", func(t *testing.T) {
		v, err := util.NewRing(10, 11)
		assert.Nil(t, err)
		for _, i := range []int{-2, 0, 2} {
			assert.Equal(t, 10, v.At(i), "should be 10 at %d", i)
		}
		for _, i := range []int{-3, -1, 1, 3} {
			assert.Equal(t, 11, v.At(i), "should be 11 at %d", i)
		}
	})

	t.Run("three elements", func(t *testing.T) {
		v, err := util.NewRing(10, 11, 12)
		assert.Nil(t, err)
		for _, i := range []int{-3, 0, 3} {
			assert.Equal(t, 10, v.At(i), "should be 10 at %d", i)
		}
		for _, i := range []int{-5, -2, 1, 4} {
			assert.Equal(t, 11, v.At(i), "should be 11 at %d", i)
		}
		for _, i := range []int{-4, -1, 2, 5} {
			assert.Equal(t, 12, v.At(i), "should be 12 at %d", i)
		}
	})
}
