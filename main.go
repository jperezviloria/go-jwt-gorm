package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jperezviloria/go-jwt-gorm/router"
)

func main() {

	app := fiber.New()

	//middleware
	app.Use(logger.New())

	//router
	router.UserRouter(app)
	router.AuthRouter(app)

	//Setting to listen server port
	app.Listen(":3000")

}
