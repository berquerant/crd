package main

import (
	"fmt"

	"github.com/berquerant/crd/chord"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.AddCommand(genCmdAttr)
	genCmdAttr.Flags().UintP("maxDegree", "d", 20, "max degrees to generate")
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate codes and files",
}

var genCmdAttr = &cobra.Command{
	Use:   "attr",
	Short: "generate attributes",
	RunE: func(cmd *cobra.Command, _ []string) error {
		output, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer output.Close()
		max, _ := cmd.Flags().GetUint("maxDegree")
		attrs := chord.GenerateAttributes(max)
		b, err := yaml.Marshal(attrs)
		if err != nil {
			return err
		}
		fmt.Fprintf(output, "%s", b)
		return nil
	},
}
