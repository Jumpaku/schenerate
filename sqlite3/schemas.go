package sqlite3

type Schemas []Schema

type Schema struct {
	Name        string
	Type        string
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
	Key       []string
	Reference ForeignKeyReference
}

type ForeignKeyReference struct {
	Table string
	Key   []string
}

type Index struct {
	Name   string
	Origin IndexOrigin
	Unique bool
	Key    []IndexKeyElem
}

// IndexOrigin is a type for the origin of an index.
// https://www.sqlite.org/pragma.html#pragma_index_list
// > "c" if the index was created by a CREATE INDEX statement, "u" if the index was created by a UNIQUE constraint, or "pk" if the index was created by a PRIMARY KEY constraint.
type IndexOrigin string

const (
	IndexOriginPrimaryKey       IndexOrigin = "pk"
	IndexOriginCreateIndex      IndexOrigin = "c"
	IndexOriginUniqueConstraint IndexOrigin = "u"
)

type IndexKeyElem struct {
	Name string
	Desc bool
}
