package db

import (
	"context"
	"gorise/models"
)

func (c *DatabaseClient) UpdateContactByName(uModel models.UpdateContactByName) error {
	if uModel.PhoneNumber != nil {
		_, err := c.dbClient.Exec(context.Background(), UpdateContactPhoneByNameQuery, uModel.Name, uModel.LastName, *uModel.PhoneNumber)
		if err != nil {
			return err
		}
	}

	if uModel.Address != nil {
		_, err := c.dbClient.Exec(context.Background(), UpdateContactAddressByNameQuery, uModel.Name, uModel.LastName, *uModel.Address)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *DatabaseClient) UpdateContactByPhone(uModel models.UpdateContactByPhone) error {
	if uModel.Name != nil {
		_, err := c.dbClient.Exec(context.Background(), UpdateContactNameByPhoneQuery, uModel.PhoneNumber, *uModel.Name)
		if err != nil {
			return err
		}
	}

	if uModel.LastName != nil {
		_, err := c.dbClient.Exec(context.Background(), UpdateContactLastNameByPhoneQuery, uModel.PhoneNumber, *uModel.LastName)
		if err != nil {
			return err
		}
	}

	if uModel.Address != nil {
		_, err := c.dbClient.Exec(context.Background(), UpdateContactAddressByPhoneQuery, uModel.PhoneNumber, *uModel.Address)
		if err != nil {
			return err
		}
	}

	return nil
}
