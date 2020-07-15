package main

import (
	_ "groupevent/cmd/eventserver/router"
	_ "groupevent/internal/pkg/cfg"
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}
