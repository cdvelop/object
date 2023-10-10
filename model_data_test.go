package object_test

type kv map[string]string

var (
	objectTestData = map[string]*struct {
		Object    string
		Data      map[string]string
		ItsNew    bool   //validar como nuevo (todo)
		ItsUpdate bool   //validar como  actualización o solo campos requeridos presentes
		Expected  string //resultado esperado
	}{
		"todos los campos correctos?": {TableName1,
			kv{nameKey: "Luis", genderKey: "D", descriptionKey: "no tiene", rutKey: "73528171-2"},
			true, false, ""},

		"tabla existe?": {"tablex",
			kv{nameKey: "Marco", genderKey: "D", descriptionKey: "no tiene"},
			true, false, "error"},

		"todos los campos, nombre correcto?": {TableName1,
			kv{nameKey: "Luis#1", genderKey: "D", descriptionKey: "no tiene"},
			true, false, "error"},

		"solo campos requeridos?": {TableName1,
			kv{nameKey: "Maria Joaquina", genderKey: "D"},
			true, false, ""},

		"id + todos los campos correctos?": {TableName1,
			kv{id_user_key: "123456789", nameKey: "Luis", genderKey: "D", descriptionKey: "no tiene"},
			true, false, ""},

		"new todos los campos?": {TableName1,
			kv{nameKey: "Luis"},
			true, false, "error"},

		"actualización nombre solo texto correcto?": {TableName1,
			kv{nameKey: "1u50"},
			false, true, "error"},

		"llave primaria valida en update?": {TableName1,
			kv{id_user_key: "OR"},
			false, true, "error"},

		"llave primaria existente en update?": {TableName1,
			kv{genderKey: "D"},
			false, true, "error"},

		// // test en tabla 2
		"campos correctos? tabla 2 ": {TableName2,
			kv{
				id_user_key:           "222",
				TableName2 + "_state": "inuse",
				nameKey:               "<script>",
				genderKey:             "=",
			},
			true, false, "error"},

		"todos los campos pertenecen a tabla 2?": {TableName2,
			kv{
				id_user_key:           "222",
				TableName2 + "_state": "inuse",
				nameKey:               "Juana",
			},
			true, false, "error"},

		"campos correctos tabla 2?": {TableName2,
			kv{
				id_user_key:           "222",
				TableName2 + "_state": "inuse",
			},
			true, false, ""},

		"tabla 2 nueva ok sin id ni fk id?": {TableName2,
			kv{
				TableName2 + "_state": "lost",
			},
			true, false, "error"},

		"modificar campo inalterable?": {TableName1,
			kv{id_user_key: "123456789", rutKey: "12.03"},
			false, true, "error"},
	}
)
