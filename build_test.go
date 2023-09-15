package object_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

type person struct {
	no_include string // campos sin etiqueta Legend no se incluyen en el objeto
	name       string `Legend:"Nombre" NotRenderHtml:"true" Input:"TextOnly" TextField:"ok"`
	age        int    `Legend:"Edad" Encrypted:"true" Input:"Number"`
	address    int    `Legend:"Dirección" Input:"Text"`
}

type user struct {
	none string //estructura sin campos
}

type staff struct {
	name  string `Legend:"Nombre"`
	other *model.Field
}

func (person) SetObjectInDomAfterDelete(data ...map[string]string) (container_id, tags string) {
	return
}
func (person) Delete(data ...map[string]string) (out []map[string]string, err error) {
	return
}

func TestBuildObjectFromStruct(t *testing.T) {

	module := &model.Module{ModuleName: "Patient"}

	inputs := []*model.Input{
		input.Text(),
		input.Number(),
		input.TextOnly(),
	}

	new_staff := staff{}

	dataTest := map[string]struct {
		model_struct interface{}
		expected     *model.Object
		err          string
	}{
		"1- estructura person expected_object1 solo con handlers front back delete": {
			model_struct: person{},
			expected: &model.Object{
				Name:           "person",
				TextFieldNames: []string{"name"},
				Fields: []model.Field{
					// {Name: "FullName", Legend: "Nombre"},
					{Name: "name", Legend: "Nombre", NotRenderHtml: true, Input: input.TextOnly()},
					{Name: "age", Legend: "Edad", Encrypted: true, Input: input.Number()},
					{Name: "address", Legend: "Dirección", Input: input.Text()},
				},
				Module:          module,
				BackendHandler:  model.BackendHandler{DeleteApi: person{}},
				FrontendHandler: model.FrontendHandler{AfterDelete: person{}},
			},
			err: "",
		},
		"2- estructura user solo un campo sin modulo se espera error": {
			model_struct: user{},
			expected: &model.Object{
				Name: "user",
			},
			err: "error estructura principal no ingresada (verifica si los campos tiene el tag 'Legend')",
		},
		"3- estructura staff ya inicializada solo un campo con modulo se espera ok": {
			model_struct: new_staff,
			expected: &model.Object{
				Name: "staff",
				Fields: []model.Field{
					{Name: "name", Legend: "Nombre"},
				},
				Module: module,
			},
			err: "",
		},
	}

	for prueba, data := range dataTest {
		t.Run((prueba), func(t *testing.T) {

			err := object.New(data.model_struct, inputs, module)

			if err != nil {

				if data.err != err.Error() {
					log.Fatalf("\n-se obtuvo:\n%v", err)
				}

			} else {

				var resp *model.Object

				for _, o := range module.Objects {
					if o.Name == data.expected.Name {
						resp = o
						break
					}
				}

				if !object.AreIdentical(resp, data.expected) {

					fmt.Println(prueba)
					fmt.Printf("\n-se esperaba:\n%v\n-pero se obtuvo:\n%v\n", data.expected, resp)

					log.Fatalln("error:", err)
				}
			}

		})
	}

}
