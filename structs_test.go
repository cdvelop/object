package object_test

import (
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

func TestCompleteFieldValuesFromChildrenStructTWO(t *testing.T) {
	// 2-  CASO DONDE SE NECESITA integrar el modulo al objeto
	handlers := &model.Handlers{
		AppInfo: model.AppInfo{
			Business_name: "NN",
		},
	}

	module := &model.Module{
		ModuleName: "client",
		Title:      "Cliente",
		Areas:      map[string]string{"s": "OK"},
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
	if err != "" {
		t.Fatal(err)
	}

	if doc.Table == "" {
		t.Fatal("se esperaba acceder al campo Table del objeto de forma directa")
		return
	}
	// fmt.Printf("Dirección de memoria ya no repetida document: %p\n", doc)

	if len(handlers.GetObjects()) != 1 {
		t.Fatal("se esperaba que creara un objeto en handlers")
		return
	}

	if doc.Handlers == nil {
		t.Fatal("se esperaba que Handlers no fuera nulo")
		return
	}
	if doc.Business_name == "" {
		t.Fatal("se esperaba que la variable Business_name no estuviese vacía")
		return
	}
	// fmt.Println("OBJETO:", doc.ObjectName, "Modulo:", doc.ModuleName)
	// fmt.Println("RESULTADO Object:", doc.Object)

}
