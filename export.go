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
	"strings"
	"fmt"
	"sort"
)

//DOTString Generate GraphViz DOT string
func (g *Graph) DOTString() string {
	lines := []string{"digraph G {"}
	var nodeNames []string
	for name, _ := range g.nodes {
		nodeNames = append(nodeNames, name)
	}
	sort.Strings(nodeNames)
	for _, name := range nodeNames {
		node := g.nodes[name]
		if len(node) == 0 {
			lines = append(lines, fmt.Sprintf(" %q", name))
		} else {
			var edges []string
			for edge, _ := range node {
				edges = append(edges, edge)
			}
			sort.Strings(edges)
			for _, edge := range edges {
				lines = append(lines, fmt.Sprintf(" %q -> %q;", name, edge))
			}
		}
	}
	return strings.Join(append(lines, "}"), "\n")
}
