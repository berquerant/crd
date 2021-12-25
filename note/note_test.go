package note_test

import (
	"fmt"
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/stretchr/testify/assert"
)

func TestSPN(t *testing.T) {
	t.Run("Bound", func(t *testing.T) {
		for _, tc := range []*struct {
			name   string
			origin note.SPN  // self
			target note.Note // the note we need
			upper  note.SPN  // the upper SPN with the target
			lower  note.SPN  // the lower SPN with the target
		}{
			{
				name:   "name of origin equals target",
				origin: note.NewSPN(note.NewNote(note.C, note.Natural), 4),
				target: note.NewNote(note.C, note.Natural),
				upper:  note.NewSPN(note.NewNote(note.C, note.Natural), 5),
				lower:  note.NewSPN(note.NewNote(note.C, note.Natural), 3),
			},
			{
				name:   "name of origin higher than target",
				origin: note.NewSPN(note.NewNote(note.D, note.Natural), 4),
				target: note.NewNote(note.C, note.Natural),
				upper:  note.NewSPN(note.NewNote(note.C, note.Natural), 5),
				lower:  note.NewSPN(note.NewNote(note.C, note.Natural), 4),
			},
			{
				name:   "name of origin lower than target",
				origin: note.NewSPN(note.NewNote(note.C, note.Natural), 4),
				target: note.NewNote(note.D, note.Natural),
				upper:  note.NewSPN(note.NewNote(note.D, note.Natural), 4),
				lower:  note.NewSPN(note.NewNote(note.D, note.Natural), 3),
			},
		} {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				gotUpper := tc.origin.UpperBound(tc.target)
				gotLower := tc.origin.LowerBound(tc.target)
				assert.True(t, tc.upper.Equal(gotUpper), fmt.Sprintf("upper %s", gotUpper))
				assert.True(t, tc.lower.Equal(gotLower), fmt.Sprintf("lower %s", gotLower))
			})
		}
	})
}
