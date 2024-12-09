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
	for row := 1; row < g.rowCount()-1; row++ {
		for column := 1; column < g.columnCount()-1; column++ {
			if isAnXmas(g, row, column) {
				result++
			}
		}
	}
	return result
}

func isAnXmas(g grid, row, column int) bool {
	if g.at(row, column) != 'A' {
		return false
	}

	areDiagonalsEqualTo := func(topLeft, topRight, bottomLeft, bottomRight rune) bool {
		return g.at(row-1, column-1) == topLeft &&
			g.at(row-1, column+1) == topRight &&
			g.at(row+1, column-1) == bottomLeft &&
			g.at(row+1, column+1) == bottomRight
	}

	return areDiagonalsEqualTo('M', 'M', 'S', 'S') ||
		areDiagonalsEqualTo('M', 'S', 'M', 'S') ||
		areDiagonalsEqualTo('S', 'S', 'M', 'M') ||
		areDiagonalsEqualTo('S', 'M', 'S', 'M')
}

type grid struct {
	g [][]rune
}

func makeGrid(s string) grid {
	return grid{
		g: mapSlice(strings.Split(s, "\n"), runes),
	}
}

func runes(value string) []rune {
	return []rune(value)
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

func (g *grid) rowCount() int {
	return len(g.g)
}

func (g *grid) columnCount() int {
	return len(g.g[0])
}

func mapSlice[I, O any](values []I, f func(value I) O) []O {
	var result []O
	for _, value := range values {
		result = append(result, f(value))
	}
	return result
}
