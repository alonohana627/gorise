package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

// Global client.
// Pool for multiple requests at once.
//If it is not a pgxpool - every action will be blocked and multiple requests to the DB are not possible.

var Client *pgxpool.Pool

type connString struct {
	username string
	password string
	url      string
	port     string
	dbName   string
}

func composeConnString(c connString) string {
	return "postgres://" + c.username + ":" + c.password + "@" + c.url + ":" + c.port + "/" + c.dbName
}

func InitClient() (*pgxpool.Pool, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_URL")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	//username := "root"
	//password := "root"
	//url := "localhost"
	//port := "5432"
	//dbName := "phonebook"

	c := connString{username, password, url, port, dbName}
	connectionString := composeConnString(c)

	var err error
	Client, err = pgxpool.New(context.Background(), connectionString)

	return Client, err
}
