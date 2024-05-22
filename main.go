package main

import (
	"managment/database"
	"managment/routes" // importing the routes package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.DBconn()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
		AllowOrigins:     "http://localhost:3000",
	}))

	routes.Setup(app) // A routes package/folder is created with 'Setup' function created.

	app.Listen(":8000")
}
