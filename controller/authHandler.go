package controller

import (
	"fiber/database"
	"fiber/models/entity"
	"fiber/models/request"
	"fiber/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(ctx *fiber.Ctx) error {
	Loginrequest := new(request.LoginRequest)

	if err := ctx.BodyParser(Loginrequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	validate := validator.New()
	errVakliate := validate.Struct(Loginrequest)
	if errVakliate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errVakliate.Error(),
		})
	}

	// cek email available
	var user entity.User
	err := database.DB.First(&user, "email = ?", Loginrequest.Email).Error
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email not found",
		})

	}

	// cek validation password
	if !utils.CheckHasedPassword(Loginrequest.Password, user.Password) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password not match",
		})
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errgenratetoken := utils.GenerateToken(&claims)

	if errgenratetoken != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cannot generate token",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})

}
