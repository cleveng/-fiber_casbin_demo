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
		ModelFilePath: "config/rbac_model.conf",
		PolicyAdapter: adapter,
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "用户尚未登录，请求非法...", "status_code": 404})
		},
		Forbidden: func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "用户权限不足...", "status_code": 404})
		},
		Lookup: func(c *fiber.Ctx) string {
			// get subject from BasicAuth, JWT, Cookie etc in real world
			return "admin"
		},
	})
	//adapter.LoadPolicy()
	return authz
}
