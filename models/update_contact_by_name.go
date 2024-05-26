package models

import (
	"encoding/json"
	"errors"
	"log"
)

type UpdateContactByName struct {
	Name        string  `json:"name"`
	LastName    string  `json:"lastName"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Address     *string `json:"address,omitempty"`
}

func (c *UpdateContactByName) Validate() error {
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
	if c.PhoneNumber != nil && !isValidPhoneNumber(*c.PhoneNumber) {
		return errors.New("phone number must contain only digits and hyphens")
	}
	return nil
}

func (c *UpdateContactByName) UnmarshalJSON(data []byte) error {
	type Alias UpdateContactByName
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
