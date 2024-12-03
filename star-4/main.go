package main

import (
	"fmt"
	"io"
	"iter"
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

	var result int

	for _, report := range strings.Split(s, "\n") {
		report = strings.TrimFunc(report, unicode.IsSpace)
		if len(report) == 0 {
			continue
		}

		levels, err := toInts(strings.Fields(report))
		if err != nil {
			return err
		}

		for ls := range problemDampenedLevels(levels) {
			if isSafe(ls) {
				result++
				break
			}
		}
	}

	_, _ = fmt.Fprintf(stdout, "%v\n", result)
	return nil
}

func isSafe(levels []int) bool {
	if !consecutiveDiffLessThanFour(levels) {
		return false
	}
	if isStrictlySortedAscending(levels) {
		return true
	}
	if isStrictlySortedDescending(levels) {
		return true
	}
	return false
}

func problemDampenedLevels(levels []int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		if !yield(levels) {
			return
		}

		for i := 0; i < len(levels); i++ {
			next := slices.Clone(levels)
			next = slices.Delete(next, i, i+1)

			if !yield(next) {
				return
			}
		}
	}
}

func isStrictlySortedAscending(values []int) bool {
	for i := 0; i < len(values)-1; i++ {
		if values[i] >= values[i+1] {
			return false
		}
	}
	return true
}

func isStrictlySortedDescending(values []int) bool {
	for i := 0; i < len(values)-1; i++ {
		if values[i] <= values[i+1] {
			return false
		}
	}
	return true
}

func toInts(values []string) ([]int, error) {
	result := make([]int, 0, len(values))
	for _, value := range values {
		toInt, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		result = append(result, toInt)
	}
	return result, nil
}

func consecutiveDiffLessThanFour(values []int) bool {
	if len(values) == 0 || len(values) == 1 {
		return true
	}

	for i := 0; i < len(values)-1; i++ {
		if abs(values[i]-values[i+1]) > 3 {
			return false
		}
	}
	return true
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
