package main

import (
	"log/slog"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

func main() {
	defer midi.CloseDriver()
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Err", slog.Any("err", err))
	}
}
