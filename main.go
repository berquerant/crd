package main

import (
	"os"

	"github.com/berquerant/crd/cmd"
	"github.com/berquerant/crd/logger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Get().Error("%v", err)
		os.Exit(1)
	}
}
