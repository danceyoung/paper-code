package updater

import (
	"log"
	"paper-code/examples/groupevent/internal/eventtimer"
)

func UpdateEventStatus() {
	log.Println("updating status of status...")
	err := eventtimer.UpdateEventStatus()
	if err != nil {
		panic(err)
	}
}
