package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          string  `gorm:"type:varchar(36);primaryKey;"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Sold        bool    `json:"sold"`
	UserID      string  `json:"user_id"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.NewString()
	return
}
