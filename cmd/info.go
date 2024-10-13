package main

import (
	"fmt"

	"github.com/berquerant/crd/desc"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/op"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(infoCmdAttr, infoCmdChord, infoKeyCmd)

	infoCmdAttr.AddCommand(infoCmdAttrDescribe)
	setRootNoteFlag(infoCmdAttrDescribe)
	infoCmdAttrDescribe.Flags().StringP("target", "t", "", "attribute name")
	infoCmdAttrDescribe.Flags().BoolP("precedeSharp", "s", false, "indicate applied note using sharp")

	infoKeyCmd.AddCommand(infoKeyCmdDescribe, infoKeyCmdList, infoKeyCmdConv)
	setKeyPersistentFlag(infoKeyCmd)
	infoKeyCmdConv.Flags().StringP("command", "c", "", "conversions")
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

var infoCmdAttrDescribe = &cobra.Command{
	Use:   "describe",
	Short: "describe attribute",
	Long: `describe attribute

Examples:
# describe the result of applying Minor7 to C#
crd info attr describe -t "Minor7" -r "C#"
# describe the result of applying Minor7 to C# using sharp
crd info attr describe -t "Minor7" -r "C#" -s
`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		mapper, err := newChordMap(cmd)
		if err != nil {
			return err
		}

		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		root, err := getRootNote(cmd)
		if err != nil {
			return err
		}
		attr, _ := cmd.Flags().GetString("target")
		precedeSharp, _ := cmd.Flags().GetBool("precedeSharp")
		attrInfo, err := desc.NewAttribute(mapper).Describe(attr, root, precedeSharp)
		if err != nil {
			return err
		}

		b, err := yaml.Marshal(attrInfo)
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

var infoKeyCmd = &cobra.Command{
	Use:   "key",
	Short: `key info`,
}

var infoKeyCmdConv = &cobra.Command{
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

var infoKeyCmdList = &cobra.Command{
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

var infoKeyCmdDescribe = &cobra.Command{
	Use:   "describe",
	Short: `describe key`,
	Long: `describe key

Examples:
  crd info key describe --key "A"
  crd info key describe --key "C#m"`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		scale, err := getScale(cmd)
		if err != nil {
			return err
		}

		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		r := desc.NewKey().Describe(scale)

		b, err := yaml.Marshal(r)
		if err != nil {
			return err
		}
		_, err = out.Write(b)
		return err
	},
}
