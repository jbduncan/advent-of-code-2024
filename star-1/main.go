package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
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

	s := string(b)

	var columnA []int64
	var columnB []int64
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimFunc(line, unicode.IsSpace)
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		first, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return err
		}
		columnA = append(columnA, first)
		second, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return err
		}
		columnB = append(columnB, second)
	}

	slices.Sort(columnA)
	slices.Sort(columnB)

	var totalDistance int64
	for i, first := range columnA {
		second := columnB[i]

		distance := first - second
		if distance < 0 {
			distance = -distance
		}

		totalDistance += distance
	}

	_, _ = fmt.Fprintf(stdout, "%v\n", totalDistance)
	return nil
}
