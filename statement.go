package sqlgogen

import (
	"database/sql"
	"fmt"
)

type Statement struct {
	Stmt string
	Args []any
}

func (s Statement) ArgsMap() map[string]any {
	args := map[string]any{}
	for _, arg := range s.Args {
		switch arg := arg.(type) {
		case sql.NamedArg:
			args[arg.Name] = arg.Value
		}
	}
	x := 0
	for _, arg := range s.Args {
		switch arg := arg.(type) {
		case sql.NamedArg:
		default:
			for {
				name := fmt.Sprintf("Arg%d", x)
				if _, ok := args[name]; !ok {
					args[name] = arg
					break
				}
				x++
			}
		}
	}
	return args

}
