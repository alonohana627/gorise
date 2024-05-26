package server

import (
	"encoding/json"
	"gorise/db"
	"gorise/models"
	"log"
	"net/http"
)

// EditContactByPhone handles the editing of a contact's details based on the phone number.
//
// This function expects a PATCH request containing a JSON body with the updated contact details and the phone number.
// If the request method is not PATCH, it returns a 405 Method Not Allowed response.
//
// The function decodes the JSON body into a models.UpdateContactByPhone struct. If decoding fails or
// the body is missing required parts, it returns a 400 Bad Request response.
//
// If the contact is successfully updated in the database, it returns a success message indicating the update.
// If there is an internal server error during the update, it returns a 500 Internal Server Error response.
//
// Example:
//
//	PATCH /edit_contact
//	{
//	  "phone_number": "052-840-8722",
//	  "name": "Alon",
//	  "last_name": "Ohana",
//	  "address": "HaDror 3 Gedera"
//	}
//
// Responses:
//   - 200 OK: Contact updated successfully.
//   - 400 Bad Request: Invalid request body.
//   - 405 Method Not Allowed: Request method is not PATCH.
//   - 500 Internal Server Error: Database update error.
func EditContactByPhone(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		errorResponse(w, r, 405, "Method Not Allowed")
		return
	}

	d := json.NewDecoder(r.Body)
	updateBody := &models.UpdateContactByPhone{}
	err := d.Decode(updateBody)
	if err != nil || updateBody == nil {
		log.Println(err)
		errorResponse(w, r, 400, "Bad Request - missing parts from the body")
		return
	}

	err = db.UpdateContactByPhone(*updateBody)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 500, "Internal Server Error")
		return
	}

	successBody := models.SuccessJSON{Message: "If entry exist - updated successfully"}
	bytes, _ := json.Marshal(&successBody)
	w.Write(bytes)
}

// EditContactByName handles the editing of a contact's details based on the name.
//
// This function expects a PATCH request containing a JSON body with the updated contact details and the name.
// If the request method is not PATCH, it returns a 405 Method Not Allowed response.
//
// The function decodes the JSON body into a models.UpdateContactByName struct. If decoding fails or
// the body is missing required parts, it returns a 400 Bad Request response.
//
// If the contact is successfully updated in the database, it returns a success message indicating the update.
// If there is an internal server error during the update, it returns a 500 Internal Server Error response.
//
// Example:
//
//	PATCH /edit_contact_by_name
//	{
//	  "name": "Alon",
//	  "last_name": "Ohana",
//	  "phone_number": "0528408722",
//	  "address": "Hador 3, Gedera"
//	}
//
// Responses:
//   - 200 OK: Contact updated successfully.
//   - 400 Bad Request: Invalid request body.
//   - 405 Method Not Allowed: Request method is not PATCH.
//   - 500 Internal Server Error: Database update error.
func EditContactByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		errorResponse(w, r, 405, "Method Not Allowed")
		return
	}

	d := json.NewDecoder(r.Body)
	updateBody := &models.UpdateContactByName{}
	err := d.Decode(updateBody)
	if err != nil || updateBody == nil {
		log.Println(err)
		errorResponse(w, r, 400, "Bad Request - missing parts from the body")
		return
	}

	err = db.UpdateContactByName(*updateBody)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 500, "Internal Server Error")
		return
	}

	successBody := models.SuccessJSON{Message: "If entry exist - updated successfully"}
	bytes, _ := json.Marshal(&successBody)
	w.Write(bytes)
}
