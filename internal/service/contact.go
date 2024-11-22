package service

import (
	"context"
	"crm/internal/domain"
	"crm/internal/repository/handlers"
	"errors"
)

type ContactService struct {
	contactHandler handlers.ContactHandler
}

type ContactServiceInterface interface {
	Create(ctx context.Context, contact *domain.Contact) error
	Update(ctx context.Context, contact *domain.Contact) error
	Delete(ctx context.Context, contact *domain.Contact) error
}

func NewContactService(contactHandler handlers.ContactHandler) *ContactService {
	return &ContactService{contactHandler: contactHandler}
}

func (s *ContactService) Create(ctx context.Context, contact *domain.Contact) error {
	if contact.Name == "" {
		return errors.New("name is required")
	}

	if contact.Phone == "" {
		return errors.New("phone is required")
	}

	contact = &domain.Contact{
		Name:  contact.Name,
		Phone: contact.Phone,
	}

	return s.contactHandler.Create(ctx, contact)
}

func (s *ContactService) Update(ctx context.Context, contact *domain.Contact) error {
	findContact, err := s.contactHandler.GetByID(ctx, contact.ID)
	if err != nil {
		return err
	}

	if contact.Name == "" {
		return errors.New("name is required")
	}

	if contact.Phone == "" {
		return errors.New("phone is required")
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
		return err
	}

	return s.contactHandler.Delete(ctx, findContact)
}

func (s *ContactService) GetAll(ctx context.Context) (*[]domain.Contact, error) {
	return s.contactHandler.GetAll(ctx)
}

func (s *ContactService) GetByID(ctx context.Context, id int64) (*domain.Contact, error) {
	contact, err := s.contactHandler.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
