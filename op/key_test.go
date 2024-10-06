package op_test

import (
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
	"github.com/stretchr/testify/assert"
)

func TestAccidental(t *testing.T) {
	t.Run("Tendency", func(t *testing.T) {
		for _, tc := range []struct {
			x, y, want op.Accidental
		}{
			{
				x:    op.UnknownAccidental,
				y:    op.UnknownAccidental,
				want: op.UnknownAccidental,
			},
			{
				x:    op.UnknownAccidental,
				y:    op.Natural,
				want: op.UnknownAccidental,
			},
			{
				x:    op.UnknownAccidental,
				y:    op.Sharp,
				want: op.UnknownAccidental,
			},
			{
				x:    op.UnknownAccidental,
				y:    op.Flat,
				want: op.UnknownAccidental,
			},
			{
				x:    op.Natural,
				y:    op.UnknownAccidental,
				want: op.UnknownAccidental,
			},
			{
				x:    op.Natural,
				y:    op.Natural,
				want: op.Natural,
			},
			{
				x:    op.Natural,
				y:    op.Sharp,
				want: op.Sharp,
			},
			{
				x:    op.Natural,
				y:    op.Flat,
				want: op.Flat,
			},
			{
				x:    op.Sharp,
				y:    op.UnknownAccidental,
				want: op.UnknownAccidental,
			},
			{
				x:    op.Sharp,
				y:    op.Natural,
				want: op.Sharp,
			},
			{
				x:    op.Sharp,
				y:    op.Sharp,
				want: op.Natural,
			},
			{
				x:    op.Sharp,
				y:    op.Flat,
				want: op.Sharp,
			},
			{
				x:    op.Flat,
				y:    op.UnknownAccidental,
				want: op.UnknownAccidental,
			},
			{
				x:    op.Flat,
				y:    op.Natural,
				want: op.Flat,
			},
			{
				x:    op.Flat,
				y:    op.Sharp,
				want: op.Flat,
			},
			{
				x:    op.Flat,
				y:    op.Flat,
				want: op.Natural,
			},
		} {
			assert.Equal(t, tc.want, tc.x.Tendency(tc.y), "%v, %v", tc.x, tc.y)
		}
	})
}

func TestKey(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		for _, tc := range []struct {
			s    string
			want op.Key
		}{
			{
				s: "Bbm",
				want: op.Key{
					Name:       note.B,
					Accidental: op.Flat,
					Minor:      true,
				},
			},
			{
				s: "Am",
				want: op.Key{
					Name:       note.A,
					Minor:      true,
					Accidental: op.Natural,
				},
			},
			{
				s: "Bb",
				want: op.Key{
					Name:       note.B,
					Accidental: op.Flat,
				},
			},
			{
				s: "C",
				want: op.Key{
					Name:       note.C,
					Accidental: op.Natural,
				},
			},
		} {
			t.Run(tc.s, func(t *testing.T) {
				got, err := op.ParseKey(tc.s)
				if !assert.Nil(t, err) {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})
}
