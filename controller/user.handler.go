package controller

import (
	"fiber/database"
	"fiber/model/entity"
	"fiber/model/request"
	"fiber/model/response"
	"log"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

func UserControllerRead(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"hello": "world",
	})
	
}
func UserHandlerGetAll(ctx *fiber.Ctx) error  {
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


func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	
	var user entity.User

	err := database.DB.First(&user, "id =?",userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find user",
			"error": err,
		})
	}
	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"user": userResponse,
	})


}


func UserHandlerUpdateByid(ctx *fiber.Ctx) error {
	userUpdate := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userUpdate); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error": err,
		})
	}

	var user entity.User
	userId := ctx.Params("id")
	// cek availabel
	err := database.DB.First(&user, "id =?",userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find user",
			"error": err,
		})
		
	}

	// update user
	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}
	user.Address = userUpdate.Address
	user.Phone = userUpdate.Phone
	user.Email = userUpdate.Email

	err = database.DB.Save(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot update user",
			"error": err,
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"updateUser": user,

})
}