package util_test

import (
	"strings"
	"testing"

	"github.com/berquerant/crd/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestRat(t *testing.T) {
	for _, tc := range []struct {
		s string
		r util.Rat
	}{
		{
			s: `6/4`,
			r: util.NewRat(6, 4),
		},
		{
			s: `3/4`,
			r: util.NewRat(3, 4),
		},
		{
			s: `1/4`,
			r: util.NewRat(1, 4),
		},
		{
			s: `"2"`,
			r: util.NewRat(2, 1),
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			a, err := yaml.Marshal(tc.r)
			assert.Nil(t, err)
			assert.Equal(t, tc.s, strings.Trim(string(a), "\n"))

			var r util.Rat
			assert.Nil(t, yaml.Unmarshal([]byte(tc.s), &r))
			assert.Equal(t, tc.r, r)
		})
	}
}
