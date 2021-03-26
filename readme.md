### fiber 使用 casbin 的代码示例

### [中文文档](readme_zh.md)

### vendor list
> [fiber](https://github.com/gofiber/fiber)
> [fiber-casbin](https://github.com/arsmn/fiber-casbin)
> [gorm](https://gorm.io/)

> You can clone this package or download zip code ,  then  you can run `go mod tidy && go mod vendor` to fix code vendor library.
> dev.env is database config file , you must to modify yours。
> i use the vagrant of the development environment，if you use docker, you must to modify database connection host.

### create database Or modify the dev.env `DB_DATABASE`
```
CREATE DATABASE testing default character set utf8;
```

### directory structure [likes laravel project]
```
app			
	-controllers  
	-middlewares  
	-models       
config					
routers					
services				
tests						
vendor					
```

### Usage
```
go run main.go	// listen: "localhost:10183" 
```
######  the command will be create the table `casbin_rule`

### Casbin Middleware
```
func Casbin() *fibercasbin.CasbinMiddleware {
	db := config.DB //global db_connection
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &models.CasbinRule{})
	authz := fibercasbin.New(fibercasbin.Config{
		Enforcer: config.Enforcer,	// global casbin_enforcer own define
		//Mode: fibercasbin.ModeEnforcer, // version2.71 already removed
		ModelFilePath: "config/rbac_model.conf",	// your rbac config filepath
		PolicyAdapter: adapter,
	})
	return authz
}
```


> The demo code provide three url of the request routers
+ `/v1/add` add a Policy
+ `/v1/remove` remove a Policy
+ `/v1/test` test the Policy

```
func Router(Router fiber.Router) {
	authz := middlewares.Casbin()
	v := Router.Group("/v1")
	{
		v.Get("/add", v1.Add)
		v.Get("/remove", v1.Remove)
		v.Get("/test", authz.RoutePermission(), v1.Test)
	}
}
```

