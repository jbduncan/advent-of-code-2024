package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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
