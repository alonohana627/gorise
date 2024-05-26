package db

import (
	"context"
	"log"
)

func (c *DatabaseClient) InitializeTables() error {
	if c.isTableExists() {
		return nil
	}

	log.Println("`contact` table does not exist in `phonebook` DB! Creating contact table...")
	_, err := c.dbClient.Query(context.Background(), createContactTableQuery)
	return err
}

func (c *DatabaseClient) isTableExists() bool {
	var exists bool
	err := c.dbClient.QueryRow(context.Background(), isTableExistsQuery, "contact").Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}
