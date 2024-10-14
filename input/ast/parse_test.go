package ast_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/berquerant/crd/input/ast"
	"github.com/berquerant/ybase"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// ast.SetDebug(1)
	// slog.SetLogLoggerLevel(slog.LevelDebug)

	for _, tc := range []struct {
		title string
		input string
		want  string
	}{
		{
			title: "Comment with base",
			input: `G_7/B[1] ; dominant motion
C[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "G"
      },
      "symbol": {
        "symbol": "7"
      },
      "base": {
        "degree": {
          "degree": "B"
        }
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    },
    {
      "degree": {
        "degree": "C"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "Comment",
			input: `G_7[1] ; dominant motion
C[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "G"
      },
      "symbol": {
        "symbol": "7"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    },
    {
      "degree": {
        "degree": "C"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "Chords and R",
			input: `6bmaj7[4] 5[4] 1_7/5[2] R[2]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "6",
        "accidental": "b"
      },
      "symbol": {
        "symbol": "maj7"
      },
      "values": {
        "values": [
          {
            "num": "4"
          }
        ]
      }
    },
    {
      "degree": {
        "degree": "5"
      },
      "values": {
        "values": [
          {
            "num": "4"
          }
        ]
      }
    },
    {
      "degree": {
        "degree": "1"
      },
      "symbol": {
        "symbol": "7"
      },
      "base": {
        "degree": {
          "degree": "5"
        }
      },
      "values": {
        "values": [
          {
            "num": "2"
          }
        ]
      }
    },
    {
      "values": {
        "values": [
          {
            "num": "2"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "With symbol and accidental and base",
			input: `Dbmaj7/F[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "D",
        "accidental": "b"
      },
      "symbol": {
        "symbol": "maj7"
      },
      "base": {
        "degree": {
          "degree": "F"
        }
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "With symbol and accidental",
			input: `Dbmaj7[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "D",
        "accidental": "b"
      },
      "symbol": {
        "symbol": "maj7"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "With base",
			input: `1/3[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "1"
      },
      "base": {
        "degree": {
          "degree": "3"
        }
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "With underscore",
			input: `D_7[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "D"
      },
      "symbol": {
        "symbol": "7"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "With symbol",
			input: `Dmaj7[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "D"
      },
      "symbol": {
        "symbol": "maj7"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "With accidental",
			input: `D#[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "D",
        "accidental": "#"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "Chord and Rest with meta",
			input: `C[1,1/2]{c=1} R[3]{d=meta,e=3}`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "C"
      },
      "values": {
        "values": [
          {
            "num": "1"
          },
          {
            "num": "1",
            "denom": "2"
          }
        ]
      },
      "meta": {
        "data": [
          {
            "key": "c",
            "value": "1"
          }
        ]
      }
    },
    {
      "values": {
         "values": [
           {
             "num": "3"
           }
         ]
      },
      "meta": {
        "data": [
          {
            "key": "d",
            "value": "meta"
          },
          {
            "key": "e",
            "value": "3"
          }
        ]
      }
    }
  ]
}`,
		},
		{
			title: "Chord and Rest",
			input: `C[1,1/2] R[3]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "C"
      },
      "values": {
        "values": [
          {
            "num": "1"
          },
          {
            "num": "1",
            "denom": "2"
          }
        ]
      }
    },
    {
      "values": {
         "values": [
           {
             "num": "3"
           }
         ]
       }
    }
  ]
}`,
		},
		{
			title: "single Rest",
			input: `R[2]`,
			want: `{
  "list": [
     {
       "values": {
         "values": [
           {
             "num": "2"
           }
         ]
       }
     }
  ]
}`,
		},
		{
			title: "single Chord",
			input: `C[1]`,
			want: `{
  "list": [
    {
      "degree": {
        "degree": "C"
      },
      "values": {
        "values": [
          {
            "num": "1"
          }
        ]
      }
    }
  ]
}`,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			var want map[string]any
			if !assert.Nil(t, json.Unmarshal([]byte(tc.want), &want)) {
				return
			}

			lex := ast.NewLexer(bytes.NewBufferString(tc.input))
			_ = ast.Parse(lex)
			if !assert.Nil(t, lex.Err(), "%s", lex.Err()) {
				return
			}
			got := lex.Result

			visitor := ast.NewMapVisitor(func(token ybase.Token) (any, bool) {
				if token == nil {
					return nil, false
				}
				return token.Value(), true
			})
			got.Accept(visitor)

			gotBytes, err := json.Marshal(visitor.Result())
			if !assert.Nil(t, err) {
				return
			}
			var result map[string]any
			if !assert.Nil(t, json.Unmarshal([]byte(gotBytes), &result)) {
				return
			}

			assert.Equal(t, want, result)

			{
				b, _ := json.Marshal(got)
				t.Logf("GOT: %s", b)
			}
			{
				b, _ := json.Marshal(result)
				t.Logf("RESULT: %s", b)
			}
		})
	}
}
