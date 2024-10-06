package astconv_test

import (
	"testing"

	"github.com/berquerant/crd/astconv"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/input/ast"
	"github.com/stretchr/testify/assert"
)

func TestASTClassifier(t *testing.T) {
	var (
		restTree = &ast.ChordList{
			List: []ast.ChordOrRest{
				&ast.Rest{},
			},
		}
		degreeTree = &ast.ChordList{
			List: []ast.ChordOrRest{
				&ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "1",
						},
					},
				},
			},
		}
		syllableTree = &ast.ChordList{
			List: []ast.ChordOrRest{
				&ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "C",
						},
					},
				},
			},
		}
		inconsistentTree = &ast.ChordList{
			List: []ast.ChordOrRest{
				&ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "1",
						},
					},
				},
				&ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "C",
						},
					},
				},
			},
		}
		inconsistentBaseTree = &ast.ChordList{
			List: []ast.ChordOrRest{
				&ast.Chord{
					Degree: &ast.ChordDegree{
						Degree: &ast.Token{
							VValue: "1",
						},
					},
					Base: &ast.ChordBase{
						Degree: &ast.ChordDegree{
							Degree: &ast.Token{
								VValue: "C",
							},
						},
					},
				},
			},
		}
	)

	for _, tc := range []struct {
		title string
		tree  ast.Node
		want  astconv.ASTType
		err   error
	}{
		{
			title: "no chords",
			tree:  restTree,
			err:   errorx.ErrInvalid,
		},
		{
			title: "degree",
			tree:  degreeTree,
			want:  astconv.DegreeAST,
		},
		{
			title: "syllable",
			tree:  syllableTree,
			want:  astconv.SyllableAST,
		},
		{
			title: "inconsistent",
			tree:  inconsistentTree,
			err:   errorx.ErrInvalid,
		},
		{
			title: "inconsistent base",
			tree:  inconsistentBaseTree,
			err:   errorx.ErrInvalid,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			c := astconv.NewASTClassifier()
			got, err := c.Classify(tc.tree)
			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
