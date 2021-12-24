package cmd

import (
	"os"

	"github.com/berquerant/crd/cc"
	"github.com/spf13/cobra"
)

var unparseCommand = &cobra.Command{
	Use:     "unparse",
	Short:   "Normalize score.",
	Long:    "Parse score and print the score as string.",
	Example: "echo SCORE | crd unparse [flags]",
	RunE: func(cmd *cobra.Command, _ []string) error {
		verbose, _ := cmd.Flags().GetInt("verbose")
		l := cc.NewLexer(os.Stdin)
		l.Debug(verbose)
		debugger := cc.NewDebugger(l)
		debugger.Unparse()
		return nil
	},
}

func init() {
	rootCommand.AddCommand(unparseCommand)
}
