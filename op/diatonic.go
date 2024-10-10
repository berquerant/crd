package op

type DiatonicChorder struct {
	scale *Scale
}

func NewDiatonicChorder(scale *Scale) *DiatonicChorder {
	return &DiatonicChorder{
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

func (dc DiatonicChorder) Sevenths() [7]DiatonicChord {
	return dc.generate(dc.seventhNames())
}

func (dc DiatonicChorder) Triads() [7]DiatonicChord {
	return dc.generate(dc.triadNames())
}

func (dc DiatonicChorder) generate(names [7]string) [7]DiatonicChord {
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

func (dc DiatonicChorder) seventhNames() [7]string {
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

func (dc DiatonicChorder) triadNames() [7]string {
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
