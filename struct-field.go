package object

import (
	"reflect"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func SetFieldsStructToSameName(modelStructs ...interface{}) error {
	for _, m := range modelStructs {
		t := reflect.TypeOf(m)

		if t.Kind() != reflect.Ptr {
			return model.Error("SetFieldsStructToSameName error debes ingresar las estructuras como punteros")
		}

		st := t.Elem()

		if st.Kind() != reflect.Struct {
			return model.Error("el puntero debe ser de tipo estructura válida")
		}

		structValue := reflect.New(st).Elem()

		for i := 0; i < structValue.NumField(); i++ {
			field := structValue.Field(i)
			fieldType := field.Type()
			nameValue := st.Field(i).Name

			nameValue = strings.LowerCaseFirstLetter(nameValue)

			if fieldType.Kind() == reflect.String {
				field.SetString(nameValue)
			}
		}

		interfaceValue := reflect.ValueOf(m)
		interfaceValue.Elem().Set(structValue)
	}

	return nil
}
