package domain

import "github.com/google/uuid"

type Service struct {
	ID                    uuid.UUID `json:"id"`
	CompanyID             uuid.UUID `json:"companyId"`
	LocationID            uuid.UUID `json:"locationId"`
	Name                  string    `json:"name"`
	Description           *string   `json:"description"`
	Type                  *string   `json:"type"`
	ImageURL              *string   `json:"imageUrl"`
	Duration              int       `json:"duration"`
	Padding               int       `json:"padding"`
	FeeAmount             int       `json:"feeAmount"`
	FeeTaxable            bool      `json:"feeTaxable"`
	CancellationFeeAmount int       `json:"cancellationFeeAmount"`
	CancellationCutOff    int       `json:"cancellationCutOff"`
	NonRefundable         bool      `json:"nonRefundable"`
}

func NewService(companyId uuid.UUID, locationId uuid.UUID, name string, description *string, duration int, feeAmount int, cancellationFeeAmount int, CancellationCutOff int) Service {
	return Service{
		ID:                    uuid.New(),
		CompanyID:             companyId,
		LocationID:            locationId,
		Name:                  name,
		Description:           description,
		Duration:              duration,
		FeeAmount:             feeAmount,
		CancellationFeeAmount: cancellationFeeAmount,
		CancellationCutOff:    CancellationCutOff,
	}
}
