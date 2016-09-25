// Package tree solvse an Exercism challenge.
package tree

import (
	"errors"
)

// Required for testing.
const testVersion = 3

// A Record is pretty self explanatory.
type Record struct {
	ID, Parent int
}

// A Node is part of the tree we're building from the Records.
type Node struct {
	ID       int
	Children []*Node
}

// Build takes a slice of Records and turns that slice into a tree. The root
// record always has ID 0.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// Make an array of Node's that I'll connect later.
	tree := make([]*Node, len(records))
	for _, r := range records {
		if r.ID < 0 || r.ID >= len(tree) {
			return nil, errors.New("invalid ID number")
		}
		if tree[r.ID] != nil {
			return nil, errors.New("duplicate ID number")
		}
		tree[r.ID] = new(Node)
		tree[r.ID].ID = r.ID
	}
	// Sort the records array by r.ID.
	for i := 0; i < len(records); {
		current := records[i]
		if current.ID == i {
			i++
			continue
		}
		// Swap records[i] with records[current.ID].
		records[i], records[current.ID] = records[current.ID], records[i]
	}
	// Run various checks on the records array.
	if records[0].Parent != 0 {
		return nil, errors.New("root node has parent")
	}
	for i := 1; i < len(records); i++ {
		if records[i].Parent > i {
			return nil, errors.New("record ID is lower than parent ID")
		}
	}
	// Add all the necessary links in the tree array.
	for i := 1; i < len(records); i++ {
		r := records[i]
		n := tree[i]
		tree[r.Parent].Children = append(tree[r.Parent].Children, n)
	}
	// Make sure the tree is connected.
	if count := countDecendents(tree[0]); count != len(tree) {
		return nil, errors.New("tree is not connected")
	}
	return tree[0], nil
}

// countDecendents counts the number of decendents of a *Node.
func countDecendents(n *Node) int {
	total := 1
	for _, c := range n.Children {
		total += countDecendents(c)
	}
	return total
}
