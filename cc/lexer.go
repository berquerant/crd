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
	// ParseUint8 parses a strng as an uint8.
	// Returns 0 and reports the error if failed.
	ParseUint8(x string) uint8

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
	expectInt bool // if true, expect that the next token will be an integer
}

func NewLexer(r io.Reader) Lexer {
	yyErrorVerbose = true // YYERROR_VERBOSE
	return &lexer{
		pos:    NewPos(1, 0, 0),
		reader: bufio.NewReader(r),
	}
}

func (s *lexer) Scan() int {
	s.discardIgnored()
	if s.expectInt {
		s.setExpectInt(false)
		if s.scanDigits() {
			return INT
		}
		s.setExpectInt(false) // int expected but not found
	}
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
		s.setExpectInt(true) // for reading note value
		s.next()
		return LBRA
	case ']':
		s.next()
		return RBRA
	case '/':
		s.next()
		switch s.Peek() {
		case '/':
			s.discardRemainingRow()
			s.ResetBuffer()
			return s.Scan()
		case '*':
			s.discardMultilineComment()
			s.ResetBuffer()
			return s.Scan()
		default:
			s.setExpectInt(true)
			return SLASH
		}
	case 'b', '♭':
		s.next()
		return FLAT
	case '#', '♯':
		s.next()
		return SHARP
	case '+':
		s.setExpectInt(true) // for reading accidental
		s.next()
		return PLUS
	case '-':
		s.setExpectInt(true) // for reading accidental
		s.next()
		return MINUS
	}
	if s.scanDigits() {
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

func (s *lexer) setExpectInt(v bool) {
	s.debugf("[expectInt] %v", v)
	s.expectInt = v
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
			switch s.Peek() {
			case 'a', 'e', 'i': // for maj, major, minor, meter
				continue
			default:
				return MINOR
			}
		case "minor":
			return MINOR
		case "meter":
			return METER
		case "M", "major":
			return MAJOR
		case "maj":
			switch s.Peek() {
			case 'o':
				continue // for major
			default:
				return MAJOR
			}
		case "aug":
			return AUGMENTED
		case "dim":
			return DIMINISHED
		case "sus":
			return SUSPENDED
		case "tempo":
			return TEMPO
		case "key":
			return KEY
		}
	}
	s.errorf("unknown ident %s", s.Buffer())
	return EOF
}

func (s *lexer) discardMultilineComment() {
	var expectEOC bool
	for x := s.Peek(); !(expectEOC && x == '/'); x = s.Peek() {
		expectEOC = x == '*'
		_ = s.Discard()
	}
	_ = s.Discard()
}

func (s *lexer) discardRemainingRow() {
	for x := s.Peek(); x != '\n'; x = s.Peek() {
		_ = s.Discard()
	}
	_ = s.Discard()
}

func (s *lexer) discardIgnored() {
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

func (s *lexer) ParseUint8(x string) uint8 {
	u, err := strconv.ParseUint(x, 10, 8)
	if err != nil {
		s.errorf("cannot parse %s as uint8", x)
		return 0
	}
	return uint8(u)
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
	s.debugf("[Lex] %s", lval.token)
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
