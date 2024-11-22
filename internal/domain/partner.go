package domain

type Partner struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"notnull"`
	Description string    `json:"description" gorm:"notnull"`
	Contacts    []Contact `json:"contacts_ids" gorm:"foreignKey:PartnerID"`
}

type CreatePartner struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Contacts    []Contact `json:"contacts_ids"`
}

type LoginPartner struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Contacts    []Contact `json:"contacts_ids"`
}
