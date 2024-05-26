package server

import (
	"encoding/json"
	"gorise/models"
	"net/http"
)

func errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, body string) {
	errorBody := models.ErrorJSON{
		Message:  body,
		ErrorNum: statusCode,
	}
	bytes, _ := json.Marshal(&errorBody)

	w.WriteHeader(statusCode)
	w.Write(bytes)
}
