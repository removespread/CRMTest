package domain

type Contact struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"notnull"`
	Phone       string `json:"phone" gorm:"notnull"`
	Description string `json:"description"`
}

type CreateContact struct {
	Name        string `json:"name" binding:"required" gorm:"notnull"`
	Phone       string `json:"phone" binding:"required" gorm:"notnull"`
	Description string `json:"description"`
}

type UpdateContact struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
