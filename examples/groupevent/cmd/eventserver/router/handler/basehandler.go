package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type BaseHandler struct{}

func (bh *BaseHandler) responseWith(rw http.ResponseWriter, resbody interface{}, err error) {
	if err != nil {
		temp := make(map[string]interface{})
		temp["code"] = -1
		temp["msg"] = err.Error()
		temp["data"] = nil
		tempbytes, err := json.Marshal(temp)
		if err != nil {
			panic(err)
		}
		io.WriteString(rw, string(tempbytes))
	} else {
		temp := make(map[string]interface{})
		temp["code"] = 0
		temp["msg"] = ""
		temp["data"] = resbody
		tempbytes, err := json.Marshal(temp)
		if err != nil {
			panic(err)
		}
		io.WriteString(rw, string(tempbytes))
	}
}
