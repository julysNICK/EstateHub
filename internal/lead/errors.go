package lead

import "errors"

var (
	ErrLeadNotFound        = errors.New("lead not found")
	ErrInvalidLeadData     = errors.New("invalid lead data")
	ErrFailedToCreateLead  = errors.New("failed to create lead")
	ErrFailedToUpdateLead  = errors.New("failed to update lead")
	ErrFailedToDeleteLead  = errors.New("failed to delete lead")
	ErrFailedToListLeads   = errors.New("failed to list leads")
	ErrFailedToGetLead     = errors.New("failed to get lead")
	ErrTitleRequired       = errors.New("title is required")
	ErrDescriptionRequired = errors.New("description is required")
	ErrFailedToGetLeads    = errors.New("failed to get leads")
	ErrInvalidName         = errors.New("invalid name")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrInvalidPhone        = errors.New("invalid phone")
	ErrInvalidPropertyID   = errors.New("invalid property ID")
	ErrInvalidAgencyID     = errors.New("invalid agency ID")
	ErrInvalidBrokerID     = errors.New("invalid broker ID")
)
