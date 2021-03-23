# -fiber_casbin_demo
> I use the vagrant env' to dev' golang project
### First CREATE DATABASE testing
```
CREATE DATABASE testing default character set utf8;
```

> NewAdapterByDBUseTableName >> casbinRule table names is ``casbin_rules``

### install mod vendor

### set dev.env config from your env setting，likes db config

### go run main.go

+ Open browser Input url one :  http://127.0.0.1:10183/v1/add [ AddPolicy admin can get this router ]
+ Open browser Input url two:  http://127.0.0.1:10183/v1/test [ test admin can be request this router `v1/test`]
+ Open browser Input url three:  http://127.0.0.1:10183/v1/add [ RemovePolicy admin get this router `v1/test` ]

when i open one of the step, the casbin_rules add one Policy to admin can request the router `v1/test`
and open the second of the stop, no request permission. no effect in time.
when i run main.go code again. the v1/test can be request .
when i open the third of the stop , delete the router policy , no effect. 
