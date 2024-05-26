package server

import (
	"encoding/json"
	"gorise/db"
	"gorise/models"
	"log"
	"net/http"
)

// SearchContact handles searching for contacts in the phonebook based on search criteria.
//
// This function expects a POST request containing a JSON body with the search criteria.
// If the request method is not POST, it returns a 405 Method Not Allowed response.
//
// The function decodes the JSON body into a models.SearchContactModel struct. If decoding fails or
// the body is missing required parts, it returns a 400 Bad Request response.
//
// It calls db.SearchContact to search for contacts that match the provided criteria.
// If there is an error during the search, it returns a 500 Internal Server Error response.
//
// The list of matching contacts is then marshaled into JSON format and sent in the response.
//
// Example:
//
//	POST /search_contact
//	{
//	  "name": "Alon",
//	  "last_name": "Ohana",
//	  "phone_number": "052-840-8722",
//	  "address": "HaDror 3 Gedera"
//	}
//
// Responses:
//   - 200 OK: Successfully retrieved list of matching contacts.
//   - 400 Bad Request: Invalid request body.
//   - 405 Method Not Allowed: Request method is not POST.
//   - 500 Internal Server Error: Error searching contacts or marshaling response.
func SearchContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorResponse(w, r, 405, "Method is not allowed")
		return
	}

	d := json.NewDecoder(r.Body)
	searchBody := &models.SearchContactModel{}
	err := d.Decode(searchBody)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 400, "Bad Request - missing parts from the body")
		return
	}

	contacts, err := db.SearchContact(*searchBody)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 500, "Internal Server Error")
		return
	}

	bytes, err := json.Marshal(&contacts)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 500, "Internal Error")
		return
	}
	w.Write(bytes)
}
