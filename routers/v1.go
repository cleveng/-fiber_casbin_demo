package routers

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/svcg/-fiber_casbin_demo/app/controllers/v1"
	"github.com/svcg/-fiber_casbin_demo/app/middlewares"
)

func InitVRouter(Router fiber.Router) {
	authz := middlewares.Casbin()
	v := Router.Group("/v1")
	{
		v.Get("/add", v1.Add)
		v.Get("/remove", v1.Remove)
		v.Get("/test", authz.RoutePermission(), v1.Test)
	}
}
