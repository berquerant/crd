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
		// Instrument appends an operation that sets instrument.
		Instrument(name string)
		// Append appends an operation that adds notes.
		Append(tones []note.Semitone, value note.Value)
		// Rest appends an operation that adds a rest.
		Rest(value note.Value)
		// Key appends an operation that adds key change.
		Key(name note.Name, accidental note.Accidental, isMinor bool)
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

func (s *writer) Instrument(name string) {
	s.add(func(w *mw.SMF) error {
		logger.Get().Debug("[Instrument] %s", name)
		return mw.Instrument(w, name)
	})
}

func (s *writer) Key(name note.Name, accidental note.Accidental, isMinor bool) {
	k := getKey(name, accidental, isMinor)()
	s.add(func(w *mw.SMF) error {
		logger.Get().Debug("[Key] %s%s minor = %s", name, accidental, isMinor)
		return mw.Key(w, k)
	})
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
		w     Writer
		trans note.Semitone
	}
)

func NewASTWriter(w Writer) ASTWriter {
	return &astWriter{
		w: w,
	}
}

func (s *astWriter) WriteNode(node ast.Node) {
	logger.Get().Debug("[WriteNode] %s", node)
	switch node := node.(type) {
	case *ast.Transposition:
		logger.Get().Debug("[WriteNode] trans %d", node.Semitone)
		s.trans = node.Semitone
	case *ast.Instrument:
		s.w.Instrument(node.Name)
	case *ast.Key:
		s.writeKey(node)
	case *ast.Meter:
		s.w.Meter(node.Num, node.Denom)
	case *ast.Tempo:
		s.w.BPM(node.BPM)
	case *ast.Rest:
		s.w.Rest(node.Value)
	case *ast.Chord:
		s.writeChord(node)
	default:
		logger.Get().Warn("[WriteNode] unknown node %s", node)
	}
}

func (s *astWriter) writeChord(node *ast.Chord) {
	if s.trans == 0 {
		if node.ChordBase != nil {
			s.w.Text(fmt.Sprintf("%s%son%s", node.ChordNote, node.ChordOption, node.ChordBase))
		} else {
			s.w.Text(fmt.Sprintf("%s%s", node.ChordNote, node.ChordOption))
		}
		s.w.Append(node.Semitones(), node.Value)
		return
	}
	// transposition
	n := &ast.ChordNote{
		SPN: (node.ChordNote.SPN.Semitone() + s.trans).SPN(),
	}
	if node.ChordBase != nil {
		s.w.Text(fmt.Sprintf("%s%son%s", n, node.ChordOption, &ast.ChordBase{
			Note: (node.ChordBase.Note.Semitone() + s.trans).Note(),
		}))
	} else {
		s.w.Text(fmt.Sprintf("%s%s", n, node.ChordOption))
	}
	v := node.Semitones()
	t := make([]note.Semitone, len(v))
	for i, x := range v {
		t[i] = x + s.trans
	}
	s.w.Append(t, node.Value)
}

func (s *astWriter) writeKey(node *ast.Key) {
	if s.trans == 0 {
		s.w.Key(node.Key.Name(), node.Key.Accidental(), node.Key.IsMinor())
		return
	}
	// transposition
	n := note.NewNote(node.Key.Name(), node.Key.Accidental())
	x := (n.Semitone() + s.trans).Note()
	s.w.Key(x.Name(), x.Accidental(), node.Key.IsMinor())
}

func (s *astWriter) Writer() Writer { return s.w }
