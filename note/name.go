package note

import (
	"errors"
	"slices"

	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/util"
)

type Name int

const (
	UnknownName Name = iota
	C
	D
	E
	F
	G
	A
	B
)

var (
	nameStringMap = map[Name]string{
		C: "C",
		D: "D",
		E: "E",
		F: "F",
		G: "G",
		A: "A",
		B: "B",
	}
	stringNameMap = util.MustInverseMap(nameStringMap)
)

func NewName(s string) Name {
	if x, ok := stringNameMap[s]; ok {
		return x
	}
	return UnknownName
}

func (n Name) String() string {
	return nameStringMap[n]
}

func (n Name) MarshalYAML() (any, error) {
	return n.String(), nil
}

func (n Name) MarshalJSON() ([]byte, error) {
	return []byte(n.String()), nil
}

var (
	ErrUnknownName  = errors.New("UnknownName")
	nameSemitoneMap = map[Name]Semitone{
		C: 0,
		D: 2,
		E: 4,
		F: 5,
		G: 7,
		A: 9,
		B: 11,
	}
)

func (n Name) Semitone() Semitone {
	if v, ok := nameSemitoneMap[n]; ok {
		return v
	}
	logx.Panic(ErrUnknownName)
	return -1
}

func (n Name) AddDegree(degree int) (Name, Octave) {
	if -1 <= degree && degree <= 1 {
		// for convinience, include 0
		return n, 0
	}

	names := []Name{C, D, E, F, G, A, B}
	index := slices.Index(names, n)
	if index < 0 {
		logx.Panic(ErrUnknownName)
	}

	newIndex := index
	if degree > 0 {
		// 1 degree is identical
		newIndex += degree - 1
	} else {
		newIndex += degree + 1
	}

	wantIndex := newIndex % len(names)
	octave := newIndex / len(names)
	if wantIndex < 0 {
		// modulus may return [1-len(names), 0)
		wantIndex += len(names)
		octave--
	}

	return names[wantIndex], Octave(octave)
}

var (
	nameRing = util.MustNewRing(
		C,
		D,
		E,
		F,
		G,
		A,
		B,
	)
)

func (x Name) GetDegree(y Name) (uint, bool) {
	if x == UnknownName || y == UnknownName {
		return 0, false
	}
	if x == y {
		return 1, true
	}
	var (
		xi = -1
		yi = -1
		i  int
	)

	for xi < 0 {
		if nameRing.At(i) == x {
			xi = i
			break
		}
		i++
	}
	for yi < 0 {
		if nameRing.At(i) == y {
			yi = i
			break
		}
		i++
	}

	return uint(yi - xi + 1), true
}
