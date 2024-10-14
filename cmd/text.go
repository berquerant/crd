package main

import (
	"fmt"
	"io"

	"github.com/berquerant/crd/astconv"
	"github.com/berquerant/crd/input"
	"github.com/berquerant/crd/input/ast"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(textCmd)
	textCmd.AddCommand(textCmdParse, textCmdConv)
	textCmdConv.AddCommand(textCmdConvDegree, textCmdConvSyllable)
	setKeyPersistentFlag(textCmdConvSyllable)
}

var textCmd = &cobra.Command{
	Use:   "text",
	Short: `text processor`,
}

var textCmdConv = &cobra.Command{
	Use:   "conv",
	Short: `convert text to yaml`,
	Long: `convert text to yaml

Spaces, newlines, any text between ; and a newline are ignored.
Syntax details: input/ast/chords.y.

Examples:
# C triad, 1 beat
C[1]
# Bb minor triad, 2 beats
Bbm[2]
# D A7/E E Rest (D major)
# '_' is required when chord symbol is a number
echo 'D[1] A_7/E[1] E[2] R[1]' | crd text conv syllable --key D | crd write midi -o out.midi
# metadata
# modulation to A minor
C[1]{key=Am}
# modulation and bpm change
C[1]{key=Am,bpm=200}
# available metadata keys:
# txt: text
C[1]{txt=message}
# lic: lyric
C[1]{lic=some lyric}
# mrk: marker
C[1]{mrk=some marker}
# bpm: BPM
C[1]{bpm=180}
# vel: velocity
C[1]{vel=ff}
# mtr: meter
C[1]{mtr=5/4}`,
}

var textCmdConvSyllable = &cobra.Command{
	Use:   "syllable [FILE]",
	Short: `convert syllable chords to degree yaml`,
	Long: `convert syllable chords to degree yaml

Examples:
  # D A7 E (C major)
  # 1 1 2 (beat)
  echo 'D[1] A_7[1] E[2]' | crd text conv syllable
  # D A7/E E (D major)
  echo 'D[1] A_7/E[1] E[2]' | crd text conv syllable --key D`,
	RunE: func(cmd *cobra.Command, args []string) error {
		scale, err := getScale(cmd)
		if err != nil {
			return err
		}

		tArgs, err := newTextCmdArgs(cmd, args)
		if err != nil {
			return err
		}
		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		converter := astconv.NewSyllableASTConverter(scale)
		return tArgs.convert(out, converter)
	},
}

var textCmdConvDegree = &cobra.Command{
	Use:   "degree [FILE]",
	Short: `convert degree chords to degree yaml`,
	Long: `convert degree chords to degree yaml

Examples:
  # II VI7/V III
  # 1 1 2 (beat)
  echo '2[1] 6_7/5[1] 3[2] | crd test conv degree'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tArgs, err := newTextCmdArgs(cmd, args)
		if err != nil {
			return err
		}
		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		converter := astconv.NewDegreeASTConverter()
		return tArgs.convert(out, converter)
	},
}

func (args textCmdArgs) convert(w io.Writer, converter astconv.Converter) error {
	if _, err := astconv.NewASTClassifier().Classify(args.tree); err != nil {
		return err
	}

	result := make([]*input.Instance, len(args.tree.List))
	for i, x := range args.tree.List {
		y, err := converter.Convert(x)
		if err != nil {
			return fmt.Errorf("%w: ChordOrRest at index %d", err, i)
		}
		result[i] = y
	}

	b, err := yaml.Marshal(result)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

var textCmdParse = &cobra.Command{
	Use:   "parse [FILE]",
	Short: `parse text as AST`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tArgs, err := newTextCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		return writeYamlOutput(cmd, tArgs.tree)
	},
}

type textCmdArgs struct {
	tree *ast.ChordList
}

func newTextCmdArgs(_ *cobra.Command, args []string) (*textCmdArgs, error) {
	tree, err := parseTextFromArgs(args)
	if err != nil {
		return nil, err
	}

	return &textCmdArgs{
		tree: tree,
	}, nil
}

func parseTextFromArgs(args []string) (*ast.ChordList, error) {
	var result *ast.ChordList
	if err := readFileOrStdinFromArgs(args, func(r io.ReadCloser) error {
		list, err := parseText(r)
		if err != nil {
			return err
		}
		result = list
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}
