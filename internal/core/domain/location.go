package domain

import "github.com/google/uuid"

type Location struct {
	ID               uuid.UUID           `json:"id"`
	CompanyID        uuid.UUID           `json:"companyId"`
	Name             *string             `json:"name"`
	Phone            *string             `json:"phone"`
	Email            *string             `json:"email"`
	Website          *string             `json:"website"`
	ImageURL         *string             `json:"imageUrl"`
	TimezoneID       *string             `json:"timezoneId"`
	BusinessHours    Availability        `json:"businessHours"`
	BusinessHolidays *[]BusinessHolidays `json:"businessHolidays"`
}

type Address struct {
	AddressLine1 string  `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	PostalCode   string  `json:"postalCode"`
}

type Mon struct {
	IsOpen    bool `json:"isOpen"`
	StartTime int  `json:"startTime"`
	EndTime   int  `json:"endTime"`
	Is24Hours bool `json:"is24Hours"`
}
type Tue struct {
	IsOpen    bool `json:"isOpen"`
	StartTime int  `json:"startTime"`
	EndTime   int  `json:"endTime"`
	Is24Hours bool `json:"is24Hours"`
}
type Wed struct {
	IsOpen    bool `json:"isOpen"`
	StartTime int  `json:"startTime"`
	EndTime   int  `json:"endTime"`
	Is24Hours bool `json:"is24Hours"`
}
type Thu struct {
	IsOpen    bool `json:"isOpen"`
	StartTime int  `json:"startTime"`
	EndTime   int  `json:"endTime"`
	Is24Hours bool `json:"is24Hours"`
}
type Fri struct {
	IsOpen    bool `json:"isOpen"`
	StartTime int  `json:"startTime"`
	EndTime   int  `json:"endTime"`
	Is24Hours bool `json:"is24Hours"`
}
type Sat struct {
	IsOpen    bool `json:"isOpen"`
	StartTime int  `json:"startTime"`
	EndTime   int  `json:"endTime"`
	Is24Hours bool `json:"is24Hours"`
}
type Sun struct {
	IsOpen    bool `json:"isOpen"`
	StartTime int  `json:"startTime"`
	EndTime   int  `json:"endTime"`
	Is24Hours bool `json:"is24Hours"`
}
type BusinessHolidays struct {
	ID             uuid.UUID `json:"id"`
	HolidayName    string    `json:"holidayName"`
	BusinessClosed bool      `json:"businessClosed"`
}
type Availability struct {
	Mon Mon `json:"mon"`
	Tue Tue `json:"tue"`
	Wed Wed `json:"wed"`
	Thu Thu `json:"thu"`
	Fri Fri `json:"fri"`
	Sat Sat `json:"sat"`
	Sun Sun `json:"sun"`
}

func NewLocation(companyId uuid.UUID, availability Availability, name *string) Location {
	return Location{
		ID:            uuid.New(),
		CompanyID:     companyId,
		Name:          name,
		BusinessHours: availability,
	}
}
