package routes

import (
	"net/http"
	"strconv"
	"time"

	"../models"
	"github.com/gorilla/mux"
)

var userModel models.UserModel = models.UserModel{}

// HandleGetOneUser Fetches a user
func HandleGetOneUser(w http.ResponseWriter, r *http.Request) {
	defer SyncLogs(time.Now())

	params := mux.Vars(r)
	paramID, _ := strconv.Atoi(params["id"])
	userID := uint64(paramID)

	user := userModel.FindByID(userID)
	if user == nil {
		body := models.HTTPError{Message: "Not found"}
		models.SendJSON(http.StatusNotFound, body, w)
		return
	}

	models.SendJSON(http.StatusOK, *user, w)
}

// HandleGetUsers Fetches all users
func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	defer SyncLogs(time.Now())
	users := userModel.FindAll()
	models.SendJSON(http.StatusOK, users, w)
}
