package note_test

import (
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/stretchr/testify/assert"
)

func TestSemitone(t *testing.T) {
	for _, tc := range []struct {
		s      note.Semitone
		octave note.Octave
		want   note.Semitone
	}{
		{
			s:      -1,
			octave: -1,
			want:   11,
		},
		{
			s:      13,
			octave: 1,
			want:   1,
		},
		{
			s:      12,
			octave: 1,
		},
		{
			s:    7,
			want: 7,
		},
		{
			s:    1,
			want: 1,
		},
		{
			s: 0,
		},
	} {
		gotOctave := tc.s.Octave()
		got := tc.s.WithoutOctave()
		assert.Equal(t, tc.octave, gotOctave)
		assert.Equal(t, tc.want, got)
		assert.Equal(t, tc.s, got+gotOctave.Semitone())
	}
}
