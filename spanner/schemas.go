package spanner

import (
	"github.com/Jumpaku/schenerate/graph"
	"github.com/samber/lo"
)

type Schemas []Schema

func (s Schemas) BuildGraph() graph.Graph[Schema] {
	schemaMap := make(map[string]int)
	for i, schema := range s {
		schemaMap[schema.Name] = i
	}

	dep := make([][]int, len(s))
	for u, schema := range s {
		var d []int
		if schema.Parent != "" {
			d = append(d, schemaMap[schema.Parent])
		}
		d = append(d, lo.Map(schema.ForeignKeys, func(fk ForeignKey, _ int) int {
			return schemaMap[fk.Reference.Table]
		})...)
		dep[u] = lo.Uniq(d)
	}

	return graph.NewGraph[Schema](s, dep)
}

type Schema struct {
	Name        string
	Type        string
	Parent      string
	Columns     []Column
	PrimaryKey  []string
	ForeignKeys []ForeignKey
	Indexes     []Index
}

type Column struct {
	Name     string
	Type     string
	Nullable bool
}

type ForeignKey struct {
	Name      string
	Key       []string
	Reference ForeignKeyReference
}

type ForeignKeyReference struct {
	Table string
	Key   []string
}

type Index struct {
	Name   string
	Unique bool
	Key    []IndexKeyElem
}

type IndexKeyElem struct {
	Name string
	Desc bool
}
