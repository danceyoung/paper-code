package main

import (
	"fmt"
	"groupevent/cmd/eventtimer/updater"
	"log"
	"time"
)

func main() {
	//execute updating when 00:00:00 everyday
	h, m, s := time.Now().Clock()
	firstDuration, _ := time.ParseDuration(fmt.Sprintf("%dh%dm%ds", 23-h, 60-m, 60-s))
	resetDuration, _ := time.ParseDuration(fmt.Sprintf("%dh%dm%ds", 24, 0, 0))
	timer := time.NewTimer(firstDuration)
	for {
		select {
		case <-timer.C:
			timer.Reset(resetDuration)
			log.Println("update: ", time.Now())
			updater.UpdateEventStatus()
		}
	}
}
