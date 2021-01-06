package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jperezviloria/go-jwt-gorm/handler"
)

func UserRouter(app *fiber.App) {

	app.Get("/user", handler.GetUser)
	app.Post("/user", handler.CreateUser)
	app.Get("/user/all", handler.GetAllUsers)
}
