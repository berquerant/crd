package cc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"unicode"

	"github.com/berquerant/crd/ast"
	"github.com/berquerant/crd/logger"
)

func Parse(lexer Lexer) int { return yyParse(lexer) }

const EOF = -1

type Lexer interface {
	SetResult(result *ast.Score)
	Result() *ast.Score

	Scan() int
	Buffer() string
	ResetBuffer()

	Debug(level int)
	Err() error

	// ParseInt parses a string as an integer.
	// Returns 0 and reports the error if failed.
	ParseInt(x string) int

	/* Implements yyLexer. */
	Lex(lval *yySymType) int
	Error(msg string) // yyerror
}

type lexer struct {
	pos       Pos
	reader    *bufio.Reader
	buf       bytes.Buffer
	result    *ast.Score
	err       error
	isReadInt bool
}

func NewLexer(r io.Reader) Lexer {
	yyErrorVerbose = true // YYERROR_VERBOSE
	return &lexer{
		pos:    NewPos(1, 0, 0),
		reader: bufio.NewReader(r),
	}
}

func (s *lexer) Scan() int {
	s.scanIgnored()
	switch s.Peek() {
	case EOF:
		return EOF
	case 'C':
		s.next()
		return C
	case 'D':
		s.next()
		return D
	case 'E':
		s.next()
		return E
	case 'F':
		s.next()
		return F
	case 'G':
		s.next()
		return G
	case 'A':
		s.next()
		return A
	case 'B':
		s.next()
		return B
	case 'R':
		s.next()
		return REST
	case '[':
		s.isReadInt = true // for reading note value
		s.next()
		return LBRA
	case ']':
		s.next()
		return RBRA
	case '/':
		s.isReadInt = true // for reading note value denominator
		s.next()
		return SLASH
	case 'b', '♭':
		s.next()
		return FLAT
	case '#', '♯':
		s.next()
		return SHARP
	case '+':
		s.isReadInt = true // for reading accidental
		s.next()
		return PLUS
	case '-':
		s.isReadInt = true // for reading accidental
		s.next()
		return MINUS
	}
	if s.scanDigits() {
		if s.isReadInt {
			s.isReadInt = false
			return INT
		}
		/* forth, sixth and seventh preceed the note value and accidentaled (added note or chord accidental) in this repository */
		switch s.Buffer() {
		case "4":
			return FORTH
		case "6":
			return SIXTH
		case "7":
			return SEVENTH
		default:
			s.errorf("failed to read chord option")
			return EOF
		}
	}
	return s.scanIdent()
}

func (*lexer) isAlpha(r rune) bool { return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' }
func (*lexer) isDigit(r rune) bool { return '0' <= r && r <= '9' }

func (s *lexer) scanIdent() int {
	if !s.isAlpha(s.Peek()) {
		return EOF
	}
	for x := s.Peek(); s.isAlpha(x); x = s.Peek() {
		s.next()
		switch s.Buffer() {
		case "m":
			return MINOR
		case "M", "maj":
			return MAJOR
		case "aug":
			return AUGMENTED
		case "dim":
			return DIMINISHED
		case "sus":
			return SUSPENDED
		}
	}
	s.errorf("unknown ident %s", s.Buffer())
	return EOF
}

func (s *lexer) scanIgnored() {
	for x := s.Peek(); unicode.IsSpace(x) || x == '|' || x == '｜'; x = s.Peek() {
		_ = s.Discard()
	}
}

func (s *lexer) scanDigits() bool {
	if !unicode.IsDigit(s.Peek()) {
		return false
	}
	for x := s.Peek(); s.isDigit(x); x = s.Peek() {
		s.next()
	}
	return true
}

func (s *lexer) ParseInt(x string) int {
	i, err := strconv.Atoi(x)
	if err != nil {
		s.errorf("cannot parse %s as int", x)
		return 0
	}
	return i
}

func (s *lexer) Discard() rune {
	r, _, err := s.reader.ReadRune()
	s.debugf("[Discard] %q %v", r, err)
	if err != nil {
		if err != io.EOF {
			s.errorf("[Discard] from reader %v", err)
		}
		return EOF
	}
	s.pos = s.pos.Add(r)
	return r
}

func (s *lexer) Peek() rune {
	r, _, err := s.reader.ReadRune()
	s.debugf("[Peek] %q %v", r, err)
	if err != nil {
		if err != io.EOF {
			s.errorf("[Peek] from reader %v", err)
		}
		return EOF
	}
	if err := s.reader.UnreadRune(); err != nil {
		s.errorf("[Peek] failed to unread %v", err)
		return EOF
	}
	return r
}

func (s *lexer) Next() rune {
	r, _, err := s.reader.ReadRune()
	s.debugf("[Next] %q %v", r, err)
	if err != nil {
		if err != io.EOF {
			s.errorf("[Next] from reader %v", err)
		}
		return EOF
	}
	s.pos = s.pos.Add(r)
	if _, err := s.buf.WriteRune(r); err != nil {
		s.errorf("[Next] failed to write buffer %v", err)
		return EOF
	}
	return r
}

func (s *lexer) Lex(lval *yySymType) int {
	if s.err != nil {
		return EOF
	}
	t := s.Scan()
	v := s.Buffer()
	lval.token = NewToken(t, v)
	s.ResetBuffer()
	return t
}

func (s *lexer) next()                                  { _ = s.Next() }
func (s *lexer) Buffer() string                         { return s.buf.String() }
func (s *lexer) ResetBuffer()                           { s.buf.Reset() }
func (s *lexer) Err() error                             { return s.err }
func (s *lexer) SetResult(result *ast.Score)            { s.result = result }
func (s *lexer) Result() *ast.Score                     { return s.result }
func (s *lexer) errorf(format string, v ...interface{}) { s.Error(fmt.Sprintf(format, v...)) }
func (s *lexer) Error(msg string) {
	s.err = fmt.Errorf("[lex][%s][%s] %s", s.pos, s.buf.String(), msg)
	s.debugf("[Error] %v", s.err)
}
func (s *lexer) debugf(format string, v ...interface{}) {
	logger.Get().Debug("[lex][%s][%s] %s", s.pos, s.buf.String(), fmt.Sprintf(format, v...))
}

func (s *lexer) Debug(level int) {
	if level < 0 {
		return
	}
	yyDebug = level // YYDEBUG
}
