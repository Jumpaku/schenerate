package sqlite3

import (
	"context"
	"fmt"
	"reflect"
)

func QueryRows[RecordStruct any](ctx context.Context, q queryer, stmt string, params []any) (records []RecordStruct, err error) {
	{
		var rs RecordStruct
		rv := reflect.ValueOf(rs)
		if !rv.IsValid() || rv.Kind() != reflect.Struct {
			return nil, fmt.Errorf("RecordStruct must be a struct")
		}
	}
	records, err = query[RecordStruct](ctx, q, stmt, params...)
	if err != nil {
		return nil, fmt.Errorf(`fail to query rows: %w`, err)
	}
	return records, nil
}
