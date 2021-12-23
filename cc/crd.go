// Code generated by goyacc -o cc/crd.go -v cc/crd.output cc/crd.y. DO NOT EDIT.

//line cc/crd.y:2
package cc

import __yyfmt__ "fmt"

//line cc/crd.y:2

import (
	"github.com/berquerant/crd/ast"
	"github.com/berquerant/crd/note"
	"math/big"
)

//line cc/crd.y:11
type yySymType struct {
	yys          int
	score        *ast.Score
	nodeList     []ast.Node
	node         ast.Node
	token        Token
	chordNote    *ast.ChordNote
	chordOption  *ast.ChordOption
	value        note.Value
	name         note.Name
	augmented    bool
	diminished   bool
	minor        bool
	major        bool
	seventh      bool
	sixth        bool
	suspended    bool
	forth        bool
	accidental   note.Accidental
	accidentaled int
}

const REST = 57346
const MINOR = 57347
const MAJOR = 57348
const SEVENTH = 57349
const DIMINISHED = 57350
const AUGMENTED = 57351
const SIXTH = 57352
const SUSPENDED = 57353
const FORTH = 57354
const MINUS = 57355
const PLUS = 57356
const SHARP = 57357
const FLAT = 57358
const SLASH = 57359
const INT = 57360
const LBRA = 57361
const RBRA = 57362
const C = 57363
const D = 57364
const E = 57365
const F = 57366
const G = 57367
const A = 57368
const B = 57369

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"REST",
	"MINOR",
	"MAJOR",
	"SEVENTH",
	"DIMINISHED",
	"AUGMENTED",
	"SIXTH",
	"SUSPENDED",
	"FORTH",
	"MINUS",
	"PLUS",
	"SHARP",
	"FLAT",
	"SLASH",
	"INT",
	"LBRA",
	"RBRA",
	"C",
	"D",
	"E",
	"F",
	"G",
	"A",
	"B",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line cc/crd.y:214

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 53

var yyAct = [...]int{
	6, 30, 36, 53, 29, 17, 49, 47, 46, 48,
	18, 44, 52, 51, 42, 50, 33, 9, 10, 11,
	12, 13, 14, 15, 25, 26, 23, 24, 40, 21,
	28, 38, 35, 32, 3, 45, 43, 16, 41, 39,
	37, 34, 31, 27, 20, 22, 8, 19, 7, 5,
	4, 2, 1,
}

var yyPact = [...]int{
	-4, -1000, -4, -1000, -1000, -1000, -9, 20, 11, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 6, -9,
	22, -1000, -1000, -1000, -1000, -16, -1000, 28, -1000, -1000,
	-2, 26, -1000, -18, 24, -1000, -1000, 18, -1000, 3,
	-1000, -1, -1000, -7, -1000, -1000, -3, -5, -6, -15,
	-1000, -1000, -1000, -1000,
}

var yyPgo = [...]int{
	0, 52, 51, 34, 50, 49, 5, 48, 47, 46,
	45, 44, 43, 42, 41, 40, 39, 38, 36, 35,
}

var yyR1 = [...]int{
	0, 1, 2, 2, 3, 3, 4, 5, 7, 9,
	9, 9, 9, 9, 9, 9, 10, 10, 10, 8,
	11, 11, 12, 12, 13, 13, 14, 14, 15, 15,
	16, 16, 17, 17, 18, 18, 19, 19, 19, 19,
	19, 6, 6,
}

var yyR2 = [...]int{
	0, 1, 1, 2, 1, 1, 2, 3, 2, 1,
	1, 1, 1, 1, 1, 1, 0, 1, 1, 9,
	0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
	0, 1, 0, 1, 0, 1, 0, 2, 2, 2,
	2, 3, 5,
}

var yyChk = [...]int{
	-1000, -1, -2, -3, -4, -5, 4, -7, -9, 21,
	22, 23, 24, 25, 26, 27, -3, -6, 19, -8,
	-11, 9, -10, 15, 16, 18, -6, -12, 8, 20,
	17, -13, 5, 18, -14, 6, 20, -15, 7, -16,
	10, -17, 11, -18, 12, -19, 15, 14, 16, 13,
	18, 18, 18, 18,
}

var yyDef = [...]int{
	0, -2, 1, 2, 4, 5, 0, 20, 16, 9,
	10, 11, 12, 13, 14, 15, 3, 6, 0, 0,
	22, 21, 8, 17, 18, 0, 7, 24, 23, 41,
	0, 26, 25, 0, 28, 27, 42, 30, 29, 32,
	31, 34, 33, 36, 35, 19, 0, 0, 0, 0,
	37, 38, 39, 40,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:78
		{
			x := &ast.Score{NodeList: yyDollar[1].nodeList}
			yylex.(Lexer).SetResult(x)
			yyVAL.score = x
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:85
		{
			yyVAL.nodeList = []ast.Node{yyDollar[1].node}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line cc/crd.y:88
		{
			yyVAL.nodeList = append(yyDollar[1].nodeList, yyDollar[2].node)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line cc/crd.y:96
		{
			yyVAL.node = &ast.Rest{Value: yyDollar[2].value}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line cc/crd.y:101
		{
			yyVAL.node = &ast.Chord{
				ChordNote:   yyDollar[1].chordNote,
				ChordOption: yyDollar[2].chordOption,
				Value:       yyDollar[3].value,
			}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line cc/crd.y:110
		{
			yyVAL.chordNote = &ast.ChordNote{
				Name:       yyDollar[1].name,
				Octave:     note.Octave(4),
				Accidental: yyDollar[2].accidental,
			}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:119
		{
			yyVAL.name = note.C
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:120
		{
			yyVAL.name = note.D
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:121
		{
			yyVAL.name = note.E
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:122
		{
			yyVAL.name = note.F
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:123
		{
			yyVAL.name = note.G
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:124
		{
			yyVAL.name = note.A
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:125
		{
			yyVAL.name = note.B
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:128
		{
			yyVAL.accidental = note.Natural
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:129
		{
			yyVAL.accidental = note.Sharp
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:130
		{
			yyVAL.accidental = note.Flat
		}
	case 19:
		yyDollar = yyS[yypt-9 : yypt+1]
//line cc/crd.y:141
		{
			yyVAL.chordOption = &ast.ChordOption{
				IsAugmented:  yyDollar[1].augmented,
				IsDiminished: yyDollar[2].diminished,
				IsMinor:      yyDollar[3].minor,
				IsMajor:      yyDollar[4].major,
				IsSeventh:    yyDollar[5].seventh,
				IsSixth:      yyDollar[6].sixth,
				IsSuspended:  yyDollar[7].suspended,
				IsForth:      yyDollar[8].forth,
				Accidentaled: yyDollar[9].accidentaled,
			}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:156
		{
			yyVAL.augmented = false
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:157
		{
			yyVAL.augmented = true
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:160
		{
			yyVAL.diminished = false
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:161
		{
			yyVAL.diminished = true
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:164
		{
			yyVAL.minor = false
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:165
		{
			yyVAL.minor = true
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:168
		{
			yyVAL.major = false
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:169
		{
			yyVAL.major = true
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:172
		{
			yyVAL.seventh = false
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:173
		{
			yyVAL.seventh = true
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:176
		{
			yyVAL.sixth = false
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:177
		{
			yyVAL.sixth = true
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:180
		{
			yyVAL.suspended = false
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:181
		{
			yyVAL.suspended = true
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:184
		{
			yyVAL.forth = false
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line cc/crd.y:185
		{
			yyVAL.forth = true
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
//line cc/crd.y:188
		{
			yyVAL.accidentaled = 0
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line cc/crd.y:189
		{
			yyVAL.accidentaled = yylex.(Lexer).ParseInt(yyDollar[2].token.Value())
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line cc/crd.y:192
		{
			yyVAL.accidentaled = yylex.(Lexer).ParseInt(yyDollar[2].token.Value())
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line cc/crd.y:195
		{
			yyVAL.accidentaled = -yylex.(Lexer).ParseInt(yyDollar[2].token.Value())
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line cc/crd.y:198
		{
			yyVAL.accidentaled = -yylex.(Lexer).ParseInt(yyDollar[2].token.Value())
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line cc/crd.y:203
		{
			i := yylex.(Lexer).ParseInt(yyDollar[2].token.Value())
			yyVAL.value = note.NewValue(new(big.Rat).SetInt64(int64(i)))
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
//line cc/crd.y:207
		{
			l := yylex.(Lexer)
			x := l.ParseInt(yyDollar[2].token.Value())
			y := l.ParseInt(yyDollar[4].token.Value())
			yyVAL.value = note.NewValue(big.NewRat(int64(x), int64(y)))
		}
	}
	goto yystack /* stack new state and value */
}
