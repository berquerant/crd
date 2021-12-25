package ast

import (
	"fmt"
	"strings"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/note"
)

type Node interface {
	IsNode()
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

//go:generate marker -method IsNode -type Key,Tempo,Meter,Rest,Chord -output ast_marker_generated.go

type Key struct {
	Key note.Key
}

func (s *Key) Equal(other Node) bool {
	if x, ok := other.(*Key); ok {
		return s.Key.Equal(x.Key)
	}
	return false
}

func (s *Key) String() string { return fmt.Sprintf("key[%s]", s.Key) }

type Tempo struct {
	BPM int
}

func (s *Tempo) String() string { return fmt.Sprintf("tempo[%d]", s.BPM) }

func (s *Tempo) Equal(other Node) bool {
	if x, ok := other.(*Tempo); ok {
		return s.BPM == x.BPM
	}
	return false
}

type Meter struct {
	Num, Denom uint8
}

func (s *Meter) String() string { return fmt.Sprintf("meter[%d/%d]", s.Num, s.Denom) }

func (s *Meter) Equal(other Node) bool {
	if x, ok := other.(*Meter); ok {
		return s.Num == x.Num && s.Denom == x.Denom
	}
	return false
}

type Rest struct {
	Value note.Value
}

func (s *Rest) String() string { return fmt.Sprintf("R[%s]", s.Value) }

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
	ChordBase   *ChordBase
}

func (s *Chord) Semitones() []note.Semitone {
	c, ok := s.ChordOption.Name()
	if !ok {
		return nil
	}
	v := c.SPNs(s.ChordNote.Semitone().SPN())
	u := make([]note.Semitone, len(v)+1)
	u[0] = s.base().Semitone()
	for i, x := range v {
		u[i+1] = x.Semitone()
	}
	return u
}

func (s *Chord) base() note.SPN {
	n := s.ChordNote.Semitone().SPN()
	if s.ChordBase == nil {
		return n.LowerBound(s.ChordNote.SPN.Note())
	}
	return n.LowerBound(s.ChordBase.Note)
}

func (s *Chord) String() string {
	if s.ChordBase != nil {
		return fmt.Sprintf("%s%s/%s[%s]", s.ChordNote, s.ChordOption, s.ChordBase, s.Value)
	}
	return fmt.Sprintf("%s%s[%s]", s.ChordNote, s.ChordOption, s.Value)
}

func (s *Chord) Equal(other Node) bool {
	if x, ok := other.(*Chord); ok {
		return s.ChordNote.Equal(x.ChordNote) && s.ChordOption.Equal(x.ChordOption) && s.Value.Equal(x.Value) && s.ChordBase.Equal(x.ChordBase)
	}
	return false
}

type ChordBase struct {
	Note note.Note
}

func (s *ChordBase) String() string {
	if s.Note.Accidental() == note.Natural {
		return s.Note.Name().String()
	}
	return fmt.Sprintf("%s%s", s.Note.Name(), s.Note.Accidental())
}
func (s *ChordBase) Equal(other *ChordBase) bool {
	if s == nil {
		return other == nil
	}
	return s.Note.Equal(other.Note)
}

type ChordNote struct {
	SPN note.SPN
}

func (s *ChordNote) Semitone() note.Semitone { return s.SPN.Semitone() }

// TODO: apply octave
func (s *ChordNote) String() string {
	if s.SPN.Note().Accidental() == note.Natural {
		return s.SPN.Note().Name().String()
	}
	return fmt.Sprintf("%s%s", s.SPN.Note().Name(), s.SPN.Note().Accidental())
}

func (s *ChordNote) Equal(other *ChordNote) bool { return s.SPN.Equal(other.SPN) }

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
