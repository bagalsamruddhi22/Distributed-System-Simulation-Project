package main

import (
	"fmt"
)

type Node struct {
	ID     int
	Active bool
}

func CheckNodeHealth(node *Node) bool {
	// Simulate health check
	return node.Active
}

func HandleNodeFailure(node *Node) {
	fmt.Printf("Handling failure for Node %d...\n", node.ID)
	// Implement recovery or task reassignment logic here
}

func main() {
	node := &Node{ID: 1, Active: false}
	if !CheckNodeHealth(node) {
		HandleNodeFailure(node)
	} else {
		fmt.Printf("Node %d is healthy.\n", node.ID)
	}
}
