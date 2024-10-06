package play

import (
	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
)

type Key struct {
	key  op.Key
	cmap chord.Mapper
}

func NewKey(key op.Key, cmap chord.Mapper) Key {
	return Key{
		key:  key,
		cmap: cmap,
	}
}

func (k Key) Apply(c op.Chord) ([]MIDINoteNumber, error) {
	attrs, ok := k.cmap.GetAttributes(c.Chord.Name)
	if !ok {
		return nil, errorx.NotFound("Chord %s", c.Chord.Name)
	}
	cd, ok := c.Degree.Semitone()
	if !ok {
		return nil, errorx.Conversion("Chord %s degree", c)
	}

	var (
		result []MIDINoteNumber
		add    = func(x MIDINoteNumber) {
			result = append(result, x)
		}
		// C4 + key
		keyNumber = MiddleC.MIDINoteNumber() + MIDINoteNumber(k.key.Semitone())
		// root of chord
		rootNumber = keyNumber + MIDINoteNumber(cd)
	)

	// base note
	{
		b, ok := c.Base.Semitone()
		if !ok {
			return nil, errorx.Conversion("Chord %s invalid base", c)
		}
		add(rootNumber + MIDINoteNumber(b) - MIDINoteNumber(note.Octave(1).Semitone()))
	}
	// chord notes
	for _, a := range attrs {
		b, ok := a.Semitone()
		if !ok {
			return nil, errorx.Conversion("Chord %s invalid attribute %s", c, a.Name)
		}
		add(rootNumber + MIDINoteNumber(b))
	}
	return result, nil
}
