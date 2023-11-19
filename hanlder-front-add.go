package object

import (
	"github.com/cdvelop/model"
)

func addFrontHandlers(o *model.Object, new_struct interface{}) {

	// if _, ok := new_struct.(model.StoreData); ok {
	// 	o.StoreData = new_struct.(model.StoreData)
	// }

	if _, ok := new_struct.(model.AfterCreate); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterCreate")
		o.AfterCreate = new_struct.(model.AfterCreate)
	}
	if _, ok := new_struct.(model.AfterUpdate); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterUpdate")
		o.AfterUpdate = new_struct.(model.AfterUpdate)
	}
	if _, ok := new_struct.(model.AfterDelete); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterDelete")
		o.AfterDelete = new_struct.(model.AfterDelete)
	}

	if _, ok := new_struct.(model.AfterClicked); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterClicked")
		o.AfterClicked = new_struct.(model.AfterClicked)
	}

	if _, ok := new_struct.(model.NotifyBootData); ok {
		o.NotifyBootData = new_struct.(model.NotifyBootData)
	}

	if _, ok := new_struct.(model.ViewHandler); ok {
		o.ViewHandler = new_struct.(model.ViewHandler)
	}

	if _, ok := new_struct.(model.ViewReset); ok {
		o.ViewReset = new_struct.(model.ViewReset)
	}

}
