package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorise/models"
	"log"
	"os"
)

type DatabaseClient struct {
	dbClient *pgxpool.Pool
}

type DatabaseEndpoints interface {
	InitializeTables() error
	InsertContact(models.Contact) (*models.Contact, int64, error)
	GetContactByName(string) (*models.Contact, error)
	GetContacts(int) ([]*models.Contact, error)
	GetAllContacts() ([]*models.Contact, error)
	UpdateContactByName(models.UpdateContactByName) error
	UpdateContactByPhone(models.UpdateContactByPhone) error
	DeleteContact(string) (bool, error)
}

func (c *DatabaseClient) Close() {
	c.dbClient.Close()
}

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

func InitClient() (*DatabaseClient, error) {
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

	dbClient, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Panic(err)
	}

	client := DatabaseClient{
		dbClient: dbClient,
	}

	return &client, err
}
