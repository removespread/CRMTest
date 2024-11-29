package postgres

import (
	"context"
	"crm/internal/domain"

	"gorm.io/gorm"
)

type BidHandler struct {
	db *gorm.DB
}

func NewBidHandler(db *gorm.DB) *BidHandler {
	return &BidHandler{db: db}
}

func (h *BidHandler) Create(ctx context.Context, bid *domain.Bid) error {
	return h.db.Create(&bid).Error
}

func (h *BidHandler) Update(ctx context.Context, bid *domain.Bid) error {
	return h.db.Save(&bid).Error
}

func (h *BidHandler) Delete(ctx context.Context, bid *domain.Bid) error {
	return h.db.Delete(&bid).Error
}

func (h *BidHandler) GetAll(ctx context.Context) (*[]domain.Bid, error) {
	var bids []domain.Bid
	err := h.db.Find(&bids).Error
	if err != nil {
		return nil, err
	}
	return &bids, nil
}

func (h *BidHandler) GetByID(ctx context.Context, id int64) (*domain.Bid, error) {
	var bid domain.Bid
	err := h.db.First(&bid, id).Error
	if err != nil {
		return nil, err
	}
	return &bid, nil
}
