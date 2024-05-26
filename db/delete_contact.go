package db

import (
	"context"
)

func (c *DatabaseClient) DeleteContact(phoneNumber string) (bool, error) {
	dbResp, err := c.dbClient.Exec(context.Background(), DeleteContactByPhone, phoneNumber)
	if err != nil {
		return false, err
	}

	return dbResp.Delete(), err
}
