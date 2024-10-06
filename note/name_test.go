package note_test

import (
	"fmt"
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	t.Run("GetDegree", func(t *testing.T) {
		for _, tc := range []struct {
			x, y note.Name
			want uint
		}{
			{
				x:    note.C,
				y:    note.C,
				want: 1,
			},
			{
				x:    note.C,
				y:    note.D,
				want: 2,
			},
			{
				x:    note.C,
				y:    note.B,
				want: 7,
			},
			{
				x:    note.B,
				y:    note.C,
				want: 2,
			},
			{
				x:    note.A,
				y:    note.D,
				want: 4,
			},
			{
				x:    note.D,
				y:    note.A,
				want: 5,
			},
		} {
			t.Run(fmt.Sprintf("%s_%s", tc.x, tc.y), func(t *testing.T) {
				got, ok := tc.x.GetDegree(tc.y)
				assert.True(t, ok)
				assert.Equal(t, tc.want, got)
			})
		}
	})

	t.Run("AddDegree", func(t *testing.T) {
		for _, tc := range []struct {
			title      string
			n          note.Name
			degree     int
			wantName   note.Name
			wantOctave note.Octave
		}{
			{
				title:      "A+10",
				n:          note.A,
				degree:     10,
				wantName:   note.C,
				wantOctave: 2,
			},
			{
				title:      "A+3",
				n:          note.A,
				degree:     3,
				wantName:   note.C,
				wantOctave: 1,
			},
			{
				title:      "D-10",
				n:          note.D,
				degree:     -10,
				wantName:   note.B,
				wantOctave: -2,
			},
			{
				title:      "D-3",
				n:          note.D,
				degree:     -3,
				wantName:   note.B,
				wantOctave: -1,
			},
			{
				title:    "D-2",
				n:        note.D,
				degree:   -2,
				wantName: note.C,
			},
			{
				title:    "D+2",
				n:        note.D,
				degree:   2,
				wantName: note.E,
			},
			{
				title:    "1",
				n:        note.C,
				degree:   1,
				wantName: note.C,
			},
			{
				title:    "-1",
				n:        note.C,
				degree:   -1,
				wantName: note.C,
			},
			{
				title:    "0",
				n:        note.C,
				wantName: note.C,
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				gotName, gotOctave := tc.n.AddDegree(tc.degree)
				assert.Equal(t, tc.wantName, gotName)
				assert.Equal(t, tc.wantOctave, gotOctave)
			})
		}
	})
}
