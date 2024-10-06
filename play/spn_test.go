package play_test

import (
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/play"
	"github.com/stretchr/testify/assert"
)

func TestSPN(t *testing.T) {
	t.Run("MIDINoteNumber", func(t *testing.T) {
		for _, tc := range []struct {
			spn play.SPN
			n   play.MIDINoteNumber
		}{
			{
				spn: play.NewSPN(note.F, 4, note.Flat),
				n:   64,
			},
			{
				spn: play.NewSPN(note.C, 4, note.Sharp),
				n:   61,
			},
			{
				spn: play.NewSPN(note.D, 4, note.Natural),
				n:   62,
			},
			{
				spn: play.NewSPN(note.C, 4, note.Natural),
				n:   60,
			},
		} {
			t.Run(tc.spn.String(), func(t *testing.T) {
				assert.Equal(t, tc.n, tc.spn.MIDINoteNumber())
			})
		}
	})
	t.Run("String", func(t *testing.T) {
		for _, tc := range []struct {
			spn  play.SPN
			want string
		}{
			{
				spn:  play.NewSPN(note.D, 4, note.Flat),
				want: "D4♭",
			},
			{
				spn:  play.NewSPN(note.C, 4, note.Sharp),
				want: "C4♯",
			},
			{
				spn:  play.NewSPN(note.D, 4, note.Natural),
				want: "D4",
			},
			{
				spn:  play.NewSPN(note.C, 4, note.Natural),
				want: "C4",
			},
		} {
			t.Run(tc.want, func(t *testing.T) {
				assert.Equal(t, tc.want, tc.spn.String())
			})
		}
	})
}
