package object

// arguments: main struct and *model.Module. inputs: []*model.Input, *model.MainHandler
func AddToHandlerFromStructs(model_structs ...interface{}) (err string) {

	st_founds, module, err := getStructFromInterface("AddToHandlerFromStructs", model_structs...)
	if err != "" {
		return err
	}

	// agregamos el modulo al manejador solo si el modulo fue ingresado
	if module != nil && handlers != nil {
		handlers.AddModules(module)

	}

	for _, sf := range st_founds {

		new_object, err := sf.buildObject(module)
		if err != "" {
			return err
		}

		// agregamos el nuevo objeto al manejador solo si el modulo  y el manejador es valido
		if handlers != nil && module != nil {
			module.MainHandler = handlers

			module.AddObjectsToModule(new_object)

			handlers.AddModules(module)
		}

	}

	return ""
}
