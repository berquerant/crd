package main_test

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndToEnd(t *testing.T) {
	e := newExecutor(t)
	defer e.close()

	if err := run(os.Stdout, nil, e.cmd, "-h"); err != nil {
		t.Fatalf("%s help %v", e.cmd, err)
	}

	t.Run("text", func(t *testing.T) {
		const (
			degreeText    = `2[1] 6_7[1] 3[2]`
			syllableText  = `D[1] A_7[1] E[2]`
			degreeOnCYaml = `- chord:
    degree: "2"
    name: ""
  values:
    - "1"
- chord:
    degree: "6"
    name: "7"
  values:
    - "1"
- chord:
    degree: "3"
    name: ""
  values:
    - "2"
`
			degreeOnDYaml = `- chord:
    degree: "1"
    name: ""
  values:
    - "1"
- chord:
    degree: "5"
    name: "7"
  values:
    - "1"
- chord:
    degree: "2"
    name: ""
  values:
    - "2"
`
		)
		t.Run("conv", func(t *testing.T) {
			for _, tc := range []struct {
				title string
				input string
				opt   []string
				want  string
			}{
				{
					title: "degree",
					input: degreeText,
					opt:   []string{"degree"},
					want:  degreeOnCYaml,
				},
				{
					title: "syllable on C",
					input: syllableText,
					opt:   []string{"syllable"},
					want:  degreeOnCYaml,
				},
				{
					title: "syllable on D",
					input: syllableText,
					opt:   []string{"syllable", "--key", "D"},
					want:  degreeOnDYaml,
				},
			} {
				t.Run(tc.title, func(t *testing.T) {
					stdin := bytes.NewBufferString(tc.input)
					var stdout bytes.Buffer
					args := append([]string{"text", "conv"}, tc.opt...)
					assert.Nil(t, run(&stdout, stdin, e.cmd, args...))
					assert.Equal(t, tc.want, stdout.String())
				})
			}
		})

		t.Run("parse", func(t *testing.T) {
			for _, tc := range []struct {
				title string
				input string
			}{
				{
					title: "degree",
					input: degreeText,
				},
				{
					title: "syllable",
					input: syllableText,
				},
			} {
				t.Run(tc.title, func(t *testing.T) {
					stdin := bytes.NewBufferString(tc.input)
					assert.Nil(t, run(os.Stdout, stdin, e.cmd, "text", "parse"))
				})
			}
		})
	})

	t.Run("write", func(t *testing.T) {
		const (
			someChordsYaml = `- chord:
    name: MajorTriad
    degree: "1"
  values:
    - "1"
  key: C
- chord:
    name: DominantSeventh
    degree: "5"
  values:
    - "1"
- chord:
    name: MajorTriad
    degree: "4"
  values:
    - "2"`
		)
		t.Run("event", func(t *testing.T) {
			stdin := bytes.NewBufferString(someChordsYaml)
			assert.Nil(t, run(os.Stdout, stdin, e.cmd, "write", "event"))
		})
		t.Run("parse", func(t *testing.T) {
			const (
				parsedYaml = `- chord:
    degree: "1"
    chord:
        name: MajorTriad
        meta:
            display: ""
        attributes:
            - Perfect1
            - Major3
            - Perfect5
    base: "1"
  values:
    - "1"
  key: C
- chord:
    degree: "5"
    chord:
        name: DominantSeventh
        meta:
            display: "7"
        attributes:
            - Minor7
        extends: MajorTriad
    base: "1"
  values:
    - "1"
- chord:
    degree: "4"
    chord:
        name: MajorTriad
        meta:
            display: ""
        attributes:
            - Perfect1
            - Major3
            - Perfect5
    base: "1"
  values:
    - "2"
`
			)
			stdin := bytes.NewBufferString(someChordsYaml)
			var stdout bytes.Buffer
			assert.Nil(t, run(&stdout, stdin, e.cmd, "write", "parse"))
			assert.Equal(t, parsedYaml, stdout.String())
		})
	})

	t.Run("midi", func(t *testing.T) {
		t.Run("port", func(t *testing.T) {
			t.Run("in", func(t *testing.T) {
				assert.Nil(t, run(os.Stdout, nil, e.cmd, "midi", "port", "in"))
			})
			t.Run("out", func(t *testing.T) {
				assert.Nil(t, run(os.Stdout, nil, e.cmd, "midi", "port", "out"))
			})
		})
	})
}

func run(w io.Writer, r io.Reader, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = "."
	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type executor struct {
	dir string
	cmd string
}

func newExecutor(t *testing.T) *executor {
	t.Helper()
	e := &executor{}
	e.init(t)
	return e
}

func (e *executor) init(t *testing.T) {
	t.Helper()
	dir, err := os.MkdirTemp("", "crd")
	if err != nil {
		t.Fatal(err)
	}
	cmd := filepath.Join(dir, "crd")
	// build crd command
	if err := run(os.Stdout, nil, "go", "build", "-o", cmd); err != nil {
		t.Fatal(err)
	}
	e.dir = dir
	e.cmd = cmd
}

func (e *executor) close() {
	os.RemoveAll(e.dir)
}
