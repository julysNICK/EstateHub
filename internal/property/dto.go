package property

import "time"

type CreatePropertyRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Address     string  `json:"address" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Size        float64 `json:"size" validate:"required"`
	Rooms       int     `json:"rooms" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	AgencyID    string  `json:"agency_id" validate:"required"`
	BrokerID    string  `json:"broker_id" validate:"required"`
}

type UpdatePropertyRequest struct {
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Address     *string  `json:"address,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Size        *float64 `json:"size,omitempty"`
	Rooms       *int     `json:"rooms,omitempty"`
	Type        *string  `json:"type,omitempty"`
	Status      *string  `json:"status,omitempty"`
	AgencyID    *string  `json:"agency_id,omitempty"`
	BrokerID    *string  `json:"broker_id,omitempty"`
}

type PropertyResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Price       float64   `json:"price"`
	Size        float64   `json:"size"`
	Rooms       int       `json:"rooms"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	AgencyID    string    `json:"agency_id"`
	BrokerID    string    `json:"broker_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListPropertyResponse struct {
	Properties []PropertyResponse `json:"properties"`
}

func (r *CreatePropertyRequest) Validate() error {
	if r.Title == "" {
		return ErrTitleRequired
	}
	if r.Description == "" {
		return ErrDescriptionRequired
	}
	if r.Address == "" {
		return ErrAddressRequired
	}
	if r.Price <= 0 {
		return ErrInvalidPrice
	}

	if r.Size <= 0 {
		return ErrInvalidSize
	}

	if r.Rooms <= 0 {
		return ErrInvalidRooms
	}

	if r.Type == "" {
		return ErrTypeRequired
	}

	if r.Status == "" {
		return ErrStatusRequired
	}
	if r.AgencyID == "" {
		return ErrInvalidAgencyID
	}

	if r.BrokerID == "" {
		return ErrInvalidBrokerID
	}
	return nil
}

func ToPropertyResponse(p *Property) *PropertyResponse {
	return &PropertyResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Address:     p.Address,
		Price:       p.Price,
		Size:        p.Size,
		Rooms:       p.Rooms,
		Type:        p.Type,
		Status:      p.Status,
		AgencyID:    p.AgencyID,
		BrokerID:    p.BrokerID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}

}

func ToPropertyResponses(properties []Property) []PropertyResponse {
	responses := make([]PropertyResponse, len(properties))

	for i, property := range properties {
		responses[i] = *ToPropertyResponse(&property)
	}

	return responses
}
