package handlers

import (
	"context"
	"crm/internal/domain"

	"gorm.io/gorm"
)

type AccountHandler struct {
	db *gorm.DB
}

func NewAccountHandler(db *gorm.DB) *AccountHandler {
	return &AccountHandler{db: db}
}

func (h *AccountHandler) Create(ctx context.Context, account *domain.RegisterAccount) error {
	return h.db.Create(&account).Error
}

func (h *AccountHandler) Update(ctx context.Context, account *domain.Account) error {
	return h.db.Save(&account).Error
}

func (h *AccountHandler) Delete(ctx context.Context, account *domain.Account) error {
	return h.db.Delete(&account).Error
}

func (h *AccountHandler) GetAll(ctx context.Context) (*[]domain.Account, error) {
	var accounts []domain.Account
	err := h.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return &accounts, nil
}

func (h *AccountHandler) GetByID(ctx context.Context, id int64) (*domain.Account, error) {
	var account domain.Account
	err := h.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (h *AccountHandler) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	var account domain.Account
	err := h.db.Where("email = ?", email).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
