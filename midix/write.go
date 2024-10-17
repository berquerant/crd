package midix

import (
	"io"
	"math"

	"github.com/berquerant/crd/errorx"
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
	Text(text string)
	Lyric(text string)
	Marker(text string)
	Close()
	Rest(value float64)
	WriteTo(out io.Writer) (int64, error)
}

type MIDIWriter struct {
	clock            smf.MetricTicks
	tickDelta        uint32
	quoaterNoteTicks uint32
	set              *TrackSetController
}

func NewWriter(ticksPerQuoaterNote uint16, set *TrackSetController) *MIDIWriter {
	clock := smf.MetricTicks(ticksPerQuoaterNote)
	w := &MIDIWriter{
		clock:            clock,
		quoaterNoteTicks: clock.Ticks4th(),
		set:              set,
	}
	w.init()
	return w
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

func (w *MIDIWriter) add(op *TrackOp) {
	w.set.Add(op)
}
func (w *MIDIWriter) addMeta(tickDelta uint32, opFunc OpFunc) {
	w.add(NewTrackOp(tickDelta, NewMetaTrack(), opFunc))
}
func (w *MIDIWriter) addFixed(tickDelta uint32, trackNo int, opFunc OpFunc) {
	w.add(NewTrackOp(tickDelta, NewFixedTrack(trackNo), opFunc))
}
func (w *MIDIWriter) addAll(tickDelta uint32, opFunc OpFunc) {
	w.set.Distribute(NewTrackOp(tickDelta, NewMetaTrack(), opFunc))
}

func (w *MIDIWriter) init() {
	w.addMeta(0, &MetaTrackSequenceName{
		Text: DefaultTrackSequenceName,
	})
	w.addMeta(0, &MetaInstrument{
		Text: DefaultInstrument,
	})
	w.addMeta(0, &ProgramChange{
		Channel: 0,
		Program: DefaultProgram,
	})
}

func (w MIDIWriter) WriteTo(out io.Writer) (int64, error) {
	s := smf.New()
	s.TimeFormat = w.clock
	for i := range w.set.Set().Len() {
		var t smf.Track
		w.set.Set().Get(i).Apply(&t)
		if err := s.Add(t); err != nil {
			return 0, err
		}
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
	for i, k := range key {
		if i == 0 {
			w.addFixed(ticks, i, &NoteOn{
				Channel:  0,
				Key:      k,
				Velocity: velocity,
			})
			continue
		}
		w.addFixed(0, i, &NoteOn{
			Channel:  0,
			Key:      k,
			Velocity: velocity,
		})
	}
	for i, k := range key {
		if i == 0 {
			w.addFixed(nextTicks, i, &NoteOff{
				Channel: 0,
				Key:     k,
			})
			continue
		}
		w.addFixed(0, i, &NoteOff{
			Channel: 0,
			Key:     k,
		})
	}
	return nil
}

func (w *MIDIWriter) Rest(value float64) {
	w.addTickDelta(w.newTicks(value))
}

func (w *MIDIWriter) Tempo(bpm int) {
	w.addMeta(w.getTickDeltaAndClear(), &MetaTempo{
		BPM: float64(bpm),
	})
}

func (w *MIDIWriter) Meter(num, denom uint8) {
	w.addMeta(w.getTickDeltaAndClear(), &MetaMeter{
		Num:   num,
		Denom: denom,
	})
}

func (w *MIDIWriter) Key(key uint8, isMajor bool, num uint8, isFlat bool) {
	w.addMeta(w.getTickDeltaAndClear(), &MetaKey{
		Key:     key,
		IsMajor: isMajor,
		Num:     num,
		IsFlat:  isFlat,
	})
}

func (w *MIDIWriter) Text(text string) {
	w.addMeta(w.getTickDeltaAndClear(), &MetaText{
		Text: text,
	})
}

func (w *MIDIWriter) Lyric(text string) {
	w.addMeta(w.getTickDeltaAndClear(), &MetaLyric{
		Text: text,
	})
}

func (w *MIDIWriter) Marker(text string) {
	w.addMeta(w.getTickDeltaAndClear(), &MetaMarker{
		Text: text,
	})
}

func (w *MIDIWriter) Close() {
	w.addAll(0, &Close{})
}
