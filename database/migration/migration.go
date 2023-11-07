package migration

import (
	
	"fiber/database"
	"fiber/model/entity"
	
	"fmt"
	"log"
)

func Migration()  {

	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Migration has been processed")
	
	
}