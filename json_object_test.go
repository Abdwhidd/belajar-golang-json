package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Abdul",
		MiddleName: "Wahid",
		LastName:   "Kahar",
		Age:        21,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
