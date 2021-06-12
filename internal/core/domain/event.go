package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID                 uuid.UUID         `json:"id"`
	LocationID         uuid.UUID         `json:"locationId"`
	ServiceID          uuid.UUID         `json:"serviceId"`
	ResourceID         uuid.UUID         `json:"resourceId"`
	ResourceImageURL   *string           `json:"resourceImageUrl"`
	CustomerID         uuid.UUID         `json:"customerId"`
	CreatedAt          time.Time         `json:"createdAt"`
	StartAt            time.Time         `json:"startAt"`
	EndAt              time.Time         `json:"endAt"`
	Status             string            `json:"status"`
	ConfirmationNumber string            `json:"confirmationNumber"`
	BookedBy           *uuid.UUID        `json:"bookedBy"`
	Confirmed          bool              `json:"confirmed"`
	CustomerMessage    string            `json:"customerMessage"`
	Notes              string            `json:"notes"`
	LastModifiedOn     *time.Time        `json:"lastModifiedOn"`
	LastModifiedBy     *time.Time        `json:"lastModifiedBy"`
	StripeChargeID     string            `json:"stripeChargeId"`
	StripeRefundID     *string           `json:"stripeRefundId"`
	PaymentStatus      int               `json:"paymentStatus"`
	Type               string            `json:"type"`
	Reason             string            `json:"reason"`
	Repeats            bool              `json:"repeats"`
	Repeat             Repeat            `json:"repeat"`
	CustomFields       map[string]string `json:"customFields"`
}

type Repeat struct {
	Frequency string `json:"frequency"`
	Interval  int    `json:"interval"`
	Weekdays  string `json:"weekdays"`
	MonthDay  string `json:"monthDay"`
	MonthType string `json:"monthType"`
}

func NewEvent(locationId uuid.UUID, serviceId uuid.UUID, resourceId uuid.UUID, customerId uuid.UUID, startAt time.Time, endAt time.Time, eventType string) Event {
	return Event{
		ID:         uuid.New(),
		LocationID: locationId,
		ServiceID:  serviceId,
		ResourceID: resourceId,
		CustomerID: customerId,
		CreatedAt:  time.Now(),
		StartAt:    startAt,
		EndAt:      endAt,
		Type:       eventType,
	}
}
