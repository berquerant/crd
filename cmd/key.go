package main

import (
	"github.com/berquerant/crd/op"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(keyCmd)
	keyCmd.AddCommand(keyCmdDescribe, keyCmdList)
	setKeyPersistentFlag(keyCmd)
}

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: `key info`,
}

var keyCmdList = &cobra.Command{
	Use:   "list [FILE]",
	Short: `list all keys`,
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		b, err := yaml.Marshal(op.AllScales())
		if err != nil {
			return err
		}
		_, err = out.Write(b)
		return err
	},
}

var keyCmdDescribe = &cobra.Command{
	Use:   "describe [FILE]",
	Short: `describe key`,
	Long: `describe key

Examples:
  crd key describe --key "A"
  crd key describe --key "C#m"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		scale, err := getScale(cmd)
		if err != nil {
			return err
		}

		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		b, err := yaml.Marshal(scale)
		if err != nil {
			return err
		}
		_, err = out.Write(b)
		return err
	},
}
