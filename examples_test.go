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

import "fmt"

func ExampleGraph_TopSort() {
	fmt.Println(`
import fmt
// Initialize the graph
graph := topsort.NewGraph()

// Add nodes
graph.AddNode("A", "B", "C", "D", "E")

// Add edges
graph.AddEdge("A", "B")
graph.AddEdge("B", "C")
graph.AddEdge("B", "D")
graph.AddEdge("E", "D")

// Topologically sort only node A in dependency order and your edges, but not sort D and E.
results, err := graph.TopSort("A")
if err != nil {
    panic(err)
}
fmt.Println(results) // => [C D B A]

// Topologically sort all nodes in the graph
results, err := graph.TopSort()
if err != nil {
    panic(err)
}
fmt.Println(results) // => [B C D B A E]
	`)
}

func ExampleGraph_DepthFirst() {
	fmt.Println(`
import fmt
// Initialize the graph
graph := topsort.NewGraph()

// Add nodes
graph.AddNode("A", "B", "C", "D", "E")

// Add edges
graph.AddEdge("A", "B")
graph.AddEdge("B", "C")
graph.AddEdge("B", "D")
graph.AddEdge("E", "D")

// in depth-first order
results, err = graph.DepthFirst("A")
if err != nil {
    panic(err)
}
fmt.Println(results) // => [A B D C]

// all nodes in depth-first order
results, err = graph.DepthFirst()
if err != nil {
    panic(err)
}
fmt.Println(results) // => [E A B D C]
	`)
}

func ExampleGraph_AddEdgeTuple() {
	fmt.Println(`
import fmt
// Initialize the graph
graph := topsort.NewGraph()

// Add nodes
graph.AddNode("A", "B")

// Add edges: A -> B and B -> C
graph.AddEdgeEdgeTuple([2]string{"A", "B"}, [2]string{"B", "C"})
	`)
}

func ExampleGraph_ParseString() {
	fmt.Println(`
import fmt
// Initialize the graph
graph := topsort.NewGraph()

// default separators
graph.ParseString("A>B,B>C", "", "")
// custom separators
graph.ParseString("A-B|B-C", "-", "|")
	`)
}

func ExampleGraph_ParseLines() {
	fmt.Println(`
import fmt
// Initialize the graph
graph := topsort.NewGraph()

lines := []string{"A>B", "B>C,C>D", ""}
i := -1
lineReader := func() string {
	i++
	return lines[i]
}

// default separators
graph.ParseLines("", "", lineReader)

// custom separators
lines := []string{"A-B", "B-C|C-D", ""}
i = -1
graph.ParseLines("-", "|", lineReader)
	`)
}

func ExampleGraph_DOTString() {
	data := `// digraph G {
//  "A" -> "B";
//  "B" -> "C";
//  "B" -> "D";
//  "C"
//  "D"
//  "E" -> "D";
// }`
	fmt.Println(`
import fmt
// Initialize the graph
graph := topsort.NewGraph()
graph := NewGraph()
graph.ParseString("A>B,E,B>C,B>D,E>D", "", "")
println(graph.DOTString())
` + data)
}
