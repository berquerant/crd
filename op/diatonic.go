package op

type DiatonicChorder interface {
	Sevenths() [7]DiatonicChord
	Triads() [7]DiatonicChord
}

var (
	_ DiatonicChorder = &DiatonicChorderImpl{}
)

type DiatonicChorderImpl struct {
	scale *Scale
}

func NewDiatonicChorder(scale *Scale) *DiatonicChorderImpl {
	return &DiatonicChorderImpl{
		scale: scale,
	}
}

type DiatonicChord struct {
	Note *ScaleNote
	Name string
}

func (dc DiatonicChord) String() string {
	return dc.Note.String() + dc.Name
}

func (dc DiatonicChord) MarshalYAML() (any, error) {
	return dc.String(), nil
}

func (dc DiatonicChorderImpl) Sevenths() [7]DiatonicChord {
	return dc.generate(dc.seventhNames())
}

func (dc DiatonicChorderImpl) Triads() [7]DiatonicChord {
	return dc.generate(dc.triadNames())
}

func (dc DiatonicChorderImpl) generate(names [7]string) [7]DiatonicChord {
	var (
		r     [7]DiatonicChord
		notes = dc.scale.Notes
	)
	for i, x := range notes {
		r[i] = DiatonicChord{
			Note: x,
			Name: names[i],
		}
	}
	return r
}

func (dc DiatonicChorderImpl) seventhNames() [7]string {
	if dc.scale.Key.Minor {
		return [7]string{
			"m7",
			"m7b5",
			"maj7",
			"m7",
			"m7",
			"maj7",
			"_7",
		}
	}

	return [7]string{
		"maj7",
		"m7",
		"m7",
		"maj7",
		"_7",
		"m7",
		"m7b5",
	}
}

func (dc DiatonicChorderImpl) triadNames() [7]string {
	if dc.scale.Key.Minor {
		return [7]string{
			"m",
			"dim",
			"",
			"m",
			"m",
			"",
			"",
		}
	}

	return [7]string{
		"",
		"m",
		"m",
		"",
		"",
		"m",
		"dim",
	}
}
