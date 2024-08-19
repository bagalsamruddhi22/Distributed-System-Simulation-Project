package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Node struct {
	ID     int
	Active bool
}

type Cluster struct {
	Nodes []*Node
	mu    sync.Mutex
}

func (c *Cluster) FaultTolerance() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, node := range c.Nodes {
		if !node.Active {
			fmt.Printf("Node %d failed. Rebalancing...\n", node.ID)
			// Handle rebalancing logic
		}
	}
}

func main() {
	cluster := Cluster{
		Nodes: []*Node{
			{ID: 1, Active: true},
			{ID: 2, Active: true},
			{ID: 3, Active: false},
		},
	}
	cluster.FaultTolerance()
}
