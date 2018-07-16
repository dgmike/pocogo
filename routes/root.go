package routes

import (
	"encoding/json"
	"net/http"
)

// HandleGetRoot Fetchs a GET on /api/
func HandleGetRoot(w http.ResponseWriter, r *http.Request) {
	metadata := map[string]interface{}{
		"name":    "users api",
		"version": "1.0",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(metadata)
}
