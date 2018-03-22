// Copyright 2013 Steven Le. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package topsort

import (
	"fmt"
	"strings"
)

type Graph struct {
	nodes map[string]node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]node),
	}
}

func (g *Graph) AddNode(name string) {
	if !g.ContainsNode(name) {
		g.nodes[name] = make(node)
	}
}

func (g *Graph) AddEdge(from string, to string) error {
	f, ok := g.nodes[from]
	if !ok {
		return fmt.Errorf("Node %q not found", from)
	}
	_, ok = g.nodes[to]
	if !ok {
		return fmt.Errorf("Node %q not found", to)
	}

	f.addEdge(to)
	return nil
}

func (g *Graph) ContainsNode(name string) bool {
	_, ok := g.nodes[name]
	return ok
}

// TopSort Topological node classifier.
// The order is defined by the default ordering of the key in the edges map and edges of edges.
// For dependency order, invert the result.
func (g *Graph) TopSort(name string) ([]string, error) {
	results := newOrderedSet()
	err := g.visit(name, results, nil)
	if err != nil {
		return nil, err
	}
	return results.items, nil
}

// TopSortDependency Topological node in dependency order.
// The order is defined by the default ordering of the key in the edges map and edges of edges.
func (g *Graph) TopSortDependency(name string) ([]string, error) {
	results, err := g.TopSort(name)
	if err != nil {
		return nil, err
	}
	Reverse(results)
	return results, nil
}

// TopSortAll Topological nodes classifier of all nodes in the graph, including those that have no edge.
// The order is defined by the default ordering of the key in the node map and its edges.
// For dependency order, invert the result.
func (g *Graph) TopSortAll() ([]string, error) {
	results := newOrderedSet()
	for name := range g.nodes {
		if !results.has(name) {
			err := g.visit(name, results, nil)
			if err != nil {
				return nil, err
			}
		}
	}
	return results.items, nil
}

// TopSortAll Topological nodes classifier of all nodes in the graph, including those that have no edge,
// in dependency order.
// The order is defined by the default ordering of the key in the node map and its edges.
func (g *Graph) TopSortAllDependency() ([]string, error) {
	results, err := g.TopSortAll()
	if err != nil {
		return nil, err
	}
	Reverse(results)
	return results, nil
}

func (g *Graph) visit(name string, results *orderedset, visited *orderedset) error {
	if visited == nil {
		visited = newOrderedSet()
	}

	added := visited.add(name)
	if !added {
		index := visited.index(name)
		cycle := append(visited.items[index:], name)
		return fmt.Errorf("Cycle error: %s", strings.Join(cycle, " -> "))
	}

	n := g.nodes[name]
	for _, edge := range n.edges() {
		err := g.visit(edge, results, visited.copy())
		if err != nil {
			return err
		}
	}

	results.add(name)
	return nil
}

type node map[string]bool

func (n node) addEdge(name string) {
	n[name] = true
}

func (n node) edges() []string {
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	return keys
}

type orderedset struct {
	indexes map[string]int
	items   []string
	length  int
}

func newOrderedSet() *orderedset {
	return &orderedset{
		indexes: make(map[string]int),
		length:  0,
	}
}

func (s *orderedset) add(item string) bool {
	_, ok := s.indexes[item]
	if !ok {
		s.indexes[item] = s.length
		s.items = append(s.items, item)
    s.length++
	}
	return !ok
}

func (s *orderedset) copy() *orderedset {
	clone := newOrderedSet()
	for _, item := range s.items {
		clone.add(item)
	}
	return clone
}

func (s *orderedset) has(item string) bool {
	_, ok := s.indexes[item]
	return ok
}

func (s *orderedset) index(item string) int {
	index, ok := s.indexes[item]
	if ok {
		return index
	}
	return -1
}
