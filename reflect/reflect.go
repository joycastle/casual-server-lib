package reflect

import (
	"reflect"
)

var intTypes = map[reflect.Kind]struct{}{
	reflect.Int:    struct{}{},
	reflect.Int8:   struct{}{},
	reflect.Int16:  struct{}{},
	reflect.Int32:  struct{}{},
	reflect.Int64:  struct{}{},
	reflect.Uint:   struct{}{},
	reflect.Uint8:  struct{}{},
	reflect.Uint16: struct{}{},
	reflect.Uint32: struct{}{},
	reflect.Uint64: struct{}{},
}

func IsIntType(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	_, ok := intTypes[kind]
	return ok
}
