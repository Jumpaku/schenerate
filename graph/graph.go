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
// This method returns nil, true if the cycle is detected.
// Otherwise, it returns ordered indexes the schemas so that u comes after v for all dependencies (u,v) each of which is the relationship where the schema at index u depends on the schema at index v.
// Related description: https://en.wikipedia.org/wiki/Topological_sorting#Kahn's_algorithm
func (g Graph[Schema]) TopologicalSort() (orderedIndexes []int, cyclic bool) {
	inDegree := make([]int, g.Len())
	for _, vs := range g.dependency {
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

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		orderedIndexes = append(orderedIndexes, u)
		for _, v := range g.dependency[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(orderedIndexes) != g.Len() {
		return nil, true // Cycle detected
	}

	return orderedIndexes, false
}
