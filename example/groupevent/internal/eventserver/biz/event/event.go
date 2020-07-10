package event

import (
	"errors"
	"paper-code/example/groupevent/internal/eventserver/data"

	"time"
)

//how errors are handled:

//1 NOT allowed to panic an application.
//2 Wrap errors with context if not being handled.
//3 Minority of handling errors happen here.

// how recovering panics:

//Can not recover from panics

type EventM struct {
	Name             string `json:"name"`
	StartDate        string `json:"start_date"`
	ExpiredOn        string `json:"expired_on"`
	MemberCountLimit int    `json:"member_count_limit"`
	Address          string `json:"address"`
	Desc             string `json:"desc"`
}

func NewAEvent(ev EventM) error {
	if _, err := time.Parse("2006-01-02", ev.StartDate); err != nil {
		return errors.New("date param is invalid")
	}
	if _, err := time.Parse("2006-01-02", ev.ExpiredOn); err != nil {
		return errors.New("date param is invalid")
	}

	if len(ev.Name) == 0 || len(ev.Address) == 0 || len(ev.Desc) == 0 || ev.MemberCountLimit == 0 {
		return errors.New("params are not enough")
	}

	if err := data.NewAEvent(ev.Name, ev.StartDate, ev.ExpiredOn, ev.MemberCountLimit, ev.Address, ev.Desc); err != nil {
		return errors.New("a error was accurred when new a event. " + err.Error())
	}

	return nil
}

func JoinAEvent(eventId string, name, gm, studentId, college, level, profession string) error {
	if len(eventId) == 0 || len(studentId) == 0 {
		return errors.New("params are not enough")
	}
	// if err := data.JoinAEvent(eventId, name, gm, studentId, college, level, profession); err != nil {
	// 	return errors.New("a error was accurred when join a event")
	// }
	return nil
}
func EventsBy(studentId string) {
	// data.EventsBy(studentId)
}
