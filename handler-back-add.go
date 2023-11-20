package object

import (
	"github.com/cdvelop/model"
)

func addBackHandlers(o *model.Object, new_struct interface{}) {

	if _, ok := new_struct.(model.BootResponse); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE BackendHandler BootResponse")
		o.ObjectBackend.BootResponse = new_struct.(model.BootResponse)
	}
	if _, ok := new_struct.(model.CreateApi); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE BackendHandler CreateApi")
		o.ObjectBackend.CreateApi = new_struct.(model.CreateApi)
	}
	if _, ok := new_struct.(model.ReadApi); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE BackendHandler ReadApi")
		o.ObjectBackend.ReadApi = new_struct.(model.ReadApi)
	}
	if _, ok := new_struct.(model.UpdateApi); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE BackendHandler UpdateApi")
		o.ObjectBackend.UpdateApi = new_struct.(model.UpdateApi)
	}
	if _, ok := new_struct.(model.DeleteApi); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE BackendHandler DeleteApi")
		o.ObjectBackend.DeleteApi = new_struct.(model.DeleteApi)
	}

}
