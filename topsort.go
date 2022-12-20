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

type Graph[Key comparable] struct {
	nodes map[Key]nodeimpl[Key]
}

func NewGraph[Key comparable]() *Graph[Key] {
	return &Graph[Key]{
		nodes: make(map[Key]nodeimpl[Key]),
	}
}

func (g *Graph[Key]) AddNode(key Key) {
	if !g.ContainsNode(key) {
		g.nodes[key] = make(nodeimpl[Key])
	}
}

func (g *Graph[Key]) getOrAddNode(node Key) nodeimpl[Key] {
	n, ok := g.nodes[node]
	if !ok {
		n = make(nodeimpl[Key])
		g.nodes[node] = n
	}
	return n
}

func (g *Graph[Key]) AddEdge(from Key, to Key) error {
	f := g.getOrAddNode(from)
	g.AddNode(to)
	f.addEdge(to)
	return nil
}

func (g *Graph[Key]) ContainsNode(key Key) bool {
	_, ok := g.nodes[key]
	return ok
}

func (g *Graph[Key]) TopSort(key Key) ([]Key, error) {
	results := newOrderedSet[Key]()
	err := g.visit(key, results, nil)
	if err != nil {
		return nil, err
	}
	return results.items, nil
}

func (g *Graph[Key]) visit(key Key, results *orderedset[Key], visited *orderedset[Key]) error {
	if visited == nil {
		visited = newOrderedSet[Key]()
	}

	added := visited.add(key)
	if !added {
		index := visited.index(key)
		cycle := append(visited.items[index:], key)
		strs := make([]string, len(cycle))
		for i, k := range cycle {
			strs[i] = fmt.Sprintf("%v", k)
		}
		return fmt.Errorf("Cycle error: %s", strings.Join(strs, " -> "))
	}

	n := g.nodes[key]
	for _, edge := range n.edges() {
		err := g.visit(edge, results, visited.copy())
		if err != nil {
			return err
		}
	}

	results.add(key)
	return nil
}

type nodeimpl[Key comparable] map[Key]bool

func (n nodeimpl[Key]) addEdge(key Key) {
	n[key] = true
}

func (n nodeimpl[Key]) edges() []Key {
	var keys []Key
	for k := range n {
		keys = append(keys, k)
	}
	return keys
}

type orderedset[Key comparable] struct {
	indexes map[Key]int
	items   []Key
	length  int
}

func newOrderedSet[Key comparable]() *orderedset[Key] {
	return &orderedset[Key]{
		indexes: make(map[Key]int),
		length:  0,
	}
}

func (s *orderedset[Key]) add(item Key) bool {
	_, ok := s.indexes[item]
	if !ok {
		s.indexes[item] = s.length
		s.items = append(s.items, item)
		s.length++
	}
	return !ok
}

func (s *orderedset[Key]) copy() *orderedset[Key] {
	clone := newOrderedSet[Key]()
	for _, item := range s.items {
		clone.add(item)
	}
	return clone
}

func (s *orderedset[Key]) index(item Key) int {
	index, ok := s.indexes[item]
	if ok {
		return index
	}
	return -1
}
