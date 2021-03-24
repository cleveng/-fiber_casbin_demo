package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/svcg/-fiber_casbin_demo/config"
	"github.com/svcg/-fiber_casbin_demo/routers"
	"github.com/svcg/-fiber_casbin_demo/services"
	"log"
	"os"
)

var r *fiber.App


func main() {
	config.ConnectDB() // 初始化数据库
	//config.InitCasbin()
	services.InitTables()

	app := setupHTTP()
	log.Fatal(app.Listen("localhost:10183")) // Start server
}

func setupHTTP() *fiber.App {
	r = fiber.New()
	r.Use(logger.New(logger.Config{
		Next:       nil,
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
		Output:     os.Stderr,
	}))
	r.Use(cors.New(cors.Config{
		//AllowOrigins:     "http://localhost:10083, http://localhost:10099", //
		AllowOrigins:     "*",                                  //
		AllowMethods:     "GET,POST,OPTIONS,UPDATE,DELETE,PUT", //
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           15,
	}))
	routers.SetupRoutes(r)
	return r
}
