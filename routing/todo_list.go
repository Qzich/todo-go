package routing

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/qzich/todo/handlers"
)

const (
	GetTodoListsUri   = "/lists"
	GetTodoListUri    = "/list/<id>"
	AddTodoListUri    = "/list"
	RemoveTodoListUri = "/list/<id>"
	UpdateTodoListUri = "/list/<id>"
)

func TodoListRoutes(routerGroup *routing.RouteGroup) {

	routerGroup.Post(AddTodoListUri, handlers.AddTodoListHandler)

	routerGroup.Get(GetTodoListUri, handlers.GetTodoListHandler)

	routerGroup.Get(GetTodoListsUri, handlers.GetTodoListsHandler)

	routerGroup.Delete(RemoveTodoListUri, handlers.RemoveTodoListHandler)

	routerGroup.Post(UpdateTodoListUri, handlers.AddTodoListHandler)
}
