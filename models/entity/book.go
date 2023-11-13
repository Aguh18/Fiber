package entity

import (
	"time"
	"gorm.io/gorm"
)

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Author      string `gorm:"type:varchar(255)" json:"author"`
	Cover		string `gorm:"type:varchar(255)" json:"cover"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updatedat"`
	DeletedAt gorm.DeletedAt `json:"deletedat" gorm:"index"`

}