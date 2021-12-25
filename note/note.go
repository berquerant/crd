package note

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type Accidental int

const (
	Natural Accidental = iota
	Sharp
	Flat
	DoubleSharp
	DoubleFlat
)

func (s Accidental) String() string {
	switch s {
	case Natural:
		return "♮"
	case Sharp:
		return "♯"
	case Flat:
		return "♭"
	case DoubleSharp:
		return "𝄪"
	case DoubleFlat:
		return "𝄫"
	default:
		panic("Unknown Accidental!")
	}
}

//go:generate stringer -type Name -output name_stringer_generated.go

// Name is a musical note name.
type Name int

const (
	C Name = iota
	D
	E
	F
	G
	A
	B
)

// Semitone returns a number of the semutones between C and this name.
func (s Name) Semitone() Semitone {
	switch s {
	case C:
		return 0
	case D:
		return 2
	case E:
		return 4
	case F:
		return 5
	case G:
		return 7
	case A:
		return 9
	case B:
		return 11
	default:
		panic("Unknown Name!")
	}
}

type Semitone int

// Sign applies an accidental to self.
func (s Semitone) Sign(a Accidental) Semitone {
	switch a {
	case Natural:
		return s
	case Sharp:
		return s + 1
	case Flat:
		return s - 1
	case DoubleSharp:
		return s + 2
	case DoubleFlat:
		return s - 2
	default:
		panic("Unknown Accidental!")
	}
}

// Octave extracts the octave.
func (s Semitone) Octave() Octave {
	i := int(s)
	if i > 0 {
		return Octave(i / 12)
	}
	return Octave(int(s)/12) - 1
}

// Reminder extracts the number of the semitones within the octave.
func (s Semitone) Reminder() Semitone {
	i := int(s)
	if i > 0 {
		return s % 12
	}
	return s%12 + 12
}

func (s Semitone) NoteWithAccidental(accidental Accidental) (Note, Octave) {
	p := selectNameToDisplay(s, accidental)
	return NewNote(p.name, p.accidental), p.octave
}

func (s Semitone) Note() Note {
	n, _ := s.NoteWithAccidental(Natural) // octave diff must be 0
	return n
}

func (s Semitone) SPNWithAccidental(accidental Accidental) SPN {
	n, oct := s.NoteWithAccidental(accidental)
	return NewSPN(n, oct)
}

// SPN converts semitones as a SPN.
func (s Semitone) SPN() SPN { return s.SPNWithAccidental(Natural) }

// Note represents a musical note.
type Note interface {
	Name() Name
	Accidental() Accidental
	Semitone() Semitone
	Equal(other Note) bool
}

type note struct {
	name       Name
	accidental Accidental
}

func NewNote(name Name, accidental Accidental) Note {
	return &note{
		name:       name,
		accidental: accidental,
	}
}

func (s *note) Name() Name             { return s.name }
func (s *note) Accidental() Accidental { return s.accidental }
func (s *note) Semitone() Semitone     { return s.name.Semitone().Sign(s.accidental) }
func (s *note) String() string         { return fmt.Sprintf("%s%s", s.name, s.accidental) }
func (s *note) Equal(other Note) bool {
	return s.Name() == other.Name() && s.Accidental() == other.Accidental()
}
func (s *note) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":       s.name,
		"accidental": s.accidental,
	})
}

type Octave int

// SPN represents Scientific Pitch Notation.
type SPN interface {
	Octave() Octave
	Note() Note
	Semitone() Semitone
	Equal(other SPN) bool
	// LowerBound returns the highest SPN with target that is lower than self.
	LowerBound(target Note) SPN
	// UpperBound returns the lowest SPN with target that is higher than self.
	UpperBound(target Note) SPN
}

type spn struct {
	octave Octave
	note   Note
}

func NewSPN(note Note, octave Octave) SPN {
	return &spn{
		octave: octave,
		note:   note,
	}
}

func (s *spn) Octave() Octave { return s.octave }
func (s *spn) Note() Note     { return s.note }
func (s *spn) Equal(other SPN) bool {
	return s.Octave() == other.Octave() && s.Note().Equal(other.Note())
}
func (s *spn) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"octave": s.octave,
		"note":   s.note,
	})
}

func (s *spn) UpperBound(target Note) SPN {
	var (
		i     Octave = 9 // from midi highest octave
		bound        = s.Semitone()
	)
	for {
		x := NewSPN(target, i-1)
		if x.Semitone() <= bound {
			return NewSPN(target, i)
		}
		i--
	}
}

func (s *spn) LowerBound(target Note) SPN {
	var (
		i     Octave
		bound = s.Semitone()
	)
	for {
		x := NewSPN(target, i+1)
		if x.Semitone() >= bound {
			return NewSPN(target, i)
		}
		i++
	}
}

// Semitone returns a number of semitones when C0 is 0.
func (s *spn) Semitone() Semitone { return Semitone(12*s.octave) + s.note.Semitone() }

func (s *spn) String() string { return fmt.Sprintf("%s%d", s.note, s.octave) }

type (
	// Value represents a note value.
	Value interface {
		Beat() Beat
		Raw() *big.Rat
		Equal(other Value) bool
	}
	Beat interface {
		Value() Value
		Raw() *big.Rat
		Equal(other Beat) bool
	}
)

type (
	value struct {
		*big.Rat
	}
	beat struct {
		*big.Rat
	}
)

func NewValue(r *big.Rat) Value { return &value{r} }

func (s *value) Beat() Beat {
	return &beat{
		new(big.Rat).Mul(
			s.Rat,
			new(big.Rat).SetInt64(4),
		),
	}
}
func (s *value) Raw() *big.Rat          { return s.Rat }
func (s *value) String() string         { return s.Rat.RatString() }
func (s *value) Equal(other Value) bool { return s.Raw().RatString() == other.Raw().RatString() }

func NewBeat(r *big.Rat) Beat { return &beat{r} }

func (s *beat) Value() Value {
	return &value{
		new(big.Rat).Mul(
			s.Rat,
			big.NewRat(1, 4),
		),
	}
}
func (s *beat) Raw() *big.Rat         { return s.Rat }
func (s *beat) String() string        { return s.Rat.RatString() }
func (s *beat) Equal(other Beat) bool { return s.Raw().RatString() == other.Raw().RatString() }

// Key represents a musical key.
type Key interface {
	Name() Name
	Accidental() Accidental
	IsMinor() bool
	Equal(other Key) bool
}

func NewKey(name Name, accidental Accidental, isMinor bool) Key {
	return &key{
		name:       name,
		accidental: accidental,
		isMinor:    isMinor,
	}
}

type key struct {
	name       Name
	accidental Accidental
	isMinor    bool
}

func (s *key) Name() Name             { return s.name }
func (s *key) Accidental() Accidental { return s.accidental }
func (s *key) IsMinor() bool          { return s.isMinor }
func (s *key) Equal(other Key) bool {
	return s.Name() == other.Name() && s.Accidental() == other.Accidental() && s.IsMinor() == other.IsMinor()
}
func (s *key) String() string {
	var b strings.Builder
	b.WriteString(s.name.String())
	b.WriteString(s.accidental.String())
	if s.isMinor {
		b.WriteString("minor")
	} else {
		b.WriteString("major")
	}
	return b.String()
}
func (s *key) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":       s.Name,
		"accidental": s.accidental,
		"isMinor":    s.isMinor,
	})
}
