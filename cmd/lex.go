package cmd

import (
	"os"

	"github.com/berquerant/crd/cc"
	"github.com/spf13/cobra"
)

var lexCommand = &cobra.Command{
	Use:     "lex",
	Short:   "Lex stdin.",
	Long:    "Lex stdin and print the result.",
	Example: "echo SCORE | crd lex [flags]",
	RunE: func(cmd *cobra.Command, _ []string) error {
		verbose, _ := cmd.Flags().GetInt("verbose")
		l := cc.NewLexer(os.Stdin)
		l.Debug(verbose)
		debugger := cc.NewDebugger(l)
		debugger.Lex()
		return nil
	},
}

func init() {
	rootCommand.AddCommand(lexCommand)
}
