package util

import (
	"reflect"
	"testing"
)

func TestAddNode(t *testing.T) {
	nodes := []string{"A", "B"}
	graph := NewGraph[string]()
	for _, node := range nodes {
		graph.AddNode(node)
	}
	// Ensure nodes exist
	for _, node := range nodes {
		if _, exists := graph.adjacencyList[node]; !exists {
			t.Errorf("Node %s was not added to Graph", node)
		}
	}
	// Ensure no vertices are present
	for _, node := range nodes {
		if len(graph.adjacencyList[node]) != 0 {
			t.Errorf("Node %s has edges, but should be empty", node)
		}
	}

}

func TestAddEdge(t *testing.T) {
	graph := NewGraph[int]()
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	// Check if vertices are present
	if _, exists := graph.adjacencyList[1]; !exists {
		t.Error("Vertex '1' was not added")
	}
	if _, exists := graph.adjacencyList[2]; !exists {
		t.Error("Vertex '2' was not added")
	}
	if len(graph.adjacencyList[1]) != 1 || graph.adjacencyList[1][0] != 2 {
		t.Errorf("Edge from 1 to 2 is incorrect, got: %v", graph.adjacencyList[1])
	}
}

func TestGetNodes(t *testing.T) {
	nodes := []string{"A", "B"}
	graph := NewGraph[string]()
	for _, node := range nodes {
		graph.AddNode(node)
	}
	got := graph.Nodes()
	if !reflect.DeepEqual(got, nodes) {
		t.Errorf("got %v, want %v", got, nodes)
	}
}

func TestGetEdges(t *testing.T) {
	graph := NewGraph[int]()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(3, 4)

	edges, err := graph.Edges(1)
	if err != nil {
		t.Error("Error encountered, but nil expected")
	}
	if !reflect.DeepEqual(edges, []int{2, 3}) {
		t.Errorf("Did not obtain the expected edges")
	}
	// Get edges of an invalid node
	_, err = graph.Edges(99)
	if err == nil {
		t.Errorf("Error expected, but no error obtained")
	}

}

func TestTopologicalSort_ValidDAG(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddNode("E") // Node with no edges

	order, err := graph.TopologicalSort()
	if err != nil {
		t.Errorf("Unexpected error for valid DAG: %v", err)
	}

	// Check that all nodes are in the result
	expectedNodes := []string{"A", "B", "C", "D", "E"}
	nodeSet := make(map[string]bool)
	for _, node := range order {
		nodeSet[node] = true
	}
	for _, node := range expectedNodes {
		if !nodeSet[node] {
			t.Errorf("Missing node %s in topological order", node)
		}
	}

	// Validate order
	nodePosition := make(map[string]int)
	for i, node := range order {
		nodePosition[node] = i
	}
	if nodePosition["A"] > nodePosition["B"] ||
		nodePosition["B"] > nodePosition["C"] ||
		nodePosition["C"] > nodePosition["D"] {
		t.Errorf("Order does not respect dependencies: %v", order)
	}
}
