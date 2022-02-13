package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// bikin struct untuk data Customer
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestJSONObject(t *testing.T) {

	// bikin data struct Customer nya
	customer := Customer{
		FirstName:  "Bayu",
		MiddleName: "Bagus",
		LastName:   "Bagaswara",
	}

	// encode dari tipe data customer menjadi data JSON
	bytes, _ := json.Marshal(customer)

	// kita konversi dari bytes menjadi string
	fmt.Println(string(bytes))
}
