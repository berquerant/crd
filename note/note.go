package note

import (
	"fmt"
	"regexp"

	"github.com/berquerant/crd/errorx"
)

type Note struct {
	Name       Name
	Accidental Accidental
}

func NewNote(name Name, accidental Accidental) Note {
	return Note{
		Name:       name,
		Accidental: accidental,
	}
}

func (n Note) AddDegree(d Degree, precedeSharp bool) (Note, Octave, error) {
	ds, ok := d.Semitone()
	if !ok {
		var t Note
		return t, 0, errorx.Invalid("Degree %v", d)
	}

	s := n.Semitone() + ds
	so, oct := s.WithoutOctave(), s.Octave()
	find := func(a ...Accidental) (Note, Octave, error) {
		for _, b := range a {
			if x, ok := n.findNameBySemitone(so, b); ok {
				return Note{
					Name:       x,
					Accidental: b,
				}, oct, nil
			}
		}
		var t Note
		return t, 0, errorx.Invalid("Cannot add %v to %v", d, n)
	}

	if precedeSharp {
		return find(Natural, Sharp, Flat)
	} else {
		return find(Natural, Flat, Sharp)
	}
}

func (Note) findNameBySemitone(want Semitone, accidental Accidental) (Name, bool) {
	for _, v := range []Name{C, D, E, F, G, A, B} {
		x := v.Semitone() + accidental.Semitone()
		if x == want {
			return v, true
		}
	}
	return UnknownName, false
}

func (n Note) Semitone() Semitone {
	return n.Name.Semitone() + n.Accidental.Semitone()
}

func (n Note) String() string {
	if n.Accidental == Natural {
		return n.Name.String()
	}
	return fmt.Sprintf("%v%s", n.Name, n.Accidental.String(true))
}

func (n Note) MarshalYAML() (any, error) {
	return n.String(), nil
}

var (
	noteRegex = regexp.MustCompile(`([A-G])([#b]?)`)
)

func ParseNote(s string) (Note, error) {
	var defaultNote Note
	matched := noteRegex.FindAllStringSubmatch(s, -1)
	if len(matched) == 0 {
		return defaultNote, errorx.Invalid("Note %s", s)
	}

	m := matched[0][1:]
	name := NewName(m[0])
	if m[1] == "" {
		return Note{
			Name:       name,
			Accidental: Natural,
		}, nil
	}

	accidental := NewAccidental(m[1])

	return Note{
		Name:       name,
		Accidental: accidental,
	}, nil
}
