package member

import (
	"errors"

	"groupevent/internal/eventserver/data"
)

func MembersBy(eventid string) ([]map[string]string, error) {
	if len(eventid) == 0 {
		return nil, errors.New("members by: params are not enough")
	}

	return data.MembersBy(eventid)
}
