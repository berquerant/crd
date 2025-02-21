package astconv

import (
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/input/ast"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/util"
)

//go:generate go tool stringer -type ASTType -output validate_stringer_generated.go
type ASTType int

const (
	UnknownASTType ASTType = iota
	SyllableAST
	DegreeAST
)

type ASTClassifier interface {
	// Classify AST type.
	Classify(v ast.Node) (ASTType, error)
}

type ASTTypeClassifier struct{}

func NewASTClassifier() *ASTTypeClassifier {
	return &ASTTypeClassifier{}
}

var (
	_ ASTClassifier = &ASTTypeClassifier{}
)

func (c ASTTypeClassifier) Classify(v ast.Node) (ASTType, error) {
	var (
		isInit  = true
		astType ASTType
	)

	for x := range ast.NewIterVisitor().All(v) {
		degree, ok := x.(*ast.ChordDegree)
		if !ok {
			continue
		}

		t := c.degreeType(degree)
		if t == UnknownASTType {
			return UnknownASTType, errorx.Invalid("AST type is unknown: at %s", degree.Degree)
		}
		if isInit {
			isInit = false
			astType = t
			continue
		}
		if t != astType {
			return UnknownASTType, errorx.Invalid("AST type is inconsistent: %s to %s, at %s",
				astType, t, degree.Degree)
		}
	}

	if astType == UnknownASTType {
		return UnknownASTType, errorx.Invalid("AST type is unknown")
	}

	return astType, nil
}

func (ASTTypeClassifier) degreeType(v *ast.ChordDegree) ASTType {
	s := v.Degree.Value()

	if x := note.NewName(s); x != note.UnknownName {
		return SyllableAST
	}

	if _, err := util.ParseUint(s); err == nil {
		return DegreeAST
	}

	return UnknownASTType
}
