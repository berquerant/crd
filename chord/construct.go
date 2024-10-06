package chord

type Builder struct {
	attrs  []Attribute
	chords []Chord
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Attribute(a Attribute) { b.attrs = append(b.attrs, a) }
func (b *Builder) Chord(c Chord)         { b.chords = append(b.chords, c) }

func (b Builder) Build() (*Map, error) {
	var (
		attrs  = map[string]Attribute{}
		chords = map[string]Chord{}
	)

	for _, a := range b.attrs {
		attrs[a.Name] = a
	}
	for _, c := range b.chords {
		chords[c.Name] = c
		chords[c.Meta.Display] = c
	}

	return NewMap(attrs, chords)
}

func (b Builder) UnwrapAttributes() []Attribute { return b.attrs }
func (b Builder) UnwrapChords() []Chord         { return b.chords }
