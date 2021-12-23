package midi

import (
	"fmt"

	"github.com/berquerant/crd/ast"
	"github.com/berquerant/crd/logger"
	"github.com/berquerant/crd/note"
	mw "gitlab.com/gomidi/midi/writer"
)

type (
	// Operation mutates SMF.
	Operation func(*mw.SMF) error

	// Writer is a midi operation accumulator.
	Writer interface {
		// BPM appends an operation that sets BPM.
		BPM(bpm int)
		// Meter appends an operation that sets the time signature.
		Meter(numerator, denominator uint8)
		// Text appends an operation that adds a text.
		Text(text string)
		// Velocity appends an operation that sets velocity.
		Velocity(velocity uint8)
		// Append appends an operation that adds notes.
		Append(tones []note.Semitone, value note.Value)
		// Rest appends an operation that adds a rest.
		Rest(value note.Value)
		// Operations returns the accumulated operations.
		Operations() []Operation
	}
)

func NewWriter() Writer {
	return &writer{
		ops:      []Operation{},
		velocity: 50,
	}
}

type writer struct {
	ops      []Operation
	velocity uint8
}

func (s *writer) Append(tones []note.Semitone, value note.Value) {
	numbers := make([]uint8, len(tones))
	for i, t := range tones {
		numbers[i] = s.noteNumber(t)
	}
	v := s.velocity
	f := s.forward(value)
	s.add(func(w *mw.SMF) error {
		for _, n := range numbers {
			logger.Get().Debug("[NoteOn] %d", n)
			if err := mw.NoteOn(w, n, v); err != nil {
				return err
			}
		}
		_ = f(w)
		for _, n := range numbers {
			logger.Get().Debug("[NoteOff] %d", n)
			if err := mw.NoteOff(w, n); err != nil {
				return err
			}
		}
		return nil
	})
}

func (*writer) noteNumber(tone note.Semitone) uint8 {
	center := note.NewSPN(note.NewNote(note.C, note.Natural), 4)
	diff := center.Semitone() - 60
	return uint8(tone - diff)
}

func (s *writer) Rest(value note.Value) {
	f := s.forward(value)
	s.add(func(w *mw.SMF) error {
		_ = f(w)
		return nil
	})
}

func (*writer) forward(value note.Value) Operation {
	n := value.Raw().Num().Uint64()
	d := value.Raw().Denom().Uint64()
	return func(w *mw.SMF) error {
		logger.Get().Debug("[forward] %d/%d", n, d)
		mw.Forward(w, 0,
			uint32(n),
			uint32(d),
		)
		return nil
	}
}

func (s *writer) Text(text string) {
	s.add(func(w *mw.SMF) error {
		logger.Get().Debug("[Text] %s", text)
		return mw.Text(w, text)
	})
}

func (s *writer) Meter(numerator, denominator uint8) {
	s.add(func(w *mw.SMF) error {
		logger.Get().Debug("[Meter] %d/%d", numerator, denominator)
		return mw.Meter(w, numerator, denominator)
	})
}

func (s *writer) BPM(bpm int) {
	s.add(func(w *mw.SMF) error {
		logger.Get().Debug("[BPM] %d", bpm)
		return mw.TempoBPM(w, float64(bpm))
	})
}

func (s *writer) add(op Operation) { s.ops = append(s.ops, op) }
func (s *writer) Velocity(velocity uint8) {
	logger.Get().Debug("[Velocity] %d", velocity)
	s.velocity = velocity
}
func (s *writer) Operations() []Operation { return s.ops }

type (
	// ASTWriter converts AST into midi operations.
	ASTWriter interface {
		WriteNode(node ast.Node)
		Writer() Writer
	}

	astWriter struct {
		w Writer
	}
)

func NewASTWriter(w Writer) ASTWriter {
	return &astWriter{
		w: w,
	}
}

func (s *astWriter) WriteNode(node ast.Node) {
	switch node := node.(type) {
	case *ast.Rest:
		logger.Get().Debug("[WriteNode] %s", node)
		s.w.Rest(node.Value)
	case *ast.Chord:
		t := fmt.Sprintf("%s%s", node.ChordNote, node.ChordOption)
		logger.Get().Debug("[WriteNode] %s", t)
		s.w.Text(t)
		s.w.Append(node.Semitones(), node.Value)
	}
}

func (s *astWriter) Writer() Writer { return s.w }
