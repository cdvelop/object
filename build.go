package object

import (
	"reflect"

	"github.com/cdvelop/model"
)

// arguments: main struct and *model.Module. optional: []*model.Input
func New(model_structs ...interface{}) (*model.Object, error) {

	if len(model_structs) < 2 {
		return nil, model.Error("error tienes que ingresar mínimo una estructura y un puntero de *model.Module como argumentos.")
	}

	var inputs_found []*model.Input

	var main_struct reflect.Type
	var new_struct interface{}
	var module *model.Module

	for _, m := range model_structs {

		t := reflect.TypeOf(m)

		// fmt.Println("KIND: ", t.Kind())

		switch t.Kind() {

		case reflect.Struct:

			if isTheMainStructure(t) {
				main_struct = t
				new_struct = m
			}

		case reflect.Slice:
			// fmt.Println("Slice ")
			sliceValue := reflect.ValueOf(m)
			for i := 0; i < sliceValue.Len(); i++ {
				// fmt.Println("VALOR", sliceValue.Index(i))
				item := sliceValue.Index(i).Interface()
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
			return nil, model.Error("error tipo:", t.Kind().String(), ". no implementado.")
		}
	}

	if main_struct == nil {
		return nil, model.Error("error estructura principal no ingresada (verifica si los campos tiene el tag 'Legend')")
	}

	new_fields, TextFieldNames, err := buildFieldsObject(main_struct, inputs_found...)
	if err != nil {
		if main_struct.NumField() != 0 {
			return nil, model.Error(err.Error())
		}
	}

	new_object := model.Object{
		Name:            main_struct.Name(),
		TextFieldNames:  TextFieldNames,
		Fields:          new_fields,
		Module:          module,
		BackendHandler:  model.BackendHandler{},
		FrontendHandler: model.FrontendHandler{},
	}

	addFrontHandlers(&new_object, new_struct)

	addBackHandlers(&new_object, new_struct)

	return &new_object, nil
}
