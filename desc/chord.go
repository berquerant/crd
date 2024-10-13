package desc

import (
	"fmt"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/note"
)

type Chord struct {
	mapper chord.Mapper
	attr   *Attribute
}

func NewChord(mapper chord.Mapper, attr *Attribute) *Chord {
	return &Chord{
		mapper: mapper,
		attr:   attr,
	}
}

type ChordInfo struct {
	Chord      chord.Chord      `yaml:"chord"`
	Root       note.Note        `yaml:"root"`
	Attributes []*AttributeInfo `yaml:"attributes"`
}

func (c Chord) Describe(nameOrDisplay string, root note.Note, precedeSharp bool) (*ChordInfo, error) {
	cd, ok := c.mapper.GetChord(nameOrDisplay)
	if !ok {
		return nil, errorx.NotFound("Chord %s", nameOrDisplay)
	}

	cdAttrs, _ := c.mapper.GetChordAttributes(nameOrDisplay)

	attrs := make([]*AttributeInfo, len(cdAttrs))
	for i, a := range cdAttrs {
		x, err := c.attr.Describe(a.Name, root, precedeSharp)
		if err != nil {
			return nil, fmt.Errorf("%w: Describe Chord %s", err, nameOrDisplay)
		}
		attrs[i] = x
	}

	return &ChordInfo{
		Chord:      cd,
		Root:       root,
		Attributes: attrs,
	}, nil
}
