package object_test

import (
	"log"
	"testing"
)

func Test_GetDataTestObject(t *testing.T) {

	var required = 2

	all_data, err := Object1().TestData(required, false, false)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("DATA OBTENIDA: ", all_data)

	if len(all_data) != required {
		log.Fatal("se esperaba ", required, "pero se obtuvo: ", len(all_data))
	}

	for _, data := range all_data {
		if len(data) != len(Object1().Fields) {
			log.Fatal("se esperaban ", len(Object1().Fields), " campos pero se obtuvo: ", len(data))
		}
	}

	// 2 prueba saltándose el id
	all_data2, err := Object1().TestData(required, true, false)

	expected := len(Object1().Fields) - 1
	for _, data := range all_data2 {
		if len(data) != expected {
			log.Fatal("(test 2) se esperaban ", len(Object1().Fields), " campos pero se obtuvo: ", expected)
		}
	}

	// fmt.Println("DATA OBTENIDA: ", all_data)
}
