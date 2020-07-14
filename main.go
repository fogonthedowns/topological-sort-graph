package main

import (
	"container/list"
	"fmt"
)

// Topological Sort

// Diirected Graph

// Set - all the visited vertex
// Stack - all the vertexes in topolotical order

// a -> v
// a -> c
// b -> d
// b -> c
// d -> z
// z -> f
// c -> e
// c -> q
// q -> r
// e -> f

type node map[string]bool
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

func (g *Graph) ContainsNode(name string) bool {
	_, ok := g.nodes[name]
	return ok
}

func (g *Graph) AddEdge(from string, to string) error {
	f, ok := g.nodes[from]
	if !ok {
		return fmt.Errorf("From Node %q not found. call AddNode(name string)", from)
	}

	_, ok = g.nodes[to]
	if !ok {
		return fmt.Errorf("to Node %q not found. call AddNode(name string)", to)
	}

	// uses from Node
	// maps from Node, to to Node
	f[to] = true
	return nil
}

func main() {
	g := initGraph()
	// track visited verticces
	visited := make(map[string]bool)
	// sorted stack
	stack := list.New()
	// range over all vertices
	for vertex, _ := range g.nodes {
		if visited[vertex] {
			continue
		}
		g.topSort(vertex, stack, visited)
	}

	// iterate over the stack
	for e := stack.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
}

func (g *Graph) topSort(vertex string, stack *list.List, visited map[string]bool) {
	// add to visited
	visited[vertex] = true
	// visit all the children first
	for child, _ := range g.nodes[vertex] {
		// continue to loop through the children if the child has been visited already
		if visited[child] == true {
			continue
		}
		// recursively call the function
		g.topSort(child, stack, visited)
		//stack.PushFront(child)
	}
	// after the children are visited push into the sorted stack
	stack.PushFront(vertex)
}

func initGraph() *Graph {
	g := NewGraph()
	n := []string{"a", "v", "b", "f", "c", "d", "q", "e", "z", "r", "f"}
	for _, val := range n {
		g.AddNode(val)
	}
	g.AddEdge("a", "v")
	g.AddEdge("a", "c")
	g.AddEdge("b", "c")
	g.AddEdge("b", "d")
	g.AddEdge("d", "z")
	g.AddEdge("q", "r")
	g.AddEdge("c", "e")
	g.AddEdge("c", "q")
	g.AddEdge("e", "f")
	g.AddEdge("z", "f")
	return g
}
