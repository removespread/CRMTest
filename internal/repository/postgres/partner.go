package postgres

import (
	"context"
	"crm/internal/domain"

	"gorm.io/gorm"
)

type PartnerHandler struct {
	db *gorm.DB
}

func NewPartnerHandler(db *gorm.DB) *PartnerHandler {
	return &PartnerHandler{db: db}
}

func (h *PartnerHandler) Create(ctx context.Context, partner *domain.Partner) error {
	return h.db.Create(&partner).Error
}

func (h *PartnerHandler) Update(ctx context.Context, partner *domain.Partner) error {
	return h.db.Save(&partner).Error
}

func (h *PartnerHandler) Delete(ctx context.Context, partner *domain.Partner) error {
	return h.db.Delete(&partner).Error
}

func (h *PartnerHandler) GetAll(ctx context.Context) (*[]domain.Partner, error) {
	var partners []domain.Partner
	err := h.db.Find(&partners).Error
	if err != nil {
		return nil, err
	}
	return &partners, nil
}

func (h *PartnerHandler) GetByID(ctx context.Context, id int64) (*domain.Partner, error) {
	var partner domain.Partner
	err := h.db.First(&partner, id).Error
	if err != nil {
		return nil, err
	}
	return &partner, nil
}
