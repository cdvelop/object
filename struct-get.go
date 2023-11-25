package object

import (
	"reflect"

	"github.com/cdvelop/model"
)

var handlers *model.Handlers

func getStructFromInterface(calling_function_name string, model_structs ...interface{}) (structs_found []*structFound, m *model.Module, err string) {

	for _, st := range model_structs {

		t := reflect.TypeOf(st)

		switch t.Kind() {

		case reflect.Struct:

			return nil, nil, "error debes de ingresar las estructuras como punteros en " + calling_function_name

		case reflect.Ptr:

			module_value := reflect.ValueOf(st).Interface()

			if module_pointer, ok := module_value.(*model.Module); ok {
				// fmt.Println("ESTRUCTURA ES UN PUNTERO MODULO: ", module_pointer)
				m = module_pointer
			} else if handlers_found, ok := module_value.(*model.Handlers); ok {
				// fmt.Println("ESTRUCTURA ES handlers_found: ", module_pointer)
				handlers = handlers_found

			} else {

				// puede que se enviaron las estructuras principales como punteros
				// Obtén el tipo subyacente al puntero
				elem_type := t.Elem()

				// Verifica si el tipo subyacente es una estructura
				if elem_type.Kind() == reflect.Struct {
					// fmt.Println("El puntero es de tipo estructura válida:", t.Name())

					structs_found = append(structs_found, &structFound{
						struct_int: st,
						struct_ref: elem_type,
					})
				}
			}
		}
	}

	if len(structs_found) == 0 {
		return nil, nil, "error ninguna estructura valida ingresada en " + calling_function_name
	}

	return
}
