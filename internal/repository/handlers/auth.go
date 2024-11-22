package handlers

import (
	"context"
	"crm/internal/domain"
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

}

func (h *AuthHandler) Login(ctx context.Context, account *domain.AccountLoginWithEmail) error {

}
