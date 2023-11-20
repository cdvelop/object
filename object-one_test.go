package object_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

// se espera que se agregue el objeto al modulo. y use sus inputs
func TestBuildOneObject(t *testing.T) {
	input_text := input.Text()

	module := &model.Module{
		ModuleName:        "",
		Title:             "",
		IconID:            "",
		UI:                nil,
		HeaderInputTarget: "",
		Areas:             []byte{},
		Objects:           []*model.Object{},
		Inputs:            []*model.Input{input_text},
	}

	type user struct {
		Id   string
		Name string `Legend:"Nombre" Input:"Text"`
	}

	obj_expected := &model.Object{
		ObjectName: "user",
		Table:      "user",
		// PrincipalFieldsName: []string{"id", "name"},
		Fields: []model.Field{
			// {Name: "id"},
			{Name: "name", Legend: "Nombre", Input: input_text},
		},
		Module: module,
	}

	u := &user{}

	obj_resp, err := object.BuildObjectFromStruct(u, module)
	if err != nil {
		t.Fatal(err)
	}

	if !object.AreIdentical(obj_resp, obj_expected) {

		fmt.Println("TestBuildOneObject")
		fmt.Printf("\n-se esperaba:\n%v\n\n-pero se obtuvo:\n%v\n", obj_expected, obj_resp)

		log.Fatalln()
	}

	if len(module.Objects) != 1 {
		fmt.Printf("error se esperaba que se agregara 1 objeto al modulo")
		log.Fatalln()
	}

}
