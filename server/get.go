package server

import (
	"encoding/json"
	"gorise/db"
	"log"
	"net/http"
	"strconv"
)

func getOffset(offset string) int {
	offsetNum := int64(0)
	if offset != "" {
		var err error
		offsetNum, err = strconv.ParseInt(offset, 10, 32)
		if err != nil {
			offsetNum = 0
		}
		if offsetNum < 0 {
			offsetNum = 0
		}
	}

	return int(offsetNum)
}

// GetContacts handles retrieving a list of contacts from the phonebook.
//
// This function expects a GET request. If the request method is not GET, it returns a 405 Method Not Allowed response.
//
// The function retrieves the offset value from the query parameters to determine the starting point for fetching contacts.
//
// It calls db.GetContacts to fetch the list of contacts starting from the specified offset.
// If there is an error during the fetch, it returns a 500 Internal Server Error response.
//
// The list of contacts is then marshaled into JSON format and sent in the response.
//
// Example:
//
//	GET /contacts?offset=10
//
// Query Parameters:
//   - offset (optional): The starting point for fetching contacts. Default is 0.
//
// Responses:
//   - 200 OK: Successfully retrieved contacts list.
//   - 405 Method Not Allowed: Request method is not GET.
//   - 500 Internal Server Error: Error fetching contacts or marshaling response.
func GetContacts(w http.ResponseWriter, r *http.Request, client *db.DatabaseClient) {
	if r.Method != "GET" {
		errorResponse(w, r, 405, "Method is not allowed")
		return
	}

	offsetNum := getOffset(r.URL.Query().Get("offset"))

	contacts, err := client.GetContacts(offsetNum)
	if err != nil {
		log.Println(err)
		errorResponse(w, r, 500, "Internal Error")
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
