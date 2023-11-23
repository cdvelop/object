package object_test

import (
	"testing"

	"github.com/cdvelop/object"
)

func TestFileInputCase(t *testing.T) {

	type stock struct {
		Id string
	}
	s := &stock{}

	err := object.SetFieldsStructToSameName(s)
	if err != "" {
		t.Fatal(err)
	}

	// fmt.Println("RESULTADO:", f)
	if s.Id != "id" {
		t.Fatal("Se esperaba id, pero se obtuvo:", s.Id)
	}
}
