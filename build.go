package object

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func buildObject(sf *structFound, m *model.Module, no_add_object_to_module bool) (*model.Object, error) {

	if sf == nil {
		return nil, model.Error("estructura nil en buildObject")
	}

	obj_name := strings.ToLowerCaseAlphabet(sf.struct_ref.Name())

	var module_name string

	if m != nil && m.ModuleName != "" {
		module_name = m.ModuleName + "."
	}

	new_object := model.Object{
		Name:            module_name + obj_name,
		Table:           obj_name,
		Module:          m, // se permite modulo nulo, solo que no sera agregado a ningún lado, util para crear tablas con el objeto
		BackendHandler:  model.BackendHandler{},
		FrontendHandler: model.FrontendHandler{},
	}

	err := sf.setStructField(&new_object, handlers)
	if err != nil {
		return nil, err
	}

	addFrontHandlers(&new_object, sf.struct_int)

	addBackHandlers(&new_object, sf.struct_int)

	addBasicHandlers(&new_object, sf.struct_int)

	if m != nil && !no_add_object_to_module {
		m.Objects = append(m.Objects, &new_object)
	}

	return &new_object, nil
}
