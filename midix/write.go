package midix

import (
	"io"
	"math"

	"github.com/berquerant/crd/errorx"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
)

const (
	DefaultTrackSequenceName   = "Piano"
	DefaultInstrument          = "Piano"
	DefaultTicksPerQuoaterNote = 960
)

var (
	_              Writer = &MIDIWriter{}
	DefaultProgram        = gm.Instr_AcousticGrandPiano.Value()
)

type Writer interface {
	Note(value float64, velocity uint8, key ...uint8) error
	Tempo(bpm int)
	Meter(num, denom uint8)
	Key(key uint8, isMajor bool, num uint8, isFlat bool)
	Close()
	Rest(value float64)
	WriteTo(out io.Writer) (int64, error)
}

type (
	Track     = smf.Track
	operation func(*Track)
)

type MIDIWriter struct {
	clock            smf.MetricTicks
	tickDelta        uint32
	quoaterNoteTicks uint32
	ops              []operation
}

func NewWriter(ticksPerQuoaterNote uint16) *MIDIWriter {
	clock := smf.MetricTicks(ticksPerQuoaterNote)
	w := &MIDIWriter{
		clock:            clock,
		quoaterNoteTicks: clock.Ticks4th(),
	}
	w.init()
	return w
}

func (w *MIDIWriter) init() {
	w.add(func(t *Track) {
		// add metadata
		t.Add(0, smf.MetaTrackSequenceName(DefaultTrackSequenceName))
		t.Add(0, smf.MetaInstrument(DefaultInstrument))
		t.Add(0, midi.ProgramChange(0, DefaultProgram))
	})
}

func (w MIDIWriter) WriteTo(out io.Writer) (int64, error) {
	var track smf.Track
	for _, op := range w.ops {
		op(&track)
	}
	s := smf.New()
	s.TimeFormat = w.clock
	if err := s.Add(track); err != nil {
		return 0, err
	}
	return s.WriteTo(out)
}

func (w *MIDIWriter) Note(value float64, velocity uint8, key ...uint8) error {
	if len(key) == 0 {
		return errorx.Invalid("midi note requires keys")
	}

	var (
		ticks     = w.getTickDeltaAndClear()
		nextTicks = w.newTicks(value)
	)
	w.add(func(t *Track) {
		for i, k := range key {
			if i == 0 {
				t.Add(ticks, midi.NoteOn(0, k, velocity))
				continue
			}
			t.Add(0, midi.NoteOn(0, k, velocity))
		}
		for i, k := range key {
			if i == 0 {
				t.Add(nextTicks, midi.NoteOff(0, k))
				continue
			}
			t.Add(0, midi.NoteOff(0, k))
		}
	})
	return nil
}

func (w *MIDIWriter) Rest(value float64) {
	w.addTickDelta(w.newTicks(value))
}

func (w *MIDIWriter) Tempo(bpm int) {
	x := w.getTickDeltaAndClear()
	w.add(func(t *Track) {
		t.Add(x, smf.MetaTempo(float64(bpm)))
	})
}

func (w *MIDIWriter) Meter(num, denom uint8) {
	x := w.getTickDeltaAndClear()
	w.add(func(t *Track) {
		t.Add(x, smf.MetaMeter(num, denom))
	})
}

func (w *MIDIWriter) Key(key uint8, isMajor bool, num uint8, isFlat bool) {
	x := w.getTickDeltaAndClear()
	w.add(func(t *Track) {
		t.Add(x, smf.MetaKey(key, isMajor, num, isFlat))
	})
}

func (w *MIDIWriter) Close() {
	w.add(func(t *Track) {
		t.Close(0)
	})
}

func (w *MIDIWriter) getTickDeltaAndClear() uint32 {
	x := w.tickDelta
	w.tickDelta = 0
	return x
}

func (w *MIDIWriter) addTickDelta(t uint32) { w.tickDelta += t }
func (w MIDIWriter) newTicks(multiplier float64) uint32 {
	return uint32(math.Round(float64(w.quoaterNoteTicks) * multiplier))
}

func (w *MIDIWriter) add(op operation) {
	w.ops = append(w.ops, op)
}
