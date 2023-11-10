package route

import (
	"fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("/", controller.UserControllerRead)
	r.Get("/user", controller.UserHandlerGetAll)
	r.Get("/user/:id", controller.UserHandlerGetById)
	r.Post("/user", controller.UserHandlerCreate)
	r.Put("/user/:id", controller.UserHandlerUpdateByid)  
	r.Put("/user/:id/update-email", controller.UserHandlerUpdateEmailByid)  
	r.Delete("/user/:id", controller.UserHandlerDeleteByid)


}
