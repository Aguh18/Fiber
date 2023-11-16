package route

import (
	"fiber/controller"
	"fiber/middleware"
	"fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", "public/asset")

	r.Post("/login", controller.Login)

	r.Get("/", middleware.Auth, controller.UserControllerRead)
	r.Get("/user", middleware.Auth, controller.UserHandlerGetAll)
	r.Get("/user/:id", middleware.Auth, controller.UserHandlerGetById)
	r.Post("/user", controller.UserHandlerCreate)
	r.Put("/user/:id", middleware.Auth, controller.UserHandlerUpdateByid)
	r.Put("/user/:id/update-email", middleware.Auth, controller.UserHandlerUpdateEmailByid)
	r.Delete("/user/:id", middleware.Auth, controller.UserHandlerDeleteByid)

	// book
	r.Post("/book", utils.HandleSingleRequest, controller.BookHandlerCreate)

	// multiple books
	r.Post("/gallery", utils.HandleMultipleRequest, controller.PhotoHandlerCreate)
	r.Delete("/gallery/:id", controller.PhotoHandlerDeleteById)

}
