package input_test

import (
	"testing"

	"github.com/berquerant/crd/input"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
	"github.com/berquerant/crd/util"
	"github.com/stretchr/testify/assert"
)

func TestChordMetaTextModifier(t *testing.T) {
	for _, tc := range []struct {
		title    string
		instance *input.Instance
		want     *op.Meta
	}{
		{
			title: "Imaj7/V",
			instance: &input.Instance{
				Chord: &input.Chord{
					Degree: note.MustNewDegree(1, note.PerfectDegree),
					Chord:  "maj7",
					Base:   util.Ptr(note.MustNewDegree(5, note.PerfectDegree)),
				},
			},
			want: op.NewMeta(input.MetaTextKey, "1maj7/5"),
		},
		{
			title: "Vmaj7",
			instance: &input.Instance{
				Chord: &input.Chord{
					Degree: note.MustNewDegree(5, note.PerfectDegree),
					Chord:  "maj7",
				},
			},
			want: op.NewMeta(input.MetaTextKey, "5maj7"),
		},
		{
			title: "I",
			instance: &input.Instance{
				Chord: &input.Chord{
					Degree: note.MustNewDegree(1, note.PerfectDegree),
				},
			},
			want: op.NewMeta(input.MetaTextKey, "1"),
		},
		{
			title:    "no chord",
			instance: &input.Instance{},
			want:     nil,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			m := input.NewChordMetaTextMofidier("", "/")
			if !assert.Nil(t, m.Modify(tc.instance)) {
				return
			}
			assert.Equal(t, tc.want, tc.instance.Meta)
		})
	}
}
