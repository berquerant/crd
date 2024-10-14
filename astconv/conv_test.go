package astconv_test

import (
	"testing"

	"github.com/berquerant/crd/astconv"
	"github.com/berquerant/crd/input"
	"github.com/berquerant/crd/input/ast"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
	"github.com/stretchr/testify/assert"
)

func TestMetaConverter(t *testing.T) {
	newMeta := func(keyValue ...string) *op.Meta {
		d := map[string]string{}
		var i int
		for i < len(keyValue) {
			key := keyValue[i]
			i++
			value := keyValue[i]
			i++
			d[key] = value
		}
		r := op.Meta(d)
		return &r
	}

	for _, tc := range []struct {
		title string
		tree  *ast.ChordMeta
		want  *op.Meta
	}{
		{
			title: "nil",
			tree:  nil,
			want:  nil,
		},
		{
			title: "no data",
			tree:  &ast.ChordMeta{},
			want:  nil,
		},
		{
			title: "one pair",
			tree: &ast.ChordMeta{
				Data: []*ast.ChordMetadata{
					{
						Key: &ast.Token{
							VValue: "k",
						},
						Value: &ast.Token{
							VValue: "v",
						},
					},
				},
			},
			want: newMeta("k", "v"),
		},
		{
			title: "two pairs",
			tree: &ast.ChordMeta{
				Data: []*ast.ChordMetadata{
					{
						Key: &ast.Token{
							VValue: "k",
						},
						Value: &ast.Token{
							VValue: "v",
						},
					},
					{
						Key: &ast.Token{
							VValue: "k2",
						},
						Value: &ast.Token{
							VValue: "v2",
						},
					},
				},
			},
			want: newMeta("k", "v", "k2", "v2"),
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			var c astconv.MetaConverterImpl
			got := c.Convert(tc.tree)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValuesConverter(t *testing.T) {
	for _, tc := range []struct {
		title string
		tree  *ast.ChordValues
		want  []note.Value
	}{
		{
			title: "values",
			tree: &ast.ChordValues{
				Values: []*ast.ChordValue{
					{
						Num: &ast.Token{
							VValue: "1",
						},
					},
					{
						Num: &ast.Token{
							VValue: "1",
						},
						Denom: &ast.Token{
							VValue: "2",
						},
					},
				},
			},
			want: []note.Value{
				note.MustNewValue(1, 1),
				note.MustNewValue(1, 2),
			},
		},
		{
			title: "single",
			tree: &ast.ChordValues{
				Values: []*ast.ChordValue{
					{
						Num: &ast.Token{
							VValue: "1",
						},
					},
				},
			},
			want: []note.Value{
				note.MustNewValue(1, 1),
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			var c astconv.ValuesConverterImpl
			got, err := c.Convert(tc.tree)
			if !assert.Nil(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestChordConverter(t *testing.T) {
	t.Run("Syllable", func(t *testing.T) {
		for _, tc := range []struct {
			title string
			key   op.Key
			tree  *ast.Chord
			want  *input.Chord
		}{
			{
				title: "B# on Cb",
				key:   op.MustParseKey("Cb"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "B",
						},
						Accidental: &ast.Token{
							VValue: "#",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 7,
						Name:  note.DoublyAugmentedDegree,
					},
				},
			},
			{
				title: "B on Cb",
				key:   op.MustParseKey("Cb"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "B",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 7,
						Name:  note.AugmentedDegree,
					},
				},
			},
			{
				title: "E on C#m",
				key:   op.MustParseKey("C#m"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "E",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 3,
						Name:  note.MinorDegree,
					},
				},
			},
			{
				title: "Eb on C",
				key:   op.MustParseKey("C"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "E",
						},
						Accidental: &ast.Token{
							VValue: "b",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 3,
						Name:  note.MinorDegree,
					},
				},
			},
			{
				title: "D# on C",
				key:   op.MustParseKey("C"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "D",
						},
						Accidental: &ast.Token{
							VValue: "#",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 2,
						Name:  note.AugmentedDegree,
					},
				},
			},
			{
				title: "D on B",
				key:   op.MustParseKey("B"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "D",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 3,
						Name:  note.MinorDegree,
					},
				},
			},
			{
				title: "D on C",
				key:   op.MustParseKey("C"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "D",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 2,
						Name:  note.MajorDegree,
					},
				},
			},
			{
				title: "C on C",
				key:   op.MustParseKey("C"),
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "C",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 1,
						Name:  note.PerfectDegree,
					},
				},
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				scale, err := op.NewScale(tc.key)
				if !assert.Nil(t, err) {
					return
				}
				c := astconv.NewSyllableChordConverter(scale)
				got, err := c.Convert(tc.tree)
				if !assert.Nil(t, err, "%#v", err) {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})

	t.Run("Degree", func(t *testing.T) {
		for _, tc := range []struct {
			title string
			tree  *ast.Chord
			want  *input.Chord
		}{
			{
				title: "I/III",
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "1",
						},
					},
					Base: &ast.ChordBase{
						Degree: &ast.ChordDegree{
							Degree: &ast.Token{
								VValue: "3",
							},
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 1,
						Name:  note.PerfectDegree,
					},
					Base: &note.Degree{
						Value: 3,
						Name:  note.MajorDegree,
					},
				},
			},
			{
				title: "IIIb",
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "3",
						},
						Accidental: &ast.Token{
							VValue: "b",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 3,
						Name:  note.MinorDegree,
					},
				},
			},
			{
				title: "IIm",
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "2",
						},
					},
					Symbol: &ast.ChordSymbol{
						Symbol: &ast.Token{
							VValue: "m",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 2,
						Name:  note.MajorDegree,
					},
					Chord: "m",
				},
			},
			{
				title: "I",
				tree: &ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "1",
						},
					},
				},
				want: &input.Chord{
					Degree: note.Degree{
						Value: 1,
						Name:  note.PerfectDegree,
					},
				},
			},
		} {
			t.Run(tc.title, func(t *testing.T) {
				var c astconv.DegreeChordConverter
				got, err := c.Convert(tc.tree)
				if !assert.Nil(t, err) {
					return
				}
				assert.Equal(t, tc.want, got)
			})
		}
	})
}
