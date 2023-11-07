package controller

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"

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


func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error": err,
		})
	}
	validate := validator.New()
	
	
	errvalid := validate.Struct(user)
	if errvalid != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errvalid.Error(),
		})
}

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateuser := database.DB.Create(&newUser).Error
	if errCreateuser != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot create user",
			"error": errCreateuser,
	})
}
return ctx.JSON(fiber.Map{
	"massaged": "successful",
	"user": newUser,
})
}