package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph_TopologicalSort(t *testing.T) {
	type testcase[Schema any] struct {
		name               string
		sut                Graph[Schema]
		wantOrderedIndexes []int
		wantCyclic         bool
	}
	tests := []testcase[string]{
		{
			name: "empty",
			sut: NewGraph[string](
				[]string{},
				[][]int{},
			),
			wantOrderedIndexes: []int{},
			wantCyclic:         false,
		},
		{
			name: "self loop",
			sut: NewGraph[string](
				[]string{"A"},
				[][]int{{0}},
			),
			wantOrderedIndexes: nil,
			wantCyclic:         true,
		},
		{
			name: "cyclic",
			sut: NewGraph[string](
				[]string{"A", "B", "C"},
				[][]int{{1}, {2}, {0}},
			),
			wantOrderedIndexes: nil,
			wantCyclic:         true,
		},
		{
			name: "single tree",
			sut: NewGraph[string](
				[]string{"A", "B", "C"},
				[][]int{{}, {0}, {0}},
			),
			wantOrderedIndexes: []int{0, 1, 2},
			wantCyclic:         false,
		},
		{
			name: "diamond",
			sut: NewGraph[string](
				[]string{"A", "B", "C", "D"},
				[][]int{{}, {0}, {0}, {1, 2}},
			),
			wantOrderedIndexes: []int{0, 1, 2, 3},
			wantCyclic:         false,
		},
		{
			name: "complex non-cyclic dag",
			sut: NewGraph[string](
				[]string{"A", "B", "C", "D", "E", "F", "G"},
				[][]int{{}, {0}, {0}, {}, {3}, {3}, {4, 5}},
			),
			wantOrderedIndexes: []int{0, 3, 1, 2, 4, 5, 6},
			wantCyclic:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrderedIndexes, gotCyclic := tt.sut.TopologicalSort()
			assert.Equal(t, gotCyclic, tt.wantCyclic)
			if tt.wantCyclic {
				assert.ElementsMatch(t, gotOrderedIndexes, tt.wantOrderedIndexes)
			}
		})
	}
}
