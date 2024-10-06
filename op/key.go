package op

import (
	"regexp"
	"strings"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/util"
	"gopkg.in/yaml.v3"
)

type Accidental int

const (
	UnknownAccidental Accidental = iota
	Natural
	Sharp
	Flat
)

// Tendency calculates the substantial temporary accidental of x with reference to a.
func (a Accidental) Tendency(x Accidental) Accidental {
	switch {
	case a == UnknownAccidental || x == UnknownAccidental:
		return UnknownAccidental
	case a == x:
		return Natural
	case a == Natural:
		return x
	case a == Sharp:
		return Sharp
	case a == Flat:
		return Flat
	default:
		return UnknownAccidental
	}
}

const (
	accidentalNatural = ""
	accidentalSharp   = "#"
	accidentalFlat    = "b"
)

var (
	accidentalStringMap = map[Accidental]string{
		Natural: accidentalNatural,
		Sharp:   accidentalSharp,
		Flat:    accidentalFlat,
	}
	stringAccidentalMap = util.MustInverseMap(accidentalStringMap)
)

func NewAccidental(s string) Accidental {
	if x, ok := stringAccidentalMap[s]; ok {
		return x
	}
	return Natural
}

func (a Accidental) AsNoteAccidental() note.Accidental {
	switch a {
	case Sharp:
		return note.Sharp
	case Flat:
		return note.Flat
	default:
		return note.Natural
	}
}

func (a Accidental) Semitone() note.Semitone {
	switch a {
	case Sharp:
		return 1
	case Flat:
		return -1
	default:
		return 0
	}
}

func (a Accidental) String() string {
	return accidentalStringMap[a]
}

func (a Accidental) MarshalYAML() (any, error) {
	return a.String(), nil
}

func (a Accidental) MarshalJSON() ([]byte, error) {
	return []byte(a.String()), nil
}

type Key struct {
	Name       note.Name  `json:"name" yaml:"name"`
	Minor      bool       `json:"minor" yaml:"minor"`
	Accidental Accidental `json:"accidental" yaml:"accidental"`
}

func (k Key) Semitone() note.Semitone {
	return k.Name.Semitone() + k.Accidental.AsNoteAccidental().Semitone()
}

func (k Key) String() string {
	ss := []string{
		k.Name.String(),
		k.Accidental.String(),
	}
	if k.Minor {
		ss = append(ss, minorKeyMark)
	}
	return strings.Join(ss, "")
}

func (k Key) MarshalYAML() (any, error) {
	return k.String(), nil
}

func (k *Key) UnmarshalYAML(value *yaml.Node) error {
	x, err := ParseKey(value.Value)
	if err != nil {
		return err
	}
	*k = x
	return nil
}

const (
	minorKeyMark = "m"
)

var (
	keyRegex = regexp.MustCompile(`([A-G])([#b]?)(m?)`)
)

func ParseKey(s string) (Key, error) {
	var (
		defaultKey Key
	)

	matched := keyRegex.FindAllStringSubmatch(s, -1)
	if len(matched) == 0 {
		return defaultKey, errorx.Invalid("Key %s", s)
	}

	m := matched[0][1:]
	name := note.NewName(m[0])
	accidental := NewAccidental(m[1])
	minor := m[2] == minorKeyMark

	return Key{
		Name:       name,
		Minor:      minor,
		Accidental: accidental,
	}, nil
}

func MustParseKey(s string) Key {
	k, err := ParseKey(s)
	logx.PanicOnError(err)
	return k
}
