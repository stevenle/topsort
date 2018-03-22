topsort
=======

Topological Sorting for Golang

Topological sorting algorithms are especially useful for dependency calculation, and so this particular implementation is mainly intended for this purpose. As a result, the direction of edges and the order of the results may seem reversed compared to other implementations of topological sorting.

For example, if:

* A depends on B
* B depends on C
* B depends on D
* E depends on D

The graph is represented as:

![Graph image](https://www.planttext.com/plantuml/img/SoWkIImgAStDuKh9J2zABCXGS5Uevb800aS5NJi59p3J2SHqHZ1Tm4nN2BDMWSiXDIy5Q0G0)

The code for this example would look something like:

```go
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

// in depth-first order
results, err = graph.DepthFirst("A")
if err != nil {
    panic(err)
}
fmt.Println(results) // => [A B D C]
```

Sort all nodes:

```go
// Topologically sort all nodes in the graph
results, err := graph.TopSort()
if err != nil {
    panic(err)
}
fmt.Println(results) // => [B C D B A E]

// all nodes in depth-first order
results, err = graph.DepthFirst()
if err != nil {
    panic(err)
}
fmt.Println(results) // => [E A B D C]
```
See `examples_test.go` for more examples.