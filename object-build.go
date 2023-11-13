package object

import "github.com/cdvelop/model"

// *model.Module optional
func BuildObjectFromStruct(no_add_object_to_module bool, model_struct ...interface{}) (*model.Object, error) {

	st_found, module, err := getStructFromInterface("BuildObjectFromStruct", model_struct...)
	if err != nil {
		return nil, err
	}

	if len(st_found) != 1 {
		return nil, model.Error("solo puedes ingresar una estructura para rear el objeto")
	}

	new_object, err := buildObject(st_found[0], module, no_add_object_to_module)
	if err != nil {
		return nil, err
	}

	return new_object, nil
}
