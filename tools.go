//go:build tools
// +build tools

package main

import (
	_ "github.com/berquerant/marker"
	_ "github.com/berquerant/mkvisitor"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "golang.org/x/tools/cmd/goyacc"
	_ "golang.org/x/tools/cmd/stringer"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "gotest.tools/gotestsum"
)
