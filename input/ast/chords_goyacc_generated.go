// Code generated by goyacc -o chords_goyacc_generated.go -v chords_goyacc_generated.output chords.y. DO NOT EDIT.

//line chords.y:2
package ast

import __yyfmt__ "fmt"

//line chords.y:2

import "github.com/berquerant/ybase"

//line chords.y:7
type yySymType struct {
	yys           int
	result        *ChordList
	chord_list    []ChordOrRest
	chord_or_rest ChordOrRest
	rest          *Rest
	chod          *Chord
	degree        *ChordDegree
	symbol        *ChordSymbol
	simple_symbol *ChordSymbol
	base          *ChordBase
	values        *ChordValues
	value         *ChordValue
	metadata      *ChordMetadata
	meta_internal *ChordMeta
	meta          *ChordMeta

	degree_head ybase.Token
	accidental  ybase.Token

	token ybase.Token
}

const SYLLABLE = 57346
const SLASH = 57347
const LBRA = 57348
const RBRA = 57349
const COMMA = 57350
const SEMICOLON = 57351
const SHARP = 57352
const FLAT = 57353
const NUMBER = 57354
const SYMBOL = 57355
const REST = 57356
const UNDERSCORE = 57357
const LCBRA = 57358
const RCBRA = 57359
const EQUAL = 57360
const METADATA = 57361

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SYLLABLE",
	"SLASH",
	"LBRA",
	"RBRA",
	"COMMA",
	"SEMICOLON",
	"SHARP",
	"FLAT",
	"NUMBER",
	"SYMBOL",
	"REST",
	"UNDERSCORE",
	"LCBRA",
	"RCBRA",
	"EQUAL",
	"METADATA",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 49

var yyAct = [...]int8{
	37, 31, 20, 21, 7, 45, 38, 42, 41, 32,
	16, 9, 15, 22, 14, 9, 16, 40, 34, 10,
	28, 6, 29, 10, 18, 19, 39, 27, 12, 30,
	25, 33, 35, 26, 27, 24, 3, 36, 17, 11,
	8, 43, 44, 23, 13, 5, 4, 2, 1,
}

var yyPact = [...]int16{
	7, -1000, 7, -1000, -1000, -1000, 22, -3, 14, -1000,
	-1000, -1000, 1, 30, -1000, 3, -1000, -1000, -1000, -1000,
	26, -1000, 15, 16, 11, -1000, -7, 1, 6, 1,
	-1000, -1000, -13, -1000, -1000, 19, 0, -1000, -11, -7,
	-1000, -13, -14, -1000, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 48, 47, 36, 46, 45, 4, 44, 14, 43,
	2, 3, 40, 38, 0, 37, 1,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 3, 3, 4, 5, 6, 6,
	12, 12, 13, 13, 7, 7, 7, 8, 9, 9,
	10, 10, 11, 11, 16, 16, 15, 15, 14,
}

var yyR2 = [...]int8{
	0, 1, 1, 2, 1, 1, 5, 7, 1, 2,
	1, 1, 1, 1, 0, 1, 2, 1, 0, 2,
	1, 3, 1, 3, 0, 3, 1, 3, 3,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, -4, -5, 14, -6, -12, 4,
	12, -3, 6, -7, -8, 15, 13, -13, 10, 11,
	-10, -11, 12, -9, 5, -8, 7, 8, 5, 6,
	-6, -16, 16, -11, 12, -10, -15, -14, 19, 7,
	17, 8, 18, -16, -14, 19,
}

var yyDef = [...]int8{
	0, -2, 1, 2, 4, 5, 0, 14, 8, 10,
	11, 3, 0, 18, 15, 0, 17, 9, 12, 13,
	0, 20, 22, 0, 0, 16, 24, 0, 0, 0,
	19, 6, 0, 21, 23, 0, 0, 26, 0, 24,
	25, 0, 0, 7, 27, 28,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19,
}

var yyTok3 = [...]int8{
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
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
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
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
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
	yyn = int(yyPact[yystate])
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
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
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
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
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
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:66
		{
			r := &ChordList{
				List: yyDollar[1].chord_list,
			}
			yylex.(*Lexer).Result = r
			yyVAL.result = r
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:75
		{
			yyVAL.chord_list = []ChordOrRest{yyDollar[1].chord_or_rest}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line chords.y:78
		{
			yyVAL.chord_list = append(yyDollar[1].chord_list, yyDollar[2].chord_or_rest)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:83
		{
			yyVAL.chord_or_rest = yyDollar[1].rest
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:86
		{
			yyVAL.chord_or_rest = yyDollar[1].chod
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
//line chords.y:95
		{
			yyVAL.rest = &Rest{
				Values: yyDollar[3].values,
				Meta:   yyDollar[5].meta,
			}
		}
	case 7:
		yyDollar = yyS[yypt-7 : yypt+1]
//line chords.y:109
		{
			yyVAL.chod = &Chord{
				Degree: yyDollar[1].degree,
				Symbol: yyDollar[2].symbol,
				Base:   yyDollar[3].base,
				Values: yyDollar[5].values,
				Meta:   yyDollar[7].meta,
			}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:120
		{
			yyVAL.degree = &ChordDegree{
				Degree: yyDollar[1].degree_head,
			}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line chords.y:125
		{
			yyVAL.degree = &ChordDegree{
				Degree:     yyDollar[1].degree_head,
				Accidental: yyDollar[2].accidental,
			}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:133
		{
			yyVAL.degree_head = NewToken(yyDollar[1].token)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:136
		{
			yyVAL.degree_head = NewToken(yyDollar[1].token)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:141
		{
			yyVAL.accidental = NewToken(yyDollar[1].token)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:144
		{
			yyVAL.accidental = NewToken(yyDollar[1].token)
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
//line chords.y:149
		{
			yyVAL.symbol = nil
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:152
		{
			yyVAL.symbol = yyDollar[1].simple_symbol
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line chords.y:155
		{
			yyVAL.symbol = yyDollar[2].simple_symbol
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:160
		{
			yyVAL.simple_symbol = &ChordSymbol{
				Symbol: NewToken(yyDollar[1].token),
			}
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
//line chords.y:167
		{
			yyVAL.base = nil
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line chords.y:170
		{
			yyVAL.base = &ChordBase{
				Degree: yyDollar[2].degree,
			}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:177
		{
			yyVAL.values = &ChordValues{
				Values: []*ChordValue{yyDollar[1].value},
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line chords.y:182
		{
			yyVAL.values = &ChordValues{
				Values: append(yyDollar[1].values.Values, yyDollar[3].value),
			}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:189
		{
			yyVAL.value = &ChordValue{
				Num: NewToken(yyDollar[1].token),
			}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line chords.y:194
		{
			yyVAL.value = &ChordValue{
				Num:   NewToken(yyDollar[1].token),
				Denom: NewToken(yyDollar[3].token),
			}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
//line chords.y:202
		{
			yyVAL.meta = nil
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line chords.y:205
		{
			yyVAL.meta = yyDollar[2].meta_internal
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line chords.y:210
		{
			yyVAL.meta_internal = &ChordMeta{
				Data: []*ChordMetadata{yyDollar[1].metadata},
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line chords.y:215
		{
			yyVAL.meta_internal = &ChordMeta{
				Data: append(yyDollar[1].meta_internal.Data, yyDollar[3].metadata),
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line chords.y:224
		{
			yyVAL.metadata = &ChordMetadata{
				Key:   NewToken(yyDollar[1].token),
				Value: NewToken(yyDollar[3].token),
			}
		}
	}
	goto yystack /* stack new state and value */
}
