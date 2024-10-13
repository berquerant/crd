package note_test

import (
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		for _, tc := range []struct {
			s    string
			want note.Note
		}{
			{
				s:    "C",
				want: note.NewNote(note.C, note.Natural),
			},
			{
				s:    "C#",
				want: note.NewNote(note.C, note.Sharp),
			},
			{
				s:    "Db",
				want: note.NewNote(note.D, note.Flat),
			},
		} {
			t.Run(tc.s, func(t *testing.T) {
				got, err := note.ParseNote(tc.s)
				if !assert.Nil(t, err) {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})

	t.Run("AddDegree", func(t *testing.T) {
		for _, tc := range []struct {
			title  string
			n      note.Note
			d      note.Degree
			sharp  bool
			want   note.Note
			octave note.Octave
		}{
			{
				title: "minor7 with flat from Db",
				n:     note.NewNote(note.D, note.Flat),
				d:     note.MustNewDegree(7, note.MinorDegree),
				sharp: false,
				want:  note.NewNote(note.B, note.Natural),
			},
			{
				title: "minor7 with flat from C#",
				n:     note.NewNote(note.C, note.Sharp),
				d:     note.MustNewDegree(7, note.MinorDegree),
				sharp: false,
				want:  note.NewNote(note.B, note.Natural),
			},
			{
				title:  "minor9 with flat",
				n:      note.NewNote(note.C, note.Natural),
				d:      note.MustNewDegree(9, note.MinorDegree),
				sharp:  false,
				want:   note.NewNote(note.D, note.Flat),
				octave: 1,
			},
			{
				title: "minor7 with flat",
				n:     note.NewNote(note.C, note.Natural),
				d:     note.MustNewDegree(7, note.MinorDegree),
				sharp: false,
				want:  note.NewNote(note.B, note.Flat),
			},
			{
				title: "perfect5 from D",
				n:     note.NewNote(note.D, note.Natural),
				d:     note.MustNewDegree(5, note.PerfectDegree),
				want:  note.NewNote(note.A, note.Natural),
			},
			{
				title: "minor2 with sharp",
				n:     note.NewNote(note.C, note.Natural),
				d:     note.MustNewDegree(2, note.MinorDegree),
				sharp: true,
				want:  note.NewNote(note.C, note.Sharp),
			},
			{
				title: "minor2 with flat",
				n:     note.NewNote(note.C, note.Natural),
				d:     note.MustNewDegree(2, note.MinorDegree),
				sharp: false,
				want:  note.NewNote(note.D, note.Flat),
			},
			{
				title: "major2",
				n:     note.NewNote(note.C, note.Natural),
				d:     note.MustNewDegree(2, note.MajorDegree),
				want:  note.NewNote(note.D, note.Natural),
			},
			{
				title: "perfect1",
				n:     note.NewNote(note.C, note.Natural),
				d:     note.MustNewDegree(1, note.PerfectDegree),
				want:  note.NewNote(note.C, note.Natural),
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				got, oct, err := tc.n.AddDegree(tc.d, tc.sharp)
				if !assert.Nil(t, err) {
					return
				}
				assert.Equal(t, tc.want, got)
				assert.Equal(t, tc.octave, oct)
			})
		}
	})
}
