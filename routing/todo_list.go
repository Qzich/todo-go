package routing

import (
	"github.com/qiangxue/fasthttp-routing"
)

func TodoListRoutes(routerGroup *routing.RouteGroup) {
	routerGroup.Get("/list/heartbeat", func(context *routing.Context) error {

		return nil
	})
}
