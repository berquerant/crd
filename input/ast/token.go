package ast

import (
	"fmt"

	"github.com/berquerant/ybase"
)

var (
	_ ybase.Pos   = &Pos{}
	_ ybase.Token = &Token{}
)

type Pos struct {
	VLine   int `json:"line" yaml:"line"`
	VColumn int `json:"col" yaml:"col"`
	VOffset int `json:"offset" yaml:"offset"`
}

func (p Pos) Line() int            { return p.VLine }
func (p Pos) Column() int          { return p.VColumn }
func (p Pos) Offset() int          { return p.VOffset }
func (p Pos) Add(_ rune) ybase.Pos { return p }

func NewPos(p ybase.Pos) ybase.Pos {
	return &Pos{
		VLine:   p.Line(),
		VColumn: p.Column(),
		VOffset: p.Offset(),
	}
}

type Token struct {
	VType  int       `json:"type" yaml:"type"`
	VValue string    `json:"value" yaml:"value"`
	VStart ybase.Pos `json:"start" yaml:"start"`
	VEnd   ybase.Pos `json:"end" yaml:"end"`
}

func (t Token) Type() int        { return t.VType }
func (t Token) Value() string    { return t.VValue }
func (t Token) Start() ybase.Pos { return t.VStart }
func (t Token) End() ybase.Pos   { return t.VEnd }
func (t Token) String() string {
	return fmt.Sprintf("type=%d value=%s start=%s end=%s",
		t.VType, t.VValue, t.VStart, t.VEnd)
}

func NewToken(t ybase.Token) ybase.Token {
	return &Token{
		VType:  t.Type(),
		VValue: t.Value(),
		VStart: NewPos(t.Start()),
		VEnd:   NewPos(t.End()),
	}
}
