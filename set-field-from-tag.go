package object

import (
	"reflect"
	"strconv"

	"github.com/cdvelop/model"
)

func setFieldFromTags(obj, value interface{}, tag_name string, o *model.Object) error {
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
			// fmt.Println("INPUT value NEW:", value_in)
			for _, mod := range o.Module.Inputs {
				// fmt.Println("INPUT value OLD:", mod.InputName)
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

	return model.Error(tag_name, value_in, "Tipo:", field_type, ", no existe en objeto:", o.Name)
}
