package name

import "github.com/samber/lo"

// Registry provides safe names which are not included in the reserved keywords.
type Registry interface {
	Safe(name string) string
}

// NewRegistry creates a new name registry.
// The parameter reserved is a list of reserved keywords.
// The parameter modify is a function to modify the name when the name is already used, which takes the specified name and the number of the modifications and returns the modified name.
func NewRegistry(reserved []string, modify func(name string, times int) string) Registry {
	return registry{
		used:      lo.SliceToMap(reserved, func(s string) (string, bool) { return s, true }),
		transform: map[string]string{},
		modify:    modify,
	}
}

type registry struct {
	used      map[string]bool
	transform map[string]string
	modify    func(name string, times int) string
}

func (r registry) Safe(name string) string {
	if _, ok := r.transform[name]; ok {
		return r.transform[name]
	}

	modName := name
	for times := 0; ; {
		if !r.used[modName] {
			r.used[modName] = true
			r.transform[modName] = modName
			return modName
		}
		times++
		newName := r.modify(name, times)
		if newName == modName {
			panic("modify function is not working properly")
		}
		modName = newName
	}
}
