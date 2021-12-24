package cmd

import (
	"os"

	"github.com/berquerant/crd/cc"
	"github.com/spf13/cobra"
)

var parseCommand = &cobra.Command{
	Use:     "parse",
	Short:   "Print AST.",
	Long:    "Parse stdin and print the AST.",
	Example: "echo SCORE | crd parse [flags]",
	RunE: func(cmd *cobra.Command, _ []string) error {
		verbose, _ := cmd.Flags().GetInt("verbose")
		l := cc.NewLexer(os.Stdin)
		l.Debug(verbose)
		debugger := cc.NewDebugger(l)
		debugger.Parse()
		return nil
	},
}

func init() {
	rootCommand.AddCommand(parseCommand)
}
