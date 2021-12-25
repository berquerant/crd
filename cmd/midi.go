package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/berquerant/crd/cc"
	"github.com/berquerant/crd/logger"
	"github.com/berquerant/crd/midi"
	"github.com/spf13/cobra"
)

const (
	defaultBPM   = 120
	defaultMeter = "4/4"
)

var midiCommand = &cobra.Command{
	Use:     "midi",
	Short:   "Generate a midi file.",
	Long:    "Generate a midi file from stdin score.",
	Example: "echo SCORE | crd midi [flags] -o FILE",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var (
			verbose, _ = cmd.Flags().GetInt("verbose")
			meter, _   = cmd.Flags().GetString("meter")
			bpm, _     = cmd.Flags().GetInt("bpm")
			out, _     = cmd.Flags().GetString("out")
		)
		mn, md, ok := parseMeter(meter)
		if !ok {
			return fmt.Errorf("invalid time signature %s", meter)
		}
		// parse input and construct AST
		lexer := cc.NewLexer(os.Stdin)
		lexer.Debug(verbose)
		status := cc.Parse(lexer)
		logger.Get().Info("parser exit with %d", status)
		if err := lexer.Err(); err != nil {
			return fmt.Errorf("parser got error %w", err)
		}
		// generate midi operations based on AST
		w := midi.NewASTWriter(midi.NewWriter())
		if bpm != defaultBPM {
			w.Writer().BPM(bpm)
		}
		if meter != defaultMeter {
			w.Writer().Meter(mn, md)
		}
		for _, n := range lexer.Result().NodeList {
			w.WriteNode(n)
		}
		// write midi file
		return midi.NewFactory(out).WriteSMF(w.Writer().Operations())
	},
}

func init() {
	midiCommand.Flags().StringP("out", "o", "crd.out.mid", "Output filepath.")
	midiCommand.Flags().IntP("bpm", "b", defaultBPM, "BPM.")
	midiCommand.Flags().StringP("meter", "m", defaultMeter, "Time signature.")
	rootCommand.AddCommand(midiCommand)
}

func parseMeter(meter string) (uint8, uint8, bool) {
	v := strings.Split(meter, "/")
	if len(v) != 2 {
		return 0, 0, false
	}
	x, err := strconv.Atoi(v[0])
	if err != nil {
		return 0, 0, false
	}
	y, err := strconv.Atoi(v[1])
	if err != nil {
		return 0, 0, false
	}
	return uint8(x), uint8(y), true
}
