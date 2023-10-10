package object_test

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/unixid"
)

const (
	TableName1 = "usuario"
	TableName2 = "shopping_cart"

	id_user_key    = "id_usuario"
	nameKey        = "name"
	genderKey      = "gender"
	descriptionKey = "description"
	rutKey         = "rut_usuario"
)

var (
	objects = map[string]model.Object{
		Object1().Name: *Object1(),
		Object2().Name: *Object2(),
	}
)

type dataGenero struct{}

func (dataGenero) SourceData() map[string]string {
	return map[string]string{"D": "Dama", "V": "Varón"}
}

func Object1() *model.Object {
	t1 := model.Object{
		Name:  TableName1,
		Table: TableName1,
		Fields: []model.Field{
			{Name: id_user_key, Input: unixid.InputPK()},
			{Name: nameKey, Input: input.Text(), Legend: "Nombre"},
			{Name: genderKey, Input: input.Radio(dataGenero{})},
			{Name: descriptionKey, Input: input.Text(), Legend: "Descripción", SkipCompletionAllowed: true},
			{Name: rutKey, Input: input.Rut(), Unique: true, SkipCompletionAllowed: true},
		},
	}

	return &t1
}

type dataState struct{}

func (dataState) SourceData() map[string]string {
	return map[string]string{
		"inuse": "En Uso", "lost": "Perdido",
	}
}

func Object2() *model.Object {

	t2 := model.Object{
		Name:  TableName2,
		Table: TableName2,
		Fields: []model.Field{
			{Name: "id_" + TableName2, Input: unixid.InputPK()},
			{Name: id_user_key, Legend: "Id Usuario", Input: unixid.InputPK()},
			{Name: TableName2 + "_state", Input: input.Radio(dataState{})},
		},
	}

	return &t2
}
