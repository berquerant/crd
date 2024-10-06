package note

import (
	"strconv"

	"gopkg.in/yaml.v3"
)

type Semitone int

func (s Semitone) MarshalYAML() (any, error) {
	return yaml.Marshal(int(s))
}

func (s *Semitone) UnmarshalYAML(value *yaml.Node) error {
	i, err := strconv.Atoi(value.Value)
	if err != nil {
		return err
	}
	*s = Semitone(i)
	return nil
}

const (
	octaveSemitones = 12
)

func (s Semitone) Octave() Octave {
	i := int(s)
	if i >= 0 {
		return Octave(i / octaveSemitones)
	}
	return Octave(int(s)/octaveSemitones) - 1
}

func (s Semitone) WithoutOctave() Semitone {
	i := int(s)
	if i >= 0 {
		return s % octaveSemitones
	}
	return s%octaveSemitones + octaveSemitones
}

type Octave int

func (o Octave) Semitone() Semitone {
	return Semitone(int(o) * octaveSemitones)
}
