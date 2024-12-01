package main

import (
	"fmt"
	"io"
	"log"
	"os"
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

	var totalSimilarityScore int64
	for _, value := range columnA {
		count := countOf(columnB, value)
		totalSimilarityScore += value * count
	}

	_, _ = fmt.Fprintf(stdout, "%v\n", totalSimilarityScore)
	return nil
}

func countOf(haystack []int64, needle int64) int64 {
	var result int64
	for _, value := range haystack {
		if value == needle {
			result++
		}
	}
	return result
}
