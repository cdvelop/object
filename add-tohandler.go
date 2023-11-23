package object

// arguments: main struct and *model.Module. inputs: []*model.Input, *model.Handlers
func AddToHandlerFromStructs(model_structs ...interface{}) (err string) {

	st_founds, module, err := getStructFromInterface("AddToHandlerFromStructs", model_structs...)
	if err != "" {
		return err
	}

	// agregamos el modulo al manejador solo si el modulo fue ingresado
	if module != nil && handlers != nil {
		handlers.AddModules(module)

		if tests != nil {
			handlers.Test = tests
		}
	}

	for _, sf := range st_founds {

		new_object, err := sf.buildObject(module)
		if err != "" {
			return err
		}

		// agregamos el nuevo objeto al manejador solo si el modulo  y el manejador es valido
		if handlers != nil && module != nil {
			handlers.AddObjects(new_object)
			module.Handlers = handlers
		}

	}

	return ""
}
