package handlers

import (
	"context"
	"crm/internal/domain"

	"gorm.io/gorm"
)

type ContactHandler struct {
	db *gorm.DB
}

func NewContactHandler(db *gorm.DB) *ContactHandler {
	return &ContactHandler{db: db}
}

func (h *ContactHandler) Create(ctx context.Context, contact *domain.Contact) error {
	return h.db.Create(&contact).Error
}

func (h *ContactHandler) Update(ctx context.Context, contact *domain.Contact) error {
	return h.db.Save(&contact).Error
}

func (h *ContactHandler) Delete(ctx context.Context, contact *domain.Contact) error {
	return h.db.Delete(&contact).Error
}

func (h *ContactHandler) GetAll(ctx context.Context) (*[]domain.Contact, error) {
	var contacts []domain.Contact
	err := h.db.Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return &contacts, nil
}

func (h *ContactHandler) GetByID(ctx context.Context, id int64) (*domain.Contact, error) {
	var contact domain.Contact
	err := h.db.First(&contact, id).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}
