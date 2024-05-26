package db

import (
	"context"
	"gorise/models"
)

func (c *DatabaseClient) InsertContact(newContact models.Contact) (*models.Contact, int64, error) {
	name := newContact.Name
	contact, err := c.GetContactByName(name)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, 0, err
	}

	if contact != nil {
		return contact, 0, nil
	}

	dbResp, err := c.dbClient.Exec(context.Background(), InsertContactQuery, newContact.Name, newContact.LastName, newContact.PhoneNumber, newContact.Address)
	if err != nil {
		return nil, 0, err
	}

	// When inserted - insert to the cache map as well
	return &newContact, dbResp.RowsAffected(), err
}
