package server

import (
	"encoding/json"
	"gorise/db"
	"gorise/models"
	"log"
	"net/http"
)

// AddContact handles the addition of a new contact to the phonebook.
//
// This function expects a POST request containing a JSON body with the contact details.
// If the request method is not POST, it returns a 405 Method Not Allowed response.
//
// The function decodes the JSON body into a models.Contact struct. If decoding fails or
// the body is missing required parts, it returns a 400 Bad Request response.
//
// If the contact is successfully inserted into the database, it returns a success message.
// If the contact already exists, it returns a message indicating the contact already exists.
//
// In case of an internal server error during insertion, it returns a 500 Internal Server Error response.
//
// Example:
//
//	POST /add_contact
//	{
//	  "name": "Alon",
//	  "last_name": "Ohana",
//	  "phone_number": "052-840-8722",
//	  "address": "HaDror 3 Gedera"
//	}
//
// Responses:
//   - 200 OK: Contact added successfully or already exists.
//   - 400 Bad Request: Invalid request body.
//   - 405 Method Not Allowed: Request method is not POST.
//   - 500 Internal Server Error: Database insertion error.
func AddContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorResponse(w, r, 405, "Method Not Allowed")
		return
	}

	d := json.NewDecoder(r.Body)
	contactBody := &models.Contact{}
	err := d.Decode(contactBody)
	if err != nil || contactBody == nil {
		log.Println(err)
		errorResponse(w, r, 400, "Bad Request - missing parts from the body")
		return
	}

	m, rowAffected, err := db.InsertContact(*contactBody)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 500, "Internal Server Error")
		return
	}

	if rowAffected == 0 {
		successBody := models.AlreadyExistJSON{
			Message: "Entry is Already Existed!",
			Contact: *m,
		}
		bytes, _ := json.Marshal(&successBody)
		w.Write(bytes)
		return
	}

	successBody := models.SuccessJSON{Message: "Inserted " + m.Name + " successfully"}
	bytes, _ := json.Marshal(&successBody)
	w.Write(bytes)

	log.Println("Added " + m.Name + " to the phonebook")
}
