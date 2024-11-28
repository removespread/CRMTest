package service

import (
	"context"
	"crm/internal/domain"
	"crm/internal/repository/handlers"

	"go.uber.org/zap"
)

type ContactService struct {
	contactHandler handlers.ContactHandler
	logger         *zap.SugaredLogger
}

type ContactServiceInterface interface {
	Create(ctx context.Context, contact *domain.Contact) error
	Update(ctx context.Context, contact *domain.Contact) error
	Delete(ctx context.Context, contact *domain.Contact) error
}

func NewContactService(contactHandler handlers.ContactHandler, logger *zap.SugaredLogger) *ContactService {
	return &ContactService{contactHandler: contactHandler, logger: logger}
}

func (s *ContactService) Create(ctx context.Context, contact *domain.Contact, logger *zap.SugaredLogger) error {
	if contact.Name == "" {
		s.logger.Errorf("Name is required")
	}

	if contact.Phone == "" {
		s.logger.Errorf("Phone is required")
	}

	contact = &domain.Contact{
		Name:  contact.Name,
		Phone: contact.Phone,
	}

	return s.contactHandler.Create(ctx, contact)
}

func (s *ContactService) Update(ctx context.Context, contact *domain.Contact, logger *zap.SugaredLogger) error {
	findContact, err := s.contactHandler.GetByID(ctx, contact.ID)
	if err != nil {
		s.logger.Errorf("Error getting contact by id: %v", err)
	}

	if contact.Name == "" {
		s.logger.Errorf("Name is required")
	}

	if contact.Phone == "" {
		s.logger.Errorf("Phone is required")
	}

	contact = &domain.Contact{
		Name:  contact.Name,
		Phone: contact.Phone,
	}

	return s.contactHandler.Update(ctx, findContact)
}

func (s *ContactService) Delete(ctx context.Context, contact *domain.Contact) error {
	findContact, err := s.contactHandler.GetByID(ctx, contact.ID)
	if err != nil {
		s.logger.Errorf("Error getting contact by id: %v", err)
	}

	return s.contactHandler.Delete(ctx, findContact)
}

func (s *ContactService) GetAll(ctx context.Context, logger *zap.SugaredLogger) (*[]domain.Contact, error) {
	return s.contactHandler.GetAll(ctx)
}

func (s *ContactService) GetByID(ctx context.Context, id int64) (*domain.Contact, error) {
	contact, err := s.contactHandler.GetByID(ctx, id)
	if err != nil {
		s.logger.Errorf("Error getting contact by id: %v", err)
	}

	return contact, nil
}
