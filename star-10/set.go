package main

import (
	"iter"
	"maps"
)

type set struct {
	s map[string]struct{}
}

func makeSet() set {
	return set{
		s: make(map[string]struct{}),
	}
}

func (s set) add(value string) {
	s.s[value] = struct{}{}
}

func (s set) has(value string) bool {
	_, ok := s.s[value]
	return ok
}

func (s set) all() iter.Seq[string] {
	return maps.Keys(s.s)
}
