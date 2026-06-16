package property

type IPropertyService interface {
	Create(req CreatePropertyRequest) (*PropertyResponse, error)
	List() (*ListPropertyResponse, error)
	GetByID(propertyId string) (*PropertyResponse, error)
	Update(propertyId string, req UpdatePropertyRequest) (*PropertyResponse, error)
	Delete(propertyId string) error
}

type PropertyService struct {
	repo IPropertyRepository
}

func NewPropertyService(repo IPropertyRepository) *PropertyService {
	return &PropertyService{
		repo: repo,
	}
}

func (s *PropertyService) Create(req CreatePropertyRequest) (*PropertyResponse, error) {
	return nil, nil
}

func (s *PropertyService) List() (*ListPropertyResponse, error) {
	return nil, nil
}

func (s *PropertyService) GetByID(propertyId string) (*PropertyResponse, error) {
	return nil, nil
}

func (s *PropertyService) Update(propertyId string, req UpdatePropertyRequest) (*PropertyResponse, error) {
	return nil, nil
}

func (s *PropertyService) Delete(propertyId string) error {
	return nil
}
