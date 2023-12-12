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
		o.FrontHandler.AfterCreate = new_struct.(model.AfterCreate)
	}
	if _, ok := new_struct.(model.AfterUpdate); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterUpdate")
		o.FrontHandler.AfterUpdate = new_struct.(model.AfterUpdate)
	}
	if _, ok := new_struct.(model.AfterDelete); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterDelete")
		o.FrontHandler.AfterDelete = new_struct.(model.AfterDelete)
	}

	if _, ok := new_struct.(model.AfterClicked); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterClicked")
		o.FrontHandler.AfterClicked = new_struct.(model.AfterClicked)
	}

	if _, ok := new_struct.(model.NotifyBootData); ok {
		o.FrontHandler.NotifyBootData = new_struct.(model.NotifyBootData)
	}

	if _, ok := new_struct.(model.ResetFrontendObjectStateAdapter); ok {
		o.FrontHandler.ResetFrontendObjectStateAdapter = new_struct.(model.ResetFrontendObjectStateAdapter)
	}

	if _, ok := new_struct.(model.ObjectViewHandler); ok {
		o.FrontHandler.ObjectViewHandler = new_struct.(model.ObjectViewHandler)
	}

	if _, ok := new_struct.(model.NotifyFormComplete); ok {
		o.FrontHandler.NotifyFormComplete = new_struct.(model.NotifyFormComplete)
	}

}
