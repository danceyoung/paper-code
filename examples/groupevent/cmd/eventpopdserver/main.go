package main

import (
	"log"
	"net/http"
	_ "paper-code/examples/groupevent/cmd/eventpopdserver/router"
	_ "paper-code/examples/groupevent/internal/pkg/cfg"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}
