package object

import (
	"github.com/cdvelop/model"
)

func addFrontHandlers(o *model.Object, new_struct interface{}) {

	// if _, ok := new_struct.(model.StoreData); ok {
	// 	o.StoreData = new_struct.(model.StoreData)
	// }

	if h, ok := new_struct.(model.AfterCreate); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterCreate")
		o.FrontHandler.AfterCreate = h
	}
	if h, ok := new_struct.(model.AfterUpdate); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterUpdate")
		o.FrontHandler.AfterUpdate = h
	}
	if h, ok := new_struct.(model.AfterDelete); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterDelete")
		o.FrontHandler.AfterDelete = h
	}

	if h, ok := new_struct.(model.AfterClickNotify); ok {
		// fmt.Println("ESTRUCTURA", o.ObjectName, "CONTIENE FrontendHandler AfterClickNotify")
		o.FrontHandler.AfterClickNotify = h
	}

	if h, ok := new_struct.(model.NotifyBootData); ok {
		o.FrontHandler.NotifyBootData = h
	}

	if h, ok := new_struct.(model.ResetFrontendObjectStateAdapter); ok {
		o.FrontHandler.ResetFrontendObjectStateAdapter = h
	}

	if h, ok := new_struct.(model.ViewHandlerObject); ok {
		o.FrontHandler.ViewHandlerObject = h
	}

	if h, ok := new_struct.(model.FormNotify); ok {
		o.FrontHandler.FormNotify = h
	}

}
