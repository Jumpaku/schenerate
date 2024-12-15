package sqlgogen

import "context"

// TableLister is an interface that lists tables.
type TableLister interface {
	List(ctx context.Context) ([]Table, error)
}
