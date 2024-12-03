package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStdout string
		wantErr    error
	}{
		{
			name:       "Star 2 example input",
			args:       []string{tempFile(t, example)},
			wantStdout: "31\n",
		},
		{
			name:       "Star 2 real input",
			args:       []string{"input.txt"},
			wantStdout: "26674158\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &strings.Builder{}
			err := run(stdout, tt.args)
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("run(): error: got %v, want %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("run(): stdout: got %v, want %v", gotStdout, tt.wantStdout)
			}
		})
	}
}

var example = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func tempFile(t *testing.T, contents string) (file string) {
	dir := t.TempDir()
	n := filepath.Join(dir, "file")
	if err := os.WriteFile(n, []byte(contents), 0o666); err != nil {
		t.Fatalf("temp file not created: %v", err)
	}
	return n
}
