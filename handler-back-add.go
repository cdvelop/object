package object

import (
	"github.com/cdvelop/model"
)

func addBackHandlers(o *model.Object, new_struct interface{}) {

	if _, ok := new_struct.(model.BootResponse); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE BackendHandler BootResponse")
		o.BootResponse = new_struct.(model.BootResponse)
	}
	if _, ok := new_struct.(model.CreateApi); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE BackendHandler CreateApi")
		o.CreateApi = new_struct.(model.CreateApi)
	}
	if _, ok := new_struct.(model.ReadApi); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE BackendHandler ReadApi")
		o.ReadApi = new_struct.(model.ReadApi)
	}
	if _, ok := new_struct.(model.UpdateApi); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE BackendHandler UpdateApi")
		o.UpdateApi = new_struct.(model.UpdateApi)
	}
	if _, ok := new_struct.(model.DeleteApi); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE BackendHandler DeleteApi")
		o.DeleteApi = new_struct.(model.DeleteApi)
	}

}
