package middlewares

import (
	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/svcg/-fiber_casbin_demo/app/models"
	"github.com/svcg/-fiber_casbin_demo/config"
	"net/http"
)

func Casbin() *fibercasbin.CasbinMiddleware {
	db := config.DB
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &models.CasbinRule{})
	authz := fibercasbin.New(fibercasbin.Config{
		Enforcer: config.Enforcer,
		//Mode: fibercasbin.ModeEnforcer, // v2.71 去掉这个参数
		ModelFilePath: "config/rbac_model.conf",
		PolicyAdapter: adapter,
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized...", "status_code": 401})
		},
		Forbidden: func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Forbidden ...", "status_code": 404})
		},
		Lookup: func(c *fiber.Ctx) string {
			// get subject from BasicAuth, JWT, Cookie etc in real world
			return "admin"
		},
	})
	return authz
}
