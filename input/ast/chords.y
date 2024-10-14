%{
package ast

import "github.com/berquerant/ybase"
%}

%union{
  result *ChordList
  chord_list []ChordOrRest
  chord_or_rest ChordOrRest
  rest *Rest
  chod *Chord
  degree *ChordDegree
  symbol *ChordSymbol
  simple_symbol *ChordSymbol
  base *ChordBase
  values *ChordValues
  value *ChordValue
  metadata *ChordMetadata
  meta_internal *ChordMeta
  meta *ChordMeta

  degree_head ybase.Token
  accidental ybase.Token

  token ybase.Token
}

%type <result> result
%type <chord_list> chord_list
%type <chord_or_rest> chord_or_rest
%type <rest> rest
%type <chod> chod
%type <degree> degree
%type <symbol> symbol
%type <simple_symbol> simple_symbol
%type <base> base
%type <values> values
%type <value> value
%type <degree_head> degree_head
%type <accidental> accidental
%type <metadata> metadata
%type <meta_internal> meta_internal
%type <meta> meta

%token <token> SYLLABLE
%token <token> SLASH
%token <token> LBRA
%token <token> RBRA
%token <token> COMMA
%token <token> SEMICOLON
%token <token> SHARP
%token <token> FLAT
%token <token> NUMBER
%token <token> SYMBOL
%token <token> REST
%token <token> UNDERSCORE
%token <token> LCBRA
%token <token> RCBRA
%token <token> EQUAL
%token <token> METADATA

%%

result:
  chord_list {
    r := &ChordList{
      List: $1,
    }
    yylex.(*Lexer).Result = r
    $$ = r
  }

chord_list:
  chord_or_rest {
    $$ = []ChordOrRest{$1}
  }
  | chord_list chord_or_rest {
    $$ = append($1, $2)
  }

chord_or_rest:
  rest {
    $$ = $1
  }
  | chod {
    $$ = $1
  }

rest:
  REST
  LBRA
  values
  RBRA
  meta {
    $$ = &Rest{
      Values: $3,
      Meta: $5,
    }
  }

chod:
  degree
  symbol
  base
  LBRA
  values
  RBRA
  meta {
    $$ = &Chord{
      Degree: $1,
      Symbol: $2,
      Base: $3,
      Values: $5,
      Meta: $7,
    }
  }

degree:
  degree_head {
    $$ = &ChordDegree{
      Degree: $1,
    }
  }
  | degree_head accidental {
    $$ = &ChordDegree{
      Degree: $1,
      Accidental: $2,
    }
  }

degree_head:
  SYLLABLE {
    $$ = NewToken($1)
  }
  | NUMBER {
    $$ = NewToken($1)
  }

accidental:
  SHARP {
   $$ = NewToken($1)
  }
  | FLAT {
   $$ = NewToken($1)
  }

symbol:
  {
    $$ = nil
  }
  | simple_symbol {
    $$ = $1
  }
  | UNDERSCORE simple_symbol {
    $$ = $2
  }

simple_symbol:
  SYMBOL {
    $$ = &ChordSymbol{
      Symbol: NewToken($1),
    }
  }

base:
  {
    $$ = nil
  }
  | SLASH degree {
    $$ = &ChordBase{
      Degree: $2,
    }
  }

values:
  value {
    $$ = &ChordValues{
      Values: []*ChordValue{$1},
    }
  }
  | values COMMA value {
    $$ = &ChordValues{
      Values: append($1.Values, $3),
    }
  }

value:
  NUMBER {
    $$ = &ChordValue{
      Num: NewToken($1),
    }
  }
  | NUMBER SLASH NUMBER {
    $$ = &ChordValue{
      Num: NewToken($1),
      Denom: NewToken($3),
    }
  }

meta:
  {
    $$ = nil
  }
  | LCBRA meta_internal RCBRA {
    $$ = $2
  }

meta_internal:
  metadata {
    $$ = &ChordMeta{
      Data: []*ChordMetadata{$1},
    }
  }
  | meta_internal COMMA metadata {
    $$ = &ChordMeta{
      Data: append($1.Data, $3),
    }
  }

metadata:
  METADATA
  EQUAL
  METADATA {
    $$ = &ChordMetadata{
      Key: NewToken($1),
      Value: NewToken($3),
    }
  }
