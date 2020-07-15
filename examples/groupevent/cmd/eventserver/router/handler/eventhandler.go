package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"paper-code/examples/groupevent/internal/eventserver/biz/event"
)

// How errors are handled:

//Allowed to panic an application.
//Wrap errors with context if not being handled.
//Majority of handling errors happen here.

type EventHandler struct {
	BaseHandler
}

func (eventH *EventHandler) NewAEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		panic(errors.New("method request is mismatched"))
	}

	bodybytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var dest event.EventM
	err = json.Unmarshal(bodybytes, &dest)
	if err != nil {
		panic(err)
	}

	err = event.NewAEvent(dest)
	if err != nil {
		panic(err)
	}

	eventH.responseWith(rw, nil, nil)
}

func (eventH *EventHandler) JoinAEvent(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		panic(err)
	}

	bodybytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var dest event.MemberM
	err = json.Unmarshal(bodybytes, &dest)
	if err != nil {
		panic(err)
	}

	err = event.JoinAEvent(req.Form.Get("event-id"), dest)
	if err != nil {
		panic(err)
	}

	eventH.responseWith(rw, nil, nil)
}

func (eventH *EventHandler) Events(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		panic(err)
	}
	events, err := event.EventsBy(req.Form.Get("student-id"))
	eventH.responseWith(rw, events, err)
}
