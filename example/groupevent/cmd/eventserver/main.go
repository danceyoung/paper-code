package main

import (
	"log"
	"net/http"
	_ "paper-code/example/groupevent/cmd/eventserver/router"
	_ "paper-code/example/groupevent/internal/pkg/cfg"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}
