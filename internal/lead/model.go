package lead

import "time"

type Lead struct {
	ID         string
	Name       string
	Email      string
	Phone      string
	PropertyID string
	AgencyID   string
	BrokerID   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
