package cc_test

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/berquerant/crd/ast"
	"github.com/berquerant/crd/cc"
	"github.com/berquerant/crd/note"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	for _, tc := range []*struct {
		name  string
		input string
		want  *ast.Score
	}{
		{
			name:  "triad",
			input: "C[1]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Chord{
						ChordNote: &ast.ChordNote{
							SPN: note.NewSPN(note.NewNote(note.C, note.Natural), 4),
						},
						ChordOption: &ast.ChordOption{},
						Value:       note.NewValue(big.NewRat(1, 1)),
					},
				},
			},
		},
		{
			name:  "rest",
			input: "R[1/4]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Rest{
						Value: note.NewValue(big.NewRat(1, 4)),
					},
				},
			},
		},
		{
			name:  "triad and rest",
			input: "C#[1/4] R[1/4] Eb[1/2]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Chord{
						ChordNote: &ast.ChordNote{
							SPN: note.NewSPN(note.NewNote(note.C, note.Sharp), 4),
						},
						ChordOption: &ast.ChordOption{},
						Value:       note.NewValue(big.NewRat(1, 4)),
					},
					&ast.Rest{
						Value: note.NewValue(big.NewRat(1, 4)),
					},
					&ast.Chord{
						ChordNote: &ast.ChordNote{
							SPN: note.NewSPN(note.NewNote(note.E, note.Flat), 4),
						},
						ChordOption: &ast.ChordOption{},
						Value:       note.NewValue(big.NewRat(1, 2)),
					},
				},
			},
		},
		{
			name:  "half diminished",
			input: "Fm7-5[1]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Chord{
						ChordNote: &ast.ChordNote{
							SPN: note.NewSPN(note.NewNote(note.F, note.Natural), 4),
						},
						ChordOption: &ast.ChordOption{
							IsMinor:      true,
							IsSeventh:    true,
							Accidentaled: -5,
						},
						Value: note.NewValue(big.NewRat(1, 1)),
					},
				},
			},
		},
		{
			name:  "meter",
			input: "meter[7/8]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Meter{
						Num:   7,
						Denom: 8,
					},
				},
			},
		},
		{
			name:  "tempo",
			input: "tempo[150]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Tempo{
						BPM: 150,
					},
				},
			},
		},
		{
			name:  "key",
			input: "key[Dbminor]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Key{
						Key: note.NewKey(
							note.D,
							note.Flat,
							true,
						),
					},
				},
			},
		},
		{
			name:  "on",
			input: "G7/F[1]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Chord{
						ChordNote: &ast.ChordNote{
							SPN: note.NewSPN(note.NewNote(note.G, note.Natural), 4),
						},
						ChordOption: &ast.ChordOption{
							IsSeventh: true,
						},
						ChordBase: &ast.ChordBase{
							Note: note.NewNote(note.F, note.Natural),
						},
						Value: note.NewValue(big.NewRat(1, 1)),
					},
				},
			},
		},
		{
			name:  "instrument",
			input: `inst["Acoustic Piano"]`,
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Instrument{
						Name: "Acoustic Piano",
					},
				},
			},
		},
		{
			name:  "transposition",
			input: "trans[2]",
			want: &ast.Score{
				NodeList: []ast.Node{
					&ast.Transposition{
						Semitone: 2,
					},
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			l := cc.NewLexer(bytes.NewBufferString(tc.input))
			cc.Parse(l)
			assert.Equal(t, len(l.Result().NodeList), len(tc.want.NodeList))
			for i, g := range l.Result().NodeList {
				assert.True(t, g.Equal(tc.want.NodeList[i]), i, fmt.Sprint(g))
			}
		})
	}
}
