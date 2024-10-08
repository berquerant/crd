package op_test

import (
	"log/slog"
	"os"
	"slices"
	"testing"

	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/op"
	"github.com/stretchr/testify/assert"
)

func TestKeyConversionChain(t *testing.T) {
	logx.Setup(os.Stdout, slog.LevelDebug)

	for _, tc := range []struct {
		title string
		key   op.Key
		chain op.KeyConversionChain
		want  []op.Key
	}{
		{
			title: "parallel parallel C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.ParallelKey,
				op.ParallelKey,
			},
			want: []op.Key{
				op.MustParseKey("C"),
			},
		},
		{
			title: "parallel C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.ParallelKey,
			},
			want: []op.Key{
				op.MustParseKey("Cm"),
			},
		},
		{
			title: "relative relative C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.RelativeKey,
				op.RelativeKey,
			},
			want: []op.Key{
				op.MustParseKey("C"),
			},
		},
		{
			title: "relative C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.RelativeKey,
			},
			want: []op.Key{
				op.MustParseKey("Am"),
			},
		},
		{
			title: "dominant C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.DominantKey,
			},
			want: []op.Key{
				op.MustParseKey("G"),
			},
		},
		{
			title: "dominant x 12 C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
				op.DominantKey,
			},
			want: []op.Key{
				op.MustParseKey("C"),
			},
		},
		{
			title: "subdominant C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.SubDominantKey,
			},
			want: []op.Key{
				op.MustParseKey("F"),
			},
		},
		{
			title: "subdominant x 12 C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
				op.SubDominantKey,
			},
			want: []op.Key{
				op.MustParseKey("C"),
			},
		},
		{
			title: "parallel dominant C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.ParallelKey,
				op.DominantKey,
			},
			want: []op.Key{
				op.MustParseKey("Gm"),
			},
		},
		{
			title: "parallel subdominant C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.ParallelKey,
				op.SubDominantKey,
			},
			want: []op.Key{
				op.MustParseKey("Fm"),
			},
		},
		{
			title: "dominant parallel C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.DominantKey,
				op.ParallelKey,
			},
			want: []op.Key{
				op.MustParseKey("Gm"),
			},
		},
		{
			title: "subdominant parallel C",
			key:   op.MustParseKey("C"),
			chain: []op.KeyConversion{
				op.SubDominantKey,
				op.ParallelKey,
			},
			want: []op.Key{
				op.MustParseKey("Fm"),
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			m, err := tc.chain.Convert(op.NewCircleOfFifth(), tc.key)
			if !assert.Nil(t, err) {
				return
			}
			got := slices.Collect(m.Keys().All())
			assert.Equal(t, tc.want, got, "want: %v got: %v", tc.want, got)
		})
	}
}
