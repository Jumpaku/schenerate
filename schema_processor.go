package sqlgogen

import "context"

// SchemaProcessHandler is a function that processes table schemas.
type SchemaProcessHandler[Schemas any] func(schemas Schemas) error

// SchemaProcessor is an interface that processes schemas.
type SchemaProcessor[Schemas any] interface {
	Process(ctx context.Context, tables []Table, handler SchemaProcessHandler[Schemas]) error
}
