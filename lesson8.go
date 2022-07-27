package main

import (
	"encoding/json"
	"fmt"
	"lesson81/true"
	"os"
)

type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   struct {
		StreetAddress string `json:"streetAddress"`
		City          string `json:"city"`
		PostalCode    int    `json:"postalCode"`
	} `json:"address"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

func main() {

	data, err := os.ReadFile("contact.json")
	if err != nil {
		panic(err)
	}

	var cont Contact
	if err = json.Unmarshal(data, &cont); err != nil {
		panic(err)
	}
	fmt.Println(cont)

	true.GetConfiguration()
}
