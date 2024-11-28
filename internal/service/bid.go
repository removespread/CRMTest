package service

import (
	"context"
	"crm/internal/domain"
	"crm/internal/repository/handlers"
	"time"

	"go.uber.org/zap"
)

type BidInterface interface {
	Create(ctx context.Context, bid *domain.Bid) error
	GetAll(ctx context.Context) (*[]domain.Bid, error)
	GetByID(ctx context.Context, id int64) (*domain.Bid, error)
	Update(ctx context.Context, bid *domain.Bid) error
	Delete(ctx context.Context, bid *domain.Bid) error
}

type BidService struct {
	bidHandler handlers.BidHandler
	logger     *zap.SugaredLogger
}

func NewBidService(bidHandler handlers.BidHandler, logger *zap.SugaredLogger) *BidService {
	return &BidService{bidHandler: bidHandler, logger: logger}
}

func (s *BidService) Create(ctx context.Context, bid *domain.Bid) error {
	bid.CreatedAt = time.Now()

	if bid.Amount <= 0 {
		s.logger.Errorf("Amount must be greater than 0")
	}

	if bid.Description == "" {
		s.logger.Errorf("Description is required")
	}

	bid = &domain.Bid{
		Description: bid.Description,
		Amount:      bid.Amount,
		CreatedAt:   bid.CreatedAt,
	}

	return s.bidHandler.Create(ctx, bid)
}

func (s *BidService) Update(ctx context.Context, bid *domain.Bid) error {
	findBid, err := s.bidHandler.GetByID(ctx, bid.ID)
	if err != nil {
		s.logger.Errorf("Error getting bid by id: %v", err)
	}

	if bid.Amount <= 0 || bid.Description == "" {
		s.logger.Errorf("Amount must be greater than 0 and description is required")
	}

	bid = &domain.Bid{
		Description: bid.Description,
		Amount:      bid.Amount,
		CreatedAt:   findBid.CreatedAt,
	}

	return s.bidHandler.Update(ctx, findBid)
}

func (s *BidService) Delete(ctx context.Context, bid *domain.Bid) error {
	findBid, err := s.bidHandler.GetByID(ctx, bid.ID)
	if err != nil {
		s.logger.Errorf("Error getting bid by id: %v", err)
	}

	return s.bidHandler.Delete(ctx, findBid)
}

func (s *BidService) GetAll(ctx context.Context) (*[]domain.Bid, error) {
	findBids, err := s.bidHandler.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("Error getting all bids: %v", err)
	}

	return findBids, nil
}

func (s *BidService) GetByID(ctx context.Context, id int64) (*domain.Bid, error) {
	findBid, err := s.bidHandler.GetByID(ctx, id)
	if err != nil {
		s.logger.Errorf("Error getting bid by id: %v", err)
	}

	return findBid, nil
}
