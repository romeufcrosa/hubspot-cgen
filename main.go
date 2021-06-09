package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v6"
)

type Contact struct {
	FirstName     string `fake:"{firstname}"`
	LastName      string `fake:"{lastname}"`
	EmailAddress  string `fake:"{email}"`
	PhoneNumber   string `fake:"{phone}"`
	StreetAddress string `fake:"{street}"`
	City          string `fake:"{city}"`
	State         string `fake:"{state}"`
	PostalCode    string `fake:"{zip}"`
}

func main() {
	recordSize := 100
	gofakeit.Seed(0)

	f, err := os.Create("hubspot.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write([]string{"First Name", "Last Name", "Email Address", "Phone Number", "Street Address", "City", "State", "Postal Code"}); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < recordSize; i++ {
		var c Contact
		gofakeit.Struct(&c)

		if err := w.Write([]string{c.FirstName, c.LastName, c.EmailAddress, c.PhoneNumber, c.StreetAddress, c.City, c.State, c.PostalCode}); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
