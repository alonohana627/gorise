package models

import (
	"encoding/json"
	"errors"
	"log"
)

type UpdateContactByPhone struct {
	Name        *string `json:"name,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	PhoneNumber string  `json:"phone_number"`
	Address     *string `json:"address,omitempty"`
}

func (c *UpdateContactByPhone) Validate() error {
	if c.Name != nil && !isValidAlphanumeric(*c.Name) {
		return errors.New("name must contain only alphanumeric characters")
	}
	if c.LastName != nil && !isValidAlphanumeric(*c.LastName) {
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

func (c *UpdateContactByPhone) UnmarshalJSON(data []byte) error {
	type Alias UpdateContactByPhone
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
