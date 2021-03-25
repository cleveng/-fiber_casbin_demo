package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svcg/-fiber_casbin_demo/config"
	"net/http"
)

func Add(ctx *fiber.Ctx) error {
	enforcer := config.Enforcer
	//authz := middlewares.Casbin()
	//authz.Enforcer.LoadPolicy()
	if ok, _ := 	enforcer.AddPolicy("admin", "/v1/test", "GET"); !ok {
		return ctx.Status(http.StatusOK).JSON("Policy exist")
	}
	return ctx.Status(http.StatusOK).JSON("add success")
}

func Remove(ctx *fiber.Ctx) error {
	enforcer := config.Enforcer
	//authz := middlewares.Casbin()
	//authz.Enforcer.LoadPolicy()
	if ok, _ := enforcer.RemovePolicy("admin", "/v1/test", "GET"); !ok {
		return ctx.Status(http.StatusOK).JSON("Policy no exist")
	}
	return ctx.Status(http.StatusOK).JSON("delete success")
}

func Test(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("I am a good man.")
}
