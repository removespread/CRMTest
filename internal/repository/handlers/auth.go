package handlers

import (
	"context"
	"crm/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandlerInterface interface {
	Register(ctx context.Context, account *domain.RegisterAccount) error
	Login(ctx context.Context, account *domain.AccountLoginWithEmail) error
}

type AuthHandler struct {
	accountHandler *AccountHandler
}

func NewAuthHandler(accountHandler *AccountHandler) *AuthHandler {
	return &AuthHandler{accountHandler: accountHandler}
}

func (h *AuthHandler) Register(ctx context.Context, account *domain.RegisterAccount) error {
	account = &domain.RegisterAccount{
		Email:    account.Email,
		Password: account.Password,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	account.Password = string(passwordHash)

	return h.accountHandler.Create(ctx, account)

}

func (h *AuthHandler) Login(ctx context.Context, account *domain.AccountLoginWithEmail) error {
	findAccount, err := h.accountHandler.GetByEmail(ctx, account.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findAccount.Password), []byte(account.Password))
	if err != nil {
		return err
	}

	return h.accountHandler.LoginAccount(ctx, findAccount)
}
