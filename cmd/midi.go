package main

import (
	"fmt"

	"github.com/berquerant/crd/midix"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(midiCmd)
	midiCmd.AddCommand(midiCmdPort)
	midiCmdPort.AddCommand(midiCmdPortIn, midiCmdPortOut)
}

var midiCmd = &cobra.Command{
	Use:   "midi",
	Short: `midi util`,
}

var midiCmdPort = &cobra.Command{
	Use: "port",
}

var midiCmdPortOut = &cobra.Command{
	Use:   "out",
	Short: `show out ports`,
	RunE: func(_ *cobra.Command, _ []string) error {
		for _, x := range midix.GetOutPortNames() {
			fmt.Println(x)
		}
		return nil
	},
}

var midiCmdPortIn = &cobra.Command{
	Use:   "in",
	Short: `show in ports`,
	RunE: func(_ *cobra.Command, _ []string) error {
		for _, x := range midix.GetInPortNames() {
			fmt.Println(x)
		}
		return nil
	},
}
