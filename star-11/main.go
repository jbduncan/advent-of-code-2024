package main

import (
	"errors"
	"fmt"
	"io"
	"iter"
	"log"
	"maps"
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
	stateMachine, err := newGuardStateMachine(g)
	if err != nil {
		return err
	}
	result := stateMachine.run()

	_, _ = fmt.Fprintf(stdout, "%v\n", result)
	return nil
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

type set[T comparable] struct {
	s map[T]struct{}
}

func makeSet[T comparable]() set[T] {
	return set[T]{
		s: make(map[T]struct{}),
	}
}

func (s set[T]) add(value T) {
	s.s[value] = struct{}{}
}

func (s set[T]) has(value T) bool {
	_, ok := s.s[value]
	return ok
}

func (s set[T]) all() iter.Seq[T] {
	return maps.Keys(s.s)
}

func (s set[T]) len() int {
	return len(s.s)
}

type cell struct {
	row    int
	column int
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

func (d direction) turnRight() direction {
	return (d + 1) % 4
}

type guardStateMachine struct {
	g        grid
	position cell
	d        direction
}

func newGuardStateMachine(g grid) (*guardStateMachine, error) {
	var position *cell
	for row, column := range g.all() {
		if g.at(row, column) == '^' {
			position = &cell{
				row:    row,
				column: column,
			}
		}
	}
	if position == nil {
		return nil, errors.New("guard position cannot be found")
	}
	return &guardStateMachine{
		g:        g,
		position: *position,
		d:        north,
	}, nil
}

func (m *guardStateMachine) run() int {
	visitedCells := makeSet[cell]()
	for {
		visitedCells.add(m.position)
		c := m.nextCell()
		if c == nil {
			return visitedCells.len()
		}
		next := m.g.at((*c).row, (*c).column)
		switch {
		case next == '.' || next == '^':
			visitedCells.add(m.position)
			m.move()
			continue
		case next == '#':
			m.d = m.d.turnRight()
			continue
		default:
			panic("unexpected")
		}
	}
}

func (m *guardStateMachine) move() {
	c := m.nextCell()
	m.position = cell{
		row:    c.row,
		column: c.column,
	}
}

func (m *guardStateMachine) nextCell() *cell {
	var r, c int
	switch m.d {
	case north:
		r, c = m.position.row-1, m.position.column
	case east:
		r, c = m.position.row, m.position.column+1
	case south:
		r, c = m.position.row+1, m.position.column
	case west:
		r, c = m.position.row, m.position.column-1
	default:
		panic("unexpected")
	}
	if m.g.has(r, c) {
		return &cell{
			row:    r,
			column: c,
		}
	}
	return nil
}

func ptr(r rune) *rune {
	return &r
}
