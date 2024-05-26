package main

import (
	"fmt"
	"gorise/db"
	"gorise/server"
	"log"
	"net/http"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Client.Close()

	routes := server.CreateRoutes()

	fmt.Println("Running the server on port 8080")
	http.ListenAndServe(":8080", routes)
}
