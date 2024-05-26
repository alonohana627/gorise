package db

import (
	"context"
	"gorise/models"
)

func InsertContact(newContact models.Contact) (*models.Contact, int64, error) {
	name := newContact.Name
	c, err := GetContactByName(name)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, 0, err
	}

	if c != nil {
		return c, 0, nil
	}

	dbResp, err := Client.Exec(context.Background(), InsertContactQuery, newContact.Name, newContact.LastName, newContact.PhoneNumber, newContact.Address)
	if err != nil {
		return nil, 0, err
	}

	// When inserted - insert to the cache map as well
	return &newContact, dbResp.RowsAffected(), err
}
