package note

import (
	"errors"
	"fmt"
	"iter"
	"strings"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/util"
	"gopkg.in/yaml.v3"
)

//go:generate go tool stringer -type DegreeName -output degree_stringer_generated.go
type DegreeName int

const (
	UnknownDegree DegreeName = iota
	MajorDegree
	MinorDegree
	PerfectDegree
	AugmentedDegree
	DiminishedDegree
	DoublyAugmentedDegree
	DoublyDiminishedDegree
)

func (d DegreeName) Coerce() CoerceDegreeName {
	if v, ok := degreeCoerceMap[d]; ok {
		return v
	}
	return UnknownCoerceDegreeName
}

type Degree struct {
	Value uint
	Name  DegreeName
}

func (d Degree) String() string {
	return fmt.Sprintf("%s%d", d.Name.Coerce(), d.Value)
}

func (d Degree) MarshalYAML() (any, error) {
	return d.String(), nil
}

func (d *Degree) UnmarshalYAML(value *yaml.Node) error {
	x, err := ParseDegree(value.Value)
	if err != nil {
		return err
	}
	*d = x
	return nil
}

var (
	perfect1 = Degree{
		Value: 1,
		Name:  PerfectDegree,
	}
	minor2 = Degree{
		Value: 2,
		Name:  MinorDegree,
	}
	major2 = Degree{
		Value: 2,
		Name:  MajorDegree,
	}
	minor3 = Degree{
		Value: 3,
		Name:  MinorDegree,
	}
	major3 = Degree{
		Value: 3,
		Name:  MajorDegree,
	}
	perfect4 = Degree{
		Value: 4,
		Name:  PerfectDegree,
	}
	augmented4 = Degree{
		Value: 4,
		Name:  AugmentedDegree,
	}
	diminished5 = Degree{
		Value: 5,
		Name:  DiminishedDegree,
	}
	perfect5 = Degree{
		Value: 5,
		Name:  PerfectDegree,
	}
	minor6 = Degree{
		Value: 6,
		Name:  MinorDegree,
	}
	major6 = Degree{
		Value: 6,
		Name:  MajorDegree,
	}
	minor7 = Degree{
		Value: 7,
		Name:  MinorDegree,
	}
	major7 = Degree{
		Value: 7,
		Name:  MajorDegree,
	}
	perfect8 = Degree{
		Value: 8,
		Name:  PerfectDegree,
	}

	degreeSemitoneMap = map[Degree]Semitone{
		perfect1:    0,
		minor2:      1,
		major2:      2,
		minor3:      3,
		major3:      4,
		perfect4:    5,
		augmented4:  6,
		diminished5: 6,
		perfect5:    7,
		minor6:      8,
		major6:      9,
		minor7:      10,
		major7:      11,
		perfect8:    12,
	}
)

func (d Degree) Semitone() (Semitone, bool) {
	if d.Value == 0 {
		return 0, false
	}
	if d.Value <= perfect8.Value {
		if v, ok := degreeSemitoneMap[d]; ok {
			return v, true
		}

		for k, v := range degreeSemitoneMap {
			if k.Value != d.Value {
				continue
			}
			switch {
			case d.Name == AugmentedDegree && (k.Name == MajorDegree || k.Name == PerfectDegree):
				return v + 1, true
			case d.Name == DiminishedDegree && (k.Name == MinorDegree || k.Name == PerfectDegree):
				return v - 1, true
			case d.Name == DoublyAugmentedDegree && (k.Name == MajorDegree || k.Name == PerfectDegree):
				return v + 2, true
			case d.Name == DoublyDiminishedDegree && (k.Name == MinorDegree || k.Name == PerfectDegree):
				return v - 2, true
			}
		}
		return 0, false
	}

	e := Degree{
		Value: d.Value - perfect8.Value + 1, // perfect1 is identical
		Name:  d.Name,
	}
	if v, ok := e.Semitone(); ok {
		return v + degreeSemitoneMap[perfect8], true
	}
	return 0, false
}

func NewDegree(value uint, name DegreeName) (Degree, bool) {
	d := Degree{
		Value: value,
		Name:  name,
	}
	if _, ok := d.Semitone(); !ok {
		return d, false
	}
	return d, true
}

func MustNewDegree(value uint, name DegreeName) Degree {
	x, ok := NewDegree(value, name)
	if !ok {
		logx.Panic(errorx.Unexpected("MustNewDegree(%d, %s)", value, name))
	}
	return x
}

type CoerceDegreeName int

const (
	UnknownCoerceDegreeName CoerceDegreeName = iota
	MajorOrPerfectCoerceDegree
	MinorOrDiminishedCoerceDegree
	AugmentedCoerceDegree
	DiminishedCoerceDegree
	DoublyAugmentedCoerceDegree
	DoublyDiminishedCoerceDegree
)

const (
	degreeNameFlatted       = accidentalFlatSimple
	degreeNameSharped       = accidentalSharpSimple
	degreeNameNatural       = ""
	degreeNameDiminished    = accidentalDoubleFlatSimple
	degreeNameDoublyFlatted = "bbb"
	degreeNameDoublySharped = accidentalDoubleSharpSimple
)

var (
	degreeCoerceMap = map[DegreeName]CoerceDegreeName{
		MajorDegree:            MajorOrPerfectCoerceDegree,
		PerfectDegree:          MajorOrPerfectCoerceDegree,
		MinorDegree:            MinorOrDiminishedCoerceDegree,
		DiminishedDegree:       DiminishedCoerceDegree,
		AugmentedDegree:        AugmentedCoerceDegree,
		DoublyAugmentedDegree:  DoublyAugmentedCoerceDegree,
		DoublyDiminishedDegree: DoublyDiminishedCoerceDegree,
	}
	stringCoerceDegreeNameMap = map[string]CoerceDegreeName{
		degreeNameFlatted:       MinorOrDiminishedCoerceDegree,
		degreeNameSharped:       AugmentedCoerceDegree,
		degreeNameNatural:       MajorOrPerfectCoerceDegree,
		degreeNameDiminished:    DiminishedCoerceDegree,
		degreeNameDoublySharped: DoublyAugmentedCoerceDegree,
		degreeNameDoublyFlatted: DoublyDiminishedCoerceDegree,
	}
	coerceDegreeNameStringMap = util.MustInverseMap(stringCoerceDegreeNameMap)
)

func NewCoerceDegree(s string) CoerceDegreeName {
	if v, ok := stringCoerceDegreeNameMap[s]; ok {
		return v
	}
	return UnknownCoerceDegreeName
}

func (c CoerceDegreeName) String() string {
	if x, ok := coerceDegreeNameStringMap[c]; ok {
		return x
	}
	logx.Panic(ErrInvalidDegree)
	return ""
}

func (c CoerceDegreeName) Degree(value uint) (Degree, bool) {
	switch c {
	case MajorOrPerfectCoerceDegree:
		if x, ok := NewDegree(value, MajorDegree); ok {
			return x, true
		}
		return NewDegree(value, PerfectDegree)
	case MinorOrDiminishedCoerceDegree:
		if x, ok := NewDegree(value, MinorDegree); ok {
			return x, true
		}
		return NewDegree(value, DiminishedDegree)
	case AugmentedCoerceDegree:
		return NewDegree(value, AugmentedDegree)
	case DiminishedCoerceDegree:
		return NewDegree(value, DiminishedDegree)
	case DoublyAugmentedCoerceDegree:
		return NewDegree(value, DoublyAugmentedDegree)
	case DoublyDiminishedCoerceDegree:
		return NewDegree(value, DoublyDiminishedDegree)
	default:
		var d Degree
		return d, false
	}
}

var (
	ErrInvalidDegree = errors.New("InvalidDegree")
)

func ParseDegree(s string) (Degree, error) {
	var (
		errContinue = errors.New("Continue")
		parse       = func(x, symbol string) (uint, error) {
			if symbol == "" {
				return util.ParseUint(x)
			}
			if !strings.Contains(x, symbol) {
				return 0, errContinue
			}

			v := strings.Trim(x, symbol)
			if v == x {
				return 0, errContinue
			}
			return util.ParseUint(v)
		}

		value uint
		name  CoerceDegreeName
	)

	for _, c := range []struct {
		symbol string
		name   CoerceDegreeName
	}{
		{
			symbol: degreeNameDoublyFlatted,
			name:   DoublyDiminishedCoerceDegree,
		},
		{
			symbol: degreeNameDoublySharped,
			name:   DoublyAugmentedCoerceDegree,
		},
		{
			symbol: degreeNameDiminished,
			name:   DiminishedCoerceDegree,
		},
		{
			symbol: degreeNameFlatted,
			name:   MinorOrDiminishedCoerceDegree,
		},
		{
			symbol: degreeNameSharped,
			name:   AugmentedCoerceDegree,
		},
		{
			name: MajorOrPerfectCoerceDegree,
		},
	} {
		val, err := parse(s, c.symbol)
		if err == nil {
			name = c.name
			value = val
			break
		}
		if errors.Is(err, errContinue) {
			continue
		}
		var d Degree
		return d, err
	}

	if x, ok := name.Degree(value); ok {
		return x, nil
	}
	var d Degree
	return d, ErrInvalidDegree
}

func GenerateDegrees(maxDegree uint) iter.Seq[Degree] {
	degreeNames := []DegreeName{
		MajorDegree,
		MinorDegree,
		PerfectDegree,
		AugmentedDegree,
		DiminishedDegree,
		DoublyAugmentedDegree,
		DoublyDiminishedDegree,
	}
	// FIXME: generate doubly augmented/diminished degrees

	return func(yield func(Degree) bool) {
		for value := range maxDegree {
			for _, name := range degreeNames {
				if d, ok := NewDegree(value, name); ok {
					if !yield(d) {
						return
					}
				}
			}
		}
	}
}
