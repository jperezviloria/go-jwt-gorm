package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jperezviloria/go-jwt-gorm/handler"
)

func AuthRouter(app *fiber.App) {

	app.Post("/login", handler.Login)
}
