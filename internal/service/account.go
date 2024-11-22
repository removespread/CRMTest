package service

import (
	"context"
	"errors"
	"time"

	"crm/internal/domain"
	"crm/internal/repository/handlers"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AccountInterface interface {
	RegisterAccount(ctx context.Context, account *domain.RegisterAccount) error
	LoginAccount(ctx context.Context, account *domain.AccountLoginWithEmail) error
}

type AccountService struct {
	accountRepo handlers.AccountHandler
	jwtSecret   string
}

func NewAccountService(accountRepo handlers.AccountHandler, jwtSecret string) *AccountService {
	return &AccountService{accountRepo: accountRepo, jwtSecret: jwtSecret}
}
func (s *AccountService) RegisterAccount(ctx context.Context, account *domain.RegisterAccount) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new Account instead of reassigning RegisterAccount
	account = &domain.RegisterAccount{
		Email:    account.Email,
		Password: string(passwordHash),
	}

	return s.accountRepo.Create(ctx, account)
}

func (s *AccountService) LoginAccount(ctx context.Context, account *domain.AccountLoginWithEmail) error {
	findAccount, err := s.accountRepo.GetByEmail(ctx, account.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findAccount.Password), []byte(account.Password))
	if err != nil {
		return errors.New("login failed / wrong email or password")
	}

	return nil
}

func (s *AccountService) GenerateToken(ctx context.Context, account *domain.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}
