package middleware

import (
	"context"
	"crm/internal/domain"
	"crm/internal/repository/postgres"
	"encoding/json"
	"net/http"
)

type AuthHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request) (int, error)
	Login(w http.ResponseWriter, r *http.Request) (int, error)
}

type AuthHandler struct {
	accountHandler postgres.AccountHandler
}

func NewAuthHandler(accountHandler postgres.AccountHandler) *AuthHandler {
	return &AuthHandler{accountHandler: accountHandler}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return http.StatusMethodNotAllowed, nil
	}

	accountRegister := &domain.RegisterAccount{}
	err := json.NewDecoder(r.Body).Decode(&accountRegister)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return http.StatusBadRequest, nil
	}

	err = h.accountHandler.Create(context.Background(), accountRegister)
	if err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, nil

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return http.StatusMethodNotAllowed, nil
	}

	accountLogin := &domain.AccountLoginWithEmail{}
	err := json.NewDecoder(r.Body).Decode(&accountLogin)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, nil
}
