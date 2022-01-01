%{
package cc

import (
  "math/big"
  "github.com/berquerant/crd/ast"
  "github.com/berquerant/crd/note"
)
%}

%union{
  score *ast.Score
  nodeList []ast.Node
  node ast.Node
  token Token
  chordNote *ast.ChordNote
  chordOption *ast.ChordOption
  chordBase *ast.ChordBase
  value note.Value
  name note.Name
  augmented bool
  diminished bool
  minor bool
  major bool
  seventh bool
  sixth bool
  suspended bool
  added bool
  ninth bool
  forth bool
  accidental note.Accidental
  accidentaled int
  keyMinor bool
  key note.Key
  transpositionDiff note.Semitone
}

%type <score> score
%type <nodeList> node_list
%type <node> node rest chord tempo meter key instrument transposition
%type <value> value
%type <chordNote>  chord_note
%type <chordOption> chord_option
%type <chordBase> chord_base
%type <name> name
%type <accidental> accidental
%type <augmented> augmented
%type <diminished> diminished
%type <minor> minor
%type <major> major
%type <seventh> seventh
%type <sixth> sixth
%type <suspended> suspended
%type <added> added
%type <ninth> ninth
%type <forth> forth
%type <accidentaled> accidentaled
%type <keyMinor> key_major_minor
%type <transpositionDiff> transposition_diff

%token <token> REST
%token <token> MINOR
%token <token> MAJOR
%token <token> SEVENTH
%token <token> DIMINISHED
%token <token> AUGMENTED
%token <token> SIXTH
%token <token> SUSPENDED
%token <token> ADDED
%token <token> NINTH
%token <token> FORTH
%token <token> MINUS
%token <token> PLUS
%token <token> SHARP
%token <token> FLAT
%token <token> SLASH
%token <token> INT /* integer */
%token <token> LBRA
%token <token> RBRA
%token <token> C
%token <token> D
%token <token> E
%token <token> F
%token <token> G
%token <token> A
%token <token> B
%token <token> TEMPO
%token <token> METER
%token <token> KEY
%token <token> INSTRUMENT
%token <token> STRING
%token <token> TRANSPOSITION

%%

score:
  node_list {
    x := &ast.Score{NodeList: $1}
    yylex.(Lexer).SetResult(x)
    $$ = x
  }

node_list:
  node {
    $$ = []ast.Node{$1}
  }
  | node_list node {
    $$ = append($1, $2)
  }

node:
  rest | chord | tempo | meter | key | instrument | transposition

transposition:
  TRANSPOSITION LBRA transposition_diff RBRA {
    $$ = &ast.Transposition{Semitone: $3}
  }

transposition_diff:
  INT {
    $$ = note.Semitone(yylex.(Lexer).ParseInt($1.Value()))
  }
  | PLUS INT {
    $$ = note.Semitone(yylex.(Lexer).ParseInt($2.Value()))
  }
  | MINUS INT {
    $$ = note.Semitone(-yylex.(Lexer).ParseInt($2.Value()))
  }

instrument:
  INSTRUMENT LBRA STRING RBRA {
    $$ = &ast.Instrument{
      Name: $3.Value(),
    }
  }

key:
  KEY LBRA name accidental key_major_minor RBRA {
    $$ = &ast.Key{
      Key: note.NewKey($3, $4, $5),
    }
  }

key_major_minor:
  { $$ = false }
  | MAJOR { $$ = false }
  | MINOR { $$ = true }

tempo:
  TEMPO LBRA INT RBRA {
    bpm := yylex.(Lexer).ParseInt($3.Value())
    $$ = &ast.Tempo{
      BPM: bpm,
    }
  }

meter:
  METER LBRA INT SLASH INT RBRA {
    l := yylex.(Lexer)
    n := l.ParseUint8($3.Value())
    d := l.ParseUint8($5.Value())
    $$ = &ast.Meter{
      Num: n,
      Denom: d,
    }
  }

rest:
  REST value {
    $$ = &ast.Rest{Value: $2}
  }

chord:
  chord_note chord_option chord_base value {
    $$ = &ast.Chord{
      ChordNote: $1,
      ChordOption: $2,
      ChordBase: $3,
      Value: $4,
    }
  }

chord_note:
  name accidental {
    n := note.NewSPN(note.NewNote($1, $2), note.Octave(4))
    $$ = &ast.ChordNote{SPN: n}
  }

name:
  C { $$ = note.C }
  | D { $$ = note.D }
  | E { $$ = note.E }
  | F { $$ = note.F }
  | G { $$ = note.G }
  | A { $$ = note.A }
  | B { $$ = note.B }

accidental:
  { $$ = note.Natural }
  | SHARP { $$ = note.Sharp }
  | FLAT { $$ = note.Flat }

chord_base:
  { $$ = nil}
  | SLASH name accidental {
    $$ = &ast.ChordBase{
      Note: note.NewNote($2, $3),
    }
  }

chord_option:
  augmented
  diminished
  minor
  major
  seventh
  sixth
  suspended
  forth
  added
  ninth
  accidentaled {
    $$ = &ast.ChordOption{
      IsAugmented: $1,
      IsDiminished: $2,
      IsMinor: $3,
      IsMajor: $4,
      IsSeventh: $5,
      IsSixth: $6,
      IsSuspended: $7,
      IsForth: $8,
      IsAdded: $9,
      IsNinth: $10,
      Accidentaled: $11,
    }
  }

augmented:
  { $$ = false }
  | AUGMENTED { $$ = true }

diminished:
  { $$ = false }
  | DIMINISHED { $$ = true }

minor:
  { $$ = false }
  | MINOR { $$ = true }

major:
  { $$ = false }
  | MAJOR { $$ = true }

seventh:
  { $$ = false }
  | SEVENTH { $$ = true }

sixth:
  { $$ = false }
  | SIXTH { $$ = true }

suspended:
  { $$ = false }
  | SUSPENDED { $$ = true }

forth:
  { $$ = false }
  | FORTH { $$ = true }

added:
  { $$ = false }
  | ADDED { $$ = true }

ninth:
  { $$ = false }
  | NINTH { $$ = true }

accidentaled:
  { $$ = 0 }
  | SHARP INT {
    $$ = yylex.(Lexer).ParseInt($2.Value())
  }
  | PLUS INT {
    $$ = yylex.(Lexer).ParseInt($2.Value())
  }
  | FLAT INT {
    $$ = -yylex.(Lexer).ParseInt($2.Value())
  }
  | MINUS INT {
    $$ = -yylex.(Lexer).ParseInt($2.Value())
  }

value:
  LBRA INT RBRA {
    i := yylex.(Lexer).ParseInt($2.Value())
    $$ = note.NewValue(new(big.Rat).SetInt64(int64(i)))
  }
  | LBRA INT SLASH INT RBRA {
    l := yylex.(Lexer)
    x := l.ParseInt($2.Value())
    y := l.ParseInt($4.Value())
    $$ = note.NewValue(big.NewRat(int64(x), int64(y)))
  }

%%
