// Package pov solves an Exercism challenge.
package pov

import (
	"fmt"
)

// Required for testing.
const testVersion = 2

// A Graph is a hashmap from nodes to the children of those nodes.
type Graph struct {
	m    map[string][]string
	arcs int
}

// New returns a pointer to a new (empty) Graph.
func New() *Graph {
	return &Graph{make(map[string][]string), 0}
}

// AddNode adds a node to the Graph g. Adding a node with a pre-existing
// label does not change the graph.
func (g *Graph) AddNode(label string) {
	if _, ok := g.m[label]; ok {
		return
	}
	g.m[label] = make([]string, 0)
}

// AddArc adds a directed edge from f to t. If t does not already exist in
// the graph, AddArc does nothing. If f does not already exist in the graph,
// a new node is created.
func (g *Graph) AddArc(f, t string) {
	if _, ok := g.m[t]; !ok {
		return
	}
	g.arcs++
	if _, ok := g.m[f]; !ok {
		g.m[f] = make([]string, 1)
		g.m[f][0] = t
		return
	}
	g.m[f] = append(g.m[f], t)
}

// ArcList returns a list of all arcs in the Graph, in the form "from -> to".
func (g *Graph) ArcList() []string {
	ret := make([]string, g.arcs)
	i := 0
	for from, edges := range g.m {
		for _, to := range edges {
			ret[i] = fmt.Sprintf("%s -> %s", from, to)
			i++
		}
	}
	return ret
}

// ChangeRoot changes the root of the Graph.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	// Copy  g.
	newG := Graph{make(map[string][]string), g.arcs}
	for k, v := range g.m {
		newG.m[k] = make([]string, len(v))
		for i, label := range v {
			newG.m[k][i] = label
		}
	}
	reversePath(newG, oldRoot, newRoot)
	return &newG
}

// reversePath reverses the path from oldRoot to newRoot.
func reversePath(g Graph, startSearch, newRoot string) bool {
	if len(g.m[startSearch]) == 0 {
		return false
	}
	for i, c := range g.m[startSearch] {
		if c == newRoot || reversePath(g, c, newRoot) {
			g.m[startSearch] = append(g.m[startSearch][:i], g.m[startSearch][i+1:]...)
			g.m[c] = append(g.m[c], startSearch)
			return true
		}
	}
	return false
}
