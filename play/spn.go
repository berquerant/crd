package play

import (
	"fmt"
	"strings"

	"github.com/berquerant/crd/note"
)

var (
	// C4
	MiddleC = SPN{
		Name:       note.C,
		Octave:     4,
		Accidental: note.Natural,
	}
)

func NewSPN(name note.Name, octave note.Octave, accidental note.Accidental) SPN {
	return SPN{
		Name:       name,
		Octave:     octave,
		Accidental: accidental,
	}
}

type SPN struct {
	Name       note.Name
	Octave     note.Octave
	Accidental note.Accidental
}

func (s SPN) String() string {
	ss := []string{
		s.Name.String(),
		fmt.Sprint(s.Octave),
	}
	if s.Accidental != note.Natural {
		ss = append(ss, s.Accidental.String(false))
	}
	return strings.Join(ss, "")
}

func (s SPN) MarshalYAML() (any, error) {
	return s.String(), nil
}

func (s SPN) MIDINoteNumber() MIDINoteNumber {
	t := (s.Octave + 1).Semitone() + s.Name.Semitone() + s.Accidental.Semitone()
	return MIDINoteNumber(t)
}

type MIDINoteNumber uint8

func (m MIDINoteNumber) MarshalYAML() (any, error) {
	return fmt.Sprint(m), nil
}
