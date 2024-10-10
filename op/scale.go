package op

import (
	"fmt"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/util"
)

type Scale struct {
	Key   Key           `json:"key" yaml:"key"`
	Notes [7]*ScaleNote `json:"notes" yaml:"notes"`
	Flat  int           `json:"flat,omitempty" yaml:"flat,omitempty"`
	Sharp int           `json:"sharp,omitempty" yaml:"sharp,omitempty"`
}

func (s Scale) Tonic() *ScaleNote {
	return s.Notes[0]
}

func (s Scale) GetNoteIndexByName(n note.Name) (int, error) {
	for i, x := range s.Notes {
		if x.Name == n {
			return i, nil
		}
	}
	return -1, errorx.Unexpected("%v is absent on %v", n, s.Key)
}

type ScaleNote struct {
	Name       note.Name  `json:"note" yaml:"note"`
	Accidental Accidental `json:"accidental,omitempty" yaml:"accidental,omitempty"`
}

func (n ScaleNote) String() string {
	return fmt.Sprintf("%v%v", n.Name, n.Accidental)
}

func (n ScaleNote) MarshalYAML() (any, error) {
	return n.String(), nil
}

func (n ScaleNote) GetDegree(x *ScaleNote, isSharp bool) (note.Degree, error) {
	s := x.Semitone() - n.Semitone()
	oct := note.Octave(1).Semitone()
	if s < 0 {
		s += oct
	}

	value, ok := n.Name.GetDegree(x.Name)
	var defaultDegree note.Degree
	if !ok {
		return defaultDegree, errorx.Invalid("UnknownName %v: ScaleNote: %v", x)
	}

	get := func(cds ...note.CoerceDegreeName) (note.Degree, bool) {
		for _, cd := range cds {
			if d, ok := cd.Degree(value); ok {
				if ds, _ := d.Semitone(); ds == s {
					return d, true
				}
			}
		}
		return defaultDegree, false
	}

	if isSharp {
		if d, ok := get(
			note.MajorOrPerfectCoerceDegree,
			note.AugmentedCoerceDegree,
			note.MinorOrDiminishedCoerceDegree,
			note.DoublyAugmentedCoerceDegree,
			note.DoublyDiminishedCoerceDegree,
		); ok {
			return d, nil
		}
	} else {
		if d, ok := get(
			note.MajorOrPerfectCoerceDegree,
			note.MinorOrDiminishedCoerceDegree,
			note.AugmentedCoerceDegree,
			note.DoublyDiminishedCoerceDegree,
			note.DoublyAugmentedCoerceDegree,
		); ok {
			return d, nil
		}
	}

	return defaultDegree, errorx.Invalid(
		"Cannot get degree from ScaleNote %v and %v, sharp = %v, semitone = %v, value = %d",
		n, x, isSharp, s, value,
	)
}

func (n ScaleNote) Semitone() note.Semitone {
	s := n.Name.Semitone()
	r := s + n.Accidental.Semitone()
	return r
}

func MustNewScale(key Key) *Scale {
	s, err := NewScale(key)
	logx.PanicOnError(err)
	return s
}

func NewScale(key Key) (*Scale, error) {
	sig, ok := keySignatures[key]
	if !ok {
		return nil, errorx.NotFound("Unknown scale: %v", key)
	}

	notes := newRawScaleNotes(key.Name)
	for _, x := range notes {
		if sig.names.In(x.Name) {
			if sig.isSharp {
				x.Accidental = Sharp
				continue
			}
			x.Accidental = Flat
			continue
		}
		x.Accidental = Natural
	}

	result := &Scale{
		Key:   key,
		Notes: notes,
	}
	if sig.isSharp {
		result.Sharp = sig.names.Len()
	} else {
		result.Flat = sig.names.Len()
	}

	return result, nil
}

func AllScales() []*Scale {
	var (
		i      int
		scales = make([]*Scale, len(keySignatures))
	)
	for k := range keySignatures {
		scales[i], _ = NewScale(k)
		i++
	}
	return scales
}

func newRawScaleNotes(tonic note.Name) [7]*ScaleNote {
	names := util.MustNewRing(note.C, note.D, note.E, note.F, note.G, note.A, note.B)

	var index int
	switch tonic {
	case note.D:
		index = 1
	case note.E:
		index = 2
	case note.F:
		index = 3
	case note.G:
		index = 4
	case note.A:
		index = 5
	case note.B:
		index = 6
	default:
		index = 0
	}

	var result [7]*ScaleNote
	for i := range 7 {
		result[i] = &ScaleNote{
			Name: names.At(index + i),
		}
	}
	return result
}

var (
	keySignatures = func() map[Key]scaleAccidentals {
		r := make(map[Key]scaleAccidentals, len(keyStringSignatures))
		for k, v := range keyStringSignatures {
			r[MustParseKey(k)] = newScaleAccidentals(v)
		}
		return r
	}()

	keyStringSignatures = map[string]int{
		// major
		"Gb": -6,
		"Db": -5,
		"C#": 7,
		"Ab": -4,
		"Eb": -3,
		"Bb": -2,
		"F":  -1,
		"C":  0,
		"G":  1,
		"D":  2,
		"A":  3,
		"E":  4,
		"B":  5,
		"Cb": -7,
		"F#": 6,
		// minor
		"Ebm": -6,
		"Bbm": -5,
		"Fm":  -4,
		"Cm":  -3,
		"Gm":  -2,
		"Dm":  -1,
		"Am":  0,
		"Em":  1,
		"Bm":  2,
		"F#m": 3,
		"C#m": 4,
		"G#m": 5,
		"D#m": 6,
	}
)

type scaleAccidentals struct {
	isSharp bool
	names   util.Set[note.Name]
}

func newScaleAccidentals(n int) scaleAccidentals {
	switch {
	case n < 0: // flat
		return scaleAccidentals{
			names: util.NewSet(flatSequence[:-n]...),
		}
	case n > 0: // sharp
		return scaleAccidentals{
			isSharp: true,
			names:   util.NewSet(flatSequence[(len(flatSequence))-n:]...),
		}
	default:
		return scaleAccidentals{
			names: util.NewSet[note.Name](),
		}
	}
}

var (
	flatSequence = []note.Name{
		note.B,
		note.E,
		note.A,
		note.D,
		note.G,
		note.C,
		note.F,
	}
)
