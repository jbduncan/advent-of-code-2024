package main

import (
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func run(stdout io.Writer, args []string) error {
	b, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	s := strings.Trim(string(b), "\n")

	g := makeGrid(s)
	result := findXmasOccurrencesCount(g)
	_, _ = fmt.Fprintf(stdout, "%v\n", result)
	return nil
}

func findXmasOccurrencesCount(g grid) int {
	var result int
	for row, column := range g.all() {
		if g.at(row, column) == 'X' {
			for rowOffset, columnOffset := range neighbourCellOffsets(g, row, column) {
				if occurs(
					g,
					row+rowOffset,
					column+columnOffset,
					rowOffset,
					columnOffset,
					xmasRunes,
					1,
				) {
					result++
				}
			}
		}
	}
	return result
}

var xmasRunes = []rune("XMAS")

func occurs(g grid, row, column, rowOffset, columnOffset int, needle []rune, index int) bool {
	for r, c, i := row, column, index; //
		i < len(xmasRunes);            //
	r, c, i = r+rowOffset, c+columnOffset, i+1 {
		if !g.has(r, c) || g.at(r, c) != needle[i] {
			return false
		}
	}
	return true
}

func neighbourCellOffsets(g grid, row, column int) iter.Seq2[int, int] {
	return func(yield func(row int, column int) bool) {
		for rowOffset := -1; rowOffset <= 1; rowOffset++ {
			for columnOffset := -1; columnOffset <= 1; columnOffset++ {
				if rowOffset == 0 && columnOffset == 0 {
					continue
				}

				if !g.has(row+rowOffset, column+columnOffset) {
					continue
				}

				if !yield(rowOffset, columnOffset) {
					return
				}
			}
		}
	}
}

type grid struct {
	g [][]rune
}

func makeGrid(s string) grid {
	result := grid{}
	for _, line := range strings.Split(s, "\n") {
		result.g = append(result.g, []rune(line))
	}
	return result
}

func (g *grid) has(row, column int) bool {
	return row >= 0 &&
		row < len(g.g) &&
		column >= 0 &&
		column < len(g.g[0])
}

func (g *grid) at(row, column int) rune {
	return g.g[row][column]
}

func (g *grid) all() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for r, row := range g.g {
			for c := range row {
				if !yield(r, c) {
					return
				}
			}
		}
	}
}
