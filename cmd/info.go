package main

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(infoCmdAttr, infoCmdChord)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show definitions",
}

var infoCmdAttr = &cobra.Command{
	Use:   "attr [FILE]",
	Short: "list attribute definitions",
	RunE: func(cmd *cobra.Command, _ []string) error {
		builder, err := newChordBuilder(cmd)
		if err != nil {
			return err
		}

		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		b, err := yaml.Marshal(builder.UnwrapAttributes())
		if err != nil {
			return err
		}
		_, err = out.Write(b)
		return err
	},
}

var infoCmdChord = &cobra.Command{
	Use:   "chord [FILE]",
	Short: "list chord definitions",
	RunE: func(cmd *cobra.Command, _ []string) error {
		builder, err := newChordBuilder(cmd)
		if err != nil {
			return err
		}

		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		b, err := yaml.Marshal(builder.UnwrapChords())
		if err != nil {
			return err
		}
		_, err = out.Write(b)
		return err
	},
}
