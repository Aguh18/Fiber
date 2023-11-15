package entity

import (
	"time"

	"gorm.io/gorm"
)



type Photo struct {

	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Image string `gorm:"type:varchar(255)" json:"image"`
	CategoryID uint64 `gorm:"type:int" json:"categoryid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updatedat"`
	DeletedAt gorm.DeletedAt `json:"deletedat" gorm:"index"`



}