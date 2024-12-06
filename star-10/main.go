package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sort"
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
		if !isValidSubTopologicalOrdering(ordering, rules) {
			newOrdering := fix(ordering, rules)
			mid := middle(newOrdering)
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

func isValidSubTopologicalOrdering(ordering []string, rules *graph) bool {
	for i := 0; i < len(ordering)-1; i++ {
		if !rules.successors(ordering[i]).has(ordering[i+1]) {
			return false
		}
	}
	return true
}

func fix(ordering []string, rules *graph) []string {
	// Valid solution 1
	//subGraph := newGraph()
	//for source := range rules.nodes() {
	//	if slices.Contains(ordering, source) {
	//		for target := range rules.successors(source).all() {
	//			if slices.Contains(ordering, target) {
	//				subGraph.putEdge(source, target)
	//			}
	//		}
	//	}
	//}
	//return topologicalOrdering(subGraph)

	// Valid solution 2
	ordering = slices.Clone(ordering)
	sort.Sort(&byRules{
		ordering: ordering,
		rules:    rules,
	})
	return ordering
}

type byRules struct {
	ordering []string
	rules    *graph
}

func (b *byRules) Len() int {
	return len(b.ordering)
}

func (b *byRules) Less(i, j int) bool {
	first := b.ordering[i]
	second := b.ordering[j]
	for source := range b.rules.nodes() {
		for target := range b.rules.successors(source).all() {
			if first == source && second == target {
				return true
			}
		}
	}
	return false
}

func (b *byRules) Swap(i, j int) {
	b.ordering[i], b.ordering[j] = b.ordering[j], b.ordering[i]
}

func topologicalOrdering(g *graph) []string {
	// Kahn's algorithm
	var result []string
	roots := rootsOf(g)
	nonRoots := nonRootsOf(g)
	for !roots.isEmpty() {
		next := roots.dequeue()
		result = append(result, next)
		for succ := range g.successors(next).all() {
			nonRoots.removeOne(succ)
			if !nonRoots.has(succ) {
				roots.enqueue(succ)
			}
		}
	}
	if !nonRoots.isEmpty() {
		panic("no topological ordering")
	}
	return result
}

func rootsOf(g *graph) queue {
	result := queue{}
	for node := range g.nodes() {
		if g.inDegree(node) == 0 {
			result.enqueue(node)
		}
	}
	return result
}

func nonRootsOf(g *graph) multiset {
	result := newMultiset()
	for node := range g.nodeToData {
		if g.inDegree(node) > 0 {
			result.add(node, g.inDegree(node))
		}
	}
	return result
}
