package controller

import (
	"fiber/models/entity"
	"fiber/models/request"
	"fmt"
	"log"

	"fiber/database"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {

	book := new(request.BookCreateRequest)

	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	validate := validator.New()

	errvalid := validate.Struct(book)
	if errvalid != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errvalid.Error(),
		})
	}

	file, err := ctx.FormFile("cover")
	if err != nil {
		log.Println("error:", err)
	}

	var filename string
	if file != nil {
		filename = file.Filename

		errsave := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
	if errsave != nil {
		log.Println("error:", errsave)
	}

	} else{
		log.Println("error:", "Nothing file to updload")
	}


	
	newbook := entity.Book{
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Cover:       filename,
	}

	errCreatebook := database.DB.Create(&newbook).Error
	if errCreatebook != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot create book",
			"error":   errCreatebook,
		})
	}
	return ctx.JSON(fiber.Map{
		"massaged": "successful",
		"book":     newbook,
	})
}
