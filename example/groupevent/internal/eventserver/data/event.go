package data

import "paper-code/example/groupevent/internal/pkg/db"

const insertaeventsql string = "INSERT INTO events(name,start_date,expired_on,member_count_limit,address,desc) VALUES (?,?,?,?,?,?)"

func NewAEvent(name, startDate, expiredOn string, countLimit int, address, desc string) error {
	_, err := db.NewDB().Exec(insertaeventsql, name, startDate, expiredOn, countLimit, address, desc)
	if err != nil {
		return err
	}
	return nil
}

func JoinAEvent(eventId string, name, gm, studentId, college, level, profession string) error {
	return nil
}

func EventsBy(studentId string) {

}
