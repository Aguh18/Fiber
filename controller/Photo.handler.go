package controller

import (
	// "fiber/database"
	// "fiber/models/entity"
	"fiber/models/request"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {

	photo := new(request.PhotoRequest)

	if err := ctx.BodyParser(photo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	validate := validator.New()

	errvalid := validate.Struct(photo)
	if errvalid != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errvalid.Error(),
		})
	}
	var filenamestring string

	// validation reuquired image
	filenames := ctx.Locals("filenames")
	if filenames == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "image is required",
		})
	}else{
		log.Print("ini dijalankan")
		filenamestring = fmt.Sprintf("%v", filenames)
		
	}
	log.Println("filename", filenamestring)

	// newphoto := entity.Photo{
	// 	Image:      filename.(string),
	// 	CategoryID: photo.CategotyId,
	// }

	// errcreatePhoto := database.DB.Create(&newphoto).Error
	// if errcreatePhoto != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Cannot create book",
	// 		"error":   errcreatePhoto,
	// 	})
	// }
	return ctx.JSON(fiber.Map{
		"massage": "successful",
		
	})

}
