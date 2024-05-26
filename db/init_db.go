package db

import (
	"log"
)

func InitDB() (*DatabaseClient, error) {
	c, err := InitClient()
	if err != nil {
		log.Fatal(err)
	}

	err = c.InitializeTables()
	if err != nil {
		log.Fatal(err)
	}
	return c, err
}
