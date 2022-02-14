package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {

	jsonString := `{"id":"P001", "name":"Mac Book Pro 2021", "price":"25000000"}`

	jsonBytes := []byte(jsonString)

	// bikin map nya dulu
	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {

	product := map[string]interface{}{
		"id":    "P001",
		"name":  "MacBook Pro 2021",
		"price": 20000000,
	}

	b, _ := json.Marshal(product)

	fmt.Println(string(b)) // hasilnya jadi JSON
}
