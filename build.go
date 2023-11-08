package object

import (
	"reflect"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

type structFound struct {
	struct_int interface{}
	struct_ref reflect.Type
}

// arguments: main struct and *model.Module. inputs: []*model.Input, *model.Handlers
func New(model_structs ...interface{}) error {

	if len(model_structs) < 2 {
		return model.Error("error tienes que ingresar mínimo una estructura y un puntero de *model.Module como argumentos.")
	}

	var structs_found []structFound

	var module *model.Module
	var handlers *model.Handlers

	for _, m := range model_structs {

		t := reflect.TypeOf(m)

		// fmt.Println("KIND: ", t.Kind())

		switch t.Kind() {

		case reflect.Struct:

			return model.Error("error debes de ingresar las estructuras como  punteros.")

		case reflect.Slice:
			// fmt.Println("Slice ")
			// sliceValue := reflect.ValueOf(m)
			// for i := 0; i < sliceValue.Len(); i++ {
			// fmt.Println("VALOR", sliceValue.Index(i))

			// item := sliceValue.Index(i).Interface()

			// Verifica si el elemento es una estructura
			// if reflect.TypeOf(item).Kind() == reflect.Struct {
			// 	// El elemento es una estructura
			// 	structs_found = append(structs_found, structFound{
			// 		struct_int: item,
			// 		struct_ref: reflect.TypeOf(item),
			// 	})
			// }

			// Verifica si el elemento es de tipo *model.Input
			// if input_item, ok := item.(*model.Input); ok {
			// 	inputs_found = append(inputs_found, input_item)

			// }
			// }

		case reflect.Ptr:

			module_value := reflect.ValueOf(m).Interface()

			if module_pointer, ok := module_value.(*model.Module); ok {
				// fmt.Println("ESTRUCTURA ES UN PUNTERO MODULO: ", module_pointer)
				module = module_pointer
			} else if handlers_found, ok := module_value.(*model.Handlers); ok {
				handlers = handlers_found

			} else {

				// puede que se enviaron las estructuras principales como punteros

				// Obtén el tipo subyacente al puntero
				elem_type := t.Elem()

				// Verifica si el tipo subyacente es una estructura
				if elem_type.Kind() == reflect.Struct {
					// fmt.Println("El puntero es de tipo estructura válida:", t.Name())
					structs_found = append(structs_found, structFound{
						struct_int: m,
						struct_ref: elem_type,
					})

				}
			}

		case reflect.Interface:

		default:
			return model.Error("error tipo:", t.Kind().String(), ". no implementado.")
		}
	}

	if len(structs_found) == 0 {
		return model.Error("error ninguna estructura valida ingresada")
	}

	if module == nil {
		return model.Error("error puntero de *model.Module no ingresado como argumento")
	}
	if handlers == nil {
		return model.Error("error puntero de *model.Handlers en modulo:", module.ModuleName, "no ingresado como argumento")
	}

	for _, sf := range structs_found {

		obj_name := strings.ToLowerCaseAlphabet(sf.struct_ref.Name())

		new_object := model.Object{
			Name:            module.ModuleName + "." + obj_name,
			Table:           obj_name,
			Module:          module,
			BackendHandler:  model.BackendHandler{},
			FrontendHandler: model.FrontendHandler{},
		}

		err := sf.setStructField(&new_object, handlers)
		if err != nil {
			return model.Error(err.Error())
		}

		addFrontHandlers(&new_object, sf.struct_int)

		addBackHandlers(&new_object, sf.struct_int)

		addBasicHandlers(&new_object, sf.struct_int)

		module.Objects = append(module.Objects, &new_object)

	}

	return nil
}
