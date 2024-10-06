package note

import (
	"errors"
	"fmt"

	"github.com/berquerant/crd/logx"
)

type Accidental int

const (
	UnknownAccidental Accidental = iota
	Natural
	Sharp
	Flat
	DoubleSharp
	DoubleFlat
)

var (
	accidentalSemitoneMap = map[Accidental]Semitone{
		Natural:     0,
		Sharp:       1,
		Flat:        -1,
		DoubleSharp: 2,
		DoubleFlat:  -2,
	}
)

func (a Accidental) Semitone() Semitone {
	if x, ok := accidentalSemitoneMap[a]; ok {
		return x
	}
	logx.Panic(fmt.Errorf("%w: %v", ErrUnknownAccidental, a))
	return 0
}

var (
	ErrUnknownAccidental = errors.New("UnknownAccidental")
)

const (
	accidentalNatural     = "♮"
	accidentalSharp       = "♯"
	accidentalFlat        = "♭"
	accidentalDoubleSharp = "𝄪"
	accidentalDoubleFlat  = "𝄫"

	accidentalNaturalSimple     = "n"
	accidentalSharpSimple       = "#"
	accidentalFlatSimple        = "b"
	accidentalDoubleSharpSimple = "##"
	accidentalDoubleFlatSimple  = "bb"

	accidentalUnknown = "UnknownAccidental"
)

type accidentalStringCell struct {
	origin string
	simple string
}

func (c accidentalStringCell) is(v string) bool {
	return v == c.origin || v == c.simple
}

func (c accidentalStringCell) get(simple bool) string {
	if simple {
		return c.simple
	}
	return c.origin
}

var (
	accidentalStringMap = map[Accidental]accidentalStringCell{
		Natural: {
			origin: accidentalNatural,
			simple: accidentalNaturalSimple,
		},
		Sharp: {
			origin: accidentalSharp,
			simple: accidentalSharpSimple,
		},
		Flat: {
			origin: accidentalFlat,
			simple: accidentalFlatSimple,
		},
		DoubleSharp: {
			origin: accidentalDoubleSharp,
			simple: accidentalDoubleSharpSimple,
		},
		DoubleFlat: {
			origin: accidentalDoubleFlat,
			simple: accidentalDoubleFlatSimple,
		},
	}
)

func NewAccidental(s string) Accidental {
	for k, v := range accidentalStringMap {
		if v.is(s) {
			return k
		}
	}
	return UnknownAccidental
}

func (a Accidental) String(simple bool) string {
	if v, ok := accidentalStringMap[a]; ok {
		return v.get(simple)
	}
	return accidentalUnknown
}
