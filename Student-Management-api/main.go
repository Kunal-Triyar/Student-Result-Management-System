package main

import (
	"student-management/database"
	"student-management/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App){
	app.Get("/student/:regno",routes.GetStudent)
	app.Post("/student",routes.CreateStudent)
}

func main() {
	app := fiber.New()

	database.ConnectToDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

	SetupRoutes(app)

	app.Listen(":3000")

}