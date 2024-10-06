package chord_test

import (
	"testing"

	"github.com/berquerant/crd/chord"
	"github.com/stretchr/testify/assert"
)

func TestChord(t *testing.T) {
	t.Run("BasicChords", func(t *testing.T) {
		_ = chord.BasicChords()
	})

	t.Run("ParseChords", func(t *testing.T) {
		for _, tc := range []struct {
			title string
			s     string
			want  []chord.Chord
		}{
			{
				title: "parse",
				s: `- name: Test
  meta:
    display: test
  extends: MajorTriad
  attributes:
    - Major11`,
				want: []chord.Chord{
					{
						Name: "Test",
						Meta: chord.Metadata{
							Display: "test",
						},
						Extends: "MajorTriad",
						Attributes: []string{
							"Major11",
						},
					},
				},
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				got, err := chord.ParseChords([]byte(tc.s))
				if !assert.Nil(t, err) {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})
}
