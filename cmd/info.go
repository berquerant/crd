package main

import (
	"fmt"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/op"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(infoCmdAttr, infoCmdChord)
	infoCmd.AddCommand(keyCmd)
	keyCmd.AddCommand(keyCmdDescribe, keyCmdList, keyCmdConv)
	setKeyPersistentFlag(keyCmd)
	keyCmdConv.Flags().StringP("command", "c", "", "conversions")
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show definitions",
}

var infoCmdAttr = &cobra.Command{
	Use:   "attr",
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
	Use:   "chord",
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

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: `key info`,
}

var keyCmdConv = &cobra.Command{
	Use:   "conv [FILE]",
	Short: `convert key`,
	Long: `convert key

Examples:
# parallel of C
crd info key conv --key "C" -c "p"
# relative of C
crd info key conv --key "C" -c "r"
# dominant of C
crd info key conv --key "C" -c "d"
# subdominant of C
crd info key conv --key "C" -c "s"
# parallel of subdominant of C
crd info key conv --key "C" -c "ps"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		command, _ := cmd.Flags().GetString("command")
		if command == "" {
			return errorx.Invalid("command required")
		}

		scale, err := getScale(cmd)
		if err != nil {
			return err
		}

		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		commandToConversion := func(c rune) op.KeyConversion {
			switch c {
			case 'p':
				return op.ParallelKey
			case 'r':
				return op.RelativeKey
			case 'd':
				return op.DominantKey
			case 's':
				return op.SubDominantKey
			default:
				return op.UnknownKeyConversion
			}
		}
		conversions := make([]op.KeyConversion, len(command))
		for i, c := range command {
			conversions[i] = commandToConversion(c)
		}

		circle := op.NewCircleOfFifth()
		result, err := op.KeyConversionChain(conversions).Convert(circle, scale.Key)
		if err != nil {
			return err
		}

		for x := range result.Keys().All() {
			if _, err := fmt.Fprintf(out, "%v\n", x); err != nil {
				return err
			}
		}
		return nil
	},
}

var keyCmdList = &cobra.Command{
	Use:   "list",
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
	Use:   "describe",
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
