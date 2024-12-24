package postgres

type Schemas []Schema

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
