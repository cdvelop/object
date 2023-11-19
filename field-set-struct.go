package object

func SetFieldsStructToSameName(model_structs ...interface{}) error {

	st_founds, _, err := getStructFromInterface("SetFieldsStructToSameName", model_structs...)
	if err != nil {
		return err
	}

	for _, sf := range st_founds {
		err := sf.setStructField()
		if err != nil {
			return err
		}
	}

	return nil
}
