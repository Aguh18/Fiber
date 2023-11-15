package utils

import (
	"fmt"
	"log"
	
	"github.com/gofiber/fiber/v2"
)

func HandleSingleRequest(ctx *fiber.Ctx) error {

	file, err := ctx.FormFile("cover")
	if err != nil {
		log.Println("error:", err)
	}

	var filename *string
	if file != nil {
		filename = &file.Filename

		errsave := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", *filename))
		if errsave != nil {
			log.Println("error:", errsave)
		}

	} else {
		log.Println("error:", "Nothing file to updload")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()

}


func HandleMultipleRequest(ctx *fiber.Ctx) error {
	form, errform := ctx.MultipartForm()
	

	if errform != nil {
		log.Println("error:", errform)
	}

	files := form.File["photos"]
	var filenames []string
	for i , file := range files{ 
		
		var filename string
	if file != nil {
		filename = fmt.Sprintf("%s-%d", file.Filename, i)

		errsave := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
		if errsave != nil {
			log.Println("error:", errsave)
		}

	} else {
		log.Println("error:", "Nothing file to updload")
	}

	if filename != "" {
		filenames = append(filenames, filename)
	} 

	

	

	}
	log.Println("filenames", filenames)
	ctx.Locals("filenames", filenames)
	return ctx.Next()


}
