package data

import "paper-code/example/groupevent/internal/pkg/db"

const insertaeventsql string = "INSERT INTO events(name,start_date,expired_on,member_count_limit,address,`desc`) VALUES (?,?,?,?,?,?)"
const insertajoinsql string = `INSERT INTO event_members (event_id, student_name, student_id, g_m, college, level, profession) VALUES (?,?,?,?,?,?,?);`
const selecteventsbystudentidsql string = "SELECT * FROM event_members where student_id=?"

func NewAEvent(name, startDate, expiredOn string, countLimit int, address, desc string) error {
	_, err := db.NewDB().Exec(insertaeventsql, name, startDate, expiredOn, countLimit, address, desc)
	if err != nil {
		return err
	}
	return nil
}

func JoinAEvent(eventId string, name, gm, studentId, college, level, profession string) error {
	_, err := db.NewDB().Exec(insertajoinsql, eventId, name, studentId, gm, college, level, profession)
	if err != nil {
		return err
	}
	return nil
}

func EventsBy(studentId string) error {
	rows, err := db.NewDB().Query(selecteventsbystudentidsql, studentId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
