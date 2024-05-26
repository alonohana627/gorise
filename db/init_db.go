package db

import "log"

func InitDB() error {
	_, err := InitClient()
	if err != nil {
		log.Fatal(err)
	}

	err = InitializeTables()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
