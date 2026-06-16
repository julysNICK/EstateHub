package property

import "errors"

var (
	ErrPropertyNotFound       = errors.New("property not found")
	ErrInvalidPropertyData    = errors.New("invalid property data")
	ErrFailedToCreateProperty = errors.New("failed to create property")
	ErrFailedToUpdateProperty = errors.New("failed to update property")
	ErrFailedToDeleteProperty = errors.New("failed to delete property")
	ErrFailedToListProperties = errors.New("failed to list properties")
	ErrFailedToGetProperty    = errors.New("failed to get property")
	ErrTitleRequired          = errors.New("title is required")
	ErrDescriptionRequired    = errors.New("description is required")
	ErrAddressRequired        = errors.New("address is required")
	ErrFailedToGetProperties  = errors.New("failed to get properties")
	ErrInvalidTitle           = errors.New("invalid title")
	ErrInvalidDescription     = errors.New("invalid description")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidPrice           = errors.New("invalid price")
	ErrInvalidSize            = errors.New("invalid size")
	ErrInvalidRooms           = errors.New("invalid rooms")
	ErrInvalidType            = errors.New("invalid type")
	ErrInvalidStatus          = errors.New("invalid status")
	ErrInvalidAgencyID        = errors.New("invalid agency ID")
	ErrInvalidBrokerID        = errors.New("invalid broker ID")
	ErrTypeRequired           = errors.New("type is required")
	ErrStatusRequired         = errors.New("status is required")
)
