# goZero demo 实验

> time: 2021.7.31

本样例基于 go-zero 微服务治理框架，使用　goctl 脚手架工具极速搭建了一个简单的 user 微服务模块，附带提供　blog.http 测试接口。


初始化基础代码

```bash
//初始化 api 模板
goctl api -o blog.api
//修改模板，生成　api 代码
goctl api go --api blog.api --dir .

//初始化 rpc 模板
goctl rpc template -o user.proto
//修改模板，生成 rpc 代码
goctl rpc proto --src user.proto --dir .

//编写sql 基于sql 生成CURD代码
goctl model mysql ddl -c -src user.sql -dir .
```

基于初始化代码，编写基本的业务逻辑,运行代码..

```bash
go mod tidy
go run users/api/blog.go -f users/api/etc/blog-api.yaml
go run users/rpc/user.go -f users/rpc/etc/user.yaml
```

最后基于 API 接口编写测试接口 blog.http.