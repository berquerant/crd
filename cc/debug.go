package cc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/berquerant/crd/ast"
	"github.com/berquerant/crd/logger"
	"github.com/berquerant/crd/note"
)

func NewDebugger(lexer Lexer, opt ...DebuggerOption) *Debugger {
	x := &Debugger{
		lexer: lexer,
	}
	for _, o := range opt {
		o(x)
	}
	return x
}

func WithTransposition(v note.Semitone) DebuggerOption {
	return func(d *Debugger) {
		d.transposition = v
	}
}

type (
	DebuggerOption func(*Debugger)
	Debugger       struct {
		lexer         Lexer
		transposition note.Semitone
	}
)

// Lex prints the process of the lexical analysis.
func (s *Debugger) Lex() {
	for {
		t := s.lexer.Scan()
		if err := s.lexer.Err(); err != nil {
			logger.Get().Error("%v", err)
			return
		}
		if t == EOF {
			logger.Get().Info("lexer finished successfully")
			return
		}
		tok := NewToken(t, s.lexer.Buffer())
		s.lexer.ResetBuffer()
		fmt.Println(tok)
	}
}

// Parse converts the input score into AST.
func (s *Debugger) Parse() {
	status := Parse(s.lexer)
	logger.Get().Info("parser exit with %d", status)
	if err := s.lexer.Err(); err != nil {
		logger.Get().Error("%v", err)
		return
	}
	for _, n := range s.lexer.Result().NodeList {
		v, _ := json.Marshal(n)
		fmt.Printf("%s\t%s\n", n, v)
	}
}

// Unparse normalizes the input score.
func (s *Debugger) Unparse() {
	status := Parse(s.lexer)
	logger.Get().Info("parser exit with %d", status)
	if err := s.lexer.Err(); err != nil {
		logger.Get().Error("%v", err)
		return
	}
	fmt.Printf("%s\n", s.lexer.Result())
}

// Semitones prints the semitones of the score.
func (s *Debugger) Semitones() {
	status := Parse(s.lexer)
	logger.Get().Info("parser exit with %d", status)
	if err := s.lexer.Err(); err != nil {
		logger.Get().Error("%v", err)
		return
	}
	for _, n := range s.lexer.Result().NodeList {
		switch n := n.(type) {
		case *ast.Key:
			v := (n.Key.Name().Semitone().Sign(n.Key.Accidental()) + s.transposition).Note()
			fmt.Println(&ast.Key{
				Key: note.NewKey(v.Name(), v.Accidental(), n.Key.IsMinor()),
			})
		case *ast.Chord:
			var (
				c = &ast.Chord{
					ChordNote: &ast.ChordNote{
						SPN: (n.ChordNote.SPN.Semitone() + s.transposition).SPN(),
					},
					ChordOption: n.ChordOption,
					ChordBase: func() *ast.ChordBase {
						if n.ChordBase == nil {
							return nil
						}
						return &ast.ChordBase{
							Note: (n.ChordBase.Note.Semitone() + s.transposition).Note(),
						}
					}(),
					Value: n.Value,
				}
			)
			tones := c.Semitones()
			v := make([]string, len(tones))
			for i, x := range tones {
				v[i] = strconv.Itoa(int(x))
			}
			fmt.Printf("%s\t%s\n", c, strings.Join(v, " "))
		default:
			fmt.Println(n)
		}
	}
}
