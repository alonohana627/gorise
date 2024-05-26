package models

import (
	"encoding/json"
	"errors"
	"log"
)

type Contact struct {
	Name        string  `json:"name"`
	LastName    string  `json:"lastName"`
	PhoneNumber string  `json:"phone_number"`
	Address     *string `json:"address,omitempty"`
}

func (c *Contact) Validate() error {
	if !isValidAlphanumeric(c.Name) {
		return errors.New("name must contain only alphanumeric characters")
	}
	if !isValidAlphanumeric(c.LastName) {
		return errors.New("last name must contain only alphanumeric characters")
	}
	if c.Address != nil && !isValidAlphanumeric(*c.Address) {
		log.Println(*c.Address)
		return errors.New("address must contain only alphanumeric characters")
	}
	if !isValidPhoneNumber(c.PhoneNumber) {
		return errors.New("phone number must contain only digits and hyphens")
	}
	return nil
}

func (c *Contact) UnmarshalJSON(data []byte) error {
	type Alias Contact
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return c.Validate()
}

type SearchContactModel struct {
	Name        *string `json:"name,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}
