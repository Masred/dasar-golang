package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Masred",
		MiddleName: "Ganteng",
		LastName:   "Banget",
		Hobbies:    []string{"Gaming", "Hiking"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Masred","MiddleName":"Ganteng","LastName":"Banget","Age":0,"Married":false,"Hobbies":["Gaming","Hiking"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Masred",
		Addresses: []Address{
			{
				Street:     "Jalan belum jadi",
				Country:    "Indo",
				PostalCode: "65346",
			},
			{
				Street:     "Jalan Sudah jadi",
				Country:    "Cina",
				PostalCode: "4423",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Masred","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan belum jadi","Country":"Indo","PostalCode":"65346"},{"Street":"Jalan Sudah jadi","Country":"Cina","PostalCode":"4423"}]}`
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
