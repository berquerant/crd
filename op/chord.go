package op

import (
	"strings"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/note"
)

var (
	defaultChordBase, _ = note.NewDegree(1, note.PerfectDegree)
)

type Chord struct {
	Degree note.Degree `yaml:"degree"`
	Chord  chord.Chord `yaml:"chord"`
	Base   note.Degree `yaml:"base,omitempty"`
}

func NewChord(degree note.Degree, c chord.Chord, base *note.Degree) Chord {
	r := Chord{
		Degree: degree,
		Chord:  c,
	}
	if base != nil {
		r.Base = *base
	} else {
		r.Base = defaultChordBase
	}
	return r
}

func (c Chord) String() string {
	ss := []string{
		c.Degree.String(),
		c.Chord.Meta.Display,
	}
	if c.Base != defaultChordBase {
		ss = append(ss, "/"+c.Base.String())
	}
	return strings.Join(ss, "")
}
