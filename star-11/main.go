package main

import (
	"fmt"
	"io"
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
	_ = s

	var result int

	_, _ = fmt.Fprintf(stdout, "%v\n", result)
	return nil
}
