package object_test

import (
	"fmt"
	"testing"

	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

type Patient struct {
	Id          string
	PatientName string
	address     string // campo en minúscula no se asignara su valor
	Phone       int
	Email       string `Legend:"Correo" Input:"Text"` // solo este campo se creara en el objeto
}

func TestCompleteFieldValuesFromChildrenStructONE(t *testing.T) {
	handlers := &model.Handlers{}

	module := &model.Module{
		ModuleName: "client",
		Title:      "Cliente",
		Areas:      []byte{'s'},
		Objects:    []*model.Object{},
		Inputs:     []*model.Input{input.Text()},
	}

	type document struct {
		Number string
		Patient
	}

	doc := &document{}

	err := object.AddToHandlerFromStructs(doc, handlers, module)
	if err != nil {
		t.Fatal(err)
	}

	if doc.Id != "id" {
		t.Fatal("Se esperaba que doc.Id fuera 'id', pero es:", doc.Id)
	}

	if doc.PatientName != "patientname" {
		t.Fatal("Se esperaba que doc.PatientName fuera 'patientname', pero es:", doc.Patient.PatientName)
	}
	if doc.Phone != 0 {
		t.Fatal("Se esperaba que doc.Phone fuera '0', pero es:", doc.Phone)
	}

	if doc.address != "" {
		t.Fatal("Se esperaba que doc.address vació '', pero es:", doc.address)
	}

	//1-  se espera la creación del campo Email
	obj_expected := &model.Object{
		ObjectName:          module.ModuleName + ".document",
		Table:               "document",
		PrincipalFieldsName: []string{doc.Number, doc.Id, doc.PatientName, "phone", doc.Email},
		Fields: []model.Field{
			{Name: doc.Email, Legend: "Correo", Input: module.Inputs[0]},
		},
		Module: module,
	}

	obj_result := handlers.GetObjects()[0]
	if !object.AreIdentical(obj_result, obj_expected) {
		fmt.Printf("\n-se esperaba:\n%v\n\n-pero se obtuvo:\n%v\n", obj_expected, obj_result)
		t.Fatal()
		return
	}

}

func TestCompleteFieldValuesFromChildrenStructTWO(t *testing.T) {
	// 2-  CASO DONDE SE NECESITA integrar el modulo al objeto
	handlers := &model.Handlers{}

	module := &model.Module{
		ModuleName: "client",
		Title:      "Cliente",
		Areas:      []byte{'s'},
		Objects:    []*model.Object{},
		Inputs:     []*model.Input{input.Text()},
	}

	type document struct {
		*model.Object
		Number string
		Patient
	}

	doc := &document{}

	err := object.AddToHandlerFromStructs(doc, handlers, module)
	if err != nil {
		t.Fatal(err)
	}

	if doc.Object == nil {
		t.Fatal("se esperaba que el objeto no sea nulo pero es", doc.Object)
		return
	}

	if doc.Handlers == nil {
		t.Fatal("se esperaba que Handlers no sea nulo pero es", doc.Handlers)
		return
	}

	// fmt.Println("RESULTADO OBJETO:", doc.ObjectName, "Modulo:", doc.ModuleName)

}
