package midix_test

import (
	"testing"

	"github.com/berquerant/crd/midix"
	"github.com/stretchr/testify/assert"
)

func TestTrackNoSelector(t *testing.T) {
	for _, tc := range []struct {
		trackNo int
		opType  midix.OpType
		want    int
	}{
		{
			trackNo: 1,
			opType:  midix.NewMetaTrack(),
			want:    0,
		},
		{
			trackNo: 1,
			opType:  midix.NewFixedTrack(0),
			want:    0,
		},
		{
			trackNo: 1,
			opType:  midix.NewFixedTrack(1),
			want:    0,
		},
		{
			trackNo: 2,
			opType:  midix.NewMetaTrack(),
			want:    0,
		},
		{
			trackNo: 2,
			opType:  midix.NewFixedTrack(0),
			want:    1,
		},
		{
			trackNo: 2,
			opType:  midix.NewFixedTrack(1),
			want:    1,
		},
		{
			trackNo: 2,
			opType:  midix.NewFixedTrack(2),
			want:    1,
		},
		{
			trackNo: 3,
			opType:  midix.NewMetaTrack(),
			want:    0,
		},
		{
			trackNo: 3,
			opType:  midix.NewFixedTrack(0),
			want:    1,
		},
		{
			trackNo: 3,
			opType:  midix.NewFixedTrack(1),
			want:    2,
		},
		{
			trackNo: 3,
			opType:  midix.NewFixedTrack(2),
			want:    1,
		},
	} {
		s, err := midix.NewTrackNoSelector(tc.trackNo)
		if !assert.Nil(t, err) {
			continue
		}
		assert.Equal(t, tc.want, s.Select(tc.opType))
	}
}
