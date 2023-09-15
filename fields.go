package object

import (
	"reflect"
	"strconv"

	"github.com/cdvelop/model"
)

func buildFieldsObject(t reflect.Type, inputs ...*model.Input) (fields []model.Field, TextFieldNames []string, err error) {

	field_names := getModelFieldNames()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		var add_field bool

		new_field := model.Field{
			Name: field.Name,
		}

		// Leer las etiquetas de cada campo y completar el new_field
		for _, name := range field_names {

			// fmt.Println("ETIQUETA:", tag)

			value := field.Tag.Get(name)

			if name == "Legend" && value != "" {
				add_field = true
			}

			if value != "" {
				// fmt.Printf("Campo: %s, Etiqueta: %s, Valor: %s\n", field.Name, tag, value)
				// fmt.Printf("TIPO: %s\n", typ)

				// fmt.Println("INGRESAR INPUTS AQUI")
				err := setFieldFromTags(&new_field, value, name, inputs...)
				if err != nil {
					return nil, nil, err
				}
			}
		}

		// verificamos si contiene el campo la etiqueta TextField
		if field.Tag.Get("TextField") != "" {
			TextFieldNames = append(TextFieldNames, field.Name)
		}

		if add_field {
			fields = append(fields, new_field)
		}

	}

	return
}

func getModelFieldNames() (names []string) {
	t := reflect.TypeOf(model.Field{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		names = append(names, field.Name)
	}

	return
}

func setFieldFromTags(obj, value interface{}, tag_name string, inputs ...*model.Input) error {
	// Obtener el valor reflect.Value de obj
	val := reflect.ValueOf(obj).Elem()

	// Obtener el campo por su nombre
	field := val.FieldByName(tag_name)

	field_type := field.Type().String()

	var value_in string

	// Verificar si el campo existe y es exportado
	if !field.IsValid() || !field.CanSet() {
		return model.Error("Campo", tag_name, "no encontrado o no se puede modificar")
	}

	switch field_type {
	case "string":
		if valueType, ok := value.(string); ok {
			field.SetString(valueType)
			return nil
		}

	case "int":
		if valueType, ok := value.(int); ok {
			field.SetInt(int64(valueType))
			return nil
		}

	case "bool":
		var bool_value, ok bool
		var err error
		if bool_value, ok = value.(bool); ok {

		} else {
			if bool_string, ok := value.(string); ok {
				// fmt.Println("ES STRING NO BOOL", bool_string, tag_name)
				bool_value, err = strconv.ParseBool(bool_string)
			}
		}

		if err == nil {
			field.SetBool(bool_value)
			return nil
		}

	case "*model.Input":

		if tag_name == "Input" {
			value_in = value.(string)
			// fmt.Println("tag_name", tag_name, "TIPO", field_type)
			// fmt.Println("value_in", value_in)
			for _, mod := range inputs {
				// fmt.Println("MODEL INPUT:", mod.InputName)
				if value_in == mod.InputName {
					// Crear un reflect.Value para el puntero
					ptrValue := reflect.ValueOf(mod)
					// Asignar el valor del puntero al campo
					field.Set(ptrValue)
					return nil
				}
			}
		}
	default:

	}

	return model.Error("error etiqueta:", tag_name, ":", value_in, "Tipo:", field_type, ", no soportada")
}
