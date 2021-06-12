package domain

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	CreatedAt     time.Time  `json:"createdAt"`
	DeletedStatus bool       `json:"deleted"`
	DeletedAt     *time.Time `json:"deletedAt"`
	AddressLine1  *string    `json:"addressLine1"`
	AddressLine2  *string    `json:"addressLine2"`
	City          *string    `json:"city"`
	State         *string    `json:"state"`
	PostalCode    *string    `json:"postalCode"`
	Country       *string    `json:"country"`
	Phone         *string    `json:"phone"`
	Email         string     `json:"email"`
	Website       *string    `json:"website"`
	TimezoneID    *string    `json:"timezoneId"`
}

func NewCompany(name string, email string) Company {
	return Company{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		DeletedAt: nil,
	}
}
