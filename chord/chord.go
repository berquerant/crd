package chord

import "github.com/berquerant/crd/note"

//go:generate stringer -type Chord -output chord_stringer_generated.go

type Chord int

const (
	MajorTriad Chord = iota
	MinorTriad
	DiminishedTriad
	AugmentedTriad
	DominantSeventh
	MajorSeventh
	MinorSeventh
	MinorMajorSeventh
	DiminishedSeventh
	HalfDiminishedSeventh
	AugmentedSeventh
	AugmentedMajorSeventh
	AddSixth
	AddMinorSixth
	SuspendedForth
)

// Semitones returns the components of the chord as a list of semitones.
func (s Chord) Semitones() []note.Semitone {
	switch s {
	case MajorTriad:
		return []note.Semitone{0, 4, 7}
	case MinorTriad:
		return []note.Semitone{0, 3, 7}
	case DiminishedTriad:
		return []note.Semitone{0, 3, 6}
	case AugmentedTriad:
		return []note.Semitone{0, 4, 8}
	case DominantSeventh:
		return []note.Semitone{0, 4, 7, 10}
	case MajorSeventh:
		return []note.Semitone{0, 4, 7, 11}
	case MinorSeventh:
		return []note.Semitone{0, 3, 7, 10}
	case MinorMajorSeventh:
		return []note.Semitone{0, 3, 7, 11}
	case DiminishedSeventh:
		return []note.Semitone{0, 3, 6, 9}
	case HalfDiminishedSeventh:
		return []note.Semitone{0, 3, 6, 10}
	case AugmentedSeventh:
		return []note.Semitone{0, 4, 8, 10}
	case AugmentedMajorSeventh:
		return []note.Semitone{0, 4, 8, 11}
	case AddSixth:
		return []note.Semitone{0, 4, 7, 9}
	case AddMinorSixth:
		return []note.Semitone{0, 3, 7, 9}
	case SuspendedForth:
		return []note.Semitone{0, 5, 7}
	default:
		panic("Unknown Chord!")
	}
}

func (s Chord) SPNs(root note.SPN) []note.SPN {
	semitones := s.Semitones()
	spns := make([]note.SPN, len(semitones))
	for i, st := range semitones {
		spns[i] = (root.Semitone() + st).SPN()
	}
	return spns
}
