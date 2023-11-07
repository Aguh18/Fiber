package route

import (
	"fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("/", controller.UserControllerRead)
}

func  UserGetAll(r *fiber.App)  {
	r.Get("/user", controller.GetAllHandler)
}

func  UserPost(r *fiber.App)  {
	r.Post("/user", controller.UserHandlerCreate)
}
