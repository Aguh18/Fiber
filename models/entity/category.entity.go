package entity

import (
	
	"time"

	"gorm.io/gorm"
)



type Category struct {

	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Photos []Photo `json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updatedat"`
	DeletedAt gorm.DeletedAt `json:"deletedat" gorm:"index"`



}