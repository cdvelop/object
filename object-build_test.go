package object_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

// se espera que no se agregue el objeto al modulo. pero que use sus inputs
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
		Name string `Input:"Text"`
	}

	obj_expected := &model.Object{
		Name:  "user",
		Table: "user",
		Fields: []model.Field{
			{Name: "id"},
			{Name: "name", Input: input_text},
		},
		Module: module,
	}

	u := &user{}

	obj_resp, err := object.BuildObjectFromStruct(true, u, module)
	if err != nil {
		t.Fatal(err)
	}

	if !object.AreIdentical(obj_resp, obj_expected) {

		fmt.Println("TestBuildOneObject")
		fmt.Printf("\n-se esperaba:\n%v\n\n-pero se obtuvo:\n%v\n", obj_expected, obj_resp)

		log.Fatalln()
	}

	if len(module.Objects) != 0 {
		fmt.Printf("error se esperaba que no se agregara el objeto al modulo")
		log.Fatalln()
	}

}
