// Code generated by "stringer -type ASTType -output validate_stringer_generated.go"; DO NOT EDIT.

package astconv

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownASTType-0]
	_ = x[SyllableAST-1]
	_ = x[DegreeAST-2]
}

const _ASTType_name = "UnknownASTTypeSyllableASTDegreeAST"

var _ASTType_index = [...]uint8{0, 14, 25, 34}

func (i ASTType) String() string {
	if i < 0 || i >= ASTType(len(_ASTType_index)-1) {
		return "ASTType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ASTType_name[_ASTType_index[i]:_ASTType_index[i+1]]
}
