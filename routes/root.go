package routes

import (
	"net/http"
	"time"

	"../models"
)

// HandleGetRoot Fetchs a GET on /api/
func HandleGetRoot(w http.ResponseWriter, r *http.Request) {
	defer SyncLogs(time.Now())

	metadata := map[string]interface{}{
		"name":    "users api",
		"version": "1.0",
	}

	models.SendJSON(http.StatusOK, metadata, w)
}
