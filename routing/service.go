package routing

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/qzich/todo/handlers"
)

const (
	ServerHeartbeatUri = "/heartbeat"
)

func ServiceRoutes(routerGroup *routing.RouteGroup) {
	routerGroup.Get(ServerHeartbeatUri, handlers.ServiceHeartbeatHandler)
}
