package server

import (
	"encoding/json"
	"gorise/db"
	"gorise/models"
	"log"
	"net/http"
)

// DeleteContact handles the deletion of a contact from the phonebook.
//
// This function expects a DELETE request containing a JSON body with succession status.
// If the request method is not DELETE, it returns a 405 Method Not Allowed response.
//
// The function decodes the JSON body into a models.UpdateContactByPhone struct. If decoding fails or
// the body is missing required parts, it returns a 400 Bad Request response.
//
// If the contact is successfully deleted from the database, it returns a success message.
// If the contact is not found, it returns a message indicating that the contact was not found.
//
// In case of an error during deletion, it returns a 500 Internal Error.
//
// Example:
//
//	DELETE /delete_contact
//	{
//	  "phone_number": "052-840-8722"
//	}
//
// Responses:
//   - 200 OK: Contact deleted successfully or not found.
//   - 400 Bad Request: Invalid request body or deletion error.
//   - 405 Method Not Allowed: Request method is not DELETE.
//   - 500 Internal Error.
func DeleteContact(w http.ResponseWriter, r *http.Request, client *db.DatabaseClient) {
	if r.Method != "DELETE" {
		errorResponse(w, r, 405, "Method Not Allowed")
		return
	}

	d := json.NewDecoder(r.Body)
	phoneBody := &models.UpdateContactByPhone{}
	err := d.Decode(phoneBody)
	if err != nil || phoneBody == nil {
		log.Println(err)
		errorResponse(w, r, 400, "Bad Request - missing parts from the body")
		return
	}

	deleted, err := client.DeleteContact(phoneBody.PhoneNumber)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 500, "Internal Error")
		return
	}

	if deleted {
		successBody := models.SuccessJSON{Message: "deleted successfully"}
		bytes, _ := json.Marshal(&successBody)
		w.Write(bytes)
		log.Println("Deleted Successfully")
		return
	}

	successBody := models.SuccessJSON{Message: "not found"}
	bytes, _ := json.Marshal(&successBody)
	w.Write(bytes)
}
