package object

import (
	"reflect"

	"github.com/cdvelop/model"
)

type structFound struct {
	struct_int interface{}
	struct_ref reflect.Type

	o *model.Object
}
