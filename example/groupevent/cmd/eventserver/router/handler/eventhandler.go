package handler

import "net/http"

type EventHandler struct {
	BaseHandler
}

func (eventH *EventHandler) NewAEvent(rw http.ResponseWriter, req *http.Request) {

}

func (eventH *EventHandler) JoinAEvent(rw http.ResponseWriter, req *http.Request) {

}

func (eventH *EventHandler) Events(rw http.ResponseWriter, req *http.Request) {

}
