package main

import (
	"fmt"
	"gorise/db"
	"gorise/server"
	"log"
	"net/http"
)

func main() {
	client, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	routes := server.CreateRoutes(client)

	fmt.Println("Running the server on port 8080")
	http.ListenAndServe(":8080", routes)
}
