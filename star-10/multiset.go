package main

import "fmt"

type multiset struct {
	m map[string]int
}

func newMultiset() multiset {
	return multiset{m: map[string]int{}}
}

func (m multiset) add(value string, count int) {
	if count <= 0 {
		panic(fmt.Sprintf("count is non-positive: %d", count))
	}

	m.m[value] += count
}

func (m multiset) has(value string) bool {
	return m.count(value) > 0
}

func (m multiset) count(value string) int {
	return m.m[value]
}

func (m multiset) removeOne(value string) {
	if _, ok := m.m[value]; !ok {
		panic("multiset does not have value")
	}

	m.m[value]--
	if m.m[value] == 0 {
		delete(m.m, value)
	}
}

func (m multiset) isEmpty() bool {
	return len(m.m) == 0
}
