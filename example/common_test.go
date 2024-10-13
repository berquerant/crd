package example_test

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type executor struct {
	cmd string
	dir string
}

const ignoreStdout = "ignore.stdout"

func (e *executor) do(t *testing.T, want string, r io.Reader, args ...string) {
	var w bytes.Buffer
	if !assert.Nil(t, run(&w, r, e.cmd, args...)) {
		return
	}

	got := w.String()
	switch want {
	case ignoreStdout:
		return
	case "":
		t.Log(got)
	default:
		assert.Equal(t, want, got)
	}
}

func (e *executor) close() {
	os.RemoveAll(e.dir)
}

func (e *executor) init(t *testing.T) bool {
	dir := t.TempDir()

	cmd := filepath.Join(dir, "crd")
	if !assert.Nil(t, run(os.Stdout, nil, "go", "build", "-o", cmd, "../cmd/")) {
		return false
	}
	e.dir = dir
	e.cmd = cmd

	return true
}

func run(w io.Writer, r io.Reader, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = "."
	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
