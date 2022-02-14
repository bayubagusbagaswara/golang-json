package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {

	product := Product{
		Id:       "P001",
		Name:     "MacBook Pro 2021",
		ImageURL: "http://example.com/image.png",
	}

	b, _ := json.Marshal(product)
	fmt.Println(string(b))

}
func TestJSONTagDecode(t *testing.T) {

	jsonString := `{"id":"P001","name":"MacBook Pro 2021","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
