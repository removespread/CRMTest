package domain

import "time"

type Bid struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	Amount      int       `json:"amount" gorm:"notnull"`
	CreatedAt   time.Time `json:"created_at" gorm:"notnull"`
}

type CreateBidInput struct {
	Description string `json:"description" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
}
