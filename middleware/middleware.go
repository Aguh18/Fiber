package middleware

import (
	"fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error  {

	token := ctx.Get("x-token")

	if token == "" {
		
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeTOken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		
	}

	role := claims["role"].(string)
	if  role != "admin" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "forbiden acces",
		})
	}
	
	
	ctx.Locals("user", claims)

	return ctx.Next()
}

func PermessionCreate(ctx *fiber.Ctx) error   {
	


	return ctx.Next()
}
