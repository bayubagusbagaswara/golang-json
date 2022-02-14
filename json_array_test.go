package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	// bikin data struct
	customer := Customer{
		FirstName:  "Bayu",
		MiddleName: "Bagus",
		LastName:   "Bagaswara",
		Hobbies:    []string{"Game", "Reading", "Coding"},
	}

	// encode menjadi JSON
	bytes, _ := json.Marshal(customer)

	// konversi bytes ke string
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {

	// bikin jsonString
	jsonString := `{"FirstName":"Bayu","MiddleName":"Bagus","LastName":"Bagaswara","Age":0,"Married":false,"Hobbies":["Game","Reading","Coding"]}`

	// konversi string ke byte
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	// decode dari byte menjadi data struct
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Bayu",
		// bikin addresses, dimana didalam addresses kita bikin lagi address yang banyak (lebih dari satu)
		// jadi bentuknya adalah Array of Object
		Addresses: []Address{
			{
				Street:     "Delima Street",
				Country:    "Indonesia",
				PostalCode: "60111",
			},
			{
				Street:     "Pahlawan Street",
				Country:    "Indonesia",
				PostalCode: "64551",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {

	jsonString := `{"FirstName":"Bayu","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Delima Street","Country":"Indonesia","PostalCode":"60111"},{"Street":"Pahlawan Street","Country":"Indonesia","PostalCode":"64551"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)

}
func TestOnlyJSONArrayComplexDecode(t *testing.T) {

	// JSON yang isinya adalah Array of String
	jsonString := `[{"Street":"Delima Street","Country":"Indonesia","PostalCode":"60111"},{"Street":"Pahlawan Street","Country":"Indonesia","PostalCode":"64551"}]`

	jsonBytes := []byte(jsonString)

	// kita ingin ubah menjadi Addresses
	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Delima Street",
			Country:    "Indonesia",
			PostalCode: "60111",
		},
		{
			Street:     "Pahlawan Street",
			Country:    "Indonesia",
			PostalCode: "64551",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes)) // keluarnya langsung json array
}
