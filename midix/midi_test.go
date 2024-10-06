package midix_test

import (
	"bytes"
	"testing"

	"github.com/berquerant/crd/midix"
	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	w := midix.NewWriter(midix.DefaultTicksPerQuoaterNote)
	w.Meter(4, 4)
	w.Tempo(120)
	assert.Nil(t, w.Note(1, 100, 60, 64))
	assert.Nil(t, w.Note(2, 100, 62, 66))
	w.Rest(1)
	assert.Nil(t, w.Note(1, 100, 64, 68))
	w.Close()

	var buf bytes.Buffer
	_, err := w.WriteTo(&buf)
	assert.Nil(t, err)

	type row struct {
		no      int
		ticks   int64
		value   int64
		message string
	}

	want := []row{
		{
			message: `MetaTrackName text: "Piano"`,
		},
		{
			message: `MetaInstrument text: "Piano"`,
		},
		{
			message: `ProgramChange channel: 0 program: 0`,
		},
		{
			message: `MetaTimeSig meter: 4/4`,
		},
		{
			message: `MetaTempo bpm: 120.00`,
		},
		{
			message: `NoteOn channel: 0 key: 60 velocity: 100`,
		},
		{
			message: `NoteOn channel: 0 key: 64 velocity: 100`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote,
			value:   1,
			message: `NoteOff channel: 0 key: 60`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote,
			value:   1,
			message: `NoteOff channel: 0 key: 64`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote,
			value:   1,
			message: `NoteOn channel: 0 key: 62 velocity: 100`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote,
			value:   1,
			message: `NoteOn channel: 0 key: 66 velocity: 100`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote * 3,
			value:   3,
			message: `NoteOff channel: 0 key: 62`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote * 3,
			value:   3,
			message: `NoteOff channel: 0 key: 66`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote * 4,
			value:   4,
			message: `NoteOn channel: 0 key: 64 velocity: 100`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote * 4,
			value:   4,
			message: `NoteOn channel: 0 key: 68 velocity: 100`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote * 5,
			value:   5,
			message: `NoteOff channel: 0 key: 64`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote * 5,
			value:   5,
			message: `NoteOff channel: 0 key: 68`,
		},
		{
			ticks:   midix.DefaultTicksPerQuoaterNote * 5,
			value:   5,
			message: `MetaEndOfTrack`,
		},
	}

	rd := midix.NewReader()
	var got []row
	for ev := range rd.Events(bytes.NewReader(buf.Bytes())) {
		got = append(got, row{
			no:      ev.TrackNo,
			ticks:   ev.AbsTicks,
			value:   ev.AbsTicks / midix.DefaultTicksPerQuoaterNote,
			message: ev.Message.String(),
		})
	}

	assert.Equal(t, want, got)
}
