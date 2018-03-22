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

// TopSort Topological node classifier.
// The order is defined by the default ordering of the key in the edges map and edges of edges.
// If `names` is empty, sort all nodes in the graph.
func (g *Graph) TopSort(names ...string) ([]string, error) {
	var err error
	results := newOrderedSet()

	if len(names) == 0 {
		for name, _ := range g.nodes {
			names = append(names, name)
		}
	}

	for _, name := range names {
		if !results.has(name) {
			err = g.visit(name, results, nil)
			if err != nil {
				return nil, err
			}
		}
	}
	return results.items, nil
}

// DepthFirst Depth-first classifier.
// The order is defined by the default ordering of the key in the edges map and edges of edges.
// If `names` is empty, classify all nodes in the graph.
func (g *Graph) DepthFirst(names ...string) ([]string, error) {
	results, err := g.TopSort(names...)
	if err != nil {
		return nil, err
	}
	Reverse(results)
	return results, nil
}
