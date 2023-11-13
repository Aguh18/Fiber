package migration

import (
	
	"fiber/database"
	"fiber/models/entity"
	
	
	"log"
)

func Migration()  {

	err := database.DB.AutoMigrate(&entity.User{}, entity.Book{})
	if err != nil {
		log.Println(err)
	}
	
}