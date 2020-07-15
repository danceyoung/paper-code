package handler

import (
	"groupevent/internal/eventserver/biz/member"
	"log"
	"net/http"
)

type MemberHandler struct {
	BaseHandler
}

func (mh *MemberHandler) Members(rw http.ResponseWriter, req *http.Request) {
	log.Println(req.ParseForm())
	members, err := member.MembersBy(req.Form.Get("event-id"))
	mh.responseWith(rw, members, err)
}
