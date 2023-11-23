package object

func SetFieldsStructToSameName(model_structs ...interface{}) (err string) {

	st_founds, _, err := getStructFromInterface("SetFieldsStructToSameName", model_structs...)
	if err != "" {
		return err
	}

	for _, sf := range st_founds {
		err := sf.setStructField()
		if err != "" {
			return err
		}
	}

	return ""
}
