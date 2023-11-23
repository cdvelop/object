package object

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func (sf *structFound) buildObject(module *model.Module) (o *model.Object, err string) {

	if sf == nil {
		return nil, "estructura nil en buildObject"
	}

	obj_name := strings.ToLowerCaseAlphabet(sf.struct_ref.Name())

	var module_name string

	if module != nil && module.ModuleName != "" {
		module_name = module.ModuleName + "."
	}

	sf.o = &model.Object{
		ObjectName: module_name + obj_name,
		Table:      obj_name,
		Module:     module,
	}

	err = sf.setStructField()
	if err != "" {
		return nil, err
	}

	addFrontHandlers(sf.o, sf.struct_int)

	addBackHandlers(sf.o, sf.struct_int)

	addBasicHandlers(sf.o, sf.struct_int)

	if module != nil {
		module.Objects = append(module.Objects, sf.o)
	}

	return sf.o, ""
}
