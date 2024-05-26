package db

import (
	"context"
	"fmt"
	"gorise/models"
	"log"
)

func SearchContact(searchModel models.SearchContactModel) ([]*models.Contact, error) {
	var firstName string
	var lastNameQuery string
	var phone string
	var searchName string

	contactList := make([]*models.Contact, 0)

	if searchModel.Name != nil {
		firstName = *searchModel.Name
	}
	if searchModel.LastName != nil {
		lastNameQuery = *searchModel.LastName
	}
	if firstName != "" || lastNameQuery != "" {
		searchName = "%" + firstName + lastNameQuery + "%"
	}

	if searchModel.PhoneNumber != nil {
		phone = "%" + *searchModel.PhoneNumber + "%"
	}

	rows, err := Client.Query(context.Background(), SearchContactByNameQuery, searchName, phone)
	if err != nil {
		log.Println(err)
		return contactList, err
	}

	var name, lastName, phoneNumber string
	var address *string
	for rows.Next() {
		err = rows.Scan(&name, &lastName, &phoneNumber, &address)
		if err != nil {
			fmt.Println(err)
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
