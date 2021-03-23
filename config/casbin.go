package config

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/svcg/-fiber_casbin_demo/app/models"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(DB, &models.CasbinRule{})
	Enforcer, _ = casbin.NewEnforcer("config/rbac_model.conf", adapter)
	Enforcer.EnableAutoSave(true)
}
