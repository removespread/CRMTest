package domain

type Account struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique, notnull"`
	Password string `json:"password" gorm:"notnull"`
}

type RegisterAccount struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AccountLoginWithEmail struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AccountTokenResponse struct {
	Token string `json:"token"`
}
