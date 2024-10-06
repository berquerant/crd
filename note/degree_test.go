package note_test

import (
	"testing"

	"github.com/berquerant/crd/note"
	"github.com/stretchr/testify/assert"
)

func TestDegree(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		for _, tc := range []struct {
			s    string
			want note.Degree
		}{
			{
				s: "5##",
				want: note.Degree{
					Name:  note.DoublyAugmentedDegree,
					Value: 5,
				},
			},
			{
				s: "4bbb",
				want: note.Degree{
					Name:  note.DoublyDiminishedDegree,
					Value: 4,
				},
			},
			{
				s: "5#",
				want: note.Degree{
					Name:  note.AugmentedDegree,
					Value: 5,
				},
			},
			{
				s: "4bb",
				want: note.Degree{
					Name:  note.DiminishedDegree,
					Value: 4,
				},
			},
			{
				s: "2b",
				want: note.Degree{
					Name:  note.MinorDegree,
					Value: 2,
				},
			},
			{
				s: "2",
				want: note.Degree{
					Name:  note.MajorDegree,
					Value: 2,
				},
			},
			{
				s: "1",
				want: note.Degree{
					Name:  note.PerfectDegree,
					Value: 1,
				},
			},
		} {
			t.Run(tc.s, func(t *testing.T) {
				got, err := note.ParseDegree(tc.s)
				assert.Nil(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})

	t.Run("Semitone", func(t *testing.T) {
		for _, tc := range []struct {
			title string
			d     note.Degree
			want  note.Semitone
		}{
			{
				title: "doublyaugmented7",
				d: note.Degree{
					Name:  note.DoublyAugmentedDegree,
					Value: 7,
				},
				want: 13,
			},
			{
				title: "doublydiminished7",
				d: note.Degree{
					Name:  note.DoublyDiminishedDegree,
					Value: 7,
				},
				want: 8,
			},
			{
				title: "diminished7",
				d: note.Degree{
					Name:  note.DiminishedDegree,
					Value: 7,
				},
				want: 9,
			},
			{
				title: "minor7",
				d: note.Degree{
					Name:  note.MinorDegree,
					Value: 7,
				},
				want: 10,
			},
			{
				title: "augmented7",
				d: note.Degree{
					Name:  note.AugmentedDegree,
					Value: 7,
				},
				want: 12,
			},
			{
				title: "major7",
				d: note.Degree{
					Name:  note.MajorDegree,
					Value: 7,
				},
				want: 11,
			},
			{
				title: "diminished4",
				d: note.Degree{
					Name:  note.DiminishedDegree,
					Value: 4,
				},
				want: 4,
			},
			{
				title: "augmented4",
				d: note.Degree{
					Name:  note.AugmentedDegree,
					Value: 4,
				},
				want: 6,
			},
			{
				title: "perfect4",
				d: note.Degree{
					Name:  note.PerfectDegree,
					Value: 4,
				},
				want: 5,
			},
			{
				title: "perfect1",
				d: note.Degree{
					Name:  note.PerfectDegree,
					Value: 1,
				},
				want: 0,
			},
			{
				title: "perfect11",
				d: note.Degree{
					Name:  note.PerfectDegree,
					Value: 11,
				},
				want: 17,
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				got, ok := tc.d.Semitone()
				if !assert.True(t, ok) {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})

	t.Run("Coerce", func(t *testing.T) {
		for _, tc := range []struct {
			title string
			name  note.CoerceDegreeName
			value uint
			want  note.Degree
			notOK bool
		}{
			{
				title: "doublyaugmented5",
				name:  note.DoublyAugmentedCoerceDegree,
				value: 5,
				want: note.Degree{
					Name:  note.DoublyAugmentedDegree,
					Value: 5,
				},
			},
			{
				title: "doublydiminished7",
				name:  note.DoublyDiminishedCoerceDegree,
				value: 7,
				want: note.Degree{
					Name:  note.DoublyDiminishedDegree,
					Value: 7,
				},
			},
			{
				title: "diminished7",
				name:  note.DiminishedCoerceDegree,
				value: 7,
				want: note.Degree{
					Name:  note.DiminishedDegree,
					Value: 7,
				},
			},
			{
				title: "perfect11",
				name:  note.MajorOrPerfectCoerceDegree,
				value: 11,
				want: note.Degree{
					Name:  note.PerfectDegree,
					Value: 11,
				},
			},
			{
				title: "augmented5",
				name:  note.AugmentedCoerceDegree,
				value: 5,
				want: note.Degree{
					Name:  note.AugmentedDegree,
					Value: 5,
				},
			},
			{
				title: "augmented4",
				name:  note.AugmentedCoerceDegree,
				value: 4,
				want: note.Degree{
					Name:  note.AugmentedDegree,
					Value: 4,
				},
			},
			{
				title: "diminished4",
				name:  note.MinorOrDiminishedCoerceDegree,
				value: 4,
				want: note.Degree{
					Name:  note.DiminishedDegree,
					Value: 4,
				},
			},
			{
				title: "diminished5",
				name:  note.MinorOrDiminishedCoerceDegree,
				value: 5,
				want: note.Degree{
					Name:  note.DiminishedDegree,
					Value: 5,
				},
			},
			{
				title: "minor3",
				name:  note.MinorOrDiminishedCoerceDegree,
				value: 3,
				want: note.Degree{
					Name:  note.MinorDegree,
					Value: 3,
				},
			},
			{
				title: "major2",
				name:  note.MajorOrPerfectCoerceDegree,
				value: 2,
				want: note.Degree{
					Name:  note.MajorDegree,
					Value: 2,
				},
			},
			{
				title: "perfect1",
				name:  note.MajorOrPerfectCoerceDegree,
				value: 1,
				want: note.Degree{
					Name:  note.PerfectDegree,
					Value: 1,
				},
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				got, ok := tc.name.Degree(tc.value)
				assert.Equal(t, tc.notOK, !ok)
				if tc.notOK {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})
}
