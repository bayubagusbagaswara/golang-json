package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// parameternya adalah interface kosong
func logJson(data interface{}) {
	// lakukan encode dari sebuah tipe data menjadi JSON
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	// bytes adalah byte slice
	fmt.Println(string(bytes))
}

// testing
func TestEncode(t *testing.T) {

	// parameternya tipe data string
	logJson("Bayu")
	// parameternya tipe data integer
	logJson(12)
	// parameternya tipe data boolean
	logJson(true)
	// parameternya tipe data string slice
	logJson([]string{"Bayu", "Bagus", "Bagaswara"})

}
