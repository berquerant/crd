package chord_test

import (
	"io"
	"os"
	"testing"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/note"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestAttribute(t *testing.T) {
	t.Run("GenerateAttributes", func(t *testing.T) {
		attrs := chord.GenerateAttributes(20)
		b, err := yaml.Marshal(attrs)
		if !assert.Nil(t, err) {
			return
		}

		var got []chord.Attribute
		if !assert.Nil(t, yaml.Unmarshal(b, &got)) {
			return
		}

		f, err := os.Open("attribute.yml")
		if !assert.Nil(t, err) {
			return
		}
		defer f.Close()
		wb, err := io.ReadAll(f)
		if !assert.Nil(t, err) {
			return
		}

		var want []chord.Attribute
		if !assert.Nil(t, yaml.Unmarshal(wb, &want)) {
			return
		}

		assert.Equal(t, want, got)
	})

	t.Run("BasicAttributes", func(t *testing.T) {
		_ = chord.BasicAttributes()
	})
	t.Run("ParseAttributes", func(t *testing.T) {
		for _, tc := range []struct {
			title string
			s     string
			want  []chord.Attribute
		}{
			{
				title: "parse",
				s: `- name: b
  degree: "4#"
- name: c
  degree: "bb19"`,
				want: []chord.Attribute{
					{
						Name: "b",
						Degree: chord.Degree{
							Value: 4,
							Name:  note.AugmentedDegree,
						},
					},
					{
						Name: "c",
						Degree: chord.Degree{
							Value: 19,
							Name:  note.DiminishedDegree,
						},
					},
				},
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				got, err := chord.ParseAttributes([]byte(tc.s))
				if !assert.Nil(t, err) {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})
	t.Run("Semitone", func(t *testing.T) {
		for _, tc := range []struct {
			a    chord.Attribute
			want chord.Semitone
		}{
			{
				a: chord.Attribute{
					Name: "b",
					Degree: chord.Degree{
						Value: 4,
						Name:  note.AugmentedDegree,
					},
				},
				want: 6,
			},
		} {
			got, ok := tc.a.Semitone()
			if !assert.True(t, ok) {
				continue
			}
			assert.Equal(t, tc.want, got)
		}
	})
}
