package route

import (
	"fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("/", controller.UserControllerRead)
}

func  UserGetAll(r *fiber.App)  {
	r.Get("/user", controller.UserHandlerGetAll)
}

func  UserPost(r *fiber.App)  {
	r.Post("/user", controller.UserHandlerCreate)
}

func UserGetByid(r *fiber.App)  {
	r.Get("/user/:id", controller.UserHandlerGetById)
	
}

func UserUpdateById(r *fiber.App) {
	r.Put("/user/:id", controller.UserHandlerUpdateByid)  
	
}
