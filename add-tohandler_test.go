package object_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
	"github.com/cdvelop/unixid"
)

// NOTA:
// campos sin inputs definidos se incluyen en el objeto
// al tener tag Legend se agregan al slice string []PrincipalFieldName
type person struct {
	Id_person  string `Legend:"Id" Input:"InputPK"`
	no_include string
	name       string `NotRenderHtml:"true" Input:"TextOnly" ` // no incluye empieza con minúscula
	Age        int    `Encrypted:"true"`                       // tipo int no se incluye au que tenga mayúscula, no tiene input ni legend
	Address    string `Legend:"Dirección" Input:"Text"`
	Cars       string `Legend:"Vehículos" Input:"Text" SourceTable:"cars"`
	Other      string // no tiene ninguna etiqueta solo se le asigna como valor el nombre se su campo y como campo principal
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

	handlers := &model.Handlers{
		ThemeAdapter:    nil,
		DataBaseAdapter: nil,
		TimeAdapter:     nil,
		DomAdapter:      nil,
		FetchAdapter:    nil,
		AuthAdapter:     nil,
		Logger:          nil,
	}

	add_inputs := []*model.Input{
		unixid.InputPK(),
		input.Text(),
		input.Number(),
		input.TextOnly(),
	}

	mod_one := &model.Module{ModuleName: "mod_one", Inputs: add_inputs}

	mod_three := &model.Module{ModuleName: "mod_three"}
	mod_four := &model.Module{ModuleName: "mod_four"}
	mod_five := &model.Module{ModuleName: "mod_five"}

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
		handlers     *model.Handlers
		model_struct []interface{}
		expected     []*model.Object
		err          string
	}{
		"1- estructura person expected_object1 solo con handlers front back delete": {
			module:       mod_one,
			handlers:     handlers,
			model_struct: []interface{}{new_person},
			expected: []*model.Object{
				{
					Name:                mod_one.ModuleName + ".person",
					Table:               "person",
					PrincipalFieldsName: []string{"id_person", "age", "address", "cars", "other"},
					Fields: []model.Field{
						{Name: "id_person", Legend: "Id", Input: unixid.InputPK()},
						{Name: "address", Legend: "Dirección", Input: input.Text()},
						{Name: "cars", Legend: "Vehículos", Input: input.Text(), SourceTable: "cars"},
						// {Name: "other"},
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
			err: "error debes de ingresar las estructuras como punteros en AddToHandlerFromStructs",
		},
		"3- estructura staff ya inicializada, un campo, sin tags, modulo y 1 handler front se espera ok": {
			module:       mod_three,
			handlers:     handlers,
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
			handlers:     handlers,
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
			handlers:     handlers,
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
			new_data = append(new_data, data.module)
			new_data = append(new_data, data.handlers)

			err := object.AddToHandlerFromStructs(new_data...)
			if err != nil {
				if data.err != err.Error() {
					log.Fatalf("\n-mensaje error diferente:\n-se esperaba:%v\n-pero se obtuvo:\n%v\n", data.err, err.Error())
				}
			} else {

				if len(data.expected) != len(data.module.Objects) {

					fmt.Printf("-se esperaba:%v -pero se obtuvo:%v objeto(s)\n\n", len(data.expected), len(data.module.Objects))

					for i, o := range data.module.Objects {
						fmt.Println("---", i)
						fmt.Println(o)
						fmt.Println()
					}

					log.Fatalln()
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

func TestStructWhitOutModule(t *testing.T) {
	// estructura sin *model.Module se debería crear objeto pero no ingresarlo
	type stock struct {
		Object *model.Object
		Id     string
		Name   string `Legend:"Nombre"`
	}

	s := &stock{}

	err := object.AddToHandlerFromStructs(s)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println("RESULTADO:", f)
	if s.Object == nil {
		t.Fatal("Se esperaba Objeto asignado pero se obtuvo:", s.Object)
	}

	if s.Id != "id" {
		t.Fatal("Se esperaba campo Id con valor: 'id' pero se obtuvo:", s.Id)
	}
	if s.Name != "name" {
		t.Fatal("Se esperaba campo Name con valor: 'name' pero se obtuvo:", s.Name)
	}

	if len(s.Object.Fields) != 1 {
		t.Fatal("Se esperaba 1 campo creado pero se obtuvo:", len(s.Object.Fields))
	}

	// fmt.Println(s)
}
