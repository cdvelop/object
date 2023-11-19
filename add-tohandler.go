package object

var noAddObjectFields bool // por defecto siempre se agregan los campos validos al objeto
// arguments: main struct and *model.Module. inputs: []*model.Input, *model.Handlers
// NOTA: si la estructura tiene el campo "noAddObjectFields" de tipo bool. se crea el objeto pero sin campos. solo se agregaran los nombres a principal fields si corresponde
func AddToHandlerFromStructs(model_structs ...interface{}) error {

	st_founds, module, err := getStructFromInterface("AddToHandlerFromStructs", model_structs...)
	if err != nil {
		return err
	}

	// agregamos el modulo al manejador solo si el modulo fue ingresado
	if module != nil && handlers != nil {
		handlers.AddModules(module)
	}

	for _, sf := range st_founds {
		noAddObjectFields = false

		new_object, err := sf.buildObject(module)
		if err != nil {
			return err
		}

		// agregamos el nuevo objeto al manejador solo si el modulo  y el manejador es valido
		if handlers != nil && module != nil {
			handlers.AddObjects(new_object)
		}

	}

	return nil
}
