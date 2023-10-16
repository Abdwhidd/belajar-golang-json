package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"Abdul","MiddleName":"Wahid","LastName":"Kahar","Age":21}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
}
