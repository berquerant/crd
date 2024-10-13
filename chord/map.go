package chord

import (
	"errors"

	"github.com/berquerant/crd/errorx"
)

// Mapper is a Chord dictionary.
type Mapper interface {
	GetChord(nameOrDisplay string) (Chord, bool)
	GetChordAttributes(nameOrDiaplay string) ([]Attribute, bool)
	GetAttribute(name string) (Attribute, bool)
}

var (
	_ Mapper = &Map{}
)

func NewMap(
	attributes map[string]Attribute,
	chords map[string]Chord,
) (*Map, error) {
	m := &Map{
		attributes: attributes,
		chords:     chords,
	}
	if err := m.validate(); err != nil {
		return nil, err
	}
	return m, nil
}

type Map struct {
	attributes map[string]Attribute
	chords     map[string]Chord
}

func (m Map) validate() error {
	errs := []error{}
	for _, c := range m.chords {
		for _, a := range c.Attributes {
			if _, ok := m.attributes[a]; !ok {
				errs = append(errs, errorx.Invalid("Chord %s Attribute %s not found", c.Name, a))
			}
		}
		if x := c.Extends; x != "" {
			if _, ok := m.chords[x]; !ok {
				errs = append(errs, errorx.Invalid("Chord %s Extends %s not found", c.Name, x))
			}
		}
	}

	return errors.Join(errs...)
}

func (m Map) GetChord(nameOrDisplay string) (Chord, bool) {
	c, ok := m.chords[nameOrDisplay]
	return c, ok
}

func (m Map) GetChordAttributes(nameOrDisplay string) ([]Attribute, bool) {
	c, ok := m.chords[nameOrDisplay]
	if !ok {
		return nil, false
	}

	var attrs []Attribute
	if x := c.Extends; x != "" {
		if xs, ok := m.GetChordAttributes(x); ok {
			attrs = append(attrs, xs...)
		}
	}

	for _, a := range c.Attributes {
		attrs = append(attrs, m.attributes[a])
	}
	return attrs, true
}

func (m Map) GetAttribute(name string) (Attribute, bool) {
	a, ok := m.attributes[name]
	return a, ok
}
