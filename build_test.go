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
	name       string `Legend:"Nombre" NotRenderHtml:"true" PrincipalField:"1" Input:"TextOnly" `
	age        int    `Legend:"Edad" Encrypted:"true" Input:"Number"`
	Address    string `Legend:"Dirección" Input:"Text"`
	Cars       string `Legend:"Vehículos" Input:"Text" SourceTable:"cars"`
}

func (person) SetObjectInDomAfterDelete(data ...map[string]string) (err error) {
	return
}
func (person) Delete(u *model.User, data ...map[string]string) (out []map[string]string, err error) {
	return
}

type user struct {
	none string //estructura sin campos
}

type staff struct {
	name  string
	other *model.Field
}

func (staff) Delete(u *model.User, data ...map[string]string) (out []map[string]string, err error) {
	return
}

type product struct {
	Object *model.Object
	name   string
}

type stock struct {
	Object *model.Object
	id     string
}

func TestBuildObjectFromStruct(t *testing.T) {

	mod_one := &model.Module{ModuleName: "mod_one"}
	mod_three := &model.Module{ModuleName: "mod_three"}
	mod_four := &model.Module{ModuleName: "mod_four"}
	mod_five := &model.Module{ModuleName: "mod_five"}

	inputs := []*model.Input{
		input.Text(),
		input.Number(),
		input.TextOnly(),
	}

	new_person := &person{}

	new_staff3 := &staff{}
	new_user4 := &user{}
	new_staff4 := &staff{}

	new_product := product{
		Object: nil,
		name:   "",
	}

	new_stock := stock{
		Object: nil,
		id:     "",
	}

	dataTest := map[string]struct {
		module       *model.Module
		model_struct []interface{}
		expected     []*model.Object
		err          string
	}{
		"1- estructura person expected_object1 solo con handlers front back delete": {
			module:       mod_one,
			model_struct: []interface{}{new_person},
			expected: []*model.Object{
				{
					Name:                mod_one.ModuleName + ".person",
					Table:               "person",
					NamePrincipalFields: []string{"name"},
					Fields: []model.Field{
						// {Name: "FullName", Legend: "Nombre"},
						{Name: "name", Legend: "Nombre", NotRenderHtml: true, Input: input.TextOnly()},
						{Name: "age", Legend: "Edad", Encrypted: true, Input: input.Number()},
						{Name: "address", Legend: "Dirección", Input: input.Text()},
						{Name: "cars", Legend: "Vehículos", Input: input.Text(), SourceTable: "cars"},
					},
					Module:          mod_one,
					BackendHandler:  model.BackendHandler{DeleteApi: new_person},
					FrontendHandler: model.FrontendHandler{AfterDelete: new_person},
				},
			},
			err: "",
		},
		"2- estructura user solo un campo sin modulo ni estructura como puntero se espera error": {
			module:       nil,
			model_struct: []interface{}{user{}},
			expected: []*model.Object{
				{
					Name: "user",
				},
			},
			err: "error debes de ingresar las estructuras como  punteros.",
		},
		"3- estructura staff ya inicializada, un campo, sin tags, modulo y 1 handler front se espera ok": {
			module:       mod_three,
			model_struct: []interface{}{new_staff3},
			expected: []*model.Object{
				{
					Name:           mod_three.ModuleName + ".staff",
					Table:          "staff",
					Module:         mod_three,
					BackendHandler: model.BackendHandler{DeleteApi: new_staff3},
				},
			},
			err: "",
		},
		"4- ingreso de 2 estructuras staff y user solo un campo con modulo se espera ok": {
			module:       mod_four,
			model_struct: []interface{}{new_staff4, new_user4},
			expected: []*model.Object{
				{
					Name:           mod_four.ModuleName + ".staff",
					Table:          "staff",
					Module:         mod_four,
					BackendHandler: model.BackendHandler{DeleteApi: new_staff4},
				},
				{
					Name:   mod_four.ModuleName + ".user",
					Table:  "user",
					Module: mod_four,
				},
			},
			err: "",
		},
		"5- 2 estructuras product y stock como punteros se espera agregar el puntero del objeto a cada una de ellas": {
			module:       mod_five,
			model_struct: []interface{}{&new_product, &new_stock},
			expected: []*model.Object{
				{
					Name:   mod_five.ModuleName + ".product",
					Table:  "product",
					Module: mod_five,
				},
				{
					Name:   mod_five.ModuleName + ".stock",
					Table:  "stock",
					Module: mod_five,
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
					log.Fatalf("\n-no se esperaba error pero se obtuvo:\n%v\n", err.Error())
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

				for i, obj_expected := range data.expected {

					obj_resp := data.module.Objects[i]

					if !object.AreIdentical(obj_resp, obj_expected) {

						fmt.Println(prueba)
						fmt.Printf("\n-se esperaba:\n%v\n\n-pero se obtuvo:\n%v\n", obj_expected, obj_resp)

						log.Fatalln()
					}

				}
			}
		})
	}

	if new_product.Object == nil {
		log.Fatalln("se esperaba puntero valido de objeto en estructura product")
	}

	if new_stock.Object == nil {
		log.Fatalln("se esperaba puntero valido de objeto en estructura stock")
	}

	if new_person.Address != "address" {
		log.Fatalln("se esperaba que el campo Address en estructura person fuera ahora en minúscula")

	}

}
