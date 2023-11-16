package object

import (
	"reflect"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func (sf structFound) setStructField(o *model.Object, h *model.Handlers) error {

	// Crear una instancia vacía del tipo subyacente
	structValue := reflect.New(sf.struct_ref).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := field.Type()

		name_value := sf.struct_ref.Field(i).Name

		// fmt.Println("VALOR NOMBRE:", name_value)

		// Verificar si el campo "Object" existe en la estructura
		if name_value == "Object" && fieldType == reflect.TypeOf((*model.Object)(nil)) {
			field.Set(reflect.ValueOf(o)) // Asignar el campo "Object" a la estructura
		} else if h != nil && name_value == "App" && fieldType == reflect.TypeOf((*model.Handlers)(nil)) {
			field.Set(reflect.ValueOf(h)) // Asignar el campo "Handlers" a la estructura
		} else {

			// primera letra en minúscula
			if newChar, ok := strings.VALID_LETTERS[name_value[0]]; ok {
				// Verificar si el campo es de tipo string
				if fieldType.Kind() == reflect.String {

					name_value = string(newChar) + name_value[1:]

					// Asignar el nombre del campo como valor
					field.SetString(name_value)

					// Obtener y mostrar el valor de la etiqueta del campo
					fieldTag := sf.struct_ref.Field(i).Tag

					err := addObjectFields(o, name_value, fieldTag)
					if err != nil {
						return err
					}
				}
			}
		}

	}

	// Obtener una referencia a la interfaz original
	interfaceValue := reflect.ValueOf(sf.struct_int)

	// Actualizar el valor en la interfaz con la estructura modificada
	interfaceValue.Elem().Set(structValue)

	return nil
}

func addObjectFields(o *model.Object, name_value string, fieldTag reflect.StructTag) error {
	new_field := model.Field{
		Name: name_value,
	}

	for _, name := range getModelFieldNames() {
		value := fieldTag.Get(name)
		if value != "" {
			err := setFieldFromTags(&new_field, value, name, o)
			if err != nil {
				return err
			}
		}
	}

	if fieldTag.Get("Legend") != "" {
		o.PrincipalFieldsName = append(o.PrincipalFieldsName, name_value)
	}

	o.Fields = append(o.Fields, new_field)

	return nil
}

func getModelFieldNames() (names []string) {
	t := reflect.TypeOf(model.Field{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		names = append(names, field.Name)
	}

	return
}
