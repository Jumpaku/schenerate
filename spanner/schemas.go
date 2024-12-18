package spanner

type Schemas []Schema

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
