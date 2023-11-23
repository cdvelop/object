package object_test

import (
	"log"
	"testing"
)

func Test_Validate(t *testing.T) {
	for prueba, data := range objectTestData {
		t.Run((prueba), func(t *testing.T) {

			object := objects[data.Object]

			err := object.ValidateData(data.ItsNew, data.ItsUpdate, data.Data)

			if err == "" {
				if data.Expected != "" {
					log.Fatalf("\n=>expected:[%v]\n=>result:[%v]", data.Expected, err)
				}
			} else {
				if data.Expected != "error" {
					log.Fatalf("\n=>expected:[%v]\n=>result:[%v]", data.Expected, err)
				}
			}

			// fmt.Printf("%v %v\n", message, ok)
		})
	}
}
