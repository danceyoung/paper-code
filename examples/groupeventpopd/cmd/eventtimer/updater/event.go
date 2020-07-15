package updater

import (
	"groupeventpopd/internal/eventtimer"
	"log"
)

func UpdateEventStatus() {
	log.Println("updating status of status...")
	err := eventtimer.UpdateEventStatus()
	if err != nil {
		panic(err)
	}
}
