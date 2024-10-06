package util_test

import (
	"testing"

	"github.com/berquerant/crd/util"
	"github.com/stretchr/testify/assert"
)

func TestOpt(t *testing.T) {
	var (
		value = util.NewOpt(1)

		called      bool
		genCallback = func(want int, msg string) func(int) {
			return func(got int) {
				called = true
				assert.Equal(t, want, got, msg)
			}
		}
		assertCalled = func(want bool, msg string) {
			assert.Equal(t, want, called, msg)
			called = false
		}
	)

	value.WhenUpdated(genCallback(1, "init"))
	assertCalled(true, "init")

	value.WhenUpdated(genCallback(1, "retry"))
	assertCalled(false, "retry")
	assert.Equal(t, 1, value.Unwrap())

	value.Update(2)
	assert.Equal(t, 2, value.Unwrap())
	value.WhenUpdated(genCallback(2, "update"))
	assertCalled(true, "update")
}
