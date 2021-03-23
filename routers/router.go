package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
	"os"
)

func SetupRoutes(app *fiber.App) {
	// 静态资源 Static file server
	app.Static("/public", "./public")
	app.Use(cors.New(cors.Config{
		//	AllowOrigins:     "http://localhost:10083, http://localhost:10099", //
		AllowOrigins:     "*", //
		AllowMethods:     "GET,POST,OPTIONS,UPDATE,DELETE,PUT",         //
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           15,
	}))
	app.Use(logger.New(logger.Config{
		Next:       nil,
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
		Output:     os.Stderr,
	}))
	prefix := app.Group("")
	InitVRouter(prefix)
	//InitWebRouter(prefix)
	// 404 Handler
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Not Found"}) // => 404 "Not Found"
	})
}
