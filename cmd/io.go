package main

import (
	"bytes"
	"io"
	"os"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/input/ast"
	"github.com/berquerant/crd/util"
	"github.com/spf13/cobra"
)

//
// common io utilities
//

func getOutput(cmd *cobra.Command) (io.WriteCloser, error) {
	x, _ := cmd.Flags().GetString("output")
	if x == "" {
		return os.Stdout, nil
	}

	f, err := os.Create(x)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func readFileOrStdinFromArgs(args []string, f func(io.ReadCloser) error) error {
	switch len(args) {
	case 0:
		return readFileOrStdin("", f)
	case 1:
		return readFileOrStdin(args[0], f)
	default:
		return errorx.Unexpected(`Should only specify at most one file. If the filename is " - " or if it is not specified, it will read from standard input`)
	}
}

func readFileOrStdin(name string, f func(io.ReadCloser) error) error {
	switch name {
	case "-", "":
		return f(os.Stdin)
	default:
		fp, err := os.Open(name)
		if err != nil {
			return err
		}
		return f(fp)
	}
}

func newChordBuilder(cmd *cobra.Command) (*chord.Builder, error) {
	var (
		builder    = chord.NewBuilder()
		attrFiles  = getAttributeFlag(cmd)
		chordFiles = getChordFlag(cmd)
	)

	for _, x := range chord.BasicAttributes() {
		builder.Attribute(x)
	}
	for _, x := range chord.BasicChords() {
		builder.Chord(x)
	}
	for _, f := range attrFiles {
		attrs, err := util.OpenAndParse(f, chord.ParseAttributes)
		if err != nil {
			return nil, err
		}
		for _, x := range attrs {
			builder.Attribute(x)
		}
	}
	for _, f := range chordFiles {
		chords, err := util.OpenAndParse(f, chord.ParseChords)
		if err != nil {
			return nil, err
		}
		for _, x := range chords {
			builder.Chord(x)
		}
	}

	return builder, nil
}

func newChordMap(cmd *cobra.Command) (chord.Mapper, error) {
	b, err := newChordBuilder(cmd)
	if err != nil {
		return nil, err
	}
	return b.Build()
}

func parseText(r io.Reader) (*ast.ChordList, error) {
	lex := ast.NewLexer(r)
	_ = ast.Parse(lex)
	return lex.Result, lex.Err()
}

func parseTextOneChordSymbol(r io.Reader) (*ast.Chord, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	buf.WriteString("[1]") // expect C, Dm to C[1], Dm[1]
	tree, err := parseText(buf)
	if err != nil {
		return nil, err
	}

	if len(tree.List) != 1 {
		return nil, errorx.Invalid("Parse one chord: %s", b)
	}
	c := tree.List[0]
	cd, ok := c.(*ast.Chord)
	if !ok {
		return nil, errorx.Invalid("Parse one chord: %s", b)
	}
	return cd, nil
}
