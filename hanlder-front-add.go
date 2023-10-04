package object

import (
	"github.com/cdvelop/model"
)

func addFrontHandlers(o *model.Object, new_struct interface{}) {

	if _, ok := new_struct.(model.AfterCreate); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE FrontendHandler AfterCreate")
		o.AfterCreate = new_struct.(model.AfterCreate)
	}
	if _, ok := new_struct.(model.AfterUpdate); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE FrontendHandler AfterUpdate")
		o.AfterUpdate = new_struct.(model.AfterUpdate)
	}
	if _, ok := new_struct.(model.AfterDelete); ok {
		// fmt.Println("ESTRUCTURA", o.Name, "CONTIENE FrontendHandler AfterDelete")
		o.AfterDelete = new_struct.(model.AfterDelete)
	}

	if _, ok := new_struct.(model.NotifyBootData); ok {
		o.NotifyBootData = new_struct.(model.NotifyBootData)
	}

	if _, ok := new_struct.(model.ViewHandler); ok {
		o.ViewHandler = new_struct.(model.ViewHandler)
	}

}
