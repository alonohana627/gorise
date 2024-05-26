package db

import (
	"context"
)

func DeleteContact(phoneNumber string) (bool, error) {
	dbResp, err := Client.Exec(context.Background(), DeleteContactByPhone, phoneNumber)
	if err != nil {
		return false, err
	}

	return dbResp.Delete(), err
}
