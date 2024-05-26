package server

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Path: "+r.URL.Path, " RemoteAddr: "+r.RemoteAddr, " Method: "+r.Method)
		next(w, r)
	}
}

func CreateRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/create-contact", LoggerMiddleware(AddContact))
	mux.HandleFunc("/api/v1/read-contacts", LoggerMiddleware(GetContacts))
	mux.HandleFunc("/api/v1/update-contact-by-name", LoggerMiddleware(EditContactByName))
	mux.HandleFunc("/api/v1/update-contact-by-phone", LoggerMiddleware(EditContactByPhone))
	mux.HandleFunc("/api/v1/delete-contact", LoggerMiddleware(DeleteContact))
	mux.HandleFunc("/api/v1/search-contacts", LoggerMiddleware(SearchContact))

	return mux
}
