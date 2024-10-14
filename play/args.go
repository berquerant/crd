package play

import (
	"github.com/berquerant/crd/midix"
	"github.com/berquerant/crd/op"
	"github.com/berquerant/crd/util"
)

type midiArgs struct {
	bpm      *util.Opt[op.BPM]
	meter    *util.Opt[op.Meter]
	velocity *util.Opt[op.DynamicSign]
	key      *util.Opt[op.Key]
	meta     *util.Opt[op.Meta]
}

func newMidiArgs() *midiArgs {
	return &midiArgs{
		bpm:      util.NewOpt(defaultBPM),
		meter:    util.NewOpt(defaultMeter),
		velocity: util.NewOpt(defaultVelocity),
		key:      util.NewOpt(defaultKey),
		meta:     util.NewOpt(defaultMeta),
	}
}

func (m *midiArgs) writeWhenUpdated(w midix.Writer) {
	m.bpm.WhenUpdated(func(v op.BPM) {
		w.Tempo(int(v))
	})
	m.meter.WhenUpdated(func(v op.Meter) {
		w.Meter(uint8(v.Num), uint8(v.Denom))
	})
	m.key.WhenUpdated(func(v op.Key) {
		scale := op.MustNewScale(v)
		key := uint8(scale.Tonic().Semitone())
		isMajor := !scale.Key.Minor
		num := uint8(scale.Flat + scale.Sharp)
		isFlat := scale.Flat > 0
		w.Key(key, isMajor, num, isFlat)
	})
	m.meta.WhenUpdated(func(v op.Meta) {
		if x := v.Get("txt"); x != "" {
			w.Text(x)
		}
		if x := v.Get("lic"); x != "" {
			w.Lyric(x)
		}
		if x := v.Get("mrk"); x != "" {
			w.Marker(x)
		}
	})
}

func (m midiArgs) getKey() op.Key     { return m.key.Unwrap() }
func (m midiArgs) getVelocity() uint8 { return uint8(m.velocity.Unwrap().Velocity()) }

func (m *midiArgs) update(instance op.Instance) {
	if x := instance.BPM; x != nil {
		m.bpm.Update(*x)
	}
	if x := instance.Meter; x != nil {
		m.meter.Update(*x)
	}
	if x := instance.Velocity; x != nil {
		m.velocity.Update(*x)
	}
	if x := instance.Key; x != nil {
		m.key.Update(*x)
	}
	if x := instance.Meta; x != nil {
		m.meta.Update(*x)
	}
}
