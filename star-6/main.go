package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		log.Fatalf("%v\n", err)
	}
}

var mulInstructionRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func run(stdout io.Writer, args []string) error {
	b, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	s := string(b)

	s = dontDoBlocksRemoved(s)

	matches := mulInstructionRegex.FindAllStringSubmatch(s, -1)

	var result int
	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		}
		second, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}

		result += first * second
	}

	_, _ = fmt.Fprintf(stdout, "%v\n", result)
	return nil
}

func dontDoBlocksRemoved(s string) string {
	var w strings.Builder
	start := 0
	do := false
	dont := false
	for i := range s {
		if !dont {
			endFound := i == len(s)-1
			dontFound := s[i:min(i+7, len(s)-1)] == "don't()"
			if endFound || dontFound {
				w.WriteString(s[start:i])
				do = false
				dont = true
				continue
			}
		}

		if !do {
			doFound := s[i:min(i+4, len(s)-1)] == "do()"
			if doFound {
				start = i + 4
				do = true
				dont = false
			}
		}
	}
	return w.String()
}
