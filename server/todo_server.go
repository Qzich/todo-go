package server

import (
	"github.com/valyala/fasthttp"
	"github.com/qiangxue/fasthttp-routing"
)

type TodoServer struct {
	addressToListen string
	routerGroup     *routing.RouteGroup
	router          *routing.Router
}

type ListenAndHandle func(addr string, handler fasthttp.RequestHandler) error

type RouteConfigure func(*routing.RouteGroup)

func New(address string, routerGroupName string, router *routing.Router) *TodoServer {
	routerGroup := router.Group(routerGroupName)

	return &TodoServer{
		addressToListen: address,
		routerGroup:     routerGroup,
		router:          router,
	}
}

func (this *TodoServer) Run(listenAndHandle ListenAndHandle) error {

	return listenAndHandle(this.addressToListen, this.router.HandleRequest)
}

func (this *TodoServer) ApplyRoutes(routeConfigure RouteConfigure) *TodoServer {
	routeConfigure(this.routerGroup)

	return this
}
