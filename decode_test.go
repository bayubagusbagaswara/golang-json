package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	// bikin dulu jsonString, ini adalah data dengan format JSON
	jsonString := `{
		"FirstName" : "Bayu",
		"MiddleName" : "Bagus",
		"LastName" : "Bagaswara"
	}`

	// bikin json bytenya, untuk menyimpan data JSON diatas
	jsonBytes := []byte(jsonString)

	// bikin Customer kosong, bisa bikin pointer
	// nantinya secara otomatis JSON akan diparsing menurut field di Customer
	customer := &Customer{}

	// lakukan konversi menggunakan Unmarshal
	// balikannya adalah data error
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
