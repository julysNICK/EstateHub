package lead

import (
	"context"
	utils "estatehub-api/internal/utils"
	"strings"
)

type Service struct {
	repo LeadRepository
}

func NewService(repo LeadRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateLeadRequest) (LeadResponse, error) {
	lead := Lead{
		ID:         utils.GenerateID(),
		Name:       strings.TrimSpace(req.Name),
		Email:      strings.TrimSpace(req.Email),
		Phone:      strings.TrimSpace(req.Phone),
		PropertyID: strings.TrimSpace(req.PropertyID),
		AgencyID:   strings.TrimSpace(req.AgencyID),
		BrokerID:   strings.TrimSpace(req.BrokerID),
		CreatedAt:  utils.Now(),
		UpdatedAt:  utils.Now(),
	}

	createdLead, err := s.repo.Create(ctx, &lead)
	if err != nil {
		return LeadResponse{}, ErrFailedToCreateLead
	}

	return ToLeadResponse(&createdLead), nil
}

func (s *Service) GetById(ctx context.Context, id string) (LeadResponse, error) {
	lead, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return LeadResponse{}, ErrLeadNotFound
	}

	return ToLeadResponse(&lead), nil
}

func (s *Service) List(ctx context.Context) ([]LeadResponse, error) {
	leads, err := s.repo.List(ctx)
	if err != nil {
		return nil, ErrFailedToGetLeads
	}

	var leadResponses []LeadResponse
	for _, lead := range leads {
		leadResponses = append(leadResponses, ToLeadResponse(&lead))
	}

	return leadResponses, nil
}

func (s *Service) Update(ctx context.Context, id string, req UpdateLeadRequest) (LeadResponse, error) {
	existingLead, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return LeadResponse{}, ErrLeadNotFound
	}

	existingLead.Name = strings.TrimSpace(*req.Name)
	existingLead.Email = strings.TrimSpace(*req.Email)
	existingLead.Phone = strings.TrimSpace(*req.Phone)
	existingLead.PropertyID = strings.TrimSpace(*req.PropertyID)
	existingLead.UpdatedAt = utils.Now()

	updatedLead, err := s.repo.Update(ctx, &existingLead)
	if err != nil {
		return LeadResponse{}, ErrFailedToUpdateLead
	}
	return ToLeadResponse(&updatedLead), nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		return ErrFailedToDeleteLead
	}

	return nil
}
