package op_test

import (
	"testing"

	"github.com/berquerant/crd/op"
	"github.com/stretchr/testify/assert"
)

func TestDiatonicChorder(t *testing.T) {
	for _, tc := range []struct {
		scale        *op.Scale
		wantTriads   []string
		wantSevenths []string
	}{
		{
			scale: op.MustNewScale(op.MustParseKey("C")),
			wantTriads: []string{
				"C",
				"Dm",
				"Em",
				"F",
				"G",
				"Am",
				"Bdim",
			},
			wantSevenths: []string{
				"Cmaj7",
				"Dm7",
				"Em7",
				"Fmaj7",
				"G_7",
				"Am7",
				"Bm7b5",
			},
		},
		{
			scale: op.MustNewScale(op.MustParseKey("Am")),
			wantTriads: []string{
				"Am",
				"Bdim",
				"C",
				"Dm",
				"Em",
				"F",
				"G",
			},
			wantSevenths: []string{
				"Am7",
				"Bm7b5",
				"Cmaj7",
				"Dm7",
				"Em7",
				"Fmaj7",
				"G_7",
			},
		},
	} {
		t.Run(tc.scale.Key.String(), func(t *testing.T) {
			dc := op.NewDiatonicChorder(tc.scale)

			gotTriads := []string{}
			for _, x := range dc.Triads() {
				gotTriads = append(gotTriads, x.String())
			}
			assert.Equal(t, tc.wantTriads, gotTriads)

			gotSevenths := []string{}
			for _, x := range dc.Sevenths() {
				gotSevenths = append(gotSevenths, x.String())
			}
			assert.Equal(t, tc.wantSevenths, gotSevenths)
		})
	}
}
