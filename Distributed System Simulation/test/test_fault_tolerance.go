package main

import (
	"testing"
)

func TestCheckNodeHealth(t *testing.T) {
	node := &Node{ID: 1, Active: false}
	if CheckNodeHealth(node) {
		t.Errorf("Expected node %d to be unhealthy", node.ID)
	}
}
