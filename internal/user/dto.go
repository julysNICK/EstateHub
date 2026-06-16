package user

import (
	"estatehub-api/internal/platform/shared"
	"time"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Type     string `json:"type" validate:"required,oneof=agency broker"`
	Password string `json:"password" validate:"required,min=6"`
	AgencyID string `json:"agency_id,omitempty"`
	BrokerID string `json:"broker_id,omitempty"`
	Cnpj     string `json:"cnpj,omitempty"`
	Creci    string `json:"creci,omitempty"`
	Cpf      string `json:"cpf,omitempty"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Type     *string `json:"type,omitempty" validate:"omitempty,oneof=agency broker"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=6"`
	AgencyID *string `json:"agency_id,omitempty"`
	BrokerID *string `json:"broker_id,omitempty"`
	Cnpj     *string `json:"cnpj,omitempty"`
	Creci    *string `json:"creci,omitempty"`
	Cpf      *string `json:"cpf,omitempty"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AgencyResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	AgencyID string `json:"agency_id"`
	Cnpj     string `json:"cnpj"`
}

type BrokerResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	BrokerID string `json:"broker_id"`
	Creci    string `json:"creci"`
	Cpf      string `json:"cpf"`
}

type ListUserRequest struct {
	Name        string
	Email       string
	PageRequest shared.PageRequest
}

type ListAgencyResponse struct {
	Agencies      []AgencyResponse `json:"agencies"`
	NextPageToken string           `json:"next_page_token,omitempty"`
}

type ListBrokerResponse struct {
	Brokers       []BrokerResponse `json:"brokers"`
	NextPageToken string           `json:"next_page_token,omitempty"`
}
type ListUserResponse struct {
	Users         []UserResponse `json:"users"`
	NextPageToken string         `json:"next_page_token,omitempty"`
}

func (r *CreateUserRequest) Validate() error {
	// Implement validation logic here, e.g., using a validation library
	return nil
}

func (r *UpdateUserRequest) Validate() error {
	// Implement validation logic here, e.g., using a validation library
	return nil
}

func ToUserResponse(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Type:      user.Type,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []User) []UserResponse {
	responses := make([]UserResponse, len(users))

	for i, user := range users {
		responses[i] = ToUserResponse(&user)
	}

	return responses
}

func ToAgencyResponse(agency *Agency) AgencyResponse {
	return AgencyResponse{
		ID:       agency.ID,
		UserID:   agency.UserID,
		AgencyID: agency.AgencyID,
		Cnpj:     agency.Cnpj,
	}
}

func ToAgencyResponses(agencies []Agency) []AgencyResponse {
	responses := make([]AgencyResponse, len(agencies))

	for i, agency := range agencies {
		responses[i] = ToAgencyResponse(&agency)
	}

	return responses
}

func ToBrokerResponse(broker *Broker) BrokerResponse {
	return BrokerResponse{
		ID:       broker.ID,
		UserID:   broker.UserID,
		BrokerID: broker.BrokerID,
		Creci:    broker.Creci,
		Cpf:      broker.Cpf,
	}
}

func ToBrokerResponses(brokers []Broker) []BrokerResponse {
	responses := make([]BrokerResponse, len(brokers))

	for i, broker := range brokers {
		responses[i] = ToBrokerResponse(&broker)
	}

	return responses
}

func ToListAgencyResponse(agencies []Agency, nextPageToken string) ListAgencyResponse {
	return ListAgencyResponse{
		Agencies:      ToAgencyResponses(agencies),
		NextPageToken: nextPageToken,
	}
}

func ToListBrokerResponse(brokers []Broker, nextPageToken string) ListBrokerResponse {
	return ListBrokerResponse{
		Brokers:       ToBrokerResponses(brokers),
		NextPageToken: nextPageToken,
	}
}

func ToListUserResponse(users []User, nextPageToken string) ListUserResponse {
	return ListUserResponse{
		Users:         ToUserResponses(users),
		NextPageToken: nextPageToken,
	}
}
