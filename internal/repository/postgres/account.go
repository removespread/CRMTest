package postgres

import (
	"context"
	"crm/internal/domain"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountHandler struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

type AccountHandlerInterface interface {
	Create(ctx context.Context, account *domain.RegisterAccount) error
	Update(ctx context.Context, account *domain.Account) error
	Delete(ctx context.Context, account *domain.Account) error
	GetAll(ctx context.Context) (*[]domain.Account, error)
	GetByID(ctx context.Context, id int64) (*domain.Account, error)
	GetByEmail(ctx context.Context, email string) (*domain.Account, error)
	LoginAccount(ctx context.Context, account *domain.Account) error
}

func NewAccountHandler(db *gorm.DB, logger *zap.SugaredLogger) *AccountHandler {
	return &AccountHandler{db: db, logger: logger}
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
		h.logger.Errorf("Error getting all accounts: %v", err)
	}
	return &accounts, nil
}

func (h *AccountHandler) GetByID(ctx context.Context, id int64) (*domain.Account, error) {
	var account domain.Account
	err := h.db.First(&account, id).Error
	if err != nil {
		h.logger.Errorf("Error getting account by id: %v", err)
	}
	return &account, nil
}

func (h *AccountHandler) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	var account domain.Account
	err := h.db.Where("email = ?", email).First(&account).Error
	if err != nil {
		h.logger.Errorf("Error getting account by email: %v", err)
	}
	return &account, nil
}

func (h *AccountHandler) LoginAccount(ctx context.Context, account *domain.Account) error {
	return h.db.Save(&account).Error
}
