package main

import (
	"github.com/qzich/todo/server"
	"github.com/valyala/fasthttp"
	fasthttpRouting "github.com/qiangxue/fasthttp-routing"
	"github.com/qzich/todo/routing"
	"github.com/qzich/todo/validators"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.CustomTypeTagMap.Set("TodoValidColors", validators.ValidateTodoListColor)
}

func main() {

	todoServer := server.New(":8080", "/api", fasthttpRouting.New())

	todoServer.ApplyRoutes(routing.ServiceRoutes)

	todoServer.ApplyRoutes(routing.TodoListRoutes)

	todoServer.ApplyRoutes(routing.TodoListItemsRoutes)

	panic(todoServer.Run(fasthttp.ListenAndServe))
}
