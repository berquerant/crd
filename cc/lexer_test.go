package cc_test

import (
	"bytes"
	"testing"

	"github.com/berquerant/crd/cc"
	"github.com/stretchr/testify/assert"
)

func newTokens(v ...interface{}) []cc.Token {
	var r []cc.Token
	for i := 0; i < len(v); i++ {
		t := v[i].(int)
		i++
		v := v[i].(string)
		r = append(r, cc.NewToken(t, v))
	}
	return r
}

func TestLexer(t *testing.T) {
	for _, tc := range []*struct {
		name  string
		input string
		want  []cc.Token
	}{
		{
			name:  "triad",
			input: "C",
			want:  newTokens(cc.C, "C"),
		},
		{
			name:  "seventh",
			input: "D7",
			want: newTokens(
				cc.D, "D",
				cc.SEVENTH, "7",
			),
		},
		{
			name:  "triads",
			input: "D# Fb Gm A#dim Baug",
			want: newTokens(
				cc.D, "D",
				cc.SHARP, "#",
				cc.F, "F",
				cc.FLAT, "b",
				cc.G, "G",
				cc.MINOR, "m",
				cc.A, "A",
				cc.SHARP, "#",
				cc.DIMINISHED, "dim",
				cc.B, "B",
				cc.AUGMENTED, "aug",
			),
		},
		{
			name:  "accidental",
			input: "C#dim-5[1/4]",
			want: newTokens(
				cc.C, "C",
				cc.SHARP, "#",
				cc.DIMINISHED, "dim",
				cc.MINUS, "-",
				cc.INT, "5",
				cc.LBRA, "[",
				cc.INT, "1",
				cc.SLASH, "/",
				cc.INT, "4",
				cc.RBRA, "]",
			),
		},
		{
			name: "ignore line comment",
			input: `C[1] // to be ignored
D[1]`,
			want: newTokens(
				cc.C, "C",
				cc.LBRA, "[",
				cc.INT, "1",
				cc.RBRA, "]",
				cc.D, "D",
				cc.LBRA, "[",
				cc.INT, "1",
				cc.RBRA, "]",
			),
		},
		{
			name: "ignore multiline comment",
			input: `C[1] /* to be ignored
desc
*/ D[1]`,
			want: newTokens(
				cc.C, "C",
				cc.LBRA, "[",
				cc.INT, "1",
				cc.RBRA, "]",
				cc.D, "D",
				cc.LBRA, "[",
				cc.INT, "1",
				cc.RBRA, "]",
			),
		},
		{
			name:  "maj",
			input: "Bmaj7[1]",
			want: newTokens(
				cc.B, "B",
				cc.MAJOR, "maj",
				cc.SEVENTH, "7",
				cc.LBRA, "[",
				cc.INT, "1",
				cc.RBRA, "]",
			),
		},
		{
			name:  "meter",
			input: "meter[3/4]",
			want: newTokens(
				cc.METER, "meter",
				cc.LBRA, "[",
				cc.INT, "3",
				cc.SLASH, "/",
				cc.INT, "4",
				cc.RBRA, "]",
			),
		},
		{
			name:  "key",
			input: "key[C#] key[Dminor] key[Amajor]",
			want: newTokens(
				cc.KEY, "key",
				cc.LBRA, "[",
				cc.C, "C",
				cc.SHARP, "#",
				cc.RBRA, "]",
				cc.KEY, "key",
				cc.LBRA, "[",
				cc.D, "D",
				cc.MINOR, "minor",
				cc.RBRA, "]",
				cc.KEY, "key",
				cc.LBRA, "[",
				cc.A, "A",
				cc.MAJOR, "major",
				cc.RBRA, "]",
			),
		},
		{
			name:  "on",
			input: "C6/E[1]",
			want: newTokens(
				cc.C, "C",
				cc.SIXTH, "6",
				cc.SLASH, "/",
				cc.E, "E",
				cc.LBRA, "[",
				cc.INT, "1",
				cc.RBRA, "]",
			),
		},
		{
			name:  "instrument",
			input: `inst["Acoustic Piano"]`,
			want: newTokens(
				cc.INSTRUMENT, "inst",
				cc.LBRA, "[",
				cc.STRING, "Acoustic Piano",
				cc.RBRA, "]",
			),
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
				g := cc.NewToken(tok, l.Buffer())
				t.Logf("got token %s", g)
				got = append(got, g)
				l.ResetBuffer()
			}
			if err := l.Err(); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, len(tc.want), len(got))
			for i, w := range tc.want {
				g := got[i]
				assert.Equal(t, w.Type(), g.Type(), i)
				assert.Equal(t, w.Value(), g.Value(), i)
			}
		})
	}
}
