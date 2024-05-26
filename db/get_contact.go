package db

import (
	"context"
	"gorise/models"
	"log"
)

func (c *DatabaseClient) GetContactByName(name string) (*models.Contact, error) {
	var firstName string
	var lastName string
	var phoneNumber string
	var address *string

	err := c.dbClient.QueryRow(context.Background(), GetSpecificContactByNameQuery, name).Scan(&firstName, &lastName, &phoneNumber, &address)
	if err != nil {
		return nil, err
	}

	contact := new(models.Contact)
	contact.Name = firstName
	contact.LastName = lastName
	contact.PhoneNumber = phoneNumber
	contact.Address = address

	return contact, nil
}

// Not using cache here. Reason - pagination.

func (c *DatabaseClient) GetContacts(offset int) ([]*models.Contact, error) {
	contactList := make([]*models.Contact, 0)
	rows, err := c.dbClient.Query(context.Background(), GetContactByNameQuery, offset)
	if err != nil {
		log.Println(err)
		return contactList, err
	}

	var name, phoneNumber, lastName string
	var address *string

	for rows.Next() {
		err = rows.Scan(&name, &lastName, &phoneNumber, &address)
		if err != nil {
			log.Println(err)
			continue
		}

		contact := new(models.Contact)

		contact.Name = name
		contact.LastName = lastName
		contact.PhoneNumber = phoneNumber
		contact.Address = address

		contactList = append(contactList, contact)
	}

	return contactList, nil
}

// Not exposed outside, meant for caching only.

func (c *DatabaseClient) GetAllContacts() ([]*models.Contact, error) {
	contactList := make([]*models.Contact, 0)
	rows, err := c.dbClient.Query(context.Background(), GetAllContactsQuery)
	if err != nil {
		log.Println(err)
		return contactList, err
	}

	var name, phoneNumber, lastName string
	var address *string

	for rows.Next() {
		err = rows.Scan(&name, &lastName, &phoneNumber, &address)
		if err != nil {
			log.Println(err)
			continue
		}

		contact := new(models.Contact)

		contact.Name = name
		contact.LastName = lastName
		contact.PhoneNumber = phoneNumber
		contact.Address = address

		contactList = append(contactList, contact)
	}

	return contactList, nil
}
