package db

import (
	"context"
	"log"
)

func InitializeTables() error {
	_, err := InitClient()
	if err != nil {
		return err
	}

	if isTableExists() {
		return err
	}

	log.Println("`contact` table does not exist in `phonebook` DB! Creating contact table...")
	_, err = Client.Query(context.Background(), createContactTableQuery)
	return err
}

func isTableExists() bool {
	var exists bool
	err := Client.QueryRow(context.Background(), isTableExistsQuery, "contact").Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}
