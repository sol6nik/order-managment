package routes

import (
	"managment/controllers" // importing the routes package

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	users := app.Group("/user")

	users.Get("/user", controllers.User)
	users.Post("/register", controllers.Register)
	users.Post("/login", controllers.Login)
	users.Post("/logout", controllers.Logout)

	order := app.Group("/order")
	order.Post("/orders", controllers.CreateOrder)
	order.Get("/orders", controllers.GetOrders)
	order.Get("/orders/:id", controllers.GetOrder)
	order.Put("/orders/:id", controllers.UpdateOrder)
	order.Delete("/orders/:id", controllers.DeleteOrder)
}
