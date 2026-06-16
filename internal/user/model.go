package user

import "time"

type User struct {
	ID           string
	Name         string
	Email        string
	Type         string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Agency struct {
	ID       string
	UserID   string
	AgencyID string
	Cnpj     string
}

type Broker struct {
	ID       string
	UserID   string
	BrokerID string
	Creci    string
	Cpf      string
}
