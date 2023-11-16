package controller

import (
	// "fiber/database"
	// "fiber/models/entity"
	"fiber/database"
	"fiber/models/entity"
	"fiber/models/request"
	"fiber/utils"

	// "fmt"
	// "log"

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
	// var filenamestring string

	// validation reuquired image
	filenames := ctx.Locals("filenames")
	if filenames == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		// log.Print("ini dijalankan")
		// filenamestring = fmt.Sprintf("%v", filenames)
		filenamedata := filenames.([]string)
		for _, filename := range filenamedata {
			newphoto := entity.Photo{
				Image:      filename,
				CategoryID: photo.CategotyId,
			}

			errcreatePhoto := database.DB.Create(&newphoto).Error
			if errcreatePhoto != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Cannot create book",
					"error":   errcreatePhoto,
				})
			}
		}

	}

	return ctx.JSON(fiber.Map{
		"massage": "successful",
	})

}

func PhotoHandlerDeleteById(ctx *fiber.Ctx) error {
	Photoid := ctx.Params("id")

	var photo entity.Photo

	err := database.DB.Debug().First(&photo, "id = ?", Photoid).Error
	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Photo not found",
			"error":   err,
		})
	}

	// Remove file Handler
	errdeletefile := utils.HandleRemoveFile(photo.Image)
	if errdeletefile != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot delete file",
			"error":   errdeletefile,
		})}

	errDelete := database.DB.Delete(&photo).Error
	if errDelete != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot delete photo",
			"error":   errDelete,
		})

	}

	return ctx.JSON(fiber.Map{
		"message": "Data succes delete",
	})

}
