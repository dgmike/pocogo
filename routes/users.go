package routes

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id    int64
	Name  string
	Email string
}

type Users []User

// func (u *Users) findByID(id int) (*User, error) {
// 	for i := 0; i < len(*u); i++ {
// 		if *u[i].Id == id {
// 			return &*u[i], nil
// 		}
// 	}

// 	return nil, error("invalid ID")
// }

var UsersList Users = Users{
	User{
		Id:    1,
		Name:  "Aurino Victor",
		Email: "aurino.victor@catho.com",
	},
	User{
		Id:    2,
		Name:  "Michael Armando",
		Email: "michael.armando@catho.com",
	},
}

// HandleGetOneUser Fetches a user
func HandleGetOneUser(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// userID, _ := strconv.Atoi(vars["id"])

	user := UsersList[0]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(user)
}
