package object

import (
	"reflect"

	"github.com/cdvelop/model"
)

func AreIdentical(obj1, obj2 *model.Object) bool {

	if obj1.Name != obj2.Name || len(obj1.Fields) != len(obj2.Fields) {
		return false
	}

	for i := range obj1.Fields {
		if !reflect.DeepEqual(obj1.Fields[i], obj2.Fields[i]) {
			return false
		}
	}

	if len(obj1.TextFieldNames) != len(obj2.TextFieldNames) {
		return false
	}

	if obj1.FrontendHandler != obj2.FrontendHandler {
		return false
	}

	if obj1.BackendHandler != obj2.BackendHandler {
		return false
	}

	return obj1.Module == obj2.Module

}
