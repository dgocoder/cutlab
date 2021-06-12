package domain

import (
	"time"

	"github.com/google/uuid"
)

type Resource struct {
	ID                  uuid.UUID          `json:"id"`
	LocationID          uuid.UUID          `json:"locationId"`
	CompanyID           uuid.UUID          `json:"companyId"`
	Name                string             `json:"name"`
	Email               string             `json:"email"`
	Description         *string            `json:"description"`
	ImageURL            *string            `json:"imageUrl"`
	DeletedStatus       bool               `json:"deletedStatus"`
	DeletedAt           *time.Time         `json:"deletedAt"`
	IgnoreBusinessHours bool               `json:"ignoreBusinessHours"`
	BioLink             *string            `json:"bioLink"`
	TimezoneID          string             `json:"timezoneId"`
	Phone               *Phone             `json:"phone"`
	Address             *Address           `json:"address"`
	Availability        Availability       `json:"availability"`
	CustomFields        *map[string]string `json:"customFields"`
}
type Phone struct {
	PhoneType        *string `json:"phoneType"`
	HomePhone        *string `json:"homePhone"`
	MobilePhone      *string `json:"mobilePhone"`
	BusinessPhone    *string `json:"businessPhone"`
	BusinessPhoneExt *string `json:"businessPhoneExt"`
}

func NewResource(name string, locationId uuid.UUID, companyId uuid.UUID, email string) Resource {
	return Resource{
		ID:         uuid.New(),
		Name:       name,
		LocationID: locationId,
		CompanyID:  companyId,
		Email:      email,
	}
}
