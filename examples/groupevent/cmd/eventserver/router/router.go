package router

import (
	"groupevent/cmd/eventserver/router/handler"
	"groupevent/pkg/middleware"
	"net/http"
)

const prepath string = "/group-event-sericed"

func init() {
	http.Handle(prepath+"/events/new", middleware.HandlerConv(new(handler.EventHandler).NewAEvent))
	http.Handle(prepath+"/event/join", middleware.HandlerConv(new(handler.EventHandler).JoinAEvent))
	http.Handle(prepath+"/events", middleware.HandlerConv(new(handler.EventHandler).Events))
	http.Handle(prepath+"/members", middleware.HandlerConv(new(handler.MemberHandler).Members))
}
