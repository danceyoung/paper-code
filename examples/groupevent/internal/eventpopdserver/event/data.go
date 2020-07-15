package event

import (
	"paper-code/examples/groupevent/internal/pkg/db"
)

const insertaeventsql string = "INSERT INTO events(name,start_date,expired_on,member_count_limit,address,`desc`) VALUES (?,?,?,?,?,?)"
const insertajoinsql string = `INSERT INTO event_members (event_id, student_name, student_id, g_m, college, level, profession) VALUES (?,?,?,?,?,?,?);`
const selecteventsbystudentidsql string = "SELECT events.name,events.start_date,events.expired_on,events.member_count_limit,events.address,events.`desc`FROM event_members, events where events.id=event_members.event_id and student_id=?"

func newAEvent(name, startDate, expiredOn string, countLimit int, address, desc string) error {
	_, err := db.NewDB().Exec(insertaeventsql, name, startDate, expiredOn, countLimit, address, desc)
	if err != nil {
		return err
	}
	return nil
}

func joinAEvent(eventId string, name, gm, studentId, college, level, profession string) error {
	_, err := db.NewDB().Exec(insertajoinsql, eventId, name, studentId, gm, college, level, profession)
	if err != nil {
		return err
	}
	return nil
}

func eventsBy(studentId string) ([]map[string]interface{}, error) {
	rows, err := db.NewDB().Query(selecteventsbystudentidsql, studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []map[string]interface{}
	for rows.Next() {
		var (
			name, address, desc  string
			startDate, expiredOn string
			countLimit           int
		)
		err = rows.Scan(&name, &startDate, &expiredOn, &countLimit, &address, &desc)
		if err != nil {
			return nil, err
		}

		temp := make(map[string]interface{})
		temp["name"] = name
		temp["startDate"] = startDate
		temp["expiredOn"] = expiredOn
		temp["countLimit"] = countLimit
		temp["address"] = address
		temp["desc"] = desc
		result = append(result, temp)

	}

	return result, nil
}
