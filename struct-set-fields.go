package object

import (
	"reflect"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func (sf *structFound) setStructField() error {

	// Crear una instancia vacía del tipo subyacente
	structValue := reflect.New(sf.struct_ref).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		var add_principal_field bool
		field := structValue.Field(i)
		fieldType := field.Type()

		name_value := sf.struct_ref.Field(i).Name

		switch fieldType.Kind() {

		// tipos de campos mas usados
		case reflect.Bool, reflect.Int, reflect.Int64:
			if _, ok := strings.VALID_LETTERS[name_value[0]]; ok {
				add_principal_field = true
			}

			if fieldType.Kind() == reflect.Bool && name_value == "noAddObjectFields" {
				noAddObjectFields = true
				// fmt.Println("CAMPO BOOL:", name_value, "VALOR:", field.Bool())
			}

		case reflect.Ptr: //campo puntero

			if name_value == "Object" && fieldType == reflect.TypeOf((*model.Object)(nil)) {
				field.Set(reflect.ValueOf(sf.o)) // Asignar el campo "Object" a la estructura
			} else if handlers != nil && name_value == "App" && fieldType == reflect.TypeOf((*model.Handlers)(nil)) {
				field.Set(reflect.ValueOf(handlers)) // Asignar el campo "Handlers" a la estructura

			}

		case reflect.Struct: // campo estructura
			// Crear una instancia vacía del tipo subyacente
			// fmt.Println("CAMPOS ESTRUCTURA HIJA:", daughterStruct)
			daughterStruct := reflect.New(fieldType).Elem()

			new_st_Found := structFound{
				struct_int: daughterStruct.Addr().Interface(),
				struct_ref: fieldType,
				o:          sf.o,
			}

			err := new_st_Found.setStructField()
			if err != nil {
				return err
			}

			// fmt.Println("EL CAMPO:", name_value, "ES OTRA ESTRUCTURA")
			// Llamar a la función que procesará la estructura hija
			// err := SetFieldsStructToSameName(daughterStruct.Addr().Interface())
			// if err != nil {
			// 	return err
			// }
			field.Set(daughterStruct)

		case reflect.String: // campo tipo string
			if setValueReflectStringField(&name_value, &field) {
				add_principal_field = true
				// Obtener y mostrar el valor de la etiqueta del campo
				fieldTag := sf.struct_ref.Field(i).Tag

				err := sf.addObjectFields(name_value, fieldTag)
				if err != nil {
					return err
				}
			}

		}

		sf.addPrincipalField(name_value, add_principal_field)
		// fmt.Println("VALOR NOMBRE:", name_value)

	}

	// Obtener una referencia a la interfaz original
	interfaceValue := reflect.ValueOf(sf.struct_int)

	// Actualizar el valor en la interfaz con la estructura modificada
	interfaceValue.Elem().Set(structValue)

	return nil
}

func (sf structFound) addObjectFields(name_value string, fieldTag reflect.StructTag) error {
	// fmt.Println("VALOR noAddObjectFields", noAddObjectFields)

	if !noAddObjectFields {

		new_field := model.Field{
			Name: name_value,
		}

		for _, name := range getModelFieldNames() {
			value := fieldTag.Get(name)
			if value != "" {
				err := sf.setFieldFromTags(&new_field, value, name)
				if err != nil {
					return err
				}
			}
		}

		if fieldTag.Get("Legend") != "" {

			sf.o.Fields = append(sf.o.Fields, new_field)
		}
	}

	return nil
}

func (sf structFound) addPrincipalField(name_value string, add bool) {
	if add && sf.o != nil {
		new_name := strings.LowerCaseFirstLetter(name_value)
		sf.o.PrincipalFieldsName = append(sf.o.PrincipalFieldsName, new_name)
	}
}

func getModelFieldNames() (names []string) {
	t := reflect.TypeOf(model.Field{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		names = append(names, field.Name)
	}

	return
}
