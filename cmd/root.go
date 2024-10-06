package main

import (
	"log/slog"
	"os"

	"github.com/berquerant/crd/input/ast"
	"github.com/berquerant/crd/logx"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug")
	rootCmd.PersistentFlags().StringP("output", "o", "", "output file")
	setAttributePersistentFlag(rootCmd)
	setChordPersistentFlag(rootCmd)
}

var rootCmd = &cobra.Command{
	Use:   "crd",
	Short: "text2midi",
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		debugEnabled, _ := cmd.Flags().GetBool("debug")
		logLevel := slog.LevelInfo
		if debugEnabled {
			logLevel = slog.LevelDebug
			ast.SetDebug(1)
		}
		logx.Setup(os.Stderr, logLevel)
	},
}
