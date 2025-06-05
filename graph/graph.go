package graph

type Graph[Schema any] struct {
	schemas    []Schema
	dependency [][]int
}

func NewGraph[Schema any](schemas []Schema, dependency [][]int) Graph[Schema] {
	return Graph[Schema]{
		schemas:    schemas,
		dependency: dependency,
	}
}

// Get returns the schema at the given index.
func (g Graph[Schema]) Get(index int) Schema {
	return g.schemas[index]
}

// Len returns the number of schemas in the graph.
func (g Graph[Schema]) Len() int {
	return len(g.schemas)
}

// References returns the indexes of the schemas that the schema at index depends on.
func (g Graph[Schema]) References(index int) []int {
	return g.dependency[index]
}

// TopologicalSort performs a topological sort on the graph.
// Returns (nil, true) if a cycle is detected in the graph.
// Otherwise, returns (orderedIndexes, false), where orderedIndexes is a slice of schema indexes sorted such that
// the index u comes after the index v for all dependencies (u, v), each of which is the relationship that the schema
// at the index v depends on the schema at the index u.
// Related description: https://en.wikipedia.org/wiki/Topological_sorting#Kahn's_algorithm
func (g Graph[Schema]) TopologicalSort() (orderedIndexes []int, cyclic bool) {
	dep := make([][]int, g.Len())
	for u, vs := range g.dependency {
		for _, v := range vs {
			dep[v] = append(dep[v], u)
		}
	}
	inDegree := make([]int, len(dep))
	for _, vs := range dep {
		for _, v := range vs {
			inDegree[v]++
		}
	}

	var q []int
	for i, n := range inDegree {
		if n == 0 {
			q = append(q, i)
		}
	}

	orderedIndexes = []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		orderedIndexes = append(orderedIndexes, u)
		for _, v := range dep[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(orderedIndexes) != len(dep) {
		return nil, true // Cycle detected
	}

	return orderedIndexes, false
}
