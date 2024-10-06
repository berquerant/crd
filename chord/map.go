package chord

import (
	"errors"

	"github.com/berquerant/crd/errorx"
)

// Mapper is a Chord dictionary.
type Mapper interface {
	Get(nameOrDisplay string) (Chord, bool)
	GetAttributes(nameOrDiaplay string) ([]Attribute, bool)
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

func (m Map) Get(nameOrDisplay string) (Chord, bool) {
	c, ok := m.chords[nameOrDisplay]
	return c, ok
}

func (m Map) GetAttributes(nameOrDisplay string) ([]Attribute, bool) {
	c, ok := m.chords[nameOrDisplay]
	if !ok {
		return nil, false
	}

	var attrs []Attribute
	if x := c.Extends; x != "" {
		if xs, ok := m.GetAttributes(x); ok {
			attrs = append(attrs, xs...)
		}
	}

	for _, a := range c.Attributes {
		attrs = append(attrs, m.attributes[a])
	}
	return attrs, true
}
