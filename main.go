package main

import (
	"github.com/dev-hyunsang/effective-gofiber-error-handling/cmd"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Middleware(app *fiber.App) {
	app.Post("/update", cmd.Create)
	app.Post("/read", cmd.AllToDoRead)
	app.Post("/read/:uuid", cmd.ParametersRead)
	app.Post("/update", cmd.Update)
	app.Delete("/delete/:uuid", cmd.Delete)
}

func main() {
	app := fiber.New()

	Middleware(app)

	if err := app.Listen(":3000"); err != nil {
		log.Println("Failed to Server Starting")
		log.Fatalln(err)
	}
}
