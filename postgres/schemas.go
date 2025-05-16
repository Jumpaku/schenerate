package postgres

import (
	"github.com/Jumpaku/schenerate/graph"
	"github.com/samber/lo"
)

type Schemas []Schema

func (s Schemas) BuildGraph() graph.Graph[Schema] {
	type key struct {
		schema string
		table  string
	}
	schemaMap := make(map[key]int)
	for i, schema := range s {
		schemaMap[key{schema: schema.Schema, table: schema.Name}] = i
	}

	dep := make([][]int, len(s))
	for u, schema := range s {
		d := lo.Map(schema.ForeignKeys, func(fk ForeignKey, _ int) int {
			return schemaMap[key{schema: fk.Reference.Schema, table: fk.Reference.Table}]
		})
		dep[u] = lo.Uniq(d)
	}

	return graph.NewGraph[Schema](s, dep)
}

type Schema struct {
	Schema      string
	Name        string
	Type        string
	Columns     []Column
	PrimaryKey  []string
	ForeignKeys []ForeignKey
	UniqueKeys  []UniqueKey
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
	Schema string
	Table  string
	Key    []string
}

type UniqueKey struct {
	Name string
	Key  []string
}
type Index struct {
	Name   string
	Unique bool
	Key    []string
}

type IndexKeyElem struct {
	Name string
	Desc bool
}
