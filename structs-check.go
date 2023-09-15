package object

import (
	"reflect"
)

func isTheMainStructure(t reflect.Type) bool {

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("Legend") != "" {
			return true
		}
	}
	return false
}
