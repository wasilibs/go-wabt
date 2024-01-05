package wabt

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestRuns(t *testing.T) {
	// Just do simple sanity checking that the executables can run.

	dirs, err := os.ReadDir("cmd")
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		name := dir.Name()
		t.Run(name, func(t *testing.T) {
			out := bytes.Buffer{}
			cmd := exec.Command("go", "run", "./cmd/"+name)
			cmd.Stdout = &out
			cmd.Stderr = &out
			// All commands expect a file, so they always exit with an error, if the error message
			// has the correct command name, it's probably OK.
			err := cmd.Run()
			if err == nil {
				t.Error("expected error, got nil")
			}
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				if exitErr.ExitCode() != 1 {
					t.Errorf("expected exit code 1, got %d", exitErr.ExitCode())
				}
			} else {
				t.Errorf("expected *exec.ExitError, got %T", err)
			}

			if !strings.Contains(out.String(), name) {
				t.Errorf("expected error message to contain %q, got %q", name, out.String())
			}
		})
	}
}
