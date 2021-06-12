package domain

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID                uuid.UUID         `json:"id"`
	Email             string            `json:"email"`
	Lastname          string            `json:"lastname"`
	Firstname         string            `json:"firstname"`
	Gender            *string           `json:"gender"`
	Birthdate         *time.Time        `json:"birthdate"`
	LastVisitDate     *time.Time        `json:"lastVisitDate"`
	CreatedAt         time.Time         `json:"createdAt"`
	VerificationDate  *time.Time        `json:"verificationDate"`
	Disabled          bool              `json:"disabled"`
	DeletedStatus     bool              `json:"deletedStatus"`
	DeletedAt         *time.Time        `json:"deletedAt"`
	PaymentCustomerId *string           `json:"paymentCustomerId"`
	Contact           Phone             `json:"contact"`
	Address           Address           `json:"address"`
	CustomFields      map[string]string `json:"customFields"`
}

func NewCustomer(email string, lastName string, firstName string) Customer {
	return Customer{
		ID:        uuid.New(),
		Email:     email,
		Lastname:  lastName,
		Firstname: firstName,
		CreatedAt: time.Now().UTC(),
	}
}
