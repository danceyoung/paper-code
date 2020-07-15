package router

import (
	"net/http"
	"paper-code/examples/groupevent/cmd/eventpopdserver/router/handler"

	"paper-code/examples/groupevent/pkg/middleware"
)

// "paper-code/examples/groupevent/pkg/middleware"

const prepath string = "/group-event-popd-serviced"

func init() {

	http.Handle(prepath+"/events/new", middleware.HandlerConv(new(handler.EventHandler).NewAEvent))
	http.Handle(prepath+"/event/join", middleware.HandlerConv(new(handler.EventHandler).JoinAEvent))
	http.Handle(prepath+"/events", middleware.HandlerConv(new(handler.EventHandler).Events))
	http.Handle(prepath+"/members", middleware.HandlerConv(new(handler.MemberHandler).Members))
}
