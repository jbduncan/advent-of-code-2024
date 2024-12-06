package main

import (
	"iter"
	"maps"
)

type nodeData struct {
	inDegree   int
	successors set
}

func newGraph() *graph {
	return &graph{
		nodeToData: make(map[string]*nodeData),
	}
}

type graph struct {
	nodeToData map[string]*nodeData
}

func (g *graph) addNode(node string) {
	if _, ok := g.nodeToData[node]; !ok {
		g.nodeToData[node] = &nodeData{
			successors: makeSet(),
		}
	}
}

func (g *graph) putEdge(source, target string) {
	g.addNode(source)
	g.addNode(target)

	successors := g.nodeToData[source].successors
	if !successors.has(target) {
		successors.add(target)
		g.nodeToData[target].inDegree++
	}
}
func (g *graph) nodes() iter.Seq[string] {
	return maps.Keys(g.nodeToData)
}

func (g *graph) successors(node string) set {
	nd, ok := g.nodeToData[node]
	if !ok {
		panic("node is not in graph")
	}

	return nd.successors
}

func (g *graph) inDegree(node string) int {
	nd, ok := g.nodeToData[node]
	if !ok {
		return 0
	}
	return nd.inDegree
}
