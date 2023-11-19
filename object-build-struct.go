package object

import "github.com/cdvelop/model"

// *model.Module optional
func BuildObjectFromStruct(model_struct ...interface{}) (*model.Object, error) {

	st_found, module, err := getStructFromInterface("BuildObjectFromStruct", model_struct...)
	if err != nil {
		return nil, err
	}

	if len(st_found) != 1 {
		return nil, model.Error("solo puedes ingresar una estructura para crear el objeto")
	}

	new_object, err := st_found[0].buildObject(module)
	if err != nil {
		return nil, err
	}

	return new_object, nil
}
