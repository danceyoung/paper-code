package main

import (
	_ "groupeventpopd/cmd/eventserver/router"
	_ "groupeventpopd/internal/pkg/cfg"
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}
