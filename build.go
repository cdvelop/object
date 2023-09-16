package object

import (
	"reflect"

	"github.com/cdvelop/model"
)

type structFound struct {
	struct_int interface{}
	struct_ref reflect.Type
}

// arguments: main struct and *model.Module. optional: []*model.Input
func New(model_structs ...interface{}) error {

	if len(model_structs) < 2 {
		return model.Error("error tienes que ingresar mínimo una estructura y un puntero de *model.Module como argumentos.")
	}

	var inputs_found []*model.Input

	var structs_found []structFound

	var module *model.Module

	for _, m := range model_structs {

		t := reflect.TypeOf(m)

		// fmt.Println("KIND: ", t.Kind())

		switch t.Kind() {

		case reflect.Struct:

			structs_found = append(structs_found, structFound{
				struct_int: m,
				struct_ref: t,
			})

		case reflect.Slice:
			// fmt.Println("Slice ")
			sliceValue := reflect.ValueOf(m)
			for i := 0; i < sliceValue.Len(); i++ {
				// fmt.Println("VALOR", sliceValue.Index(i))

				item := sliceValue.Index(i).Interface()

				// Verifica si el elemento es una estructura
				// if reflect.TypeOf(item).Kind() == reflect.Struct {
				// 	// El elemento es una estructura
				// 	structs_found = append(structs_found, structFound{
				// 		struct_int: item,
				// 		struct_ref: reflect.TypeOf(item),
				// 	})
				// }

				// Verifica si el elemento es de tipo *model.Input
				if input_item, ok := item.(*model.Input); ok {
					inputs_found = append(inputs_found, input_item)

				}
			}

		case reflect.Ptr:
			// fmt.Println("PUNTERO:", t.Name())

			module_value := reflect.ValueOf(m).Interface()

			if module_pointer, ok := module_value.(*model.Module); ok {
				// fmt.Println("ESTRUCTURA ES UN PUNTERO MODULO: ", module_pointer)
				module = module_pointer
			}

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

	for _, sf := range structs_found {

		new_fields, TextFieldNames, err := buildFieldsObject(sf.struct_ref, inputs_found...)
		if err != nil {
			if sf.struct_ref.NumField() != 0 {
				return model.Error(err.Error())
			}
		}

		new_object := &model.Object{
			Name:            sf.struct_ref.Name(),
			TextFieldNames:  TextFieldNames,
			Fields:          new_fields,
			Module:          module,
			BackendHandler:  model.BackendHandler{},
			FrontendHandler: model.FrontendHandler{},
		}

		addFrontHandlers(new_object, sf.struct_int)

		addBackHandlers(new_object, sf.struct_int)

		module.Objects = append(module.Objects, new_object)

	}

	return nil
}
