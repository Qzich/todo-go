package routing

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/qzich/todo/handlers"
)

const (
	GetTodoListItemsUri   = "/list/<listId>/items"
	GetTodoListItemUri    = "/list/<listId>/item/<itemId>"
	AddTodoListItemUri    = "/list/<listId>/item"
	RemoveTodoListItemUri = "/list/<listId>/item/<itemId>"
	UpdateTodoListItemUri = "/list/<listId>/item/<itemId>"
)

func TodoListItemsRoutes(routerGroup *routing.RouteGroup) {

	routerGroup.Post(AddTodoListItemUri, handlers.AddTodoListItemHandler)

	routerGroup.Post(UpdateTodoListItemUri, handlers.AddTodoListItemHandler)

	routerGroup.Delete(RemoveTodoListItemUri, handlers.RemoveTodoListItemHandler)

	routerGroup.Get(GetTodoListItemUri, handlers.GetTodoListItemHandler)

	routerGroup.Get(GetTodoListItemsUri, handlers.GetTodoListItemsHandler)

}
