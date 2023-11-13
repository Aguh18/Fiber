package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"index:unique"`
	Address    string `json:"address"`
	Password 	string `json:"password" `
	Phone     string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updatedat"`
	DeletedAt gorm.DeletedAt `json:"deletedat" gorm:"index"`
}
