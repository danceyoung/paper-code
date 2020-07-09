package main

import (
	"net/http"
	_ "paper-code/example/groupevent/cmd/eventserver/router"
)

func main() {
	http.ListenAndServe(":8080", nil)
}
