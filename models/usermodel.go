package models

// UserModel is a User manager
type UserModel struct{}

// User is the user entity
type User struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Users is a list of users
type Users []User

// FindByID finds a user by ID
func (u UserModel) FindByID(id uint64) *User {
	for _, item := range UsersList {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

// FindAll select all users
func (u UserModel) FindAll() *Users {
	return &UsersList
}

// UsersList List of all users
var UsersList = Users{
	{
		1,
		"Rubens",
		"rubens.ribeiro@catho.com",
	},
	{
		2,
		"João",
		"joao.armando@catho.com",
	},
}
