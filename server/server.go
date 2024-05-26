package server

import (
	"gorise/db"
	"log"
	"net/http"
)

type NextFunc func(http.ResponseWriter, *http.Request, *db.DatabaseClient)

func LoggerMiddleware(next NextFunc, client *db.DatabaseClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Path: "+r.URL.Path, " RemoteAddr: "+r.RemoteAddr, " Method: "+r.Method)
		next(w, r, client)
	}
}

func CreateRoutes(client *db.DatabaseClient) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/create-contact", LoggerMiddleware(AddContact, client))
	mux.HandleFunc("/api/v1/read-contacts", LoggerMiddleware(GetContacts, client))
	mux.HandleFunc("/api/v1/update-contact-by-name", LoggerMiddleware(EditContactByName, client))
	mux.HandleFunc("/api/v1/update-contact-by-phone", LoggerMiddleware(EditContactByPhone, client))
	mux.HandleFunc("/api/v1/delete-contact", LoggerMiddleware(DeleteContact, client))
	mux.HandleFunc("/api/v1/search-contacts", LoggerMiddleware(SearchContact, client))

	return mux
}
