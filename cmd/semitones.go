package cmd

import (
	"os"

	"github.com/berquerant/crd/cc"
	"github.com/spf13/cobra"
)

var semitonesCommand = &cobra.Command{
	Use:     "semitones",
	Short:   "Print semitones.",
	Long:    "Print the chords of the score as the semitone lists.",
	Example: "echo SCORE | crd semitones [flags]",
	RunE: func(cmd *cobra.Command, _ []string) error {
		verbose, _ := cmd.Flags().GetInt("verbose")
		l := cc.NewLexer(os.Stdin)
		l.Debug(verbose)
		debugger := cc.NewDebugger(l)
		debugger.Semitones()
		return nil
	},
}

func init() {
	rootCommand.AddCommand(semitonesCommand)
}
