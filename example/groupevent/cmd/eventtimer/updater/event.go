package updater

import (
	"log"
	"paper-code/example/groupevent/internal/eventtimer"
)

func UpdateEventStatus() {
	log.Println("updating status of status...")
	eventtimer.UpdateEventStatus()
}
