package main

import (
	"github.com/qzich/todo/server"
	"github.com/valyala/fasthttp"
	fasthttpRouting "github.com/qiangxue/fasthttp-routing"
	"github.com/qzich/todo/routing"
)

func main() {
	todoServer := server.New(":8080", "/api", fasthttpRouting.New())

	todoServer.ApplyRoutes(routing.TodoListRoutes)

	panic(todoServer.Run(fasthttp.ListenAndServe))
}
