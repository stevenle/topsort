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
	"testing"
)

func testGraph(graph *Graph, t *testing.T) {
	if len(graph.nodes) != 5 {
		t.Error("Expected 5 nodes.")
	}
	node, ok := graph.nodes["A"]
	if !ok {
		t.Error("Node A not found.")
	}
	if !node.hasEdge("B") {
		t.Error("Node A not have B edge.")
	}

	node, ok = graph.nodes["B"]
	if !node.hasEdge("C") {
		t.Error("Node B not have C edge.")
	}

	node, ok = graph.nodes["C"]
	if !node.hasEdge("D") {
		t.Error("Node C not have D edge.")
	}
}

func TestGraph_ParseStringDefaultSeparators(t *testing.T) {
	graph := NewGraph()
	graph.ParseString("A>B,E,B>C,C>D", "", "")
	testGraph(graph, t)
}

func TestGraph_ParseStringCustomSeparators(t *testing.T) {
	graph := NewGraph()
	graph.ParseString("A-B|E|B-C|C-D", "-", "|")
	testGraph(graph, t)
}

func TestGraph_ParseLinesDefaultSeparators(t *testing.T) {
	graph := NewGraph()
	lines := []string{"A>B,E", "B>C,C>D", ""}
	var i int;
	lineReader := func() string {
		line := lines[i]
		i++
		return line
	}
	graph.ParseLines("", "", lineReader)
	testGraph(graph, t)
}

func TestGraph_ParseLinesCustomSeparators(t *testing.T) {
	graph := NewGraph()
	lines := []string{"A-B|E|C-D", "B-C", ""}
	var i int;
	lineReader := func() string {
		line := lines[i]
		i++
		return line
	}
	graph.ParseLines("-", "|", lineReader)
	testGraph(graph, t)
}