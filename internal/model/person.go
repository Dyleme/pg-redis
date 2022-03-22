package model

import (
	"fmt"
)

type Person struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Phone   string  `json:"phone"`
}

func ValidatePerson(person *Person) error {
	if person.Name == "" {
		return fmt.Errorf("empty name")
	}

	if person.Phone == "" {
		return fmt.Errorf("empty phone")
	}

	if err := ValidateAddress(&person.Address); err != nil {
		return fmt.Errorf("address: %w", err)
	}

	return nil
}

type Address struct {
	Country    string `json:"country"`
	City       string `json:"city"`
	Street     string `json:"street"`
	House      string `json:"house"`
	Apartments string `json:"apartments"`
}

func ValidateAddress(address *Address) error {
	if address.Country == "" {
		return fmt.Errorf("empty country")
	}

	if address.City == "" {
		return fmt.Errorf("empty city")
	}

	if address.Street == "" {
		return fmt.Errorf("empty street")
	}

	if address.House == "" {
		return fmt.Errorf("empty house")
	}

	return nil
}
