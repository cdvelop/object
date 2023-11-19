package object

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func (sf *structFound) buildObject(module *model.Module) (*model.Object, error) {

	if sf == nil {
		return nil, model.Error("estructura nil en buildObject")
	}

	obj_name := strings.ToLowerCaseAlphabet(sf.struct_ref.Name())

	var module_name string

	if module != nil && module.ModuleName != "" {
		module_name = module.ModuleName + "."
	}

	sf.o = &model.Object{
		Name:            module_name + obj_name,
		Table:           obj_name,
		Module:          module, // se permite modulo nulo, solo que no sera agregado a ningún lado, util para crear tablas con el objeto
		BackendHandler:  model.BackendHandler{},
		FrontendHandler: model.FrontendHandler{},
	}

	err := sf.setStructField()
	if err != nil {
		return nil, err
	}

	addFrontHandlers(sf.o, sf.struct_int)

	addBackHandlers(sf.o, sf.struct_int)

	addBasicHandlers(sf.o, sf.struct_int)

	if module != nil {
		module.Objects = append(module.Objects, sf.o)
	}

	return sf.o, nil
}
