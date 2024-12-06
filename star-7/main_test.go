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
			name:       "Star 7 example input 1",
			args:       []string{tempFile(t, example1)},
			wantStdout: "4\n",
		},
		{
			name:       "Star 7 example input 2",
			args:       []string{tempFile(t, example2)},
			wantStdout: "18\n",
		},
		{
			name:       "Star 7 example input 3",
			args:       []string{tempFile(t, example3)},
			wantStdout: "18\n",
		},
		{
			name:       "Star 7 real input",
			args:       []string{"input.txt"},
			wantStdout: "2524\n",
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

var example1 = `
..X...
.SAMX.
.A..A.
XMAS.S
.X....
`

var example2 = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

var example3 = `
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
`

func tempFile(t *testing.T, contents string) (file string) {
	dir := t.TempDir()
	n := filepath.Join(dir, "file")
	if err := os.WriteFile(n, []byte(contents), 0o666); err != nil {
		t.Fatalf("temp file not created: %v", err)
	}
	return n
}
