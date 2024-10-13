package desc

import (
	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/note"
)

type Attribute struct {
	mapper chord.Mapper
}

func NewAttribute(mapper chord.Mapper) *Attribute {
	return &Attribute{
		mapper: mapper,
	}
}

type AttributeInfo struct {
	Attribute             chord.Attribute `yaml:"attribute"`
	Semitone              note.Semitone   `yaml:"semitone"`
	SemitoneWithoutOctave note.Semitone   `yaml:"semitone_without_octave"`
	Root                  note.Note       `yaml:"root"`
	Applied               note.Note       `yaml:"applied"`
	OctaveDiff            note.Octave     `yaml:"octave_diff,omitempty"`
}

func (a Attribute) Describe(name string, root note.Note, precedeSharp bool) (*AttributeInfo, error) {
	attr, ok := a.mapper.GetAttribute(name)
	if !ok {
		return nil, errorx.NotFound("Attribute %s", name)
	}
	semi, _ := attr.Semitone()
	applied, octDiff, err := root.AddDegree(attr.Degree, precedeSharp)
	if err != nil {
		return nil, err
	}

	return &AttributeInfo{
		Attribute:             attr,
		Semitone:              semi,
		SemitoneWithoutOctave: semi.WithoutOctave(),
		Root:                  root,
		Applied:               applied,
		OctaveDiff:            octDiff,
	}, nil
}
