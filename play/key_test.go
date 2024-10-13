package play_test

import (
	"testing"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
	"github.com/berquerant/crd/play"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMapper struct {
	mock.Mock
}

func (m *mockMapper) GetChordAttributes(name string) ([]chord.Attribute, bool) {
	r := m.Called(name)
	return r.Get(0).([]chord.Attribute), r.Get(1).(bool)
}

func (*mockMapper) GetChord(name string) (chord.Chord, bool) {
	// unused
	var c chord.Chord
	return c, false
}

func TestKey(t *testing.T) {
	t.Run("Apply", func(t *testing.T) {
		var (
			cMajor      = op.MustParseKey("C")
			dMajor      = op.MustParseKey("D")
			perfect1, _ = note.ParseDegree("1")
			major2, _   = note.ParseDegree("2")
			major3, _   = note.ParseDegree("3")
			testChord   = chord.Chord{
				Name: "Test",
				Meta: chord.Metadata{
					Display: "",
				},
			}
			// for test
			testChordAttrs = []chord.Attribute{
				{
					Name:   "Major3",
					Degree: major3,
				},
			}
		)

		for _, tc := range []struct {
			title string
			key   op.Key
			c     op.Chord
			attrs []chord.Attribute
			want  []play.MIDINoteNumber
			err   error
		}{
			{
				title: "test chord 1 on C with base 2",
				key:   cMajor,
				c:     op.NewChord(perfect1, testChord, &major2),
				attrs: testChordAttrs,
				want: []play.MIDINoteNumber{
					50, // D3 from base
					64, // E4 from attr
				},
			},
			{
				title: "test chord 1 on D",
				key:   dMajor,
				c:     op.NewChord(perfect1, testChord, nil),
				attrs: testChordAttrs,
				want: []play.MIDINoteNumber{
					50, // D3 from base
					66, // F#4 from attr
				},
			},
			{
				title: "test chord 2 on C",
				key:   cMajor,
				c:     op.NewChord(major2, testChord, nil),
				attrs: testChordAttrs,
				want: []play.MIDINoteNumber{
					50, // D3 from base
					66, // F#4 from attr
				},
			},
			{
				title: "test chord 1 on C",
				key:   cMajor,
				c:     op.NewChord(perfect1, testChord, nil),
				attrs: testChordAttrs,
				want: []play.MIDINoteNumber{
					48, // C3 from base
					64, // E4 from attr
				},
			},
			{
				title: "chord not found",
				key:   cMajor,
				c:     op.NewChord(perfect1, testChord, nil),
				err:   errorx.ErrNotFound,
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				var mapper mockMapper
				mapper.On("GetChordAttributes", tc.c.Chord.Name).Return(tc.attrs, len(tc.attrs) > 0)
				key := play.NewKey(tc.key, &mapper)
				got, err := key.Apply(tc.c)
				if tc.err != nil {
					assert.ErrorIs(t, err, tc.err)
					return
				}
				assert.Nil(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}
