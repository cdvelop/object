package object

import (
	"reflect"

	"github.com/cdvelop/model"
)

func AreIdentical(obj1, obj2 *model.Object) bool {

	// fmt.Println("OBJETO 1", obj1.Name)
	// fmt.Println("OBJETO 2", obj2.Name)

	if obj1.ObjectName != obj2.ObjectName || len(obj1.Fields) != len(obj2.Fields) {
		return false
	}

	for i := range obj1.Fields {
		if !reflect.DeepEqual(obj1.Fields[i], obj2.Fields[i]) {
			return false
		}
	}

	if len(obj1.PrincipalFieldsName) != len(obj2.PrincipalFieldsName) {
		return false
	}

	if obj1.FrontHandler != obj2.FrontHandler {
		return false
	}

	if obj1.BackHandler != obj2.BackHandler {
		return false
	}

	return reflect.DeepEqual(obj1, obj2)

}
