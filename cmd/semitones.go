package cmd

import (
	"os"

	"github.com/berquerant/crd/cc"
	"github.com/berquerant/crd/note"
	"github.com/spf13/cobra"
)

var semitonesCommand = &cobra.Command{
	Use:     "semitones",
	Short:   "Print semitones.",
	Long:    "Print the chords of the score as the semitone lists.",
	Example: "echo SCORE | crd semitones [flags]",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var (
			verbose, _ = cmd.Flags().GetInt("verbose")
			trans, _   = cmd.Flags().GetInt("trans")
		)
		l := cc.NewLexer(os.Stdin)
		l.Debug(verbose)
		debugger := cc.NewDebugger(l, cc.WithTransposition(note.Semitone(trans)))
		debugger.Semitones()
		return nil
	},
}

func init() {
	semitonesCommand.Flags().IntP("trans", "t", 0, "Transposition")
	rootCommand.AddCommand(semitonesCommand)
}
