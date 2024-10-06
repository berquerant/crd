package play

import (
	"fmt"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/midix"
	"github.com/berquerant/crd/op"
)

var (
	defaultBPM, _   = op.NewBPM(100)
	defaultKey      = op.MustParseKey("C")
	defaultMeter    = op.MustNewMeter(4, 4)
	defaultVelocity = op.MezzoPiano
)

type Writer interface {
	Write(w midix.Writer, instances []op.Instance) error
}

var (
	_ Writer = &MIDIWriter{}
)

func NewWriter(cmap chord.Mapper, newKey func(op.Key) Key) *MIDIWriter {
	return &MIDIWriter{
		cmap:   cmap,
		newKey: newKey,
	}
}

type MIDIWriter struct {
	cmap   chord.Mapper
	newKey func(op.Key) Key
}

func (m MIDIWriter) Write(w midix.Writer, instances []op.Instance) error {
	if len(instances) == 0 {
		return errorx.Invalid("MIDIWriter requires 1 or more instances")
	}

	args := newMidiArgs()

	for i, instance := range instances {
		if err := instance.Validate(); err != nil {
			return fmt.Errorf("%w: instance[%d]", err, i)
		}

		args.update(instance)
		// apply control changes
		args.writeWhenUpdated(w)

		var value float64
		for _, v := range instance.Values {
			value += v.Float()
		}

		if instance.IsRest() {
			w.Rest(value)
			continue
		}

		key := m.newKey(args.getKey())
		numbers, err := key.Apply(*instance.Chord)
		if err != nil {
			return fmt.Errorf("%w: instance[%d]", err, i)
		}
		midiKeys := make([]uint8, len(numbers))
		for i, x := range numbers {
			midiKeys[i] = uint8(x)
		}

		if err := w.Note(value, args.getVelocity(), midiKeys...); err != nil {
			return fmt.Errorf("%w: instance[%d]", err, i)
		}
	}

	w.Close()
	return nil
}
