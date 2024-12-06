package main

import (
	"fmt"
	"io"
	"iter"
	"log"
	"maps"
	"os"
	"strconv"
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

	parts := strings.SplitN(s, "\n\n", 2)
	rules := toRules(parts[0])
	orderings := toOrderings(parts[1])

	var result int
	for _, ordering := range orderings {
		if rules.isValidSubTopologicalOrdering(ordering) {
			mid := middle(ordering)
			midAsInt, err := strconv.Atoi(mid)
			if err != nil {
				return err
			}
			result += midAsInt
		}
	}

	_, _ = fmt.Fprintf(stdout, "%v\n", result)
	return nil
}

func toRules(s string) *graph {
	g := newGraph()
	for _, rule := range strings.Split(s, "\n") {
		parts := strings.SplitN(rule, "|", 2)
		g.putEdge(parts[0], parts[1])
	}
	return g
}

func toOrderings(s string) [][]string {
	var result [][]string
	for _, ordering := range strings.Split(s, "\n") {
		result = append(result, strings.Split(ordering, ","))
	}
	return result
}

func middle(values []string) string {
	return values[len(values)/2]
}

type set map[string]struct{}

func makeSet() set {
	return make(set)
}

func (s set) add(value string) {
	s[value] = struct{}{}
}

func (s set) has(value string) bool {
	_, ok := s[value]
	return ok
}

func newGraph() *graph {
	return &graph{
		nodeToSuccessors: make(map[string]set),
	}
}

type graph struct {
	nodeToSuccessors map[string]set
}

func (g *graph) addNode(node string) {
	if _, ok := g.nodeToSuccessors[node]; !ok {
		g.nodeToSuccessors[node] = makeSet()
	}
}

func (g *graph) putEdge(source, target string) {
	g.addNode(source)
	g.addNode(target)

	successors := g.nodeToSuccessors[source]
	if !successors.has(target) {
		successors.add(target)
	}
}

func (g *graph) nodes() iter.Seq[string] {
	return maps.Keys(g.nodeToSuccessors)
}

func (g *graph) successors(node string) set {
	successors, ok := g.nodeToSuccessors[node]
	if !ok {
		panic("node is not in graph")
	}

	return successors
}

func (g *graph) isValidSubTopologicalOrdering(ordering []string) bool {
	for i := 0; i < len(ordering)-1; i++ {
		if !g.successors(ordering[i]).has(ordering[i+1]) {
			return false
		}
	}
	return true
}
