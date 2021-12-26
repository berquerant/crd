package midi_test

import (
	"math/big"
	"testing"

	"github.com/berquerant/crd/ast"
	"github.com/berquerant/crd/midi"
	"github.com/berquerant/crd/note"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

type (
	BPMArg struct {
		BPM int
	}
	MeterArg struct {
		Num, Denom uint8
	}
	TextArg struct {
		Text string
	}
	VelocityArg struct {
		Velocity uint8
	}
	InstrumentArg struct {
		Name string
	}
	AppendArg struct {
		Tones []note.Semitone
		Value note.Value
	}
	RestArg struct {
		Value note.Value
	}
	KeyArg struct {
		Name       note.Name
		Accidental note.Accidental
		IsMinor    bool
	}
)

type MockWriter struct {
	BPMArg        BPMArg
	MeterArg      MeterArg
	TextArg       TextArg
	VelocityArg   VelocityArg
	InstrumentArg InstrumentArg
	AppendArg     AppendArg
	RestArg       RestArg
	KeyArg        KeyArg
}

func (s *MockWriter) BPM(bpm int)             { s.BPMArg = BPMArg{bpm} }
func (s *MockWriter) Meter(num, denom uint8)  { s.MeterArg = MeterArg{num, denom} }
func (s *MockWriter) Text(text string)        { s.TextArg = TextArg{text} }
func (s *MockWriter) Velocity(velocity uint8) { s.VelocityArg = VelocityArg{velocity} }
func (s *MockWriter) Instrument(name string)  { s.InstrumentArg = InstrumentArg{name} }
func (s *MockWriter) Append(tones []note.Semitone, value note.Value) {
	s.AppendArg = AppendArg{tones, value}
}
func (s *MockWriter) Rest(value note.Value) { s.RestArg = RestArg{value} }
func (s *MockWriter) Key(name note.Name, accidental note.Accidental, isMinor bool) {
	s.KeyArg = KeyArg{name, accidental, isMinor}
}
func (*MockWriter) Operations() []midi.Operation { return nil }

func TestASTWriter(t *testing.T) {
	for _, tc := range []*struct {
		name string
		node ast.Node
		want MockWriter
	}{
		{
			name: "chord C/D trans",
			node: &ast.Chord{
				ChordNote: &ast.ChordNote{
					SPN: note.NewSPN(note.NewNote(note.C, note.Natural), 4),
				},
				ChordOption: &ast.ChordOption{},
				ChordBase: &ast.ChordBase{
					Note: note.NewNote(note.D, note.Natural),
				},
				Value: note.NewValue(big.NewRat(1, 1)),
			},
			want: MockWriter{
				AppendArg: AppendArg{
					Tones: []note.Semitone{40, 50, 54, 57},
					Value: note.NewValue(big.NewRat(1, 1)),
				},
				TextArg: TextArg{
					Text: "DonE",
				},
			},
		},
		{
			name: "chord C# trans",
			node: &ast.Chord{
				ChordNote: &ast.ChordNote{
					SPN: note.NewSPN(note.NewNote(note.C, note.Sharp), 4),
				},
				ChordOption: &ast.ChordOption{},
				Value:       note.NewValue(big.NewRat(1, 1)),
			},
			want: MockWriter{
				AppendArg: AppendArg{
					Tones: []note.Semitone{39, 51, 55, 58},
					Value: note.NewValue(big.NewRat(1, 1)),
				},
				TextArg: TextArg{
					Text: "D♯",
				},
			},
		},
		{
			name: "chord C trans",
			node: &ast.Chord{
				ChordNote: &ast.ChordNote{
					SPN: note.NewSPN(note.NewNote(note.C, note.Natural), 4),
				},
				ChordOption: &ast.ChordOption{},
				Value:       note.NewValue(big.NewRat(1, 1)),
			},
			want: MockWriter{
				AppendArg: AppendArg{
					Tones: []note.Semitone{38, 50, 54, 57},
					Value: note.NewValue(big.NewRat(1, 1)),
				},
				TextArg: TextArg{
					Text: "D",
				},
			},
		},
		{
			name: "key trans",
			node: &ast.Key{
				Key: note.NewKey(note.F, note.Natural, false),
			},
			want: MockWriter{
				KeyArg: KeyArg{
					Name:       note.G,
					Accidental: note.Natural,
					IsMinor:    false,
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var w MockWriter
			aw := midi.NewASTWriter(&w)
			aw.WriteNode(&ast.Transposition{
				Semitone: 2,
			})
			aw.WriteNode(tc.node)
			assert.Equal(t, "", cmp.Diff(tc.want, w))
		})
	}
	for _, tc := range []*struct {
		name string
		node ast.Node
		want MockWriter
	}{
		{
			name: "chord C/D",
			node: &ast.Chord{
				ChordNote: &ast.ChordNote{
					SPN: note.NewSPN(note.NewNote(note.C, note.Natural), 4),
				},
				ChordOption: &ast.ChordOption{},
				ChordBase: &ast.ChordBase{
					Note: note.NewNote(note.D, note.Natural),
				},
				Value: note.NewValue(big.NewRat(1, 1)),
			},
			want: MockWriter{
				AppendArg: AppendArg{
					Tones: []note.Semitone{38, 48, 52, 55},
					Value: note.NewValue(big.NewRat(1, 1)),
				},
				TextArg: TextArg{
					Text: "ConD",
				},
			},
		},
		{
			name: "chord C#",
			node: &ast.Chord{
				ChordNote: &ast.ChordNote{
					SPN: note.NewSPN(note.NewNote(note.C, note.Sharp), 4),
				},
				ChordOption: &ast.ChordOption{},
				Value:       note.NewValue(big.NewRat(1, 1)),
			},
			want: MockWriter{
				AppendArg: AppendArg{
					Tones: []note.Semitone{37, 49, 53, 56},
					Value: note.NewValue(big.NewRat(1, 1)),
				},
				TextArg: TextArg{
					Text: "C♯",
				},
			},
		},
		{
			name: "chord C",
			node: &ast.Chord{
				ChordNote: &ast.ChordNote{
					SPN: note.NewSPN(note.NewNote(note.C, note.Natural), 4),
				},
				ChordOption: &ast.ChordOption{},
				Value:       note.NewValue(big.NewRat(1, 1)),
			},
			want: MockWriter{
				AppendArg: AppendArg{
					Tones: []note.Semitone{36, 48, 52, 55},
					Value: note.NewValue(big.NewRat(1, 1)),
				},
				TextArg: TextArg{
					Text: "C",
				},
			},
		},
		{
			name: "key",
			node: &ast.Key{
				Key: note.NewKey(note.F, note.Flat, false),
			},
			want: MockWriter{
				KeyArg: KeyArg{
					Name:       note.F,
					Accidental: note.Flat,
					IsMinor:    false,
				},
			},
		},
		{
			name: "rest",
			node: &ast.Rest{
				Value: note.NewValue(big.NewRat(1, 2)),
			},
			want: MockWriter{
				RestArg: RestArg{
					Value: note.NewValue(big.NewRat(1, 2)),
				},
			},
		},
		{
			name: "meter",
			node: &ast.Meter{
				Num:   1,
				Denom: 4,
			},
			want: MockWriter{
				MeterArg: MeterArg{
					Num:   1,
					Denom: 4,
				},
			},
		},
		{
			name: "bpm",
			node: &ast.Tempo{
				BPM: 100,
			},
			want: MockWriter{
				BPMArg: BPMArg{
					BPM: 100,
				},
			},
		},
		{
			name: "instrument",
			node: &ast.Instrument{
				Name: "Clavi",
			},
			want: MockWriter{
				InstrumentArg: InstrumentArg{
					Name: "Clavi",
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var w MockWriter
			midi.NewASTWriter(&w).WriteNode(tc.node)
			assert.Equal(t, "", cmp.Diff(tc.want, w))
		})
	}
}
