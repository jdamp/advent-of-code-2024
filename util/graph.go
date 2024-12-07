package util

import "errors"

// Graph represents a generic directed graph
type Graph[T comparable] struct {
	adjacencyList map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	adjacencyList := make(map[T][]T)
	return &Graph[T]{
		adjacencyList,
	}
}

func (g *Graph[T]) AddNode(node T) {
	if _, exists := g.adjacencyList[node]; !exists {
		g.adjacencyList[node] = []T{}
	}
}

// AddEdge adds a directed edge from source to target and creates the Nodes if they do not exist
func (g *Graph[T]) AddEdge(source, target T) {
	g.AddNode(source)
	g.AddNode(target)
	g.adjacencyList[source] = append(g.adjacencyList[source], target)
}

// Nodes returns the nodes for a given graph
func (g *Graph[T]) Nodes() []T {
	keys := make([]T, 0, len(g.adjacencyList))
	for k := range g.adjacencyList {
		keys = append(keys, k)
	}
	return keys
}

// Edges returns the edges for a given node
func (g *Graph[T]) Edges(node T) ([]T, error) {
	nodes, exists := g.adjacencyList[node]
	if !exists {
		return nil, errors.New("node does not exist in Graph")
	}
	return nodes, nil
}

// Topological sort performs a topological sort of the Graph using Kahn's algorithm
func (g *Graph[T]) TopologicalSort() ([]T, error) {
	// Step 1: Compute in-degrees of all nodes
	inDegree := make(map[T]int)
	for _, node := range g.Nodes() {
		inDegree[node] = 0
	}
	for _, node := range g.Nodes() {
		edges, _ := g.Edges(node)
		for _, neighbor := range edges {
			inDegree[neighbor]++
		}
	}
	// Step 2: Initialize a queue with all nodes having in-degree 0
	queue := []T{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}
	topoOrder := []T{}
	// Step 3: Process the graph
	var node T
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]
		topoOrder = append(topoOrder, node)
		neighbors, _ := g.Edges(node)
		for _, neighbor := range neighbors {
			inDegree[neighbor] -= 1
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(topoOrder) != len(g.Nodes()) {
		return nil, errors.New("the graph contains a cycle; topological sorting is not possible")
	}

	return topoOrder, nil
}
