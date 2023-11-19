package object

import (
	"reflect"

	"github.com/cdvelop/strings"
)

func setValueReflectStringField(name_value *string, field *reflect.Value) bool {
	if name_value != nil && field != nil {

		var new_name = *name_value
		// que la primera letra sea mayúscula
		if _, ok := strings.VALID_LETTERS[new_name[0]]; ok {

			new_name = strings.ToLowerCaseAlphabet(new_name)
			// Asignar el nombre del campo como valor
			field.SetString(new_name)
			*name_value = new_name
			return true
		}

	}
	return false
}
