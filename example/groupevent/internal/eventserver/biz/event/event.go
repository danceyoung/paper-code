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

func NewAEvent(name, startDate, expiredOn string, countLimit int, address, desc string) error {
	if _, err := time.Parse("2006-01-02", startDate); err != nil {
		return errors.New("date param is invalid")
	}
	if _, err := time.Parse("2006-01-02", expiredOn); err != nil {
		return errors.New("date param is invalid")
	}

	if len(name) == 0 || len(address) == 0 || len(desc) == 0 || countLimit == 0 {
		return errors.New("params are not enough")
	}

	if err := data.NewAEvent(name, startDate, expiredOn, countLimit, address, desc); err != nil {
		return errors.New("a error was accurred when new a event")
	}

	return nil
}

func JoinAEvent(eventId string, name, gm, studentId, college, level, profession string) error {
	if len(eventId) == 0 || len(studentId) == 0 {
		return errors.New("params are not enough")
	}
	if err := data.JoinAEvent(eventId, name, gm, studentId, college, level, profession); err != nil {
		return errors.New("a error was accurred when join a event")
	}
	return nil
}
func EventsBy(studentId string) {
	data.EventsBy(studentId)
}
