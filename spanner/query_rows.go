package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"reflect"
)

func QueryRows[RecordStruct any](ctx context.Context, q queryer, stmt spanner.Statement) (records []RecordStruct, err error) {
	{
		var rs RecordStruct
		rv := reflect.ValueOf(rs)
		if !rv.IsValid() || rv.Kind() != reflect.Struct {
			return nil, fmt.Errorf("RecordStruct must be a struct")
		}
	}
	tx := q.client.ReadOnlyTransaction()
	defer tx.Close()

	records, err = query[RecordStruct](ctx, tx, stmt)
	if err != nil {
		return nil, fmt.Errorf(`fail to query rows: %w`, err)
	}
	return records, nil
}
