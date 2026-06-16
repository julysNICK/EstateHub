package lead

import "time"

type CreateLeadRequest struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required"`
	PropertyID string `json:"property_id" validate:"required"`
	AgencyID   string `json:"agency_id" validate:"required"`
	BrokerID   string `json:"broker_id" validate:"required"`
}

type UpdateLeadRequest struct {
	Name       *string `json:"name,omitempty"`
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	PropertyID *string `json:"property_id,omitempty"`
	AgencyID   *string `json:"agency_id,omitempty"`
	BrokerID   *string `json:"broker_id,omitempty"`
}

type LeadResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	PropertyID string    `json:"property_id"`
	AgencyID   string    `json:"agency_id"`
	BrokerID   string    `json:"broker_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ListLeadResponse struct {
	Leads []LeadResponse `json:"leads"`
}

func (r *CreateLeadRequest) Validate() error {
	if r.Name == "" {
		return ErrInvalidName
	}
	if r.Email == "" {
		return ErrInvalidEmail
	}

	if r.Phone == "" {
		return ErrInvalidPhone
	}

	if r.PropertyID == "" {
		return ErrInvalidPropertyID
	}

	if r.AgencyID == "" {
		return ErrInvalidAgencyID
	}
	if r.BrokerID == "" {
		return ErrInvalidBrokerID
	}

	return nil
}

func (r *UpdateLeadRequest) Validate() error {
	if r.Name != nil && *r.Name == "" {
		return ErrInvalidName
	}

	if r.Email != nil && *r.Email == "" {
		return ErrInvalidEmail
	}
	if r.Phone != nil && *r.Phone == "" {
		return ErrInvalidPhone
	}
	if r.PropertyID != nil && *r.PropertyID == "" {
		return ErrInvalidPropertyID
	}
	if r.AgencyID != nil && *r.AgencyID == "" {
		return ErrInvalidAgencyID
	}
	if r.BrokerID != nil && *r.BrokerID == "" {
		return ErrInvalidBrokerID
	}

	return nil
}
func ToLeadResponse(lead *Lead) LeadResponse {
	return LeadResponse{
		ID:         lead.ID,
		Name:       lead.Name,
		Email:      lead.Email,
		Phone:      lead.Phone,
		PropertyID: lead.PropertyID,
		AgencyID:   lead.AgencyID,
		BrokerID:   lead.BrokerID,
		CreatedAt:  lead.CreatedAt,
		UpdatedAt:  lead.UpdatedAt,
	}
}

func ToLeadResponses(leads []Lead) []LeadResponse {
	responses := make([]LeadResponse, len(leads))

	for i, lead := range leads {
		responses[i] = ToLeadResponse(&lead)
	}

	return responses
}
