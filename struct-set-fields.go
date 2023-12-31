package object

import (
	"reflect"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func (sf *structFound) setStructField() (err string) {

	// Crear una instancia vacía del tipo subyacente
	structValue := reflect.New(sf.struct_ref).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := field.Type()

		name_value := sf.struct_ref.Field(i).Name

		// si la primera letra no es Mayúscula continuamos a la siguiente
		if _, ok := strings.LettersUpperLowerCase()[rune(name_value[0])]; !ok {
			continue
		}

		switch fieldType.Kind() {
		// case reflect.Bool, reflect.Int, reflect.Int64:

		case reflect.Ptr: //campo puntero

			if name_value == "Object" && fieldType == reflect.TypeOf((*model.Object)(nil)) {
				field.Set(reflect.ValueOf(sf.o)) // Asignar el campo "Object" a la estructura
			} else if handlers != nil && name_value == "App" && fieldType == reflect.TypeOf((*model.MainHandler)(nil)) {
				field.Set(reflect.ValueOf(handlers)) // Asignar el campo "MainHandler" a la estructura

			}

		case reflect.Struct: // campo estructura

			if knownName(name_value) {
				continue
			}
			// fmt.Println("EL CAMPO:", name_value, "ES OTRA ESTRUCTURA")

			// fmt.Println("CAMPOS ESTRUCTURA HIJA:", daughterStruct)
			daughterStruct := reflect.New(fieldType).Elem()

			new_st_Found := structFound{
				struct_int: daughterStruct.Addr().Interface(),
				struct_ref: fieldType,
				o:          sf.o,
			}

			// Llamar a la función que procesará la estructura hija
			err = new_st_Found.setStructField()
			if err != "" {
				return
			}

			field.Set(daughterStruct)

		case reflect.String: // campo tipo string
			if setValueReflectStringField(&name_value, &field) {
				// add_principal_field = true
				// Obtener y mostrar el valor de la etiqueta del campo
				fieldTag := sf.struct_ref.Field(i).Tag

				err := sf.addObjectFields(name_value, fieldTag)
				if err != "" {
					return err
				}
			}
		}
		// fmt.Println("VALOR NOMBRE:", name_value)
	}

	// Obtener una referencia a la interfaz original
	interfaceValue := reflect.ValueOf(sf.struct_int)

	// Actualizar el valor en la interfaz con la estructura modificada
	interfaceValue.Elem().Set(structValue)

	return ""
}

func (sf structFound) addObjectFields(name_value string, fieldTag reflect.StructTag) (err string) {

	new_field := model.Field{
		Name: name_value,
	}

	for _, name := range getModelFieldNames() {
		value := fieldTag.Get(name)
		if value != "" {
			err := sf.setFieldFromTags(&new_field, value, name)
			if err != "" {
				return err
			}
		}
	}

	if fieldTag.Get("Legend") != "" {

		sf.o.Fields = append(sf.o.Fields, new_field)
	}

	return ""
}

func getModelFieldNames() (names []string) {
	t := reflect.TypeOf(model.Field{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		names = append(names, field.Name)
	}

	return
}
