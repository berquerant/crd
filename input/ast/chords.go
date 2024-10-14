package ast

//go:generate go run golang.org/x/tools/cmd/goyacc -o chords_goyacc_generated.go -v chords_goyacc_generated.output chords.y
//go:generate go run github.com/berquerant/mkvisitor -output chords_mkvisitor_generated.go -type ChordList,Chord,ChordDegree,ChordSymbol,ChordBase,ChordValues,ChordValue,Rest,ChordMeta,ChordMetadata
//go:generate go run github.com/berquerant/marker -output chords_marker_generated.go -method IsNode -type ChordList,Chord,ChordDegree,ChordSymbol,ChordBase,ChordValues,ChordValue,Rest,ChordMeta,ChordMetadata

func Parse(lexer *Lexer) int {
	return yyParse(lexer)
}

func SetDebug(level int) {
	yyDebug = level
}
