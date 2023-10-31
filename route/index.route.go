package route

import (
	"fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("/user", controller.UserControllerRead)
}
