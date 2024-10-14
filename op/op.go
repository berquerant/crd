package op

import (
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/note"
)

type Instance struct {
	Chord  *Chord       `yaml:"chord,omitempty"`
	Values []note.Value `yaml:"values"`

	BPM      *BPM         `yaml:"bpm,omitempty"`
	Velocity *DynamicSign `yaml:"velocity,omitempty"`
	Meter    *Meter       `yaml:"meter,omitempty"`
	Key      *Key         `yaml:"key,omitempty"`
	Meta     *Meta        `yaml:"meta,omitempty"`
}

func (i Instance) IsRest() bool {
	return i.Chord == nil
}

func (i Instance) Validate() error {
	if len(i.Values) == 0 {
		return errorx.Invalid("Instance should have values")
	}
	return nil
}

type Meta map[string]string

func (m Meta) Get(key string) string {
	return m[key]
}

func (m Meta) Set(key, value string) {
	m[key] = value
}

func NewMeta(keyValues ...string) *Meta {
	d := map[string]string{}
	var i int
	for i < len(keyValues) {
		k := keyValues[i]
		i++
		v := keyValues[i]
		i++
		d[k] = v
	}
	x := Meta(d)
	return &x
}
