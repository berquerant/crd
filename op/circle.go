package op

import (
	"fmt"
	"iter"
	"log/slog"
	"maps"
	"slices"
	"strings"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/util"
)

type CircleMember struct {
	scales map[Key]*Scale
}

func (c CircleMember) String() string {
	var (
		i    int
		keys = make([]string, len(c.scales))
	)
	for k := range c.scales {
		keys[i] = k.String()
		i++
	}
	return strings.Join(keys, ", ")
}

func (c CircleMember) Keys() util.Set[Key] {
	return util.NewSet(slices.Collect(maps.Keys(c.scales))...)
}

func (c CircleMember) Get(key Key) (*Scale, bool) {
	v, ok := c.scales[key]
	return v, ok
}

func (c CircleMember) Head() *Scale {
	return slices.Collect(maps.Values(c.scales))[0]
}

func NewCircleMember(ss ...*Scale) CircleMember {
	scales := map[Key]*Scale{}
	for _, x := range ss {
		scales[x.Key] = x
	}
	return CircleMember{
		scales: scales,
	}
}

func NewCircle(r util.Ring[CircleMember]) Circle {
	return Circle{
		r: r,
	}
}

type Circle struct {
	r util.Ring[CircleMember]
}

func (c Circle) Index(key Key) (int, bool) {
	for i := range c.r.Len() {
		if c.r.At(i).Keys().In(key) {
			return i, true
		}
	}
	return -1, false
}

func (c Circle) At(i int) CircleMember {
	return c.r.At(i)
}

func (c Circle) All() iter.Seq[CircleMember] {
	return c.r.All()
}

type circleMemberSeed []string

func (c circleMemberSeed) member() CircleMember {
	xs := make([]*Scale, len(c))
	for i, x := range c {
		xs[i] = MustNewScale(MustParseKey(x))
	}
	return NewCircleMember(xs...)
}

type circleSeed []circleMemberSeed

func (c circleSeed) circle() Circle {
	xs := make([]CircleMember, len(c))
	for i, x := range c {
		xs[i] = x.member()
	}
	return NewCircle(util.MustNewRing(xs...))
}

func newMajorCircle() Circle {
	seeds := []circleMemberSeed{
		circleMemberSeed([]string{"C"}),
		circleMemberSeed([]string{"G"}),
		circleMemberSeed([]string{"D"}),
		circleMemberSeed([]string{"A"}),
		circleMemberSeed([]string{"E"}),
		circleMemberSeed([]string{"B", "Cb"}),
		circleMemberSeed([]string{"Gb", "F#"}),
		circleMemberSeed([]string{"Db", "C#"}),
		circleMemberSeed([]string{"Ab"}),
		circleMemberSeed([]string{"Eb"}),
		circleMemberSeed([]string{"Bb"}),
		circleMemberSeed([]string{"F"}),
	}
	return circleSeed(seeds).circle()
}

func newMinorCircle() Circle {
	seeds := []circleMemberSeed{
		circleMemberSeed([]string{"Am"}),
		circleMemberSeed([]string{"Em"}),
		circleMemberSeed([]string{"Bm"}),
		circleMemberSeed([]string{"F#m"}),
		circleMemberSeed([]string{"C#m"}),
		circleMemberSeed([]string{"G#m"}),
		circleMemberSeed([]string{"Ebm", "D#m"}),
		circleMemberSeed([]string{"Bbm"}),
		circleMemberSeed([]string{"Fm"}),
		circleMemberSeed([]string{"Cm"}),
		circleMemberSeed([]string{"Gm"}),
		circleMemberSeed([]string{"Dm"}),
	}
	return circleSeed(seeds).circle()
}

type CircleOfFifth struct {
	Minors Circle
	Majors Circle
}

func NewCircleOfFifth() CircleOfFifth {
	return CircleOfFifth{
		Minors: newMinorCircle(),
		Majors: newMajorCircle(),
	}
}

func (c CircleOfFifth) index(key Key) (int, error) {
	if key.Minor {
		if x, ok := c.Minors.Index(key); ok {
			return x, nil
		}
	} else {
		if x, ok := c.Majors.Index(key); ok {
			return x, nil
		}
	}
	return -1, errorx.NotFound("%v", key)
}

func (c CircleOfFifth) find(key Key, isMinor bool, indexDelta int) (CircleMember, error) {
	index, err := c.index(key)
	if err != nil {
		var defaultValue CircleMember
		return defaultValue, err
	}

	if isMinor {
		return c.Minors.At(index + indexDelta), nil
	}
	return c.Majors.At(index + indexDelta), nil
}

// Parallel returns a key that has the same tonic but with the major and minor modes swapped.
func (c CircleOfFifth) Parallel(key Key) (CircleMember, error) {
	var delta int
	if key.Minor {
		delta = 3
	} else {
		delta = -3
	}

	r, err := c.find(key, !key.Minor, delta)
	if err != nil {
		return r, fmt.Errorf("%w: Parallel of %v", err, key)
	}
	return r, nil
}

// Relative returns a key that has the same signature but with the major and minor modes swapped.
func (c CircleOfFifth) Relative(key Key) (CircleMember, error) {
	r, err := c.find(key, !key.Minor, 0)
	if err != nil {
		return r, fmt.Errorf("%w: Relative of %v", err, key)
	}
	return r, nil
}

func (c CircleOfFifth) Dominant(key Key) (CircleMember, error) {
	r, err := c.find(key, key.Minor, 1)
	if err != nil {
		return r, fmt.Errorf("%w: Dominant of %v", err, key)
	}
	return r, nil
}

func (c CircleOfFifth) SubDominant(key Key) (CircleMember, error) {
	r, err := c.find(key, key.Minor, -1)
	if err != nil {
		return r, fmt.Errorf("%w: SubDominant of %v", err, key)
	}
	return r, nil
}

//go:generate go run golang.org/x/tools/cmd/stringer -type KeyConversion -output circle_stringer_generated.go
type KeyConversion int

const (
	UnknownKeyConversion KeyConversion = iota
	ParallelKey
	RelativeKey
	DominantKey
	SubDominantKey
)

func (k KeyConversion) Converter(c CircleOfFifth) func(Key) (CircleMember, error) {
	switch k {
	case ParallelKey:
		return c.Parallel
	case RelativeKey:
		return c.Relative
	case DominantKey:
		return c.Dominant
	case SubDominantKey:
		return c.SubDominant
	default:
		return func(_ Key) (CircleMember, error) {
			var m CircleMember
			return m, errorx.Unexpected("Unknown KeyConversion")
		}
	}
}

type KeyConversionChain []KeyConversion

func (cc KeyConversionChain) Convert(c CircleOfFifth, key Key) (CircleMember, error) {
	scale, err := NewScale(key)
	if err != nil {
		var m CircleMember
		return m, err
	}

	var (
		m    = NewCircleMember(scale)
		rErr error
	)
	for i, x := range cc {
		for k := range m.Keys().All() {
			ks, _ := m.Get(k)
			next, err := x.Converter(c)(ks.Key)

			slog.Debug("KeyConversion",
				slog.String("conversion", x.String()),
				slog.String("from", k.String()),
				slog.String("to", next.String()),
				logx.Err(err),
			)
			if err == nil {
				m = next
				rErr = nil
				break
			}

			rErr = fmt.Errorf("%w: key conversion[%d]", err, i)
		}

		if rErr != nil {
			var m CircleMember
			return m, rErr
		}
	}

	return m, nil
}
