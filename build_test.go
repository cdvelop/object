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
	name       string `Legend:"Nombre" NotRenderHtml:"true" TextField:"1" Input:"TextOnly" `
	age        int    `Legend:"Edad" Encrypted:"true" Input:"Number"`
	address    int    `Legend:"Dirección" Input:"Text"`
}

func (person) SetObjectInDomAfterDelete(data ...map[string]string) (container_id, tags string) {
	return
}
func (person) Delete(data ...map[string]string) (out []map[string]string, err error) {
	return
}

type user struct {
	none string //estructura sin campos
}

type staff struct {
	name  string
	other *model.Field
}

func (staff) Delete(data ...map[string]string) (out []map[string]string, err error) {
	return
}

func TestBuildObjectFromStruct(t *testing.T) {

	module_one := &model.Module{ModuleName: "module_one"}
	module_three := &model.Module{ModuleName: "module_three"}
	module_four := &model.Module{ModuleName: "module_four"}

	inputs := []*model.Input{
		input.Text(),
		input.Number(),
		input.TextOnly(),
	}

	new_staff := staff{}

	dataTest := map[string]struct {
		module       *model.Module
		model_struct []interface{}
		expected     []*model.Object
		err          string
	}{
		"1- estructura person expected_object1 solo con handlers front back delete": {
			module:       module_one,
			model_struct: []interface{}{person{}},
			expected: []*model.Object{
				{
					Name:           "person",
					TextFieldNames: []string{"name"},
					Fields: []model.Field{
						// {Name: "FullName", Legend: "Nombre"},
						{Name: "name", Legend: "Nombre", NotRenderHtml: true, Input: input.TextOnly()},
						{Name: "age", Legend: "Edad", Encrypted: true, Input: input.Number()},
						{Name: "address", Legend: "Dirección", Input: input.Text()},
					},
					Module:          module_one,
					BackendHandler:  model.BackendHandler{DeleteApi: person{}},
					FrontendHandler: model.FrontendHandler{AfterDelete: person{}},
				},
			},
			err: "",
		},
		"2- estructura user solo un campo sin modulo se espera error": {
			module:       nil,
			model_struct: []interface{}{user{}},
			expected: []*model.Object{
				{
					Name: "user",
				},
			},
			err: "error puntero de *model.Module no ingresado como argumento",
		},
		"3- estructura staff ya inicializada, un campo, sin tags, modulo y 1 handler front se espera ok": {
			module:       module_three,
			model_struct: []interface{}{new_staff},
			expected: []*model.Object{
				{
					Name:           "staff",
					Module:         module_three,
					BackendHandler: model.BackendHandler{DeleteApi: new_staff},
				},
			},
			err: "",
		},
		"4- ingreso de 2 estructuras staff y user solo un campo con modulo se espera ok": {
			module:       module_four,
			model_struct: []interface{}{staff{}, user{}},
			expected: []*model.Object{
				{
					Name:           "staff",
					Module:         module_four,
					BackendHandler: model.BackendHandler{DeleteApi: staff{}},
				},
				{
					Name:   "user",
					Module: module_four,
				},
			},
			err: "",
		},
	}

	for prueba, data := range dataTest {
		t.Run((prueba), func(t *testing.T) {

			var new_data []interface{}

			new_data = append(new_data, data.model_struct...)
			new_data = append(new_data, inputs)
			new_data = append(new_data, data.module)

			err := object.New(new_data...)
			if err != nil {
				if data.err != err.Error() {
					log.Fatalf("\n-se esperaba:\n%v\n-pero se obtuvo:\n%v\n", data.err, err.Error())
				}
			} else {

				if len(data.expected) != len(data.module.Objects) {
					fmt.Printf("-se esperaba:%v -pero se obtuvo:%v objeto(s)\n\n", len(data.expected), len(data.module.Objects))

					for _, o := range data.module.Objects {
						fmt.Println(o)
						fmt.Println()
					}

					log.Fatal()
				}

				for _, obj_expected := range data.expected {

					for _, obj_resp := range data.module.Objects {

						if obj_resp.Name == obj_expected.Name {

							if !object.AreIdentical(obj_resp, obj_expected) {

								fmt.Println(prueba)
								fmt.Printf("\n-se esperaba:\n%v\n\n-pero se obtuvo:\n%v\n", obj_expected, obj_resp)

								log.Fatalln()
							}

							break
						}
					}
				}
			}
		})
	}
}
