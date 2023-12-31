package response

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID        		uint   `json:"id" gorm:"primaryKey"`
	Name     		string `json:"name"`
	Email     		string `json:"email"`
	Address    		string `json:"address"`
	Password string `json:"-" gorm:"column:password"`
	Phone     		string `json:"phone"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time	`json:"updatedat"`
	DeletedAt 		gorm.DeletedAt `json:"deletedat" gorm:"index, column:deleted_at"`
}
