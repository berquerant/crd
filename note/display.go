package note

type nameAndDiff struct {
	name Name
	diff Octave
}

var semitoneWithAccidentalMap = map[Semitone]map[Accidental]*nameAndDiff{
	0: { // B#, C, Dbb
		Sharp:      {name: B, diff: -1},
		Natural:    {name: C},
		DoubleFlat: {name: D},
	},
	1: { // B##, C#, Db
		DoubleSharp: {name: B, diff: -1},
		Sharp:       {name: C},
		Flat:        {name: D},
	},
	2: { // C##, D, Ebb
		DoubleSharp: {name: C},
		Natural:     {name: D},
		DoubleFlat:  {name: E},
	},
	3: { // D#, Eb, Fbb
		Sharp:      {name: D},
		Flat:       {name: E},
		DoubleFlat: {name: F},
	},
	4: { // D##, E, Fb
		DoubleSharp: {name: D},
		Natural:     {name: E},
		Flat:        {name: F},
	},
	5: { // E#, F, Gbb
		Sharp:      {name: E},
		Natural:    {name: F},
		DoubleFlat: {name: G},
	},
	6: { // E##, F#, Gb
		DoubleSharp: {name: E},
		Sharp:       {name: F},
		Flat:        {name: G},
	},
	7: { // F##, G, Abb
		DoubleSharp: {name: F},
		Natural:     {name: G},
		DoubleFlat:  {name: A},
	},
	8: { // G#, Ab
		Sharp: {name: G},
		Flat:  {name: A},
	},
	9: { // G##, A, Bbb
		DoubleSharp: {name: G},
		Natural:     {name: A},
		DoubleFlat:  {name: B},
	},
	10: { // A#, Bb, Cbb
		Sharp:      {name: A},
		Flat:       {name: B},
		DoubleFlat: {name: C, diff: 1},
	},
	11: { // A##, B, Cb
		DoubleSharp: {name: A},
		Natural:     {name: B},
		Flat:        {name: C, diff: 1},
	},
}

var accidentalPriorities = map[Accidental][]Accidental{
	Natural:     {Natural, Sharp, Flat, DoubleSharp, DoubleFlat},
	Sharp:       {Sharp, Natural, DoubleSharp, Flat, DoubleFlat},
	DoubleSharp: {DoubleSharp, Sharp, Natural, Flat, DoubleFlat},
	Flat:        {Flat, Natural, DoubleFlat, Sharp, DoubleSharp},
	DoubleFlat:  {DoubleFlat, Flat, Natural, Sharp, DoubleSharp},
}

type displayPair struct {
	name       Name
	octave     Octave
	accidental Accidental
}

func selectNameToDisplay(semitone Semitone, accidental Accidental) *displayPair {
	sem := semitone.Reminder()
	for _, a := range accidentalPriorities[accidental] {
		if x, exist := semitoneWithAccidentalMap[sem][a]; exist {
			return &displayPair{
				name:       x.name,
				octave:     semitone.Octave() + x.diff,
				accidental: a,
			}
		}
	}
	panic("Cannot find name to display!")
}
