### fiber 使用 casbin 的代码示例

### [英文文档](readme.md)

### 依赖
> [fiber](https://github.com/gofiber/fiber)
> [fiber-casbin](https://github.com/arsmn/fiber-casbin)
> [gorm](https://gorm.io/)

> 本示例代码使用以上依赖库，git clone 仓库，或者直接下载 zip 文件 后，执行 go mod tidy && go mod vendor
> dev.env 为数据库配置文件，需要修改成您的配置
> 我使用 vagrant 作为开发环境，如果你使用 docker 需要指定数据库localhost

### 创建数据库
```
CREATE DATABASE testing default character set utf8;
```

### 目录结构 [类 laravel]
```
app			
	-controllers  控制器
	-middlewares  中间件
	-models       声明struct
config					配置
routers					路由
services				服务类
tests						测试类
vendor					依赖
```

### 使用
```
go run main.go	// listen: "localhost:10183" 
```
######  会自动创建所使用的表 casbin_rule

### Casbin 中间件
```
func Casbin() *fibercasbin.CasbinMiddleware {
	db := config.DB //全局数据库db_connection
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &models.CasbinRule{})
	authz := fibercasbin.New(fibercasbin.Config{
		Enforcer: config.Enforcer,	// 全局 casbin_enforcer 这个是我们定义的
		//Mode: fibercasbin.ModeEnforcer, // v2.71 已去掉这个参数
		ModelFilePath: "config/rbac_model.conf",	//读取的rbac配置
		PolicyAdapter: adapter,
	})
	return authz
}
```


> 示例代码提供了三个路由url
+ `/v1/add` 添加 一个 Policy
+ `/v1/remove` 移出指定的Policy
+ `/v1/test` 测试指定Policy

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
