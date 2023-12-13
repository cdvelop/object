package object

import (
	"github.com/cdvelop/model"
)

var object_fields []string
var module_fields []string
var handlers_fields []string

func knownName(name string) bool {
	if handlers != nil && len(handlers_fields) == 0 {
		names := model.GetFieldNamesFrom(model.Handlers{})
		names = append(names, "Handlers")
		handlers_fields = names
	}

	for _, n := range handlers_fields {
		if name == n {
			return true
		}
	}

	if len(object_fields) == 0 {
		object_fields = model.GetFieldNamesFrom(model.Object{})
		object_fields = append(object_fields, "Object")
		object_fields = append(object_fields, "O")
		// fmt.Println("object_fields", object_fields)
	}

	// fmt.Println("CAMPO CONOCIDO?:", name)

	for _, n := range object_fields {
		if name == n {
			return true
		}
	}
	if len(module_fields) == 0 {
		module_fields = model.GetFieldNamesFrom(model.Module{})
		module_fields = append(module_fields, "Module")
		// fmt.Println("module_fields", module_fields)
	}

	for _, n := range module_fields {
		if name == n {
			return true
		}
	}

	return false
}
