package handler

import (
	"log"
	"net/http"
	"paper-code/example/groupevent/internal/eventserver/biz/member"
)

type MemberHandler struct {
	BaseHandler
}

func (mh *MemberHandler) Members(rw http.ResponseWriter, req *http.Request) {
	log.Println(req.ParseForm())
	members, err := member.MembersBy(req.Form.Get("event-id"))
	mh.responseWith(rw, members, err)
}
