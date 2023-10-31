package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Adress    string `json:"adress"`
	Phone     string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updatedat"`
	DeletedAt gorm.DeletedAt `json:"deletedat" gorm:"index"`
}
