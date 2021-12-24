package ast

import (
	"fmt"
	"strings"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/note"
)

type Node interface {
	IsNode()
	Semitones() []note.Semitone
	Equal(other Node) bool
}

type Score struct {
	NodeList []Node
}

func (s *Score) String() string {
	v := make([]string, len(s.NodeList))
	for i, n := range s.NodeList {
		v[i] = fmt.Sprint(n)
	}
	return strings.Join(v, " ")
}

//go:generate marker -method IsNode -type Rest,Chord -output ast_marker_generated.go

type Rest struct {
	Value note.Value
}

func (s *Rest) Semitones() []note.Semitone { return nil }
func (s *Rest) String() string             { return fmt.Sprintf("R[%s]", s.Value) }

func (s *Rest) Equal(other Node) bool {
	if x, ok := other.(*Rest); ok {
		return s.Value.Equal(x.Value)
	}
	return false
}

type Chord struct {
	ChordNote   *ChordNote
	ChordOption *ChordOption
	Value       note.Value
}

func (s *Chord) Semitones() []note.Semitone {
	c, ok := s.ChordOption.Name()
	if !ok {
		return nil
	}
	v := c.SPNs(s.ChordNote.Semitone().SPN())
	u := make([]note.Semitone, len(v))
	for i, x := range v {
		u[i] = x.Semitone()
	}
	return u
}

func (s *Chord) String() string { return fmt.Sprintf("%s%s[%s]", s.ChordNote, s.ChordOption, s.Value) }

func (s *Chord) Equal(other Node) bool {
	if x, ok := other.(*Chord); ok {
		return s.ChordNote.Equal(x.ChordNote) && s.ChordOption.Equal(x.ChordOption) && s.Value.Equal(x.Value)
	}
	return false
}

type ChordNote struct {
	Name       note.Name
	Octave     note.Octave
	Accidental note.Accidental
}

func (s *ChordNote) Semitone() note.Semitone {
	return note.NewSPN(note.NewNote(s.Name, s.Accidental), s.Octave).Semitone()
}

// TODO: apply octave
func (s *ChordNote) String() string {
	if s.Accidental == note.Natural {
		return s.Name.String()
	}
	return fmt.Sprintf("%s%s", s.Name, s.Accidental)
}

func (s *ChordNote) Equal(other *ChordNote) bool {
	return s.Name == other.Name && s.Octave == other.Octave && s.Accidental == other.Accidental
}

type ChordOption struct {
	IsMajor      bool
	IsMinor      bool
	IsSeventh    bool
	IsDiminished bool
	IsAugmented  bool
	IsSixth      bool
	IsSuspended  bool
	IsForth      bool
	Accidentaled int
}

func (s *ChordOption) Equal(other *ChordOption) bool {
	if s == nil && other == nil {
		return true
	}
	return s.IsMajor == other.IsMajor &&
		s.IsMinor == other.IsMinor &&
		s.IsSeventh == other.IsSeventh &&
		s.IsDiminished == other.IsDiminished &&
		s.IsAugmented == other.IsAugmented &&
		s.IsSixth == other.IsSixth &&
		s.IsSuspended == other.IsSuspended &&
		s.IsForth == other.IsForth &&
		s.Accidentaled == other.Accidentaled
}

var chordOptionToChord = []struct {
	option ChordOption
	name   chord.Chord
}{
	{
		option: ChordOption{},
		name:   chord.MajorTriad,
	},
	{
		option: ChordOption{IsMinor: true},
		name:   chord.MinorTriad,
	},
	{
		option: ChordOption{IsDiminished: true},
		name:   chord.DiminishedTriad,
	},
	{
		option: ChordOption{IsAugmented: true},
		name:   chord.AugmentedTriad,
	},
	{
		option: ChordOption{IsSeventh: true},
		name:   chord.DominantSeventh,
	},
	{
		option: ChordOption{IsSeventh: true, IsMinor: true},
		name:   chord.MinorSeventh,
	},
	{
		option: ChordOption{IsMajor: true, IsSeventh: true},
		name:   chord.MajorSeventh,
	},
	{
		option: ChordOption{IsMinor: true, IsMajor: true, IsSeventh: true},
		name:   chord.MinorMajorSeventh,
	},
	{
		option: ChordOption{IsDiminished: true, IsSeventh: true},
		name:   chord.DiminishedSeventh,
	},
	{
		option: ChordOption{IsMinor: true, IsSeventh: true, Accidentaled: -5},
		name:   chord.HalfDiminishedSeventh,
	},
	{
		option: ChordOption{IsAugmented: true, IsSeventh: true},
		name:   chord.AugmentedSeventh,
	},
	{
		option: ChordOption{IsMajor: true, IsSeventh: true, IsAugmented: true},
		name:   chord.AugmentedMajorSeventh,
	},
	{
		option: ChordOption{IsSixth: true},
		name:   chord.AddSixth,
	},
	{
		option: ChordOption{IsSixth: true, IsMinor: true},
		name:   chord.AddMinorSixth,
	},
	{
		option: ChordOption{IsSuspended: true, IsForth: true},
		name:   chord.SuspendedForth,
	},
}

func (s *ChordOption) Name() (chord.Chord, bool) {
	for _, c := range chordOptionToChord {
		if s.Equal(&c.option) {
			return c.name, true
		}
	}
	return chord.MajorTriad, false
}

func (s *ChordOption) String() string {
	var (
		b     strings.Builder
		write = func(x string) { b.WriteString(x) }
	)
	if s.IsMinor {
		write("m")
	}
	if s.IsMajor {
		write("M")
	}
	if s.IsDiminished {
		write("dim")
	}
	if s.IsAugmented {
		write("aug")
	}
	if s.IsSuspended {
		write("sus")
	}
	if s.IsForth {
		write("4")
	}
	if s.IsSixth {
		write("6")
	}
	if s.IsSeventh {
		write("7")
	}
	switch {
	case s.Accidentaled > 0:
		write(fmt.Sprintf("+%d", s.Accidentaled))
	case s.Accidentaled < 0:
		write(fmt.Sprint(s.Accidentaled))
	}
	return b.String()
}
