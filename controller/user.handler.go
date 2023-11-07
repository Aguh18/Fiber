package controller

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/entity/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserControllerRead(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"hello": "world",
	})
	
}
func GetAllHandler(ctx *fiber.Ctx) error  {
	var users []entity.User

	err :=  database.DB.Find(&users).Error

	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(users)
	
	
}
func CreateUser (ctx *fiber.Ctx) error{
	
	user := new(request.User)

	if err:= ctx.BodyParser(user); err != nil {
		log.Println(err)
		
	}
	NewUser := entity.User{
		Name : user.Name,
		Email : user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateuser := database.DB.Create(&NewUser).Error
	if errCreateuser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Create user failed",
	})
	
}
return ctx.Status(200).JSON(fiber.Map{
	"message": "Create user success",})

}