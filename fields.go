package object

import (
	"reflect"

	"github.com/cdvelop/model"
)

func (sf structFound) setStructField(o *model.Object, inputs ...*model.Input) error {

	// Crear una instancia vacía del tipo subyacente
	structValue := reflect.New(sf.struct_ref).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := field.Type()

		name_value := sf.struct_ref.Field(i).Name

		// fmt.Println("VALOR NOMBRE:", name_value)

		// Verificar si el campo "Object" existe en la estructura
		if name_value == "Object" && fieldType == reflect.TypeOf((*model.Object)(nil)) {
			field.Set(reflect.ValueOf(o)) // Asignar el campo "Object" en la estructura
		} else {

			// primera letra en minúscula
			if newChar, ok := valid_letters[name_value[0]]; ok {

				name_value = string(newChar) + name_value[1:]

				// Verificar si el campo es de tipo string
				if fieldType.Kind() == reflect.String {
					// Asignar el nombre del campo como valor
					field.SetString(name_value)
				}

			}
			// Obtener y mostrar el valor de la etiqueta del campo
			fieldTag := sf.struct_ref.Field(i).Tag

			err := addObjectFields(o, name_value, fieldTag, inputs...)
			if err != nil {
				return err
			}

		}

	}

	// Obtener una referencia a la interfaz original
	interfaceValue := reflect.ValueOf(sf.struct_int)

	// Actualizar el valor en la interfaz con la estructura modificada
	interfaceValue.Elem().Set(structValue)

	return nil
}

func addObjectFields(o *model.Object, name_value string, fieldTag reflect.StructTag, inputs ...*model.Input) error {
	new_field := model.Field{
		Name: name_value,
	}

	var add_field bool

	for _, name := range getModelFieldNames() {
		value := fieldTag.Get(name)
		if value != "" {
			err := setFieldFromTags(&new_field, value, name, inputs...)
			if err != nil {
				return err
			}

			add_field = true
		}
	}

	if fieldTag.Get("TextField") != "" {
		add_field = true
		o.TextFieldNames = append(o.TextFieldNames, name_value)
	}

	if add_field {
		o.Fields = append(o.Fields, new_field)
	}

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

var valid_letters = map[byte]byte{
	'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i',
	'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r',
	'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z',
}
