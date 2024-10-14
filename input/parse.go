package input

import (
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
)

type Instance struct {
	Chord  *Chord       `yaml:"chord,omitempty"`
	Values []note.Value `yaml:"values"`

	BPM      *op.BPM         `yaml:"bpm,omitempty"`
	Velocity *op.DynamicSign `yaml:"velocity,omitempty"`
	Meter    *op.Meter       `yaml:"meter,omitempty"`
	Key      *op.Key         `yaml:"key,omitempty"`
	Meta     *op.Meta        `yaml:"meta,omitempty"`
}

type Chord struct {
	Degree note.Degree  `yaml:"degree"`
	Chord  string       `yaml:"name"`
	Base   *note.Degree `yaml:"base,omitempty"`
}
