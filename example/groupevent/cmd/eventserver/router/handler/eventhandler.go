package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"paper-code/example/groupevent/internal/eventserver/biz/event"
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

}

func (eventH *EventHandler) JoinAEvent(rw http.ResponseWriter, req *http.Request) {

}

func (eventH *EventHandler) Events(rw http.ResponseWriter, req *http.Request) {

}
