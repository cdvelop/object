package object

import "github.com/cdvelop/model"

func addBasicHandlers(o *model.Object, new_struct interface{}) {
	if _, ok := new_struct.(model.PrinterHandler); ok {
		o.PrinterHandler = new_struct.(model.PrinterHandler)
	}

}
