package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"type:varchar(36);primaryKey;"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Items    []Item `json:"items"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
