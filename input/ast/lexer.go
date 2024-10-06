package ast

import (
	"io"
	"log/slog"
	"strings"
	"unicode"

	"github.com/berquerant/ybase"
)

func NewLexer(r io.Reader) *Lexer {
	yyErrorVerbose = true
	scanner := &LexScanner{}
	lex := &Lexer{
		LexScanner: scanner,
		Lexer: ybase.NewLexer(
			ybase.NewScanner(
				ybase.NewReader(
					r,
					slog.Debug,
				),
				scanner.ScanFunc,
			),
		),
	}
	scanner.publishError = lex.Error
	return lex
}

var (
	_ yyLexer = &Lexer{}
)

type Lexer struct {
	ybase.Lexer
	*LexScanner
	Result *ChordList
}

func (lex *Lexer) Lex(lval *yySymType) int {
	return lex.DoLex(func(tok ybase.Token) {
		lval.token = tok
	})
}

type LexScanner struct {
	expectSymbol bool
	publishError func(string)
}

func (lex *LexScanner) SetExpectSymbol(v bool) {
	lex.expectSymbol = v
}

func (lex *LexScanner) ScanFunc(r ybase.Reader) int {
	r.DiscardWhile(unicode.IsSpace)

	if lex.expectSymbol {
		if lex.scanSymbol(r) {
			lex.SetExpectSymbol(false)
			return SYMBOL
		}
		lex.publishError("LexScanner: expect symbol failure")
		return ybase.EOF
	}

	nextRet := func(t int) int {
		_ = r.Next()
		return t
	}
	switch r.Peek() {
	case ';': // comment
		r.DiscardWhile(func(r rune) bool { return r != '\n' })
		return lex.ScanFunc(r)
	case 'C', 'D', 'E', 'F', 'G', 'A', 'B':
		return nextRet(SYLLABLE)
	case 'R':
		return nextRet(REST)
	case '/':
		return nextRet(SLASH)
	case '[':
		return nextRet(LBRA)
	case ']':
		return nextRet(RBRA)
	case ',':
		return nextRet(COMMA)
	case '#', '♯':
		return nextRet(SHARP)
	case 'b', '♭':
		return nextRet(FLAT)
	case '_':
		// when the symbol consists only of numbers, place an underscore in front of it
		// like: 5_7 (V7), G_7 (G7)
		lex.SetExpectSymbol(true)
		return nextRet(UNDERSCORE)
	}

	switch {
	case lex.scanDigits(r):
		return NUMBER
	case lex.scanSymbol(r):
		return SYMBOL
	default:
		return ybase.EOF
	}
}

func (LexScanner) isDigit(r rune) bool { return '0' <= r && r <= '9' }

func (lex LexScanner) scanDigits(r ybase.Reader) bool {
	slog.Debug("LexScanner: scanDigits try")
	if !lex.isDigit(r.Peek()) {
		return false
	}
	slog.Debug("LexScanner: scanDigits")
	r.NextWhile(lex.isDigit)
	return true
}

func (lex LexScanner) isSymbolRune(r rune) bool {
	// not beginning of slash chord, values, comment
	// no spaces
	return !lex.isBeginningOfNextOfSymbol(r) && !unicode.IsSpace(r)
}

func (LexScanner) isBeginningOfNextOfSymbol(r rune) bool {
	return strings.ContainsRune("/[_;", r)
}

func (lex LexScanner) scanSymbol(r ybase.Reader) bool {
	slog.Debug("LexScanner: scanSymbol try")

	x := r.Peek()
	switch {
	case x == ybase.EOF:
		return false
	case !lex.isSymbolRune(x):
		return false
	default:
		slog.Debug("LexScanner: scanSymbol")
		r.NextWhile(lex.isSymbolRune)
		return true
	}
}
