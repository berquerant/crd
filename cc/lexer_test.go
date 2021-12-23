package cc_test

import (
	"bytes"
	"testing"

	"github.com/berquerant/crd/cc"
	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	for _, tc := range []*struct {
		name  string
		input string
		want  []cc.Token
	}{
		{
			name:  "triad",
			input: "C",
			want: []cc.Token{
				cc.NewToken(cc.C, "C"),
			},
		},
		{
			name:  "seventh",
			input: "D7",
			want: []cc.Token{
				cc.NewToken(cc.D, "D"),
				cc.NewToken(cc.SEVENTH, "7"),
			},
		},
		{
			name:  "triads",
			input: "D# Fb Gm A#dim Baug",
			want: []cc.Token{
				cc.NewToken(cc.D, "D"),
				cc.NewToken(cc.SHARP, "#"),
				cc.NewToken(cc.F, "F"),
				cc.NewToken(cc.FLAT, "b"),
				cc.NewToken(cc.G, "G"),
				cc.NewToken(cc.MINOR, "m"),
				cc.NewToken(cc.A, "A"),
				cc.NewToken(cc.SHARP, "#"),
				cc.NewToken(cc.DIMINISHED, "dim"),
				cc.NewToken(cc.B, "B"),
				cc.NewToken(cc.AUGMENTED, "aug"),
			},
		},
		{
			name:  "accidental",
			input: "C#dim-5[1/4]",
			want: []cc.Token{
				cc.NewToken(cc.C, "C"),
				cc.NewToken(cc.SHARP, "#"),
				cc.NewToken(cc.DIMINISHED, "dim"),
				cc.NewToken(cc.MINUS, "-"),
				cc.NewToken(cc.INT, "5"),
				cc.NewToken(cc.LBRA, "["),
				cc.NewToken(cc.INT, "1"),
				cc.NewToken(cc.SLASH, "/"),
				cc.NewToken(cc.INT, "4"),
				cc.NewToken(cc.RBRA, "]"),
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			l := cc.NewLexer(bytes.NewBufferString(tc.input))
			got := []cc.Token{}
			for {
				tok := l.Scan()
				if tok == cc.EOF {
					break
				}
				got = append(got, cc.NewToken(tok, l.Buffer()))
				l.ResetBuffer()
			}
			if err := l.Err(); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, len(tc.want), len(got))
			for i, w := range tc.want {
				g := got[i]
				assert.Equal(t, w.Type(), g.Type())
				assert.Equal(t, w.Value(), g.Value())
			}
		})
	}
}
