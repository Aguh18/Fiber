package migration

import (
	"fiber/database"
	"fiber/models/entity"
	"fmt"

	"log"
)

func Migration()  {

	err := database.DB.AutoMigrate(&entity.User{}, entity.Book{}, entity.Category{}, entity.Photo{})
	if err != nil {
		log.Println(err)
		fmt.Print("Error migrate database")
	}
	
}