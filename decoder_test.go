package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	// jadi ini konversi langsung dari reader

	reader, _ := os.Open("Customer.json")

	// bikin json Decoder nya
	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	// decode, mirip Unmarshal
	decoder.Decode(customer)

	fmt.Println(customer)
}
