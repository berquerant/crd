package ast

import "github.com/berquerant/ybase"

type Node interface {
	IsNode()
}

type ChordList struct {
	List []ChordOrRest `json:"list" yaml:"list"`
}

type ChordOrRest interface {
	IsChordOrRest()
}

type Rest struct {
	Values *ChordValues `json:"values" yaml:"values"`
	Meta   *ChordMeta   `json:"meta,omitempty" yaml:"meta,omitempty"`
}

func (*Rest) IsChordOrRest() {}

type Chord struct {
	Degree *ChordDegree `json:"degree" yaml:"degree"`
	Symbol *ChordSymbol `json:"symbol,omitempty" yaml:"symbol,omitempty"`
	Base   *ChordBase   `json:"base,omitempty" yaml:"base,omitempty"`
	Values *ChordValues `json:"values" yaml:"values"`
	Meta   *ChordMeta   `json:"meta,omitempty" yaml:"meta,omitempty"`
}

func (*Chord) IsChordOrRest() {}

type ChordDegree struct {
	Degree     ybase.Token `json:"degree" yaml:"degree"`
	Accidental ybase.Token `json:"accidental,omitempty" yaml:"accidental,omitempty"`
}

type ChordSymbol struct {
	Symbol ybase.Token `json:"symbol" yaml:"symbol"`
}

type ChordBase struct {
	Degree *ChordDegree `json:"degree" yaml:"degree"`
}

type ChordValues struct {
	Values []*ChordValue `json:"values" yaml:"values"`
}

type ChordValue struct {
	Num   ybase.Token `json:"num" yaml:"num"`
	Denom ybase.Token `json:"denom,omitempty" yaml:"denom,omitempty"`
}

type ChordMeta struct {
	Data []*ChordMetadata `json:"data,omitempty" yaml:"data,omitempty"`
}

type ChordMetadata struct {
	Key   ybase.Token `json:"key" yaml:"key"`
	Value ybase.Token `json:"value" yaml:"value"`
}
