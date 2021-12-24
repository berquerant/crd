package cmd

import (
	"github.com/berquerant/crd/logger"
	"github.com/spf13/cobra"
)

var rootCommand = newRootCommand()

func newRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "echo SCORE | crd COMMAND",
		Short: "Generate a midi file.",
		Long: `Generate a midi file or print some info from stdin.

Example of SCORE:
  C[1/2] G[1/2] Am[1/2] Em[1/2] F[1/2] C[1/2] F[1/2] G[1/2]

SCORE is a sequence of chords, chord V forms SCORE like
V V | V V V
Vertical bars, spaces and newlines are only for readability.

A chord is a string formatted like NOTE OPTION [VALUE].
NOTE is a note name, C, D#, Gb, ...
VALUE ia a relative note value, for example 1 means a whole note, 1/2 means a half note.
OPTION is a chord option, like m (minor). Available options are below:

  m    minor triad
  aug  augmented triad
  dim  diminished triad
  7    dominant seventh
  m7   minor seventh
  M7   major seventh
  mM7  minor major seventh
  dim7 diminished seventh
  m7-5 half diminished seventh
  aug7 augmented seventh
  6    add sixth
  m6   add minor sixth
  sus4 suspended forth`,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			verbose, _ := cmd.Flags().GetInt("verbose")
			if verbose < 1 {
				logger.Get().SetLevel(logger.Linfo)
			} else {
				logger.Get().SetLevel(logger.Ldebug)
			}
		},
	}
}

func init() {
	rootCommand.PersistentFlags().IntP("verbose", "v", 0, "Logging verbosity.")
}

func Execute() error {
	return rootCommand.Execute()
}
