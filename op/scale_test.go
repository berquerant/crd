package op_test

import (
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
	"github.com/stretchr/testify/assert"
)

func TestScale(t *testing.T) {
	for _, tc := range []struct {
		key  op.Key
		want *op.Scale
	}{
		{
			key: op.MustParseKey("C"),
			want: &op.Scale{
				Key: op.MustParseKey("C"),
				Notes: [7]*op.ScaleNote{
					{
						Name:       note.C,
						Accidental: op.Natural,
					},
					{
						Name:       note.D,
						Accidental: op.Natural,
					},
					{
						Name:       note.E,
						Accidental: op.Natural,
					},
					{
						Name:       note.F,
						Accidental: op.Natural,
					},
					{
						Name:       note.G,
						Accidental: op.Natural,
					},
					{
						Name:       note.A,
						Accidental: op.Natural,
					},
					{
						Name:       note.B,
						Accidental: op.Natural,
					},
				},
			},
		},
		{
			key: op.MustParseKey("A"),
			want: &op.Scale{
				Key:   op.MustParseKey("A"),
				Sharp: 3,
				Notes: [7]*op.ScaleNote{
					{
						Name:       note.A,
						Accidental: op.Natural,
					},
					{
						Name:       note.B,
						Accidental: op.Natural,
					},
					{
						Name:       note.C,
						Accidental: op.Sharp,
					},
					{
						Name:       note.D,
						Accidental: op.Natural,
					},
					{
						Name:       note.E,
						Accidental: op.Natural,
					},
					{
						Name:       note.F,
						Accidental: op.Sharp,
					},
					{
						Name:       note.G,
						Accidental: op.Sharp,
					},
				},
			},
		},
		{
			key: op.MustParseKey("Gm"),
			want: &op.Scale{
				Key:  op.MustParseKey("Gm"),
				Flat: 2,
				Notes: [7]*op.ScaleNote{
					{
						Name:       note.G,
						Accidental: op.Natural,
					},
					{
						Name:       note.A,
						Accidental: op.Natural,
					},
					{
						Name:       note.B,
						Accidental: op.Flat,
					},
					{
						Name:       note.C,
						Accidental: op.Natural,
					},
					{
						Name:       note.D,
						Accidental: op.Natural,
					},
					{
						Name:       note.E,
						Accidental: op.Flat,
					},
					{
						Name:       note.F,
						Accidental: op.Natural,
					},
				},
			},
		},
	} {
		t.Run(tc.key.String(), func(t *testing.T) {
			got, err := op.NewScale(tc.key)
			if !assert.Nil(t, err) {
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
