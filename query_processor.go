package sql_gogen_lib

import "context"

// QueryProcessHandler is a function that processes a statement and records.
type QueryProcessHandler[Record any] func(records []Record) error
type QueryProcessor[Record any] interface {
	Process(ctx context.Context, stmt Statement, handler QueryProcessHandler[Record]) error
}
