package util_test

import (
	"testing"

	"github.com/berquerant/crd/util"
	"github.com/stretchr/testify/assert"
)

func TestInverseMap(t *testing.T) {
	for _, tc := range []struct {
		d map[string]int
		q map[int]string
	}{
		{
			d: map[string]int{
				"b": 2,
				"a": 1,
			},
			q: map[int]string{
				1: "a",
				2: "b",
			},
		},
		{
			d: map[string]int{
				"a": 1,
			},
			q: map[int]string{
				1: "a",
			},
		},
		{
			d: map[string]int{},
			q: map[int]string{},
		},
	} {
		{
			got, err := util.InverseMap(tc.d)
			assert.Nil(t, err)
			assert.Equal(t, tc.q, got)
		}
		{
			got, err := util.InverseMap(tc.q)
			assert.Nil(t, err)
			assert.Equal(t, tc.d, got)
		}
	}
}
