package property

import "time"

type Property struct {
	ID          string
	Title       string
	Description string
	Address     string
	Price       float64
	Size        float64
	Rooms       int
	Type        string
	Status      string
	AgencyID    string
	BrokerID    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
