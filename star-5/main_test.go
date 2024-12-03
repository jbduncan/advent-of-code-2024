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
			name:       "Star 5 example input",
			args:       []string{tempFile(t, example)},
			wantStdout: "161\n",
		},
		{
			name:       "Star 5 real input",
			args:       []string{"input.txt"},
			wantStdout: "175615763\n",
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
				t.Errorf("run(): stdout: got %#v, want %#v", gotStdout, tt.wantStdout)
			}
		})
	}
}

var example = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func tempFile(t *testing.T, contents string) (file string) {
	dir := t.TempDir()
	n := filepath.Join(dir, "file")
	if err := os.WriteFile(n, []byte(contents), 0o666); err != nil {
		t.Fatalf("temp file not created: %v", err)
	}
	return n
}
