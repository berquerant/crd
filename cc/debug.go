package cc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/berquerant/crd/logger"
)

func NewDebugger(lexer Lexer) *Debugger {
	return &Debugger{
		lexer: lexer,
	}
}

type Debugger struct {
	lexer Lexer
}

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
		tones := n.Semitones()
		v := make([]string, len(tones))
		for i, x := range tones {
			v[i] = strconv.Itoa(int(x))
		}
		fmt.Printf("%s\t%s\n", n, strings.Join(v, " "))
	}
}
