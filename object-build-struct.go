package object

import "github.com/cdvelop/model"

// *model.Module optional
func BuildObjectFromStruct(model_struct ...interface{}) (new_object *model.Object, err string) {

	st_found, module, err := getStructFromInterface("BuildObjectFromStruct", model_struct...)
	if err != "" {
		return nil, err
	}

	if len(st_found) != 1 {
		return nil, "solo puedes ingresar una estructura para crear el objeto"
	}

	new_object, err = st_found[0].buildObject(module)

	return
}
