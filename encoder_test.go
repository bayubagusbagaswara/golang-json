package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")

	// bikin encoder
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Bayu",
		MiddleName: "Bagus",
		LastName:   "Bagaswara",
	}

	// kita encode data customernya, dalam arti kita lansung jadikan json data customernya
	encoder.Encode(customer)

	fmt.Println(customer)

}
