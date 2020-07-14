package member

import (
	"errors"
	"log"

	"paper-code/example/groupevent/internal/eventserver/data"
)

func MembersBy(eventid string) error {
	if len(eventid) == 0 {
		return errors.New("members by: params are not enough")
	}
	log.Println(data.MembersBy(eventid))

	return nil
}
