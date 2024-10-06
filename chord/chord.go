package chord

import (
	_ "embed"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"gopkg.in/yaml.v3"
)

type Chord struct {
	Name string   `yaml:"name"`
	Meta Metadata `yaml:"meta"`
	// attribute names
	Attributes []string `yaml:"attributes,omitempty"`
	// chord name
	Extends string `yaml:"extends,omitempty"`
}

func (c Chord) validate() error {
	if c.Name == "" {
		return errorx.Invalid("Chord should have name")
	}
	if c.Meta.Display == "" && c.Name != "MajorTriad" {
		return errorx.Invalid("Chord should have display name except major triad")
	}
	if len(c.Attributes) == 0 && c.Extends == "" {
		return errorx.Invalid("Chord should have attributes or extends")
	}
	return nil
}

type Metadata struct {
	Display string `yaml:"display"`
}

//go:embed chord.yml
var basicChords []byte

func BasicChords() []Chord {
	chords, err := ParseChords(basicChords)
	logx.PanicOnError(err)
	return chords
}

func ParseChords(b []byte) ([]Chord, error) {
	var chords []Chord
	if err := yaml.Unmarshal(b, &chords); err != nil {
		return nil, err
	}
	for _, c := range chords {
		if err := c.validate(); err != nil {
			return nil, err
		}
	}
	return chords, nil
}
