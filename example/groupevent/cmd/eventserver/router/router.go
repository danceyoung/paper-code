package router

import (
	"net/http"
	"paper-code/example/groupevent/cmd/eventserver/router/handler"
	"paper-code/example/groupevent/pkg/middleware"
)

const prepath string = "/group-event-sericed"

func init() {
	http.Handle(prepath+"/events/new", middleware.HandlerConv(new(handler.EventHandler).NewAEvent))
	http.Handle(prepath+"/events/new", middleware.HandlerConv(new(handler.EventHandler).JoinAEvent))
	http.Handle(prepath+"/events/new", middleware.HandlerConv(new(handler.EventHandler).Events))
}
