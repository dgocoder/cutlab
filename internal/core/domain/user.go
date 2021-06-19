package domain

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Phone *Phone    `json:"phone"`
}

func NewUser(name string, email string, phone *Phone) User {
	return User{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
		Phone: phone,
	}
}
