package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HTTPError response error
type HTTPError struct {
	Message string `json:"message"`
}

// SendJSON Sends a JSON response throw the response writer
func SendJSON(status int, body interface{}, response http.ResponseWriter) {
	jsonResponse, err := json.Marshal(body)
	if err != nil {
		response.Header().Set("Content-Type", "application/vnd.error+json; charset=UTF-8")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(HTTPError{Message: "Internal Server Error"})

		log.Printf("Erro: %s", err)
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(status)
	fmt.Fprint(response, string(jsonResponse))
}
