package service

import (
	"context"
	"crm/internal/domain"
	"crm/internal/repository/handlers"

	"go.uber.org/zap"
)

type PartnerService struct {
	partnerHandler *handlers.PartnerHandler
	sliceContacts  *handlers.ContactHandler
	logger         *zap.SugaredLogger
}

type PartnerInterface interface {
	Create(ctx context.Context, partner *domain.Partner) error
	Update(ctx context.Context, partner *domain.Partner) error
	Delete(ctx context.Context, partner *domain.Partner) error
	GetAll(ctx context.Context) (*[]domain.Partner, error)
	GetByID(ctx context.Context, id int64) (*domain.Partner, error)
}

func NewPartnerService(partnerHandler *handlers.PartnerHandler, sliceContacts *handlers.ContactHandler, logger *zap.SugaredLogger) *PartnerService {
	return &PartnerService{partnerHandler: partnerHandler, sliceContacts: sliceContacts, logger: logger}
}

func (s *PartnerService) Create(ctx context.Context, partner *domain.Partner) error {
	if partner.Name == "" {
		s.logger.Errorf("Name is required")
	}

	if len(partner.Contacts) == 0 {
		s.logger.Errorf("Contacts are required")
	}

	var sliceContacts []domain.Contact

	for _, contactID := range partner.Contacts {
		contact, err := s.sliceContacts.GetByID(ctx, contactID.ID)
		if err != nil {
			s.logger.Errorf("Error getting contact by id: %v", err)
		}

		sliceContacts = append(sliceContacts, *contact)
	}

	partner = &domain.Partner{
		Name:        partner.Name,
		Description: partner.Description,
		Contacts:    sliceContacts,
	}

	return s.partnerHandler.Create(ctx, partner)
}

func (s *PartnerService) Update(ctx context.Context, partner *domain.Partner) error {
	_, err := s.partnerHandler.GetByID(ctx, partner.ID)
	if err != nil {
		return err
	}

	if partner.Name == "" {
		s.logger.Errorf("Name is required")
	}

	var sliceContacts []domain.Contact
	for _, contactID := range partner.Contacts {
		contact, err := s.sliceContacts.GetByID(ctx, contactID.ID)
		if err != nil {
			s.logger.Errorf("Error getting contact by id: %v", err)
		}

		sliceContacts = append(sliceContacts, *contact)
	}

	partner = &domain.Partner{
		Name:        partner.Name,
		Description: partner.Description,
		Contacts:    sliceContacts,
	}

	if len(partner.Contacts) == 0 {
		s.logger.Errorf("Contacts are required")
	}

	return s.partnerHandler.Update(ctx, partner)
}

func (s *PartnerService) Delete(ctx context.Context, partner *domain.Partner) error {
	findPartner, err := s.partnerHandler.GetByID(ctx, partner.ID)
	if err != nil {
		return err
	}

	return s.partnerHandler.Delete(ctx, findPartner)
}

func (s *PartnerService) GetAll(ctx context.Context) (*[]domain.Partner, error) {
	return s.partnerHandler.GetAll(ctx)
}

func (s *PartnerService) GetByID(ctx context.Context, id int64) (*domain.Partner, error) {
	return s.partnerHandler.GetByID(ctx, id)
}
